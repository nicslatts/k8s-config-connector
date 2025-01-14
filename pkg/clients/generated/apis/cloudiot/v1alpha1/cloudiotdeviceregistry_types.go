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

type DeviceregistryCredentials struct {
	/* A public key certificate format and data. */
	PublicKeyCertificate DeviceregistryPublicKeyCertificate `json:"publicKeyCertificate"`
}

type DeviceregistryEventNotificationConfigs struct {
	/* PubSub topic name to publish device events. */
	PubsubTopicName string `json:"pubsubTopicName"`

	/* If the subfolder name matches this string exactly, this
	configuration will be used. The string must not include the
	leading '/' character. If empty, all strings are matched. Empty
	value can only be used for the last 'event_notification_configs'
	item. */
	// +optional
	SubfolderMatches *string `json:"subfolderMatches,omitempty"`
}

type DeviceregistryHttpConfig struct {
}

type DeviceregistryMqttConfig struct {
}

type DeviceregistryPublicKeyCertificate struct {
}

type DeviceregistryStateNotificationConfig struct {
}

type CloudIOTDeviceRegistrySpec struct {
	/* List of public key certificates to authenticate devices. */
	// +optional
	Credentials []DeviceregistryCredentials `json:"credentials,omitempty"`

	/* List of configurations for event notifications, such as PubSub topics
	to publish device events to. */
	// +optional
	EventNotificationConfigs []DeviceregistryEventNotificationConfigs `json:"eventNotificationConfigs,omitempty"`

	/* Activate or deactivate HTTP. */
	// +optional
	HttpConfig *DeviceregistryHttpConfig `json:"httpConfig,omitempty"`

	/* The default logging verbosity for activity from devices in this
	registry. Specifies which events should be written to logs. For
	example, if the LogLevel is ERROR, only events that terminate in
	errors will be logged. LogLevel is inclusive; enabling INFO logging
	will also enable ERROR logging. Default value: "NONE" Possible values: ["NONE", "ERROR", "INFO", "DEBUG"]. */
	// +optional
	LogLevel *string `json:"logLevel,omitempty"`

	/* Activate or deactivate MQTT. */
	// +optional
	MqttConfig *DeviceregistryMqttConfig `json:"mqttConfig,omitempty"`

	/* Immutable. */
	// +optional
	Project *string `json:"project,omitempty"`

	/* Immutable. The region in which the created registry should reside.
	If it is not provided, the provider region is used. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* A PubSub topic to publish device state updates. */
	// +optional
	StateNotificationConfig *DeviceregistryStateNotificationConfig `json:"stateNotificationConfig,omitempty"`
}

type CloudIOTDeviceRegistryStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudIOTDeviceRegistry's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudiotdeviceregistry;gcpcloudiotdeviceregistries
// +kubebuilder:subresource:status

// CloudIOTDeviceRegistry is the Schema for the cloudiot API
// +k8s:openapi-gen=true
type CloudIOTDeviceRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIOTDeviceRegistrySpec   `json:"spec,omitempty"`
	Status CloudIOTDeviceRegistryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudIOTDeviceRegistryList contains a list of CloudIOTDeviceRegistry
type CloudIOTDeviceRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIOTDeviceRegistry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudIOTDeviceRegistry{}, &CloudIOTDeviceRegistryList{})
}
