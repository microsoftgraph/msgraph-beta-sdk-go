package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder provides operations to manage the customExtension property of the microsoft.graph.customExtensionHandler entity.
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters indicates which custom workflow extension is executed at this stage. Nullable. Supports $expand.
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}/customExtensionHandlers/{customExtensionHandler%2Did}/customExtension{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get indicates which custom workflow extension is executed at this stage. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a CustomAccessPackageWorkflowExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomAccessPackageWorkflowExtensionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation indicates which custom workflow extension is executed at this stage. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemCustomextensionhandlersItemCustomextensionCustomExtensionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
