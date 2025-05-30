// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
)

type KafkaLogPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *KafkaLogPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type KafkaLogPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *KafkaLogPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type KafkaLogPluginOrdering struct {
	After  *KafkaLogPluginAfter  `json:"after,omitempty"`
	Before *KafkaLogPluginBefore `json:"before,omitempty"`
}

func (o *KafkaLogPluginOrdering) GetAfter() *KafkaLogPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *KafkaLogPluginOrdering) GetBefore() *KafkaLogPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type KafkaLogPluginPartials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

func (o *KafkaLogPluginPartials) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KafkaLogPluginPartials) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *KafkaLogPluginPartials) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// KafkaLogPluginMechanism - The SASL authentication mechanism.  Supported options: `PLAIN`, `SCRAM-SHA-256` or `SCRAM-SHA-512`.
type KafkaLogPluginMechanism string

const (
	KafkaLogPluginMechanismPlain       KafkaLogPluginMechanism = "PLAIN"
	KafkaLogPluginMechanismScramSha256 KafkaLogPluginMechanism = "SCRAM-SHA-256"
	KafkaLogPluginMechanismScramSha512 KafkaLogPluginMechanism = "SCRAM-SHA-512"
)

func (e KafkaLogPluginMechanism) ToPointer() *KafkaLogPluginMechanism {
	return &e
}
func (e *KafkaLogPluginMechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PLAIN":
		fallthrough
	case "SCRAM-SHA-256":
		fallthrough
	case "SCRAM-SHA-512":
		*e = KafkaLogPluginMechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KafkaLogPluginMechanism: %v", v)
	}
}

// KafkaLogPluginStrategy - The authentication strategy for the plugin, the only option for the value is `sasl`.
type KafkaLogPluginStrategy string

const (
	KafkaLogPluginStrategySasl KafkaLogPluginStrategy = "sasl"
)

func (e KafkaLogPluginStrategy) ToPointer() *KafkaLogPluginStrategy {
	return &e
}
func (e *KafkaLogPluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sasl":
		*e = KafkaLogPluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KafkaLogPluginStrategy: %v", v)
	}
}

type KafkaLogPluginAuthentication struct {
	// The SASL authentication mechanism.  Supported options: `PLAIN`, `SCRAM-SHA-256` or `SCRAM-SHA-512`.
	Mechanism *KafkaLogPluginMechanism `json:"mechanism,omitempty"`
	// Password for SASL authentication.
	Password *string `json:"password,omitempty"`
	// The authentication strategy for the plugin, the only option for the value is `sasl`.
	Strategy *KafkaLogPluginStrategy `json:"strategy,omitempty"`
	// Enable this to indicate `DelegationToken` authentication
	Tokenauth *bool `json:"tokenauth,omitempty"`
	// Username for SASL authentication.
	User *string `json:"user,omitempty"`
}

func (o *KafkaLogPluginAuthentication) GetMechanism() *KafkaLogPluginMechanism {
	if o == nil {
		return nil
	}
	return o.Mechanism
}

func (o *KafkaLogPluginAuthentication) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *KafkaLogPluginAuthentication) GetStrategy() *KafkaLogPluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *KafkaLogPluginAuthentication) GetTokenauth() *bool {
	if o == nil {
		return nil
	}
	return o.Tokenauth
}

func (o *KafkaLogPluginAuthentication) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

type KafkaLogPluginBootstrapServers struct {
	// A string representing a host name, such as example.com.
	Host string `json:"host"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port int64 `json:"port"`
}

func (o *KafkaLogPluginBootstrapServers) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *KafkaLogPluginBootstrapServers) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

// KafkaLogPluginProducerRequestAcks - The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments; 1 for only the leader; and -1 for the full ISR (In-Sync Replica set).
type KafkaLogPluginProducerRequestAcks int64

const (
	KafkaLogPluginProducerRequestAcksMinus1 KafkaLogPluginProducerRequestAcks = -1
	KafkaLogPluginProducerRequestAcksZero   KafkaLogPluginProducerRequestAcks = 0
	KafkaLogPluginProducerRequestAcksOne    KafkaLogPluginProducerRequestAcks = 1
)

func (e KafkaLogPluginProducerRequestAcks) ToPointer() *KafkaLogPluginProducerRequestAcks {
	return &e
}
func (e *KafkaLogPluginProducerRequestAcks) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 0:
		fallthrough
	case 1:
		*e = KafkaLogPluginProducerRequestAcks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KafkaLogPluginProducerRequestAcks: %v", v)
	}
}

type KafkaLogPluginSecurity struct {
	// UUID of certificate entity for mTLS authentication.
	CertificateID *string `json:"certificate_id,omitempty"`
	// Enables TLS.
	Ssl *bool `json:"ssl,omitempty"`
}

func (o *KafkaLogPluginSecurity) GetCertificateID() *string {
	if o == nil {
		return nil
	}
	return o.CertificateID
}

func (o *KafkaLogPluginSecurity) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

type KafkaLogPluginConfig struct {
	Authentication *KafkaLogPluginAuthentication `json:"authentication,omitempty"`
	// Set of bootstrap brokers in a `{host: host, port: port}` list format.
	BootstrapServers []KafkaLogPluginBootstrapServers `json:"bootstrap_servers,omitempty"`
	// An identifier for the Kafka cluster. By default, this field generates a random string. You can also set your own custom cluster identifier.  If more than one Kafka plugin is configured without a `cluster_name` (that is, if the default autogenerated value is removed), these plugins will use the same producer, and by extension, the same cluster. Logs will be sent to the leader of the cluster.
	ClusterName *string `json:"cluster_name,omitempty"`
	// Lua code as a key-value map
	CustomFieldsByLua map[string]any `json:"custom_fields_by_lua,omitempty"`
	Keepalive         *int64         `json:"keepalive,omitempty"`
	KeepaliveEnabled  *bool          `json:"keepalive_enabled,omitempty"`
	// Flag to enable asynchronous mode.
	ProducerAsync *bool `json:"producer_async,omitempty"`
	// Maximum number of messages that can be buffered in memory in asynchronous mode.
	ProducerAsyncBufferingLimitsMessagesInMemory *int64 `json:"producer_async_buffering_limits_messages_in_memory,omitempty"`
	// Maximum time interval in milliseconds between buffer flushes in asynchronous mode.
	ProducerAsyncFlushTimeout *int64 `json:"producer_async_flush_timeout,omitempty"`
	// The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments; 1 for only the leader; and -1 for the full ISR (In-Sync Replica set).
	ProducerRequestAcks *KafkaLogPluginProducerRequestAcks `json:"producer_request_acks,omitempty"`
	// Maximum size of a Produce request in bytes.
	ProducerRequestLimitsBytesPerRequest *int64 `json:"producer_request_limits_bytes_per_request,omitempty"`
	// Maximum number of messages to include into a single Produce request.
	ProducerRequestLimitsMessagesPerRequest *int64 `json:"producer_request_limits_messages_per_request,omitempty"`
	// Backoff interval between retry attempts in milliseconds.
	ProducerRequestRetriesBackoffTimeout *int64 `json:"producer_request_retries_backoff_timeout,omitempty"`
	// Maximum number of retry attempts per single Produce request.
	ProducerRequestRetriesMaxAttempts *int64 `json:"producer_request_retries_max_attempts,omitempty"`
	// Time to wait for a Produce response in milliseconds
	ProducerRequestTimeout *int64                  `json:"producer_request_timeout,omitempty"`
	Security               *KafkaLogPluginSecurity `json:"security,omitempty"`
	// Socket timeout in milliseconds.
	Timeout *int64 `json:"timeout,omitempty"`
	// The Kafka topic to publish to.
	Topic *string `json:"topic,omitempty"`
}

func (o *KafkaLogPluginConfig) GetAuthentication() *KafkaLogPluginAuthentication {
	if o == nil {
		return nil
	}
	return o.Authentication
}

func (o *KafkaLogPluginConfig) GetBootstrapServers() []KafkaLogPluginBootstrapServers {
	if o == nil {
		return nil
	}
	return o.BootstrapServers
}

func (o *KafkaLogPluginConfig) GetClusterName() *string {
	if o == nil {
		return nil
	}
	return o.ClusterName
}

func (o *KafkaLogPluginConfig) GetCustomFieldsByLua() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldsByLua
}

func (o *KafkaLogPluginConfig) GetKeepalive() *int64 {
	if o == nil {
		return nil
	}
	return o.Keepalive
}

func (o *KafkaLogPluginConfig) GetKeepaliveEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.KeepaliveEnabled
}

func (o *KafkaLogPluginConfig) GetProducerAsync() *bool {
	if o == nil {
		return nil
	}
	return o.ProducerAsync
}

func (o *KafkaLogPluginConfig) GetProducerAsyncBufferingLimitsMessagesInMemory() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerAsyncBufferingLimitsMessagesInMemory
}

func (o *KafkaLogPluginConfig) GetProducerAsyncFlushTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerAsyncFlushTimeout
}

func (o *KafkaLogPluginConfig) GetProducerRequestAcks() *KafkaLogPluginProducerRequestAcks {
	if o == nil {
		return nil
	}
	return o.ProducerRequestAcks
}

func (o *KafkaLogPluginConfig) GetProducerRequestLimitsBytesPerRequest() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestLimitsBytesPerRequest
}

func (o *KafkaLogPluginConfig) GetProducerRequestLimitsMessagesPerRequest() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestLimitsMessagesPerRequest
}

func (o *KafkaLogPluginConfig) GetProducerRequestRetriesBackoffTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestRetriesBackoffTimeout
}

func (o *KafkaLogPluginConfig) GetProducerRequestRetriesMaxAttempts() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestRetriesMaxAttempts
}

func (o *KafkaLogPluginConfig) GetProducerRequestTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ProducerRequestTimeout
}

func (o *KafkaLogPluginConfig) GetSecurity() *KafkaLogPluginSecurity {
	if o == nil {
		return nil
	}
	return o.Security
}

func (o *KafkaLogPluginConfig) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *KafkaLogPluginConfig) GetTopic() *string {
	if o == nil {
		return nil
	}
	return o.Topic
}

// KafkaLogPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type KafkaLogPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *KafkaLogPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type KafkaLogPluginProtocols string

const (
	KafkaLogPluginProtocolsGrpc  KafkaLogPluginProtocols = "grpc"
	KafkaLogPluginProtocolsGrpcs KafkaLogPluginProtocols = "grpcs"
	KafkaLogPluginProtocolsHTTP  KafkaLogPluginProtocols = "http"
	KafkaLogPluginProtocolsHTTPS KafkaLogPluginProtocols = "https"
	KafkaLogPluginProtocolsWs    KafkaLogPluginProtocols = "ws"
	KafkaLogPluginProtocolsWss   KafkaLogPluginProtocols = "wss"
)

func (e KafkaLogPluginProtocols) ToPointer() *KafkaLogPluginProtocols {
	return &e
}
func (e *KafkaLogPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = KafkaLogPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KafkaLogPluginProtocols: %v", v)
	}
}

// KafkaLogPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type KafkaLogPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *KafkaLogPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// KafkaLogPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type KafkaLogPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *KafkaLogPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// KafkaLogPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type KafkaLogPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                    `json:"enabled,omitempty"`
	ID           *string                  `json:"id,omitempty"`
	InstanceName *string                  `json:"instance_name,omitempty"`
	name         string                   `const:"kafka-log" json:"name"`
	Ordering     *KafkaLogPluginOrdering  `json:"ordering,omitempty"`
	Partials     []KafkaLogPluginPartials `json:"partials,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                `json:"updated_at,omitempty"`
	Config    *KafkaLogPluginConfig `json:"config,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *KafkaLogPluginConsumer `json:"consumer,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support tcp and tls.
	Protocols []KafkaLogPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *KafkaLogPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *KafkaLogPluginService `json:"service,omitempty"`
}

func (k KafkaLogPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KafkaLogPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KafkaLogPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *KafkaLogPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *KafkaLogPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KafkaLogPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *KafkaLogPlugin) GetName() string {
	return "kafka-log"
}

func (o *KafkaLogPlugin) GetOrdering() *KafkaLogPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *KafkaLogPlugin) GetPartials() []KafkaLogPluginPartials {
	if o == nil {
		return nil
	}
	return o.Partials
}

func (o *KafkaLogPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *KafkaLogPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *KafkaLogPlugin) GetConfig() *KafkaLogPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *KafkaLogPlugin) GetConsumer() *KafkaLogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *KafkaLogPlugin) GetProtocols() []KafkaLogPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *KafkaLogPlugin) GetRoute() *KafkaLogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *KafkaLogPlugin) GetService() *KafkaLogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
