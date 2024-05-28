package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder provides operations to manage the mfaCompletions property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
type UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderGetQueryParameters get a list of daily MFA completions on apps registered in your tenant configured for Microsoft Entra External ID for customers.
type UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderGetQueryParameters struct {
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
// UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderGetQueryParameters
}
// ByMfaCompletionMetricId provides operations to manage the mfaCompletions property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserinsightsDailyMfacompletionsMfaCompletionMetricItemRequestBuilder when successful
func (m *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) ByMfaCompletionMetricId(mfaCompletionMetricId string)(*UserinsightsDailyMfacompletionsMfaCompletionMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mfaCompletionMetricId != "" {
        urlTplParams["mfaCompletionMetric%2Did"] = mfaCompletionMetricId
    }
    return NewUserinsightsDailyMfacompletionsMfaCompletionMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderInternal instantiates a new UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder and sets the default values.
func NewUserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) {
    m := &UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/daily/mfaCompletions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder instantiates a new UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder and sets the default values.
func NewUserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserinsightsDailyMfacompletionsCountRequestBuilder when successful
func (m *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) Count()(*UserinsightsDailyMfacompletionsCountRequestBuilder) {
    return NewUserinsightsDailyMfacompletionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of daily MFA completions on apps registered in your tenant configured for Microsoft Entra External ID for customers.
// returns a MfaCompletionMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/dailyuserinsightmetricsroot-list-mfacompletions?view=graph-rest-beta
func (m *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) Get(ctx context.Context, requestConfiguration *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaCompletionMetricCollectionResponseable, error) {
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
// ToGetRequestInformation get a list of daily MFA completions on apps registered in your tenant configured for Microsoft Entra External ID for customers.
// returns a *RequestInformation when successful
func (m *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder when successful
func (m *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) WithUrl(rawUrl string)(*UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) {
    return NewUserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
