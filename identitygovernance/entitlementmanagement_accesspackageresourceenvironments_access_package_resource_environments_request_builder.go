package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder provides operations to manage the accessPackageResourceEnvironments property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderGetQueryParameters retrieve a list of accessPackageResourceEnvironment objects and their properties.
type EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessPackageResourceEnvironmentId provides operations to manage the accessPackageResourceEnvironments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) ByAccessPackageResourceEnvironmentId(accessPackageResourceEnvironmentId string)(*EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceEnvironmentId != "" {
        urlTplParams["accessPackageResourceEnvironment%2Did"] = accessPackageResourceEnvironmentId
    }
    return NewEntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) {
    m := &EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceEnvironments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder instantiates a new EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackageresourceenvironmentsCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) Count()(*EntitlementmanagementAccesspackageresourceenvironmentsCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourceenvironmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of accessPackageResourceEnvironment objects and their properties.
// returns a AccessPackageResourceEnvironmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/entitlementmanagement-list-accesspackageresourceenvironment?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceEnvironmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentCollectionResponseable), nil
}
// Post create new navigation property to accessPackageResourceEnvironments for identityGovernance
// returns a AccessPackageResourceEnvironmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, requestConfiguration *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable), nil
}
// ToGetRequestInformation retrieve a list of accessPackageResourceEnvironment objects and their properties.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to accessPackageResourceEnvironments for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, requestConfiguration *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
