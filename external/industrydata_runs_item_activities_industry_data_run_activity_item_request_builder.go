package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder provides operations to manage the activities property of the microsoft.graph.industryData.industryDataRun entity.
type IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetQueryParameters the set of activities performed during the run.
type IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetQueryParameters
}
// Activity provides operations to manage the activity property of the microsoft.graph.industryData.industryDataRunActivity entity.
// returns a *IndustrydataRunsItemActivitiesItemActivityRequestBuilder when successful
func (m *IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) Activity()(*IndustrydataRunsItemActivitiesItemActivityRequestBuilder) {
    return NewIndustrydataRunsItemActivitiesItemActivityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderInternal instantiates a new IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder and sets the default values.
func NewIndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) {
    m := &IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs/{industryDataRun%2Did}/activities/{industryDataRunActivity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder instantiates a new IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder and sets the default values.
func NewIndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the set of activities performed during the run.
// returns a IndustryDataRunActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataRunActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunActivityable), nil
}
// ToGetRequestInformation the set of activities performed during the run.
// returns a *RequestInformation when successful
func (m *IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder when successful
func (m *IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) WithUrl(rawUrl string)(*IndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) {
    return NewIndustrydataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
