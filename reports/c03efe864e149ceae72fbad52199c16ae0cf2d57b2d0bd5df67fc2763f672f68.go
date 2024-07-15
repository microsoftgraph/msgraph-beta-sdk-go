package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder provides operations to manage the inactiveUsersByApplication property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderGetQueryParameters get inactiveUsersByApplication from reports
type UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderGetQueryParameters
}
// NewUserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderInternal instantiates a new UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder) {
    m := &UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/inactiveUsersByApplication/{monthlyInactiveUsersByApplicationMetric%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder instantiates a new UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get inactiveUsersByApplication from reports
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a MonthlyInactiveUsersByApplicationMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyInactiveUsersByApplicationMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMonthlyInactiveUsersByApplicationMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyInactiveUsersByApplicationMetricable), nil
}
// ToGetRequestInformation get inactiveUsersByApplication from reports
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder when successful
func (m *UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder) {
    return NewUserInsightsMonthlyInactiveUsersByApplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
