// Code generated from semantic convention specification. DO NOT EDIT.

// Package httpconv provides types and functionality for OpenTelemetry semantic
// conventions in the "http" namespace.
package http

import (
  "fmt"
  "github.com/prometheus/client_golang/prometheus"
)

type Attribute interface {
  ID() string
  Value() string
}

// ErrorTypeAttr is an attribute conforming to the error.type semantic conventions. It represents the describes a class of error the operation ended with
type ErrorTypeAttr string

var (
	// ErrorTypeOther is a fallback error value to be used when the instrumentation
	// doesn't define a custom value.
	ErrorTypeOther ErrorTypeAttr = "_OTHER"
)
func (a ErrorTypeAttr) ID() string {
  return "error.type"
}
func (a ErrorTypeAttr) Value() string {
  return fmt.Sprintf("%v", a)
}

// ConnectionStateAttr is an attribute conforming to the http.connection.state semantic conventions. It represents the state of the HTTP connection in the HTTP connection pool
type ConnectionStateAttr string

var (
	// ConnectionStateActive is the active state.
	ConnectionStateActive ConnectionStateAttr = "active"
	// ConnectionStateIdle is the idle state.
	ConnectionStateIdle ConnectionStateAttr = "idle"
)
func (a ConnectionStateAttr) ID() string {
  return "http.connection.state"
}
func (a ConnectionStateAttr) Value() string {
  return fmt.Sprintf("%v", a)
}

// RequestMethodAttr is an attribute conforming to the http.request.method semantic conventions. It represents the HTTP request method
type RequestMethodAttr string

var (
	// RequestMethodConnect is the CONNECT method.
	RequestMethodConnect RequestMethodAttr = "CONNECT"
	// RequestMethodDelete is the DELETE method.
	RequestMethodDelete RequestMethodAttr = "DELETE"
	// RequestMethodGet is the GET method.
	RequestMethodGet RequestMethodAttr = "GET"
	// RequestMethodHead is the HEAD method.
	RequestMethodHead RequestMethodAttr = "HEAD"
	// RequestMethodOptions is the OPTIONS method.
	RequestMethodOptions RequestMethodAttr = "OPTIONS"
	// RequestMethodPatch is the PATCH method.
	RequestMethodPatch RequestMethodAttr = "PATCH"
	// RequestMethodPost is the POST method.
	RequestMethodPost RequestMethodAttr = "POST"
	// RequestMethodPut is the PUT method.
	RequestMethodPut RequestMethodAttr = "PUT"
	// RequestMethodTrace is the TRACE method.
	RequestMethodTrace RequestMethodAttr = "TRACE"
	// RequestMethodOther is the any HTTP method that the instrumentation has no
	// prior knowledge of.
	RequestMethodOther RequestMethodAttr = "_OTHER"
)
func (a RequestMethodAttr) ID() string {
  return "http.request.method"
}
func (a RequestMethodAttr) Value() string {
  return fmt.Sprintf("%v", a)
}

// ResponseStatusCodeAttr is an attribute conforming to the http.response.status_code semantic conventions. It represents the [HTTP response status code]
//
// [HTTP response status code]: https://tools.ietf.org/html/rfc7231#section-6
type ResponseStatusCodeAttr int
func (a ResponseStatusCodeAttr) ID() string {
  return "http.response.status_code"
}
func (a ResponseStatusCodeAttr) Value() string {
  return fmt.Sprintf("%d", a)
}

// RouteAttr is an attribute conforming to the http.route semantic conventions. It represents the matched route, that is, the path template in the format used by the respective server framework
type RouteAttr string
func (a RouteAttr) ID() string {
  return "http.route"
}
func (a RouteAttr) Value() string {
  return string(a)
}

// NetworkPeerAddressAttr is an attribute conforming to the network.peer.address semantic conventions. It represents the peer address of the network connection - IP address or Unix domain socket name
type NetworkPeerAddressAttr string
func (a NetworkPeerAddressAttr) ID() string {
  return "network.peer.address"
}
func (a NetworkPeerAddressAttr) Value() string {
  return string(a)
}

// NetworkProtocolNameAttr is an attribute conforming to the network.protocol.name semantic conventions. It represents the [OSI application layer] or non-OSI equivalent
//
// [OSI application layer]: https://wikipedia.org/wiki/Application_layer
type NetworkProtocolNameAttr string
func (a NetworkProtocolNameAttr) ID() string {
  return "network.protocol.name"
}
func (a NetworkProtocolNameAttr) Value() string {
  return string(a)
}

// NetworkProtocolVersionAttr is an attribute conforming to the network.protocol.version semantic conventions. It represents the actual version of the protocol used for network communication
type NetworkProtocolVersionAttr string
func (a NetworkProtocolVersionAttr) ID() string {
  return "network.protocol.version"
}
func (a NetworkProtocolVersionAttr) Value() string {
  return string(a)
}

// ServerAddressAttr is an attribute conforming to the server.address semantic conventions. It represents the server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name
type ServerAddressAttr string
func (a ServerAddressAttr) ID() string {
  return "server.address"
}
func (a ServerAddressAttr) Value() string {
  return string(a)
}

// ServerPortAttr is an attribute conforming to the server.port semantic conventions. It represents the port identifier of the ["URI origin"] HTTP request is sent to
//
// ["URI origin"]: https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin
type ServerPortAttr int
func (a ServerPortAttr) ID() string {
  return "server.port"
}
func (a ServerPortAttr) Value() string {
  return fmt.Sprintf("%d", a)
}

// UrlSchemeAttr is an attribute conforming to the url.scheme semantic conventions. It represents the [URI scheme] component identifying the used protocol
//
// [URI scheme]: https://www.rfc-editor.org/rfc/rfc3986#section-3.1
type UrlSchemeAttr string
func (a UrlSchemeAttr) ID() string {
  return "url.scheme"
}
func (a UrlSchemeAttr) Value() string {
  return string(a)
}

// UrlTemplateAttr is an attribute conforming to the url.template semantic conventions. It represents the low-cardinality template of an [absolute path reference]
//
// [absolute path reference]: https://www.rfc-editor.org/rfc/rfc3986#section-4.2
type UrlTemplateAttr string
func (a UrlTemplateAttr) ID() string {
  return "url.template"
}
func (a UrlTemplateAttr) Value() string {
  return string(a)
}

// UserAgentSyntheticTypeAttr is an attribute conforming to the user_agent.synthetic.type semantic conventions. It represents the specifies the category of synthetic traffic, such as tests or bots
type UserAgentSyntheticTypeAttr string

var (
	// UserAgentSyntheticTypeBot is the bot source.
	UserAgentSyntheticTypeBot UserAgentSyntheticTypeAttr = "bot"
	// UserAgentSyntheticTypeTest is the synthetic test source.
	UserAgentSyntheticTypeTest UserAgentSyntheticTypeAttr = "test"
)
func (a UserAgentSyntheticTypeAttr) ID() string {
  return "user_agent.synthetic.type"
}
func (a UserAgentSyntheticTypeAttr) Value() string {
  return fmt.Sprintf("%v", a)
}

// ClientActiveRequests is an instrument used to record metric values conforming to the "http.client.active_requests" semantic conventions. It represents the number of active HTTP requests
type ClientActiveRequests struct {
  *prometheus.GaugeVec
}

// NewClientActiveRequests returns a new ClientActiveRequests instrument
func NewClientActiveRequests() ClientActiveRequests {
  labels := []string{
     "http.request.method",
     "server.port",
     "server.address",
     "url.scheme",
     "url.template",
     }
  return ClientActiveRequests{
     GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
         Name: "http_client_active_requests",
         Help: "Number of active HTTP requests.",
  }, labels)}
}


type ClientActiveRequestsAttr interface {
  Attribute
  implClientActiveRequests()
}
func (a RequestMethodAttr) implClientActiveRequests() {}
func (a ServerPortAttr) implClientActiveRequests() {}
func (a ServerAddressAttr) implClientActiveRequests() {}
func (a UrlSchemeAttr) implClientActiveRequests() {}
func (a UrlTemplateAttr) implClientActiveRequests() {}

func (m ClientActiveRequests) With(
    port ServerPortAttr,
    address ServerAddressAttr,
    extra ...ClientActiveRequestsAttr,
) prometheus.Gauge {
    labels := prometheus.Labels{
        "server.port":  port.Value(),
        "server.address":  address.Value(),
        "http.request.method":  "",
        "url.scheme":  "",
        "url.template":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.GaugeVec.With(labels)
}// ClientConnectionDuration is an instrument used to record metric values conforming to the "http.client.connection.duration" semantic conventions. It represents the duration of the successfully established outbound HTTP connections
type ClientConnectionDuration struct {
  *prometheus.HistogramVec
}

// NewClientConnectionDuration returns a new ClientConnectionDuration instrument
func NewClientConnectionDuration() ClientConnectionDuration {
  labels := []string{
     "network.peer.address",
     "server.port",
     "server.address",
     "url.scheme",
     "network.protocol.version",
     }
  return ClientConnectionDuration{
     HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
         Name: "http_client_connection_duration",
         Help: "The duration of the successfully established outbound HTTP connections.",
  }, labels)}
}


type ClientConnectionDurationAttr interface {
  Attribute
  implClientConnectionDuration()
}
func (a NetworkPeerAddressAttr) implClientConnectionDuration() {}
func (a ServerPortAttr) implClientConnectionDuration() {}
func (a ServerAddressAttr) implClientConnectionDuration() {}
func (a UrlSchemeAttr) implClientConnectionDuration() {}
func (a NetworkProtocolVersionAttr) implClientConnectionDuration() {}

func (m ClientConnectionDuration) With(
    port ServerPortAttr,
    address ServerAddressAttr,
    extra ...ClientConnectionDurationAttr,
) prometheus.Observer {
    labels := prometheus.Labels{
        "server.port":  port.Value(),
        "server.address":  address.Value(),
        "network.peer.address":  "",
        "url.scheme":  "",
        "network.protocol.version":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.HistogramVec.With(labels)
}// ClientOpenConnections is an instrument used to record metric values conforming to the "http.client.open_connections" semantic conventions. It represents the number of outbound HTTP connections that are currently active or idle on the client
type ClientOpenConnections struct {
  *prometheus.GaugeVec
}

// NewClientOpenConnections returns a new ClientOpenConnections instrument
func NewClientOpenConnections() ClientOpenConnections {
  labels := []string{
     "network.peer.address",
     "server.port",
     "server.address",
     "http.connection.state",
     "url.scheme",
     "network.protocol.version",
     }
  return ClientOpenConnections{
     GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
         Name: "http_client_open_connections",
         Help: "Number of outbound HTTP connections that are currently active or idle on the client.",
  }, labels)}
}


type ClientOpenConnectionsAttr interface {
  Attribute
  implClientOpenConnections()
}
func (a NetworkPeerAddressAttr) implClientOpenConnections() {}
func (a ServerPortAttr) implClientOpenConnections() {}
func (a ServerAddressAttr) implClientOpenConnections() {}
func (a ConnectionStateAttr) implClientOpenConnections() {}
func (a UrlSchemeAttr) implClientOpenConnections() {}
func (a NetworkProtocolVersionAttr) implClientOpenConnections() {}

func (m ClientOpenConnections) With(
    port ServerPortAttr,
    address ServerAddressAttr,
    connectionState ConnectionStateAttr,
    extra ...ClientOpenConnectionsAttr,
) prometheus.Gauge {
    labels := prometheus.Labels{
        "server.port":  port.Value(),
        "server.address":  address.Value(),
        "http.connection.state":  connectionState.Value(),
        "network.peer.address":  "",
        "url.scheme":  "",
        "network.protocol.version":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.GaugeVec.With(labels)
}// ClientRequestBodySize is an instrument used to record metric values conforming to the "http.client.request.body.size" semantic conventions. It represents the size of HTTP client request bodies
type ClientRequestBodySize struct {
  *prometheus.HistogramVec
}

// NewClientRequestBodySize returns a new ClientRequestBodySize instrument
func NewClientRequestBodySize() ClientRequestBodySize {
  labels := []string{
     "error.type",
     "http.request.method",
     "server.address",
     "server.port",
     "url.scheme",
     "network.protocol.version",
     "url.template",
     "http.response.status_code",
     "network.protocol.name",
     }
  return ClientRequestBodySize{
     HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
         Name: "http_client_request_body_size",
         Help: "Size of HTTP client request bodies.",
  }, labels)}
}


type ClientRequestBodySizeAttr interface {
  Attribute
  implClientRequestBodySize()
}
func (a ErrorTypeAttr) implClientRequestBodySize() {}
func (a RequestMethodAttr) implClientRequestBodySize() {}
func (a ServerAddressAttr) implClientRequestBodySize() {}
func (a ServerPortAttr) implClientRequestBodySize() {}
func (a UrlSchemeAttr) implClientRequestBodySize() {}
func (a NetworkProtocolVersionAttr) implClientRequestBodySize() {}
func (a UrlTemplateAttr) implClientRequestBodySize() {}
func (a ResponseStatusCodeAttr) implClientRequestBodySize() {}
func (a NetworkProtocolNameAttr) implClientRequestBodySize() {}

func (m ClientRequestBodySize) With(
    requestMethod RequestMethodAttr,
    address ServerAddressAttr,
    port ServerPortAttr,
    extra ...ClientRequestBodySizeAttr,
) prometheus.Observer {
    labels := prometheus.Labels{
        "http.request.method":  requestMethod.Value(),
        "server.address":  address.Value(),
        "server.port":  port.Value(),
        "error.type":  "",
        "url.scheme":  "",
        "network.protocol.version":  "",
        "url.template":  "",
        "http.response.status_code":  "",
        "network.protocol.name":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.HistogramVec.With(labels)
}// ClientRequestDuration is an instrument used to record metric values conforming to the "http.client.request.duration" semantic conventions. It represents the duration of HTTP client requests
type ClientRequestDuration struct {
  *prometheus.HistogramVec
}

// NewClientRequestDuration returns a new ClientRequestDuration instrument
func NewClientRequestDuration() ClientRequestDuration {
  labels := []string{
     "error.type",
     "http.request.method",
     "server.address",
     "server.port",
     "url.scheme",
     "network.protocol.version",
     "url.template",
     "http.response.status_code",
     "network.protocol.name",
     }
  return ClientRequestDuration{
     HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
         Name: "http_client_request_duration",
         Help: "Duration of HTTP client requests.",
  }, labels)}
}


type ClientRequestDurationAttr interface {
  Attribute
  implClientRequestDuration()
}
func (a ErrorTypeAttr) implClientRequestDuration() {}
func (a RequestMethodAttr) implClientRequestDuration() {}
func (a ServerAddressAttr) implClientRequestDuration() {}
func (a ServerPortAttr) implClientRequestDuration() {}
func (a UrlSchemeAttr) implClientRequestDuration() {}
func (a NetworkProtocolVersionAttr) implClientRequestDuration() {}
func (a UrlTemplateAttr) implClientRequestDuration() {}
func (a ResponseStatusCodeAttr) implClientRequestDuration() {}
func (a NetworkProtocolNameAttr) implClientRequestDuration() {}

func (m ClientRequestDuration) With(
    requestMethod RequestMethodAttr,
    address ServerAddressAttr,
    port ServerPortAttr,
    extra ...ClientRequestDurationAttr,
) prometheus.Observer {
    labels := prometheus.Labels{
        "http.request.method":  requestMethod.Value(),
        "server.address":  address.Value(),
        "server.port":  port.Value(),
        "error.type":  "",
        "url.scheme":  "",
        "network.protocol.version":  "",
        "url.template":  "",
        "http.response.status_code":  "",
        "network.protocol.name":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.HistogramVec.With(labels)
}// ClientResponseBodySize is an instrument used to record metric values conforming to the "http.client.response.body.size" semantic conventions. It represents the size of HTTP client response bodies
type ClientResponseBodySize struct {
  *prometheus.HistogramVec
}

// NewClientResponseBodySize returns a new ClientResponseBodySize instrument
func NewClientResponseBodySize() ClientResponseBodySize {
  labels := []string{
     "error.type",
     "http.request.method",
     "server.address",
     "server.port",
     "url.scheme",
     "network.protocol.version",
     "url.template",
     "http.response.status_code",
     "network.protocol.name",
     }
  return ClientResponseBodySize{
     HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
         Name: "http_client_response_body_size",
         Help: "Size of HTTP client response bodies.",
  }, labels)}
}


type ClientResponseBodySizeAttr interface {
  Attribute
  implClientResponseBodySize()
}
func (a ErrorTypeAttr) implClientResponseBodySize() {}
func (a RequestMethodAttr) implClientResponseBodySize() {}
func (a ServerAddressAttr) implClientResponseBodySize() {}
func (a ServerPortAttr) implClientResponseBodySize() {}
func (a UrlSchemeAttr) implClientResponseBodySize() {}
func (a NetworkProtocolVersionAttr) implClientResponseBodySize() {}
func (a UrlTemplateAttr) implClientResponseBodySize() {}
func (a ResponseStatusCodeAttr) implClientResponseBodySize() {}
func (a NetworkProtocolNameAttr) implClientResponseBodySize() {}

func (m ClientResponseBodySize) With(
    requestMethod RequestMethodAttr,
    address ServerAddressAttr,
    port ServerPortAttr,
    extra ...ClientResponseBodySizeAttr,
) prometheus.Observer {
    labels := prometheus.Labels{
        "http.request.method":  requestMethod.Value(),
        "server.address":  address.Value(),
        "server.port":  port.Value(),
        "error.type":  "",
        "url.scheme":  "",
        "network.protocol.version":  "",
        "url.template":  "",
        "http.response.status_code":  "",
        "network.protocol.name":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.HistogramVec.With(labels)
}// ServerActiveRequests is an instrument used to record metric values conforming to the "http.server.active_requests" semantic conventions. It represents the number of active HTTP server requests
type ServerActiveRequests struct {
  *prometheus.GaugeVec
}

// NewServerActiveRequests returns a new ServerActiveRequests instrument
func NewServerActiveRequests() ServerActiveRequests {
  labels := []string{
     "http.request.method",
     "server.address",
     "server.port",
     "url.scheme",
     }
  return ServerActiveRequests{
     GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
         Name: "http_server_active_requests",
         Help: "Number of active HTTP server requests.",
  }, labels)}
}


type ServerActiveRequestsAttr interface {
  Attribute
  implServerActiveRequests()
}
func (a RequestMethodAttr) implServerActiveRequests() {}
func (a ServerAddressAttr) implServerActiveRequests() {}
func (a ServerPortAttr) implServerActiveRequests() {}
func (a UrlSchemeAttr) implServerActiveRequests() {}

func (m ServerActiveRequests) With(
    requestMethod RequestMethodAttr,
    scheme UrlSchemeAttr,
    extra ...ServerActiveRequestsAttr,
) prometheus.Gauge {
    labels := prometheus.Labels{
        "http.request.method":  requestMethod.Value(),
        "url.scheme":  scheme.Value(),
        "server.address":  "",
        "server.port":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.GaugeVec.With(labels)
}// ServerRequestBodySize is an instrument used to record metric values conforming to the "http.server.request.body.size" semantic conventions. It represents the size of HTTP server request bodies
type ServerRequestBodySize struct {
  *prometheus.HistogramVec
}

// NewServerRequestBodySize returns a new ServerRequestBodySize instrument
func NewServerRequestBodySize() ServerRequestBodySize {
  labels := []string{
     "error.type",
     "http.request.method",
     "server.address",
     "server.port",
     "user_agent.synthetic.type",
     "url.scheme",
     "network.protocol.version",
     "http.route",
     "http.response.status_code",
     "network.protocol.name",
     }
  return ServerRequestBodySize{
     HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
         Name: "http_server_request_body_size",
         Help: "Size of HTTP server request bodies.",
  }, labels)}
}


type ServerRequestBodySizeAttr interface {
  Attribute
  implServerRequestBodySize()
}
func (a ErrorTypeAttr) implServerRequestBodySize() {}
func (a RequestMethodAttr) implServerRequestBodySize() {}
func (a ServerAddressAttr) implServerRequestBodySize() {}
func (a ServerPortAttr) implServerRequestBodySize() {}
func (a UserAgentSyntheticTypeAttr) implServerRequestBodySize() {}
func (a UrlSchemeAttr) implServerRequestBodySize() {}
func (a NetworkProtocolVersionAttr) implServerRequestBodySize() {}
func (a RouteAttr) implServerRequestBodySize() {}
func (a ResponseStatusCodeAttr) implServerRequestBodySize() {}
func (a NetworkProtocolNameAttr) implServerRequestBodySize() {}

func (m ServerRequestBodySize) With(
    requestMethod RequestMethodAttr,
    scheme UrlSchemeAttr,
    extra ...ServerRequestBodySizeAttr,
) prometheus.Observer {
    labels := prometheus.Labels{
        "http.request.method":  requestMethod.Value(),
        "url.scheme":  scheme.Value(),
        "error.type":  "",
        "server.address":  "",
        "server.port":  "",
        "user_agent.synthetic.type":  "",
        "network.protocol.version":  "",
        "http.route":  "",
        "http.response.status_code":  "",
        "network.protocol.name":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.HistogramVec.With(labels)
}// ServerRequestDuration is an instrument used to record metric values conforming to the "http.server.request.duration" semantic conventions. It represents the duration of HTTP server requests
type ServerRequestDuration struct {
  *prometheus.HistogramVec
}

// NewServerRequestDuration returns a new ServerRequestDuration instrument
func NewServerRequestDuration() ServerRequestDuration {
  labels := []string{
     "error.type",
     "http.request.method",
     "server.address",
     "server.port",
     "user_agent.synthetic.type",
     "url.scheme",
     "network.protocol.version",
     "http.route",
     "http.response.status_code",
     "network.protocol.name",
     }
  return ServerRequestDuration{
     HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
         Name: "http_server_request_duration",
         Help: "Duration of HTTP server requests.",
  }, labels)}
}


type ServerRequestDurationAttr interface {
  Attribute
  implServerRequestDuration()
}
func (a ErrorTypeAttr) implServerRequestDuration() {}
func (a RequestMethodAttr) implServerRequestDuration() {}
func (a ServerAddressAttr) implServerRequestDuration() {}
func (a ServerPortAttr) implServerRequestDuration() {}
func (a UserAgentSyntheticTypeAttr) implServerRequestDuration() {}
func (a UrlSchemeAttr) implServerRequestDuration() {}
func (a NetworkProtocolVersionAttr) implServerRequestDuration() {}
func (a RouteAttr) implServerRequestDuration() {}
func (a ResponseStatusCodeAttr) implServerRequestDuration() {}
func (a NetworkProtocolNameAttr) implServerRequestDuration() {}

func (m ServerRequestDuration) With(
    requestMethod RequestMethodAttr,
    scheme UrlSchemeAttr,
    extra ...ServerRequestDurationAttr,
) prometheus.Observer {
    labels := prometheus.Labels{
        "http.request.method":  requestMethod.Value(),
        "url.scheme":  scheme.Value(),
        "error.type":  "",
        "server.address":  "",
        "server.port":  "",
        "user_agent.synthetic.type":  "",
        "network.protocol.version":  "",
        "http.route":  "",
        "http.response.status_code":  "",
        "network.protocol.name":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.HistogramVec.With(labels)
}// ServerResponseBodySize is an instrument used to record metric values conforming to the "http.server.response.body.size" semantic conventions. It represents the size of HTTP server response bodies
type ServerResponseBodySize struct {
  *prometheus.HistogramVec
}

// NewServerResponseBodySize returns a new ServerResponseBodySize instrument
func NewServerResponseBodySize() ServerResponseBodySize {
  labels := []string{
     "error.type",
     "http.request.method",
     "server.address",
     "server.port",
     "user_agent.synthetic.type",
     "url.scheme",
     "network.protocol.version",
     "http.route",
     "http.response.status_code",
     "network.protocol.name",
     }
  return ServerResponseBodySize{
     HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
         Name: "http_server_response_body_size",
         Help: "Size of HTTP server response bodies.",
  }, labels)}
}


type ServerResponseBodySizeAttr interface {
  Attribute
  implServerResponseBodySize()
}
func (a ErrorTypeAttr) implServerResponseBodySize() {}
func (a RequestMethodAttr) implServerResponseBodySize() {}
func (a ServerAddressAttr) implServerResponseBodySize() {}
func (a ServerPortAttr) implServerResponseBodySize() {}
func (a UserAgentSyntheticTypeAttr) implServerResponseBodySize() {}
func (a UrlSchemeAttr) implServerResponseBodySize() {}
func (a NetworkProtocolVersionAttr) implServerResponseBodySize() {}
func (a RouteAttr) implServerResponseBodySize() {}
func (a ResponseStatusCodeAttr) implServerResponseBodySize() {}
func (a NetworkProtocolNameAttr) implServerResponseBodySize() {}

func (m ServerResponseBodySize) With(
    requestMethod RequestMethodAttr,
    scheme UrlSchemeAttr,
    extra ...ServerResponseBodySizeAttr,
) prometheus.Observer {
    labels := prometheus.Labels{
        "http.request.method":  requestMethod.Value(),
        "url.scheme":  scheme.Value(),
        "error.type":  "",
        "server.address":  "",
        "server.port":  "",
        "user_agent.synthetic.type":  "",
        "network.protocol.version":  "",
        "http.route":  "",
        "http.response.status_code":  "",
        "network.protocol.name":  "",
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.HistogramVec.With(labels)
}
/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "all_attrs": [
            {
                "brief": "Describes a class of error the operation ended with.\n",
                "examples": [
                    "timeout",
                    "java.net.UnknownHostException",
                    "server_certificate_invalid",
                    "500",
                ],
                "name": "error.type",
                "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                "requirement_level": {
                    "conditionally_required": "If request has ended with an error.",
                },
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                            "id": "other",
                            "stability": "stable",
                            "value": "_OTHER",
                        },
                    ],
                },
            },
            {
                "brief": "State of the HTTP connection in the HTTP connection pool.",
                "examples": [
                    "active",
                    "idle",
                ],
                "name": "http.connection.state",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "active state.",
                            "id": "active",
                            "stability": "development",
                            "value": "active",
                        },
                        {
                            "brief": "idle state.",
                            "id": "idle",
                            "stability": "development",
                            "value": "idle",
                        },
                    ],
                },
            },
            {
                "brief": "HTTP request method.",
                "examples": [
                    "GET",
                    "POST",
                    "HEAD",
                ],
                "name": "http.request.method",
                "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "CONNECT method.",
                            "id": "connect",
                            "stability": "stable",
                            "value": "CONNECT",
                        },
                        {
                            "brief": "DELETE method.",
                            "id": "delete",
                            "stability": "stable",
                            "value": "DELETE",
                        },
                        {
                            "brief": "GET method.",
                            "id": "get",
                            "stability": "stable",
                            "value": "GET",
                        },
                        {
                            "brief": "HEAD method.",
                            "id": "head",
                            "stability": "stable",
                            "value": "HEAD",
                        },
                        {
                            "brief": "OPTIONS method.",
                            "id": "options",
                            "stability": "stable",
                            "value": "OPTIONS",
                        },
                        {
                            "brief": "PATCH method.",
                            "id": "patch",
                            "stability": "stable",
                            "value": "PATCH",
                        },
                        {
                            "brief": "POST method.",
                            "id": "post",
                            "stability": "stable",
                            "value": "POST",
                        },
                        {
                            "brief": "PUT method.",
                            "id": "put",
                            "stability": "stable",
                            "value": "PUT",
                        },
                        {
                            "brief": "TRACE method.",
                            "id": "trace",
                            "stability": "stable",
                            "value": "TRACE",
                        },
                        {
                            "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                            "id": "other",
                            "stability": "stable",
                            "value": "_OTHER",
                        },
                    ],
                },
            },
            {
                "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                "examples": [
                    200,
                ],
                "name": "http.response.status_code",
                "requirement_level": {
                    "conditionally_required": "If and only if one was received/sent.",
                },
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                "examples": [
                    "/users/:userID?",
                    "{controller}/{action}/{id?}",
                ],
                "name": "http.route",
                "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                "requirement_level": {
                    "conditionally_required": "If and only if it's available",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Peer address of the network connection - IP address or Unix domain socket name.",
                "examples": [
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "network.peer.address",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                "examples": [
                    "http",
                    "spdy",
                ],
                "name": "network.protocol.name",
                "note": "The value SHOULD be normalized to lowercase.",
                "requirement_level": {
                    "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The actual version of the protocol used for network communication.",
                "examples": [
                    "1.1",
                    "2",
                ],
                "name": "network.protocol.version",
                "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                "examples": [
                    "http",
                    "https",
                ],
                "name": "url.scheme",
                "requirement_level": "opt_in",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).\n",
                "examples": [
                    "/users/{id}",
                    "/users/:id",
                    "/users?id={id}",
                ],
                "name": "url.template",
                "note": "The `url.template` MUST have low cardinality. It is not usually available on HTTP clients, but may be known by the application or specialized HTTP instrumentation.\n",
                "requirement_level": {
                    "conditionally_required": "If available.",
                },
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "Specifies the category of synthetic traffic, such as tests or bots.\n",
                "name": "user_agent.synthetic.type",
                "note": "This attribute MAY be derived from the contents of the `user_agent.original` attribute. Components that populate the attribute are responsible for determining what they consider to be synthetic bot or test traffic. This attribute can either be set for self-identification purposes, or on telemetry detected to be generated as a result of a synthetic request. This attribute is useful for distinguishing between genuine client traffic and synthetic traffic generated by bots or tests.\n",
                "requirement_level": "opt_in",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Bot source.",
                            "id": "bot",
                            "stability": "development",
                            "value": "bot",
                        },
                        {
                            "brief": "Synthetic test source.",
                            "id": "test",
                            "stability": "development",
                            "value": "test",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "metrics": [
                {
                    "attributes": [
                        {
                            "brief": "HTTP request method.",
                            "examples": [
                                "GET",
                                "POST",
                                "HEAD",
                            ],
                            "name": "http.request.method",
                            "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "CONNECT method.",
                                        "id": "connect",
                                        "stability": "stable",
                                        "value": "CONNECT",
                                    },
                                    {
                                        "brief": "DELETE method.",
                                        "id": "delete",
                                        "stability": "stable",
                                        "value": "DELETE",
                                    },
                                    {
                                        "brief": "GET method.",
                                        "id": "get",
                                        "stability": "stable",
                                        "value": "GET",
                                    },
                                    {
                                        "brief": "HEAD method.",
                                        "id": "head",
                                        "stability": "stable",
                                        "value": "HEAD",
                                    },
                                    {
                                        "brief": "OPTIONS method.",
                                        "id": "options",
                                        "stability": "stable",
                                        "value": "OPTIONS",
                                    },
                                    {
                                        "brief": "PATCH method.",
                                        "id": "patch",
                                        "stability": "stable",
                                        "value": "PATCH",
                                    },
                                    {
                                        "brief": "POST method.",
                                        "id": "post",
                                        "stability": "stable",
                                        "value": "POST",
                                    },
                                    {
                                        "brief": "PUT method.",
                                        "id": "put",
                                        "stability": "stable",
                                        "value": "PUT",
                                    },
                                    {
                                        "brief": "TRACE method.",
                                        "id": "trace",
                                        "stability": "stable",
                                        "value": "TRACE",
                                    },
                                    {
                                        "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).\n",
                            "examples": [
                                "/users/{id}",
                                "/users/:id",
                                "/users?id={id}",
                            ],
                            "name": "url.template",
                            "note": "The `url.template` MUST have low cardinality. It is not usually available on HTTP clients, but may be known by the application or specialized HTTP instrumentation.\n",
                            "requirement_level": {
                                "conditionally_required": "If available.",
                            },
                            "stability": "development",
                            "type": "string",
                        },
                    ],
                    "brief": "Number of active HTTP requests.",
                    "id": "metric.http.client.active_requests",
                    "instrument": "updowncounter",
                    "lineage": {
                        "attributes": {
                            "http.request.method": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                            "url.template": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.client.active_requests",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "metric",
                    "unit": "{request}",
                },
                {
                    "attributes": [
                        {
                            "brief": "Peer address of the network connection - IP address or Unix domain socket name.",
                            "examples": [
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "network.peer.address",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The actual version of the protocol used for network communication.",
                            "examples": [
                                "1.1",
                                "2",
                            ],
                            "name": "network.protocol.version",
                            "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                    ],
                    "brief": "The duration of the successfully established outbound HTTP connections.",
                    "id": "metric.http.client.connection.duration",
                    "instrument": "histogram",
                    "lineage": {
                        "attributes": {
                            "network.peer.address": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "network.protocol.version": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.client.connection.duration",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "metric",
                    "unit": "s",
                },
                {
                    "attributes": [
                        {
                            "brief": "Peer address of the network connection - IP address or Unix domain socket name.",
                            "examples": [
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "network.peer.address",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The actual version of the protocol used for network communication.",
                            "examples": [
                                "1.1",
                                "2",
                            ],
                            "name": "network.protocol.version",
                            "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "State of the HTTP connection in the HTTP connection pool.",
                            "examples": [
                                "active",
                                "idle",
                            ],
                            "name": "http.connection.state",
                            "requirement_level": "required",
                            "stability": "development",
                            "type": {
                                "members": [
                                    {
                                        "brief": "active state.",
                                        "id": "active",
                                        "stability": "development",
                                        "value": "active",
                                    },
                                    {
                                        "brief": "idle state.",
                                        "id": "idle",
                                        "stability": "development",
                                        "value": "idle",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                    ],
                    "brief": "Number of outbound HTTP connections that are currently active or idle on the client.",
                    "id": "metric.http.client.open_connections",
                    "instrument": "updowncounter",
                    "lineage": {
                        "attributes": {
                            "http.connection.state": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "network.peer.address": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "network.protocol.version": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.client.open_connections",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "metric",
                    "unit": "{connection}",
                },
                {
                    "attributes": [
                        {
                            "brief": "Describes a class of error the operation ended with.\n",
                            "examples": [
                                "timeout",
                                "java.net.UnknownHostException",
                                "server_certificate_invalid",
                                "500",
                            ],
                            "name": "error.type",
                            "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                            "requirement_level": {
                                "conditionally_required": "If request has ended with an error.",
                            },
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "HTTP request method.",
                            "examples": [
                                "GET",
                                "POST",
                                "HEAD",
                            ],
                            "name": "http.request.method",
                            "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "CONNECT method.",
                                        "id": "connect",
                                        "stability": "stable",
                                        "value": "CONNECT",
                                    },
                                    {
                                        "brief": "DELETE method.",
                                        "id": "delete",
                                        "stability": "stable",
                                        "value": "DELETE",
                                    },
                                    {
                                        "brief": "GET method.",
                                        "id": "get",
                                        "stability": "stable",
                                        "value": "GET",
                                    },
                                    {
                                        "brief": "HEAD method.",
                                        "id": "head",
                                        "stability": "stable",
                                        "value": "HEAD",
                                    },
                                    {
                                        "brief": "OPTIONS method.",
                                        "id": "options",
                                        "stability": "stable",
                                        "value": "OPTIONS",
                                    },
                                    {
                                        "brief": "PATCH method.",
                                        "id": "patch",
                                        "stability": "stable",
                                        "value": "PATCH",
                                    },
                                    {
                                        "brief": "POST method.",
                                        "id": "post",
                                        "stability": "stable",
                                        "value": "POST",
                                    },
                                    {
                                        "brief": "PUT method.",
                                        "id": "put",
                                        "stability": "stable",
                                        "value": "PUT",
                                    },
                                    {
                                        "brief": "TRACE method.",
                                        "id": "trace",
                                        "stability": "stable",
                                        "value": "TRACE",
                                    },
                                    {
                                        "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                            "examples": [
                                200,
                            ],
                            "name": "http.response.status_code",
                            "requirement_level": {
                                "conditionally_required": "If and only if one was received/sent.",
                            },
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                            "examples": [
                                "http",
                                "spdy",
                            ],
                            "name": "network.protocol.name",
                            "note": "The value SHOULD be normalized to lowercase.",
                            "requirement_level": {
                                "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The actual version of the protocol used for network communication.",
                            "examples": [
                                "1.0",
                                "1.1",
                                "2",
                                "3",
                            ],
                            "name": "network.protocol.version",
                            "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Host identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "If an HTTP client request is explicitly made to an IP address, e.g. `http://x.x.x.x:8080`, then `server.address` SHOULD be the IP address `x.x.x.x`. A DNS lookup SHOULD NOT be used.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).\n",
                            "examples": [
                                "/users/{id}",
                                "/users/:id",
                                "/users?id={id}",
                            ],
                            "name": "url.template",
                            "note": "The `url.template` MUST have low cardinality. It is not usually available on HTTP clients, but may be known by the application or specialized HTTP instrumentation.\n",
                            "requirement_level": {
                                "conditionally_required": "If available.",
                            },
                            "stability": "development",
                            "type": "string",
                        },
                    ],
                    "brief": "Size of HTTP client request bodies.",
                    "id": "metric.http.client.request.body.size",
                    "instrument": "histogram",
                    "lineage": {
                        "attributes": {
                            "error.type": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.error",
                            },
                            "http.request.method": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.response.status_code": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "network.protocol.name": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "network.protocol.version": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                ],
                                "source_group": "registry.network",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                            "url.template": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.client.request.body.size",
                    "note": "The size of the request payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length) header. For requests using transport encoding, this should be the compressed size.\n",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "metric",
                    "unit": "By",
                },
                {
                    "attributes": [
                        {
                            "brief": "Describes a class of error the operation ended with.\n",
                            "examples": [
                                "timeout",
                                "java.net.UnknownHostException",
                                "server_certificate_invalid",
                                "500",
                            ],
                            "name": "error.type",
                            "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                            "requirement_level": {
                                "conditionally_required": "If request has ended with an error.",
                            },
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "HTTP request method.",
                            "examples": [
                                "GET",
                                "POST",
                                "HEAD",
                            ],
                            "name": "http.request.method",
                            "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "CONNECT method.",
                                        "id": "connect",
                                        "stability": "stable",
                                        "value": "CONNECT",
                                    },
                                    {
                                        "brief": "DELETE method.",
                                        "id": "delete",
                                        "stability": "stable",
                                        "value": "DELETE",
                                    },
                                    {
                                        "brief": "GET method.",
                                        "id": "get",
                                        "stability": "stable",
                                        "value": "GET",
                                    },
                                    {
                                        "brief": "HEAD method.",
                                        "id": "head",
                                        "stability": "stable",
                                        "value": "HEAD",
                                    },
                                    {
                                        "brief": "OPTIONS method.",
                                        "id": "options",
                                        "stability": "stable",
                                        "value": "OPTIONS",
                                    },
                                    {
                                        "brief": "PATCH method.",
                                        "id": "patch",
                                        "stability": "stable",
                                        "value": "PATCH",
                                    },
                                    {
                                        "brief": "POST method.",
                                        "id": "post",
                                        "stability": "stable",
                                        "value": "POST",
                                    },
                                    {
                                        "brief": "PUT method.",
                                        "id": "put",
                                        "stability": "stable",
                                        "value": "PUT",
                                    },
                                    {
                                        "brief": "TRACE method.",
                                        "id": "trace",
                                        "stability": "stable",
                                        "value": "TRACE",
                                    },
                                    {
                                        "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                            "examples": [
                                200,
                            ],
                            "name": "http.response.status_code",
                            "requirement_level": {
                                "conditionally_required": "If and only if one was received/sent.",
                            },
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                            "examples": [
                                "http",
                                "spdy",
                            ],
                            "name": "network.protocol.name",
                            "note": "The value SHOULD be normalized to lowercase.",
                            "requirement_level": {
                                "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The actual version of the protocol used for network communication.",
                            "examples": [
                                "1.0",
                                "1.1",
                                "2",
                                "3",
                            ],
                            "name": "network.protocol.version",
                            "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Host identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "If an HTTP client request is explicitly made to an IP address, e.g. `http://x.x.x.x:8080`, then `server.address` SHOULD be the IP address `x.x.x.x`. A DNS lookup SHOULD NOT be used.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).\n",
                            "examples": [
                                "/users/{id}",
                                "/users/:id",
                                "/users?id={id}",
                            ],
                            "name": "url.template",
                            "note": "The `url.template` MUST have low cardinality. It is not usually available on HTTP clients, but may be known by the application or specialized HTTP instrumentation.\n",
                            "requirement_level": "opt_in",
                            "stability": "development",
                            "type": "string",
                        },
                    ],
                    "brief": "Duration of HTTP client requests.",
                    "id": "metric.http.client.request.duration",
                    "instrument": "histogram",
                    "lineage": {
                        "attributes": {
                            "error.type": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.error",
                            },
                            "http.request.method": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.response.status_code": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "network.protocol.name": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "network.protocol.version": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                ],
                                "source_group": "registry.network",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                            "url.template": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.client.request.duration",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": "metric",
                    "unit": "s",
                },
                {
                    "attributes": [
                        {
                            "brief": "Describes a class of error the operation ended with.\n",
                            "examples": [
                                "timeout",
                                "java.net.UnknownHostException",
                                "server_certificate_invalid",
                                "500",
                            ],
                            "name": "error.type",
                            "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                            "requirement_level": {
                                "conditionally_required": "If request has ended with an error.",
                            },
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "HTTP request method.",
                            "examples": [
                                "GET",
                                "POST",
                                "HEAD",
                            ],
                            "name": "http.request.method",
                            "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "CONNECT method.",
                                        "id": "connect",
                                        "stability": "stable",
                                        "value": "CONNECT",
                                    },
                                    {
                                        "brief": "DELETE method.",
                                        "id": "delete",
                                        "stability": "stable",
                                        "value": "DELETE",
                                    },
                                    {
                                        "brief": "GET method.",
                                        "id": "get",
                                        "stability": "stable",
                                        "value": "GET",
                                    },
                                    {
                                        "brief": "HEAD method.",
                                        "id": "head",
                                        "stability": "stable",
                                        "value": "HEAD",
                                    },
                                    {
                                        "brief": "OPTIONS method.",
                                        "id": "options",
                                        "stability": "stable",
                                        "value": "OPTIONS",
                                    },
                                    {
                                        "brief": "PATCH method.",
                                        "id": "patch",
                                        "stability": "stable",
                                        "value": "PATCH",
                                    },
                                    {
                                        "brief": "POST method.",
                                        "id": "post",
                                        "stability": "stable",
                                        "value": "POST",
                                    },
                                    {
                                        "brief": "PUT method.",
                                        "id": "put",
                                        "stability": "stable",
                                        "value": "PUT",
                                    },
                                    {
                                        "brief": "TRACE method.",
                                        "id": "trace",
                                        "stability": "stable",
                                        "value": "TRACE",
                                    },
                                    {
                                        "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                            "examples": [
                                200,
                            ],
                            "name": "http.response.status_code",
                            "requirement_level": {
                                "conditionally_required": "If and only if one was received/sent.",
                            },
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                            "examples": [
                                "http",
                                "spdy",
                            ],
                            "name": "network.protocol.name",
                            "note": "The value SHOULD be normalized to lowercase.",
                            "requirement_level": {
                                "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The actual version of the protocol used for network communication.",
                            "examples": [
                                "1.0",
                                "1.1",
                                "2",
                                "3",
                            ],
                            "name": "network.protocol.version",
                            "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Host identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "If an HTTP client request is explicitly made to an IP address, e.g. `http://x.x.x.x:8080`, then `server.address` SHOULD be the IP address `x.x.x.x`. A DNS lookup SHOULD NOT be used.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).\n",
                            "examples": [
                                "/users/{id}",
                                "/users/:id",
                                "/users?id={id}",
                            ],
                            "name": "url.template",
                            "note": "The `url.template` MUST have low cardinality. It is not usually available on HTTP clients, but may be known by the application or specialized HTTP instrumentation.\n",
                            "requirement_level": {
                                "conditionally_required": "If available.",
                            },
                            "stability": "development",
                            "type": "string",
                        },
                    ],
                    "brief": "Size of HTTP client response bodies.",
                    "id": "metric.http.client.response.body.size",
                    "instrument": "histogram",
                    "lineage": {
                        "attributes": {
                            "error.type": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.error",
                            },
                            "http.request.method": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.response.status_code": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "network.protocol.name": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "network.protocol.version": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                ],
                                "source_group": "registry.network",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                            "url.template": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.client.response.body.size",
                    "note": "The size of the response payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length) header. For requests using transport encoding, this should be the compressed size.\n",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "metric",
                    "unit": "By",
                },
                {
                    "attributes": [
                        {
                            "brief": "HTTP request method.",
                            "examples": [
                                "GET",
                                "POST",
                                "HEAD",
                            ],
                            "name": "http.request.method",
                            "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "CONNECT method.",
                                        "id": "connect",
                                        "stability": "stable",
                                        "value": "CONNECT",
                                    },
                                    {
                                        "brief": "DELETE method.",
                                        "id": "delete",
                                        "stability": "stable",
                                        "value": "DELETE",
                                    },
                                    {
                                        "brief": "GET method.",
                                        "id": "get",
                                        "stability": "stable",
                                        "value": "GET",
                                    },
                                    {
                                        "brief": "HEAD method.",
                                        "id": "head",
                                        "stability": "stable",
                                        "value": "HEAD",
                                    },
                                    {
                                        "brief": "OPTIONS method.",
                                        "id": "options",
                                        "stability": "stable",
                                        "value": "OPTIONS",
                                    },
                                    {
                                        "brief": "PATCH method.",
                                        "id": "patch",
                                        "stability": "stable",
                                        "value": "PATCH",
                                    },
                                    {
                                        "brief": "POST method.",
                                        "id": "post",
                                        "stability": "stable",
                                        "value": "POST",
                                    },
                                    {
                                        "brief": "PUT method.",
                                        "id": "put",
                                        "stability": "stable",
                                        "value": "PUT",
                                    },
                                    {
                                        "brief": "TRACE method.",
                                        "id": "trace",
                                        "stability": "stable",
                                        "value": "TRACE",
                                    },
                                    {
                                        "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "Name of the local HTTP server that received the request.\n",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port of the local HTTP server that received the request.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                    ],
                    "brief": "Number of active HTTP server requests.",
                    "id": "metric.http.server.active_requests",
                    "instrument": "updowncounter",
                    "lineage": {
                        "attributes": {
                            "http.request.method": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.server.active_requests",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "metric",
                    "unit": "{request}",
                },
                {
                    "attributes": [
                        {
                            "brief": "Describes a class of error the operation ended with.\n",
                            "examples": [
                                "timeout",
                                "java.net.UnknownHostException",
                                "server_certificate_invalid",
                                "500",
                            ],
                            "name": "error.type",
                            "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                            "requirement_level": {
                                "conditionally_required": "If request has ended with an error.",
                            },
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "HTTP request method.",
                            "examples": [
                                "GET",
                                "POST",
                                "HEAD",
                            ],
                            "name": "http.request.method",
                            "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "CONNECT method.",
                                        "id": "connect",
                                        "stability": "stable",
                                        "value": "CONNECT",
                                    },
                                    {
                                        "brief": "DELETE method.",
                                        "id": "delete",
                                        "stability": "stable",
                                        "value": "DELETE",
                                    },
                                    {
                                        "brief": "GET method.",
                                        "id": "get",
                                        "stability": "stable",
                                        "value": "GET",
                                    },
                                    {
                                        "brief": "HEAD method.",
                                        "id": "head",
                                        "stability": "stable",
                                        "value": "HEAD",
                                    },
                                    {
                                        "brief": "OPTIONS method.",
                                        "id": "options",
                                        "stability": "stable",
                                        "value": "OPTIONS",
                                    },
                                    {
                                        "brief": "PATCH method.",
                                        "id": "patch",
                                        "stability": "stable",
                                        "value": "PATCH",
                                    },
                                    {
                                        "brief": "POST method.",
                                        "id": "post",
                                        "stability": "stable",
                                        "value": "POST",
                                    },
                                    {
                                        "brief": "PUT method.",
                                        "id": "put",
                                        "stability": "stable",
                                        "value": "PUT",
                                    },
                                    {
                                        "brief": "TRACE method.",
                                        "id": "trace",
                                        "stability": "stable",
                                        "value": "TRACE",
                                    },
                                    {
                                        "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                            "examples": [
                                200,
                            ],
                            "name": "http.response.status_code",
                            "requirement_level": {
                                "conditionally_required": "If and only if one was received/sent.",
                            },
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                            "examples": [
                                "/users/:userID?",
                                "{controller}/{action}/{id?}",
                            ],
                            "name": "http.route",
                            "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                            "requirement_level": {
                                "conditionally_required": "If and only if it's available",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                            "examples": [
                                "http",
                                "spdy",
                            ],
                            "name": "network.protocol.name",
                            "note": "The value SHOULD be normalized to lowercase.",
                            "requirement_level": {
                                "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The actual version of the protocol used for network communication.",
                            "examples": [
                                "1.0",
                                "1.1",
                                "2",
                                "3",
                            ],
                            "name": "network.protocol.version",
                            "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Name of the local HTTP server that received the request.\n",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port of the local HTTP server that received the request.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "Specifies the category of synthetic traffic, such as tests or bots.\n",
                            "name": "user_agent.synthetic.type",
                            "note": "This attribute MAY be derived from the contents of the `user_agent.original` attribute. Components that populate the attribute are responsible for determining what they consider to be synthetic bot or test traffic. This attribute can either be set for self-identification purposes, or on telemetry detected to be generated as a result of a synthetic request. This attribute is useful for distinguishing between genuine client traffic and synthetic traffic generated by bots or tests.\n",
                            "requirement_level": "opt_in",
                            "stability": "development",
                            "type": {
                                "members": [
                                    {
                                        "brief": "Bot source.",
                                        "id": "bot",
                                        "stability": "development",
                                        "value": "bot",
                                    },
                                    {
                                        "brief": "Synthetic test source.",
                                        "id": "test",
                                        "stability": "development",
                                        "value": "test",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "note": "The scheme of the original client request, if known (e.g. from [Forwarded#proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/Forwarded#proto), [X-Forwarded-Proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/X-Forwarded-Proto), or a similar header). Otherwise, the scheme of the immediate peer request.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                    ],
                    "brief": "Size of HTTP server request bodies.",
                    "id": "metric.http.server.request.body.size",
                    "instrument": "histogram",
                    "lineage": {
                        "attributes": {
                            "error.type": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.error",
                            },
                            "http.request.method": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.response.status_code": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.route": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "network.protocol.name": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "network.protocol.version": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                ],
                                "source_group": "registry.network",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                            "user_agent.synthetic.type": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.user_agent.os",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.server.request.body.size",
                    "note": "The size of the request payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length) header. For requests using transport encoding, this should be the compressed size.\n",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "metric",
                    "unit": "By",
                },
                {
                    "attributes": [
                        {
                            "brief": "Describes a class of error the operation ended with.\n",
                            "examples": [
                                "timeout",
                                "java.net.UnknownHostException",
                                "server_certificate_invalid",
                                "500",
                            ],
                            "name": "error.type",
                            "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                            "requirement_level": {
                                "conditionally_required": "If request has ended with an error.",
                            },
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "HTTP request method.",
                            "examples": [
                                "GET",
                                "POST",
                                "HEAD",
                            ],
                            "name": "http.request.method",
                            "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "CONNECT method.",
                                        "id": "connect",
                                        "stability": "stable",
                                        "value": "CONNECT",
                                    },
                                    {
                                        "brief": "DELETE method.",
                                        "id": "delete",
                                        "stability": "stable",
                                        "value": "DELETE",
                                    },
                                    {
                                        "brief": "GET method.",
                                        "id": "get",
                                        "stability": "stable",
                                        "value": "GET",
                                    },
                                    {
                                        "brief": "HEAD method.",
                                        "id": "head",
                                        "stability": "stable",
                                        "value": "HEAD",
                                    },
                                    {
                                        "brief": "OPTIONS method.",
                                        "id": "options",
                                        "stability": "stable",
                                        "value": "OPTIONS",
                                    },
                                    {
                                        "brief": "PATCH method.",
                                        "id": "patch",
                                        "stability": "stable",
                                        "value": "PATCH",
                                    },
                                    {
                                        "brief": "POST method.",
                                        "id": "post",
                                        "stability": "stable",
                                        "value": "POST",
                                    },
                                    {
                                        "brief": "PUT method.",
                                        "id": "put",
                                        "stability": "stable",
                                        "value": "PUT",
                                    },
                                    {
                                        "brief": "TRACE method.",
                                        "id": "trace",
                                        "stability": "stable",
                                        "value": "TRACE",
                                    },
                                    {
                                        "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                            "examples": [
                                200,
                            ],
                            "name": "http.response.status_code",
                            "requirement_level": {
                                "conditionally_required": "If and only if one was received/sent.",
                            },
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                            "examples": [
                                "/users/:userID?",
                                "{controller}/{action}/{id?}",
                            ],
                            "name": "http.route",
                            "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                            "requirement_level": {
                                "conditionally_required": "If and only if it's available",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                            "examples": [
                                "http",
                                "spdy",
                            ],
                            "name": "network.protocol.name",
                            "note": "The value SHOULD be normalized to lowercase.",
                            "requirement_level": {
                                "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The actual version of the protocol used for network communication.",
                            "examples": [
                                "1.0",
                                "1.1",
                                "2",
                                "3",
                            ],
                            "name": "network.protocol.version",
                            "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Name of the local HTTP server that received the request.\n",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port of the local HTTP server that received the request.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "Specifies the category of synthetic traffic, such as tests or bots.\n",
                            "name": "user_agent.synthetic.type",
                            "note": "This attribute MAY be derived from the contents of the `user_agent.original` attribute. Components that populate the attribute are responsible for determining what they consider to be synthetic bot or test traffic. This attribute can either be set for self-identification purposes, or on telemetry detected to be generated as a result of a synthetic request. This attribute is useful for distinguishing between genuine client traffic and synthetic traffic generated by bots or tests.\n",
                            "requirement_level": "opt_in",
                            "stability": "development",
                            "type": {
                                "members": [
                                    {
                                        "brief": "Bot source.",
                                        "id": "bot",
                                        "stability": "development",
                                        "value": "bot",
                                    },
                                    {
                                        "brief": "Synthetic test source.",
                                        "id": "test",
                                        "stability": "development",
                                        "value": "test",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "note": "The scheme of the original client request, if known (e.g. from [Forwarded#proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/Forwarded#proto), [X-Forwarded-Proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/X-Forwarded-Proto), or a similar header). Otherwise, the scheme of the immediate peer request.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                    ],
                    "brief": "Duration of HTTP server requests.",
                    "id": "metric.http.server.request.duration",
                    "instrument": "histogram",
                    "lineage": {
                        "attributes": {
                            "error.type": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.error",
                            },
                            "http.request.method": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.response.status_code": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.route": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "network.protocol.name": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "network.protocol.version": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                ],
                                "source_group": "registry.network",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                            "user_agent.synthetic.type": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.user_agent.os",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.server.request.duration",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": "metric",
                    "unit": "s",
                },
                {
                    "attributes": [
                        {
                            "brief": "Describes a class of error the operation ended with.\n",
                            "examples": [
                                "timeout",
                                "java.net.UnknownHostException",
                                "server_certificate_invalid",
                                "500",
                            ],
                            "name": "error.type",
                            "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                            "requirement_level": {
                                "conditionally_required": "If request has ended with an error.",
                            },
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "HTTP request method.",
                            "examples": [
                                "GET",
                                "POST",
                                "HEAD",
                            ],
                            "name": "http.request.method",
                            "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": {
                                "members": [
                                    {
                                        "brief": "CONNECT method.",
                                        "id": "connect",
                                        "stability": "stable",
                                        "value": "CONNECT",
                                    },
                                    {
                                        "brief": "DELETE method.",
                                        "id": "delete",
                                        "stability": "stable",
                                        "value": "DELETE",
                                    },
                                    {
                                        "brief": "GET method.",
                                        "id": "get",
                                        "stability": "stable",
                                        "value": "GET",
                                    },
                                    {
                                        "brief": "HEAD method.",
                                        "id": "head",
                                        "stability": "stable",
                                        "value": "HEAD",
                                    },
                                    {
                                        "brief": "OPTIONS method.",
                                        "id": "options",
                                        "stability": "stable",
                                        "value": "OPTIONS",
                                    },
                                    {
                                        "brief": "PATCH method.",
                                        "id": "patch",
                                        "stability": "stable",
                                        "value": "PATCH",
                                    },
                                    {
                                        "brief": "POST method.",
                                        "id": "post",
                                        "stability": "stable",
                                        "value": "POST",
                                    },
                                    {
                                        "brief": "PUT method.",
                                        "id": "put",
                                        "stability": "stable",
                                        "value": "PUT",
                                    },
                                    {
                                        "brief": "TRACE method.",
                                        "id": "trace",
                                        "stability": "stable",
                                        "value": "TRACE",
                                    },
                                    {
                                        "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                        "id": "other",
                                        "stability": "stable",
                                        "value": "_OTHER",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                            "examples": [
                                200,
                            ],
                            "name": "http.response.status_code",
                            "requirement_level": {
                                "conditionally_required": "If and only if one was received/sent.",
                            },
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                            "examples": [
                                "/users/:userID?",
                                "{controller}/{action}/{id?}",
                            ],
                            "name": "http.route",
                            "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                            "requirement_level": {
                                "conditionally_required": "If and only if it's available",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                            "examples": [
                                "http",
                                "spdy",
                            ],
                            "name": "network.protocol.name",
                            "note": "The value SHOULD be normalized to lowercase.",
                            "requirement_level": {
                                "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                            },
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "The actual version of the protocol used for network communication.",
                            "examples": [
                                "1.0",
                                "1.1",
                                "2",
                                "3",
                            ],
                            "name": "network.protocol.version",
                            "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                            "requirement_level": "recommended",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Name of the local HTTP server that received the request.\n",
                            "examples": [
                                "example.com",
                                "10.1.2.80",
                                "/tmp/my.sock",
                            ],
                            "name": "server.address",
                            "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "string",
                        },
                        {
                            "brief": "Port of the local HTTP server that received the request.\n",
                            "examples": [
                                80,
                                8080,
                                443,
                            ],
                            "name": "server.port",
                            "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                            "requirement_level": "opt_in",
                            "stability": "stable",
                            "type": "int",
                        },
                        {
                            "brief": "Specifies the category of synthetic traffic, such as tests or bots.\n",
                            "name": "user_agent.synthetic.type",
                            "note": "This attribute MAY be derived from the contents of the `user_agent.original` attribute. Components that populate the attribute are responsible for determining what they consider to be synthetic bot or test traffic. This attribute can either be set for self-identification purposes, or on telemetry detected to be generated as a result of a synthetic request. This attribute is useful for distinguishing between genuine client traffic and synthetic traffic generated by bots or tests.\n",
                            "requirement_level": "opt_in",
                            "stability": "development",
                            "type": {
                                "members": [
                                    {
                                        "brief": "Bot source.",
                                        "id": "bot",
                                        "stability": "development",
                                        "value": "bot",
                                    },
                                    {
                                        "brief": "Synthetic test source.",
                                        "id": "test",
                                        "stability": "development",
                                        "value": "test",
                                    },
                                ],
                            },
                        },
                        {
                            "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                            "examples": [
                                "http",
                                "https",
                            ],
                            "name": "url.scheme",
                            "note": "The scheme of the original client request, if known (e.g. from [Forwarded#proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/Forwarded#proto), [X-Forwarded-Proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/X-Forwarded-Proto), or a similar header). Otherwise, the scheme of the immediate peer request.\n",
                            "requirement_level": "required",
                            "stability": "stable",
                            "type": "string",
                        },
                    ],
                    "brief": "Size of HTTP server response bodies.",
                    "id": "metric.http.server.response.body.size",
                    "instrument": "histogram",
                    "lineage": {
                        "attributes": {
                            "error.type": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.error",
                            },
                            "http.request.method": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.response.status_code": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "http.route": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.http",
                            },
                            "network.protocol.name": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "requirement_level",
                                ],
                                "source_group": "registry.network",
                            },
                            "network.protocol.version": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                ],
                                "source_group": "registry.network",
                            },
                            "server.address": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "server.port": {
                                "inherited_fields": [
                                    "examples",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "brief",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.server",
                            },
                            "url.scheme": {
                                "inherited_fields": [
                                    "brief",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "examples",
                                    "note",
                                    "requirement_level",
                                ],
                                "source_group": "registry.url",
                            },
                            "user_agent.synthetic.type": {
                                "inherited_fields": [
                                    "brief",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.user_agent.os",
                            },
                        },
                        "provenance": {
                            "path": "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[model]/http/metrics.yaml",
                            "registry_id": "otel",
                        },
                    },
                    "metric_name": "http.server.response.body.size",
                    "note": "The size of the response payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length) header. For requests using transport encoding, this should be the compressed size.\n",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "metric",
                    "unit": "By",
                },
            ],
            "root_namespace": "http",
        },
        "for_each_attr": <macro for_each_attr>,
        "h": ,
    },
    env: Environment {
        globals: {
            "concat_if": weaver_forge::extensions::util::concat_if,
            "cycler": minijinja_contrib::globals::cycler,
            "debug": minijinja::functions::builtins::debug,
            "dict": minijinja::functions::builtins::dict,
            "joiner": minijinja_contrib::globals::joiner,
            "namespace": minijinja::functions::builtins::namespace,
            "params": {
                "included_namespaces": "http",
                "params": {
                    "included_namespaces": "http",
                },
            },
            "range": minijinja::functions::builtins::range,
            "template": {},
        },
        tests: [
            "!=",
            "<",
            "<=",
            "==",
            ">",
            ">=",
            "array",
            "boolean",
            "defined",
            "deprecated",
            "divisibleby",
            "endingwith",
            "enum",
            "enum_type",
            "eq",
            "equalto",
            "escaped",
            "even",
            "experimental",
            "false",
            "filter",
            "float",
            "ge",
            "greaterthan",
            "gt",
            "in",
            "int",
            "integer",
            "iterable",
            "le",
            "lessthan",
            "lower",
            "lt",
            "mapping",
            "ne",
            "none",
            "number",
            "odd",
            "safe",
            "sameas",
            "sequence",
            "simple_type",
            "stable",
            "startingwith",
            "string",
            "template_type",
            "test",
            "true",
            "undefined",
            "upper",
        ],
        filters: [
            "abs",
            "acronym",
            "ansi_bg_black",
            "ansi_bg_blue",
            "ansi_bg_bright_black",
            "ansi_bg_bright_blue",
            "ansi_bg_bright_cyan",
            "ansi_bg_bright_green",
            "ansi_bg_bright_magenta",
            "ansi_bg_bright_red",
            "ansi_bg_bright_white",
            "ansi_bg_bright_yellow",
            "ansi_bg_cyan",
            "ansi_bg_green",
            "ansi_bg_magenta",
            "ansi_bg_red",
            "ansi_bg_white",
            "ansi_bg_yellow",
            "ansi_black",
            "ansi_blue",
            "ansi_bold",
            "ansi_bright_black",
            "ansi_bright_blue",
            "ansi_bright_cyan",
            "ansi_bright_green",
            "ansi_bright_magenta",
            "ansi_bright_red",
            "ansi_bright_white",
            "ansi_bright_yellow",
            "ansi_cyan",
            "ansi_green",
            "ansi_italic",
            "ansi_magenta",
            "ansi_red",
            "ansi_strikethrough",
            "ansi_underline",
            "ansi_white",
            "ansi_yellow",
            "attr",
            "attribute_id",
            "attribute_namespace",
            "attribute_registry_file",
            "attribute_registry_namespace",
            "attribute_registry_title",
            "attribute_sort",
            "batch",
            "body_fields",
            "bool",
            "camel_case",
            "camel_case_const",
            "capitalize",
            "capitalize_first",
            "chain",
            "comment",
            "comment_with_prefix",
            "count",
            "d",
            "default",
            "dictsort",
            "e",
            "enum_type",
            "escape",
            "filesizeformat",
            "first",
            "flatten",
            "float",
            "groupby",
            "indent",
            "instantiated_type",
            "int",
            "items",
            "join",
            "kebab_case",
            "kebab_case_const",
            "last",
            "length",
            "lines",
            "list",
            "lower",
            "lower_case",
            "map",
            "map_text",
            "markdown_to_html",
            "max",
            "metric_namespace",
            "min",
            "not_required",
            "pascal_case",
            "pascal_case_const",
            "pluralize",
            "pprint",
            "print_member_value",
            "regex_replace",
            "reject",
            "rejectattr",
            "replace",
            "required",
            "reverse",
            "round",
            "safe",
            "screaming_kebab_case",
            "screaming_snake_case",
            "screaming_snake_case_const",
            "select",
            "selectattr",
            "slice",
            "snake_case",
            "snake_case_const",
            "sort",
            "split",
            "split_id",
            "string",
            "striptags",
            "sum",
            "title",
            "title_case",
            "tojson",
            "toyaml",
            "trim",
            "truncate",
            "type_mapping",
            "unique",
            "upper",
            "upper_case",
            "urlencode",
        ],
        templates: [
            "helpers.j2",
            "vec.go.j2",
        ],
    },
}
*/