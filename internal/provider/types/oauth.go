// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Oauth struct {
	Audience      []types.String          `tfsdk:"audience"`
	ClientID      types.String            `tfsdk:"client_id"`
	ClientSecret  types.String            `tfsdk:"client_secret"`
	GrantType     types.String            `tfsdk:"grant_type"`
	Password      types.String            `tfsdk:"password"`
	Scopes        []types.String          `tfsdk:"scopes"`
	TokenEndpoint types.String            `tfsdk:"token_endpoint"`
	TokenHeaders  map[string]types.String `tfsdk:"token_headers"`
	TokenPostArgs map[string]types.String `tfsdk:"token_post_args"`
	Username      types.String            `tfsdk:"username"`
}
