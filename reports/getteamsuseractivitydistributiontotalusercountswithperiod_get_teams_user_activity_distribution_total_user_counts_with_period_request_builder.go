package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder provides operations to call the getTeamsUserActivityDistributionTotalUserCounts method.
type GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal instantiates a new GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    m := &GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getTeamsUserActivityDistributionTotalUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder instantiates a new GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getTeamsUserActivityDistributionTotalUserCounts
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getTeamsUserActivityDistributionTotalUserCounts
// returns a *RequestInformation when successful
func (m *GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder when successful
func (m *GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewGetteamsuseractivitydistributiontotalusercountswithperiodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
