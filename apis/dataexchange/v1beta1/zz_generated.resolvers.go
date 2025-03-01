/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Revision.
func (mg *Revision) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataSetID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DataSetIDRef,
		Selector:     mg.Spec.ForProvider.DataSetIDSelector,
		To: reference.To{
			List:    &DataSetList{},
			Managed: &DataSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DataSetID")
	}
	mg.Spec.ForProvider.DataSetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DataSetIDRef = rsp.ResolvedReference

	return nil
}
