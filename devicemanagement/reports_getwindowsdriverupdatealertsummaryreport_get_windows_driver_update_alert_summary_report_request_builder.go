package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder provides operations to call the getWindowsDriverUpdateAlertSummaryReport method.
type ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilderInternal instantiates a new ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder and sets the default values.
func NewReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) {
    m := &ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getWindowsDriverUpdateAlertSummaryReport", pathParameters),
    }
    return m
}
// NewReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder instantiates a new ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder and sets the default values.
func NewReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getWindowsDriverUpdateAlertSummaryReport
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) Post(ctx context.Context, body ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action getWindowsDriverUpdateAlertSummaryReport
// returns a *RequestInformation when successful
func (m *ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder when successful
func (m *ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsGetwindowsdriverupdatealertsummaryreportGetWindowsDriverUpdateAlertSummaryReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
