package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder provides operations to call the getDeviceStatusByCompliancePolicySettingReport method.
type ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilderInternal instantiates a new GetDeviceStatusByCompliancePolicySettingReportRequestBuilder and sets the default values.
func NewReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder) {
    m := &ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getDeviceStatusByCompliancePolicySettingReport", pathParameters),
    }
    return m
}
// NewReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder instantiates a new GetDeviceStatusByCompliancePolicySettingReportRequestBuilder and sets the default values.
func NewReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getDeviceStatusByCompliancePolicySettingReport
func (m *ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder) Post(ctx context.Context, body ReportsGetDeviceStatusByCompliancePolicySettingReportPostRequestBodyable, requestConfiguration *ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getDeviceStatusByCompliancePolicySettingReport
func (m *ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetDeviceStatusByCompliancePolicySettingReportPostRequestBodyable, requestConfiguration *ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder) {
    return NewReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
