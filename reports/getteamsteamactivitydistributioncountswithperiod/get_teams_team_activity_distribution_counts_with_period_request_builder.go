package getteamsteamactivitydistributioncountswithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder provides operations to call the getTeamsTeamActivityDistributionCounts method.
type GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderGetOptions options for Get
type GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal instantiates a new GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    m := &GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsTeamActivityDistributionCounts(period='{period}')";
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
// NewGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder instantiates a new GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder and sets the default values.
func NewGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getTeamsTeamActivityDistributionCounts
func (m *GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getTeamsTeamActivityDistributionCounts
func (m *GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) Get(options *GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderGetOptions)(GetTeamsTeamActivityDistributionCountsWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetTeamsTeamActivityDistributionCountsWithPeriodResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetTeamsTeamActivityDistributionCountsWithPeriodResponseable), nil
}
