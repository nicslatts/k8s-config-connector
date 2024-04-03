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

package logging

import (
	"context"
	"fmt"

	api "google.golang.org/api/logging/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

const ctrlName = "logmetric-controller"

// LogMetric from api "google.golang.org/api/logging/v2":
// // LogMetric: Describes a logs-based metric. The value of the metric is
// // the number of log entries that match a logs filter in a given time
// // interval.Logs-based metrics can also be used to extract values from
// // logs and create a distribution of the values. The distribution
// // records the statistics of the extracted values along with an optional
// // histogram of the values as specified by the bucket options.
// type LogMetric struct {
// 	// BucketName: Optional. The resource name of the Log Bucket that owns
// 	// the Log Metric. Only Log Buckets in projects are supported. The
// 	// bucket has to be in the same project as the metric.For
// 	// example:projects/my-project/locations/global/buckets/my-bucketIf
// 	// empty, then the Log Metric is considered a non-Bucket Log Metric.
// 	BucketName string `json:"bucketName,omitempty"`

// 	// BucketOptions: Optional. The bucket_options are required when the
// 	// logs-based metric is using a DISTRIBUTION value type and it describes
// 	// the bucket boundaries used to create a histogram of the extracted
// 	// values.
// 	BucketOptions *BucketOptions `json:"bucketOptions,omitempty"`

// 	// CreateTime: Output only. The creation timestamp of the metric.This
// 	// field may not be present for older metrics.
// 	CreateTime string `json:"createTime,omitempty"`

// 	// Description: Optional. A description of this metric, which is used in
// 	// documentation. The maximum length of the description is 8000
// 	// characters.
// 	Description string `json:"description,omitempty"`

// 	// Disabled: Optional. If set to True, then this metric is disabled and
// 	// it does not generate any points.
// 	Disabled bool `json:"disabled,omitempty"`

// 	// Filter: Required. An advanced logs filter
// 	// (https://cloud.google.com/logging/docs/view/advanced_filters) which
// 	// is used to match log entries. Example: "resource.type=gae_app AND
// 	// severity>=ERROR" The maximum length of the filter is 20000
// 	// characters.
// 	Filter string `json:"filter,omitempty"`

// 	// LabelExtractors: Optional. A map from a label key string to an
// 	// extractor expression which is used to extract data from a log entry
// 	// field and assign as the label value. Each label key specified in the
// 	// LabelDescriptor must have an associated extractor expression in this
// 	// map. The syntax of the extractor expression is the same as for the
// 	// value_extractor field.The extracted value is converted to the type
// 	// defined in the label descriptor. If either the extraction or the type
// 	// conversion fails, the label will have a default value. The default
// 	// value for a string label is an empty string, for an integer label its
// 	// 0, and for a boolean label its false.Note that there are upper bounds
// 	// on the maximum number of labels and the number of active time series
// 	// that are allowed in a project.
// 	LabelExtractors map[string]string `json:"labelExtractors,omitempty"`

// 	// MetricDescriptor: Optional. The metric descriptor associated with the
// 	// logs-based metric. If unspecified, it uses a default metric
// 	// descriptor with a DELTA metric kind, INT64 value type, with no labels
// 	// and a unit of "1". Such a metric counts the number of log entries
// 	// matching the filter expression.The name, type, and description fields
// 	// in the metric_descriptor are output only, and is constructed using
// 	// the name and description field in the LogMetric.To create a
// 	// logs-based metric that records a distribution of log values, a DELTA
// 	// metric kind with a DISTRIBUTION value type must be used along with a
// 	// value_extractor expression in the LogMetric.Each label in the metric
// 	// descriptor must have a matching label name as the key and an
// 	// extractor expression as the value in the label_extractors map.The
// 	// metric_kind and value_type fields in the metric_descriptor cannot be
// 	// updated once initially configured. New labels can be added in the
// 	// metric_descriptor, but existing labels cannot be modified except for
// 	// their description.
// 	MetricDescriptor *MetricDescriptor `json:"metricDescriptor,omitempty"`

// 	// Name: Required. The client-assigned metric identifier. Examples:
// 	// "error_count", "nginx/requests".Metric identifiers are limited to 100
// 	// characters and can include only the following characters: A-Z, a-z,
// 	// 0-9, and the special characters _-.,+!*',()%/. The forward-slash
// 	// character (/) denotes a hierarchy of name pieces, and it cannot be
// 	// the first character of the name.This field is the [METRIC_ID] part of
// 	// a metric resource name in the format
// 	// "projects/PROJECT_ID/metrics/METRIC_ID". Example: If the resource
// 	// name of a metric is "projects/my-project/metrics/nginx%2Frequests",
// 	// this field's value is "nginx/requests".
// 	Name string `json:"name,omitempty"`

// 	// UpdateTime: Output only. The last update timestamp of the metric.This
// 	// field may not be present for older metrics.
// 	UpdateTime string `json:"updateTime,omitempty"`

// 	// ValueExtractor: Optional. A value_extractor is required when using a
// 	// distribution logs-based metric to extract the values to record from a
// 	// log entry. Two functions are supported for value extraction:
// 	// EXTRACT(field) or REGEXP_EXTRACT(field, regex). The arguments are:
// 	// field: The name of the log entry field from which the value is to be
// 	// extracted. regex: A regular expression using the Google RE2 syntax
// 	// (https://github.com/google/re2/wiki/Syntax) with a single capture
// 	// group to extract data from the specified log entry field. The value
// 	// of the field is converted to a string before applying the regex. It
// 	// is an error to specify a regex that does not include exactly one
// 	// capture group.The result of the extraction must be convertible to a
// 	// double type, as the distribution always records double values. If
// 	// either the extraction or the conversion to double fails, then those
// 	// values are not recorded in the distribution.Example:
// 	// REGEXP_EXTRACT(jsonPayload.request, ".*quantity=(\d+).*")
// 	ValueExtractor string `json:"valueExtractor,omitempty"`

// 	// Version: Deprecated. The API version that created or updated this
// 	// metric. The v2 format is used by default and cannot be changed.
// 	//
// 	// Possible values:
// 	//   "V2" - Logging API v2.
// 	//   "V1" - Logging API v1.
// 	Version string `json:"version,omitempty"`

// 	// ServerResponse contains the HTTP response code and headers from the
// 	// server.
// 	googleapi.ServerResponse `json:"-"`

// 	// ForceSendFields is a list of field names (e.g. "BucketName") to
// 	// unconditionally include in API requests. By default, fields with
// 	// empty or default values are omitted from API requests. However, any
// 	// non-pointer, non-interface field appearing in ForceSendFields will be
// 	// sent to the server regardless of whether the field is empty or not.
// 	// This may be used to include empty fields in Patch requests.
// 	ForceSendFields []string `json:"-"`

// 	// NullFields is a list of field names (e.g. "BucketName") to include in
// 	// API requests with the JSON null value. By default, fields with empty
// 	// values are omitted from API requests. However, any field with an
// 	// empty value appearing in NullFields will be sent to the server as
// 	// null. It is an error if a field in this list has a non-empty value.
// 	// This may be used to include null fields in Patch requests.
// 	NullFields []string `json:"-"`
// }

// AddLogMetricController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddLogMetricController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.LoggingLogMetricGVK

	// todo(acpana): plumb context throughout direct
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &logMetricModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m)
}

type logMetricModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &logMetricModel{}

type logMetricAdapter struct {
	resourceID string
	parentID   string

	desired         *krm.LoggingLogMetric
	actual          *api.LogMetric
	logMetricClient *api.ProjectsMetricsService
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &logMetricAdapter{}

// AdapterForObject implements the Model interface.
func (m *logMetricModel) AdapterForObject(ctx context.Context, client client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	projectMetricsService, err := m.newProjectMetricsService(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.LoggingLogMetric{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// todo acpana: any other validations?
	if obj.Spec.ProjectRef.External == "" {
		// todo acpana: resolve the external ref
		return nil, fmt.Errorf("project external reference is not set")
	}

	return &logMetricAdapter{
		resourceID:      ValueOf(obj.Spec.ResourceID),
		parentID:        obj.Spec.ProjectRef.External,
		desired:         obj,
		logMetricClient: projectMetricsService,
	}, nil
}

func (a *logMetricAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	logMetric, err := a.logMetricClient.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting logMetric %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = logMetric

	return true, nil
}

// Delete implements the Adapter interface.
func (a *logMetricAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deletd
	if a.resourceID == "" {
		return false, nil
	}

	_, err := a.logMetricClient.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting log metric %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

func (a *logMetricAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating object", "u", u)

	project := a.desired.Spec.ProjectRef.External
	// todo acpana this looks like a good candidate for factored out validation;
	// a shared validator? validate exists/ check set? validate is well formed?
	if project == "" {
		return fmt.Errorf("project is empty")
	}
	name := a.desired.GetName()
	if name == "" {
		return fmt.Errorf("name is empty")
	}
	lm := &api.LogMetric{
		Name: MakeFQN(project, name),
	}

	req := a.logMetricClient.Create(fmt.Sprintf("projects/%s", project), lm)
	log.Info("creating logMetric", "request", req)
	created, err := req.Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("logMetric %s creation failed: %w", lm.Name, err)
	}

	log.V(2).Info("created logMetric", "logMetric", created)

	resourceID := created.Name
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	status := &krm.LoggingLogMetricStatus{}
	if err := logMetricStatusToKRM(created, status); err != nil {
		return err
	}

	return setStatus(u, status)
}

func logMetricStatusToKRM(in *api.LogMetric, out *krm.LoggingLogMetricStatus) error {
	out.CreateTime = &in.CreateTime
	out.UpdateTime = &in.UpdateTime
	return nil
}

func (a *logMetricAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	update := &api.LogMetric{}
	update.Name = a.fullyQualifiedName()

	// todo: acpana do the rest of the field updates
	if ValueOf(a.desired.Spec.Description) != a.actual.Description {
		update.Description = ValueOf(a.desired.Spec.Description)
	}

	// DANGER: this is an upsert; it will create the LogMetric if it doesn't exists
	// but this behavior is consistent with the Terraform backed behavior we provide for this resource.
	// todo acpana: look for / switch to a better method and/or use etags etc
	_, err := a.logMetricClient.Update(a.fullyQualifiedName(), update).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("logMetric update failed: %w", err)
	}

	return nil
}

func (a *logMetricAdapter) fullyQualifiedName() string {
	return MakeFQN(a.parentID, a.resourceID)
}

// MakeFQN constructions a fully qualified name for a LogMetric resources
// to be used in API calls. The format expected is: "projects/[PROJECT_ID]/metrics/[METRIC_ID]".
// Func assumes values are well formed and validated.
func MakeFQN(projectID, metricID string) string {
	return fmt.Sprintf("projects/%s/metrics/%s", projectID, metricID)
}
