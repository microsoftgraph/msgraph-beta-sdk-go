package getteamsdeviceusagedistributiontotalusercountswithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder provides operations to call the getTeamsDeviceUsageDistributionTotalUserCounts method.
type GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal instantiates a new GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    m := &GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsDeviceUsageDistributionTotalUserCounts(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams[""] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder instantiates a new GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getTeamsDeviceUsageDistributionTotalUserCounts
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getTeamsDeviceUsageDistributionTotalUserCounts
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// GetWithResponseHandler invoke function getTeamsDeviceUsageDistributionTotalUserCounts
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) GetWithResponseHandler(requestConfiguration *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getTeamsDeviceUsageDistributionTotalUserCounts
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) GetWithResponseHandler(requestConfiguration *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodResponseable), nil
}
