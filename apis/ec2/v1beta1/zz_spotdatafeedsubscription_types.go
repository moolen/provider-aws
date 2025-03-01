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

type SpotDatafeedSubscriptionInitParameters struct {

	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Path of folder inside bucket to place spot pricing data.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type SpotDatafeedSubscriptionObservation struct {

	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Path of folder inside bucket to place spot pricing data.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type SpotDatafeedSubscriptionParameters struct {

	// The Amazon S3 bucket in which to store the Spot instance data feed.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Path of folder inside bucket to place spot pricing data.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// SpotDatafeedSubscriptionSpec defines the desired state of SpotDatafeedSubscription
type SpotDatafeedSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpotDatafeedSubscriptionParameters `json:"forProvider"`
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
	InitProvider SpotDatafeedSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// SpotDatafeedSubscriptionStatus defines the observed state of SpotDatafeedSubscription.
type SpotDatafeedSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpotDatafeedSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpotDatafeedSubscription is the Schema for the SpotDatafeedSubscriptions API. Provides a Spot Datafeed Subscription resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SpotDatafeedSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	Spec   SpotDatafeedSubscriptionSpec   `json:"spec"`
	Status SpotDatafeedSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotDatafeedSubscriptionList contains a list of SpotDatafeedSubscriptions
type SpotDatafeedSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotDatafeedSubscription `json:"items"`
}

// Repository type metadata.
var (
	SpotDatafeedSubscription_Kind             = "SpotDatafeedSubscription"
	SpotDatafeedSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpotDatafeedSubscription_Kind}.String()
	SpotDatafeedSubscription_KindAPIVersion   = SpotDatafeedSubscription_Kind + "." + CRDGroupVersion.String()
	SpotDatafeedSubscription_GroupVersionKind = CRDGroupVersion.WithKind(SpotDatafeedSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&SpotDatafeedSubscription{}, &SpotDatafeedSubscriptionList{})
}
