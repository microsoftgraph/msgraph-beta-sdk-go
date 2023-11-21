package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder provides operations to manage the activeUsersBreakdown property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
type UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetQueryParameters get activeUsersBreakdown from reports
type UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetQueryParameters
}
// NewUserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderInternal instantiates a new ActiveUsersBreakdownMetricItemRequestBuilder and sets the default values.
func NewUserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) {
    m := &UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/daily/activeUsersBreakdown/{activeUsersBreakdownMetric%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder instantiates a new ActiveUsersBreakdownMetricItemRequestBuilder and sets the default values.
func NewUserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get activeUsersBreakdown from reports
func (m *UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActiveUsersBreakdownMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActiveUsersBreakdownMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActiveUsersBreakdownMetricable), nil
}
// ToGetRequestInformation get activeUsersBreakdown from reports
func (m *UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) {
    return NewUserInsightsDailyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
