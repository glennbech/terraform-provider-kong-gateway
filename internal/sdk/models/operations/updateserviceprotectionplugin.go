// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateServiceprotectionPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                string                         `pathParam:"style=simple,explode=false,name=PluginId"`
	ServiceProtectionPlugin shared.ServiceProtectionPlugin `request:"mediaType=application/json"`
}

func (o *UpdateServiceprotectionPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateServiceprotectionPluginRequest) GetServiceProtectionPlugin() shared.ServiceProtectionPlugin {
	if o == nil {
		return shared.ServiceProtectionPlugin{}
	}
	return o.ServiceProtectionPlugin
}

type UpdateServiceprotectionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ServiceProtection plugin
	ServiceProtectionPlugin *shared.ServiceProtectionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateServiceprotectionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateServiceprotectionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateServiceprotectionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateServiceprotectionPluginResponse) GetServiceProtectionPlugin() *shared.ServiceProtectionPlugin {
	if o == nil {
		return nil
	}
	return o.ServiceProtectionPlugin
}

func (o *UpdateServiceprotectionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
