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

type TagOptionResourceAssociationInitParameters struct {
}

type TagOptionResourceAssociationObservation struct {

	// Identifier of the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the resource.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Creation time of the resource.
	ResourceCreatedTime *string `json:"resourceCreatedTime,omitempty" tf:"resource_created_time,omitempty"`

	// Description of the resource.
	ResourceDescription *string `json:"resourceDescription,omitempty" tf:"resource_description,omitempty"`

	// Resource identifier.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Description of the resource.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	// Tag Option identifier.
	TagOptionID *string `json:"tagOptionId,omitempty" tf:"tag_option_id,omitempty"`
}

type TagOptionResourceAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Resource identifier.
	// +crossplane:generate:reference:type=Product
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Product to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Product to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// Tag Option identifier.
	// +crossplane:generate:reference:type=TagOption
	// +kubebuilder:validation:Optional
	TagOptionID *string `json:"tagOptionId,omitempty" tf:"tag_option_id,omitempty"`

	// Reference to a TagOption to populate tagOptionId.
	// +kubebuilder:validation:Optional
	TagOptionIDRef *v1.Reference `json:"tagOptionIdRef,omitempty" tf:"-"`

	// Selector for a TagOption to populate tagOptionId.
	// +kubebuilder:validation:Optional
	TagOptionIDSelector *v1.Selector `json:"tagOptionIdSelector,omitempty" tf:"-"`
}

// TagOptionResourceAssociationSpec defines the desired state of TagOptionResourceAssociation
type TagOptionResourceAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagOptionResourceAssociationParameters `json:"forProvider"`
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
	InitProvider TagOptionResourceAssociationInitParameters `json:"initProvider,omitempty"`
}

// TagOptionResourceAssociationStatus defines the observed state of TagOptionResourceAssociation.
type TagOptionResourceAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagOptionResourceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TagOptionResourceAssociation is the Schema for the TagOptionResourceAssociations API. Manages a Service Catalog Tag Option Resource Association
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TagOptionResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagOptionResourceAssociationSpec   `json:"spec"`
	Status            TagOptionResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagOptionResourceAssociationList contains a list of TagOptionResourceAssociations
type TagOptionResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagOptionResourceAssociation `json:"items"`
}

// Repository type metadata.
var (
	TagOptionResourceAssociation_Kind             = "TagOptionResourceAssociation"
	TagOptionResourceAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TagOptionResourceAssociation_Kind}.String()
	TagOptionResourceAssociation_KindAPIVersion   = TagOptionResourceAssociation_Kind + "." + CRDGroupVersion.String()
	TagOptionResourceAssociation_GroupVersionKind = CRDGroupVersion.WithKind(TagOptionResourceAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&TagOptionResourceAssociation{}, &TagOptionResourceAssociationList{})
}
