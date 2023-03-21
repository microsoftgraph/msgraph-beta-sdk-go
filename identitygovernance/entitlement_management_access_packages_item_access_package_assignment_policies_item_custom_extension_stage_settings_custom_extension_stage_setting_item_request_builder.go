package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder provides operations to manage the customExtensionStageSettings property of the microsoft.graph.accessPackageAssignmentPolicy entity.
type EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderGetQueryParameters the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
type EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderGetQueryParameters
}
// NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderInternal instantiates a new CustomExtensionStageSettingItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionStageSettings/{customExtensionStageSetting%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder instantiates a new CustomExtensionStageSettingItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
func (m *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionStageSettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomExtensionStageSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionStageSettingable), nil
}
// ToGetRequestInformation the collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
func (m *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionStageSettingsCustomExtensionStageSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
