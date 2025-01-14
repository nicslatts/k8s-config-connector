// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FeaturestoreEncryptionSpec struct {
	/* The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created. */
	KmsKeyName string `json:"kmsKeyName"`
}

type FeaturestoreOnlineServingConfig struct {
	/* The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating. */
	// +optional
	FixedNodeCount *int `json:"fixedNodeCount,omitempty"`

	/* Online serving scaling configuration. Only one of fixedNodeCount and scaling can be set. Setting one will reset the other. */
	// +optional
	Scaling *FeaturestoreScaling `json:"scaling,omitempty"`
}

type FeaturestoreScaling struct {
	/* The maximum number of nodes to scale up to. Must be greater than minNodeCount, and less than or equal to 10 times of 'minNodeCount'. */
	MaxNodeCount int `json:"maxNodeCount"`

	/* The minimum number of nodes to scale down to. Must be greater than or equal to 1. */
	MinNodeCount int `json:"minNodeCount"`
}

type VertexAIFeaturestoreSpec struct {
	/* If set, both of the online and offline data storage will be secured by this key. */
	// +optional
	EncryptionSpec *FeaturestoreEncryptionSpec `json:"encryptionSpec,omitempty"`

	/* If set to true, any EntityTypes and Features for this Featurestore will also be deleted. */
	// +optional
	ForceDestroy *bool `json:"forceDestroy,omitempty"`

	/* Config for online serving resources. */
	// +optional
	OnlineServingConfig *FeaturestoreOnlineServingConfig `json:"onlineServingConfig,omitempty"`

	/* TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a featurestore. If not set, default to 4000 days. */
	// +optional
	OnlineStorageTtlDays *int `json:"onlineStorageTtlDays,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The region of the dataset. eg us-central1. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type VertexAIFeaturestoreStatus struct {
	/* Conditions represent the latest available observations of the
	   VertexAIFeaturestore's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Used to perform consistent read-modify-write updates. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaifeaturestore;gcpvertexaifeaturestores
// +kubebuilder:subresource:status

// VertexAIFeaturestore is the Schema for the vertexai API
// +k8s:openapi-gen=true
type VertexAIFeaturestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIFeaturestoreSpec   `json:"spec,omitempty"`
	Status VertexAIFeaturestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VertexAIFeaturestoreList contains a list of VertexAIFeaturestore
type VertexAIFeaturestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIFeaturestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIFeaturestore{}, &VertexAIFeaturestoreList{})
}
