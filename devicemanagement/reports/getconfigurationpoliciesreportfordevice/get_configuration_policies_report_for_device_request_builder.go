package getconfigurationpoliciesreportfordevice

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetConfigurationPoliciesReportForDeviceRequestBuilder provides operations to call the getConfigurationPoliciesReportForDevice method.
type GetConfigurationPoliciesReportForDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetConfigurationPoliciesReportForDeviceRequestBuilderInternal instantiates a new GetConfigurationPoliciesReportForDeviceRequestBuilder and sets the default values.
func NewGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationPoliciesReportForDeviceRequestBuilder) {
    m := &GetConfigurationPoliciesReportForDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getConfigurationPoliciesReportForDevice";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetConfigurationPoliciesReportForDeviceRequestBuilder instantiates a new GetConfigurationPoliciesReportForDeviceRequestBuilder and sets the default values.
func NewGetConfigurationPoliciesReportForDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationPoliciesReportForDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigurationPoliciesReportForDevice
func (m *GetConfigurationPoliciesReportForDeviceRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigurationPoliciesReportForDeviceRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigurationPoliciesReportForDevice
func (m *GetConfigurationPoliciesReportForDeviceRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigurationPoliciesReportForDeviceRequestBodyable, requestConfiguration *GetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action getConfigurationPoliciesReportForDevice
func (m *GetConfigurationPoliciesReportForDeviceRequestBuilder) PostWithResponseHandler(body GetConfigurationPoliciesReportForDeviceRequestBodyable, requestConfiguration *GetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration)(GetConfigurationPoliciesReportForDeviceResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getConfigurationPoliciesReportForDevice
func (m *GetConfigurationPoliciesReportForDeviceRequestBuilder) PostWithResponseHandler(body GetConfigurationPoliciesReportForDeviceRequestBodyable, requestConfiguration *GetConfigurationPoliciesReportForDeviceRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetConfigurationPoliciesReportForDeviceResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetConfigurationPoliciesReportForDeviceResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetConfigurationPoliciesReportForDeviceResponseable), nil
}
