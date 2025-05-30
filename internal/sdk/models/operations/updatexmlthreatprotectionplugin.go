// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateXmlthreatprotectionPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                  string                           `pathParam:"style=simple,explode=false,name=PluginId"`
	XMLThreatProtectionPlugin shared.XMLThreatProtectionPlugin `request:"mediaType=application/json"`
}

func (o *UpdateXmlthreatprotectionPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateXmlthreatprotectionPluginRequest) GetXMLThreatProtectionPlugin() shared.XMLThreatProtectionPlugin {
	if o == nil {
		return shared.XMLThreatProtectionPlugin{}
	}
	return o.XMLThreatProtectionPlugin
}

type UpdateXmlthreatprotectionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// XmlThreatProtection plugin
	XMLThreatProtectionPlugin *shared.XMLThreatProtectionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateXmlthreatprotectionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateXmlthreatprotectionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateXmlthreatprotectionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateXmlthreatprotectionPluginResponse) GetXMLThreatProtectionPlugin() *shared.XMLThreatProtectionPlugin {
	if o == nil {
		return nil
	}
	return o.XMLThreatProtectionPlugin
}

func (o *UpdateXmlthreatprotectionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
