package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder provides operations to manage the activities property of the microsoft.graph.industryData.industryDataRun entity.
type IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetQueryParameters the set of activities performed during the run.
type IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetQueryParameters
}
// Activity provides operations to manage the activity property of the microsoft.graph.industryData.industryDataRunActivity entity.
func (m *IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) Activity()(*IndustryDataRunsItemActivitiesItemActivityRequestBuilder) {
    return NewIndustryDataRunsItemActivitiesItemActivityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderInternal instantiates a new IndustryDataRunActivityItemRequestBuilder and sets the default values.
func NewIndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) {
    m := &IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs/{industryDataRun%2Did}/activities/{industryDataRunActivity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder instantiates a new IndustryDataRunActivityItemRequestBuilder and sets the default values.
func NewIndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the set of activities performed during the run.
func (m *IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) WithUrl(rawUrl string)(*IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) {
    return NewIndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
