// Code generated from semantic convention specification. DO NOT EDIT.

// Package httpconv provides types and functionality for OpenTelemetry semantic
// conventions in the "library" namespace.
package library

import (
  "github.com/prometheus/client_golang/prometheus"
)

type Attribute interface {
  ID() string
  Value() string
}

// LocationAttr is an attribute conforming to the library.location semantic conventions. It represents the location identifier of the library
type LocationAttr string
func (a LocationAttr) ID() string {
  return "library.location"
}
func (a LocationAttr) Value() string {
  return string(a)
}

// BooksIndexed is an instrument used to record metric values conforming to the "library.books_indexed" semantic conventions. It represents a counter of the number of books indexed
type BooksIndexed struct {
  *prometheus.CounterVec
}

// NewBooksIndexed returns a new BooksIndexed instrument
func NewBooksIndexed() BooksIndexed {
  labels := []string{
     "library.location",
     }
  return BooksIndexed{
     CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
         Name: "library_books_indexed",
         Help: "A counter of the number of books indexed",
  }, labels)}
}


type BooksIndexedAttr interface {
  Attribute
  implBooksIndexed()
}
func (a LocationAttr) implBooksIndexed() {}

func (m BooksIndexed) With(
    location LocationAttr,
    extra ...BooksIndexedAttr,
) prometheus.Counter {
    labels := prometheus.Labels{
        "library.location":  location.Value(),
    }
    for _,v := range extra {
        labels[v.ID()] = v.Value()
    }
    
    return m.CounterVec.With(labels)
}
/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "all_attrs": [
            {
                "brief": "Location identifier of the library",
                "examples": [
                    "Vienna/1110",
                ],
                "name": "library.location",
                "requirement_level": "required",
                "stability": "development",
                "type": "string",
            },
        ],
        "ctx": {
            "metrics": [
                {
                    "attributes": [
                        {
                            "brief": "Location identifier of the library",
                            "examples": [
                                "Vienna/1110",
                            ],
                            "name": "library.location",
                            "requirement_level": "required",
                            "stability": "development",
                            "type": "string",
                        },
                    ],
                    "brief": "A counter of the number of books indexed",
                    "id": "metric.library.books_indexed",
                    "instrument": "counter",
                    "lineage": {
                        "attributes": {
                            "library.location": {
                                "inherited_fields": [
                                    "brief",
                                    "examples",
                                    "note",
                                    "stability",
                                ],
                                "locally_overridden_fields": [
                                    "requirement_level",
                                ],
                                "source_group": "registry.library",
                            },
                        },
                        "provenance": {
                            "path": "./semconv/metrics.yaml",
                            "registry_id": "demo-library",
                        },
                    },
                    "metric_name": "library.books_indexed",
                    "root_namespace": "library",
                    "stability": "development",
                    "type": "metric",
                    "unit": "1",
                },
            ],
            "root_namespace": "library",
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
                "included_namespaces": [],
                "params": {
                    "included_namespaces": [],
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
            "vec.go.j2",
            "helpers.j2",
        ],
    },
}
*/