// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type GetKafkaconsumePluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
}

func (o *GetKafkaconsumePluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

type GetKafkaconsumePluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// KafkaConsume plugin
	KafkaConsumePlugin *shared.KafkaConsumePlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetKafkaconsumePluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKafkaconsumePluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKafkaconsumePluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetKafkaconsumePluginResponse) GetKafkaConsumePlugin() *shared.KafkaConsumePlugin {
	if o == nil {
		return nil
	}
	return o.KafkaConsumePlugin
}

func (o *GetKafkaconsumePluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
