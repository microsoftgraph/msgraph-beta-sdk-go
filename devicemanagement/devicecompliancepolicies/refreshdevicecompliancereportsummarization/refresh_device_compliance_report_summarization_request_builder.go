package refreshdevicecompliancereportsummarization

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RefreshDeviceComplianceReportSummarizationRequestBuilder provides operations to call the refreshDeviceComplianceReportSummarization method.
type RefreshDeviceComplianceReportSummarizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RefreshDeviceComplianceReportSummarizationRequestBuilderPostOptions options for Post
type RefreshDeviceComplianceReportSummarizationRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
// CreatePostRequestInformation invoke action refreshDeviceComplianceReportSummarization
func (m *RefreshDeviceComplianceReportSummarizationRequestBuilder) CreatePostRequestInformation(options *RefreshDeviceComplianceReportSummarizationRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Post invoke action refreshDeviceComplianceReportSummarization
func (m *RefreshDeviceComplianceReportSummarizationRequestBuilder) Post(options *RefreshDeviceComplianceReportSummarizationRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
