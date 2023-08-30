package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageResourceEnvironment entity.
type EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderGetQueryParameters read-only. Required.
type EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderGetQueryParameters
}
// ByAccessPackageResourceId provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageResourceEnvironment entity.
func (m *EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) ByAccessPackageResourceId(accessPackageResourceId string)(*EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesAccessPackageResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceId != "" {
        urlTplParams["accessPackageResource%2Did"] = accessPackageResourceId
    }
    return NewEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesAccessPackageResourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderInternal instantiates a new AccessPackageResourcesRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) {
    m := &EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/{accessPackageResourceEnvironment%2Did}/accessPackageResources{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder instantiates a new AccessPackageResourcesRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) Count()(*EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesCountRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Required.
func (m *EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceCollectionResponseable), nil
}
// ToGetRequestInformation read-only. Required.
func (m *EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourceEnvironmentsItemAccessPackageResourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
