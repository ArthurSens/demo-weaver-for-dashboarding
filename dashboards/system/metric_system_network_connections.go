package system

import (
       "github.com/prometheus/client_golang/prometheus"
      "github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/network"
)

// Deprecated, use `system.network.connection.count` instead.
type NetworkConnections struct {
     *prometheus.GaugeVec
     extra NetworkConnectionsExtra
}



func NewNetworkConnections() NetworkConnections {
     labels := []string{network.AttrConnectionState("").Key(),network.AttrInterfaceName("").Key(),network.AttrTransport("").Key(),}
     return NetworkConnections{ 
        GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
            Name: "system_network_connections",
            Help: "Deprecated, use `system.network.connection.count` instead.",
     }, labels)}
}

func (m NetworkConnections) With(extras ...interface {NetworkConnectionState() network.AttrConnectionState;NetworkInterfaceName() network.AttrInterfaceName;NetworkTransport() network.AttrTransport;}) prometheus.Gauge { 
        if extras == nil { extras = append(extras, m.extra) }
        extra := extras[0]
    
    return m.GaugeVec.WithLabelValues(extra.NetworkConnectionState().Value(),extra.NetworkInterfaceName().Value(),extra.NetworkTransport().Value(),)
}

// Deprecated: Use [NetworkConnections.With] instead
func (m NetworkConnections) WithLabelValues(lvs ...string) prometheus.Gauge {
    return m.GaugeVec.WithLabelValues(lvs...)
}

func (a NetworkConnections) WithNetworkConnectionState (attr interface { NetworkConnectionState() network.AttrConnectionState } ) NetworkConnections {
    a.extra.AttrNetworkConnectionState = attr.NetworkConnectionState()
    return a
}
func (a NetworkConnections) WithNetworkInterfaceName (attr interface { NetworkInterfaceName() network.AttrInterfaceName } ) NetworkConnections {
    a.extra.AttrNetworkInterfaceName = attr.NetworkInterfaceName()
    return a
}
func (a NetworkConnections) WithNetworkTransport (attr interface { NetworkTransport() network.AttrTransport } ) NetworkConnections {
    a.extra.AttrNetworkTransport = attr.NetworkTransport()
    return a
}


type NetworkConnectionsExtra struct {
// The state of network connection
    AttrNetworkConnectionState network.AttrConnectionState `otel:"network.connection.state"`// The network interface name
    AttrNetworkInterfaceName network.AttrInterfaceName `otel:"network.interface.name"`// [OSI transport layer] or [inter-process communication method]
//
// [OSI transport layer]: https://wikipedia.org/wiki/Transport_layer
// [inter-process communication method]: https://wikipedia.org/wiki/Inter-process_communication
    AttrNetworkTransport network.AttrTransport `otel:"network.transport"`
}

func (a NetworkConnectionsExtra) NetworkConnectionState() network.AttrConnectionState {return a.AttrNetworkConnectionState}
func (a NetworkConnectionsExtra) NetworkInterfaceName() network.AttrInterfaceName {return a.AttrNetworkInterfaceName}
func (a NetworkConnectionsExtra) NetworkTransport() network.AttrTransport {return a.AttrNetworkTransport}


/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "NetworkConnectionsExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "network.connections",
        "Type": "NetworkConnections",
        "attributes": [
            {
                "brief": "The state of network connection",
                "examples": [
                    "close_wait",
                ],
                "name": "network.connection.state",
                "note": "Connection states are defined as part of the [rfc9293](https://datatracker.ietf.org/doc/html/rfc9293#section-3.3.2)",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "id": "closed",
                            "stability": "development",
                            "value": "closed",
                        },
                        {
                            "id": "close_wait",
                            "stability": "development",
                            "value": "close_wait",
                        },
                        {
                            "id": "closing",
                            "stability": "development",
                            "value": "closing",
                        },
                        {
                            "id": "established",
                            "stability": "development",
                            "value": "established",
                        },
                        {
                            "id": "fin_wait_1",
                            "stability": "development",
                            "value": "fin_wait_1",
                        },
                        {
                            "id": "fin_wait_2",
                            "stability": "development",
                            "value": "fin_wait_2",
                        },
                        {
                            "id": "last_ack",
                            "stability": "development",
                            "value": "last_ack",
                        },
                        {
                            "id": "listen",
                            "stability": "development",
                            "value": "listen",
                        },
                        {
                            "id": "syn_received",
                            "stability": "development",
                            "value": "syn_received",
                        },
                        {
                            "id": "syn_sent",
                            "stability": "development",
                            "value": "syn_sent",
                        },
                        {
                            "id": "time_wait",
                            "stability": "development",
                            "value": "time_wait",
                        },
                    ],
                },
            },
            {
                "brief": "The network interface name.",
                "examples": [
                    "lo",
                    "eth0",
                ],
                "name": "network.interface.name",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "[OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).\n",
                "examples": [
                    "tcp",
                    "udp",
                ],
                "name": "network.transport",
                "note": "The value SHOULD be normalized to lowercase.\n\nConsider always setting the transport when setting a port number, since\na port number is ambiguous without knowing the transport. For example\ndifferent processes could be listening on TCP port 12345 and UDP port 12345.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "TCP",
                            "id": "tcp",
                            "stability": "stable",
                            "value": "tcp",
                        },
                        {
                            "brief": "UDP",
                            "id": "udp",
                            "stability": "stable",
                            "value": "udp",
                        },
                        {
                            "brief": "Named or anonymous pipe.",
                            "id": "pipe",
                            "stability": "stable",
                            "value": "pipe",
                        },
                        {
                            "brief": "Unix domain socket",
                            "id": "unix",
                            "stability": "stable",
                            "value": "unix",
                        },
                        {
                            "brief": "QUIC",
                            "id": "quic",
                            "stability": "stable",
                            "value": "quic",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "[OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).\n",
                    "examples": [
                        "tcp",
                        "udp",
                    ],
                    "name": "network.transport",
                    "note": "The value SHOULD be normalized to lowercase.\n\nConsider always setting the transport when setting a port number, since\na port number is ambiguous without knowing the transport. For example\ndifferent processes could be listening on TCP port 12345 and UDP port 12345.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "TCP",
                                "id": "tcp",
                                "stability": "stable",
                                "value": "tcp",
                            },
                            {
                                "brief": "UDP",
                                "id": "udp",
                                "stability": "stable",
                                "value": "udp",
                            },
                            {
                                "brief": "Named or anonymous pipe.",
                                "id": "pipe",
                                "stability": "stable",
                                "value": "pipe",
                            },
                            {
                                "brief": "Unix domain socket",
                                "id": "unix",
                                "stability": "stable",
                                "value": "unix",
                            },
                            {
                                "brief": "QUIC",
                                "id": "quic",
                                "stability": "stable",
                                "value": "quic",
                            },
                        ],
                    },
                },
                {
                    "brief": "The network interface name.",
                    "examples": [
                        "lo",
                        "eth0",
                    ],
                    "name": "network.interface.name",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The state of network connection",
                    "examples": [
                        "close_wait",
                    ],
                    "name": "network.connection.state",
                    "note": "Connection states are defined as part of the [rfc9293](https://datatracker.ietf.org/doc/html/rfc9293#section-3.3.2)",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "closed",
                                "stability": "development",
                                "value": "closed",
                            },
                            {
                                "id": "close_wait",
                                "stability": "development",
                                "value": "close_wait",
                            },
                            {
                                "id": "closing",
                                "stability": "development",
                                "value": "closing",
                            },
                            {
                                "id": "established",
                                "stability": "development",
                                "value": "established",
                            },
                            {
                                "id": "fin_wait_1",
                                "stability": "development",
                                "value": "fin_wait_1",
                            },
                            {
                                "id": "fin_wait_2",
                                "stability": "development",
                                "value": "fin_wait_2",
                            },
                            {
                                "id": "last_ack",
                                "stability": "development",
                                "value": "last_ack",
                            },
                            {
                                "id": "listen",
                                "stability": "development",
                                "value": "listen",
                            },
                            {
                                "id": "syn_received",
                                "stability": "development",
                                "value": "syn_received",
                            },
                            {
                                "id": "syn_sent",
                                "stability": "development",
                                "value": "syn_sent",
                            },
                            {
                                "id": "time_wait",
                                "stability": "development",
                                "value": "time_wait",
                            },
                        ],
                    },
                },
            ],
            "brief": "Deprecated, use `system.network.connection.count` instead.",
            "deprecated": {
                "note": "Replaced by `system.network.connection.count`.",
                "reason": "renamed",
                "renamed_to": "system.network.connection.count",
            },
            "entity_associations": [
                "host",
            ],
            "id": "metric.system.network.connections",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "network.connection.state": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.network",
                    },
                    "network.interface.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.network",
                    },
                    "network.transport": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.network",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/system/deprecated/metrics-deprecated.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "system.network.connections",
            "root_namespace": "system",
            "stability": "development",
            "type": "metric",
            "unit": "{connection}",
        },
        "for_each_attr": <macro for_each_attr>,
        "module": "github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards",
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
                "params": {},
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
            "vec.go.j2",
        ],
    },
}
*/