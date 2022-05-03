package getwindowsqualityupdatealertsummaryreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetWindowsQualityUpdateAlertSummaryReportRequestBuilder provides operations to call the getWindowsQualityUpdateAlertSummaryReport method.
type GetWindowsQualityUpdateAlertSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetWindowsQualityUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetWindowsQualityUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal instantiates a new GetWindowsQualityUpdateAlertSummaryReportRequestBuilder and sets the default values.
func NewGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetWindowsQualityUpdateAlertSummaryReportRequestBuilder) {
    m := &GetWindowsQualityUpdateAlertSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getWindowsQualityUpdateAlertSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetWindowsQualityUpdateAlertSummaryReportRequestBuilder instantiates a new GetWindowsQualityUpdateAlertSummaryReportRequestBuilder and sets the default values.
func NewGetWindowsQualityUpdateAlertSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetWindowsQualityUpdateAlertSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getWindowsQualityUpdateAlertSummaryReport
func (m *GetWindowsQualityUpdateAlertSummaryReportRequestBuilder) CreatePostRequestInformation(body GetWindowsQualityUpdateAlertSummaryReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getWindowsQualityUpdateAlertSummaryReport
func (m *GetWindowsQualityUpdateAlertSummaryReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetWindowsQualityUpdateAlertSummaryReportRequestBodyable, requestConfiguration *GetWindowsQualityUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getWindowsQualityUpdateAlertSummaryReport
func (m *GetWindowsQualityUpdateAlertSummaryReportRequestBuilder) Post(body GetWindowsQualityUpdateAlertSummaryReportRequestBodyable)(GetWindowsQualityUpdateAlertSummaryReportResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getWindowsQualityUpdateAlertSummaryReport
func (m *GetWindowsQualityUpdateAlertSummaryReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetWindowsQualityUpdateAlertSummaryReportRequestBodyable, requestConfiguration *GetWindowsQualityUpdateAlertSummaryReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetWindowsQualityUpdateAlertSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetWindowsQualityUpdateAlertSummaryReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetWindowsQualityUpdateAlertSummaryReportResponseable), nil
}
