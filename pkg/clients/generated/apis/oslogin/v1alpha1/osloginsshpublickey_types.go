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

type OSLoginSSHPublicKeySpec struct {
	/* An expiration time in microseconds since epoch. */
	// +optional
	ExpirationTimeUsec *string `json:"expirationTimeUsec,omitempty"`

	/* Immutable. Public key text in SSH format, defined by RFC4253 section 6.6. */
	Key string `json:"key"`

	/* Immutable. The project ID of the Google Cloud Platform project. */
	// +optional
	Project *string `json:"project,omitempty"`

	/* Immutable. Optional. The service-generated fingerprint of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The user email. */
	User string `json:"user"`
}

type OSLoginSSHPublicKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   OSLoginSSHPublicKey's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The SHA-256 fingerprint of the SSH public key. */
	// +optional
	Fingerprint *string `json:"fingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcposloginsshpublickey;gcposloginsshpublickeys
// +kubebuilder:subresource:status

// OSLoginSSHPublicKey is the Schema for the oslogin API
// +k8s:openapi-gen=true
type OSLoginSSHPublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSLoginSSHPublicKeySpec   `json:"spec,omitempty"`
	Status OSLoginSSHPublicKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSLoginSSHPublicKeyList contains a list of OSLoginSSHPublicKey
type OSLoginSSHPublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OSLoginSSHPublicKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OSLoginSSHPublicKey{}, &OSLoginSSHPublicKeyList{})
}
