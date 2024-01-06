/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBClusterParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBClusterParameterGroupNameRef,
		Selector:     mg.Spec.ForProvider.DBClusterParameterGroupNameSelector,
		To: reference.To{
			List:    &ClusterParameterGroupList{},
			Managed: &ClusterParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBClusterParameterGroupName")
	}
	mg.Spec.ForProvider.DBClusterParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBClusterParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DBClusterParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.DBClusterParameterGroupNameRef,
		Selector:     mg.Spec.InitProvider.DBClusterParameterGroupNameSelector,
		To: reference.To{
			List:    &ClusterParameterGroupList{},
			Managed: &ClusterParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DBClusterParameterGroupName")
	}
	mg.Spec.InitProvider.DBClusterParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DBClusterParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KMSKeyIDRef,
		Selector:     mg.Spec.InitProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyID")
	}
	mg.Spec.InitProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.InitProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCSecurityGroupIds")
	}
	mg.Spec.InitProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ClusterInstance.
func (mg *ClusterInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIdentifierRef,
		Selector:     mg.Spec.InitProvider.ClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterIdentifier")
	}
	mg.Spec.InitProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterSnapshot.
func (mg *ClusterSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBClusterIdentifier),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DBClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.DBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBClusterIdentifier")
	}
	mg.Spec.ForProvider.DBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBClusterIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DBClusterIdentifier),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.DBClusterIdentifierRef,
		Selector:     mg.Spec.InitProvider.DBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DBClusterIdentifier")
	}
	mg.Spec.InitProvider.DBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventSubscription.
func (mg *EventSubscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnsTopicArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.SnsTopicArnRef,
		Selector:     mg.Spec.ForProvider.SnsTopicArnSelector,
		To: reference.To{
			List:    &v1beta12.TopicList{},
			Managed: &v1beta12.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnsTopicArn")
	}
	mg.Spec.ForProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnsTopicArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SnsTopicArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.SnsTopicArnRef,
		Selector:     mg.Spec.InitProvider.SnsTopicArnSelector,
		To: reference.To{
			List:    &v1beta12.TopicList{},
			Managed: &v1beta12.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SnsTopicArn")
	}
	mg.Spec.InitProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SnsTopicArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GlobalCluster.
func (mg *GlobalCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceDBClusterIdentifier),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.SourceDBClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.SourceDBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceDBClusterIdentifier")
	}
	mg.Spec.ForProvider.SourceDBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceDBClusterIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceDBClusterIdentifier),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.SourceDBClusterIdentifierRef,
		Selector:     mg.Spec.InitProvider.SourceDBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SourceDBClusterIdentifier")
	}
	mg.Spec.InitProvider.SourceDBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SourceDBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIdsRefs,
		Selector:      mg.Spec.ForProvider.SubnetIdsSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.SubnetIdsRefs,
		Selector:      mg.Spec.InitProvider.SubnetIdsSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetIds")
	}
	mg.Spec.InitProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetIdsRefs = mrsp.ResolvedReferences

	return nil
}
