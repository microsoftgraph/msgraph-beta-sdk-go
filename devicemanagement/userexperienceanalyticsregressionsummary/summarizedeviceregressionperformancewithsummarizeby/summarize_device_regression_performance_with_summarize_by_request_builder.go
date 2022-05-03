package summarizedeviceregressionperformancewithsummarizeby

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder provides operations to call the summarizeDeviceRegressionPerformance method.
type SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderInternal instantiates a new SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, summarizeBy *string)(*SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) {
    m := &SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsRegressionSummary/microsoft.graph.summarizeDeviceRegressionPerformance(summarizeBy='{summarizeBy}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if summarizeBy != nil {
        urlTplParams[""] = *summarizeBy
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder instantiates a new SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function summarizeDeviceRegressionPerformance
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function summarizeDeviceRegressionPerformance
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler invoke function summarizeDeviceRegressionPerformance
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) GetWithResponseHandler(requestConfiguration *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function summarizeDeviceRegressionPerformance
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) GetWithResponseHandler(requestConfiguration *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsRegressionSummaryFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable), nil
}
