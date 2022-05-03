package getteamsuseractivitytotaldistributioncountswithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder provides operations to call the getTeamsUserActivityTotalDistributionCounts method.
type GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal instantiates a new GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    m := &GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsUserActivityTotalDistributionCounts(period='{period}')";
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
// NewGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder instantiates a new GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getTeamsUserActivityTotalDistributionCounts
func (m *GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getTeamsUserActivityTotalDistributionCounts
func (m *GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// GetWithResponseHandler invoke function getTeamsUserActivityTotalDistributionCounts
func (m *GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) GetWithResponseHandler(requestConfiguration *GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderGetRequestConfiguration)(GetTeamsUserActivityTotalDistributionCountsWithPeriodResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getTeamsUserActivityTotalDistributionCounts
func (m *GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) GetWithResponseHandler(requestConfiguration *GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetTeamsUserActivityTotalDistributionCountsWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetTeamsUserActivityTotalDistributionCountsWithPeriodResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetTeamsUserActivityTotalDistributionCountsWithPeriodResponseable), nil
}
