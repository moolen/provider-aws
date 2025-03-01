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

type DestinationPolicyInitParameters struct {

	// The policy document. This is a JSON formatted string.
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// Specify true if you are updating an existing destination policy to grant permission to an organization ID instead of granting permission to individual AWS accounts.
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update,omitempty"`
}

type DestinationPolicyObservation struct {

	// The policy document. This is a JSON formatted string.
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// Specify true if you are updating an existing destination policy to grant permission to an organization ID instead of granting permission to individual AWS accounts.
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DestinationPolicyParameters struct {

	// The policy document. This is a JSON formatted string.
	// +kubebuilder:validation:Optional
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// Specify true if you are updating an existing destination policy to grant permission to an organization ID instead of granting permission to individual AWS accounts.
	// +kubebuilder:validation:Optional
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DestinationPolicySpec defines the desired state of DestinationPolicy
type DestinationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DestinationPolicyParameters `json:"forProvider"`
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
	InitProvider DestinationPolicyInitParameters `json:"initProvider,omitempty"`
}

// DestinationPolicyStatus defines the observed state of DestinationPolicy.
type DestinationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DestinationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DestinationPolicy is the Schema for the DestinationPolicys API. Provides a CloudWatch Logs destination policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DestinationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessPolicy) || (has(self.initProvider) && has(self.initProvider.accessPolicy))",message="spec.forProvider.accessPolicy is a required parameter"
	Spec   DestinationPolicySpec   `json:"spec"`
	Status DestinationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DestinationPolicyList contains a list of DestinationPolicys
type DestinationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DestinationPolicy `json:"items"`
}

// Repository type metadata.
var (
	DestinationPolicy_Kind             = "DestinationPolicy"
	DestinationPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DestinationPolicy_Kind}.String()
	DestinationPolicy_KindAPIVersion   = DestinationPolicy_Kind + "." + CRDGroupVersion.String()
	DestinationPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DestinationPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DestinationPolicy{}, &DestinationPolicyList{})
}
