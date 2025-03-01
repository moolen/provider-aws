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

type ActiveReceiptRuleSetInitParameters struct {

	// The name of the rule set
	RuleSetName *string `json:"ruleSetName,omitempty" tf:"rule_set_name,omitempty"`
}

type ActiveReceiptRuleSetObservation struct {

	// The SES receipt rule set ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The SES receipt rule set name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the rule set
	RuleSetName *string `json:"ruleSetName,omitempty" tf:"rule_set_name,omitempty"`
}

type ActiveReceiptRuleSetParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the rule set
	// +kubebuilder:validation:Optional
	RuleSetName *string `json:"ruleSetName,omitempty" tf:"rule_set_name,omitempty"`
}

// ActiveReceiptRuleSetSpec defines the desired state of ActiveReceiptRuleSet
type ActiveReceiptRuleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActiveReceiptRuleSetParameters `json:"forProvider"`
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
	InitProvider ActiveReceiptRuleSetInitParameters `json:"initProvider,omitempty"`
}

// ActiveReceiptRuleSetStatus defines the observed state of ActiveReceiptRuleSet.
type ActiveReceiptRuleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActiveReceiptRuleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveReceiptRuleSet is the Schema for the ActiveReceiptRuleSets API. Provides a resource to designate the active SES receipt rule set
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ActiveReceiptRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleSetName) || (has(self.initProvider) && has(self.initProvider.ruleSetName))",message="spec.forProvider.ruleSetName is a required parameter"
	Spec   ActiveReceiptRuleSetSpec   `json:"spec"`
	Status ActiveReceiptRuleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveReceiptRuleSetList contains a list of ActiveReceiptRuleSets
type ActiveReceiptRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveReceiptRuleSet `json:"items"`
}

// Repository type metadata.
var (
	ActiveReceiptRuleSet_Kind             = "ActiveReceiptRuleSet"
	ActiveReceiptRuleSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActiveReceiptRuleSet_Kind}.String()
	ActiveReceiptRuleSet_KindAPIVersion   = ActiveReceiptRuleSet_Kind + "." + CRDGroupVersion.String()
	ActiveReceiptRuleSet_GroupVersionKind = CRDGroupVersion.WithKind(ActiveReceiptRuleSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ActiveReceiptRuleSet{}, &ActiveReceiptRuleSetList{})
}
