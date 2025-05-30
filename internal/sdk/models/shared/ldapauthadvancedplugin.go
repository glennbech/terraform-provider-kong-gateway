// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type LdapAuthAdvancedPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *LdapAuthAdvancedPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type LdapAuthAdvancedPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *LdapAuthAdvancedPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type LdapAuthAdvancedPluginOrdering struct {
	After  *LdapAuthAdvancedPluginAfter  `json:"after,omitempty"`
	Before *LdapAuthAdvancedPluginBefore `json:"before,omitempty"`
}

func (o *LdapAuthAdvancedPluginOrdering) GetAfter() *LdapAuthAdvancedPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *LdapAuthAdvancedPluginOrdering) GetBefore() *LdapAuthAdvancedPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type LdapAuthAdvancedPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *LdapAuthAdvancedPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LdapAuthAdvancedPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LdapAuthAdvancedPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

type LdapAuthAdvancedPluginConsumerBy string

const (
	LdapAuthAdvancedPluginConsumerByCustomID LdapAuthAdvancedPluginConsumerBy = "custom_id"
	LdapAuthAdvancedPluginConsumerByUsername LdapAuthAdvancedPluginConsumerBy = "username"
)

func (e LdapAuthAdvancedPluginConsumerBy) ToPointer() *LdapAuthAdvancedPluginConsumerBy {
	return &e
}
func (e *LdapAuthAdvancedPluginConsumerBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "custom_id":
		fallthrough
	case "username":
		*e = LdapAuthAdvancedPluginConsumerBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LdapAuthAdvancedPluginConsumerBy: %v", v)
	}
}

type LdapAuthAdvancedPluginConfig struct {
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request will fail with an authentication failure `4xx`. Note that this value must refer to the consumer `id` or `username` attribute, and **not** its `custom_id`.
	Anonymous *string `json:"anonymous,omitempty"`
	// Attribute to be used to search the user; e.g., "cn".
	Attribute *string `json:"attribute,omitempty"`
	// Base DN as the starting point for the search; e.g., 'dc=example,dc=com'.
	BaseDn *string `json:"base_dn,omitempty"`
	// The DN to bind to. Used to perform LDAP search of user. This `bind_dn` should have permissions to search for the user being authenticated.
	BindDn *string `json:"bind_dn,omitempty"`
	// Cache expiry time in seconds.
	CacheTTL *float64 `json:"cache_ttl,omitempty"`
	// Whether to authenticate consumers based on `username`, `custom_id`, or both.
	ConsumerBy []LdapAuthAdvancedPluginConsumerBy `json:"consumer_by,omitempty"`
	// Whether consumer mapping is optional. If `consumer_optional=true`, the plugin will not attempt to associate a consumer with the LDAP authenticated user.
	ConsumerOptional *bool `json:"consumer_optional,omitempty"`
	// Sets a distinguished name (DN) for the entry where LDAP searches for groups begin. This field is case-insensitive.',dc=com'.
	GroupBaseDn *string `json:"group_base_dn,omitempty"`
	// Sets the attribute holding the members of the LDAP group. This field is case-sensitive.
	GroupMemberAttribute *string `json:"group_member_attribute,omitempty"`
	// Sets the attribute holding the name of a group, typically called `name` (in Active Directory) or `cn` (in OpenLDAP). This field is case-insensitive.
	GroupNameAttribute *string `json:"group_name_attribute,omitempty"`
	// The groups required to be present in the LDAP search result for successful authorization. This config parameter works in both **AND** / **OR** cases. - When `["group1 group2"]` are in the same array indices, both `group1` AND `group2` need to be present in the LDAP search result. - When `["group1", "group2"]` are in different array indices, either `group1` OR `group2` need to be present in the LDAP search result.
	GroupsRequired []string `json:"groups_required,omitempty"`
	// An optional string to use as part of the Authorization header. By default, a valid Authorization header looks like this: `Authorization: ldap base64(username:password)`. If `header_type` is set to "basic", then the Authorization header would be `Authorization: basic base64(username:password)`. Note that `header_type` can take any string, not just `'ldap'` and `'basic'`.
	HeaderType *string `json:"header_type,omitempty"`
	// An optional boolean value telling the plugin to hide the credential to the upstream server. It will be removed by Kong before proxying the request.
	HideCredentials *bool `json:"hide_credentials,omitempty"`
	// An optional value in milliseconds that defines how long an idle connection to LDAP server will live before being closed.
	Keepalive *float64 `json:"keepalive,omitempty"`
	// Host on which the LDAP server is running.
	LdapHost *string `json:"ldap_host,omitempty"`
	// The password to the LDAP server.
	LdapPassword *string `json:"ldap_password,omitempty"`
	// TCP port where the LDAP server is listening. 389 is the default port for non-SSL LDAP and AD. 636 is the port required for SSL LDAP and AD. If `ldaps` is configured, you must use port 636.
	LdapPort *float64 `json:"ldap_port,omitempty"`
	// Set it to `true` to use `ldaps`, a secure protocol (that can be configured to TLS) to connect to the LDAP server. When `ldaps` is configured, you must use port 636. If the `ldap` setting is enabled, ensure the `start_tls` setting is disabled.
	Ldaps *bool `json:"ldaps,omitempty"`
	// Displays all the LDAP search results received from the LDAP server for debugging purposes. Not recommended to be enabled in a production environment.
	LogSearchResults *bool `json:"log_search_results,omitempty"`
	// When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
	Realm *string `json:"realm,omitempty"`
	// Set it to `true` to issue StartTLS (Transport Layer Security) extended operation over `ldap` connection. If the `start_tls` setting is enabled, ensure the `ldaps` setting is disabled.
	StartTLS *bool `json:"start_tls,omitempty"`
	// An optional timeout in milliseconds when waiting for connection with LDAP server.
	Timeout *float64 `json:"timeout,omitempty"`
	// Set to `true` to authenticate LDAP server. The server certificate will be verified according to the CA certificates specified by the `lua_ssl_trusted_certificate` directive.
	VerifyLdapHost *bool `json:"verify_ldap_host,omitempty"`
}

func (o *LdapAuthAdvancedPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *LdapAuthAdvancedPluginConfig) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *LdapAuthAdvancedPluginConfig) GetBaseDn() *string {
	if o == nil {
		return nil
	}
	return o.BaseDn
}

func (o *LdapAuthAdvancedPluginConfig) GetBindDn() *string {
	if o == nil {
		return nil
	}
	return o.BindDn
}

func (o *LdapAuthAdvancedPluginConfig) GetCacheTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *LdapAuthAdvancedPluginConfig) GetConsumerBy() []LdapAuthAdvancedPluginConsumerBy {
	if o == nil {
		return nil
	}
	return o.ConsumerBy
}

func (o *LdapAuthAdvancedPluginConfig) GetConsumerOptional() *bool {
	if o == nil {
		return nil
	}
	return o.ConsumerOptional
}

func (o *LdapAuthAdvancedPluginConfig) GetGroupBaseDn() *string {
	if o == nil {
		return nil
	}
	return o.GroupBaseDn
}

func (o *LdapAuthAdvancedPluginConfig) GetGroupMemberAttribute() *string {
	if o == nil {
		return nil
	}
	return o.GroupMemberAttribute
}

func (o *LdapAuthAdvancedPluginConfig) GetGroupNameAttribute() *string {
	if o == nil {
		return nil
	}
	return o.GroupNameAttribute
}

func (o *LdapAuthAdvancedPluginConfig) GetGroupsRequired() []string {
	if o == nil {
		return nil
	}
	return o.GroupsRequired
}

func (o *LdapAuthAdvancedPluginConfig) GetHeaderType() *string {
	if o == nil {
		return nil
	}
	return o.HeaderType
}

func (o *LdapAuthAdvancedPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *LdapAuthAdvancedPluginConfig) GetKeepalive() *float64 {
	if o == nil {
		return nil
	}
	return o.Keepalive
}

func (o *LdapAuthAdvancedPluginConfig) GetLdapHost() *string {
	if o == nil {
		return nil
	}
	return o.LdapHost
}

func (o *LdapAuthAdvancedPluginConfig) GetLdapPassword() *string {
	if o == nil {
		return nil
	}
	return o.LdapPassword
}

func (o *LdapAuthAdvancedPluginConfig) GetLdapPort() *float64 {
	if o == nil {
		return nil
	}
	return o.LdapPort
}

func (o *LdapAuthAdvancedPluginConfig) GetLdaps() *bool {
	if o == nil {
		return nil
	}
	return o.Ldaps
}

func (o *LdapAuthAdvancedPluginConfig) GetLogSearchResults() *bool {
	if o == nil {
		return nil
	}
	return o.LogSearchResults
}

func (o *LdapAuthAdvancedPluginConfig) GetRealm() *string {
	if o == nil {
		return nil
	}
	return o.Realm
}

func (o *LdapAuthAdvancedPluginConfig) GetStartTLS() *bool {
	if o == nil {
		return nil
	}
	return o.StartTLS
}

func (o *LdapAuthAdvancedPluginConfig) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *LdapAuthAdvancedPluginConfig) GetVerifyLdapHost() *bool {
	if o == nil {
		return nil
	}
	return o.VerifyLdapHost
}

type LdapAuthAdvancedPluginProtocols string

const (
	LdapAuthAdvancedPluginProtocolsGrpc  LdapAuthAdvancedPluginProtocols = "grpc"
	LdapAuthAdvancedPluginProtocolsGrpcs LdapAuthAdvancedPluginProtocols = "grpcs"
	LdapAuthAdvancedPluginProtocolsHTTP  LdapAuthAdvancedPluginProtocols = "http"
	LdapAuthAdvancedPluginProtocolsHTTPS LdapAuthAdvancedPluginProtocols = "https"
	LdapAuthAdvancedPluginProtocolsWs    LdapAuthAdvancedPluginProtocols = "ws"
	LdapAuthAdvancedPluginProtocolsWss   LdapAuthAdvancedPluginProtocols = "wss"
)

func (e LdapAuthAdvancedPluginProtocols) ToPointer() *LdapAuthAdvancedPluginProtocols {
	return &e
}
func (e *LdapAuthAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = LdapAuthAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LdapAuthAdvancedPluginProtocols: %v", v)
	}
}

// LdapAuthAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type LdapAuthAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *LdapAuthAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// LdapAuthAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type LdapAuthAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *LdapAuthAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// LdapAuthAdvancedPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type LdapAuthAdvancedPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                            `json:"enabled,omitempty"`
	ID           *string                          `json:"id,omitempty"`
	InstanceName *string                          `json:"instance_name,omitempty"`
	name         string                           `const:"ldap-auth-advanced" json:"name"`
	Ordering     *LdapAuthAdvancedPluginOrdering  `json:"ordering,omitempty"`
	Partials     []LdapAuthAdvancedPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                        `json:"updated_at,omitempty"`
	Config    *LdapAuthAdvancedPluginConfig `json:"config,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support tcp and tls.
	Protocols []LdapAuthAdvancedPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *LdapAuthAdvancedPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *LdapAuthAdvancedPluginService `json:"service,omitempty"`
}

func (l LdapAuthAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LdapAuthAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LdapAuthAdvancedPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LdapAuthAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *LdapAuthAdvancedPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LdapAuthAdvancedPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *LdapAuthAdvancedPlugin) GetName() string {
	return "ldap-auth-advanced"
}

func (o *LdapAuthAdvancedPlugin) GetOrdering() *LdapAuthAdvancedPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *LdapAuthAdvancedPlugin) GetPartials() []LdapAuthAdvancedPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *LdapAuthAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *LdapAuthAdvancedPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LdapAuthAdvancedPlugin) GetConfig() *LdapAuthAdvancedPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *LdapAuthAdvancedPlugin) GetProtocols() []LdapAuthAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *LdapAuthAdvancedPlugin) GetRoute() *LdapAuthAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *LdapAuthAdvancedPlugin) GetService() *LdapAuthAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
