package gen_ai

import (
       "github.com/prometheus/client_golang/prometheus"
      "github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/server"
)

// Number of input and output tokens used.
type ClientTokenUsage struct {
     *prometheus.HistogramVec
     extra ClientTokenUsageExtra
}



func NewClientTokenUsage() ClientTokenUsage {
     labels := []string{AttrOperationName("").Key(),AttrProviderName("").Key(),AttrTokenType("").Key(),AttrRequestModel("").Key(),server.AttrPort("").Key(),AttrResponseModel("").Key(),server.AttrAddress("").Key(),}
     return ClientTokenUsage{ 
        HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
            Name: "gen_ai_client_token_usage",
            Help: "Number of input and output tokens used.",
     }, labels)}
}

func (m ClientTokenUsage) With(operationName AttrOperationName,providerName AttrProviderName,tokenKind AttrTokenType,extras ...interface {GenAiRequestModel() AttrRequestModel;ServerPort() server.AttrPort;GenAiResponseModel() AttrResponseModel;ServerAddress() server.AttrAddress;}) prometheus.Observer { 
        if extras == nil { extras = append(extras, m.extra) }
        extra := extras[0]
    
    return m.HistogramVec.WithLabelValues(operationName.Value(),providerName.Value(),tokenKind.Value(),extra.GenAiRequestModel().Value(),extra.ServerPort().Value(),extra.GenAiResponseModel().Value(),extra.ServerAddress().Value(),)
}

// Deprecated: Use [ClientTokenUsage.With] instead
func (m ClientTokenUsage) WithLabelValues(lvs ...string) prometheus.Observer {
    return m.HistogramVec.WithLabelValues(lvs...)
}

func (a ClientTokenUsage) WithRequestModel (attr interface { GenAiRequestModel() AttrRequestModel } ) ClientTokenUsage {
    a.extra.AttrRequestModel = attr.GenAiRequestModel()
    return a
}
func (a ClientTokenUsage) WithServerPort (attr interface { ServerPort() server.AttrPort } ) ClientTokenUsage {
    a.extra.AttrServerPort = attr.ServerPort()
    return a
}
func (a ClientTokenUsage) WithResponseModel (attr interface { GenAiResponseModel() AttrResponseModel } ) ClientTokenUsage {
    a.extra.AttrResponseModel = attr.GenAiResponseModel()
    return a
}
func (a ClientTokenUsage) WithServerAddress (attr interface { ServerAddress() server.AttrAddress } ) ClientTokenUsage {
    a.extra.AttrServerAddress = attr.ServerAddress()
    return a
}


type ClientTokenUsageExtra struct {
// The name of the GenAI model a request is being made to
    AttrRequestModel AttrRequestModel `otel:"gen_ai.request.model"`// GenAI server port
    AttrServerPort server.AttrPort `otel:"server.port"`// The name of the model that generated the response
    AttrResponseModel AttrResponseModel `otel:"gen_ai.response.model"`// GenAI server address
    AttrServerAddress server.AttrAddress `otel:"server.address"`
}

func (a ClientTokenUsageExtra) GenAiRequestModel() AttrRequestModel {return a.AttrRequestModel}
func (a ClientTokenUsageExtra) ServerPort() server.AttrPort {return a.AttrServerPort}
func (a ClientTokenUsageExtra) GenAiResponseModel() AttrResponseModel {return a.AttrResponseModel}
func (a ClientTokenUsageExtra) ServerAddress() server.AttrAddress {return a.AttrServerAddress}


/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientTokenUsageExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "client.token.usage",
        "Type": "ClientTokenUsage",
        "attributes": [
            {
                "brief": "The name of the operation being performed.",
                "name": "gen_ai.operation.name",
                "note": "If one of the predefined values applies, but specific system uses a different name it's RECOMMENDED to document it in the semantic conventions for specific GenAI system and use system-specific name in the instrumentation. If a different name is not documented, instrumentation libraries SHOULD use applicable predefined value.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Chat completion operation such as [OpenAI Chat API](https://platform.openai.com/docs/api-reference/chat)",
                            "id": "chat",
                            "stability": "development",
                            "value": "chat",
                        },
                        {
                            "brief": "Multimodal content generation operation such as [Gemini Generate Content](https://ai.google.dev/api/generate-content)",
                            "id": "generate_content",
                            "stability": "development",
                            "value": "generate_content",
                        },
                        {
                            "brief": "Text completions operation such as [OpenAI Completions API (Legacy)](https://platform.openai.com/docs/api-reference/completions)",
                            "id": "text_completion",
                            "stability": "development",
                            "value": "text_completion",
                        },
                        {
                            "brief": "Embeddings operation such as [OpenAI Create embeddings API](https://platform.openai.com/docs/api-reference/embeddings/create)",
                            "id": "embeddings",
                            "stability": "development",
                            "value": "embeddings",
                        },
                        {
                            "brief": "Create GenAI agent",
                            "id": "create_agent",
                            "stability": "development",
                            "value": "create_agent",
                        },
                        {
                            "brief": "Invoke GenAI agent",
                            "id": "invoke_agent",
                            "stability": "development",
                            "value": "invoke_agent",
                        },
                        {
                            "brief": "Execute a tool",
                            "id": "execute_tool",
                            "stability": "development",
                            "value": "execute_tool",
                        },
                    ],
                },
            },
            {
                "brief": "The Generative AI provider as identified by the client or server instrumentation.",
                "name": "gen_ai.provider.name",
                "note": "The attribute SHOULD be set based on the instrumentation's best\nknowledge and may differ from the actual model provider.\n\nMultiple providers, including Azure OpenAI, Gemini, and AI hosting platforms\nare accessible using the OpenAI REST API and corresponding client libraries,\nbut may proxy or host models from different providers.\n\nThe `gen_ai.request.model`, `gen_ai.response.model`, and `server.address`\nattributes may help identify the actual system in use.\n\nThe `gen_ai.provider.name` attribute acts as a discriminator that\nidentifies the GenAI telemetry format flavor specific to that provider\nwithin GenAI semantic conventions.\nIt SHOULD be set consistently with provider-specific attributes and signals.\nFor example, GenAI spans, metrics, and events related to AWS Bedrock\nshould have the `gen_ai.provider.name` set to `aws.bedrock` and include\napplicable `aws.bedrock.*` attributes and are not expected to include\n`openai.*` attributes.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "[OpenAI](https://openai.com/)",
                            "id": "openai",
                            "stability": "development",
                            "value": "openai",
                        },
                        {
                            "brief": "Any Google generative AI endpoint",
                            "id": "gcp.gen_ai",
                            "note": "May be used when specific backend is unknown.\n",
                            "stability": "development",
                            "value": "gcp.gen_ai",
                        },
                        {
                            "brief": "[Vertex AI](https://cloud.google.com/vertex-ai)",
                            "id": "gcp.vertex_ai",
                            "note": "Used when accessing the 'aiplatform.googleapis.com' endpoint.\n",
                            "stability": "development",
                            "value": "gcp.vertex_ai",
                        },
                        {
                            "brief": "[Gemini](https://cloud.google.com/products/gemini)",
                            "id": "gcp.gemini",
                            "note": "Used when accessing the 'generativelanguage.googleapis.com' endpoint. Also known as the AI Studio API.\n",
                            "stability": "development",
                            "value": "gcp.gemini",
                        },
                        {
                            "brief": "[Anthropic](https://www.anthropic.com/)",
                            "id": "anthropic",
                            "stability": "development",
                            "value": "anthropic",
                        },
                        {
                            "brief": "[Cohere](https://cohere.com/)",
                            "id": "cohere",
                            "stability": "development",
                            "value": "cohere",
                        },
                        {
                            "brief": "Azure AI Inference",
                            "id": "azure.ai.inference",
                            "stability": "development",
                            "value": "azure.ai.inference",
                        },
                        {
                            "brief": "[Azure OpenAI](https://azure.microsoft.com/products/ai-services/openai-service/)",
                            "id": "azure.ai.openai",
                            "stability": "development",
                            "value": "azure.ai.openai",
                        },
                        {
                            "brief": "[IBM Watsonx AI](https://www.ibm.com/products/watsonx-ai)",
                            "id": "ibm.watsonx.ai",
                            "stability": "development",
                            "value": "ibm.watsonx.ai",
                        },
                        {
                            "brief": "[AWS Bedrock](https://aws.amazon.com/bedrock)",
                            "id": "aws.bedrock",
                            "stability": "development",
                            "value": "aws.bedrock",
                        },
                        {
                            "brief": "[Perplexity](https://www.perplexity.ai/)",
                            "id": "perplexity",
                            "stability": "development",
                            "value": "perplexity",
                        },
                        {
                            "brief": "[xAI](https://x.ai/)",
                            "id": "x_ai",
                            "stability": "development",
                            "value": "x_ai",
                        },
                        {
                            "brief": "[DeepSeek](https://www.deepseek.com/)",
                            "id": "deepseek",
                            "stability": "development",
                            "value": "deepseek",
                        },
                        {
                            "brief": "[Groq](https://groq.com/)",
                            "id": "groq",
                            "stability": "development",
                            "value": "groq",
                        },
                        {
                            "brief": "[Mistral AI](https://mistral.ai/)",
                            "id": "mistral_ai",
                            "stability": "development",
                            "value": "mistral_ai",
                        },
                    ],
                },
            },
            {
                "brief": "The type of token being counted.",
                "examples": [
                    "input",
                    "output",
                ],
                "name": "gen_ai.token.type",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Input tokens (prompt, input, etc.)",
                            "id": "input",
                            "stability": "development",
                            "value": "input",
                        },
                        {
                            "brief": "Output tokens (completion, response, etc.)",
                            "deprecated": {
                                "note": "Replaced by `output`.",
                                "reason": "unspecified",
                            },
                            "id": "completion",
                            "stability": "development",
                            "value": "output",
                        },
                        {
                            "brief": "Output tokens (completion, response, etc.)",
                            "id": "output",
                            "stability": "development",
                            "value": "output",
                        },
                    ],
                },
            },
            {
                "brief": "The name of the GenAI model a request is being made to.",
                "examples": "gpt-4",
                "name": "gen_ai.request.model",
                "requirement_level": {
                    "conditionally_required": "If available.",
                },
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "GenAI server port.",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": {
                    "conditionally_required": "If `server.address` is set.",
                },
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "The name of the model that generated the response.",
                "examples": [
                    "gpt-4-0613",
                ],
                "name": "gen_ai.response.model",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "GenAI server address.",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "recommended",
                "stability": "stable",
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
                    "brief": "GenAI server address.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "GenAI server port.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": {
                        "conditionally_required": "If `server.address` is set.",
                    },
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "The name of the model that generated the response.",
                    "examples": [
                        "gpt-4-0613",
                    ],
                    "name": "gen_ai.response.model",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the GenAI model a request is being made to.",
                    "examples": "gpt-4",
                    "name": "gen_ai.request.model",
                    "requirement_level": {
                        "conditionally_required": "If available.",
                    },
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The Generative AI provider as identified by the client or server instrumentation.",
                    "name": "gen_ai.provider.name",
                    "note": "The attribute SHOULD be set based on the instrumentation's best\nknowledge and may differ from the actual model provider.\n\nMultiple providers, including Azure OpenAI, Gemini, and AI hosting platforms\nare accessible using the OpenAI REST API and corresponding client libraries,\nbut may proxy or host models from different providers.\n\nThe `gen_ai.request.model`, `gen_ai.response.model`, and `server.address`\nattributes may help identify the actual system in use.\n\nThe `gen_ai.provider.name` attribute acts as a discriminator that\nidentifies the GenAI telemetry format flavor specific to that provider\nwithin GenAI semantic conventions.\nIt SHOULD be set consistently with provider-specific attributes and signals.\nFor example, GenAI spans, metrics, and events related to AWS Bedrock\nshould have the `gen_ai.provider.name` set to `aws.bedrock` and include\napplicable `aws.bedrock.*` attributes and are not expected to include\n`openai.*` attributes.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "[OpenAI](https://openai.com/)",
                                "id": "openai",
                                "stability": "development",
                                "value": "openai",
                            },
                            {
                                "brief": "Any Google generative AI endpoint",
                                "id": "gcp.gen_ai",
                                "note": "May be used when specific backend is unknown.\n",
                                "stability": "development",
                                "value": "gcp.gen_ai",
                            },
                            {
                                "brief": "[Vertex AI](https://cloud.google.com/vertex-ai)",
                                "id": "gcp.vertex_ai",
                                "note": "Used when accessing the 'aiplatform.googleapis.com' endpoint.\n",
                                "stability": "development",
                                "value": "gcp.vertex_ai",
                            },
                            {
                                "brief": "[Gemini](https://cloud.google.com/products/gemini)",
                                "id": "gcp.gemini",
                                "note": "Used when accessing the 'generativelanguage.googleapis.com' endpoint. Also known as the AI Studio API.\n",
                                "stability": "development",
                                "value": "gcp.gemini",
                            },
                            {
                                "brief": "[Anthropic](https://www.anthropic.com/)",
                                "id": "anthropic",
                                "stability": "development",
                                "value": "anthropic",
                            },
                            {
                                "brief": "[Cohere](https://cohere.com/)",
                                "id": "cohere",
                                "stability": "development",
                                "value": "cohere",
                            },
                            {
                                "brief": "Azure AI Inference",
                                "id": "azure.ai.inference",
                                "stability": "development",
                                "value": "azure.ai.inference",
                            },
                            {
                                "brief": "[Azure OpenAI](https://azure.microsoft.com/products/ai-services/openai-service/)",
                                "id": "azure.ai.openai",
                                "stability": "development",
                                "value": "azure.ai.openai",
                            },
                            {
                                "brief": "[IBM Watsonx AI](https://www.ibm.com/products/watsonx-ai)",
                                "id": "ibm.watsonx.ai",
                                "stability": "development",
                                "value": "ibm.watsonx.ai",
                            },
                            {
                                "brief": "[AWS Bedrock](https://aws.amazon.com/bedrock)",
                                "id": "aws.bedrock",
                                "stability": "development",
                                "value": "aws.bedrock",
                            },
                            {
                                "brief": "[Perplexity](https://www.perplexity.ai/)",
                                "id": "perplexity",
                                "stability": "development",
                                "value": "perplexity",
                            },
                            {
                                "brief": "[xAI](https://x.ai/)",
                                "id": "x_ai",
                                "stability": "development",
                                "value": "x_ai",
                            },
                            {
                                "brief": "[DeepSeek](https://www.deepseek.com/)",
                                "id": "deepseek",
                                "stability": "development",
                                "value": "deepseek",
                            },
                            {
                                "brief": "[Groq](https://groq.com/)",
                                "id": "groq",
                                "stability": "development",
                                "value": "groq",
                            },
                            {
                                "brief": "[Mistral AI](https://mistral.ai/)",
                                "id": "mistral_ai",
                                "stability": "development",
                                "value": "mistral_ai",
                            },
                        ],
                    },
                },
                {
                    "brief": "The name of the operation being performed.",
                    "name": "gen_ai.operation.name",
                    "note": "If one of the predefined values applies, but specific system uses a different name it's RECOMMENDED to document it in the semantic conventions for specific GenAI system and use system-specific name in the instrumentation. If a different name is not documented, instrumentation libraries SHOULD use applicable predefined value.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Chat completion operation such as [OpenAI Chat API](https://platform.openai.com/docs/api-reference/chat)",
                                "id": "chat",
                                "stability": "development",
                                "value": "chat",
                            },
                            {
                                "brief": "Multimodal content generation operation such as [Gemini Generate Content](https://ai.google.dev/api/generate-content)",
                                "id": "generate_content",
                                "stability": "development",
                                "value": "generate_content",
                            },
                            {
                                "brief": "Text completions operation such as [OpenAI Completions API (Legacy)](https://platform.openai.com/docs/api-reference/completions)",
                                "id": "text_completion",
                                "stability": "development",
                                "value": "text_completion",
                            },
                            {
                                "brief": "Embeddings operation such as [OpenAI Create embeddings API](https://platform.openai.com/docs/api-reference/embeddings/create)",
                                "id": "embeddings",
                                "stability": "development",
                                "value": "embeddings",
                            },
                            {
                                "brief": "Create GenAI agent",
                                "id": "create_agent",
                                "stability": "development",
                                "value": "create_agent",
                            },
                            {
                                "brief": "Invoke GenAI agent",
                                "id": "invoke_agent",
                                "stability": "development",
                                "value": "invoke_agent",
                            },
                            {
                                "brief": "Execute a tool",
                                "id": "execute_tool",
                                "stability": "development",
                                "value": "execute_tool",
                            },
                        ],
                    },
                },
                {
                    "brief": "The type of token being counted.",
                    "examples": [
                        "input",
                        "output",
                    ],
                    "name": "gen_ai.token.type",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Input tokens (prompt, input, etc.)",
                                "id": "input",
                                "stability": "development",
                                "value": "input",
                            },
                            {
                                "brief": "Output tokens (completion, response, etc.)",
                                "deprecated": {
                                    "note": "Replaced by `output`.",
                                    "reason": "unspecified",
                                },
                                "id": "completion",
                                "stability": "development",
                                "value": "output",
                            },
                            {
                                "brief": "Output tokens (completion, response, etc.)",
                                "id": "output",
                                "stability": "development",
                                "value": "output",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of input and output tokens used.",
            "id": "metric.gen_ai.client.token.usage",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "gen_ai.operation.name": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "gen_ai.provider.name": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "gen_ai.request.model": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "gen_ai.response.model": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "gen_ai.token.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "server.address": {
                        "inherited_fields": [
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/gen-ai/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "gen_ai.client.token.usage",
            "root_namespace": "gen_ai",
            "stability": "development",
            "type": "metric",
            "unit": "{token}",
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