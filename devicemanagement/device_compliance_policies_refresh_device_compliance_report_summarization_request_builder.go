package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder provides operations to call the refreshDeviceComplianceReportSummarization method.
type DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilderInternal instantiates a new RefreshDeviceComplianceReportSummarizationRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder) {
    m := &DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/refreshDeviceComplianceReportSummarization", pathParameters),
    }
    return m
}
// NewDeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder instantiates a new RefreshDeviceComplianceReportSummarizationRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action refreshDeviceComplianceReportSummarization
func (m *DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder) Post(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action refreshDeviceComplianceReportSummarization
func (m *DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesRefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
