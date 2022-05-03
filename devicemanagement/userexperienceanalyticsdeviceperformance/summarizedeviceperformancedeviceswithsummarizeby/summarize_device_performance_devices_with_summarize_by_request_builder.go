package summarizedeviceperformancedeviceswithsummarizeby

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder provides operations to call the summarizeDevicePerformanceDevices method.
type SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal instantiates a new SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, summarizeBy *string)(*SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    m := &SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsDevicePerformance/microsoft.graph.summarizeDevicePerformanceDevices(summarizeBy='{summarizeBy}')";
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
// NewSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder instantiates a new SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function summarizeDevicePerformanceDevices
func (m *SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function summarizeDevicePerformanceDevices
func (m *SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// GetWithResponseHandler invoke function summarizeDevicePerformanceDevices
func (m *SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) GetWithResponseHandler(requestConfiguration *SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration)(SummarizeDevicePerformanceDevicesWithSummarizeByResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function summarizeDevicePerformanceDevices
func (m *SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) GetWithResponseHandler(requestConfiguration *SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(SummarizeDevicePerformanceDevicesWithSummarizeByResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSummarizeDevicePerformanceDevicesWithSummarizeByResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(SummarizeDevicePerformanceDevicesWithSummarizeByResponseable), nil
}
