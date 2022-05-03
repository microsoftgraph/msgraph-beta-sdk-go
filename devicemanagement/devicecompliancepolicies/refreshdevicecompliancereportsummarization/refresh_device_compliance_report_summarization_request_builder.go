package refreshdevicecompliancereportsummarization

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RefreshDeviceComplianceReportSummarizationRequestBuilder provides operations to call the refreshDeviceComplianceReportSummarization method.
type RefreshDeviceComplianceReportSummarizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRefreshDeviceComplianceReportSummarizationRequestBuilderInternal instantiates a new RefreshDeviceComplianceReportSummarizationRequestBuilder and sets the default values.
func NewRefreshDeviceComplianceReportSummarizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RefreshDeviceComplianceReportSummarizationRequestBuilder) {
    m := &RefreshDeviceComplianceReportSummarizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/microsoft.graph.refreshDeviceComplianceReportSummarization";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRefreshDeviceComplianceReportSummarizationRequestBuilder instantiates a new RefreshDeviceComplianceReportSummarizationRequestBuilder and sets the default values.
func NewRefreshDeviceComplianceReportSummarizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RefreshDeviceComplianceReportSummarizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRefreshDeviceComplianceReportSummarizationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action refreshDeviceComplianceReportSummarization
func (m *RefreshDeviceComplianceReportSummarizationRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action refreshDeviceComplianceReportSummarization
func (m *RefreshDeviceComplianceReportSummarizationRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *RefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler invoke action refreshDeviceComplianceReportSummarization
func (m *RefreshDeviceComplianceReportSummarizationRequestBuilder) PostWithResponseHandler(requestConfiguration *RefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler invoke action refreshDeviceComplianceReportSummarization
func (m *RefreshDeviceComplianceReportSummarizationRequestBuilder) PostWithResponseHandler(requestConfiguration *RefreshDeviceComplianceReportSummarizationRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
