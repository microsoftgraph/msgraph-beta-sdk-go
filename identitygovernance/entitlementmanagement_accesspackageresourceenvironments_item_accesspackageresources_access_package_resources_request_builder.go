package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageResourceEnvironment entity.
type EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderGetQueryParameters read-only. Required.
type EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderGetQueryParameters
}
// ByAccessPackageResourceId provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageResourceEnvironment entity.
// returns a *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder) ByAccessPackageResourceId(accessPackageResourceId string)(*EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceId != "" {
        urlTplParams["accessPackageResource%2Did"] = accessPackageResourceId
    }
    return NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder) {
    m := &EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/{accessPackageResourceEnvironment%2Did}/accessPackageResources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder instantiates a new EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder) Count()(*EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Required.
// returns a AccessPackageResourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourceenvironmentsItemAccesspackageresourcesAccessPackageResourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
