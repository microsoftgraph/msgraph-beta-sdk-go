package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder provides operations to call the getStatistics method.
type IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal instantiates a new IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder and sets the default values.
func NewIndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    m := &IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs/{industryDataRun%2Did}/microsoft.graph.industryData.getStatistics()", pathParameters),
    }
    return m
}
// NewIndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder instantiates a new IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder and sets the default values.
func NewIndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get statistics for an industryDataRun.
// returns a IndustryDataRunStatisticsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-industrydatarun-getstatistics?view=graph-rest-beta
func (m *IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunStatisticsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataRunStatisticsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunStatisticsable), nil
}
// ToGetRequestInformation get statistics for an industryDataRun.
// returns a *RequestInformation when successful
func (m *IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder when successful
func (m *IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) WithUrl(rawUrl string)(*IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    return NewIndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
