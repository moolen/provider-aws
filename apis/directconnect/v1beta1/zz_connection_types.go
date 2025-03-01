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

type ConnectionInitParameters struct {

	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	Bandwidth *string `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are no_encrypt, should_encrypt, and must_encrypt.
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode,omitempty"`

	// The AWS Direct Connect location where the connection is located. See DescribeLocations for the list of AWS Direct Connect locations. Use locationCode.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the connection.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the service provider associated with the connection.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See MACsec prerequisites for more information about MAC Security (MACsec) prerequisites. Default value: false.
	RequestMacsec *bool `json:"requestMacsec,omitempty" tf:"request_macsec,omitempty"`

	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConnectionObservation struct {

	// The ARN of the connection.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	Bandwidth *string `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are no_encrypt, should_encrypt, and must_encrypt.
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode,omitempty"`

	// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy *string `json:"hasLogicalRedundancy,omitempty" tf:"has_logical_redundancy,omitempty"`

	// The ID of the connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Boolean value representing if jumbo frames have been enabled for this connection.
	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`

	// The AWS Direct Connect location where the connection is located. See DescribeLocations for the list of AWS Direct Connect locations. Use locationCode.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Boolean value indicating whether the connection supports MAC Security (MACsec).
	MacsecCapable *bool `json:"macsecCapable,omitempty" tf:"macsec_capable,omitempty"`

	// The name of the connection.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the AWS account that owns the connection.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName *string `json:"partnerName,omitempty" tf:"partner_name,omitempty"`

	// The MAC Security (MACsec) port link status of the connection.
	PortEncryptionStatus *string `json:"portEncryptionStatus,omitempty" tf:"port_encryption_status,omitempty"`

	// The name of the service provider associated with the connection.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See MACsec prerequisites for more information about MAC Security (MACsec) prerequisites. Default value: false.
	RequestMacsec *bool `json:"requestMacsec,omitempty" tf:"request_macsec,omitempty"`

	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VLAN ID.
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type ConnectionParameters struct {

	// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	// +kubebuilder:validation:Optional
	Bandwidth *string `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are no_encrypt, should_encrypt, and must_encrypt.
	// +kubebuilder:validation:Optional
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode,omitempty"`

	// The AWS Direct Connect location where the connection is located. See DescribeLocations for the list of AWS Direct Connect locations. Use locationCode.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the connection.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the service provider associated with the connection.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See MACsec prerequisites for more information about MAC Security (MACsec) prerequisites. Default value: false.
	// +kubebuilder:validation:Optional
	RequestMacsec *bool `json:"requestMacsec,omitempty" tf:"request_macsec,omitempty"`

	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
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
	InitProvider ConnectionInitParameters `json:"initProvider,omitempty"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API. Provides a Connection of Direct Connect.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bandwidth) || (has(self.initProvider) && has(self.initProvider.bandwidth))",message="spec.forProvider.bandwidth is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ConnectionSpec   `json:"spec"`
	Status ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
