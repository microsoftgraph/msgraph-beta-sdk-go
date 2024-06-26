package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.userProcessingResult entity.
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderGetQueryParameters get the task processing result from a userProcessingResult either directly or through a run.
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderGetQueryParameters struct {
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
// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderGetQueryParameters
}
// ByTaskProcessingResultId provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.userProcessingResult entity.
// returns a *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) ByTaskProcessingResultId(taskProcessingResultId string)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if taskProcessingResultId != "" {
        urlTplParams["taskProcessingResult%2Did"] = taskProcessingResultId
    }
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/userProcessingResults/{userProcessingResult%2Did}/taskProcessingResults{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsCountRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) Count()(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsCountRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the task processing result from a userProcessingResult either directly or through a run.
// returns a TaskProcessingResultCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-userprocessingresult-list-taskprocessingresults?view=graph-rest-beta
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskProcessingResultCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateTaskProcessingResultCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskProcessingResultCollectionResponseable), nil
}
// ToGetRequestInformation get the task processing result from a userProcessingResult either directly or through a run.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
