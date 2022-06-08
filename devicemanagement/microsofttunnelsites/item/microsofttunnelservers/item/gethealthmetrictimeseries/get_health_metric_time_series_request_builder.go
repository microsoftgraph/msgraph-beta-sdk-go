package gethealthmetrictimeseries

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetHealthMetricTimeSeriesRequestBuilder provides operations to call the getHealthMetricTimeSeries method.
type GetHealthMetricTimeSeriesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetHealthMetricTimeSeriesRequestBuilderInternal instantiates a new GetHealthMetricTimeSeriesRequestBuilder and sets the default values.
func NewGetHealthMetricTimeSeriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetHealthMetricTimeSeriesRequestBuilder) {
    m := &GetHealthMetricTimeSeriesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers/{microsoftTunnelServer%2Did}/microsoft.graph.getHealthMetricTimeSeries";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetHealthMetricTimeSeriesRequestBuilder instantiates a new GetHealthMetricTimeSeriesRequestBuilder and sets the default values.
func NewGetHealthMetricTimeSeriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetHealthMetricTimeSeriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetHealthMetricTimeSeriesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getHealthMetricTimeSeries
func (m *GetHealthMetricTimeSeriesRequestBuilder) CreatePostRequestInformation(body GetHealthMetricTimeSeriesPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getHealthMetricTimeSeries
func (m *GetHealthMetricTimeSeriesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *GetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getHealthMetricTimeSeries
func (m *GetHealthMetricTimeSeriesRequestBuilder) Post(body GetHealthMetricTimeSeriesPostRequestBodyable)(GetHealthMetricTimeSeriesResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getHealthMetricTimeSeries
func (m *GetHealthMetricTimeSeriesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *GetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetHealthMetricTimeSeriesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetHealthMetricTimeSeriesResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetHealthMetricTimeSeriesResponseable), nil
}
