package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
type ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderGetQueryParameters collection of tasks in the plan. Read-only. Nullable.
type ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderGetQueryParameters struct {
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
// ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderGetQueryParameters
}
// ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPlannerTaskId provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
// returns a *ItemPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder when successful
func (m *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) ByPlannerTaskId(plannerTaskId string)(*ItemPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if plannerTaskId != "" {
        urlTplParams["plannerTask%2Did"] = plannerTaskId
    }
    return NewItemPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPrimaryChannelPlannerPlansItemTasksRequestBuilderInternal instantiates a new ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder and sets the default values.
func NewItemPrimaryChannelPlannerPlansItemTasksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) {
    m := &ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/primaryChannel/planner/plans/{plannerPlan%2Did}/tasks{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPrimaryChannelPlannerPlansItemTasksRequestBuilder instantiates a new ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder and sets the default values.
func NewItemPrimaryChannelPlannerPlansItemTasksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrimaryChannelPlannerPlansItemTasksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPrimaryChannelPlannerPlansItemTasksCountRequestBuilder when successful
func (m *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) Count()(*ItemPrimaryChannelPlannerPlansItemTasksCountRequestBuilder) {
    return NewItemPrimaryChannelPlannerPlansItemTasksCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *ItemPrimaryChannelPlannerPlansItemTasksDeltaRequestBuilder when successful
func (m *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) Delta()(*ItemPrimaryChannelPlannerPlansItemTasksDeltaRequestBuilder) {
    return NewItemPrimaryChannelPlannerPlansItemTasksDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of tasks in the plan. Read-only. Nullable.
// returns a PlannerTaskCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskCollectionResponseable, error) {
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
// Post create new navigation property to tasks for teams
// returns a PlannerTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
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
// ToGetRequestInformation collection of tasks in the plan. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to tasks for teams
// returns a *RequestInformation when successful
func (m *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder when successful
func (m *ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) WithUrl(rawUrl string)(*ItemPrimaryChannelPlannerPlansItemTasksRequestBuilder) {
    return NewItemPrimaryChannelPlannerPlansItemTasksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
