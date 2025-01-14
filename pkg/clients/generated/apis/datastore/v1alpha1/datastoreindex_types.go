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

type IndexProperties struct {
	/* Immutable. The direction the index should optimize for sorting. Possible values: ["ASCENDING", "DESCENDING"]. */
	Direction string `json:"direction"`

	/* Immutable. The property name to index. */
	Name string `json:"name"`
}

type DatastoreIndexSpec struct {
	/* Immutable. Policy for including ancestors in the index. Default value: "NONE" Possible values: ["NONE", "ALL_ANCESTORS"]. */
	// +optional
	Ancestor *string `json:"ancestor,omitempty"`

	/* Immutable. The entity kind which the index applies to. */
	Kind string `json:"kind"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. An ordered list of properties to index on. */
	// +optional
	Properties []IndexProperties `json:"properties,omitempty"`

	/* Immutable. Optional. The service-generated indexId of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type DatastoreIndexStatus struct {
	/* Conditions represent the latest available observations of the
	   DatastoreIndex's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The index id. */
	// +optional
	IndexId *string `json:"indexId,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatastoreindex;gcpdatastoreindexes
// +kubebuilder:subresource:status

// DatastoreIndex is the Schema for the datastore API
// +k8s:openapi-gen=true
type DatastoreIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastoreIndexSpec   `json:"spec,omitempty"`
	Status DatastoreIndexStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatastoreIndexList contains a list of DatastoreIndex
type DatastoreIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastoreIndex `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatastoreIndex{}, &DatastoreIndexList{})
}
