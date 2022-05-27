package getwindowsqualityupdatealertsperpolicyperdevicereport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder provides operations to call the getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport method.
type GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal instantiates a new GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    m := &GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder instantiates a new GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) CreatePostRequestInformation(body GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportPostRequestBodyable, requestConfiguration *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) Post(body GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportPostRequestBodyable)([]byte, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportPostRequestBodyable, requestConfiguration *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(requestInfo, "byte", responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
