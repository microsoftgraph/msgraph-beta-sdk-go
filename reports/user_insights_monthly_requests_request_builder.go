package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyRequestsRequestBuilder provides operations to manage the requests property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyRequestsRequestBuilderGetQueryParameters get a list of monthly user requests on apps registered in your tenant configured for Microsoft Entra External ID for customers.
type UserInsightsMonthlyRequestsRequestBuilderGetQueryParameters struct {
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
// UserInsightsMonthlyRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyRequestsRequestBuilderGetQueryParameters
}
// ByUserRequestsMetricId provides operations to manage the requests property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder when successful
func (m *UserInsightsMonthlyRequestsRequestBuilder) ByUserRequestsMetricId(userRequestsMetricId string)(*UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userRequestsMetricId != "" {
        urlTplParams["userRequestsMetric%2Did"] = userRequestsMetricId
    }
    return NewUserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserInsightsMonthlyRequestsRequestBuilderInternal instantiates a new UserInsightsMonthlyRequestsRequestBuilder and sets the default values.
func NewUserInsightsMonthlyRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyRequestsRequestBuilder) {
    m := &UserInsightsMonthlyRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/requests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyRequestsRequestBuilder instantiates a new UserInsightsMonthlyRequestsRequestBuilder and sets the default values.
func NewUserInsightsMonthlyRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserInsightsMonthlyRequestsCountRequestBuilder when successful
func (m *UserInsightsMonthlyRequestsRequestBuilder) Count()(*UserInsightsMonthlyRequestsCountRequestBuilder) {
    return NewUserInsightsMonthlyRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of monthly user requests on apps registered in your tenant configured for Microsoft Entra External ID for customers.
// returns a UserRequestsMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/monthlyuserinsightmetricsroot-list-requests?view=graph-rest-1.0
func (m *UserInsightsMonthlyRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRequestsMetricCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserRequestsMetricCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRequestsMetricCollectionResponseable), nil
}
// ToGetRequestInformation get a list of monthly user requests on apps registered in your tenant configured for Microsoft Entra External ID for customers.
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserInsightsMonthlyRequestsRequestBuilder when successful
func (m *UserInsightsMonthlyRequestsRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyRequestsRequestBuilder) {
    return NewUserInsightsMonthlyRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
