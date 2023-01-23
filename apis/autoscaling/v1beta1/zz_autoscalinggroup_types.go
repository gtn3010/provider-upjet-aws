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

type AcceleratorCountObservation struct {
}

type AcceleratorCountParameters struct {

	// Maximum.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// Minimum.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type AcceleratorTotalMemoryMibObservation struct {
}

type AcceleratorTotalMemoryMibParameters struct {

	// Maximum.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// Minimum.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type AutoscalingGroupObservation struct {

	// ARN for this Auto Scaling Group
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Auto Scaling Group id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of elastic load balancer names to add to the autoscaling
	// group names. Only valid for classic load balancers. For ALBs, use target_group_arns instead.
	LoadBalancers []*string `json:"loadBalancers,omitempty" tf:"load_balancers,omitempty"`

	// Set of aws_alb_target_group ARNs, for use with Application or Network Load Balancing.
	TargetGroupArns []*string `json:"targetGroupArns,omitempty" tf:"target_group_arns,omitempty"`
}

type AutoscalingGroupParameters struct {

	// List of one or more availability zones for the group. Used for EC2-Classic, attaching a network interface via id from a launch template and default subnets when not specified with vpc_zone_identifier argument. Conflicts with vpc_zone_identifier.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Whether capacity rebalance is enabled. Otherwise, capacity rebalance is disabled.
	// +kubebuilder:validation:Optional
	CapacityRebalance *bool `json:"capacityRebalance,omitempty" tf:"capacity_rebalance,omitempty"`

	// Reserved.
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// Amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
	// +kubebuilder:validation:Optional
	DefaultCooldown *float64 `json:"defaultCooldown,omitempty" tf:"default_cooldown,omitempty"`

	// Amount of time, in seconds, until a newly launched instance can contribute to the Amazon CloudWatch metrics. This delay lets an instance finish initializing before Amazon EC2 Auto Scaling aggregates instance metrics, resulting in more reliable usage data. Set this value equal to the amount of time that it takes for resource consumption to become stable after an instance reaches the InService state. (See Set the default instance warmup for an Auto Scaling group)
	// +kubebuilder:validation:Optional
	DefaultInstanceWarmup *float64 `json:"defaultInstanceWarmup,omitempty" tf:"default_instance_warmup,omitempty"`

	// Number of Amazon EC2 instances that
	// should be running in the group. (See also Waiting for
	// Capacity below.)
	// +kubebuilder:validation:Optional
	DesiredCapacity *float64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// The unit of measurement for the value specified for desired_capacity. Supported for attribute-based instance type selection only. Valid values: "units", "vcpu", "memory-mib".
	// +kubebuilder:validation:Optional
	DesiredCapacityType *string `json:"desiredCapacityType,omitempty" tf:"desired_capacity_type,omitempty"`

	// List of metrics to collect. The allowed values are defined by the underlying AWS API.
	// +kubebuilder:validation:Optional
	EnabledMetrics []*string `json:"enabledMetrics,omitempty" tf:"enabled_metrics,omitempty"`

	// Allows deleting the Auto Scaling Group without waiting
	// for all instances in the pool to terminate.  You can force an Auto Scaling Group to delete
	// even if it's in the process of scaling a resource.  This bypasses that
	// behavior and potentially leaves resources dangling.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// If this block is configured, add a Warm Pool
	// to the specified Auto Scaling group. Defined below
	// +kubebuilder:validation:Optional
	ForceDeleteWarmPool *bool `json:"forceDeleteWarmPool,omitempty" tf:"force_delete_warm_pool,omitempty"`

	// Time (in seconds) after instance comes into service before checking health.
	// +kubebuilder:validation:Optional
	HealthCheckGracePeriod *float64 `json:"healthCheckGracePeriod,omitempty" tf:"health_check_grace_period,omitempty"`

	// "EC2" or "ELB". Controls how health checking is done.
	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// One or more
	// Lifecycle Hooks
	// to attach to the Auto Scaling Group before instances are launched. The
	// syntax is exactly the same as the separate
	// aws_autoscaling_lifecycle_hook
	// resource, without the autoscaling_group_name attribute. Please note that this will only work when creating
	// a new Auto Scaling Group. For all other use-cases, please use aws_autoscaling_lifecycle_hook resource.
	// +kubebuilder:validation:Optional
	InitialLifecycleHook []InitialLifecycleHookParameters `json:"initialLifecycleHook,omitempty" tf:"initial_lifecycle_hook,omitempty"`

	// If this block is configured, start an
	// Instance Refresh
	// when this Auto Scaling Group is updated. Defined below.
	// +kubebuilder:validation:Optional
	InstanceRefresh []InstanceRefreshParameters `json:"instanceRefresh,omitempty" tf:"instance_refresh,omitempty"`

	// Name of the launch configuration to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/autoscaling/v1beta1.LaunchConfiguration
	// +kubebuilder:validation:Optional
	LaunchConfiguration *string `json:"launchConfiguration,omitempty" tf:"launch_configuration,omitempty"`

	// Reference to a LaunchConfiguration in autoscaling to populate launchConfiguration.
	// +kubebuilder:validation:Optional
	LaunchConfigurationRef *v1.Reference `json:"launchConfigurationRef,omitempty" tf:"-"`

	// Selector for a LaunchConfiguration in autoscaling to populate launchConfiguration.
	// +kubebuilder:validation:Optional
	LaunchConfigurationSelector *v1.Selector `json:"launchConfigurationSelector,omitempty" tf:"-"`

	// Nested argument with Launch template specification to use to launch instances. See Launch Template below for more details.
	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// Maximum amount of time, in seconds, that an instance can be in service, values must be either equal to 0 or between 86400 and 31536000 seconds.
	// +kubebuilder:validation:Optional
	MaxInstanceLifetime *float64 `json:"maxInstanceLifetime,omitempty" tf:"max_instance_lifetime,omitempty"`

	// Maximum size of the Auto Scaling Group.
	// +kubebuilder:validation:Required
	MaxSize *float64 `json:"maxSize" tf:"max_size,omitempty"`

	// Granularity to associate with the metrics to collect. The only valid value is 1Minute. Default is 1Minute.
	// +kubebuilder:validation:Optional
	MetricsGranularity *string `json:"metricsGranularity,omitempty" tf:"metrics_granularity,omitempty"`

	// Updates will not wait on ELB instance number changes.
	// (See also Waiting for Capacity below.)
	// +kubebuilder:validation:Optional
	MinELBCapacity *float64 `json:"minElbCapacity,omitempty" tf:"min_elb_capacity,omitempty"`

	// Minimum size of the Auto Scaling Group.
	// (See also Waiting for Capacity below.)
	// +kubebuilder:validation:Required
	MinSize *float64 `json:"minSize" tf:"min_size,omitempty"`

	// Configuration block containing settings to define launch targets for Auto Scaling groups. See Mixed Instances Policy below for more details.
	// +kubebuilder:validation:Optional
	MixedInstancesPolicy []MixedInstancesPolicyParameters `json:"mixedInstancesPolicy,omitempty" tf:"mixed_instances_policy,omitempty"`

	// Name of the placement group into which you'll launch your instances, if any.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.PlacementGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group,omitempty"`

	// Reference to a PlacementGroup in ec2 to populate placementGroup.
	// +kubebuilder:validation:Optional
	PlacementGroupRef *v1.Reference `json:"placementGroupRef,omitempty" tf:"-"`

	// Selector for a PlacementGroup in ec2 to populate placementGroup.
	// +kubebuilder:validation:Optional
	PlacementGroupSelector *v1.Selector `json:"placementGroupSelector,omitempty" tf:"-"`

	// in protection
	// in the Amazon EC2 Auto Scaling User Guide.
	// +kubebuilder:validation:Optional
	ProtectFromScaleIn *bool `json:"protectFromScaleIn,omitempty" tf:"protect_from_scale_in,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// linked role that the ASG will use to call other AWS services
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceLinkedRoleArn *string `json:"serviceLinkedRoleArn,omitempty" tf:"service_linked_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceLinkedRoleArn.
	// +kubebuilder:validation:Optional
	ServiceLinkedRoleArnRef *v1.Reference `json:"serviceLinkedRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceLinkedRoleArn.
	// +kubebuilder:validation:Optional
	ServiceLinkedRoleArnSelector *v1.Selector `json:"serviceLinkedRoleArnSelector,omitempty" tf:"-"`

	// List of processes to suspend for the Auto Scaling Group. The allowed values are Launch, Terminate, HealthCheck, ReplaceUnhealthy, AZRebalance, AlarmNotification, ScheduledActions, AddToLoadBalancer, InstanceRefresh.
	// Note that if you suspend either the Launch or Terminate process types, it can prevent your Auto Scaling Group from functioning properly.
	// +kubebuilder:validation:Optional
	SuspendedProcesses []*string `json:"suspendedProcesses,omitempty" tf:"suspended_processes,omitempty"`

	// Configuration block(s) containing resource tags. Conflicts with tags. See Tag below for more details.
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags []map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// List of policies to decide how the instances in the Auto Scaling Group should be terminated. The allowed values are OldestInstance, NewestInstance, OldestLaunchConfiguration, ClosestToNextInstanceHour, OldestLaunchTemplate, AllocationStrategy, Default. Additionally, the ARN of a Lambda function can be specified for custom termination policies.
	// +kubebuilder:validation:Optional
	TerminationPolicies []*string `json:"terminationPolicies,omitempty" tf:"termination_policies,omitempty"`

	// List of subnet IDs to launch resources in. Subnets automatically determine which availability zones the group will reside. Conflicts with availability_zones.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	VPCZoneIdentifier []*string `json:"vpcZoneIdentifier,omitempty" tf:"vpc_zone_identifier,omitempty"`

	// References to Subnet in ec2 to populate vpcZoneIdentifier.
	// +kubebuilder:validation:Optional
	VPCZoneIdentifierRefs []v1.Reference `json:"vpcZoneIdentifierRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate vpcZoneIdentifier.
	// +kubebuilder:validation:Optional
	VPCZoneIdentifierSelector *v1.Selector `json:"vpcZoneIdentifierSelector,omitempty" tf:"-"`

	// (See also Waiting
	// for Capacity below.
	// +kubebuilder:validation:Optional
	WaitForCapacityTimeout *string `json:"waitForCapacityTimeout,omitempty" tf:"wait_for_capacity_timeout,omitempty"`

	// (Takes
	// precedence over min_elb_capacity behavior.)
	// (See also Waiting for Capacity below.)
	// +kubebuilder:validation:Optional
	WaitForELBCapacity *float64 `json:"waitForElbCapacity,omitempty" tf:"wait_for_elb_capacity,omitempty"`

	// If this block is configured, add a Warm Pool
	// to the specified Auto Scaling group. Defined below
	// +kubebuilder:validation:Optional
	WarmPool []WarmPoolParameters `json:"warmPool,omitempty" tf:"warm_pool,omitempty"`
}

type BaselineEBSBandwidthMbpsObservation struct {
}

type BaselineEBSBandwidthMbpsParameters struct {

	// Maximum.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// Minimum.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type InitialLifecycleHookObservation struct {
}

type InitialLifecycleHookParameters struct {

	// +kubebuilder:validation:Optional
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// +kubebuilder:validation:Optional
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// +kubebuilder:validation:Required
	LifecycleTransition *string `json:"lifecycleTransition" tf:"lifecycle_transition,omitempty"`

	// Name of the Auto Scaling Group. Conflicts with name_prefix.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// ARN for this Auto Scaling Group
	// +kubebuilder:validation:Optional
	NotificationTargetArn *string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`

	// ARN for this Auto Scaling Group
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type InstanceRefreshObservation struct {
}

type InstanceRefreshParameters struct {

	// Override default parameters for Instance Refresh.
	// +kubebuilder:validation:Optional
	Preferences []PreferencesParameters `json:"preferences,omitempty" tf:"preferences,omitempty"`

	// Strategy to use for instance refresh. The only allowed value is Rolling. See StartInstanceRefresh Action for more information.
	// +kubebuilder:validation:Required
	Strategy *string `json:"strategy" tf:"strategy,omitempty"`

	// Set of additional property names that will trigger an Instance Refresh. A refresh will always be triggered by a change in any of launch_configuration, launch_template, or mixed_instances_policy.
	// +kubebuilder:validation:Optional
	Triggers []*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type InstanceRequirementsObservation struct {
}

type InstanceRequirementsParameters struct {

	// Block describing the minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips). Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	AcceleratorCount []AcceleratorCountParameters `json:"acceleratorCount,omitempty" tf:"accelerator_count,omitempty"`

	// List of accelerator manufacturer names. Default is any manufacturer.
	// +kubebuilder:validation:Optional
	AcceleratorManufacturers []*string `json:"acceleratorManufacturers,omitempty" tf:"accelerator_manufacturers,omitempty"`

	// List of accelerator names. Default is any acclerator.
	// +kubebuilder:validation:Optional
	AcceleratorNames []*string `json:"acceleratorNames,omitempty" tf:"accelerator_names,omitempty"`

	// Block describing the minimum and maximum total memory of the accelerators. Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	AcceleratorTotalMemoryMib []AcceleratorTotalMemoryMibParameters `json:"acceleratorTotalMemoryMib,omitempty" tf:"accelerator_total_memory_mib,omitempty"`

	// List of accelerator types. Default is any accelerator type.
	// +kubebuilder:validation:Optional
	AcceleratorTypes []*string `json:"acceleratorTypes,omitempty" tf:"accelerator_types,omitempty"`

	// Indicate whether bare metal instace types should be included, excluded, or required. Default is excluded.
	// +kubebuilder:validation:Optional
	BareMetal *string `json:"bareMetal,omitempty" tf:"bare_metal,omitempty"`

	// Block describing the minimum and maximum baseline EBS bandwidth, in Mbps. Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	BaselineEBSBandwidthMbps []BaselineEBSBandwidthMbpsParameters `json:"baselineEbsBandwidthMbps,omitempty" tf:"baseline_ebs_bandwidth_mbps,omitempty"`

	// Indicate whether burstable performance instance types should be included, excluded, or required. Default is excluded.
	// +kubebuilder:validation:Optional
	BurstablePerformance *string `json:"burstablePerformance,omitempty" tf:"burstable_performance,omitempty"`

	// List of CPU manufacturer names. Default is any manufacturer.
	// +kubebuilder:validation:Optional
	CPUManufacturers []*string `json:"cpuManufacturers,omitempty" tf:"cpu_manufacturers,omitempty"`

	// List of instance types to exclude. You can use strings with one or more wild cards, represented by an asterisk (*). The following are examples: c5*, m5a.*, r*, *3*. For example, if you specify c5*, you are excluding the entire C5 instance family, which includes all C5a and C5n instance types. If you specify m5a.*, you are excluding all the M5a instance types, but not the M5n instance types. Maximum of 400 entries in the list; each entry is limited to 30 characters. Default is no excluded instance types.
	// +kubebuilder:validation:Optional
	ExcludedInstanceTypes []*string `json:"excludedInstanceTypes,omitempty" tf:"excluded_instance_types,omitempty"`

	// List of instance generation names. Default is any generation.
	// +kubebuilder:validation:Optional
	InstanceGenerations []*string `json:"instanceGenerations,omitempty" tf:"instance_generations,omitempty"`

	// Indicate whether instance types with local storage volumes are included, excluded, or required. Default is included.
	// +kubebuilder:validation:Optional
	LocalStorage *string `json:"localStorage,omitempty" tf:"local_storage,omitempty"`

	// List of local storage type names. Default any storage type.
	// +kubebuilder:validation:Optional
	LocalStorageTypes []*string `json:"localStorageTypes,omitempty" tf:"local_storage_types,omitempty"`

	// Block describing the minimum and maximum amount of memory (GiB) per vCPU. Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	MemoryGibPerVcpu []MemoryGibPerVcpuParameters `json:"memoryGibPerVcpu,omitempty" tf:"memory_gib_per_vcpu,omitempty"`

	// Block describing the minimum and maximum amount of memory (MiB). Default is no maximum.
	// +kubebuilder:validation:Optional
	MemoryMib []MemoryMibParameters `json:"memoryMib,omitempty" tf:"memory_mib,omitempty"`

	// Block describing the minimum and maximum number of network interfaces. Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	NetworkInterfaceCount []NetworkInterfaceCountParameters `json:"networkInterfaceCount,omitempty" tf:"network_interface_count,omitempty"`

	// Price protection threshold for On-Demand Instances. This is the maximum you’ll pay for an On-Demand Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 20.
	// +kubebuilder:validation:Optional
	OnDemandMaxPricePercentageOverLowestPrice *float64 `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" tf:"on_demand_max_price_percentage_over_lowest_price,omitempty"`

	// Indicate whether instance types must support On-Demand Instance Hibernation, either true or false. Default is false.
	// +kubebuilder:validation:Optional
	RequireHibernateSupport *bool `json:"requireHibernateSupport,omitempty" tf:"require_hibernate_support,omitempty"`

	// Price protection threshold for Spot Instances. This is the maximum you’ll pay for a Spot Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 100.
	// +kubebuilder:validation:Optional
	SpotMaxPricePercentageOverLowestPrice *float64 `json:"spotMaxPricePercentageOverLowestPrice,omitempty" tf:"spot_max_price_percentage_over_lowest_price,omitempty"`

	// Block describing the minimum and maximum total local storage (GB). Default is no minimum or maximum.
	// +kubebuilder:validation:Optional
	TotalLocalStorageGb []TotalLocalStorageGbParameters `json:"totalLocalStorageGb,omitempty" tf:"total_local_storage_gb,omitempty"`

	// Block describing the minimum and maximum number of vCPUs. Default is no maximum.
	// +kubebuilder:validation:Optional
	VcpuCount []VcpuCountParameters `json:"vcpuCount,omitempty" tf:"vcpu_count,omitempty"`
}

type InstanceReusePolicyObservation struct {
}

type InstanceReusePolicyParameters struct {

	// Whether instances in the Auto Scaling group can be returned to the warm pool on scale in.
	// +kubebuilder:validation:Optional
	ReuseOnScaleIn *bool `json:"reuseOnScaleIn,omitempty" tf:"reuse_on_scale_in,omitempty"`
}

type InstancesDistributionObservation struct {
}

type InstancesDistributionParameters struct {

	// Strategy to use when launching on-demand instances. Valid values: prioritized. Default: prioritized.
	// +kubebuilder:validation:Optional
	OnDemandAllocationStrategy *string `json:"onDemandAllocationStrategy,omitempty" tf:"on_demand_allocation_strategy,omitempty"`

	// Absolute minimum amount of desired capacity that must be fulfilled by on-demand instances. Default: 0.
	// +kubebuilder:validation:Optional
	OnDemandBaseCapacity *float64 `json:"onDemandBaseCapacity,omitempty" tf:"on_demand_base_capacity,omitempty"`

	// Percentage split between on-demand and Spot instances above the base on-demand capacity. Default: 100.
	// +kubebuilder:validation:Optional
	OnDemandPercentageAboveBaseCapacity *float64 `json:"onDemandPercentageAboveBaseCapacity,omitempty" tf:"on_demand_percentage_above_base_capacity,omitempty"`

	// How to allocate capacity across the Spot pools. Valid values: lowest-price, capacity-optimized, capacity-optimized-prioritized, and price-capacity-optimized. Default: lowest-price.
	// +kubebuilder:validation:Optional
	SpotAllocationStrategy *string `json:"spotAllocationStrategy,omitempty" tf:"spot_allocation_strategy,omitempty"`

	// Number of Spot pools per availability zone to allocate capacity. EC2 Auto Scaling selects the cheapest Spot pools and evenly allocates Spot capacity across the number of Spot pools that you specify. Only available with spot_allocation_strategy set to lowest-price. Otherwise it must be set to 0, if it has been defined before. Default: 2.
	// +kubebuilder:validation:Optional
	SpotInstancePools *float64 `json:"spotInstancePools,omitempty" tf:"spot_instance_pools,omitempty"`

	// Maximum price per unit hour that the user is willing to pay for the Spot instances. Default: an empty string which means the on-demand price.
	// +kubebuilder:validation:Optional
	SpotMaxPrice *string `json:"spotMaxPrice,omitempty" tf:"spot_max_price,omitempty"`
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// ID of the launch template. Conflicts with name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.LaunchTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a LaunchTemplate in ec2 to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a LaunchTemplate in ec2 to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Name of the launch template. Conflicts with id.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Template version. Can be version number, $Latest, or $Default. (Default: $Default).
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type LaunchTemplateSpecificationObservation struct {
}

type LaunchTemplateSpecificationParameters struct {

	// ID of the launch template. Conflicts with launch_template_name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.LaunchTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// Reference to a LaunchTemplate in ec2 to populate launchTemplateId.
	// +kubebuilder:validation:Optional
	LaunchTemplateIDRef *v1.Reference `json:"launchTemplateIdRef,omitempty" tf:"-"`

	// Selector for a LaunchTemplate in ec2 to populate launchTemplateId.
	// +kubebuilder:validation:Optional
	LaunchTemplateIDSelector *v1.Selector `json:"launchTemplateIdSelector,omitempty" tf:"-"`

	// Name of the launch template. Conflicts with launch_template_id.
	// +kubebuilder:validation:Optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// Template version. Can be version number, $Latest, or $Default. (Default: $Default).
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MemoryGibPerVcpuObservation struct {
}

type MemoryGibPerVcpuParameters struct {

	// Maximum.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// Minimum.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type MemoryMibObservation struct {
}

type MemoryMibParameters struct {

	// Maximum.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// Minimum.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type MixedInstancesPolicyLaunchTemplateObservation struct {
}

type MixedInstancesPolicyLaunchTemplateParameters struct {

	// Nested argument defines the Launch Template. Defined below.
	// +kubebuilder:validation:Required
	LaunchTemplateSpecification []LaunchTemplateSpecificationParameters `json:"launchTemplateSpecification" tf:"launch_template_specification,omitempty"`

	// List of nested arguments provides the ability to specify multiple instance types. This will override the same parameter in the launch template. For on-demand instances, Auto Scaling considers the order of preference of instance types to launch based on the order specified in the overrides list. Defined below.
	// +kubebuilder:validation:Optional
	Override []OverrideParameters `json:"override,omitempty" tf:"override,omitempty"`
}

type MixedInstancesPolicyObservation struct {
}

type MixedInstancesPolicyParameters struct {

	// Nested argument containing settings on how to mix on-demand and Spot instances in the Auto Scaling group. Defined below.
	// +kubebuilder:validation:Optional
	InstancesDistribution []InstancesDistributionParameters `json:"instancesDistribution,omitempty" tf:"instances_distribution,omitempty"`

	// Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
	// +kubebuilder:validation:Required
	LaunchTemplate []MixedInstancesPolicyLaunchTemplateParameters `json:"launchTemplate" tf:"launch_template,omitempty"`
}

type NetworkInterfaceCountObservation struct {
}

type NetworkInterfaceCountParameters struct {

	// Maximum.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// Minimum.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type OverrideLaunchTemplateSpecificationObservation struct {
}

type OverrideLaunchTemplateSpecificationParameters struct {

	// ID of the launch template. Conflicts with launch_template_name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.LaunchTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LaunchTemplateID *string `json:"launchTemplateId,omitempty" tf:"launch_template_id,omitempty"`

	// Reference to a LaunchTemplate in ec2 to populate launchTemplateId.
	// +kubebuilder:validation:Optional
	LaunchTemplateIDRef *v1.Reference `json:"launchTemplateIdRef,omitempty" tf:"-"`

	// Selector for a LaunchTemplate in ec2 to populate launchTemplateId.
	// +kubebuilder:validation:Optional
	LaunchTemplateIDSelector *v1.Selector `json:"launchTemplateIdSelector,omitempty" tf:"-"`

	// Name of the launch template. Conflicts with launch_template_id.
	// +kubebuilder:validation:Optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name,omitempty"`

	// Template version. Can be version number, $Latest, or $Default. (Default: $Default).
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OverrideObservation struct {
}

type OverrideParameters struct {

	// Override the instance type in the Launch Template with instance types that satisfy the requirements.
	// +kubebuilder:validation:Optional
	InstanceRequirements []InstanceRequirementsParameters `json:"instanceRequirements,omitempty" tf:"instance_requirements,omitempty"`

	// Override the instance type in the Launch Template.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Nested argument defines the Launch Template. Defined below.
	// +kubebuilder:validation:Optional
	LaunchTemplateSpecification []OverrideLaunchTemplateSpecificationParameters `json:"launchTemplateSpecification,omitempty" tf:"launch_template_specification,omitempty"`

	// Number of capacity units, which gives the instance type a proportional weight to other instance types.
	// +kubebuilder:validation:Optional
	WeightedCapacity *string `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type PreferencesObservation struct {
}

type PreferencesParameters struct {

	// Number of seconds to wait after a checkpoint. Defaults to 3600.
	// +kubebuilder:validation:Optional
	CheckpointDelay *string `json:"checkpointDelay,omitempty" tf:"checkpoint_delay,omitempty"`

	// List of percentages for each checkpoint. Values must be unique and in ascending order. To replace all instances, the final number must be 100.
	// +kubebuilder:validation:Optional
	CheckpointPercentages []*float64 `json:"checkpointPercentages,omitempty" tf:"checkpoint_percentages,omitempty"`

	// Number of seconds until a newly launched instance is configured and ready to use. Default behavior is to use the Auto Scaling Group's health check grace period.
	// +kubebuilder:validation:Optional
	InstanceWarmup *string `json:"instanceWarmup,omitempty" tf:"instance_warmup,omitempty"`

	// Amount of capacity in the Auto Scaling group that must remain healthy during an instance refresh to allow the operation to continue, as a percentage of the desired capacity of the Auto Scaling group. Defaults to 90.
	// +kubebuilder:validation:Optional
	MinHealthyPercentage *float64 `json:"minHealthyPercentage,omitempty" tf:"min_healthy_percentage,omitempty"`

	// Replace instances that already have your desired configuration. Defaults to false.
	// +kubebuilder:validation:Optional
	SkipMatching *bool `json:"skipMatching,omitempty" tf:"skip_matching,omitempty"`
}

type TagObservation struct {
}

type TagParameters struct {

	// Key
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Enables propagation of the tag to
	// Amazon EC2 instances launched via this ASG
	// +kubebuilder:validation:Required
	PropagateAtLaunch *bool `json:"propagateAtLaunch" tf:"propagate_at_launch,omitempty"`

	// Value
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type TotalLocalStorageGbObservation struct {
}

type TotalLocalStorageGbParameters struct {

	// Maximum.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// Minimum.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type VcpuCountObservation struct {
}

type VcpuCountParameters struct {

	// Maximum.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// Minimum.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type WarmPoolObservation struct {
}

type WarmPoolParameters struct {

	// Whether instances in the Auto Scaling group can be returned to the warm pool on scale in. The default is to terminate instances in the Auto Scaling group when the group scales in.
	// +kubebuilder:validation:Optional
	InstanceReusePolicy []InstanceReusePolicyParameters `json:"instanceReusePolicy,omitempty" tf:"instance_reuse_policy,omitempty"`

	// Total maximum number of instances that are allowed to be in the warm pool or in any state except Terminated for the Auto Scaling group.
	// +kubebuilder:validation:Optional
	MaxGroupPreparedCapacity *float64 `json:"maxGroupPreparedCapacity,omitempty" tf:"max_group_prepared_capacity,omitempty"`

	// Minimum size of the Auto Scaling Group.
	// (See also Waiting for Capacity below.)
	// +kubebuilder:validation:Optional
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Sets the instance state to transition to after the lifecycle hooks finish. Valid values are: Stopped (default), Running or Hibernated.
	// +kubebuilder:validation:Optional
	PoolState *string `json:"poolState,omitempty" tf:"pool_state,omitempty"`
}

// AutoscalingGroupSpec defines the desired state of AutoscalingGroup
type AutoscalingGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoscalingGroupParameters `json:"forProvider"`
}

// AutoscalingGroupStatus defines the observed state of AutoscalingGroup.
type AutoscalingGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoscalingGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingGroup is the Schema for the AutoscalingGroups API. Provides an Auto Scaling Group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AutoscalingGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingGroupSpec   `json:"spec"`
	Status            AutoscalingGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingGroupList contains a list of AutoscalingGroups
type AutoscalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingGroup `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingGroup_Kind             = "AutoscalingGroup"
	AutoscalingGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoscalingGroup_Kind}.String()
	AutoscalingGroup_KindAPIVersion   = AutoscalingGroup_Kind + "." + CRDGroupVersion.String()
	AutoscalingGroup_GroupVersionKind = CRDGroupVersion.WithKind(AutoscalingGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoscalingGroup{}, &AutoscalingGroupList{})
}
