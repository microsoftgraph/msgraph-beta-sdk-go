package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.plannerBucket entity.
type ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters read-only. Nullable. The collection of tasks in the bucket.
type ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetQueryParameters
}
// ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedToTaskBoardFormat provides operations to manage the assignedToTaskBoardFormat property of the microsoft.graph.plannerTask entity.
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) AssignedToTaskBoardFormat()(*ItemPlannerPlansItemBucketsItemTasksItemAssignedToTaskBoardFormatRequestBuilder) {
    return NewItemPlannerPlansItemBucketsItemTasksItemAssignedToTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BucketTaskBoardFormat provides operations to manage the bucketTaskBoardFormat property of the microsoft.graph.plannerTask entity.
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) BucketTaskBoardFormat()(*ItemPlannerPlansItemBucketsItemTasksItemBucketTaskBoardFormatRequestBuilder) {
    return NewItemPlannerPlansItemBucketsItemTasksItemBucketTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal instantiates a new PlannerTaskItemRequestBuilder and sets the default values.
func NewItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) {
    m := &ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks/{plannerTask%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder instantiates a new PlannerTaskItemRequestBuilder and sets the default values.
func NewItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property tasks for groups
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation read-only. Nullable. The collection of tasks in the bucket.
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property tasks in groups
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property tasks for groups
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Details provides operations to manage the details property of the microsoft.graph.plannerTask entity.
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Details()(*ItemPlannerPlansItemBucketsItemTasksItemDetailsRequestBuilder) {
    return NewItemPlannerPlansItemBucketsItemTasksItemDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Nullable. The collection of tasks in the bucket.
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable), nil
}
// Patch update the navigation property tasks in groups
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerTaskable), nil
}
// ProgressTaskBoardFormat provides operations to manage the progressTaskBoardFormat property of the microsoft.graph.plannerTask entity.
func (m *ItemPlannerPlansItemBucketsItemTasksPlannerTaskItemRequestBuilder) ProgressTaskBoardFormat()(*ItemPlannerPlansItemBucketsItemTasksItemProgressTaskBoardFormatRequestBuilder) {
    return NewItemPlannerPlansItemBucketsItemTasksItemProgressTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
