package system

import (
       "github.com/prometheus/client_golang/prometheus"
      "github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/cpu"
)

// Seconds each logical CPU spent on each mode.
type CpuTime struct {
     *prometheus.CounterVec
     extra CpuTimeExtra
}



func NewCpuTime() CpuTime {
     labels := []string{cpu.AttrLogicalNumber("").Key(),cpu.AttrMode("").Key(),}
     return CpuTime{ 
        CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
            Name: "system_cpu_time",
            Help: "Seconds each logical CPU spent on each mode.",
     }, labels)}
}

func (m CpuTime) With(extras ...interface {CpuLogicalNumber() cpu.AttrLogicalNumber;CpuMode() cpu.AttrMode;}) prometheus.Counter { 
        if extras == nil { extras = append(extras, m.extra) }
        extra := extras[0]
    
    return m.CounterVec.WithLabelValues(extra.CpuLogicalNumber().Value(),extra.CpuMode().Value(),)
}

// Deprecated: Use [CpuTime.With] instead
func (m CpuTime) WithLabelValues(lvs ...string) prometheus.Counter {
    return m.CounterVec.WithLabelValues(lvs...)
}

func (a CpuTime) WithCpuLogicalNumber (attr interface { CpuLogicalNumber() cpu.AttrLogicalNumber } ) CpuTime {
    a.extra.AttrCpuLogicalNumber = attr.CpuLogicalNumber()
    return a
}
func (a CpuTime) WithCpuMode (attr interface { CpuMode() cpu.AttrMode } ) CpuTime {
    a.extra.AttrCpuMode = attr.CpuMode()
    return a
}


type CpuTimeExtra struct {
// The logical CPU number [0..n-1]
    AttrCpuLogicalNumber cpu.AttrLogicalNumber `otel:"cpu.logical_number"`// The mode of the CPU
    AttrCpuMode cpu.AttrMode `otel:"cpu.mode"`
}

func (a CpuTimeExtra) CpuLogicalNumber() cpu.AttrLogicalNumber {return a.AttrCpuLogicalNumber}
func (a CpuTimeExtra) CpuMode() cpu.AttrMode {return a.AttrCpuMode}


/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "CpuTimeExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "cpu.time",
        "Type": "CpuTime",
        "attributes": [
            {
                "brief": "The logical CPU number [0..n-1]",
                "examples": [
                    1,
                ],
                "name": "cpu.logical_number",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "int",
            },
            {
                "brief": "The mode of the CPU",
                "examples": [
                    "user",
                    "system",
                ],
                "name": "cpu.mode",
                "note": "Following states SHOULD be used: `user`, `system`, `nice`, `idle`, `iowait`, `interrupt`, `steal`",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "User",
                            "id": "user",
                            "stability": "development",
                            "value": "user",
                        },
                        {
                            "brief": "System",
                            "id": "system",
                            "stability": "development",
                            "value": "system",
                        },
                        {
                            "brief": "Nice",
                            "id": "nice",
                            "stability": "development",
                            "value": "nice",
                        },
                        {
                            "brief": "Idle",
                            "id": "idle",
                            "stability": "development",
                            "value": "idle",
                        },
                        {
                            "brief": "IO Wait",
                            "id": "iowait",
                            "stability": "development",
                            "value": "iowait",
                        },
                        {
                            "brief": "Interrupt",
                            "id": "interrupt",
                            "stability": "development",
                            "value": "interrupt",
                        },
                        {
                            "brief": "Steal",
                            "id": "steal",
                            "stability": "development",
                            "value": "steal",
                        },
                        {
                            "brief": "Kernel",
                            "id": "kernel",
                            "stability": "development",
                            "value": "kernel",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "annotations": {
                "code_generation": {
                    "metric_value_type": "double",
                },
            },
            "attributes": [
                {
                    "brief": "The logical CPU number [0..n-1]",
                    "examples": [
                        1,
                    ],
                    "name": "cpu.logical_number",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The mode of the CPU",
                    "examples": [
                        "user",
                        "system",
                    ],
                    "name": "cpu.mode",
                    "note": "Following states SHOULD be used: `user`, `system`, `nice`, `idle`, `iowait`, `interrupt`, `steal`",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "User",
                                "id": "user",
                                "stability": "development",
                                "value": "user",
                            },
                            {
                                "brief": "System",
                                "id": "system",
                                "stability": "development",
                                "value": "system",
                            },
                            {
                                "brief": "Nice",
                                "id": "nice",
                                "stability": "development",
                                "value": "nice",
                            },
                            {
                                "brief": "Idle",
                                "id": "idle",
                                "stability": "development",
                                "value": "idle",
                            },
                            {
                                "brief": "IO Wait",
                                "id": "iowait",
                                "stability": "development",
                                "value": "iowait",
                            },
                            {
                                "brief": "Interrupt",
                                "id": "interrupt",
                                "stability": "development",
                                "value": "interrupt",
                            },
                            {
                                "brief": "Steal",
                                "id": "steal",
                                "stability": "development",
                                "value": "steal",
                            },
                            {
                                "brief": "Kernel",
                                "id": "kernel",
                                "stability": "development",
                                "value": "kernel",
                            },
                        ],
                    },
                },
            ],
            "brief": "Seconds each logical CPU spent on each mode.",
            "entity_associations": [
                "host",
            ],
            "id": "metric.system.cpu.time",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "cpu.logical_number": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.cpu",
                    },
                    "cpu.mode": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                        ],
                        "source_group": "registry.cpu",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/system/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "system.cpu.time",
            "root_namespace": "system",
            "stability": "development",
            "type": "metric",
            "unit": "s",
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