package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataRunsRequestBuilder provides operations to manage the runs property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustryDataRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataRunsRequestBuilderGetQueryParameters set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system. Read-only.
type IndustryDataRunsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// IndustryDataRunsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataRunsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataRunsRequestBuilderGetQueryParameters
}
// ByIndustryDataRunId provides operations to manage the runs property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustryDataRunsIndustryDataRunItemRequestBuilder when successful
func (m *IndustryDataRunsRequestBuilder) ByIndustryDataRunId(industryDataRunId string)(*IndustryDataRunsIndustryDataRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if industryDataRunId != "" {
        urlTplParams["industryDataRun%2Did"] = industryDataRunId
    }
    return NewIndustryDataRunsIndustryDataRunItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustryDataRunsRequestBuilderInternal instantiates a new IndustryDataRunsRequestBuilder and sets the default values.
func NewIndustryDataRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsRequestBuilder) {
    m := &IndustryDataRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIndustryDataRunsRequestBuilder instantiates a new IndustryDataRunsRequestBuilder and sets the default values.
func NewIndustryDataRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IndustryDataRunsCountRequestBuilder when successful
func (m *IndustryDataRunsRequestBuilder) Count()(*IndustryDataRunsCountRequestBuilder) {
    return NewIndustryDataRunsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system. Read-only.
// returns a IndustryDataRunCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustryDataRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataRunsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataRunCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunCollectionResponseable), nil
}
// MicrosoftGraphIndustryDataGetStatistics provides operations to call the getStatistics method.
// returns a *IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder when successful
func (m *IndustryDataRunsRequestBuilder) MicrosoftGraphIndustryDataGetStatistics()(*IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    return NewIndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system. Read-only.
// returns a *RequestInformation when successful
func (m *IndustryDataRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataRunsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustryDataRunsRequestBuilder when successful
func (m *IndustryDataRunsRequestBuilder) WithUrl(rawUrl string)(*IndustryDataRunsRequestBuilder) {
    return NewIndustryDataRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
