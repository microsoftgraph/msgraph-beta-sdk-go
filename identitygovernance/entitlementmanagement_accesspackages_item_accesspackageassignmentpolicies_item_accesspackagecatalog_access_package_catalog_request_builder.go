package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder provides operations to manage the accessPackageCatalog property of the microsoft.graph.accessPackageAssignmentPolicy entity.
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderGetQueryParameters get accessPackageCatalog from identityGovernance
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy%2Did}/accessPackageCatalog{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get accessPackageCatalog from identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageCatalogable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable), nil
}
// ToGetRequestInformation get accessPackageCatalog from identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackageassignmentpoliciesItemAccesspackagecatalogAccessPackageCatalogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
