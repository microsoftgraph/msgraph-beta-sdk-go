package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
type ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderGetQueryParameters collection of tasks in the plan. Read-only. Nullable.
type ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderGetQueryParameters struct {
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
// ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPlannerTaskId provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
// returns a *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksPlannerTaskItemRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) ByPlannerTaskId(plannerTaskId string)(*ItemTeamDefinitionChannelsItemPlannerPlansItemTasksPlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if plannerTaskId != "" {
        urlTplParams["plannerTask%2Did"] = plannerTaskId
    }
    return NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksPlannerTaskItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderInternal instantiates a new ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) {
    m := &ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/planner/plans/{plannerPlan%2Did}/tasks{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder instantiates a new ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksCountRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) Count()(*ItemTeamDefinitionChannelsItemPlannerPlansItemTasksCountRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksDeltaRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) Delta()(*ItemTeamDefinitionChannelsItemPlannerPlansItemTasksDeltaRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of tasks in the plan. Read-only. Nullable.
// returns a PlannerTaskCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskCollectionResponseable, error) {
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
// Post create new navigation property to tasks for teamTemplateDefinition
// returns a PlannerTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
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
func (m *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to tasks for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemPlannerPlansItemTasksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
