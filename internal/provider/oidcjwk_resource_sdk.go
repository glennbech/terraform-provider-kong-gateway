// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *OidcJwkResourceModel) ToSharedOidcJwk(ctx context.Context) (*shared.OidcJwk, diag.Diagnostics) {
	var diags diag.Diagnostics

	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var jwks *shared.Jwks
	if r.Jwks != nil {
		keys := make([]shared.Keys, 0, len(r.Jwks.Keys))
		for _, keysItem := range r.Jwks.Keys {
			alg := new(string)
			if !keysItem.Alg.IsUnknown() && !keysItem.Alg.IsNull() {
				*alg = keysItem.Alg.ValueString()
			} else {
				alg = nil
			}
			crv := new(string)
			if !keysItem.Crv.IsUnknown() && !keysItem.Crv.IsNull() {
				*crv = keysItem.Crv.ValueString()
			} else {
				crv = nil
			}
			d := new(string)
			if !keysItem.D.IsUnknown() && !keysItem.D.IsNull() {
				*d = keysItem.D.ValueString()
			} else {
				d = nil
			}
			dp := new(string)
			if !keysItem.Dp.IsUnknown() && !keysItem.Dp.IsNull() {
				*dp = keysItem.Dp.ValueString()
			} else {
				dp = nil
			}
			dq := new(string)
			if !keysItem.Dq.IsUnknown() && !keysItem.Dq.IsNull() {
				*dq = keysItem.Dq.ValueString()
			} else {
				dq = nil
			}
			e := new(string)
			if !keysItem.E.IsUnknown() && !keysItem.E.IsNull() {
				*e = keysItem.E.ValueString()
			} else {
				e = nil
			}
			issuer := new(string)
			if !keysItem.Issuer.IsUnknown() && !keysItem.Issuer.IsNull() {
				*issuer = keysItem.Issuer.ValueString()
			} else {
				issuer = nil
			}
			k := new(string)
			if !keysItem.K.IsUnknown() && !keysItem.K.IsNull() {
				*k = keysItem.K.ValueString()
			} else {
				k = nil
			}
			keyOps := make([]string, 0, len(keysItem.KeyOps))
			for _, keyOpsItem := range keysItem.KeyOps {
				keyOps = append(keyOps, keyOpsItem.ValueString())
			}
			kid := new(string)
			if !keysItem.Kid.IsUnknown() && !keysItem.Kid.IsNull() {
				*kid = keysItem.Kid.ValueString()
			} else {
				kid = nil
			}
			kty := new(string)
			if !keysItem.Kty.IsUnknown() && !keysItem.Kty.IsNull() {
				*kty = keysItem.Kty.ValueString()
			} else {
				kty = nil
			}
			n := new(string)
			if !keysItem.N.IsUnknown() && !keysItem.N.IsNull() {
				*n = keysItem.N.ValueString()
			} else {
				n = nil
			}
			oth := new(string)
			if !keysItem.Oth.IsUnknown() && !keysItem.Oth.IsNull() {
				*oth = keysItem.Oth.ValueString()
			} else {
				oth = nil
			}
			p := new(string)
			if !keysItem.P.IsUnknown() && !keysItem.P.IsNull() {
				*p = keysItem.P.ValueString()
			} else {
				p = nil
			}
			q := new(string)
			if !keysItem.Q.IsUnknown() && !keysItem.Q.IsNull() {
				*q = keysItem.Q.ValueString()
			} else {
				q = nil
			}
			qi := new(string)
			if !keysItem.Qi.IsUnknown() && !keysItem.Qi.IsNull() {
				*qi = keysItem.Qi.ValueString()
			} else {
				qi = nil
			}
			r1 := new(string)
			if !keysItem.R.IsUnknown() && !keysItem.R.IsNull() {
				*r1 = keysItem.R.ValueString()
			} else {
				r1 = nil
			}
			t := new(string)
			if !keysItem.T.IsUnknown() && !keysItem.T.IsNull() {
				*t = keysItem.T.ValueString()
			} else {
				t = nil
			}
			use := new(string)
			if !keysItem.Use.IsUnknown() && !keysItem.Use.IsNull() {
				*use = keysItem.Use.ValueString()
			} else {
				use = nil
			}
			x := new(string)
			if !keysItem.X.IsUnknown() && !keysItem.X.IsNull() {
				*x = keysItem.X.ValueString()
			} else {
				x = nil
			}
			x5c := make([]string, 0, len(keysItem.X5c))
			for _, x5cItem := range keysItem.X5c {
				x5c = append(x5c, x5cItem.ValueString())
			}
			x5t := new(string)
			if !keysItem.X5t.IsUnknown() && !keysItem.X5t.IsNull() {
				*x5t = keysItem.X5t.ValueString()
			} else {
				x5t = nil
			}
			x5tNumberS256 := new(string)
			if !keysItem.X5tNumberS256.IsUnknown() && !keysItem.X5tNumberS256.IsNull() {
				*x5tNumberS256 = keysItem.X5tNumberS256.ValueString()
			} else {
				x5tNumberS256 = nil
			}
			x5u := new(string)
			if !keysItem.X5u.IsUnknown() && !keysItem.X5u.IsNull() {
				*x5u = keysItem.X5u.ValueString()
			} else {
				x5u = nil
			}
			y := new(string)
			if !keysItem.Y.IsUnknown() && !keysItem.Y.IsNull() {
				*y = keysItem.Y.ValueString()
			} else {
				y = nil
			}
			keys = append(keys, shared.Keys{
				Alg:           alg,
				Crv:           crv,
				D:             d,
				Dp:            dp,
				Dq:            dq,
				E:             e,
				Issuer:        issuer,
				K:             k,
				KeyOps:        keyOps,
				Kid:           kid,
				Kty:           kty,
				N:             n,
				Oth:           oth,
				P:             p,
				Q:             q,
				Qi:            qi,
				R:             r1,
				T:             t,
				Use:           use,
				X:             x,
				X5c:           x5c,
				X5t:           x5t,
				X5tNumberS256: x5tNumberS256,
				X5u:           x5u,
				Y:             y,
			})
		}
		jwks = &shared.Jwks{
			Keys: keys,
		}
	}
	out := shared.OidcJwk{
		ID:   id,
		Jwks: jwks,
	}

	return &out, diags
}

func (r *OidcJwkResourceModel) ToOperationsUpsertOicJwkRequest(ctx context.Context) (*operations.UpsertOicJwkRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var oidcJwkID string
	oidcJwkID = r.ID.ValueString()

	oidcJwk, oidcJwkDiags := r.ToSharedOidcJwk(ctx)
	diags.Append(oidcJwkDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpsertOicJwkRequest{
		OidcJwkID: oidcJwkID,
		OidcJwk:   *oidcJwk,
	}

	return &out, diags
}

func (r *OidcJwkResourceModel) ToOperationsGetOicJwkRequest(ctx context.Context) (*operations.GetOicJwkRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var oidcJwkID string
	oidcJwkID = r.ID.ValueString()

	out := operations.GetOicJwkRequest{
		OidcJwkID: oidcJwkID,
	}

	return &out, diags
}

func (r *OidcJwkResourceModel) ToOperationsDeleteOicJwkRequest(ctx context.Context) (*operations.DeleteOicJwkRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var oidcJwkID string
	oidcJwkID = r.ID.ValueString()

	out := operations.DeleteOicJwkRequest{
		OidcJwkID: oidcJwkID,
	}

	return &out, diags
}

func (r *OidcJwkResourceModel) RefreshFromSharedOidcJwk(ctx context.Context, resp *shared.OidcJwk) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.ID = types.StringPointerValue(resp.ID)
		if resp.Jwks == nil {
			r.Jwks = nil
		} else {
			r.Jwks = &tfTypes.Jwks{}
			r.Jwks.Keys = []tfTypes.Keys{}
			if len(r.Jwks.Keys) > len(resp.Jwks.Keys) {
				r.Jwks.Keys = r.Jwks.Keys[:len(resp.Jwks.Keys)]
			}
			for keysCount, keysItem := range resp.Jwks.Keys {
				var keys tfTypes.Keys
				keys.Alg = types.StringPointerValue(keysItem.Alg)
				keys.Crv = types.StringPointerValue(keysItem.Crv)
				keys.D = types.StringPointerValue(keysItem.D)
				keys.Dp = types.StringPointerValue(keysItem.Dp)
				keys.Dq = types.StringPointerValue(keysItem.Dq)
				keys.E = types.StringPointerValue(keysItem.E)
				keys.Issuer = types.StringPointerValue(keysItem.Issuer)
				keys.K = types.StringPointerValue(keysItem.K)
				keys.KeyOps = make([]types.String, 0, len(keysItem.KeyOps))
				for _, v := range keysItem.KeyOps {
					keys.KeyOps = append(keys.KeyOps, types.StringValue(v))
				}
				keys.Kid = types.StringPointerValue(keysItem.Kid)
				keys.Kty = types.StringPointerValue(keysItem.Kty)
				keys.N = types.StringPointerValue(keysItem.N)
				keys.Oth = types.StringPointerValue(keysItem.Oth)
				keys.P = types.StringPointerValue(keysItem.P)
				keys.Q = types.StringPointerValue(keysItem.Q)
				keys.Qi = types.StringPointerValue(keysItem.Qi)
				keys.R = types.StringPointerValue(keysItem.R)
				keys.T = types.StringPointerValue(keysItem.T)
				keys.Use = types.StringPointerValue(keysItem.Use)
				keys.X = types.StringPointerValue(keysItem.X)
				keys.X5c = make([]types.String, 0, len(keysItem.X5c))
				for _, v := range keysItem.X5c {
					keys.X5c = append(keys.X5c, types.StringValue(v))
				}
				keys.X5t = types.StringPointerValue(keysItem.X5t)
				keys.X5tNumberS256 = types.StringPointerValue(keysItem.X5tNumberS256)
				keys.X5u = types.StringPointerValue(keysItem.X5u)
				keys.Y = types.StringPointerValue(keysItem.Y)
				if keysCount+1 > len(r.Jwks.Keys) {
					r.Jwks.Keys = append(r.Jwks.Keys, keys)
				} else {
					r.Jwks.Keys[keysCount].Alg = keys.Alg
					r.Jwks.Keys[keysCount].Crv = keys.Crv
					r.Jwks.Keys[keysCount].D = keys.D
					r.Jwks.Keys[keysCount].Dp = keys.Dp
					r.Jwks.Keys[keysCount].Dq = keys.Dq
					r.Jwks.Keys[keysCount].E = keys.E
					r.Jwks.Keys[keysCount].Issuer = keys.Issuer
					r.Jwks.Keys[keysCount].K = keys.K
					r.Jwks.Keys[keysCount].KeyOps = keys.KeyOps
					r.Jwks.Keys[keysCount].Kid = keys.Kid
					r.Jwks.Keys[keysCount].Kty = keys.Kty
					r.Jwks.Keys[keysCount].N = keys.N
					r.Jwks.Keys[keysCount].Oth = keys.Oth
					r.Jwks.Keys[keysCount].P = keys.P
					r.Jwks.Keys[keysCount].Q = keys.Q
					r.Jwks.Keys[keysCount].Qi = keys.Qi
					r.Jwks.Keys[keysCount].R = keys.R
					r.Jwks.Keys[keysCount].T = keys.T
					r.Jwks.Keys[keysCount].Use = keys.Use
					r.Jwks.Keys[keysCount].X = keys.X
					r.Jwks.Keys[keysCount].X5c = keys.X5c
					r.Jwks.Keys[keysCount].X5t = keys.X5t
					r.Jwks.Keys[keysCount].X5tNumberS256 = keys.X5tNumberS256
					r.Jwks.Keys[keysCount].X5u = keys.X5u
					r.Jwks.Keys[keysCount].Y = keys.Y
				}
			}
		}
	}

	return diags
}
