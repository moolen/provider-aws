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

type DestinationOptionsInitParameters struct {

	// The format for the flow log. Default value: plain-text. Valid values: plain-text, parquet.
	FileFormat *string `json:"fileFormat,omitempty" tf:"file_format,omitempty"`

	// Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3. Default value: false.
	HiveCompatiblePartitions *bool `json:"hiveCompatiblePartitions,omitempty" tf:"hive_compatible_partitions,omitempty"`

	// Indicates whether to partition the flow log per hour. This reduces the cost and response time for queries. Default value: false.
	PerHourPartition *bool `json:"perHourPartition,omitempty" tf:"per_hour_partition,omitempty"`
}

type DestinationOptionsObservation struct {

	// The format for the flow log. Default value: plain-text. Valid values: plain-text, parquet.
	FileFormat *string `json:"fileFormat,omitempty" tf:"file_format,omitempty"`

	// Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3. Default value: false.
	HiveCompatiblePartitions *bool `json:"hiveCompatiblePartitions,omitempty" tf:"hive_compatible_partitions,omitempty"`

	// Indicates whether to partition the flow log per hour. This reduces the cost and response time for queries. Default value: false.
	PerHourPartition *bool `json:"perHourPartition,omitempty" tf:"per_hour_partition,omitempty"`
}

type DestinationOptionsParameters struct {

	// The format for the flow log. Default value: plain-text. Valid values: plain-text, parquet.
	// +kubebuilder:validation:Optional
	FileFormat *string `json:"fileFormat,omitempty" tf:"file_format,omitempty"`

	// Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3. Default value: false.
	// +kubebuilder:validation:Optional
	HiveCompatiblePartitions *bool `json:"hiveCompatiblePartitions,omitempty" tf:"hive_compatible_partitions,omitempty"`

	// Indicates whether to partition the flow log per hour. This reduces the cost and response time for queries. Default value: false.
	// +kubebuilder:validation:Optional
	PerHourPartition *bool `json:"perHourPartition,omitempty" tf:"per_hour_partition,omitempty"`
}

type FlowLogInitParameters struct {

	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole *string `json:"deliverCrossAccountRole,omitempty" tf:"deliver_cross_account_role,omitempty"`

	// Describes the destination options for a flow log. More details below.
	DestinationOptions []DestinationOptionsInitParameters `json:"destinationOptions,omitempty" tf:"destination_options,omitempty"`

	// Elastic Network Interface ID to attach to
	EniID *string `json:"eniId,omitempty" tf:"eni_id,omitempty"`

	// The type of the logging destination. Valid values: cloud-watch-logs, s3, kinesis-data-firehose. Default: cloud-watch-logs.
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`

	// Deprecated: Use log_destination instead. The name of the CloudWatch log group. Either log_group_name or log_destination must be set.
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: 60 seconds (1 minute) or 600 seconds (10
	// minutes). Default: 600. When transit_gateway_id or transit_gateway_attachment_id is specified, max_aggregation_interval must be 60 seconds (1 minute).
	MaxAggregationInterval *float64 `json:"maxAggregationInterval,omitempty" tf:"max_aggregation_interval,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of traffic to capture. Valid values: ACCEPT,REJECT, ALL.
	TrafficType *string `json:"trafficType,omitempty" tf:"traffic_type,omitempty"`

	// Transit Gateway Attachment ID to attach to
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Transit Gateway ID to attach to
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`
}

type FlowLogObservation struct {

	// The ARN of the Flow Log.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole *string `json:"deliverCrossAccountRole,omitempty" tf:"deliver_cross_account_role,omitempty"`

	// Describes the destination options for a flow log. More details below.
	DestinationOptions []DestinationOptionsObservation `json:"destinationOptions,omitempty" tf:"destination_options,omitempty"`

	// Elastic Network Interface ID to attach to
	EniID *string `json:"eniId,omitempty" tf:"eni_id,omitempty"`

	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// The Flow Log ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the logging destination. Either log_destination or log_group_name must be set.
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// The type of the logging destination. Valid values: cloud-watch-logs, s3, kinesis-data-firehose. Default: cloud-watch-logs.
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`

	// Deprecated: Use log_destination instead. The name of the CloudWatch log group. Either log_group_name or log_destination must be set.
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: 60 seconds (1 minute) or 600 seconds (10
	// minutes). Default: 600. When transit_gateway_id or transit_gateway_attachment_id is specified, max_aggregation_interval must be 60 seconds (1 minute).
	MaxAggregationInterval *float64 `json:"maxAggregationInterval,omitempty" tf:"max_aggregation_interval,omitempty"`

	// Subnet ID to attach to
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of traffic to capture. Valid values: ACCEPT,REJECT, ALL.
	TrafficType *string `json:"trafficType,omitempty" tf:"traffic_type,omitempty"`

	// Transit Gateway Attachment ID to attach to
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Transit Gateway ID to attach to
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// VPC ID to attach to
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type FlowLogParameters struct {

	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	// +kubebuilder:validation:Optional
	DeliverCrossAccountRole *string `json:"deliverCrossAccountRole,omitempty" tf:"deliver_cross_account_role,omitempty"`

	// Describes the destination options for a flow log. More details below.
	// +kubebuilder:validation:Optional
	DestinationOptions []DestinationOptionsParameters `json:"destinationOptions,omitempty" tf:"destination_options,omitempty"`

	// Elastic Network Interface ID to attach to
	// +kubebuilder:validation:Optional
	EniID *string `json:"eniId,omitempty" tf:"eni_id,omitempty"`

	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// The ARN of the logging destination. Either log_destination or log_group_name must be set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1.Group
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// Reference to a Group in cloudwatchlogs to populate logDestination.
	// +kubebuilder:validation:Optional
	LogDestinationRef *v1.Reference `json:"logDestinationRef,omitempty" tf:"-"`

	// Selector for a Group in cloudwatchlogs to populate logDestination.
	// +kubebuilder:validation:Optional
	LogDestinationSelector *v1.Selector `json:"logDestinationSelector,omitempty" tf:"-"`

	// The type of the logging destination. Valid values: cloud-watch-logs, s3, kinesis-data-firehose. Default: cloud-watch-logs.
	// +kubebuilder:validation:Optional
	LogDestinationType *string `json:"logDestinationType,omitempty" tf:"log_destination_type,omitempty"`

	// The fields to include in the flow log record, in the order in which they should appear.
	// +kubebuilder:validation:Optional
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`

	// Deprecated: Use log_destination instead. The name of the CloudWatch log group. Either log_group_name or log_destination must be set.
	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	// The maximum interval of time
	// during which a flow of packets is captured and aggregated into a flow
	// log record. Valid Values: 60 seconds (1 minute) or 600 seconds (10
	// minutes). Default: 600. When transit_gateway_id or transit_gateway_attachment_id is specified, max_aggregation_interval must be 60 seconds (1 minute).
	// +kubebuilder:validation:Optional
	MaxAggregationInterval *float64 `json:"maxAggregationInterval,omitempty" tf:"max_aggregation_interval,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Subnet ID to attach to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of traffic to capture. Valid values: ACCEPT,REJECT, ALL.
	// +kubebuilder:validation:Optional
	TrafficType *string `json:"trafficType,omitempty" tf:"traffic_type,omitempty"`

	// Transit Gateway Attachment ID to attach to
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Transit Gateway ID to attach to
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// VPC ID to attach to
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

// FlowLogSpec defines the desired state of FlowLog
type FlowLogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlowLogParameters `json:"forProvider"`
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
	InitProvider FlowLogInitParameters `json:"initProvider,omitempty"`
}

// FlowLogStatus defines the observed state of FlowLog.
type FlowLogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlowLogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlowLog is the Schema for the FlowLogs API. Provides a VPC/Subnet/ENI Flow Log
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FlowLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlowLogSpec   `json:"spec"`
	Status            FlowLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlowLogList contains a list of FlowLogs
type FlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlowLog `json:"items"`
}

// Repository type metadata.
var (
	FlowLog_Kind             = "FlowLog"
	FlowLog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlowLog_Kind}.String()
	FlowLog_KindAPIVersion   = FlowLog_Kind + "." + CRDGroupVersion.String()
	FlowLog_GroupVersionKind = CRDGroupVersion.WithKind(FlowLog_Kind)
)

func init() {
	SchemeBuilder.Register(&FlowLog{}, &FlowLogList{})
}
