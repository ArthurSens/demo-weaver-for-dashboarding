package k8s

import (
       "github.com/prometheus/client_golang/prometheus"
)

// Pod volume usage.
type PodVolumeUsage struct {
     *prometheus.GaugeVec
     extra PodVolumeUsageExtra
}



func NewPodVolumeUsage() PodVolumeUsage {
     labels := []string{AttrVolumeName("").Key(),AttrVolumeType("").Key(),}
     return PodVolumeUsage{ 
        GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
            Name: "k8s_pod_volume_usage",
            Help: "Pod volume usage.",
     }, labels)}
}

func (m PodVolumeUsage) With(volumeName AttrVolumeName,extras ...interface {K8sVolumeType() AttrVolumeType;}) prometheus.Gauge { 
        if extras == nil { extras = append(extras, m.extra) }
        extra := extras[0]
    
    return m.GaugeVec.WithLabelValues(volumeName.Value(),extra.K8sVolumeType().Value(),)
}

// Deprecated: Use [PodVolumeUsage.With] instead
func (m PodVolumeUsage) WithLabelValues(lvs ...string) prometheus.Gauge {
    return m.GaugeVec.WithLabelValues(lvs...)
}

func (a PodVolumeUsage) WithVolumeType (attr interface { K8sVolumeType() AttrVolumeType } ) PodVolumeUsage {
    a.extra.AttrVolumeType = attr.K8sVolumeType()
    return a
}


type PodVolumeUsageExtra struct {
// The type of the K8s volume
    AttrVolumeType AttrVolumeType `otel:"k8s.volume.type"`
}

func (a PodVolumeUsageExtra) K8sVolumeType() AttrVolumeType {return a.AttrVolumeType}


/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "PodVolumeUsageExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "pod.volume.usage",
        "Type": "PodVolumeUsage",
        "attributes": [
            {
                "brief": "The name of the K8s volume.\n",
                "examples": [
                    "volume0",
                ],
                "name": "k8s.volume.name",
                "requirement_level": "required",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The type of the K8s volume.\n",
                "examples": [
                    "emptyDir",
                    "persistentVolumeClaim",
                ],
                "name": "k8s.volume.type",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "A [persistentVolumeClaim](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#persistentvolumeclaim) volume",
                            "id": "persistent_volume_claim",
                            "stability": "development",
                            "value": "persistentVolumeClaim",
                        },
                        {
                            "brief": "A [configMap](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#configmap) volume",
                            "id": "config_map",
                            "stability": "development",
                            "value": "configMap",
                        },
                        {
                            "brief": "A [downwardAPI](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#downwardapi) volume",
                            "id": "downward_api",
                            "stability": "development",
                            "value": "downwardAPI",
                        },
                        {
                            "brief": "An [emptyDir](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#emptydir) volume",
                            "id": "empty_dir",
                            "stability": "development",
                            "value": "emptyDir",
                        },
                        {
                            "brief": "A [secret](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#secret) volume",
                            "id": "secret",
                            "stability": "development",
                            "value": "secret",
                        },
                        {
                            "brief": "A [local](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#local) volume",
                            "id": "local",
                            "stability": "development",
                            "value": "local",
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
                    "brief": "The type of the K8s volume.\n",
                    "examples": [
                        "emptyDir",
                        "persistentVolumeClaim",
                    ],
                    "name": "k8s.volume.type",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "A [persistentVolumeClaim](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#persistentvolumeclaim) volume",
                                "id": "persistent_volume_claim",
                                "stability": "development",
                                "value": "persistentVolumeClaim",
                            },
                            {
                                "brief": "A [configMap](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#configmap) volume",
                                "id": "config_map",
                                "stability": "development",
                                "value": "configMap",
                            },
                            {
                                "brief": "A [downwardAPI](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#downwardapi) volume",
                                "id": "downward_api",
                                "stability": "development",
                                "value": "downwardAPI",
                            },
                            {
                                "brief": "An [emptyDir](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#emptydir) volume",
                                "id": "empty_dir",
                                "stability": "development",
                                "value": "emptyDir",
                            },
                            {
                                "brief": "A [secret](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#secret) volume",
                                "id": "secret",
                                "stability": "development",
                                "value": "secret",
                            },
                            {
                                "brief": "A [local](https://v1-30.docs.kubernetes.io/docs/concepts/storage/volumes/#local) volume",
                                "id": "local",
                                "stability": "development",
                                "value": "local",
                            },
                        ],
                    },
                },
                {
                    "brief": "The name of the K8s volume.\n",
                    "examples": [
                        "volume0",
                    ],
                    "name": "k8s.volume.name",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "Pod volume usage.",
            "entity_associations": [
                "k8s.pod",
            ],
            "id": "metric.k8s.pod.volume.usage",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "k8s.volume.name": {
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
                    "k8s.volume.type": {
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
            "metric_name": "k8s.pod.volume.usage",
            "note": "This may not equal capacity - available.\n\nThis metric is derived from the\n[VolumeStats.UsedBytes](https://pkg.go.dev/k8s.io/kubelet@v0.33.0/pkg/apis/stats/v1alpha1#VolumeStats) field\nof the [PodStats](https://pkg.go.dev/k8s.io/kubelet@v0.33.0/pkg/apis/stats/v1alpha1#PodStats) of the\nKubelet's stats API.\n",
            "root_namespace": "k8s",
            "stability": "development",
            "type": "metric",
            "unit": "By",
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