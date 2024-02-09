package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder provides operations to manage the requests property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderGetQueryParameters insights for all user requests on apps registered in the tenant for a specified period.
type UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderGetQueryParameters
}
// NewUserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderInternal instantiates a new UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder) {
    m := &UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/requests/{userRequestsMetric%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder instantiates a new UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get insights for all user requests on apps registered in the tenant for a specified period.
// returns a UserRequestsMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRequestsMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserRequestsMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRequestsMetricable), nil
}
// ToGetRequestInformation insights for all user requests on apps registered in the tenant for a specified period.
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder when successful
func (m *UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder) {
    return NewUserInsightsMonthlyRequestsUserRequestsMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
