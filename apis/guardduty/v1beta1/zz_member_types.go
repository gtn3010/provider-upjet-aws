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

type MemberInitParameters struct {

	// Boolean whether an email notification is sent to the accounts. Defaults to false.
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`

	// Email address for member account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Message for invitation.
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`

	// Boolean whether to invite the account to GuardDuty as a member. Defaults to false.
	Invite *bool `json:"invite,omitempty" tf:"invite,omitempty"`
}

type MemberObservation struct {

	// AWS account ID for member account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The detector ID of the GuardDuty account where you want to create member accounts.
	DetectorID *string `json:"detectorId,omitempty" tf:"detector_id,omitempty"`

	// Boolean whether an email notification is sent to the accounts. Defaults to false.
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`

	// Email address for member account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ID of the GuardDuty member
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Message for invitation.
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`

	// Boolean whether to invite the account to GuardDuty as a member. Defaults to false.
	Invite *bool `json:"invite,omitempty" tf:"invite,omitempty"`

	// The status of the relationship between the member account and its primary account. More information can be found in Amazon GuardDuty API Reference.
	RelationshipStatus *string `json:"relationshipStatus,omitempty" tf:"relationship_status,omitempty"`
}

type MemberParameters struct {

	// AWS account ID for member account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/guardduty/v1beta1.Detector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("account_id",true)
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Detector in guardduty to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Detector in guardduty to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// The detector ID of the GuardDuty account where you want to create member accounts.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/guardduty/v1beta1.Detector
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DetectorID *string `json:"detectorId,omitempty" tf:"detector_id,omitempty"`

	// Reference to a Detector in guardduty to populate detectorId.
	// +kubebuilder:validation:Optional
	DetectorIDRef *v1.Reference `json:"detectorIdRef,omitempty" tf:"-"`

	// Selector for a Detector in guardduty to populate detectorId.
	// +kubebuilder:validation:Optional
	DetectorIDSelector *v1.Selector `json:"detectorIdSelector,omitempty" tf:"-"`

	// Boolean whether an email notification is sent to the accounts. Defaults to false.
	// +kubebuilder:validation:Optional
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`

	// Email address for member account.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Message for invitation.
	// +kubebuilder:validation:Optional
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`

	// Boolean whether to invite the account to GuardDuty as a member. Defaults to false.
	// +kubebuilder:validation:Optional
	Invite *bool `json:"invite,omitempty" tf:"invite,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
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
	InitProvider MemberInitParameters `json:"initProvider,omitempty"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Member is the Schema for the Members API. Provides a resource to manage a GuardDuty member
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	Spec   MemberSpec   `json:"spec"`
	Status MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}
