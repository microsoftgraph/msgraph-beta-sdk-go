package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder provides operations to manage the authentications property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderGetQueryParameters get authentications from reports
type UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderGetQueryParameters
}
// NewUserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderInternal instantiates a new AuthenticationsMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder) {
    m := &UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/authentications/{authenticationsMetric%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder instantiates a new AuthenticationsMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get authentications from reports
func (m *UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationsMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationsMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationsMetricable), nil
}
// ToGetRequestInformation get authentications from reports
func (m *UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder) {
    return NewUserInsightsMonthlyAuthenticationsAuthenticationsMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
