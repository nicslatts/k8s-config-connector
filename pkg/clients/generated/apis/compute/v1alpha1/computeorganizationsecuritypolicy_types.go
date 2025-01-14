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

type ComputeOrganizationSecurityPolicySpec struct {
	/* A textual description for the organization security policy. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. A textual name of the security policy. */
	DisplayName string `json:"displayName"`

	/* Immutable. The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
	Format: organizations/{organization_id} or folders/{folder_id}. */
	Parent string `json:"parent"`

	/* Immutable. Optional. The policyId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The type indicates the intended use of the security policy.
	For organization security policies, the only supported type
	is "FIREWALL". Default value: "FIREWALL" Possible values: ["FIREWALL"]. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type ComputeOrganizationSecurityPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeOrganizationSecurityPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Fingerprint of this resource. This field is used internally during
	updates of this resource. */
	// +optional
	Fingerprint *string `json:"fingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The unique identifier for the resource. This identifier is defined by the server. */
	// +optional
	PolicyId *string `json:"policyId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeorganizationsecuritypolicy;gcpcomputeorganizationsecuritypolicies
// +kubebuilder:subresource:status

// ComputeOrganizationSecurityPolicy is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeOrganizationSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeOrganizationSecurityPolicySpec   `json:"spec,omitempty"`
	Status ComputeOrganizationSecurityPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeOrganizationSecurityPolicyList contains a list of ComputeOrganizationSecurityPolicy
type ComputeOrganizationSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeOrganizationSecurityPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeOrganizationSecurityPolicy{}, &ComputeOrganizationSecurityPolicyList{})
}
