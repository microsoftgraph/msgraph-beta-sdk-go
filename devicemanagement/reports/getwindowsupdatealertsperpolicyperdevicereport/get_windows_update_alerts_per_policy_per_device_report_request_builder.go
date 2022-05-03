package getwindowsupdatealertsperpolicyperdevicereport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder provides operations to call the getWindowsUpdateAlertsPerPolicyPerDeviceReport method.
type GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal instantiates a new GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    m := &GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getWindowsUpdateAlertsPerPolicyPerDeviceReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder instantiates a new GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getWindowsUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) CreatePostRequestInformation(body GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getWindowsUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBodyable, requestConfiguration *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getWindowsUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) Post(body GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBodyable)(GetWindowsUpdateAlertsPerPolicyPerDeviceReportResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getWindowsUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBodyable, requestConfiguration *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetWindowsUpdateAlertsPerPolicyPerDeviceReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetWindowsUpdateAlertsPerPolicyPerDeviceReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetWindowsUpdateAlertsPerPolicyPerDeviceReportResponseable), nil
}
