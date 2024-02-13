package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalSignInActivitiesRequestBuilder provides operations to manage the servicePrincipalSignInActivities property of the microsoft.graph.reportRoot entity.
type ServicePrincipalSignInActivitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicePrincipalSignInActivitiesRequestBuilderGetQueryParameters get a list of servicePrincipalSignInActivity objects that contains sign-in activity information for service principals in a Microsoft Entra tenant. You can use a service principal as a client or resource. A service principal supports delegated or app-only authentication context.
type ServicePrincipalSignInActivitiesRequestBuilderGetQueryParameters struct {
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
// ServicePrincipalSignInActivitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalSignInActivitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalSignInActivitiesRequestBuilderGetQueryParameters
}
// ServicePrincipalSignInActivitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalSignInActivitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByServicePrincipalSignInActivityId provides operations to manage the servicePrincipalSignInActivities property of the microsoft.graph.reportRoot entity.
// returns a *ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder when successful
func (m *ServicePrincipalSignInActivitiesRequestBuilder) ByServicePrincipalSignInActivityId(servicePrincipalSignInActivityId string)(*ServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if servicePrincipalSignInActivityId != "" {
        urlTplParams["servicePrincipalSignInActivity%2Did"] = servicePrincipalSignInActivityId
    }
    return NewServicePrincipalSignInActivitiesServicePrincipalSignInActivityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServicePrincipalSignInActivitiesRequestBuilderInternal instantiates a new ServicePrincipalSignInActivitiesRequestBuilder and sets the default values.
func NewServicePrincipalSignInActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalSignInActivitiesRequestBuilder) {
    m := &ServicePrincipalSignInActivitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/servicePrincipalSignInActivities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewServicePrincipalSignInActivitiesRequestBuilder instantiates a new ServicePrincipalSignInActivitiesRequestBuilder and sets the default values.
func NewServicePrincipalSignInActivitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalSignInActivitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalSignInActivitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ServicePrincipalSignInActivitiesCountRequestBuilder when successful
func (m *ServicePrincipalSignInActivitiesRequestBuilder) Count()(*ServicePrincipalSignInActivitiesCountRequestBuilder) {
    return NewServicePrincipalSignInActivitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of servicePrincipalSignInActivity objects that contains sign-in activity information for service principals in a Microsoft Entra tenant. You can use a service principal as a client or resource. A service principal supports delegated or app-only authentication context.
// returns a ServicePrincipalSignInActivityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-list-serviceprincipalsigninactivities?view=graph-rest-1.0
func (m *ServicePrincipalSignInActivitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalSignInActivitiesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalSignInActivityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityCollectionResponseable), nil
}
// Post create new navigation property to servicePrincipalSignInActivities for reports
// returns a ServicePrincipalSignInActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalSignInActivitiesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable, requestConfiguration *ServicePrincipalSignInActivitiesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalSignInActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable), nil
}
// ToGetRequestInformation get a list of servicePrincipalSignInActivity objects that contains sign-in activity information for service principals in a Microsoft Entra tenant. You can use a service principal as a client or resource. A service principal supports delegated or app-only authentication context.
// returns a *RequestInformation when successful
func (m *ServicePrincipalSignInActivitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalSignInActivitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to servicePrincipalSignInActivities for reports
// returns a *RequestInformation when successful
func (m *ServicePrincipalSignInActivitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalSignInActivityable, requestConfiguration *ServicePrincipalSignInActivitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/reports/servicePrincipalSignInActivities", m.BaseRequestBuilder.PathParameters)
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
// returns a *ServicePrincipalSignInActivitiesRequestBuilder when successful
func (m *ServicePrincipalSignInActivitiesRequestBuilder) WithUrl(rawUrl string)(*ServicePrincipalSignInActivitiesRequestBuilder) {
    return NewServicePrincipalSignInActivitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
