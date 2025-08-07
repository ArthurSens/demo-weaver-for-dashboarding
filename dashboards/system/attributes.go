package system




// Deprecated, use `cpu.logical_number` instead
type AttrCpuLogicalNumber string // system.cpu.logical_number

func (AttrCpuLogicalNumber) Development() {}
func (AttrCpuLogicalNumber) Recommended() {}
func (AttrCpuLogicalNumber) Key() string { return "system_cpu_logical_number" }
func (a AttrCpuLogicalNumber) Value() string { return string(a) }





// Deprecated, use `cpu.mode` instead
type AttrCpuState string // system.cpu.state

func (AttrCpuState) Development() {}
func (AttrCpuState) Recommended() {}
func (AttrCpuState) Key() string { return "system_cpu_state" }
func (a AttrCpuState) Value() string { return string(a) }

const CpuStateUser AttrCpuState = "user"
const CpuStateSystem AttrCpuState = "system"
const CpuStateNice AttrCpuState = "nice"
const CpuStateIdle AttrCpuState = "idle"
const CpuStateIowait AttrCpuState = "iowait"
const CpuStateInterrupt AttrCpuState = "interrupt"
const CpuStateSteal AttrCpuState = "steal"




// The device identifier
type AttrDevice string // system.device

func (AttrDevice) Development() {}
func (AttrDevice) Recommended() {}
func (AttrDevice) Key() string { return "system_device" }
func (a AttrDevice) Value() string { return string(a) }





// The filesystem mode
type AttrFilesystemMode string // system.filesystem.mode

func (AttrFilesystemMode) Development() {}
func (AttrFilesystemMode) Recommended() {}
func (AttrFilesystemMode) Key() string { return "system_filesystem_mode" }
func (a AttrFilesystemMode) Value() string { return string(a) }





// The filesystem mount path
type AttrFilesystemMountpoint string // system.filesystem.mountpoint

func (AttrFilesystemMountpoint) Development() {}
func (AttrFilesystemMountpoint) Recommended() {}
func (AttrFilesystemMountpoint) Key() string { return "system_filesystem_mountpoint" }
func (a AttrFilesystemMountpoint) Value() string { return string(a) }





// The filesystem state
type AttrFilesystemState string // system.filesystem.state

func (AttrFilesystemState) Development() {}
func (AttrFilesystemState) Recommended() {}
func (AttrFilesystemState) Key() string { return "system_filesystem_state" }
func (a AttrFilesystemState) Value() string { return string(a) }

const FilesystemStateUsed AttrFilesystemState = "used"
const FilesystemStateFree AttrFilesystemState = "free"
const FilesystemStateReserved AttrFilesystemState = "reserved"




// The filesystem type
type AttrFilesystemType string // system.filesystem.type

func (AttrFilesystemType) Development() {}
func (AttrFilesystemType) Recommended() {}
func (AttrFilesystemType) Key() string { return "system_filesystem_type" }
func (a AttrFilesystemType) Value() string { return string(a) }

const FilesystemTypeFat32 AttrFilesystemType = "fat32"
const FilesystemTypeExfat AttrFilesystemType = "exfat"
const FilesystemTypeNtfs AttrFilesystemType = "ntfs"
const FilesystemTypeRefs AttrFilesystemType = "refs"
const FilesystemTypeHfsplus AttrFilesystemType = "hfsplus"
const FilesystemTypeExt4 AttrFilesystemType = "ext4"




// The memory state
type AttrMemoryState string // system.memory.state

func (AttrMemoryState) Development() {}
func (AttrMemoryState) Recommended() {}
func (AttrMemoryState) Key() string { return "system_memory_state" }
func (a AttrMemoryState) Value() string { return string(a) }

const MemoryStateUsed AttrMemoryState = "used"
const MemoryStateFree AttrMemoryState = "free"
const MemoryStateShared AttrMemoryState = "shared"
const MemoryStateBuffers AttrMemoryState = "buffers"
const MemoryStateCached AttrMemoryState = "cached"




// Deprecated, use `network.connection.state` instead
type AttrNetworkState string // system.network.state

func (AttrNetworkState) Development() {}
func (AttrNetworkState) Recommended() {}
func (AttrNetworkState) Key() string { return "system_network_state" }
func (a AttrNetworkState) Value() string { return string(a) }

const NetworkStateClose AttrNetworkState = "close"
const NetworkStateCloseWait AttrNetworkState = "close_wait"
const NetworkStateClosing AttrNetworkState = "closing"
const NetworkStateDelete AttrNetworkState = "delete"
const NetworkStateEstablished AttrNetworkState = "established"
const NetworkStateFinWait1 AttrNetworkState = "fin_wait_1"
const NetworkStateFinWait2 AttrNetworkState = "fin_wait_2"
const NetworkStateLastAck AttrNetworkState = "last_ack"
const NetworkStateListen AttrNetworkState = "listen"
const NetworkStateSynRecv AttrNetworkState = "syn_recv"
const NetworkStateSynSent AttrNetworkState = "syn_sent"
const NetworkStateTimeWait AttrNetworkState = "time_wait"




// The paging access direction
type AttrPagingDirection string // system.paging.direction

func (AttrPagingDirection) Development() {}
func (AttrPagingDirection) Recommended() {}
func (AttrPagingDirection) Key() string { return "system_paging_direction" }
func (a AttrPagingDirection) Value() string { return string(a) }

const PagingDirectionIn AttrPagingDirection = "in"
const PagingDirectionOut AttrPagingDirection = "out"




// The memory paging state
type AttrPagingState string // system.paging.state

func (AttrPagingState) Development() {}
func (AttrPagingState) Recommended() {}
func (AttrPagingState) Key() string { return "system_paging_state" }
func (a AttrPagingState) Value() string { return string(a) }

const PagingStateUsed AttrPagingState = "used"
const PagingStateFree AttrPagingState = "free"




// The memory paging type
type AttrPagingType string // system.paging.type

func (AttrPagingType) Development() {}
func (AttrPagingType) Recommended() {}
func (AttrPagingType) Key() string { return "system_paging_type" }
func (a AttrPagingType) Value() string { return string(a) }

const PagingTypeMajor AttrPagingType = "major"
const PagingTypeMinor AttrPagingType = "minor"




// The process state, e.g., [Linux Process State Codes]
//
// [Linux Process State Codes]: https://man7.org/linux/man-pages/man1/ps.1.html#PROCESS_STATE_CODES
type AttrProcessStatus string // system.process.status

func (AttrProcessStatus) Development() {}
func (AttrProcessStatus) Recommended() {}
func (AttrProcessStatus) Key() string { return "system_process_status" }
func (a AttrProcessStatus) Value() string { return string(a) }

const ProcessStatusRunning AttrProcessStatus = "running"
const ProcessStatusSleeping AttrProcessStatus = "sleeping"
const ProcessStatusStopped AttrProcessStatus = "stopped"
const ProcessStatusDefunct AttrProcessStatus = "defunct"




// Deprecated, use `system.process.status` instead
type AttrProcessesStatus string // system.processes.status

func (AttrProcessesStatus) Development() {}
func (AttrProcessesStatus) Recommended() {}
func (AttrProcessesStatus) Key() string { return "system_processes_status" }
func (a AttrProcessesStatus) Value() string { return string(a) }

const ProcessesStatusRunning AttrProcessesStatus = "running"
const ProcessesStatusSleeping AttrProcessesStatus = "sleeping"
const ProcessesStatusStopped AttrProcessesStatus = "stopped"
const ProcessesStatusDefunct AttrProcessesStatus = "defunct"



/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Deprecated, use `cpu.logical_number` instead.",
                    "examples": [
                        1,
                    ],
                    "name": "system.cpu.logical_number",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `cpu.mode` instead.",
                    "deprecated": {
                        "note": "Replaced by `cpu.mode`.",
                        "reason": "renamed",
                        "renamed_to": "cpu.mode",
                    },
                    "examples": [
                        "idle",
                        "interrupt",
                    ],
                    "name": "system.cpu.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "user",
                                "stability": "development",
                                "value": "user",
                            },
                            {
                                "id": "system",
                                "stability": "development",
                                "value": "system",
                            },
                            {
                                "id": "nice",
                                "stability": "development",
                                "value": "nice",
                            },
                            {
                                "id": "idle",
                                "stability": "development",
                                "value": "idle",
                            },
                            {
                                "id": "iowait",
                                "stability": "development",
                                "value": "iowait",
                            },
                            {
                                "id": "interrupt",
                                "stability": "development",
                                "value": "interrupt",
                            },
                            {
                                "id": "steal",
                                "stability": "development",
                                "value": "steal",
                            },
                        ],
                    },
                },
                {
                    "brief": "The device identifier",
                    "examples": [
                        "(identifier)",
                    ],
                    "name": "system.device",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The filesystem mode",
                    "examples": [
                        "rw, ro",
                    ],
                    "name": "system.filesystem.mode",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The filesystem mount path",
                    "examples": [
                        "/mnt/data",
                    ],
                    "name": "system.filesystem.mountpoint",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The filesystem state",
                    "examples": [
                        "used",
                    ],
                    "name": "system.filesystem.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "used",
                                "stability": "development",
                                "value": "used",
                            },
                            {
                                "id": "free",
                                "stability": "development",
                                "value": "free",
                            },
                            {
                                "id": "reserved",
                                "stability": "development",
                                "value": "reserved",
                            },
                        ],
                    },
                },
                {
                    "brief": "The filesystem type",
                    "examples": [
                        "ext4",
                    ],
                    "name": "system.filesystem.type",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "fat32",
                                "stability": "development",
                                "value": "fat32",
                            },
                            {
                                "id": "exfat",
                                "stability": "development",
                                "value": "exfat",
                            },
                            {
                                "id": "ntfs",
                                "stability": "development",
                                "value": "ntfs",
                            },
                            {
                                "id": "refs",
                                "stability": "development",
                                "value": "refs",
                            },
                            {
                                "id": "hfsplus",
                                "stability": "development",
                                "value": "hfsplus",
                            },
                            {
                                "id": "ext4",
                                "stability": "development",
                                "value": "ext4",
                            },
                        ],
                    },
                },
                {
                    "brief": "The memory state",
                    "examples": [
                        "free",
                        "cached",
                    ],
                    "name": "system.memory.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Actual used virtual memory in bytes.",
                                "id": "used",
                                "note": "Calculation based on the operating system metrics. On Linux, this corresponds to \"MemTotal - MemAvailable\" from /proc/meminfo, which more accurately reflects memory in active use by applications compared to older formulas based on free, cached, and buffers. If MemAvailable is not available, a fallback to those older formulas may be used.\n",
                                "stability": "development",
                                "value": "used",
                            },
                            {
                                "id": "free",
                                "stability": "development",
                                "value": "free",
                            },
                            {
                                "deprecated": {
                                    "note": "Removed, report shared memory usage with `metric.system.memory.shared` metric",
                                    "reason": "unspecified",
                                },
                                "id": "shared",
                                "stability": "development",
                                "value": "shared",
                            },
                            {
                                "id": "buffers",
                                "stability": "development",
                                "value": "buffers",
                            },
                            {
                                "id": "cached",
                                "stability": "development",
                                "value": "cached",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `network.connection.state` instead.",
                    "deprecated": {
                        "note": "Replaced by `network.connection.state`.",
                        "reason": "renamed",
                        "renamed_to": "network.connection.state",
                    },
                    "examples": [
                        "close_wait",
                    ],
                    "name": "system.network.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "close",
                                "stability": "development",
                                "value": "close",
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
                                "id": "delete",
                                "stability": "development",
                                "value": "delete",
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
                                "id": "syn_recv",
                                "stability": "development",
                                "value": "syn_recv",
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
                    "brief": "The paging access direction",
                    "examples": [
                        "in",
                    ],
                    "name": "system.paging.direction",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "in",
                                "stability": "development",
                                "value": "in",
                            },
                            {
                                "id": "out",
                                "stability": "development",
                                "value": "out",
                            },
                        ],
                    },
                },
                {
                    "brief": "The memory paging state",
                    "examples": [
                        "free",
                    ],
                    "name": "system.paging.state",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "used",
                                "stability": "development",
                                "value": "used",
                            },
                            {
                                "id": "free",
                                "stability": "development",
                                "value": "free",
                            },
                        ],
                    },
                },
                {
                    "brief": "The memory paging type",
                    "examples": [
                        "minor",
                    ],
                    "name": "system.paging.type",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "major",
                                "stability": "development",
                                "value": "major",
                            },
                            {
                                "id": "minor",
                                "stability": "development",
                                "value": "minor",
                            },
                        ],
                    },
                },
                {
                    "brief": "The process state, e.g., [Linux Process State Codes](https://man7.org/linux/man-pages/man1/ps.1.html#PROCESS_STATE_CODES)\n",
                    "examples": [
                        "running",
                    ],
                    "name": "system.process.status",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "running",
                                "stability": "development",
                                "value": "running",
                            },
                            {
                                "id": "sleeping",
                                "stability": "development",
                                "value": "sleeping",
                            },
                            {
                                "id": "stopped",
                                "stability": "development",
                                "value": "stopped",
                            },
                            {
                                "id": "defunct",
                                "stability": "development",
                                "value": "defunct",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `system.process.status` instead.",
                    "deprecated": {
                        "note": "Replaced by `system.process.status`.",
                        "reason": "renamed",
                        "renamed_to": "system.process.status",
                    },
                    "examples": [
                        "running",
                    ],
                    "name": "system.processes.status",
                    "requirement_level": "recommended",
                    "root_namespace": "system",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "id": "running",
                                "stability": "development",
                                "value": "running",
                            },
                            {
                                "id": "sleeping",
                                "stability": "development",
                                "value": "sleeping",
                            },
                            {
                                "id": "stopped",
                                "stability": "development",
                                "value": "stopped",
                            },
                            {
                                "id": "defunct",
                                "stability": "development",
                                "value": "defunct",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "system",
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