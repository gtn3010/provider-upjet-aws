// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VoiceConnectorLoggingInitParameters struct {

	// When true, enables logging of detailed media metrics for Voice Connectors to Amazon CloudWatch logs.
	EnableMediaMetricLogs *bool `json:"enableMediaMetricLogs,omitempty" tf:"enable_media_metric_logs,omitempty"`

	// When true, enables SIP message logs for sending to Amazon CloudWatch Logs.
	EnableSIPLogs *bool `json:"enableSipLogs,omitempty" tf:"enable_sip_logs,omitempty"`
}

type VoiceConnectorLoggingObservation struct {

	// When true, enables logging of detailed media metrics for Voice Connectors to Amazon CloudWatch logs.
	EnableMediaMetricLogs *bool `json:"enableMediaMetricLogs,omitempty" tf:"enable_media_metric_logs,omitempty"`

	// When true, enables SIP message logs for sending to Amazon CloudWatch Logs.
	EnableSIPLogs *bool `json:"enableSipLogs,omitempty" tf:"enable_sip_logs,omitempty"`

	// The Amazon Chime Voice Connector ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`
}

type VoiceConnectorLoggingParameters struct {

	// When true, enables logging of detailed media metrics for Voice Connectors to Amazon CloudWatch logs.
	// +kubebuilder:validation:Optional
	EnableMediaMetricLogs *bool `json:"enableMediaMetricLogs,omitempty" tf:"enable_media_metric_logs,omitempty"`

	// When true, enables SIP message logs for sending to Amazon CloudWatch Logs.
	// +kubebuilder:validation:Optional
	EnableSIPLogs *bool `json:"enableSipLogs,omitempty" tf:"enable_sip_logs,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Chime Voice Connector ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/chime/v1beta1.VoiceConnector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`

	// Reference to a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDRef *v1.Reference `json:"voiceConnectorIdRef,omitempty" tf:"-"`

	// Selector for a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDSelector *v1.Selector `json:"voiceConnectorIdSelector,omitempty" tf:"-"`
}

// VoiceConnectorLoggingSpec defines the desired state of VoiceConnectorLogging
type VoiceConnectorLoggingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VoiceConnectorLoggingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VoiceConnectorLoggingInitParameters `json:"initProvider,omitempty"`
}

// VoiceConnectorLoggingStatus defines the observed state of VoiceConnectorLogging.
type VoiceConnectorLoggingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VoiceConnectorLoggingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorLogging is the Schema for the VoiceConnectorLoggings API. Adds a logging configuration for the specified Amazon Chime Voice Connector. The logging configuration specifies whether SIP message logs are enabled for sending to Amazon CloudWatch Logs.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VoiceConnectorLogging struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VoiceConnectorLoggingSpec   `json:"spec"`
	Status            VoiceConnectorLoggingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorLoggingList contains a list of VoiceConnectorLoggings
type VoiceConnectorLoggingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VoiceConnectorLogging `json:"items"`
}

// Repository type metadata.
var (
	VoiceConnectorLogging_Kind             = "VoiceConnectorLogging"
	VoiceConnectorLogging_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VoiceConnectorLogging_Kind}.String()
	VoiceConnectorLogging_KindAPIVersion   = VoiceConnectorLogging_Kind + "." + CRDGroupVersion.String()
	VoiceConnectorLogging_GroupVersionKind = CRDGroupVersion.WithKind(VoiceConnectorLogging_Kind)
)

func init() {
	SchemeBuilder.Register(&VoiceConnectorLogging{}, &VoiceConnectorLoggingList{})
}
