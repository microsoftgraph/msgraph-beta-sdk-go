package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetCompliancePoliciesReportForDeviceRequestBuilder provides operations to call the getCompliancePoliciesReportForDevice method.
type ReportsGetCompliancePoliciesReportForDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetCompliancePoliciesReportForDeviceRequestBuilderInternal instantiates a new GetCompliancePoliciesReportForDeviceRequestBuilder and sets the default values.
func NewReportsGetCompliancePoliciesReportForDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetCompliancePoliciesReportForDeviceRequestBuilder) {
    m := &ReportsGetCompliancePoliciesReportForDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getCompliancePoliciesReportForDevice", pathParameters),
    }
    return m
}
// NewReportsGetCompliancePoliciesReportForDeviceRequestBuilder instantiates a new GetCompliancePoliciesReportForDeviceRequestBuilder and sets the default values.
func NewReportsGetCompliancePoliciesReportForDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetCompliancePoliciesReportForDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetCompliancePoliciesReportForDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getCompliancePoliciesReportForDevice
func (m *ReportsGetCompliancePoliciesReportForDeviceRequestBuilder) Post(ctx context.Context, body ReportsGetCompliancePoliciesReportForDevicePostRequestBodyable, requestConfiguration *ReportsGetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getCompliancePoliciesReportForDevice
func (m *ReportsGetCompliancePoliciesReportForDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetCompliancePoliciesReportForDevicePostRequestBodyable, requestConfiguration *ReportsGetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ReportsGetCompliancePoliciesReportForDeviceRequestBuilder) WithUrl(rawUrl string)(*ReportsGetCompliancePoliciesReportForDeviceRequestBuilder) {
    return NewReportsGetCompliancePoliciesReportForDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
