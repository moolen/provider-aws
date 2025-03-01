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

type EnumerationValueInitParameters struct {

	// Additional values related to the slot type value. Each item must be less than or equal to 140 characters in length.
	Synonyms []*string `json:"synonyms,omitempty" tf:"synonyms,omitempty"`

	// The value of the slot type. Must be less than or equal to 140 characters in length.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnumerationValueObservation struct {

	// Additional values related to the slot type value. Each item must be less than or equal to 140 characters in length.
	Synonyms []*string `json:"synonyms,omitempty" tf:"synonyms,omitempty"`

	// The value of the slot type. Must be less than or equal to 140 characters in length.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnumerationValueParameters struct {

	// Additional values related to the slot type value. Each item must be less than or equal to 140 characters in length.
	// +kubebuilder:validation:Optional
	Synonyms []*string `json:"synonyms,omitempty" tf:"synonyms,omitempty"`

	// The value of the slot type. Must be less than or equal to 140 characters in length.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type SlotTypeInitParameters struct {

	// Determines if a new slot type version is created when the initial resource is created and on each
	// update. Defaults to false.
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version,omitempty"`

	// A description of the slot type. Must be less than or equal to 200 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of EnumerationValue objects that defines the values that
	// the slot type can take. Each value can have a list of synonyms, which are additional values that help
	// train the machine learning model about the values that it resolves for a slot. Attributes are
	// documented under enumeration_value.
	EnumerationValue []EnumerationValueInitParameters `json:"enumerationValue,omitempty" tf:"enumeration_value,omitempty"`

	// Determines the slot resolution strategy that Amazon Lex
	// uses to return slot type values. ORIGINAL_VALUE returns the value entered by the user if the user
	// value is similar to the slot value. TOP_RESOLUTION returns the first value in the resolution list
	// if there is a resolution list for the slot, otherwise null is returned. Defaults to ORIGINAL_VALUE.
	ValueSelectionStrategy *string `json:"valueSelectionStrategy,omitempty" tf:"value_selection_strategy,omitempty"`
}

type SlotTypeObservation struct {

	// Checksum identifying the version of the slot type that was created. The checksum is
	// not included as an argument because the resource will add it automatically when updating the slot type.
	Checksum *string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	// Determines if a new slot type version is created when the initial resource is created and on each
	// update. Defaults to false.
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version,omitempty"`

	// The date when the slot type version was created.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// A description of the slot type. Must be less than or equal to 200 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of EnumerationValue objects that defines the values that
	// the slot type can take. Each value can have a list of synonyms, which are additional values that help
	// train the machine learning model about the values that it resolves for a slot. Attributes are
	// documented under enumeration_value.
	EnumerationValue []EnumerationValueObservation `json:"enumerationValue,omitempty" tf:"enumeration_value,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date when the $LATEST version of this slot type was updated.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	// Determines the slot resolution strategy that Amazon Lex
	// uses to return slot type values. ORIGINAL_VALUE returns the value entered by the user if the user
	// value is similar to the slot value. TOP_RESOLUTION returns the first value in the resolution list
	// if there is a resolution list for the slot, otherwise null is returned. Defaults to ORIGINAL_VALUE.
	ValueSelectionStrategy *string `json:"valueSelectionStrategy,omitempty" tf:"value_selection_strategy,omitempty"`

	// The version of the slot type.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SlotTypeParameters struct {

	// Determines if a new slot type version is created when the initial resource is created and on each
	// update. Defaults to false.
	// +kubebuilder:validation:Optional
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version,omitempty"`

	// A description of the slot type. Must be less than or equal to 200 characters in length.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of EnumerationValue objects that defines the values that
	// the slot type can take. Each value can have a list of synonyms, which are additional values that help
	// train the machine learning model about the values that it resolves for a slot. Attributes are
	// documented under enumeration_value.
	// +kubebuilder:validation:Optional
	EnumerationValue []EnumerationValueParameters `json:"enumerationValue,omitempty" tf:"enumeration_value,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Determines the slot resolution strategy that Amazon Lex
	// uses to return slot type values. ORIGINAL_VALUE returns the value entered by the user if the user
	// value is similar to the slot value. TOP_RESOLUTION returns the first value in the resolution list
	// if there is a resolution list for the slot, otherwise null is returned. Defaults to ORIGINAL_VALUE.
	// +kubebuilder:validation:Optional
	ValueSelectionStrategy *string `json:"valueSelectionStrategy,omitempty" tf:"value_selection_strategy,omitempty"`
}

// SlotTypeSpec defines the desired state of SlotType
type SlotTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SlotTypeParameters `json:"forProvider"`
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
	InitProvider SlotTypeInitParameters `json:"initProvider,omitempty"`
}

// SlotTypeStatus defines the observed state of SlotType.
type SlotTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SlotTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SlotType is the Schema for the SlotTypes API. Provides details about a specific Amazon Lex Slot Type
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SlotType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enumerationValue) || (has(self.initProvider) && has(self.initProvider.enumerationValue))",message="spec.forProvider.enumerationValue is a required parameter"
	Spec   SlotTypeSpec   `json:"spec"`
	Status SlotTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SlotTypeList contains a list of SlotTypes
type SlotTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SlotType `json:"items"`
}

// Repository type metadata.
var (
	SlotType_Kind             = "SlotType"
	SlotType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SlotType_Kind}.String()
	SlotType_KindAPIVersion   = SlotType_Kind + "." + CRDGroupVersion.String()
	SlotType_GroupVersionKind = CRDGroupVersion.WithKind(SlotType_Kind)
)

func init() {
	SchemeBuilder.Register(&SlotType{}, &SlotTypeList{})
}
