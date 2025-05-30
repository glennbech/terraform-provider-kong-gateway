// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type LogglyPluginConfig struct {
	ClientErrorsSeverity types.String            `tfsdk:"client_errors_severity"`
	CustomFieldsByLua    map[string]types.String `tfsdk:"custom_fields_by_lua"`
	Host                 types.String            `tfsdk:"host"`
	Key                  types.String            `tfsdk:"key"`
	LogLevel             types.String            `tfsdk:"log_level"`
	Port                 types.Int64             `tfsdk:"port"`
	ServerErrorsSeverity types.String            `tfsdk:"server_errors_severity"`
	SuccessfulSeverity   types.String            `tfsdk:"successful_severity"`
	Tags                 []types.String          `tfsdk:"tags"`
	Timeout              types.Float64           `tfsdk:"timeout"`
}
