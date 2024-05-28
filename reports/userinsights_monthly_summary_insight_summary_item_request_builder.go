package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder provides operations to manage the summary property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserinsightsMonthlySummaryInsightSummaryItemRequestBuilderGetQueryParameters summary of all usage insights on apps registered in the tenant for a specified period.
type UserinsightsMonthlySummaryInsightSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserinsightsMonthlySummaryInsightSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsMonthlySummaryInsightSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserinsightsMonthlySummaryInsightSummaryItemRequestBuilderGetQueryParameters
}
// NewUserinsightsMonthlySummaryInsightSummaryItemRequestBuilderInternal instantiates a new UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder and sets the default values.
func NewUserinsightsMonthlySummaryInsightSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder) {
    m := &UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/summary/{insightSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserinsightsMonthlySummaryInsightSummaryItemRequestBuilder instantiates a new UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder and sets the default values.
func NewUserinsightsMonthlySummaryInsightSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserinsightsMonthlySummaryInsightSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get summary of all usage insights on apps registered in the tenant for a specified period.
// returns a InsightSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserinsightsMonthlySummaryInsightSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InsightSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInsightSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InsightSummaryable), nil
}
// ToGetRequestInformation summary of all usage insights on apps registered in the tenant for a specified period.
// returns a *RequestInformation when successful
func (m *UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserinsightsMonthlySummaryInsightSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder when successful
func (m *UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder) WithUrl(rawUrl string)(*UserinsightsMonthlySummaryInsightSummaryItemRequestBuilder) {
    return NewUserinsightsMonthlySummaryInsightSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
