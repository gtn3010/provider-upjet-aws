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

type ReplicaExternalKeyInitParameters struct {

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the Default Key Policy section in the AWS Key Management Service Developer Guide.
	// The default value is false.
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30.
	DeletionWindowInDays *int64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// A description of the KMS key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be false. Imported keys default to true unless expired.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the default key policy to the KMS key.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: RFC3339 time string (YYYY-MM-DDTHH:MM:SSZ)
	ValidTo *string `json:"validTo,omitempty" tf:"valid_to,omitempty"`
}

type ReplicaExternalKeyObservation struct {

	// The Amazon Resource Name (ARN) of the replica key. The key ARNs of related multi-Region keys differ only in the Region value.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the Default Key Policy section in the AWS Key Management Service Developer Guide.
	// The default value is false.
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30.
	DeletionWindowInDays *int64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// A description of the KMS key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be false. Imported keys default to true unless expired.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether the key material expires. Empty when pending key material import, otherwise KEY_MATERIAL_EXPIRES or KEY_MATERIAL_DOES_NOT_EXPIRE.
	ExpirationModel *string `json:"expirationModel,omitempty" tf:"expiration_model,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The key ID of the replica key. Related multi-Region keys have the same key ID.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The state of the replica key.
	KeyState *string `json:"keyState,omitempty" tf:"key_state,omitempty"`

	// The cryptographic operations for which you can use the KMS key. This is a shared property of multi-Region keys.
	KeyUsage *string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the default key policy to the KMS key.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn *string `json:"primaryKeyArn,omitempty" tf:"primary_key_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: RFC3339 time string (YYYY-MM-DDTHH:MM:SSZ)
	ValidTo *string `json:"validTo,omitempty" tf:"valid_to,omitempty"`
}

type ReplicaExternalKeyParameters struct {

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	// For more information, refer to the scenario in the Default Key Policy section in the AWS Key Management Service Developer Guide.
	// The default value is false.
	// +kubebuilder:validation:Optional
	BypassPolicyLockoutSafetyCheck *bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" tf:"bypass_policy_lockout_safety_check,omitempty"`

	// The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	// If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30.
	// +kubebuilder:validation:Optional
	DeletionWindowInDays *int64 `json:"deletionWindowInDays,omitempty" tf:"deletion_window_in_days,omitempty"`

	// A description of the KMS key.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be false. Imported keys default to true unless expired.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Base64 encoded 256-bit symmetric encryption key material to import. The KMS key is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
	// +kubebuilder:validation:Optional
	KeyMaterialBase64SecretRef *v1.SecretKeySelector `json:"keyMaterialBase64SecretRef,omitempty" tf:"-"`

	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the default key policy to the KMS key.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	// +crossplane:generate:reference:type=ExternalKey
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PrimaryKeyArn *string `json:"primaryKeyArn,omitempty" tf:"primary_key_arn,omitempty"`

	// Reference to a ExternalKey to populate primaryKeyArn.
	// +kubebuilder:validation:Optional
	PrimaryKeyArnRef *v1.Reference `json:"primaryKeyArnRef,omitempty" tf:"-"`

	// Selector for a ExternalKey to populate primaryKeyArn.
	// +kubebuilder:validation:Optional
	PrimaryKeyArnSelector *v1.Selector `json:"primaryKeyArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: RFC3339 time string (YYYY-MM-DDTHH:MM:SSZ)
	// +kubebuilder:validation:Optional
	ValidTo *string `json:"validTo,omitempty" tf:"valid_to,omitempty"`
}

// ReplicaExternalKeySpec defines the desired state of ReplicaExternalKey
type ReplicaExternalKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicaExternalKeyParameters `json:"forProvider"`
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
	InitProvider ReplicaExternalKeyInitParameters `json:"initProvider,omitempty"`
}

// ReplicaExternalKeyStatus defines the observed state of ReplicaExternalKey.
type ReplicaExternalKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicaExternalKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicaExternalKey is the Schema for the ReplicaExternalKeys API. Manages a KMS multi-Region replica key that uses external key material.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReplicaExternalKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicaExternalKeySpec   `json:"spec"`
	Status            ReplicaExternalKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicaExternalKeyList contains a list of ReplicaExternalKeys
type ReplicaExternalKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicaExternalKey `json:"items"`
}

// Repository type metadata.
var (
	ReplicaExternalKey_Kind             = "ReplicaExternalKey"
	ReplicaExternalKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicaExternalKey_Kind}.String()
	ReplicaExternalKey_KindAPIVersion   = ReplicaExternalKey_Kind + "." + CRDGroupVersion.String()
	ReplicaExternalKey_GroupVersionKind = CRDGroupVersion.WithKind(ReplicaExternalKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicaExternalKey{}, &ReplicaExternalKeyList{})
}
