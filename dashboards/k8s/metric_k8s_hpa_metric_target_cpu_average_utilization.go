package k8s

import (
       "github.com/prometheus/client_golang/prometheus"
)

// Target average utilization, in percentage, for CPU resource in HPA config.
type HpaMetricTargetCpuAverageUtilization struct {
     *prometheus.GaugeVec
     extra HpaMetricTargetCpuAverageUtilizationExtra
}



func NewHpaMetricTargetCpuAverageUtilization() HpaMetricTargetCpuAverageUtilization {
     labels := []string{AttrContainerName("").Key(),AttrHpaMetricType("").Key(),}
     return HpaMetricTargetCpuAverageUtilization{ 
        GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
            Name: "k8s_hpa_metric_target_cpu_average_utilization",
            Help: "Target average utilization, in percentage, for CPU resource in HPA config.",
     }, labels)}
}

func (m HpaMetricTargetCpuAverageUtilization) With(extras ...interface {K8sContainerName() AttrContainerName;K8sHpaMetricType() AttrHpaMetricType;}) prometheus.Gauge { 
        if extras == nil { extras = append(extras, m.extra) }
        extra := extras[0]
    
    return m.GaugeVec.WithLabelValues(extra.K8sContainerName().Value(),extra.K8sHpaMetricType().Value(),)
}

// Deprecated: Use [HpaMetricTargetCpuAverageUtilization.With] instead
func (m HpaMetricTargetCpuAverageUtilization) WithLabelValues(lvs ...string) prometheus.Gauge {
    return m.GaugeVec.WithLabelValues(lvs...)
}

func (a HpaMetricTargetCpuAverageUtilization) WithContainerName (attr interface { K8sContainerName() AttrContainerName } ) HpaMetricTargetCpuAverageUtilization {
    a.extra.AttrContainerName = attr.K8sContainerName()
    return a
}
func (a HpaMetricTargetCpuAverageUtilization) WithHpaMetricType (attr interface { K8sHpaMetricType() AttrHpaMetricType } ) HpaMetricTargetCpuAverageUtilization {
    a.extra.AttrHpaMetricType = attr.K8sHpaMetricType()
    return a
}


type HpaMetricTargetCpuAverageUtilizationExtra struct {
// The name of the Container from Pod specification, must be unique within a Pod. Container runtime usually uses different globally unique name (`container.name`)
    AttrContainerName AttrContainerName `otel:"k8s.container.name"`// The type of metric source for the horizontal pod autoscaler
    AttrHpaMetricType AttrHpaMetricType `otel:"k8s.hpa.metric.type"`
}

func (a HpaMetricTargetCpuAverageUtilizationExtra) K8sContainerName() AttrContainerName {return a.AttrContainerName}
func (a HpaMetricTargetCpuAverageUtilizationExtra) K8sHpaMetricType() AttrHpaMetricType {return a.AttrHpaMetricType}


/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "HpaMetricTargetCpuAverageUtilizationExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "hpa.metric.target.cpu.average_utilization",
        "Type": "HpaMetricTargetCpuAverageUtilization",
        "attributes": [
            {
                "brief": "The name of the Container from Pod specification, must be unique within a Pod. Container runtime usually uses different globally unique name (`container.name`).\n",
                "examples": [
                    "redis",
                ],
                "name": "k8s.container.name",
                "requirement_level": {
                    "conditionally_required": "if and only if k8s.hpa.metric.type is ContainerResource.",
                },
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The type of metric source for the horizontal pod autoscaler.\n",
                "examples": [
                    "Resource",
                    "ContainerResource",
                ],
                "name": "k8s.hpa.metric.type",
                "note": "This attribute reflects the `type` field of spec.metrics[] in the HPA.\n",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The type of metric source for the horizontal pod autoscaler.\n",
                    "examples": [
                        "Resource",
                        "ContainerResource",
                    ],
                    "name": "k8s.hpa.metric.type",
                    "note": "This attribute reflects the `type` field of spec.metrics[] in the HPA.\n",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the Container from Pod specification, must be unique within a Pod. Container runtime usually uses different globally unique name (`container.name`).\n",
                    "examples": [
                        "redis",
                    ],
                    "name": "k8s.container.name",
                    "requirement_level": {
                        "conditionally_required": "if and only if k8s.hpa.metric.type is ContainerResource.",
                    },
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "Target average utilization, in percentage, for CPU resource in HPA config.",
            "entity_associations": [
                "k8s.hpa",
                "k8s.namespace",
            ],
            "id": "metric.k8s.hpa.metric.target.cpu.average_utilization",
            "instrument": "gauge",
            "lineage": {
                "attributes": {
                    "k8s.container.name": {
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
                    "k8s.hpa.metric.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.k8s",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/k8s/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "k8s.hpa.metric.target.cpu.average_utilization",
            "note": "This metric aligns with the `averageUtilization` field of the\n[K8s HPA MetricTarget](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#metrictarget-v2-autoscaling).\nIf the type of the metric is [`ContainerResource`](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#support-for-metrics-apis),\nthe `k8s.container.name` attribute MUST be set to identify the specific container within the pod to which the metric applies.\n",
            "root_namespace": "k8s",
            "stability": "development",
            "type": "metric",
            "unit": "1",
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