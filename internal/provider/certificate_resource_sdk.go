// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *CertificateResourceModel) ToSharedCertificate(ctx context.Context) (*shared.Certificate, diag.Diagnostics) {
	var diags diag.Diagnostics

	var cert string
	cert = r.Cert.ValueString()

	certAlt := new(string)
	if !r.CertAlt.IsUnknown() && !r.CertAlt.IsNull() {
		*certAlt = r.CertAlt.ValueString()
	} else {
		certAlt = nil
	}
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var key string
	key = r.Key.ValueString()

	keyAlt := new(string)
	if !r.KeyAlt.IsUnknown() && !r.KeyAlt.IsNull() {
		*keyAlt = r.KeyAlt.ValueString()
	} else {
		keyAlt = nil
	}
	var snis []string
	if r.Snis != nil {
		snis = make([]string, 0, len(r.Snis))
		for _, snisItem := range r.Snis {
			snis = append(snis, snisItem.ValueString())
		}
	}
	tags := make([]string, 0, len(r.Tags))
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	out := shared.Certificate{
		Cert:      cert,
		CertAlt:   certAlt,
		CreatedAt: createdAt,
		ID:        id,
		Key:       key,
		KeyAlt:    keyAlt,
		Snis:      snis,
		Tags:      tags,
		UpdatedAt: updatedAt,
	}

	return &out, diags
}

func (r *CertificateResourceModel) ToOperationsUpsertCertificateRequest(ctx context.Context) (*operations.UpsertCertificateRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var certificateID string
	certificateID = r.ID.ValueString()

	certificate, certificateDiags := r.ToSharedCertificate(ctx)
	diags.Append(certificateDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpsertCertificateRequest{
		CertificateID: certificateID,
		Certificate:   *certificate,
	}

	return &out, diags
}

func (r *CertificateResourceModel) ToOperationsGetCertificateRequest(ctx context.Context) (*operations.GetCertificateRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var certificateID string
	certificateID = r.ID.ValueString()

	out := operations.GetCertificateRequest{
		CertificateID: certificateID,
	}

	return &out, diags
}

func (r *CertificateResourceModel) ToOperationsDeleteCertificateRequest(ctx context.Context) (*operations.DeleteCertificateRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var certificateID string
	certificateID = r.ID.ValueString()

	out := operations.DeleteCertificateRequest{
		CertificateID: certificateID,
	}

	return &out, diags
}

func (r *CertificateResourceModel) RefreshFromSharedCertificate(ctx context.Context, resp *shared.Certificate) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Cert = types.StringValue(resp.Cert)
		r.CertAlt = types.StringPointerValue(resp.CertAlt)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringValue(resp.Key)
		r.KeyAlt = types.StringPointerValue(resp.KeyAlt)
		if resp.Snis != nil {
			r.Snis = make([]types.String, 0, len(resp.Snis))
			for _, v := range resp.Snis {
				r.Snis = append(r.Snis, types.StringValue(v))
			}
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}

	return diags
}
