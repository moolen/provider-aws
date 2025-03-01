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

type CertificateInitParameters struct {

	// ARN of an ACM PCA
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// Certificate's PEM-formatted public key
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// Certificate's PEM-formatted chain
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// Domain name for which the certificate should be issued
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Amount of time to start automatic renewal process before expiration.
	// Has no effect if less than 60 days.
	// Represented by either
	// a subset of RFC 3339 duration supporting years, months, and days (e.g., P90D),
	// or a string such as 2160h.
	EarlyRenewalDuration *string `json:"earlyRenewalDuration,omitempty" tf:"early_renewal_duration,omitempty"`

	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See ACM Certificate characteristics for more details.
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`

	// Configuration block used to set certificate options. Detailed below.
	Options []OptionsInitParameters `json:"options,omitempty" tf:"options,omitempty"`

	// Set of domains that should be SANs in the issued certificate.
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Which method to use for validation.
	ValidationMethod *string `json:"validationMethod,omitempty" tf:"validation_method,omitempty"`

	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	ValidationOption []ValidationOptionInitParameters `json:"validationOption,omitempty" tf:"validation_option,omitempty"`
}

type CertificateObservation struct {

	// ARN of the certificate
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of an ACM PCA
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// Certificate's PEM-formatted public key
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// Certificate's PEM-formatted chain
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// Domain name for which the certificate should be issued
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Set of domain validation objects which can be used to complete certificate validation.
	// Can have more than one element, e.g., if SANs are defined.
	// Only set if DNS-validation was used.
	DomainValidationOptions []DomainValidationOptionsObservation `json:"domainValidationOptions,omitempty" tf:"domain_validation_options,omitempty"`

	// Amount of time to start automatic renewal process before expiration.
	// Has no effect if less than 60 days.
	// Represented by either
	// a subset of RFC 3339 duration supporting years, months, and days (e.g., P90D),
	// or a string such as 2160h.
	EarlyRenewalDuration *string `json:"earlyRenewalDuration,omitempty" tf:"early_renewal_duration,omitempty"`

	// ARN of the certificate
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See ACM Certificate characteristics for more details.
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`

	// Expiration date and time of the certificate.
	NotAfter *string `json:"notAfter,omitempty" tf:"not_after,omitempty"`

	// Start of the validity period of the certificate.
	NotBefore *string `json:"notBefore,omitempty" tf:"not_before,omitempty"`

	// Configuration block used to set certificate options. Detailed below.
	Options []OptionsObservation `json:"options,omitempty" tf:"options,omitempty"`

	// true if a Private certificate eligible for managed renewal is within the early_renewal_duration period.
	PendingRenewal *bool `json:"pendingRenewal,omitempty" tf:"pending_renewal,omitempty"`

	// Whether the certificate is eligible for managed renewal.
	RenewalEligibility *string `json:"renewalEligibility,omitempty" tf:"renewal_eligibility,omitempty"`

	// Contains information about the status of ACM's managed renewal for the certificate.
	RenewalSummary []RenewalSummaryObservation `json:"renewalSummary,omitempty" tf:"renewal_summary,omitempty"`

	// Status of the certificate.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Set of domains that should be SANs in the issued certificate.
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Source of the certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// List of addresses that received a validation email. Only set if EMAIL validation was used.
	ValidationEmails []*string `json:"validationEmails,omitempty" tf:"validation_emails,omitempty"`

	// Which method to use for validation.
	ValidationMethod *string `json:"validationMethod,omitempty" tf:"validation_method,omitempty"`

	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	ValidationOption []ValidationOptionObservation `json:"validationOption,omitempty" tf:"validation_option,omitempty"`
}

type CertificateParameters struct {

	// ARN of an ACM PCA
	// +kubebuilder:validation:Optional
	CertificateAuthorityArn *string `json:"certificateAuthorityArn,omitempty" tf:"certificate_authority_arn,omitempty"`

	// Certificate's PEM-formatted public key
	// +kubebuilder:validation:Optional
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// Certificate's PEM-formatted chain
	// +kubebuilder:validation:Optional
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// Domain name for which the certificate should be issued
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Amount of time to start automatic renewal process before expiration.
	// Has no effect if less than 60 days.
	// Represented by either
	// a subset of RFC 3339 duration supporting years, months, and days (e.g., P90D),
	// or a string such as 2160h.
	// +kubebuilder:validation:Optional
	EarlyRenewalDuration *string `json:"earlyRenewalDuration,omitempty" tf:"early_renewal_duration,omitempty"`

	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See ACM Certificate characteristics for more details.
	// +kubebuilder:validation:Optional
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`

	// Configuration block used to set certificate options. Detailed below.
	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// Certificate's PEM-formatted private key
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Set of domains that should be SANs in the issued certificate.
	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Which method to use for validation.
	// +kubebuilder:validation:Optional
	ValidationMethod *string `json:"validationMethod,omitempty" tf:"validation_method,omitempty"`

	// Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	// +kubebuilder:validation:Optional
	ValidationOption []ValidationOptionParameters `json:"validationOption,omitempty" tf:"validation_option,omitempty"`
}

type DomainValidationOptionsInitParameters struct {
}

type DomainValidationOptionsObservation struct {

	// Fully qualified domain name (FQDN) in the certificate.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The name of the DNS record to create to validate the certificate
	ResourceRecordName *string `json:"resourceRecordName,omitempty" tf:"resource_record_name,omitempty"`

	// The type of DNS record to create
	ResourceRecordType *string `json:"resourceRecordType,omitempty" tf:"resource_record_type,omitempty"`

	// The value the DNS record needs to have
	ResourceRecordValue *string `json:"resourceRecordValue,omitempty" tf:"resource_record_value,omitempty"`
}

type DomainValidationOptionsParameters struct {
}

type OptionsInitParameters struct {

	// Whether certificate details should be added to a certificate transparency log. Valid values are ENABLED or DISABLED. See https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency for more details.
	CertificateTransparencyLoggingPreference *string `json:"certificateTransparencyLoggingPreference,omitempty" tf:"certificate_transparency_logging_preference,omitempty"`
}

type OptionsObservation struct {

	// Whether certificate details should be added to a certificate transparency log. Valid values are ENABLED or DISABLED. See https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency for more details.
	CertificateTransparencyLoggingPreference *string `json:"certificateTransparencyLoggingPreference,omitempty" tf:"certificate_transparency_logging_preference,omitempty"`
}

type OptionsParameters struct {

	// Whether certificate details should be added to a certificate transparency log. Valid values are ENABLED or DISABLED. See https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency for more details.
	// +kubebuilder:validation:Optional
	CertificateTransparencyLoggingPreference *string `json:"certificateTransparencyLoggingPreference,omitempty" tf:"certificate_transparency_logging_preference,omitempty"`
}

type RenewalSummaryInitParameters struct {
}

type RenewalSummaryObservation struct {

	// The status of ACM's managed renewal of the certificate
	RenewalStatus *string `json:"renewalStatus,omitempty" tf:"renewal_status,omitempty"`

	// The reason that a renewal request was unsuccessful or is pending
	RenewalStatusReason *string `json:"renewalStatusReason,omitempty" tf:"renewal_status_reason,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type RenewalSummaryParameters struct {
}

type ValidationOptionInitParameters struct {

	// Fully qualified domain name (FQDN) in the certificate.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Domain name that you want ACM to use to send you validation emails. This domain name is the suffix of the email addresses that you want ACM to use. This must be the same as the domain_name value or a superdomain of the domain_name value. For example, if you request a certificate for "testing.example.com", you can specify "example.com" for this value.
	ValidationDomain *string `json:"validationDomain,omitempty" tf:"validation_domain,omitempty"`
}

type ValidationOptionObservation struct {

	// Fully qualified domain name (FQDN) in the certificate.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Domain name that you want ACM to use to send you validation emails. This domain name is the suffix of the email addresses that you want ACM to use. This must be the same as the domain_name value or a superdomain of the domain_name value. For example, if you request a certificate for "testing.example.com", you can specify "example.com" for this value.
	ValidationDomain *string `json:"validationDomain,omitempty" tf:"validation_domain,omitempty"`
}

type ValidationOptionParameters struct {

	// Fully qualified domain name (FQDN) in the certificate.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// Domain name that you want ACM to use to send you validation emails. This domain name is the suffix of the email addresses that you want ACM to use. This must be the same as the domain_name value or a superdomain of the domain_name value. For example, if you request a certificate for "testing.example.com", you can specify "example.com" for this value.
	// +kubebuilder:validation:Optional
	ValidationDomain *string `json:"validationDomain" tf:"validation_domain,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
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
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API. Requests and manages a certificate from Amazon Certificate Manager (ACM).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
