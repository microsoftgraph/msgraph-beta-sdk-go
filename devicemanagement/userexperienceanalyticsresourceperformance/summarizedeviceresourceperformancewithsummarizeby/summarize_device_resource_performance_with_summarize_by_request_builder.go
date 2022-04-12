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
// SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetOptions options for Get
type SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
        urlTplParams[""] = *summarizeBy
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
func (m *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) CreateGetRequestInformation(options *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
// Get invoke function summarizeDeviceResourcePerformance
func (m *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) Get(options *SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetOptions)(SummarizeDeviceResourcePerformanceWithSummarizeByResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSummarizeDeviceResourcePerformanceWithSummarizeByResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(SummarizeDeviceResourcePerformanceWithSummarizeByResponseable), nil
}
