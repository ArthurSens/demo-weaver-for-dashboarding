package cloud




// The cloud account ID the resource is assigned to
type AttrAccountId string // cloud.account.id

func (AttrAccountId) Development() {}
func (AttrAccountId) Recommended() {}
func (AttrAccountId) Key() string { return "cloud_account_id" }
func (a AttrAccountId) Value() string { return string(a) }





// Cloud regions often have multiple, isolated locations known as zones to increase availability. Availability zone represents the zone where the resource is running.
//
// Availability zones are called "zones" on Alibaba Cloud and Google Cloud
type AttrAvailabilityZone string // cloud.availability_zone

func (AttrAvailabilityZone) Development() {}
func (AttrAvailabilityZone) Recommended() {}
func (AttrAvailabilityZone) Key() string { return "cloud_availability_zone" }
func (a AttrAvailabilityZone) Value() string { return string(a) }





// The cloud platform in use.
//
// The prefix of the service SHOULD match the one specified in `cloud.provider`
type AttrPlatform string // cloud.platform

func (AttrPlatform) Development() {}
func (AttrPlatform) Recommended() {}
func (AttrPlatform) Key() string { return "cloud_platform" }
func (a AttrPlatform) Value() string { return string(a) }

const PlatformAlibabaCloudEcs AttrPlatform = "alibaba_cloud_ecs"
const PlatformAlibabaCloudFc AttrPlatform = "alibaba_cloud_fc"
const PlatformAlibabaCloudOpenshift AttrPlatform = "alibaba_cloud_openshift"
const PlatformAwsEc2 AttrPlatform = "aws_ec2"
const PlatformAwsEcs AttrPlatform = "aws_ecs"
const PlatformAwsEks AttrPlatform = "aws_eks"
const PlatformAwsLambda AttrPlatform = "aws_lambda"
const PlatformAwsElasticBeanstalk AttrPlatform = "aws_elastic_beanstalk"
const PlatformAwsAppRunner AttrPlatform = "aws_app_runner"
const PlatformAwsOpenshift AttrPlatform = "aws_openshift"
const PlatformAzureVm AttrPlatform = "azure.vm"
const PlatformAzureContainerApps AttrPlatform = "azure.container_apps"
const PlatformAzureContainerInstances AttrPlatform = "azure.container_instances"
const PlatformAzureAks AttrPlatform = "azure.aks"
const PlatformAzureFunctions AttrPlatform = "azure.functions"
const PlatformAzureAppService AttrPlatform = "azure.app_service"
const PlatformAzureOpenshift AttrPlatform = "azure.openshift"
const PlatformGcpBareMetalSolution AttrPlatform = "gcp_bare_metal_solution"
const PlatformGcpComputeEngine AttrPlatform = "gcp_compute_engine"
const PlatformGcpCloudRun AttrPlatform = "gcp_cloud_run"
const PlatformGcpKubernetesEngine AttrPlatform = "gcp_kubernetes_engine"
const PlatformGcpCloudFunctions AttrPlatform = "gcp_cloud_functions"
const PlatformGcpAppEngine AttrPlatform = "gcp_app_engine"
const PlatformGcpOpenshift AttrPlatform = "gcp_openshift"
const PlatformIbmCloudOpenshift AttrPlatform = "ibm_cloud_openshift"
const PlatformOracleCloudCompute AttrPlatform = "oracle_cloud_compute"
const PlatformOracleCloudOke AttrPlatform = "oracle_cloud_oke"
const PlatformTencentCloudCvm AttrPlatform = "tencent_cloud_cvm"
const PlatformTencentCloudEks AttrPlatform = "tencent_cloud_eks"
const PlatformTencentCloudScf AttrPlatform = "tencent_cloud_scf"




// Name of the cloud provider
type AttrProvider string // cloud.provider

func (AttrProvider) Development() {}
func (AttrProvider) Recommended() {}
func (AttrProvider) Key() string { return "cloud_provider" }
func (a AttrProvider) Value() string { return string(a) }

const ProviderAlibabaCloud AttrProvider = "alibaba_cloud"
const ProviderAws AttrProvider = "aws"
const ProviderAzure AttrProvider = "azure"
const ProviderGcp AttrProvider = "gcp"
const ProviderHeroku AttrProvider = "heroku"
const ProviderIbmCloud AttrProvider = "ibm_cloud"
const ProviderOracleCloud AttrProvider = "oracle_cloud"
const ProviderTencentCloud AttrProvider = "tencent_cloud"




// The geographical region within a cloud provider. When associated with a resource, this attribute specifies the region where the resource operates. When calling services or APIs deployed on a cloud, this attribute identifies the region where the called destination is deployed.
//
// Refer to your provider's docs to see the available regions, for example [Alibaba Cloud regions], [AWS regions], [Azure regions], [Google Cloud regions], or [Tencent Cloud regions]
//
// [Alibaba Cloud regions]: https://www.alibabacloud.com/help/doc-detail/40654.htm
// [AWS regions]: https://aws.amazon.com/about-aws/global-infrastructure/regions_az/
// [Azure regions]: https://azure.microsoft.com/global-infrastructure/geographies/
// [Google Cloud regions]: https://cloud.google.com/about/locations
// [Tencent Cloud regions]: https://www.tencentcloud.com/document/product/213/6091
type AttrRegion string // cloud.region

func (AttrRegion) Development() {}
func (AttrRegion) Recommended() {}
func (AttrRegion) Key() string { return "cloud_region" }
func (a AttrRegion) Value() string { return string(a) }





// Cloud provider-specific native identifier of the monitored cloud resource (e.g. an [ARN] on AWS, a [fully qualified resource ID] on Azure, a [full resource name] on GCP)
//
// On some cloud providers, it may not be possible to determine the full ID at startup,
// so it may be necessary to set `cloud.resource_id` as a span attribute instead.
//
// The exact value to use for `cloud.resource_id` depends on the cloud provider.
// The following well-known definitions MUST be used if you set this attribute and they apply:
//
//   - **AWS Lambda:** The function [ARN].
//     Take care not to use the "invoked ARN" directly but replace any
//     [alias suffix]
//     with the resolved function version, as the same runtime instance may be invocable with
//     multiple different aliases.
//   - **GCP:** The [URI of the resource]
//   - **Azure:** The [Fully Qualified Resource ID] of the invoked function,
//     *not* the function app, having the form
//     `/subscriptions/<SUBSCRIPTION_GUID>/resourceGroups/<RG>/providers/Microsoft.Web/sites/<FUNCAPP>/functions/<FUNC>`.
//     This means that a span attribute MUST be used, as an Azure function app can host multiple functions that would usually share
//     a TracerProvider
//
//
// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
// [fully qualified resource ID]: https://learn.microsoft.com/rest/api/resources/resources/get-by-id
// [full resource name]: https://google.aip.dev/122#full-resource-names
// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
// [alias suffix]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html
// [URI of the resource]: https://cloud.google.com/iam/docs/full-resource-names
// [Fully Qualified Resource ID]: https://learn.microsoft.com/rest/api/resources/resources/get-by-id
type AttrResourceId string // cloud.resource_id

func (AttrResourceId) Development() {}
func (AttrResourceId) Recommended() {}
func (AttrResourceId) Key() string { return "cloud_resource_id" }
func (a AttrResourceId) Value() string { return string(a) }




/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The cloud account ID the resource is assigned to.\n",
                    "examples": [
                        "111111111111",
                        "opentelemetry",
                    ],
                    "name": "cloud.account.id",
                    "requirement_level": "recommended",
                    "root_namespace": "cloud",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Cloud regions often have multiple, isolated locations known as zones to increase availability. Availability zone represents the zone where the resource is running.\n",
                    "examples": [
                        "us-east-1c",
                    ],
                    "name": "cloud.availability_zone",
                    "note": "Availability zones are called \"zones\" on Alibaba Cloud and Google Cloud.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloud",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The cloud platform in use.\n",
                    "name": "cloud.platform",
                    "note": "The prefix of the service SHOULD match the one specified in `cloud.provider`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloud",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Alibaba Cloud Elastic Compute Service",
                                "id": "alibaba_cloud_ecs",
                                "stability": "development",
                                "value": "alibaba_cloud_ecs",
                            },
                            {
                                "brief": "Alibaba Cloud Function Compute",
                                "id": "alibaba_cloud_fc",
                                "stability": "development",
                                "value": "alibaba_cloud_fc",
                            },
                            {
                                "brief": "Red Hat OpenShift on Alibaba Cloud",
                                "id": "alibaba_cloud_openshift",
                                "stability": "development",
                                "value": "alibaba_cloud_openshift",
                            },
                            {
                                "brief": "AWS Elastic Compute Cloud",
                                "id": "aws_ec2",
                                "stability": "development",
                                "value": "aws_ec2",
                            },
                            {
                                "brief": "AWS Elastic Container Service",
                                "id": "aws_ecs",
                                "stability": "development",
                                "value": "aws_ecs",
                            },
                            {
                                "brief": "AWS Elastic Kubernetes Service",
                                "id": "aws_eks",
                                "stability": "development",
                                "value": "aws_eks",
                            },
                            {
                                "brief": "AWS Lambda",
                                "id": "aws_lambda",
                                "stability": "development",
                                "value": "aws_lambda",
                            },
                            {
                                "brief": "AWS Elastic Beanstalk",
                                "id": "aws_elastic_beanstalk",
                                "stability": "development",
                                "value": "aws_elastic_beanstalk",
                            },
                            {
                                "brief": "AWS App Runner",
                                "id": "aws_app_runner",
                                "stability": "development",
                                "value": "aws_app_runner",
                            },
                            {
                                "brief": "Red Hat OpenShift on AWS (ROSA)",
                                "id": "aws_openshift",
                                "stability": "development",
                                "value": "aws_openshift",
                            },
                            {
                                "brief": "Azure Virtual Machines",
                                "id": "azure.vm",
                                "stability": "development",
                                "value": "azure.vm",
                            },
                            {
                                "brief": "Azure Container Apps",
                                "id": "azure.container_apps",
                                "stability": "development",
                                "value": "azure.container_apps",
                            },
                            {
                                "brief": "Azure Container Instances",
                                "id": "azure.container_instances",
                                "stability": "development",
                                "value": "azure.container_instances",
                            },
                            {
                                "brief": "Azure Kubernetes Service",
                                "id": "azure.aks",
                                "stability": "development",
                                "value": "azure.aks",
                            },
                            {
                                "brief": "Azure Functions",
                                "id": "azure.functions",
                                "stability": "development",
                                "value": "azure.functions",
                            },
                            {
                                "brief": "Azure App Service",
                                "id": "azure.app_service",
                                "stability": "development",
                                "value": "azure.app_service",
                            },
                            {
                                "brief": "Azure Red Hat OpenShift",
                                "id": "azure.openshift",
                                "stability": "development",
                                "value": "azure.openshift",
                            },
                            {
                                "brief": "Google Bare Metal Solution (BMS)",
                                "id": "gcp_bare_metal_solution",
                                "stability": "development",
                                "value": "gcp_bare_metal_solution",
                            },
                            {
                                "brief": "Google Cloud Compute Engine (GCE)",
                                "id": "gcp_compute_engine",
                                "stability": "development",
                                "value": "gcp_compute_engine",
                            },
                            {
                                "brief": "Google Cloud Run",
                                "id": "gcp_cloud_run",
                                "stability": "development",
                                "value": "gcp_cloud_run",
                            },
                            {
                                "brief": "Google Cloud Kubernetes Engine (GKE)",
                                "id": "gcp_kubernetes_engine",
                                "stability": "development",
                                "value": "gcp_kubernetes_engine",
                            },
                            {
                                "brief": "Google Cloud Functions (GCF)",
                                "id": "gcp_cloud_functions",
                                "stability": "development",
                                "value": "gcp_cloud_functions",
                            },
                            {
                                "brief": "Google Cloud App Engine (GAE)",
                                "id": "gcp_app_engine",
                                "stability": "development",
                                "value": "gcp_app_engine",
                            },
                            {
                                "brief": "Red Hat OpenShift on Google Cloud",
                                "id": "gcp_openshift",
                                "stability": "development",
                                "value": "gcp_openshift",
                            },
                            {
                                "brief": "Red Hat OpenShift on IBM Cloud",
                                "id": "ibm_cloud_openshift",
                                "stability": "development",
                                "value": "ibm_cloud_openshift",
                            },
                            {
                                "brief": "Compute on Oracle Cloud Infrastructure (OCI)",
                                "id": "oracle_cloud_compute",
                                "stability": "development",
                                "value": "oracle_cloud_compute",
                            },
                            {
                                "brief": "Kubernetes Engine (OKE) on Oracle Cloud Infrastructure (OCI)",
                                "id": "oracle_cloud_oke",
                                "stability": "development",
                                "value": "oracle_cloud_oke",
                            },
                            {
                                "brief": "Tencent Cloud Cloud Virtual Machine (CVM)",
                                "id": "tencent_cloud_cvm",
                                "stability": "development",
                                "value": "tencent_cloud_cvm",
                            },
                            {
                                "brief": "Tencent Cloud Elastic Kubernetes Service (EKS)",
                                "id": "tencent_cloud_eks",
                                "stability": "development",
                                "value": "tencent_cloud_eks",
                            },
                            {
                                "brief": "Tencent Cloud Serverless Cloud Function (SCF)",
                                "id": "tencent_cloud_scf",
                                "stability": "development",
                                "value": "tencent_cloud_scf",
                            },
                        ],
                    },
                },
                {
                    "brief": "Name of the cloud provider.\n",
                    "name": "cloud.provider",
                    "requirement_level": "recommended",
                    "root_namespace": "cloud",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Alibaba Cloud",
                                "id": "alibaba_cloud",
                                "stability": "development",
                                "value": "alibaba_cloud",
                            },
                            {
                                "brief": "Amazon Web Services",
                                "id": "aws",
                                "stability": "development",
                                "value": "aws",
                            },
                            {
                                "brief": "Microsoft Azure",
                                "id": "azure",
                                "stability": "development",
                                "value": "azure",
                            },
                            {
                                "brief": "Google Cloud Platform",
                                "id": "gcp",
                                "stability": "development",
                                "value": "gcp",
                            },
                            {
                                "brief": "Heroku Platform as a Service",
                                "id": "heroku",
                                "stability": "development",
                                "value": "heroku",
                            },
                            {
                                "brief": "IBM Cloud",
                                "id": "ibm_cloud",
                                "stability": "development",
                                "value": "ibm_cloud",
                            },
                            {
                                "brief": "Oracle Cloud Infrastructure (OCI)",
                                "id": "oracle_cloud",
                                "stability": "development",
                                "value": "oracle_cloud",
                            },
                            {
                                "brief": "Tencent Cloud",
                                "id": "tencent_cloud",
                                "stability": "development",
                                "value": "tencent_cloud",
                            },
                        ],
                    },
                },
                {
                    "brief": "The geographical region within a cloud provider. When associated with a resource, this attribute specifies the region where the resource operates. When calling services or APIs deployed on a cloud, this attribute identifies the region where the called destination is deployed.\n",
                    "examples": [
                        "us-central1",
                        "us-east-1",
                    ],
                    "name": "cloud.region",
                    "note": "Refer to your provider's docs to see the available regions, for example [Alibaba Cloud regions](https://www.alibabacloud.com/help/doc-detail/40654.htm), [AWS regions](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/), [Azure regions](https://azure.microsoft.com/global-infrastructure/geographies/), [Google Cloud regions](https://cloud.google.com/about/locations), or [Tencent Cloud regions](https://www.tencentcloud.com/document/product/213/6091).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloud",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Cloud provider-specific native identifier of the monitored cloud resource (e.g. an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) on AWS, a [fully qualified resource ID](https://learn.microsoft.com/rest/api/resources/resources/get-by-id) on Azure, a [full resource name](https://google.aip.dev/122#full-resource-names) on GCP)\n",
                    "examples": [
                        "arn:aws:lambda:REGION:ACCOUNT_ID:function:my-function",
                        "//run.googleapis.com/projects/PROJECT_ID/locations/LOCATION_ID/services/SERVICE_ID",
                        "/subscriptions/<SUBSCRIPTION_GUID>/resourceGroups/<RG>/providers/Microsoft.Web/sites/<FUNCAPP>/functions/<FUNC>",
                    ],
                    "name": "cloud.resource_id",
                    "note": "On some cloud providers, it may not be possible to determine the full ID at startup,\nso it may be necessary to set `cloud.resource_id` as a span attribute instead.\n\nThe exact value to use for `cloud.resource_id` depends on the cloud provider.\nThe following well-known definitions MUST be used if you set this attribute and they apply:\n\n- **AWS Lambda:** The function [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).\n  Take care not to use the \"invoked ARN\" directly but replace any\n  [alias suffix](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html)\n  with the resolved function version, as the same runtime instance may be invocable with\n  multiple different aliases.\n- **GCP:** The [URI of the resource](https://cloud.google.com/iam/docs/full-resource-names)\n- **Azure:** The [Fully Qualified Resource ID](https://learn.microsoft.com/rest/api/resources/resources/get-by-id) of the invoked function,\n  *not* the function app, having the form\n  `/subscriptions/<SUBSCRIPTION_GUID>/resourceGroups/<RG>/providers/Microsoft.Web/sites/<FUNCAPP>/functions/<FUNC>`.\n  This means that a span attribute MUST be used, as an Azure function app can host multiple functions that would usually share\n  a TracerProvider.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloud",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "cloud",
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