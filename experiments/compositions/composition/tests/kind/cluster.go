// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kind

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	kstatus "sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

const (
	ReuseIfPresent    bool = true
	RecreateIfPresent bool = false
)

var (
	scheme = runtime.NewScheme()
)

type kindCluster struct {
	name          string
	config        *rest.Config
	manifestPaths []string
	images        []string
	deployments   []types.NamespacedName
	ctx           context.Context
	client.Client
}

type KindClusterUser interface {
	Config() *rest.Config
	Name() string
	RestartWorkloads() error
	WaitForWorkloads() error
}

type KindCluster interface {
	ClusterUp() error
	Delete() error
	Exists() (bool, error)

	KindClusterUser
}

type KindClusterSet struct {
	available sync.Map
}

var kindClusterSet KindClusterSet

// NewKindCluster - return a cluster setup object
func NewKindCluster(name string, images []string, manifestPaths []string, deployments []types.NamespacedName) KindCluster {
	return &kindCluster{
		name:          name,
		manifestPaths: manifestPaths,
		images:        images,
		deployments:   deployments,
		ctx:           context.Background(),
	}
}

func ReserveCluster(t *testing.T) KindClusterUser {
	var cluster KindClusterUser
	found := false
	for !found {
		kindClusterSet.available.Range(func(k, v any) bool {
			v, loaded := kindClusterSet.available.LoadAndDelete(k)
			if !loaded {
				return true
			}
			found = true
			cluster = v.(KindClusterUser)
			return false
		})

		if found {
			continue
		}

		t.Logf("\nWaiting for a cluster to become available...")
		time.Sleep(5 * time.Second)
	}

	t.Logf("Reserved cluster %s", cluster.Name())
	return cluster
}

func ReleaseCluster(t *testing.T, cluster KindClusterUser) {
	kindClusterSet.available.Store(cluster.Name(), cluster)
	t.Logf("Released cluster %s", cluster.Name())
}

func VerifyKindIsInstalled() error {
	_, err := exec.LookPath("kind")
	if err != nil {
		return err
	}
	return nil
}

// Wait for all clusters to become ready
func (c *kindCluster) create() error {
	err := c.Delete()
	if err != nil {
		return err
	}
	clusterConfig, err := c.kindClusterDefinition()
	if err != nil {
		return err
	}
	defer os.Remove(clusterConfig)

	c.config, err = c.createCluster(clusterConfig)
	if err != nil {
		return err
	}

	c.Client, err = client.New(c.Config(), client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	return nil
}

func (c *kindCluster) registerImages() error {
	for _, image := range c.images {
		err := c.LoadImage(image)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *kindCluster) installManifests() error {
	for _, path := range c.manifestPaths {
		manifests, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		objects, err := manifest.ParseObjects(context.Background(), string(manifests))
		for _, item := range objects.Items {
			err := c.Client.Create(context.Background(), item.UnstructuredObject())
			if err != nil {
				exists := apierrors.IsAlreadyExists(err)
				if exists {
					continue
				}
				return err
			}
		}
	}
	return nil
}

// isReady - is the object ready
func isReady(ctx context.Context, c client.Client, u *unstructured.Unstructured) (bool, error) {
	key := types.NamespacedName{
		Name:      u.GetName(),
		Namespace: u.GetNamespace(),
	}
	err := c.Get(ctx, key, u)
	result := &kstatus.Result{}
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return false, err
		}
		return false, nil
	} else {
		result, err = kstatus.Compute(u)
		if err != nil {
			return false, err
		}
	}
	if result.Status != kstatus.CurrentStatus {
		return false, nil
	}
	return true, nil
}

func isDeploymentReady(ctx context.Context, c client.Client, nn types.NamespacedName) (bool, error) {
	u := unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"})
	u.SetName(nn.Name)
	u.SetNamespace(nn.Namespace)

	return isReady(ctx, c, &u)
}

func (c *kindCluster) WaitForWorkloads() error {
	start := time.Now()
	for true {
		allReady := true
		for _, workload := range c.deployments {
			ready, err := isDeploymentReady(c.ctx, c.Client, workload)
			if err != nil {
				continue
			}
			if !ready {
				allReady = false
				break
			}
		}
		if allReady {
			return nil
		}
		if time.Since(start).Seconds() > 40 {
			return fmt.Errorf("timed out waiting for operator to be ready")
		}
		time.Sleep(2)
	}
	return nil
}

func (c *kindCluster) RestartWorkloads() error {
	return nil
}

// ClusterUp: Create() + registerImages() + installManifests() + WaitForWorkloads()
func (c *kindCluster) ClusterUp() error {
	err := c.create()
	if err != nil {
		return fmt.Errorf("Error Creating Cluster. err: %v", err)
	}

	err = c.registerImages()
	if err != nil {
		return fmt.Errorf("Error Registering Images. err: %v", err)
	}

	err = c.installManifests()
	if err != nil {
		return fmt.Errorf("Error Installing Manifests. err: %v", err)
	}

	err = c.WaitForWorkloads()
	if err != nil {
		return fmt.Errorf("Error Waiting for Deloyments. err: %v", err)
	}

	kindClusterSet.available.Store(c.Name(), c)
	return nil
}

// Config return rest.Config
func (c *kindCluster) Config() *rest.Config {
	return c.config
}

// Name return name
func (c *kindCluster) Name() string {
	return c.name
}

func (c *kindCluster) String() string {
	return c.name
}

// LoadImage loads a docker image into the cluster
func (c *kindCluster) LoadImage(image string) error {
	err := exec.Command("kind", "load", "docker-image", image, "--name", c.name).Run()
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes all clusters
func (c *kindCluster) Delete() error {
	kindClusterSet.available.Delete(c.Name())
	err := exec.Command("kind", "delete", "cluster", "--name", c.name).Run()
	if err != nil {
		return err
	}
	return nil
}

// Exists checks if the cluster exists
func (c *kindCluster) Exists() (bool, error) {
	output, err := exec.Command("kind", "get", "clusters").CombinedOutput()
	if err != nil {
		return false, err
	}

	for _, cluster := range strings.Split(string(output), "\n") {
		if cluster == c.name {
			return true, nil
		}
	}
	return false, nil
}

func (c *kindCluster) kindClusterDefinition() (string, error) {
	ipAddress, err := c.getHostIPAddress()
	if err != nil {
		return "", err
	}
	clusterConfigFile, err := os.CreateTemp("", "kind-cluster.yaml")
	if err != nil {
		return "", err
	}
	defer clusterConfigFile.Close()
	kindClusterConfig := `kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  # Allow connections to the API Sever with the host IP address
  apiServerAddress: "` + ipAddress + `"`
	bytes := []byte(kindClusterConfig)
	_, err = clusterConfigFile.Write(bytes)
	return clusterConfigFile.Name(), err
}

func (c *kindCluster) createCluster(clusterConfig string) (*rest.Config, error) {
	_, err := exec.Command("kind", "create", "cluster", "--name", c.name, "--config", clusterConfig).CombinedOutput()
	if err != nil {
		return nil, err
	}

	kubeConfigFile, err := os.CreateTemp("", "kubeconfig.yaml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(kubeConfigFile.Name())
	content, err := exec.Command("kind", "get", "kubeconfig", "--name", c.name).CombinedOutput()
	if err != nil {
		return nil, err
	}
	bytes := []byte(content)
	_, err = kubeConfigFile.Write(bytes)
	kubeConfigFile.Close()

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigFile.Name()},
		&clientcmd.ConfigOverrides{
			ClusterInfo: clientcmdapi.Cluster{
				Server: "",
			},
			CurrentContext: "",
		}).ClientConfig()
}

func (c *kindCluster) getHostIPAddress() (string, error) {
	// Try getting host ip by creating a connection object and reading the localaddr
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
}
