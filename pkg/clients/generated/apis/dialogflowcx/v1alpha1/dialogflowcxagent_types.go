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

type AgentSpeechToTextSettings struct {
	/* Whether to use speech adaptation for speech recognition. */
	// +optional
	EnableSpeechAdaptation *bool `json:"enableSpeechAdaptation,omitempty"`
}

type DialogflowCXAgentSpec struct {
	/* The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration. */
	// +optional
	AvatarUri *string `json:"avatarUri,omitempty"`

	/* Immutable. The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	for a list of the currently supported language codes. This field cannot be updated after creation. */
	DefaultLanguageCode string `json:"defaultLanguageCode"`

	/* The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The human-readable name of the agent, unique within the location. */
	DisplayName string `json:"displayName"`

	/* Indicates if automatic spell correction is enabled in detect intent requests. */
	// +optional
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty"`

	/* Determines whether this agent should log conversation queries. */
	// +optional
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`

	/* Immutable. The name of the location this agent is located in.

	~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
	This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	Another options is to use global location so you don't need to manually configure location settings. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>. */
	// +optional
	SecuritySettings *string `json:"securitySettings,omitempty"`

	/* Settings related to speech recognition. */
	// +optional
	SpeechToTextSettings *AgentSpeechToTextSettings `json:"speechToTextSettings,omitempty"`

	/* The list of all languages supported by this agent (except for the default_language_code). */
	// +optional
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty"`

	/* The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	Europe/Paris. */
	TimeZone string `json:"timeZone"`
}

type DialogflowCXAgentStatus struct {
	/* Conditions represent the latest available observations of the
	   DialogflowCXAgent's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The unique identifier of the agent. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>. */
	// +optional
	StartFlow *string `json:"startFlow,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxagent;gcpdialogflowcxagents
// +kubebuilder:subresource:status

// DialogflowCXAgent is the Schema for the dialogflowcx API
// +k8s:openapi-gen=true
type DialogflowCXAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXAgentSpec   `json:"spec,omitempty"`
	Status DialogflowCXAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DialogflowCXAgentList contains a list of DialogflowCXAgent
type DialogflowCXAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXAgent{}, &DialogflowCXAgentList{})
}
