package k8s

import (
       "github.com/prometheus/client_golang/prometheus"
)

// Describes the condition of a particular Node.
type NodeConditionStatus struct {
     *prometheus.GaugeVec
     extra NodeConditionStatusExtra
}



func NewNodeConditionStatus() NodeConditionStatus {
     labels := []string{AttrNodeConditionStatus("").Key(),AttrNodeConditionType("").Key(),}
     return NodeConditionStatus{ 
        GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
            Name: "k8s_node_condition_status",
            Help: "Describes the condition of a particular Node.",
     }, labels)}
}

func (m NodeConditionStatus) With(nodeConditionStatus AttrNodeConditionStatus,nodeConditionKind AttrNodeConditionType,extras ...interface {}) prometheus.Gauge {
    return m.GaugeVec.WithLabelValues(nodeConditionStatus.Value(),nodeConditionKind.Value(),)
}

// Deprecated: Use [NodeConditionStatus.With] instead
func (m NodeConditionStatus) WithLabelValues(lvs ...string) prometheus.Gauge {
    return m.GaugeVec.WithLabelValues(lvs...)
}



type NodeConditionStatusExtra struct {

}



/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "NodeConditionStatusExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "node.condition.status",
        "Type": "NodeConditionStatus",
        "attributes": [
            {
                "brief": "The status of the condition, one of True, False, Unknown.\n",
                "examples": [
                    "true",
                    "false",
                    "unknown",
                ],
                "name": "k8s.node.condition.status",
                "note": "This attribute aligns with the `status` field of the\n[NodeCondition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#nodecondition-v1-core)\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "id": "condition_true",
                            "stability": "development",
                            "value": "true",
                        },
                        {
                            "id": "condition_false",
                            "stability": "development",
                            "value": "false",
                        },
                        {
                            "id": "condition_unknown",
                            "stability": "development",
                            "value": "unknown",
                        },
                    ],
                },
            },
            {
                "brief": "The condition type of a K8s Node.\n",
                "examples": [
                    "Ready",
                    "DiskPressure",
                ],
                "name": "k8s.node.condition.type",
                "note": "K8s Node conditions as described\nby [K8s documentation](https://v1-32.docs.kubernetes.io/docs/reference/node/node-status/#condition).\n\nThis attribute aligns with the `type` field of the\n[NodeCondition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#nodecondition-v1-core)\n\nThe set of possible values is not limited to those listed here. Managed Kubernetes environments,\nor custom controllers MAY introduce additional node condition types.\nWhen this occurs, the exact value as reported by the Kubernetes API SHOULD be used.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "The node is healthy and ready to accept pods",
                            "id": "ready",
                            "stability": "development",
                            "value": "Ready",
                        },
                        {
                            "brief": "Pressure exists on the disk size—that is, if the disk capacity is low",
                            "id": "disk_pressure",
                            "stability": "development",
                            "value": "DiskPressure",
                        },
                        {
                            "brief": "Pressure exists on the node memory—that is, if the node memory is low",
                            "id": "memory_pressure",
                            "stability": "development",
                            "value": "MemoryPressure",
                        },
                        {
                            "brief": "Pressure exists on the processes—that is, if there are too many processes on the node",
                            "id": "pid_pressure",
                            "stability": "development",
                            "value": "PIDPressure",
                        },
                        {
                            "brief": "The network for the node is not correctly configured",
                            "id": "network_unavailable",
                            "stability": "development",
                            "value": "NetworkUnavailable",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "int",
                },
            },
            "attributes": [
                {
                    "brief": "The condition type of a K8s Node.\n",
                    "examples": [
                        "Ready",
                        "DiskPressure",
                    ],
                    "name": "k8s.node.condition.type",
                    "note": "K8s Node conditions as described\nby [K8s documentation](https://v1-32.docs.kubernetes.io/docs/reference/node/node-status/#condition).\n\nThis attribute aligns with the `type` field of the\n[NodeCondition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#nodecondition-v1-core)\n\nThe set of possible values is not limited to those listed here. Managed Kubernetes environments,\nor custom controllers MAY introduce additional node condition types.\nWhen this occurs, the exact value as reported by the Kubernetes API SHOULD be used.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The node is healthy and ready to accept pods",
                                "id": "ready",
                                "stability": "development",
                                "value": "Ready",
                            },
                            {
                                "brief": "Pressure exists on the disk size—that is, if the disk capacity is low",
                                "id": "disk_pressure",
                                "stability": "development",
                                "value": "DiskPressure",
                            },
                            {
                                "brief": "Pressure exists on the node memory—that is, if the node memory is low",
                                "id": "memory_pressure",
                                "stability": "development",
                                "value": "MemoryPressure",
                            },
                            {
                                "brief": "Pressure exists on the processes—that is, if there are too many processes on the node",
                                "id": "pid_pressure",
                                "stability": "development",
                                "value": "PIDPressure",
                            },
                            {
                                "brief": "The network for the node is not correctly configured",
                                "id": "network_unavailable",
                                "stability": "development",
                                "value": "NetworkUnavailable",
                            },
                        ],
                    },
                },
                {
                    "brief": "The status of the condition, one of True, False, Unknown.\n",
                    "examples": [
                        "true",
                        "false",
                        "unknown",
                    ],
                    "name": "k8s.node.condition.status",
                    "note": "This attribute aligns with the `status` field of the\n[NodeCondition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#nodecondition-v1-core)\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "condition_true",
                                "stability": "development",
                                "value": "true",
                            },
                            {
                                "id": "condition_false",
                                "stability": "development",
                                "value": "false",
                            },
                            {
                                "id": "condition_unknown",
                                "stability": "development",
                                "value": "unknown",
                            },
                        ],
                    },
                },
            ],
            "brief": "Describes the condition of a particular Node.",
            "entity_associations": [
                "k8s.node",
            ],
            "id": "metric.k8s.node.condition.status",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "k8s.node.condition.status": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.k8s",
                    },
                    "k8s.node.condition.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.k8s",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/k8s/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "k8s.node.condition.status",
            "note": "All possible node condition pairs (type and status) will be reported at each time interval to avoid missing metrics. Condition pairs corresponding to the current conditions' statuses will be non-zero.\n",
            "root_namespace": "k8s",
            "stability": "development",
            "type": "metric",
            "unit": "{node}",
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