package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i889297e93166312716592efbb05995455ae1882e388416fee06599864e54e462 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/tasks"
    i97f4aa0504a94677b48d17585be69f3adfec35ef126990f5c0fa1deb971970ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/buckets"
    i9800485461c637fd964b8b604adecc9eb52f29415b0fdcbce97286ef735a649d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/details"
    i3466366278330596cef17ffa7a24be314ef4f15e7e9d83c16fc122c5d7bbec9d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/buckets/item"
    ie21edb59ce06faac589655d693ce3b2317a4cec27df5b0fb71362a772a5e4fbd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/tasks/item"
)

// PlannerPlanItemRequestBuilder provides operations to manage the plans property of the microsoft.graph.plannerUser entity.
type PlannerPlanItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PlannerPlanItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerPlanItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PlannerPlanItemRequestBuilderGetQueryParameters get plans from me
type PlannerPlanItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PlannerPlanItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerPlanItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PlannerPlanItemRequestBuilderGetQueryParameters
}
// PlannerPlanItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerPlanItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Buckets provides operations to manage the buckets property of the microsoft.graph.plannerPlan entity.
func (m *PlannerPlanItemRequestBuilder) Buckets()(*i97f4aa0504a94677b48d17585be69f3adfec35ef126990f5c0fa1deb971970ee.BucketsRequestBuilder) {
    return i97f4aa0504a94677b48d17585be69f3adfec35ef126990f5c0fa1deb971970ee.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BucketsById provides operations to manage the buckets property of the microsoft.graph.plannerPlan entity.
func (m *PlannerPlanItemRequestBuilder) BucketsById(id string)(*i3466366278330596cef17ffa7a24be314ef4f15e7e9d83c16fc122c5d7bbec9d.PlannerBucketItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket%2Did"] = id
    }
    return i3466366278330596cef17ffa7a24be314ef4f15e7e9d83c16fc122c5d7bbec9d.NewPlannerBucketItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPlannerPlanItemRequestBuilderInternal instantiates a new PlannerPlanItemRequestBuilder and sets the default values.
func NewPlannerPlanItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerPlanItemRequestBuilder) {
    m := &PlannerPlanItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/planner/plans/{plannerPlan%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerPlanItemRequestBuilder instantiates a new PlannerPlanItemRequestBuilder and sets the default values.
func NewPlannerPlanItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerPlanItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerPlanItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property plans for me
func (m *PlannerPlanItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PlannerPlanItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get plans from me
func (m *PlannerPlanItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PlannerPlanItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property plans in me
func (m *PlannerPlanItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, requestConfiguration *PlannerPlanItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property plans for me
func (m *PlannerPlanItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PlannerPlanItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Details provides operations to manage the details property of the microsoft.graph.plannerPlan entity.
func (m *PlannerPlanItemRequestBuilder) Details()(*i9800485461c637fd964b8b604adecc9eb52f29415b0fdcbce97286ef735a649d.DetailsRequestBuilder) {
    return i9800485461c637fd964b8b604adecc9eb52f29415b0fdcbce97286ef735a649d.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get plans from me
func (m *PlannerPlanItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PlannerPlanItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable), nil
}
// Patch update the navigation property plans in me
func (m *PlannerPlanItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, requestConfiguration *PlannerPlanItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable), nil
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
func (m *PlannerPlanItemRequestBuilder) Tasks()(*i889297e93166312716592efbb05995455ae1882e388416fee06599864e54e462.TasksRequestBuilder) {
    return i889297e93166312716592efbb05995455ae1882e388416fee06599864e54e462.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.plannerPlan entity.
func (m *PlannerPlanItemRequestBuilder) TasksById(id string)(*ie21edb59ce06faac589655d693ce3b2317a4cec27df5b0fb71362a772a5e4fbd.PlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask%2Did"] = id
    }
    return ie21edb59ce06faac589655d693ce3b2317a4cec27df5b0fb71362a772a5e4fbd.NewPlannerTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
