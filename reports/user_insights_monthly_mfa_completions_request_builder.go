package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyMfaCompletionsRequestBuilder provides operations to manage the mfaCompletions property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyMfaCompletionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyMfaCompletionsRequestBuilderGetQueryParameters get a list of monthly MFA completions on apps registered in your tenant configured for Microsoft Entra External ID for customers.
type UserInsightsMonthlyMfaCompletionsRequestBuilderGetQueryParameters struct {
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
// UserInsightsMonthlyMfaCompletionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyMfaCompletionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyMfaCompletionsRequestBuilderGetQueryParameters
}
// ByMfaCompletionMetricId provides operations to manage the mfaCompletions property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserInsightsMonthlyMfaCompletionsMfaCompletionMetricItemRequestBuilder when successful
func (m *UserInsightsMonthlyMfaCompletionsRequestBuilder) ByMfaCompletionMetricId(mfaCompletionMetricId string)(*UserInsightsMonthlyMfaCompletionsMfaCompletionMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mfaCompletionMetricId != "" {
        urlTplParams["mfaCompletionMetric%2Did"] = mfaCompletionMetricId
    }
    return NewUserInsightsMonthlyMfaCompletionsMfaCompletionMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserInsightsMonthlyMfaCompletionsRequestBuilderInternal instantiates a new UserInsightsMonthlyMfaCompletionsRequestBuilder and sets the default values.
func NewUserInsightsMonthlyMfaCompletionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyMfaCompletionsRequestBuilder) {
    m := &UserInsightsMonthlyMfaCompletionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/mfaCompletions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyMfaCompletionsRequestBuilder instantiates a new UserInsightsMonthlyMfaCompletionsRequestBuilder and sets the default values.
func NewUserInsightsMonthlyMfaCompletionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyMfaCompletionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyMfaCompletionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserInsightsMonthlyMfaCompletionsCountRequestBuilder when successful
func (m *UserInsightsMonthlyMfaCompletionsRequestBuilder) Count()(*UserInsightsMonthlyMfaCompletionsCountRequestBuilder) {
    return NewUserInsightsMonthlyMfaCompletionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of monthly MFA completions on apps registered in your tenant configured for Microsoft Entra External ID for customers.
// returns a MfaCompletionMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/monthlyuserinsightmetricsroot-list-mfacompletions?view=graph-rest-beta
func (m *UserInsightsMonthlyMfaCompletionsRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyMfaCompletionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaCompletionMetricCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMfaCompletionMetricCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaCompletionMetricCollectionResponseable), nil
}
// ToGetRequestInformation get a list of monthly MFA completions on apps registered in your tenant configured for Microsoft Entra External ID for customers.
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyMfaCompletionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyMfaCompletionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserInsightsMonthlyMfaCompletionsRequestBuilder when successful
func (m *UserInsightsMonthlyMfaCompletionsRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyMfaCompletionsRequestBuilder) {
    return NewUserInsightsMonthlyMfaCompletionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
