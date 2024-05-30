package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder provides operations to call the getCompliancePolicyDeviceSummaryReport method.
type ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilderInternal instantiates a new ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder and sets the default values.
func NewReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder) {
    m := &ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getCompliancePolicyDeviceSummaryReport", pathParameters),
    }
    return m
}
// NewReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder instantiates a new ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder and sets the default values.
func NewReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getCompliancePolicyDeviceSummaryReport
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder) Post(ctx context.Context, body ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getCompliancePolicyDeviceSummaryReport
// returns a *RequestInformation when successful
func (m *ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder when successful
func (m *ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder) {
    return NewReportsGetcompliancepolicydevicesummaryreportGetCompliancePolicyDeviceSummaryReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
