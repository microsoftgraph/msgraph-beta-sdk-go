package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i275fea45cd066f9262919593c03c0b1e1e7e71a727b5af64101dd64bc6e46b89 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/lists/item/tasks/item/linkedresources"
    i7260b89120c013953545f8577655f38687e6049054459b4979fc66317449b7d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/lists/item/tasks/item/parentlist"
    i7bfd138746038534bad1aac3dce035ec42cf5d70a13a1413d2f2e46892f38d09 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/lists/item/tasks/item/checklistitems"
    ia9d0c01958ba0799d3542155fd11213ed561dc63cc667c395497892e9e77f1d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/lists/item/tasks/item/extensions"
    ic7aa8eb0e8aefbfa6858e7f2fc465b90d431c4e0f73bb00cf8b56f796639bf40 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/lists/item/tasks/item/move"
    i34e9ef0b7eef0d6704c833332ffa8d16c019e313b1def9f065f50d25a8f8068a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/lists/item/tasks/item/linkedresources/item"
    i49f37fa0ca8dbd0779122380b5f0f7e39b7ff61cbc48e2b361a00f669235f391 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/lists/item/tasks/item/extensions/item"
    i6044c1d2eb01b0e0622e3b6c3849a33775aab29f024a15554528173596ac8f6b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/lists/item/tasks/item/checklistitems/item"
)

// BaseTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.baseTaskList entity.
type BaseTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BaseTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BaseTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BaseTaskItemRequestBuilderGetQueryParameters the tasks in this task list. Read-only. Nullable.
type BaseTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BaseTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BaseTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BaseTaskItemRequestBuilderGetQueryParameters
}
// BaseTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BaseTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChecklistItems provides operations to manage the checklistItems property of the microsoft.graph.baseTask entity.
func (m *BaseTaskItemRequestBuilder) ChecklistItems()(*i7bfd138746038534bad1aac3dce035ec42cf5d70a13a1413d2f2e46892f38d09.ChecklistItemsRequestBuilder) {
    return i7bfd138746038534bad1aac3dce035ec42cf5d70a13a1413d2f2e46892f38d09.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById provides operations to manage the checklistItems property of the microsoft.graph.baseTask entity.
func (m *BaseTaskItemRequestBuilder) ChecklistItemsById(id string)(*i6044c1d2eb01b0e0622e3b6c3849a33775aab29f024a15554528173596ac8f6b.ChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem%2Did"] = id
    }
    return i6044c1d2eb01b0e0622e3b6c3849a33775aab29f024a15554528173596ac8f6b.NewChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskItemRequestBuilderInternal instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    m := &BaseTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/tasks/lists/{baseTaskList%2Did}/tasks/{baseTask%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseTaskItemRequestBuilder instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property tasks for me
func (m *BaseTaskItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *BaseTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *BaseTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property tasks in me
func (m *BaseTaskItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable, requestConfiguration *BaseTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property tasks for me
func (m *BaseTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BaseTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.baseTask entity.
func (m *BaseTaskItemRequestBuilder) Extensions()(*ia9d0c01958ba0799d3542155fd11213ed561dc63cc667c395497892e9e77f1d3.ExtensionsRequestBuilder) {
    return ia9d0c01958ba0799d3542155fd11213ed561dc63cc667c395497892e9e77f1d3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.baseTask entity.
func (m *BaseTaskItemRequestBuilder) ExtensionsById(id string)(*i49f37fa0ca8dbd0779122380b5f0f7e39b7ff61cbc48e2b361a00f669235f391.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i49f37fa0ca8dbd0779122380b5f0f7e39b7ff61cbc48e2b361a00f669235f391.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BaseTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBaseTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable), nil
}
// LinkedResources provides operations to manage the linkedResources property of the microsoft.graph.baseTask entity.
func (m *BaseTaskItemRequestBuilder) LinkedResources()(*i275fea45cd066f9262919593c03c0b1e1e7e71a727b5af64101dd64bc6e46b89.LinkedResourcesRequestBuilder) {
    return i275fea45cd066f9262919593c03c0b1e1e7e71a727b5af64101dd64bc6e46b89.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById provides operations to manage the linkedResources property of the microsoft.graph.baseTask entity.
func (m *BaseTaskItemRequestBuilder) LinkedResourcesById(id string)(*i34e9ef0b7eef0d6704c833332ffa8d16c019e313b1def9f065f50d25a8f8068a.LinkedResource_v2ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_v2%2Did"] = id
    }
    return i34e9ef0b7eef0d6704c833332ffa8d16c019e313b1def9f065f50d25a8f8068a.NewLinkedResource_v2ItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move provides operations to call the move method.
func (m *BaseTaskItemRequestBuilder) Move()(*ic7aa8eb0e8aefbfa6858e7f2fc465b90d431c4e0f73bb00cf8b56f796639bf40.MoveRequestBuilder) {
    return ic7aa8eb0e8aefbfa6858e7f2fc465b90d431c4e0f73bb00cf8b56f796639bf40.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentList provides operations to manage the parentList property of the microsoft.graph.baseTask entity.
func (m *BaseTaskItemRequestBuilder) ParentList()(*i7260b89120c013953545f8577655f38687e6049054459b4979fc66317449b7d4.ParentListRequestBuilder) {
    return i7260b89120c013953545f8577655f38687e6049054459b4979fc66317449b7d4.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property tasks in me
func (m *BaseTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable, requestConfiguration *BaseTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBaseTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable), nil
}
