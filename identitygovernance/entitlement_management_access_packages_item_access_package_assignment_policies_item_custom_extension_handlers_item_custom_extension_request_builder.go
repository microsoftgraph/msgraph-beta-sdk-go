package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder provides operations to manage the customExtension property of the microsoft.graph.customExtensionHandler entity.
type EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderGetQueryParameters indicates which custom workflow extension will be executed at this stage. Nullable. Supports $expand.
type EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderGetQueryParameters
}
// NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderInternal instantiates a new CustomExtensionRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionHandlers/{customExtensionHandler%2Did}/customExtension{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder instantiates a new CustomExtensionRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get indicates which custom workflow extension will be executed at this stage. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomAccessPackageWorkflowExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionable), nil
}
// ToGetRequestInformation indicates which custom workflow extension will be executed at this stage. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageAssignmentPoliciesItemCustomExtensionHandlersItemCustomExtensionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
