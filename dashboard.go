package main


// Import the appropriate Grafana Foundation SDK packages
import (
	"encoding/json"
	"log"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)


func main() {
  // Define a data source reference for our testdata data source
  prometheusDataSourceRef := dashboard.DataSourceRef{
	Type: cog.ToPtr("prometheus"),
    Uid:  cog.ToPtr("prometheus"),
  }

  // Define our dashboard as strongly typed code
  builder := dashboard.NewDashboardBuilder("http").
    Uid("http-dashboard").
	Tags([]string{"built-with-weaver"}).
	Refresh("5m").
	Time("now-1h", "now").
	Timezone(common.TimeZoneBrowser).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Number of active HTTP requests.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_client_active_requests[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("The duration of the successfully established outbound HTTP connections.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_client_connection_duration[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Number of outbound HTTP connections that are currently active or idle on the client.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_client_open_connections[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Size of HTTP client request bodies.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_client_request_body_size[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Duration of HTTP client requests.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_client_request_duration[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Size of HTTP client response bodies.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_client_response_body_size[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Number of active HTTP server requests.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_server_active_requests[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Size of HTTP server request bodies.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_server_request_body_size[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Duration of HTTP server requests.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_server_request_duration[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).WithPanel(
		timeseries.NewPanelBuilder().
			Title("Size of HTTP server response bodies.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(metric_http_server_response_body_size[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	)
    
  // Build the dashboard - errors in configuration will be thrown here
  dashboard, err := builder.Build()
  if err != nil {
    log.Fatalf("failed to build dashboard: %v", err)
  }

  // Output the generated dashboard as JSON
  dashboardJson, err := json.MarshalIndent(dashboard, "", "  ")
  if err != nil {
    log.Fatalf("failed to marshal dashboard: %v", err)
  }

  log.Printf(string(dashboardJson))
}

/* {
    "metrics": [
        {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "int",
                },
            },
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
                    "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "In HTTP/1.1, when the [request target](https://www.rfc-editor.org/rfc/rfc9112.html#name-request-target)\nis passed in its [absolute-form](https://www.rfc-editor.org/rfc/rfc9112.html#section-3.2.2),\nthe `server.address` SHOULD match the host component of the request target.\n\nIn all other cases, `server.address` SHOULD match the host component of the\n`Host` header in HTTP/1.1 or the `:authority` pseudo-header in HTTP/2 and HTTP/3.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Server port number.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "In the case of HTTP/1.1, when the [request target](https://www.rfc-editor.org/rfc/rfc9112.html#name-request-target)\nis passed in its [absolute-form](https://www.rfc-editor.org/rfc/rfc9112.html#section-3.2.2),\nthe `server.port` SHOULD match the port component of the request target.\n\nIn all other cases, `server.port` SHOULD match the port component of the\n`Host` header in HTTP/1.1 or the `:authority` pseudo-header in HTTP/2 and HTTP/3.\n",
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
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.client.active_requests",
            "root_namespace": "http",
            "stability": "development",
            "type": "metric",
            "unit": "{request}",
        },
        {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "double",
                },
            },
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
                    "brief": "Server port number.",
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.client.connection.duration",
            "root_namespace": "http",
            "stability": "development",
            "type": "metric",
            "unit": "s",
        },
        {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "int",
                },
            },
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
                    "brief": "Server port number.",
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.client.open_connections",
            "root_namespace": "http",
            "stability": "development",
            "type": "metric",
            "unit": "{connection}",
        },
        {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "int",
                },
            },
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
                    "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "In HTTP/1.1, when the [request target](https://www.rfc-editor.org/rfc/rfc9112.html#name-request-target)\nis passed in its [absolute-form](https://www.rfc-editor.org/rfc/rfc9112.html#section-3.2.2),\nthe `server.address` SHOULD match the host component of the request target.\n\nIn all other cases, `server.address` SHOULD match the host component of the\n`Host` header in HTTP/1.1 or the `:authority` pseudo-header in HTTP/2 and HTTP/3.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Server port number.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "In the case of HTTP/1.1, when the [request target](https://www.rfc-editor.org/rfc/rfc9112.html#name-request-target)\nis passed in its [absolute-form](https://www.rfc-editor.org/rfc/rfc9112.html#section-3.2.2),\nthe `server.port` SHOULD match the port component of the request target.\n\nIn all other cases, `server.port` SHOULD match the port component of the\n`Host` header in HTTP/1.1 or the `:authority` pseudo-header in HTTP/2 and HTTP/3.\n",
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
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
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
            "annotations": {
                "code_generation": {
                    "metric_value_type": "double",
                },
            },
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
                    "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "In HTTP/1.1, when the [request target](https://www.rfc-editor.org/rfc/rfc9112.html#name-request-target)\nis passed in its [absolute-form](https://www.rfc-editor.org/rfc/rfc9112.html#section-3.2.2),\nthe `server.address` SHOULD match the host component of the request target.\n\nIn all other cases, `server.address` SHOULD match the host component of the\n`Host` header in HTTP/1.1 or the `:authority` pseudo-header in HTTP/2 and HTTP/3.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Server port number.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "In the case of HTTP/1.1, when the [request target](https://www.rfc-editor.org/rfc/rfc9112.html#name-request-target)\nis passed in its [absolute-form](https://www.rfc-editor.org/rfc/rfc9112.html#section-3.2.2),\nthe `server.port` SHOULD match the port component of the request target.\n\nIn all other cases, `server.port` SHOULD match the port component of the\n`Host` header in HTTP/1.1 or the `:authority` pseudo-header in HTTP/2 and HTTP/3.\n",
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
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.client.request.duration",
            "root_namespace": "http",
            "stability": "stable",
            "type": "metric",
            "unit": "s",
        },
        {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "int",
                },
            },
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
                    "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "In HTTP/1.1, when the [request target](https://www.rfc-editor.org/rfc/rfc9112.html#name-request-target)\nis passed in its [absolute-form](https://www.rfc-editor.org/rfc/rfc9112.html#section-3.2.2),\nthe `server.address` SHOULD match the host component of the request target.\n\nIn all other cases, `server.address` SHOULD match the host component of the\n`Host` header in HTTP/1.1 or the `:authority` pseudo-header in HTTP/2 and HTTP/3.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Server port number.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "In the case of HTTP/1.1, when the [request target](https://www.rfc-editor.org/rfc/rfc9112.html#name-request-target)\nis passed in its [absolute-form](https://www.rfc-editor.org/rfc/rfc9112.html#section-3.2.2),\nthe `server.port` SHOULD match the port component of the request target.\n\nIn all other cases, `server.port` SHOULD match the port component of the\n`Host` header in HTTP/1.1 or the `:authority` pseudo-header in HTTP/2 and HTTP/3.\n",
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
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
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
            "annotations": {
                "code_generation": {
                    "metric_value_type": "int",
                },
            },
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.server.active_requests",
            "root_namespace": "http",
            "stability": "development",
            "type": "metric",
            "unit": "{request}",
        },
        {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "int",
                },
            },
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
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
            "annotations": {
                "code_generation": {
                    "metric_value_type": "double",
                },
            },
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.server.request.duration",
            "root_namespace": "http",
            "stability": "stable",
            "type": "metric",
            "unit": "s",
        },
        {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "int",
                },
            },
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
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
} */