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
// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions options for Post
type GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions struct {
    // 
    Body GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBodyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
func (m *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) CreatePostRequestInformation(options *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) Post(options *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions)(GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportResponseable), nil
}
