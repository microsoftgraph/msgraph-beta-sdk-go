package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataRunsItemActivitiesItemActivityRequestBuilder provides operations to manage the activity property of the microsoft.graph.industryData.industryDataRunActivity entity.
type IndustrydataRunsItemActivitiesItemActivityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataRunsItemActivitiesItemActivityRequestBuilderGetQueryParameters the flow that was run by this activity.
type IndustrydataRunsItemActivitiesItemActivityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataRunsItemActivitiesItemActivityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataRunsItemActivitiesItemActivityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataRunsItemActivitiesItemActivityRequestBuilderGetQueryParameters
}
// NewIndustrydataRunsItemActivitiesItemActivityRequestBuilderInternal instantiates a new IndustrydataRunsItemActivitiesItemActivityRequestBuilder and sets the default values.
func NewIndustrydataRunsItemActivitiesItemActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsItemActivitiesItemActivityRequestBuilder) {
    m := &IndustrydataRunsItemActivitiesItemActivityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs/{industryDataRun%2Did}/activities/{industryDataRunActivity%2Did}/activity{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataRunsItemActivitiesItemActivityRequestBuilder instantiates a new IndustrydataRunsItemActivitiesItemActivityRequestBuilder and sets the default values.
func NewIndustrydataRunsItemActivitiesItemActivityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataRunsItemActivitiesItemActivityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataRunsItemActivitiesItemActivityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the flow that was run by this activity.
// returns a IndustryDataActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataRunsItemActivitiesItemActivityRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataRunsItemActivitiesItemActivityRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataActivityable), nil
}
// ToGetRequestInformation the flow that was run by this activity.
// returns a *RequestInformation when successful
func (m *IndustrydataRunsItemActivitiesItemActivityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataRunsItemActivitiesItemActivityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataRunsItemActivitiesItemActivityRequestBuilder when successful
func (m *IndustrydataRunsItemActivitiesItemActivityRequestBuilder) WithUrl(rawUrl string)(*IndustrydataRunsItemActivitiesItemActivityRequestBuilder) {
    return NewIndustrydataRunsItemActivitiesItemActivityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
