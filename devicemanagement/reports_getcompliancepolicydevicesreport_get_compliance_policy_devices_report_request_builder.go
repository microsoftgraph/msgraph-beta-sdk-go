package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder provides operations to call the getCompliancePolicyDevicesReport method.
type ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilderInternal instantiates a new ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder and sets the default values.
func NewReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder) {
    m := &ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getCompliancePolicyDevicesReport", pathParameters),
    }
    return m
}
// NewReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder instantiates a new ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder and sets the default values.
func NewReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getCompliancePolicyDevicesReport
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder) Post(ctx context.Context, body ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportPostRequestBodyable, requestConfiguration *ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getCompliancePolicyDevicesReport
// returns a *RequestInformation when successful
func (m *ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportPostRequestBodyable, requestConfiguration *ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder when successful
func (m *ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder) {
    return NewReportsGetcompliancepolicydevicesreportGetCompliancePolicyDevicesReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
