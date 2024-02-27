package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder provides operations to manage the activeUsersBreakdown property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetQueryParameters get activeUsersBreakdown from reports
type UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetQueryParameters
}
// NewUserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderInternal instantiates a new UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) {
    m := &UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/activeUsersBreakdown/{activeUsersBreakdownMetric%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder instantiates a new UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get activeUsersBreakdown from reports
// Deprecated: The Active Users Breakdown Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Active Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a ActiveUsersBreakdownMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActiveUsersBreakdownMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Deprecated: The Active Users Breakdown Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Active Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Active Users Breakdown Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Active Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder when successful
func (m *UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder) {
    return NewUserInsightsMonthlyActiveUsersBreakdownActiveUsersBreakdownMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
