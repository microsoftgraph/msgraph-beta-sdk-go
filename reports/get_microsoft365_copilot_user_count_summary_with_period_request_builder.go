package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder provides operations to call the getMicrosoft365CopilotUserCountSummary method.
type GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilderInternal instantiates a new GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder and sets the default values.
func NewGetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder) {
    m := &GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getMicrosoft365CopilotUserCountSummary(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder instantiates a new GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder and sets the default values.
func NewGetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the aggregated number of active and enabled users of Microsoft 365 Copilot for a specified time period.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getmicrosoft365copilotusercountsummary?view=graph-rest-beta
func (m *GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get the aggregated number of active and enabled users of Microsoft 365 Copilot for a specified time period.
// returns a *RequestInformation when successful
func (m *GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder when successful
func (m *GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder) {
    return NewGetMicrosoft365CopilotUserCountSummaryWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
