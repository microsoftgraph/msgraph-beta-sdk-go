package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder provides operations to call the getQuietTimePolicyUserSummaryReport method.
type ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetQuietTimePolicyUserSummaryReportRequestBuilderInternal instantiates a new GetQuietTimePolicyUserSummaryReportRequestBuilder and sets the default values.
func NewReportsGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) {
    m := &ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getQuietTimePolicyUserSummaryReport", pathParameters),
    }
    return m
}
// NewReportsGetQuietTimePolicyUserSummaryReportRequestBuilder instantiates a new GetQuietTimePolicyUserSummaryReportRequestBuilder and sets the default values.
func NewReportsGetQuietTimePolicyUserSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getQuietTimePolicyUserSummaryReport
func (m *ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) Post(ctx context.Context, body ReportsGetQuietTimePolicyUserSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation invoke action getQuietTimePolicyUserSummaryReport
func (m *ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetQuietTimePolicyUserSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) {
    return NewReportsGetQuietTimePolicyUserSummaryReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
