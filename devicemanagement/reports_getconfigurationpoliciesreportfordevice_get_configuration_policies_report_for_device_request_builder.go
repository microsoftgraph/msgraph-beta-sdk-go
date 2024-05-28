package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder provides operations to call the getConfigurationPoliciesReportForDevice method.
type ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilderInternal instantiates a new ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder and sets the default values.
func NewReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder) {
    m := &ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getConfigurationPoliciesReportForDevice", pathParameters),
    }
    return m
}
// NewReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder instantiates a new ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder and sets the default values.
func NewReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getConfigurationPoliciesReportForDevice
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder) Post(ctx context.Context, body ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDevicePostRequestBodyable, requestConfiguration *ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getConfigurationPoliciesReportForDevice
// returns a *RequestInformation when successful
func (m *ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDevicePostRequestBodyable, requestConfiguration *ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder when successful
func (m *ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder) WithUrl(rawUrl string)(*ReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder) {
    return NewReportsGetconfigurationpoliciesreportfordeviceGetConfigurationPoliciesReportForDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
