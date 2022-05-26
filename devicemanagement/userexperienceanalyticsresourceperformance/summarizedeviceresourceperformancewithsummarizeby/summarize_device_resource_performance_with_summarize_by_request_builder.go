package summarizedeviceresourceperformancewithsummarizeby

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder provides operations to call the summarizeDeviceResourcePerformance method.
type SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal instantiates a new SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, summarizeBy *string)(*SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    m := &SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsResourcePerformance/microsoft.graph.summarizeDeviceResourcePerformance(summarizeBy='{summarizeBy}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if summarizeBy != nil {
        urlTplParams["summarizeBy"] = *summarizeBy
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder instantiates a new SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function summarizeDeviceResourcePerformance
func (m *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function summarizeDeviceResourcePerformance
func (m *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function summarizeDeviceResourcePerformance
func (m *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) Get()(SummarizeDeviceResourcePerformanceWithSummarizeByResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function summarizeDeviceResourcePerformance
func (m *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(SummarizeDeviceResourcePerformanceWithSummarizeByResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSummarizeDeviceResourcePerformanceWithSummarizeByResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(SummarizeDeviceResourcePerformanceWithSummarizeByResponseable), nil
}
