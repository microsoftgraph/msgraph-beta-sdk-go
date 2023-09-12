package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder provides operations to call the getDeviceStatusSummaryByCompliancePolicySettingsReport method.
type ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilderInternal instantiates a new GetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder and sets the default values.
func NewReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder) {
    m := &ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getDeviceStatusSummaryByCompliancePolicySettingsReport", pathParameters),
    }
    return m
}
// NewReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder instantiates a new GetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder and sets the default values.
func NewReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getDeviceStatusSummaryByCompliancePolicySettingsReport
func (m *ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder) Post(ctx context.Context, body ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportPostRequestBodyable, requestConfiguration *ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getDeviceStatusSummaryByCompliancePolicySettingsReport
func (m *ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportPostRequestBodyable, requestConfiguration *ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder) {
    return NewReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
