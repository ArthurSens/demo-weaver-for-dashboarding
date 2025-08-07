package k8s

import (
       "github.com/prometheus/client_golang/prometheus"
)

// The object count limits in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.
type ResourcequotaObjectCountHard struct {
     *prometheus.GaugeVec
     extra ResourcequotaObjectCountHardExtra
}



func NewResourcequotaObjectCountHard() ResourcequotaObjectCountHard {
     labels := []string{AttrResourcequotaResourceName("").Key(),}
     return ResourcequotaObjectCountHard{ 
        GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
            Name: "k8s_resourcequota_object_count_hard",
            Help: "The object count limits in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.",
     }, labels)}
}

func (m ResourcequotaObjectCountHard) With(resourcequotaResourceName AttrResourcequotaResourceName,extras ...interface {}) prometheus.Gauge {
    return m.GaugeVec.WithLabelValues(resourcequotaResourceName.Value(),)
}

// Deprecated: Use [ResourcequotaObjectCountHard.With] instead
func (m ResourcequotaObjectCountHard) WithLabelValues(lvs ...string) prometheus.Gauge {
    return m.GaugeVec.WithLabelValues(lvs...)
}



type ResourcequotaObjectCountHardExtra struct {

}



/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ResourcequotaObjectCountHardExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "resourcequota.object_count.hard",
        "Type": "ResourcequotaObjectCountHard",
        "attributes": [
            {
                "brief": "The name of the K8s resource a resource quota defines.\n",
                "examples": [
                    "count/replicationcontrollers",
                ],
                "name": "k8s.resourcequota.resource_name",
                "note": "The value for this attribute can be either the full `count/<resource>[.<group>]` string (e.g., count/deployments.apps, count/pods), or, for certain core Kubernetes resources, just the resource name (e.g., pods, services, configmaps). Both forms are supported by Kubernetes for object count quotas. See [Kubernetes Resource Quotas documentation](https://kubernetes.io/docs/concepts/policy/resource-quotas/#object-count-quota) for more details.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": "string",
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
                    "brief": "The name of the K8s resource a resource quota defines.\n",
                    "examples": [
                        "count/replicationcontrollers",
                    ],
                    "name": "k8s.resourcequota.resource_name",
                    "note": "The value for this attribute can be either the full `count/<resource>[.<group>]` string (e.g., count/deployments.apps, count/pods), or, for certain core Kubernetes resources, just the resource name (e.g., pods, services, configmaps). Both forms are supported by Kubernetes for object count quotas. See [Kubernetes Resource Quotas documentation](https://kubernetes.io/docs/concepts/policy/resource-quotas/#object-count-quota) for more details.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "The object count limits in a specific namespace.\nThe value represents the configured quota limit of the resource in the namespace.\n",
            "entity_associations": [
                "k8s.resourcequota",
            ],
            "id": "metric.k8s.resourcequota.object_count.hard",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "k8s.resourcequota.resource_name": {
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
            "metric_name": "k8s.resourcequota.object_count.hard",
            "note": "This metric is retrieved from the `hard` field of the\n[K8s ResourceQuotaStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcequotastatus-v1-core).\n",
            "root_namespace": "k8s",
            "stability": "development",
            "type": "metric",
            "unit": "{object}",
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