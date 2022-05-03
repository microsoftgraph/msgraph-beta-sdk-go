package getcompliancepoliciesreportfordevice

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetCompliancePoliciesReportForDeviceRequestBuilder provides operations to call the getCompliancePoliciesReportForDevice method.
type GetCompliancePoliciesReportForDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetCompliancePoliciesReportForDeviceRequestBuilderInternal instantiates a new GetCompliancePoliciesReportForDeviceRequestBuilder and sets the default values.
func NewGetCompliancePoliciesReportForDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetCompliancePoliciesReportForDeviceRequestBuilder) {
    m := &GetCompliancePoliciesReportForDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getCompliancePoliciesReportForDevice";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetCompliancePoliciesReportForDeviceRequestBuilder instantiates a new GetCompliancePoliciesReportForDeviceRequestBuilder and sets the default values.
func NewGetCompliancePoliciesReportForDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetCompliancePoliciesReportForDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetCompliancePoliciesReportForDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getCompliancePoliciesReportForDevice
func (m *GetCompliancePoliciesReportForDeviceRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetCompliancePoliciesReportForDeviceRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getCompliancePoliciesReportForDevice
func (m *GetCompliancePoliciesReportForDeviceRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetCompliancePoliciesReportForDeviceRequestBodyable, requestConfiguration *GetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler invoke action getCompliancePoliciesReportForDevice
func (m *GetCompliancePoliciesReportForDeviceRequestBuilder) PostWithResponseHandler(body GetCompliancePoliciesReportForDeviceRequestBodyable, requestConfiguration *GetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration)(GetCompliancePoliciesReportForDeviceResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getCompliancePoliciesReportForDevice
func (m *GetCompliancePoliciesReportForDeviceRequestBuilder) PostWithResponseHandler(body GetCompliancePoliciesReportForDeviceRequestBodyable, requestConfiguration *GetCompliancePoliciesReportForDeviceRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetCompliancePoliciesReportForDeviceResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetCompliancePoliciesReportForDeviceResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetCompliancePoliciesReportForDeviceResponseable), nil
}
