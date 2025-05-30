// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type KeyAuthPluginConfig struct {
	Anonymous       types.String     `tfsdk:"anonymous"`
	HideCredentials types.Bool       `tfsdk:"hide_credentials"`
	IdentityRealms  []IdentityRealms `tfsdk:"identity_realms"`
	KeyInBody       types.Bool       `tfsdk:"key_in_body"`
	KeyInHeader     types.Bool       `tfsdk:"key_in_header"`
	KeyInQuery      types.Bool       `tfsdk:"key_in_query"`
	KeyNames        []types.String   `tfsdk:"key_names"`
	Realm           types.String     `tfsdk:"realm"`
	RunOnPreflight  types.Bool       `tfsdk:"run_on_preflight"`
}
