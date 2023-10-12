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

type VPNGatewayInitParameters_2 struct {

	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	// The Availability Zone for the virtual private gateway.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPNGatewayObservation_2 struct {

	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	// Amazon Resource Name (ARN) of the VPN Gateway.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Availability Zone for the virtual private gateway.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The ID of the VPN Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VPC ID to create in.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPNGatewayParameters_2 struct {

	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	// +kubebuilder:validation:Optional
	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	// The Availability Zone for the virtual private gateway.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC ID to create in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPNGatewaySpec defines the desired state of VPNGateway
type VPNGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNGatewayParameters_2 `json:"forProvider"`
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
	InitProvider VPNGatewayInitParameters_2 `json:"initProvider,omitempty"`
}

// VPNGatewayStatus defines the observed state of VPNGateway.
type VPNGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNGatewayObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGateway is the Schema for the VPNGateways API. Provides a resource to create a VPC VPN Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNGatewaySpec   `json:"spec"`
	Status            VPNGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayList contains a list of VPNGateways
type VPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNGateway `json:"items"`
}

// Repository type metadata.
var (
	VPNGateway_Kind             = "VPNGateway"
	VPNGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNGateway_Kind}.String()
	VPNGateway_KindAPIVersion   = VPNGateway_Kind + "." + CRDGroupVersion.String()
	VPNGateway_GroupVersionKind = CRDGroupVersion.WithKind(VPNGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNGateway{}, &VPNGatewayList{})
}
