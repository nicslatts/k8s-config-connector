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

type KMSCryptoKeyVersionSpec struct {
	/* Immutable. The name of the cryptoKey associated with the CryptoKeyVersions.
	Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}''. */
	CryptoKey string `json:"cryptoKey"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The current state of the CryptoKeyVersion. Possible values: ["PENDING_GENERATION", "ENABLED", "DISABLED", "DESTROYED", "DESTROY_SCHEDULED", "PENDING_IMPORT", "IMPORT_FAILED"]. */
	// +optional
	State *string `json:"state,omitempty"`
}

type CryptokeyversionAttestationStatus struct {
	/* The certificate chains needed to validate the attestation. */
	// +optional
	CertChains *CryptokeyversionCertChainsStatus `json:"certChains,omitempty"`

	/* The attestation data provided by the HSM when the key operation was performed. */
	// +optional
	Content *string `json:"content,omitempty"`

	/* ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level and EXTERNAL_VPC protection levels. */
	// +optional
	ExternalProtectionLevelOptions *CryptokeyversionExternalProtectionLevelOptionsStatus `json:"externalProtectionLevelOptions,omitempty"`

	/* The format of the attestation data. */
	// +optional
	Format *string `json:"format,omitempty"`
}

type CryptokeyversionCertChainsStatus struct {
	/* Cavium certificate chain corresponding to the attestation. */
	// +optional
	CaviumCerts *string `json:"caviumCerts,omitempty"`

	/* Google card certificate chain corresponding to the attestation. */
	// +optional
	GoogleCardCerts *string `json:"googleCardCerts,omitempty"`

	/* Google partition certificate chain corresponding to the attestation. */
	// +optional
	GooglePartitionCerts *string `json:"googlePartitionCerts,omitempty"`
}

type CryptokeyversionExternalProtectionLevelOptionsStatus struct {
	/* The path to the external key material on the EKM when using EkmConnection e.g., "v0/my/key". Set this field instead of externalKeyUri when using an EkmConnection. */
	// +optional
	EkmConnectionKeyPath *string `json:"ekmConnectionKeyPath,omitempty"`

	/* The URI for an external resource that this CryptoKeyVersion represents. */
	// +optional
	ExternalKeyUri *string `json:"externalKeyUri,omitempty"`
}

type KMSCryptoKeyVersionStatus struct {
	/* Conditions represent the latest available observations of the
	   KMSCryptoKeyVersion's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports. */
	// +optional
	Algorithm *string `json:"algorithm,omitempty"`

	/* Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
	Only provided for key versions with protectionLevel HSM. */
	// +optional
	Attestation []CryptokeyversionAttestationStatus `json:"attestation,omitempty"`

	/* The time this CryptoKeyVersion key material was generated. */
	// +optional
	GenerateTime *string `json:"generateTime,omitempty"`

	/* The resource name for this CryptoKeyVersion. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. */
	// +optional
	ProtectionLevel *string `json:"protectionLevel,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmscryptokeyversion;gcpkmscryptokeyversions
// +kubebuilder:subresource:status

// KMSCryptoKeyVersion is the Schema for the kms API
// +k8s:openapi-gen=true
type KMSCryptoKeyVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSCryptoKeyVersionSpec   `json:"spec,omitempty"`
	Status KMSCryptoKeyVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KMSCryptoKeyVersionList contains a list of KMSCryptoKeyVersion
type KMSCryptoKeyVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSCryptoKeyVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSCryptoKeyVersion{}, &KMSCryptoKeyVersionList{})
}
