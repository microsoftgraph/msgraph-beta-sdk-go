package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
type ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderGetQueryParameters collection of tasks in the plan. Read-only. Nullable.
type ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderGetQueryParameters
}
// ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedToTaskBoardFormat provides operations to manage the assignedToTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *ItemTeamPrimaryChannelPlannerPlansItemTasksItemAssignedToTaskBoardFormatRequestBuilder when successful
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) AssignedToTaskBoardFormat()(*ItemTeamPrimaryChannelPlannerPlansItemTasksItemAssignedToTaskBoardFormatRequestBuilder) {
    return NewItemTeamPrimaryChannelPlannerPlansItemTasksItemAssignedToTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BucketTaskBoardFormat provides operations to manage the bucketTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *ItemTeamPrimaryChannelPlannerPlansItemTasksItemBucketTaskBoardFormatRequestBuilder when successful
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) BucketTaskBoardFormat()(*ItemTeamPrimaryChannelPlannerPlansItemTasksItemBucketTaskBoardFormatRequestBuilder) {
    return NewItemTeamPrimaryChannelPlannerPlansItemTasksItemBucketTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderInternal instantiates a new ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) {
    m := &ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel/planner/plans/{plannerPlan%2Did}/tasks/{plannerTask%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder instantiates a new ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tasks for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Details provides operations to manage the details property of the microsoft.graph.plannerTask entity.
// returns a *ItemTeamPrimaryChannelPlannerPlansItemTasksItemDetailsRequestBuilder when successful
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) Details()(*ItemTeamPrimaryChannelPlannerPlansItemTasksItemDetailsRequestBuilder) {
    return NewItemTeamPrimaryChannelPlannerPlansItemTasksItemDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of tasks in the plan. Read-only. Nullable.
// returns a PlannerTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property tasks in groups
// returns a PlannerTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ProgressTaskBoardFormat provides operations to manage the progressTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *ItemTeamPrimaryChannelPlannerPlansItemTasksItemProgressTaskBoardFormatRequestBuilder when successful
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) ProgressTaskBoardFormat()(*ItemTeamPrimaryChannelPlannerPlansItemTasksItemProgressTaskBoardFormatRequestBuilder) {
    return NewItemTeamPrimaryChannelPlannerPlansItemTasksItemProgressTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property tasks for groups
// returns a *RequestInformation when successful
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of tasks in the plan. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tasks in groups
// returns a *RequestInformation when successful
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder when successful
func (m *ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) WithUrl(rawUrl string)(*ItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder) {
    return NewItemTeamPrimaryChannelPlannerPlansItemTasksPlannerTaskItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
