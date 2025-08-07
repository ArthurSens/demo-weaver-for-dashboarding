package hw




// An identifier for the hardware component, unique within the monitored host
type AttrId string // hw.id

func (AttrId) Development() {}
func (AttrId) Recommended() {}
func (AttrId) Key() string { return "hw_id" }
func (a AttrId) Value() string { return string(a) }





// An easily-recognizable name for the hardware component
type AttrName string // hw.name

func (AttrName) Development() {}
func (AttrName) Recommended() {}
func (AttrName) Key() string { return "hw_name" }
func (a AttrName) Value() string { return string(a) }





// Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
type AttrParent string // hw.parent

func (AttrParent) Development() {}
func (AttrParent) Recommended() {}
func (AttrParent) Key() string { return "hw_parent" }
func (a AttrParent) Value() string { return string(a) }





// The current state of the component
type AttrState string // hw.state

func (AttrState) Development() {}
func (AttrState) Recommended() {}
func (AttrState) Key() string { return "hw_state" }
func (a AttrState) Value() string { return string(a) }

const StateOk AttrState = "ok"
const StateDegraded AttrState = "degraded"
const StateFailed AttrState = "failed"




// Type of the component
//
// Describes the category of the hardware component for which `hw.state` is being reported. For example, `hw.type=temperature` along with `hw.state=degraded` would indicate that the temperature of the hardware component has been reported as `degraded`
type AttrType string // hw.type

func (AttrType) Development() {}
func (AttrType) Recommended() {}
func (AttrType) Key() string { return "hw_type" }
func (a AttrType) Value() string { return string(a) }

const TypeBattery AttrType = "battery"
const TypeCpu AttrType = "cpu"
const TypeDiskController AttrType = "disk_controller"
const TypeEnclosure AttrType = "enclosure"
const TypeFan AttrType = "fan"
const TypeGpu AttrType = "gpu"
const TypeLogicalDisk AttrType = "logical_disk"
const TypeMemory AttrType = "memory"
const TypeNetwork AttrType = "network"
const TypePhysicalDisk AttrType = "physical_disk"
const TypePowerSupply AttrType = "power_supply"
const TypeTapeDrive AttrType = "tape_drive"
const TypeTemperature AttrType = "temperature"
const TypeVoltage AttrType = "voltage"



/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "An identifier for the hardware component, unique within the monitored host\n",
                    "examples": [
                        "win32battery_battery_testsysa33_1",
                    ],
                    "name": "hw.id",
                    "requirement_level": "recommended",
                    "root_namespace": "hw",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "An easily-recognizable name for the hardware component\n",
                    "examples": [
                        "eth0",
                    ],
                    "name": "hw.name",
                    "requirement_level": "recommended",
                    "root_namespace": "hw",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)\n",
                    "examples": [
                        "dellStorage_perc_0",
                    ],
                    "name": "hw.parent",
                    "requirement_level": "recommended",
                    "root_namespace": "hw",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The current state of the component\n",
                    "name": "hw.state",
                    "requirement_level": "recommended",
                    "root_namespace": "hw",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Ok",
                                "id": "ok",
                                "stability": "development",
                                "value": "ok",
                            },
                            {
                                "brief": "Degraded",
                                "id": "degraded",
                                "stability": "development",
                                "value": "degraded",
                            },
                            {
                                "brief": "Failed",
                                "id": "failed",
                                "stability": "development",
                                "value": "failed",
                            },
                        ],
                    },
                },
                {
                    "brief": "Type of the component\n",
                    "name": "hw.type",
                    "note": "Describes the category of the hardware component for which `hw.state` is being reported. For example, `hw.type=temperature` along with `hw.state=degraded` would indicate that the temperature of the hardware component has been reported as `degraded`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "hw",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Battery",
                                "id": "battery",
                                "stability": "development",
                                "value": "battery",
                            },
                            {
                                "brief": "CPU",
                                "id": "cpu",
                                "stability": "development",
                                "value": "cpu",
                            },
                            {
                                "brief": "Disk controller",
                                "id": "disk_controller",
                                "stability": "development",
                                "value": "disk_controller",
                            },
                            {
                                "brief": "Enclosure",
                                "id": "enclosure",
                                "stability": "development",
                                "value": "enclosure",
                            },
                            {
                                "brief": "Fan",
                                "id": "fan",
                                "stability": "development",
                                "value": "fan",
                            },
                            {
                                "brief": "GPU",
                                "id": "gpu",
                                "stability": "development",
                                "value": "gpu",
                            },
                            {
                                "brief": "Logical disk",
                                "id": "logical_disk",
                                "stability": "development",
                                "value": "logical_disk",
                            },
                            {
                                "brief": "Memory",
                                "id": "memory",
                                "stability": "development",
                                "value": "memory",
                            },
                            {
                                "brief": "Network",
                                "id": "network",
                                "stability": "development",
                                "value": "network",
                            },
                            {
                                "brief": "Physical disk",
                                "id": "physical_disk",
                                "stability": "development",
                                "value": "physical_disk",
                            },
                            {
                                "brief": "Power supply",
                                "id": "power_supply",
                                "stability": "development",
                                "value": "power_supply",
                            },
                            {
                                "brief": "Tape drive",
                                "id": "tape_drive",
                                "stability": "development",
                                "value": "tape_drive",
                            },
                            {
                                "brief": "Temperature",
                                "id": "temperature",
                                "stability": "development",
                                "value": "temperature",
                            },
                            {
                                "brief": "Voltage",
                                "id": "voltage",
                                "stability": "development",
                                "value": "voltage",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "hw",
        },
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
            "attr.go.j2",
        ],
    },
} */