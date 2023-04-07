/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type AcceleratorCountRequest struct {
	Max *int64 `json:"max,omitempty"`

	Min *int64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type AcceleratorTotalMemoryMiBRequest struct {
	Max *int64 `json:"max,omitempty"`

	Min *int64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type Activity struct {
	ActivityID *string `json:"activityID,omitempty"`

	AutoScalingGroupARN *string `json:"autoScalingGroupARN,omitempty"`

	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	Cause *string `json:"cause,omitempty"`

	Description *string `json:"description,omitempty"`

	Details *string `json:"details,omitempty"`

	EndTime *metav1.Time `json:"endTime,omitempty"`

	StartTime *metav1.Time `json:"startTime,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`
}

// +kubebuilder:skipversion
type AdjustmentType struct {
	AdjustmentType *string `json:"adjustmentType,omitempty"`
}

// +kubebuilder:skipversion
type Alarm struct {
	AlarmARN *string `json:"alarmARN,omitempty"`

	AlarmName *string `json:"alarmName,omitempty"`
}

// +kubebuilder:skipversion
type BaselineEBSBandwidthMbpsRequest struct {
	Max *int64 `json:"max,omitempty"`

	Min *int64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type BlockDeviceMapping struct {
	DeviceName *string `json:"deviceName,omitempty"`

	VirtualName *string `json:"virtualName,omitempty"`
}

// +kubebuilder:skipversion
type DesiredConfiguration struct {
	// Describes the launch template and the version of the launch template that
	// Amazon EC2 Auto Scaling uses to launch Amazon EC2 instances. For more information
	// about launch templates, see Launch templates (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchTemplates.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate,omitempty"`
	// Use this structure to launch multiple instance types and On-Demand Instances
	// and Spot Instances within a single Auto Scaling group.
	//
	// A mixed instances policy contains information that Amazon EC2 Auto Scaling
	// can use to launch instances and help optimize your costs. For more information,
	// see Auto Scaling groups with multiple instance types and purchase options
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	MixedInstancesPolicy *MixedInstancesPolicy `json:"mixedInstancesPolicy,omitempty"`
}

// +kubebuilder:skipversion
type EBS struct {
	SnapshotID *string `json:"snapshotID,omitempty"`
}

// +kubebuilder:skipversion
type EnabledMetric struct {
	Granularity *string `json:"granularity,omitempty"`

	Metric *string `json:"metric,omitempty"`
}

// +kubebuilder:skipversion
type FailedScheduledUpdateGroupActionRequest struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`

	ScheduledActionName *string `json:"scheduledActionName,omitempty"`
}

// +kubebuilder:skipversion
type Filter struct {
	Name *string `json:"name,omitempty"`

	Values []*string `json:"values,omitempty"`
}

// +kubebuilder:skipversion
type Group struct {
	AutoScalingGroupARN *string `json:"autoScalingGroupARN,omitempty"`

	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	AvailabilityZones []*string `json:"availabilityZones,omitempty"`

	CapacityRebalance *bool `json:"capacityRebalance,omitempty"`

	Context *string `json:"context,omitempty"`

	CreatedTime *metav1.Time `json:"createdTime,omitempty"`

	DefaultCooldown *int64 `json:"defaultCooldown,omitempty"`

	DefaultInstanceWarmup *int64 `json:"defaultInstanceWarmup,omitempty"`

	DesiredCapacity *int64 `json:"desiredCapacity,omitempty"`

	DesiredCapacityType *string `json:"desiredCapacityType,omitempty"`

	EnabledMetrics []*EnabledMetric `json:"enabledMetrics,omitempty"`

	HealthCheckGracePeriod *int64 `json:"healthCheckGracePeriod,omitempty"`

	HealthCheckType *string `json:"healthCheckType,omitempty"`

	Instances []*Instance `json:"instances,omitempty"`

	LaunchConfigurationName *string `json:"launchConfigurationName,omitempty"`
	// Describes the launch template and the version of the launch template that
	// Amazon EC2 Auto Scaling uses to launch Amazon EC2 instances. For more information
	// about launch templates, see Launch templates (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchTemplates.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate,omitempty"`

	LoadBalancerNames []*string `json:"loadBalancerNames,omitempty"`

	MaxInstanceLifetime *int64 `json:"maxInstanceLifetime,omitempty"`

	MaxSize *int64 `json:"maxSize,omitempty"`

	MinSize *int64 `json:"minSize,omitempty"`
	// Use this structure to launch multiple instance types and On-Demand Instances
	// and Spot Instances within a single Auto Scaling group.
	//
	// A mixed instances policy contains information that Amazon EC2 Auto Scaling
	// can use to launch instances and help optimize your costs. For more information,
	// see Auto Scaling groups with multiple instance types and purchase options
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	MixedInstancesPolicy *MixedInstancesPolicy `json:"mixedInstancesPolicy,omitempty"`

	NewInstancesProtectedFromScaleIn *bool `json:"newInstancesProtectedFromScaleIn,omitempty"`

	PlacementGroup *string `json:"placementGroup,omitempty"`

	PredictedCapacity *int64 `json:"predictedCapacity,omitempty"`

	ServiceLinkedRoleARN *string `json:"serviceLinkedRoleARN,omitempty"`

	Status *string `json:"status,omitempty"`

	SuspendedProcesses []*SuspendedProcess `json:"suspendedProcesses,omitempty"`

	Tags []*TagDescription `json:"tags,omitempty"`

	TargetGroupARNs []*string `json:"targetGroupARNs,omitempty"`

	TerminationPolicies []*string `json:"terminationPolicies,omitempty"`

	TrafficSources []*TrafficSourceIdentifier `json:"trafficSources,omitempty"`

	VPCZoneIdentifier *string `json:"vpcZoneIdentifier,omitempty"`
	// Describes a warm pool configuration.
	WarmPoolConfiguration *WarmPoolConfiguration `json:"warmPoolConfiguration,omitempty"`

	WarmPoolSize *int64 `json:"warmPoolSize,omitempty"`
}

// +kubebuilder:skipversion
type Instance struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	HealthStatus *string `json:"healthStatus,omitempty"`

	InstanceID *string `json:"instanceID,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	LaunchConfigurationName *string `json:"launchConfigurationName,omitempty"`
	// Describes the launch template and the version of the launch template that
	// Amazon EC2 Auto Scaling uses to launch Amazon EC2 instances. For more information
	// about launch templates, see Launch templates (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchTemplates.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate,omitempty"`

	LifecycleState *string `json:"lifecycleState,omitempty"`

	ProtectedFromScaleIn *bool `json:"protectedFromScaleIn,omitempty"`

	WeightedCapacity *string `json:"weightedCapacity,omitempty"`
}

// +kubebuilder:skipversion
type InstanceDetails struct {
	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	HealthStatus *string `json:"healthStatus,omitempty"`

	InstanceID *string `json:"instanceID,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	LaunchConfigurationName *string `json:"launchConfigurationName,omitempty"`
	// Describes the launch template and the version of the launch template that
	// Amazon EC2 Auto Scaling uses to launch Amazon EC2 instances. For more information
	// about launch templates, see Launch templates (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchTemplates.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate,omitempty"`

	LifecycleState *string `json:"lifecycleState,omitempty"`

	ProtectedFromScaleIn *bool `json:"protectedFromScaleIn,omitempty"`

	WeightedCapacity *string `json:"weightedCapacity,omitempty"`
}

// +kubebuilder:skipversion
type InstanceRefresh struct {
	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	EndTime *metav1.Time `json:"endTime,omitempty"`

	InstanceRefreshID *string `json:"instanceRefreshID,omitempty"`

	StartTime *metav1.Time `json:"startTime,omitempty"`

	StatusReason *string `json:"statusReason,omitempty"`
}

// +kubebuilder:skipversion
type InstanceRequirements struct {
	// Specifies the minimum and maximum for the AcceleratorCount object when you
	// specify InstanceRequirements for an Auto Scaling group.
	AcceleratorCount *AcceleratorCountRequest `json:"acceleratorCount,omitempty"`

	AcceleratorManufacturers []*string `json:"acceleratorManufacturers,omitempty"`

	AcceleratorNames []*string `json:"acceleratorNames,omitempty"`
	// Specifies the minimum and maximum for the AcceleratorTotalMemoryMiB object
	// when you specify InstanceRequirements for an Auto Scaling group.
	AcceleratorTotalMemoryMiB *AcceleratorTotalMemoryMiBRequest `json:"acceleratorTotalMemoryMiB,omitempty"`

	AcceleratorTypes []*string `json:"acceleratorTypes,omitempty"`

	AllowedInstanceTypes []*string `json:"allowedInstanceTypes,omitempty"`

	BareMetal *string `json:"bareMetal,omitempty"`
	// Specifies the minimum and maximum for the BaselineEbsBandwidthMbps object
	// when you specify InstanceRequirements for an Auto Scaling group.
	BaselineEBSBandwidthMbps *BaselineEBSBandwidthMbpsRequest `json:"baselineEBSBandwidthMbps,omitempty"`

	BurstablePerformance *string `json:"burstablePerformance,omitempty"`

	CPUManufacturers []*string `json:"cpuManufacturers,omitempty"`

	ExcludedInstanceTypes []*string `json:"excludedInstanceTypes,omitempty"`

	InstanceGenerations []*string `json:"instanceGenerations,omitempty"`

	LocalStorage *string `json:"localStorage,omitempty"`

	LocalStorageTypes []*string `json:"localStorageTypes,omitempty"`
	// Specifies the minimum and maximum for the MemoryGiBPerVCpu object when you
	// specify InstanceRequirements for an Auto Scaling group.
	MemoryGiBPerVCPU *MemoryGiBPerVCPURequest `json:"memoryGiBPerVCPU,omitempty"`
	// Specifies the minimum and maximum for the MemoryMiB object when you specify
	// InstanceRequirements for an Auto Scaling group.
	MemoryMiB *MemoryMiBRequest `json:"memoryMiB,omitempty"`
	// Specifies the minimum and maximum for the NetworkBandwidthGbps object when
	// you specify InstanceRequirements for an Auto Scaling group.
	//
	// Setting the minimum bandwidth does not guarantee that your instance will
	// achieve the minimum bandwidth. Amazon EC2 will identify instance types that
	// support the specified minimum bandwidth, but the actual bandwidth of your
	// instance might go below the specified minimum at times. For more information,
	// see Available instance bandwidth (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-network-bandwidth.html#available-instance-bandwidth)
	// in the Amazon EC2 User Guide for Linux Instances.
	NetworkBandwidthGbps *NetworkBandwidthGbpsRequest `json:"networkBandwidthGbps,omitempty"`
	// Specifies the minimum and maximum for the NetworkInterfaceCount object when
	// you specify InstanceRequirements for an Auto Scaling group.
	NetworkInterfaceCount *NetworkInterfaceCountRequest `json:"networkInterfaceCount,omitempty"`

	OnDemandMaxPricePercentageOverLowestPrice *int64 `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty"`

	RequireHibernateSupport *bool `json:"requireHibernateSupport,omitempty"`

	SpotMaxPricePercentageOverLowestPrice *int64 `json:"spotMaxPricePercentageOverLowestPrice,omitempty"`
	// Specifies the minimum and maximum for the TotalLocalStorageGB object when
	// you specify InstanceRequirements for an Auto Scaling group.
	TotalLocalStorageGB *TotalLocalStorageGBRequest `json:"totalLocalStorageGB,omitempty"`
	// Specifies the minimum and maximum for the VCpuCount object when you specify
	// InstanceRequirements for an Auto Scaling group.
	VCPUCount *VCPUCountRequest `json:"vCPUCount,omitempty"`
}

// +kubebuilder:skipversion
type InstanceReusePolicy struct {
	ReuseOnScaleIn *bool `json:"reuseOnScaleIn,omitempty"`
}

// +kubebuilder:skipversion
type InstancesDistribution struct {
	OnDemandAllocationStrategy *string `json:"onDemandAllocationStrategy,omitempty"`

	OnDemandBaseCapacity *int64 `json:"onDemandBaseCapacity,omitempty"`

	OnDemandPercentageAboveBaseCapacity *int64 `json:"onDemandPercentageAboveBaseCapacity,omitempty"`

	SpotAllocationStrategy *string `json:"spotAllocationStrategy,omitempty"`

	SpotInstancePools *int64 `json:"spotInstancePools,omitempty"`

	SpotMaxPrice *string `json:"spotMaxPrice,omitempty"`
}

// +kubebuilder:skipversion
type LaunchConfiguration struct {
	ClassicLinkVPCID *string `json:"classicLinkVPCID,omitempty"`

	CreatedTime *metav1.Time `json:"createdTime,omitempty"`

	IAMInstanceProfile *string `json:"iamInstanceProfile,omitempty"`

	ImageID *string `json:"imageID,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	KernelID *string `json:"kernelID,omitempty"`

	KeyName *string `json:"keyName,omitempty"`

	LaunchConfigurationARN *string `json:"launchConfigurationARN,omitempty"`

	LaunchConfigurationName *string `json:"launchConfigurationName,omitempty"`

	RAMDiskID *string `json:"ramDiskID,omitempty"`
}

// +kubebuilder:skipversion
type LaunchTemplate struct {
	// Describes the launch template and the version of the launch template that
	// Amazon EC2 Auto Scaling uses to launch Amazon EC2 instances. For more information
	// about launch templates, see Launch templates (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchTemplates.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	LaunchTemplateSpecification *LaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty"`

	Overrides []*LaunchTemplateOverrides `json:"overrides,omitempty"`
}

// +kubebuilder:skipversion
type LaunchTemplateOverrides struct {
	// The attributes for the instance types for a mixed instances policy. Amazon
	// EC2 Auto Scaling uses your specified requirements to identify instance types.
	// Then, it uses your On-Demand and Spot allocation strategies to launch instances
	// from these instance types.
	//
	// When you specify multiple attributes, you get instance types that satisfy
	// all of the specified attributes. If you specify multiple values for an attribute,
	// you get instance types that satisfy any of the specified values.
	//
	// To limit the list of instance types from which Amazon EC2 Auto Scaling can
	// identify matching instance types, you can use one of the following parameters,
	// but not both in the same request:
	//
	//    * AllowedInstanceTypes - The instance types to include in the list. All
	//    other instance types are ignored, even if they match your specified attributes.
	//
	//    * ExcludedInstanceTypes - The instance types to exclude from the list,
	//    even if they match your specified attributes.
	//
	// You must specify VCpuCount and MemoryMiB. All other attributes are optional.
	// Any unspecified optional attribute is set to its default.
	//
	// For more information, see Creating an Auto Scaling group using attribute-based
	// instance type selection (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html)
	// in the Amazon EC2 Auto Scaling User Guide. For help determining which instance
	// types match your attributes before you apply them to your Auto Scaling group,
	// see Preview instance types with specified attributes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html#ec2fleet-get-instance-types-from-instance-requirements)
	// in the Amazon EC2 User Guide for Linux Instances.
	InstanceRequirements *InstanceRequirements `json:"instanceRequirements,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`
	// Describes the launch template and the version of the launch template that
	// Amazon EC2 Auto Scaling uses to launch Amazon EC2 instances. For more information
	// about launch templates, see Launch templates (https://docs.aws.amazon.com/autoscaling/ec2/userguide/LaunchTemplates.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	LaunchTemplateSpecification *LaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty"`

	WeightedCapacity *string `json:"weightedCapacity,omitempty"`
}

// +kubebuilder:skipversion
type LaunchTemplateSpecification struct {
	LaunchTemplateID *string `json:"launchTemplateID,omitempty"`

	LaunchTemplateName *string `json:"launchTemplateName,omitempty"`

	Version *string `json:"version,omitempty"`
}

// +kubebuilder:skipversion
type LifecycleHook struct {
	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	DefaultResult *string `json:"defaultResult,omitempty"`

	HeartbeatTimeout *int64 `json:"heartbeatTimeout,omitempty"`

	LifecycleHookName *string `json:"lifecycleHookName,omitempty"`

	LifecycleTransition *string `json:"lifecycleTransition,omitempty"`

	NotificationMetadata *string `json:"notificationMetadata,omitempty"`

	NotificationTargetARN *string `json:"notificationTargetARN,omitempty"`

	RoleARN *string `json:"roleARN,omitempty"`
}

// +kubebuilder:skipversion
type LifecycleHookSpecification struct {
	DefaultResult *string `json:"defaultResult,omitempty"`

	HeartbeatTimeout *int64 `json:"heartbeatTimeout,omitempty"`

	LifecycleHookName *string `json:"lifecycleHookName,omitempty"`

	LifecycleTransition *string `json:"lifecycleTransition,omitempty"`

	NotificationMetadata *string `json:"notificationMetadata,omitempty"`

	NotificationTargetARN *string `json:"notificationTargetARN,omitempty"`

	RoleARN *string `json:"roleARN,omitempty"`
}

// +kubebuilder:skipversion
type LoadBalancerState struct {
	LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type LoadBalancerTargetGroupState struct {
	LoadBalancerTargetGroupARN *string `json:"loadBalancerTargetGroupARN,omitempty"`

	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type MemoryGiBPerVCPURequest struct {
	Max *float64 `json:"max,omitempty"`

	Min *float64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type MemoryMiBRequest struct {
	Max *int64 `json:"max,omitempty"`

	Min *int64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type MetricCollectionType struct {
	Metric *string `json:"metric,omitempty"`
}

// +kubebuilder:skipversion
type MetricDataQuery struct {
	Expression *string `json:"expression,omitempty"`

	ID *string `json:"id,omitempty"`
}

// +kubebuilder:skipversion
type MetricGranularityType struct {
	Granularity *string `json:"granularity,omitempty"`
}

// +kubebuilder:skipversion
type MixedInstancesPolicy struct {
	// Use this structure to specify the distribution of On-Demand Instances and
	// Spot Instances and the allocation strategies used to fulfill On-Demand and
	// Spot capacities for a mixed instances policy.
	InstancesDistribution *InstancesDistribution `json:"instancesDistribution,omitempty"`
	// Use this structure to specify the launch templates and instance types (overrides)
	// for a mixed instances policy.
	LaunchTemplate *LaunchTemplate `json:"launchTemplate,omitempty"`
}

// +kubebuilder:skipversion
type NetworkBandwidthGbpsRequest struct {
	Max *float64 `json:"max,omitempty"`

	Min *float64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type NetworkInterfaceCountRequest struct {
	Max *int64 `json:"max,omitempty"`

	Min *int64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type NotificationConfiguration struct {
	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	NotificationType *string `json:"notificationType,omitempty"`

	TopicARN *string `json:"topicARN,omitempty"`
}

// +kubebuilder:skipversion
type PredefinedMetricSpecification struct {
	ResourceLabel *string `json:"resourceLabel,omitempty"`
}

// +kubebuilder:skipversion
type PredictiveScalingPredefinedLoadMetric struct {
	ResourceLabel *string `json:"resourceLabel,omitempty"`
}

// +kubebuilder:skipversion
type PredictiveScalingPredefinedMetricPair struct {
	ResourceLabel *string `json:"resourceLabel,omitempty"`
}

// +kubebuilder:skipversion
type PredictiveScalingPredefinedScalingMetric struct {
	ResourceLabel *string `json:"resourceLabel,omitempty"`
}

// +kubebuilder:skipversion
type ProcessType struct {
	ProcessName *string `json:"processName,omitempty"`
}

// +kubebuilder:skipversion
type ScalingPolicy struct {
	AdjustmentType *string `json:"adjustmentType,omitempty"`

	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	Cooldown *int64 `json:"cooldown,omitempty"`

	MetricAggregationType *string `json:"metricAggregationType,omitempty"`

	PolicyARN *string `json:"policyARN,omitempty"`

	PolicyName *string `json:"policyName,omitempty"`
}

// +kubebuilder:skipversion
type ScheduledUpdateGroupAction struct {
	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	DesiredCapacity *int64 `json:"desiredCapacity,omitempty"`

	EndTime *metav1.Time `json:"endTime,omitempty"`

	MaxSize *int64 `json:"maxSize,omitempty"`

	MinSize *int64 `json:"minSize,omitempty"`

	Recurrence *string `json:"recurrence,omitempty"`

	ScheduledActionARN *string `json:"scheduledActionARN,omitempty"`

	ScheduledActionName *string `json:"scheduledActionName,omitempty"`

	StartTime *metav1.Time `json:"startTime,omitempty"`

	Time *metav1.Time `json:"time,omitempty"`

	TimeZone *string `json:"timeZone,omitempty"`
}

// +kubebuilder:skipversion
type ScheduledUpdateGroupActionRequest struct {
	DesiredCapacity *int64 `json:"desiredCapacity,omitempty"`

	EndTime *metav1.Time `json:"endTime,omitempty"`

	MaxSize *int64 `json:"maxSize,omitempty"`

	MinSize *int64 `json:"minSize,omitempty"`

	Recurrence *string `json:"recurrence,omitempty"`

	ScheduledActionName *string `json:"scheduledActionName,omitempty"`

	StartTime *metav1.Time `json:"startTime,omitempty"`

	TimeZone *string `json:"timeZone,omitempty"`
}

// +kubebuilder:skipversion
type SuspendedProcess struct {
	ProcessName *string `json:"processName,omitempty"`

	SuspensionReason *string `json:"suspensionReason,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	PropagateAtLaunch *bool `json:"propagateAtLaunch,omitempty"`

	ResourceID *string `json:"resourceID,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type TagDescription struct {
	Key *string `json:"key,omitempty"`

	PropagateAtLaunch *bool `json:"propagateAtLaunch,omitempty"`

	ResourceID *string `json:"resourceID,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type TargetTrackingMetricDataQuery struct {
	Expression *string `json:"expression,omitempty"`

	ID *string `json:"id,omitempty"`
}

// +kubebuilder:skipversion
type TotalLocalStorageGBRequest struct {
	Max *float64 `json:"max,omitempty"`

	Min *float64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type TrafficSourceIdentifier struct {
	Identifier *string `json:"identifier,omitempty"`
}

// +kubebuilder:skipversion
type TrafficSourceState struct {
	State *string `json:"state,omitempty"`

	TrafficSource *string `json:"trafficSource,omitempty"`
}

// +kubebuilder:skipversion
type VCPUCountRequest struct {
	Max *int64 `json:"max,omitempty"`

	Min *int64 `json:"min,omitempty"`
}

// +kubebuilder:skipversion
type WarmPoolConfiguration struct {
	// Describes an instance reuse policy for a warm pool.
	//
	// For more information, see Warm pools for Amazon EC2 Auto Scaling (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceReusePolicy *InstanceReusePolicy `json:"instanceReusePolicy,omitempty"`

	MaxGroupPreparedCapacity *int64 `json:"maxGroupPreparedCapacity,omitempty"`

	MinSize *int64 `json:"minSize,omitempty"`

	PoolState *string `json:"poolState,omitempty"`

	Status *string `json:"status,omitempty"`
}