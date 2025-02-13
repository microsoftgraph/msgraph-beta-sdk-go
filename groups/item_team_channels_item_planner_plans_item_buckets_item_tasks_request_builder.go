package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder provides operations to manage the tasks property of the microsoft.graph.plannerBucket entity.
type ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderGetQueryParameters read-only. Nullable. The collection of tasks in the bucket.
type ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderGetQueryParameters struct {
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
// ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderGetQueryParameters
}
// ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPlannerTaskId provides operations to manage the tasks property of the microsoft.graph.plannerBucket entity.
// returns a *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder when successful
func (m *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) ByPlannerTaskId(plannerTaskId string)(*ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if plannerTaskId != "" {
        urlTplParams["plannerTask%2Did"] = plannerTaskId
    }
    return NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderInternal instantiates a new ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder and sets the default values.
func NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) {
    m := &ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder instantiates a new ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder and sets the default values.
func NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksCountRequestBuilder when successful
func (m *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) Count()(*ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksCountRequestBuilder) {
    return NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksDeltaRequestBuilder when successful
func (m *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) Delta()(*ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksDeltaRequestBuilder) {
    return NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. The collection of tasks in the bucket.
// returns a PlannerTaskCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerTaskCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskCollectionResponseable), nil
}
// Post create new navigation property to tasks for groups
// returns a PlannerTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable), nil
}
// ToGetRequestInformation read-only. Nullable. The collection of tasks in the bucket.
// returns a *RequestInformation when successful
func (m *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to tasks for groups
// returns a *RequestInformation when successful
func (m *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder when successful
func (m *ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) WithUrl(rawUrl string)(*ItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder) {
    return NewItemTeamChannelsItemPlannerPlansItemBucketsItemTasksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
