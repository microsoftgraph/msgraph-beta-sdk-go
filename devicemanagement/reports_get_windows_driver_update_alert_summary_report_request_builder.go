package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder provides operations to call the getWindowsDriverUpdateAlertSummaryReport method.
type ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilderInternal instantiates a new GetWindowsDriverUpdateAlertSummaryReportRequestBuilder and sets the default values.
func NewReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) {
    m := &ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getWindowsDriverUpdateAlertSummaryReport", pathParameters),
    }
    return m
}
// NewReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder instantiates a new GetWindowsDriverUpdateAlertSummaryReportRequestBuilder and sets the default values.
func NewReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getWindowsDriverUpdateAlertSummaryReport
func (m *ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) Post(ctx context.Context, body ReportsGetWindowsDriverUpdateAlertSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getWindowsDriverUpdateAlertSummaryReport
func (m *ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetWindowsDriverUpdateAlertSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
