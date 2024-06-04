package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.businessScenarioPlanner entity.
type BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderGetQueryParameters read the properties and relationships of a businessScenarioTask object.
type BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderGetQueryParameters
}
// BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedToTaskBoardFormat provides operations to manage the assignedToTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *BusinessscenariosItemPlannerTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder when successful
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) AssignedToTaskBoardFormat()(*BusinessscenariosItemPlannerTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) {
    return NewBusinessscenariosItemPlannerTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BucketTaskBoardFormat provides operations to manage the bucketTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *BusinessscenariosItemPlannerTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder when successful
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) BucketTaskBoardFormat()(*BusinessscenariosItemPlannerTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) {
    return NewBusinessscenariosItemPlannerTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewBusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderInternal instantiates a new BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) {
    m := &BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner/tasks/{businessScenarioTask%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder instantiates a new BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a businessScenarioTask object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/businessscenarioplanner-delete-tasks?view=graph-rest-beta
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *BusinessscenariosItemPlannerTasksItemDetailsRequestBuilder when successful
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) Details()(*BusinessscenariosItemPlannerTasksItemDetailsRequestBuilder) {
    return NewBusinessscenariosItemPlannerTasksItemDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a businessScenarioTask object.
// returns a BusinessScenarioTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/businessscenariotask-get?view=graph-rest-beta
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable), nil
}
// Patch update the properties of a businessScenarioTask object.
// returns a BusinessScenarioTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/businessscenariotask-update?view=graph-rest-beta
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable, requestConfiguration *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable), nil
}
// ProgressTaskBoardFormat provides operations to manage the progressTaskBoardFormat property of the microsoft.graph.plannerTask entity.
// returns a *BusinessscenariosItemPlannerTasksItemProgresstaskboardformatProgressTaskBoardFormatRequestBuilder when successful
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) ProgressTaskBoardFormat()(*BusinessscenariosItemPlannerTasksItemProgresstaskboardformatProgressTaskBoardFormatRequestBuilder) {
    return NewBusinessscenariosItemPlannerTasksItemProgresstaskboardformatProgressTaskBoardFormatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a businessScenarioTask object.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a businessScenarioTask object.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a businessScenarioTask object.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable, requestConfiguration *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder when successful
func (m *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) WithUrl(rawUrl string)(*BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) {
    return NewBusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
