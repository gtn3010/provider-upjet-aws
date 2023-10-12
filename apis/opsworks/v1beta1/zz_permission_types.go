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

type PermissionInitParameters struct {

	// Whether the user is allowed to use SSH to communicate with the instance
	AllowSSH *bool `json:"allowSsh,omitempty" tf:"allow_ssh,omitempty"`

	// Whether the user is allowed to use sudo to elevate privileges
	AllowSudo *bool `json:"allowSudo,omitempty" tf:"allow_sudo,omitempty"`

	// The users permission level. Mus be one of deny, show, deploy, manage, iam_only
	Level *string `json:"level,omitempty" tf:"level,omitempty"`
}

type PermissionObservation struct {

	// Whether the user is allowed to use SSH to communicate with the instance
	AllowSSH *bool `json:"allowSsh,omitempty" tf:"allow_ssh,omitempty"`

	// Whether the user is allowed to use sudo to elevate privileges
	AllowSudo *bool `json:"allowSudo,omitempty" tf:"allow_sudo,omitempty"`

	// The computed id of the permission. Please note that this is only used internally to identify the permission. This value is not used in aws.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The users permission level. Mus be one of deny, show, deploy, manage, iam_only
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// The stack to set the permissions for
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// The user's IAM ARN to set permissions for
	UserArn *string `json:"userArn,omitempty" tf:"user_arn,omitempty"`
}

type PermissionParameters struct {

	// Whether the user is allowed to use SSH to communicate with the instance
	// +kubebuilder:validation:Optional
	AllowSSH *bool `json:"allowSsh,omitempty" tf:"allow_ssh,omitempty"`

	// Whether the user is allowed to use sudo to elevate privileges
	// +kubebuilder:validation:Optional
	AllowSudo *bool `json:"allowSudo,omitempty" tf:"allow_sudo,omitempty"`

	// The users permission level. Mus be one of deny, show, deploy, manage, iam_only
	// +kubebuilder:validation:Optional
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// The stack to set the permissions for
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opsworks/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Reference to a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDRef *v1.Reference `json:"stackIdRef,omitempty" tf:"-"`

	// Selector for a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDSelector *v1.Selector `json:"stackIdSelector,omitempty" tf:"-"`

	// The user's IAM ARN to set permissions for
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	UserArn *string `json:"userArn,omitempty" tf:"user_arn,omitempty"`

	// Reference to a User in iam to populate userArn.
	// +kubebuilder:validation:Optional
	UserArnRef *v1.Reference `json:"userArnRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate userArn.
	// +kubebuilder:validation:Optional
	UserArnSelector *v1.Selector `json:"userArnSelector,omitempty" tf:"-"`
}

// PermissionSpec defines the desired state of Permission
type PermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionParameters `json:"forProvider"`
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
	InitProvider PermissionInitParameters `json:"initProvider,omitempty"`
}

// PermissionStatus defines the observed state of Permission.
type PermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Permission is the Schema for the Permissions API. Provides an OpsWorks permission resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Permission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionSpec   `json:"spec"`
	Status            PermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionList contains a list of Permissions
type PermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permission `json:"items"`
}

// Repository type metadata.
var (
	Permission_Kind             = "Permission"
	Permission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Permission_Kind}.String()
	Permission_KindAPIVersion   = Permission_Kind + "." + CRDGroupVersion.String()
	Permission_GroupVersionKind = CRDGroupVersion.WithKind(Permission_Kind)
)

func init() {
	SchemeBuilder.Register(&Permission{}, &PermissionList{})
}
