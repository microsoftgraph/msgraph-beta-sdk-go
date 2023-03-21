package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataRunsIndustryDataRunItemRequestBuilder provides operations to manage the runs property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustryDataRunsIndustryDataRunItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IndustryDataRunsIndustryDataRunItemRequestBuilderGetQueryParameters set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system. Read-only.
type IndustryDataRunsIndustryDataRunItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustryDataRunsIndustryDataRunItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataRunsIndustryDataRunItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataRunsIndustryDataRunItemRequestBuilderGetQueryParameters
}
// Activities provides operations to manage the activities property of the microsoft.graph.industryData.industryDataRun entity.
func (m *IndustryDataRunsIndustryDataRunItemRequestBuilder) Activities()(*IndustryDataRunsItemActivitiesRequestBuilder) {
    return NewIndustryDataRunsItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.industryData.industryDataRun entity.
func (m *IndustryDataRunsIndustryDataRunItemRequestBuilder) ActivitiesById(id string)(*IndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["industryDataRunActivity%2Did"] = id
    }
    return NewIndustryDataRunsItemActivitiesIndustryDataRunActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewIndustryDataRunsIndustryDataRunItemRequestBuilderInternal instantiates a new IndustryDataRunItemRequestBuilder and sets the default values.
func NewIndustryDataRunsIndustryDataRunItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsIndustryDataRunItemRequestBuilder) {
    m := &IndustryDataRunsIndustryDataRunItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/external/industryData/runs/{industryDataRun%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewIndustryDataRunsIndustryDataRunItemRequestBuilder instantiates a new IndustryDataRunItemRequestBuilder and sets the default values.
func NewIndustryDataRunsIndustryDataRunItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsIndustryDataRunItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataRunsIndustryDataRunItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system. Read-only.
func (m *IndustryDataRunsIndustryDataRunItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataRunsIndustryDataRunItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataRunFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunable), nil
}
// IndustryDataGetStatistics provides operations to call the getStatistics method.
func (m *IndustryDataRunsIndustryDataRunItemRequestBuilder) IndustryDataGetStatistics()(*IndustryDataRunsItemIndustryDataGetStatisticsRequestBuilder) {
    return NewIndustryDataRunsItemIndustryDataGetStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system. Read-only.
func (m *IndustryDataRunsIndustryDataRunItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataRunsIndustryDataRunItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
