package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder provides operations to call the getStatistics method.
type IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal instantiates a new IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder and sets the default values.
func NewIndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    m := &IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs/microsoft.graph.industryData.getStatistics()", pathParameters),
    }
    return m
}
// NewIndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder instantiates a new IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder and sets the default values.
func NewIndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getStatistics
// returns a IndustryDataRunStatisticsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunStatisticsable, error) {
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
// ToGetRequestInformation invoke function getStatistics
// returns a *RequestInformation when successful
func (m *IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder when successful
func (m *IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) WithUrl(rawUrl string)(*IndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    return NewIndustrydataRunsMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
