package otel

import (
       "github.com/prometheus/client_golang/prometheus"
)

// The number of created spans.
type SdkSpanStarted struct {
     *prometheus.CounterVec
     extra SdkSpanStartedExtra
}



func NewSdkSpanStarted() SdkSpanStarted {
     labels := []string{AttrSpanParentOrigin("").Key(),AttrSpanSamplingResult("").Key(),}
     return SdkSpanStarted{ 
        CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
            Name: "otel_sdk_span_started",
            Help: "The number of created spans.",
     }, labels)}
}

func (m SdkSpanStarted) With(extras ...interface {OtelSpanParentOrigin() AttrSpanParentOrigin;OtelSpanSamplingResult() AttrSpanSamplingResult;}) prometheus.Counter { 
        if extras == nil { extras = append(extras, m.extra) }
        extra := extras[0]
    
    return m.CounterVec.WithLabelValues(extra.OtelSpanParentOrigin().Value(),extra.OtelSpanSamplingResult().Value(),)
}

// Deprecated: Use [SdkSpanStarted.With] instead
func (m SdkSpanStarted) WithLabelValues(lvs ...string) prometheus.Counter {
    return m.CounterVec.WithLabelValues(lvs...)
}

func (a SdkSpanStarted) WithSpanParentOrigin (attr interface { OtelSpanParentOrigin() AttrSpanParentOrigin } ) SdkSpanStarted {
    a.extra.AttrSpanParentOrigin = attr.OtelSpanParentOrigin()
    return a
}
func (a SdkSpanStarted) WithSpanSamplingResult (attr interface { OtelSpanSamplingResult() AttrSpanSamplingResult } ) SdkSpanStarted {
    a.extra.AttrSpanSamplingResult = attr.OtelSpanSamplingResult()
    return a
}


type SdkSpanStartedExtra struct {
// Determines whether the span has a parent span, and if so, [whether it is a remote parent]
//
// [whether it is a remote parent]: https://opentelemetry.io/docs/specs/otel/trace/api/#isremote
    AttrSpanParentOrigin AttrSpanParentOrigin `otel:"otel.span.parent.origin"`// The result value of the sampler for this span
    AttrSpanSamplingResult AttrSpanSamplingResult `otel:"otel.span.sampling_result"`
}

func (a SdkSpanStartedExtra) OtelSpanParentOrigin() AttrSpanParentOrigin {return a.AttrSpanParentOrigin}
func (a SdkSpanStartedExtra) OtelSpanSamplingResult() AttrSpanSamplingResult {return a.AttrSpanSamplingResult}


/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "SdkSpanStartedExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "sdk.span.started",
        "Type": "SdkSpanStarted",
        "attributes": [
            {
                "brief": "Determines whether the span has a parent span, and if so, [whether it is a remote parent](https://opentelemetry.io/docs/specs/otel/trace/api/#isremote)",
                "name": "otel.span.parent.origin",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "The span does not have a parent, it is a root span",
                            "id": "none",
                            "stability": "development",
                            "value": "none",
                        },
                        {
                            "brief": "The span has a parent and the parent's span context [isRemote()](https://opentelemetry.io/docs/specs/otel/trace/api/#isremote) is false",
                            "id": "local",
                            "stability": "development",
                            "value": "local",
                        },
                        {
                            "brief": "The span has a parent and the parent's span context [isRemote()](https://opentelemetry.io/docs/specs/otel/trace/api/#isremote) is true",
                            "id": "remote",
                            "stability": "development",
                            "value": "remote",
                        },
                    ],
                },
            },
            {
                "brief": "The result value of the sampler for this span",
                "name": "otel.span.sampling_result",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "The span is not sampled and not recording",
                            "id": "drop",
                            "stability": "development",
                            "value": "DROP",
                        },
                        {
                            "brief": "The span is not sampled, but recording",
                            "id": "record_only",
                            "stability": "development",
                            "value": "RECORD_ONLY",
                        },
                        {
                            "brief": "The span is sampled and recording",
                            "id": "record_and_sample",
                            "stability": "development",
                            "value": "RECORD_AND_SAMPLE",
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
                    "brief": "The result value of the sampler for this span",
                    "name": "otel.span.sampling_result",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The span is not sampled and not recording",
                                "id": "drop",
                                "stability": "development",
                                "value": "DROP",
                            },
                            {
                                "brief": "The span is not sampled, but recording",
                                "id": "record_only",
                                "stability": "development",
                                "value": "RECORD_ONLY",
                            },
                            {
                                "brief": "The span is sampled and recording",
                                "id": "record_and_sample",
                                "stability": "development",
                                "value": "RECORD_AND_SAMPLE",
                            },
                        ],
                    },
                },
                {
                    "brief": "Determines whether the span has a parent span, and if so, [whether it is a remote parent](https://opentelemetry.io/docs/specs/otel/trace/api/#isremote)",
                    "name": "otel.span.parent.origin",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The span does not have a parent, it is a root span",
                                "id": "none",
                                "stability": "development",
                                "value": "none",
                            },
                            {
                                "brief": "The span has a parent and the parent's span context [isRemote()](https://opentelemetry.io/docs/specs/otel/trace/api/#isremote) is false",
                                "id": "local",
                                "stability": "development",
                                "value": "local",
                            },
                            {
                                "brief": "The span has a parent and the parent's span context [isRemote()](https://opentelemetry.io/docs/specs/otel/trace/api/#isremote) is true",
                                "id": "remote",
                                "stability": "development",
                                "value": "remote",
                            },
                        ],
                    },
                },
            ],
            "brief": "The number of created spans.",
            "id": "metric.otel.sdk.span.started",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "otel.span.parent.origin": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.otel",
                    },
                    "otel.span.sampling_result": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.otel",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/otel/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "otel.sdk.span.started",
            "note": "Implementations MUST record this metric for all spans, even for non-recording ones.\n",
            "root_namespace": "otel",
            "stability": "development",
            "type": "metric",
            "unit": "{span}",
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