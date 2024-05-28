package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataRunsIndustryDataRunItemRequestBuilder provides operations to manage the runs property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustrydataRunsIndustryDataRunItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataRunsIndustryDataRunItemRequestBuilderGetQueryParameters read the properties and relationships of an industryDataRun object.
type IndustrydataRunsIndustryDataRunItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataRunsIndustryDataRunItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataRunsIndustryDataRunItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataRunsIndustryDataRunItemRequestBuilderGetQueryParameters
}
// Activities provides operations to manage the activities property of the microsoft.graph.industryData.industryDataRun entity.
// returns a *IndustrydataRunsItemActivitiesRequestBuilder when successful
func (m *IndustrydataRunsIndustryDataRunItemRequestBuilder) Activities()(*IndustrydataRunsItemActivitiesRequestBuilder) {
    return NewIndustrydataRunsItemActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustrydataRunsIndustryDataRunItemRequestBuilderInternal instantiates a new IndustrydataRunsIndustryDataRunItemRequestBuilder and sets the default values.
func NewIndustrydataRunsIndustryDataRunItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsIndustryDataRunItemRequestBuilder) {
    m := &IndustrydataRunsIndustryDataRunItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs/{industryDataRun%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataRunsIndustryDataRunItemRequestBuilder instantiates a new IndustrydataRunsIndustryDataRunItemRequestBuilder and sets the default values.
func NewIndustrydataRunsIndustryDataRunItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsIndustryDataRunItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataRunsIndustryDataRunItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of an industryDataRun object.
// returns a IndustryDataRunable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-industrydatarun-get?view=graph-rest-beta
func (m *IndustrydataRunsIndustryDataRunItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataRunsIndustryDataRunItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataRunFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunable), nil
}
// MicrosoftGraphIndustryDataGetStatistics provides operations to call the getStatistics method.
// returns a *IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder when successful
func (m *IndustrydataRunsIndustryDataRunItemRequestBuilder) MicrosoftGraphIndustryDataGetStatistics()(*IndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    return NewIndustrydataRunsItemMicrosoftgraphindustrydatagetstatisticsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read the properties and relationships of an industryDataRun object.
// returns a *RequestInformation when successful
func (m *IndustrydataRunsIndustryDataRunItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataRunsIndustryDataRunItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IndustrydataRunsIndustryDataRunItemRequestBuilder when successful
func (m *IndustrydataRunsIndustryDataRunItemRequestBuilder) WithUrl(rawUrl string)(*IndustrydataRunsIndustryDataRunItemRequestBuilder) {
    return NewIndustrydataRunsIndustryDataRunItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
