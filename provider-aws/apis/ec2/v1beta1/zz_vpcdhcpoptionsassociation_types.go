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

type VPCDHCPOptionsAssociationObservation struct {

	// The ID of the DHCP Options Set Association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPCDHCPOptionsAssociationParameters struct {

	// The ID of the DHCP Options Set to associate to the VPC.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPCDHCPOptions
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DHCPOptionsID *string `json:"dhcpOptionsId,omitempty" tf:"dhcp_options_id,omitempty"`

	// Reference to a VPCDHCPOptions in ec2 to populate dhcpOptionsId.
	// +kubebuilder:validation:Optional
	DHCPOptionsIDRef *v1.Reference `json:"dhcpOptionsIdRef,omitempty" tf:"-"`

	// Selector for a VPCDHCPOptions in ec2 to populate dhcpOptionsId.
	// +kubebuilder:validation:Optional
	DHCPOptionsIDSelector *v1.Selector `json:"dhcpOptionsIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the VPC to which we would like to associate a DHCP Options Set.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPCDHCPOptionsAssociationSpec defines the desired state of VPCDHCPOptionsAssociation
type VPCDHCPOptionsAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCDHCPOptionsAssociationParameters `json:"forProvider"`
}

// VPCDHCPOptionsAssociationStatus defines the observed state of VPCDHCPOptionsAssociation.
type VPCDHCPOptionsAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCDHCPOptionsAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCDHCPOptionsAssociation is the Schema for the VPCDHCPOptionsAssociations API. Provides a VPC DHCP Options Association resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCDHCPOptionsAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCDHCPOptionsAssociationSpec   `json:"spec"`
	Status            VPCDHCPOptionsAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCDHCPOptionsAssociationList contains a list of VPCDHCPOptionsAssociations
type VPCDHCPOptionsAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCDHCPOptionsAssociation `json:"items"`
}

// Repository type metadata.
var (
	VPCDHCPOptionsAssociation_Kind             = "VPCDHCPOptionsAssociation"
	VPCDHCPOptionsAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCDHCPOptionsAssociation_Kind}.String()
	VPCDHCPOptionsAssociation_KindAPIVersion   = VPCDHCPOptionsAssociation_Kind + "." + CRDGroupVersion.String()
	VPCDHCPOptionsAssociation_GroupVersionKind = CRDGroupVersion.WithKind(VPCDHCPOptionsAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCDHCPOptionsAssociation{}, &VPCDHCPOptionsAssociationList{})
}