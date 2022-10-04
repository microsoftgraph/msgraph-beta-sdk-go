package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3adc90057c795befd8028e54019c8fc8d20d74d0cdd16cef85704caf4cba6c78 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/checklistitems"
    i4b453bade987cdc7d78bf2dbba94d4dbc45a5e405d3622a58a4d26c09b49e08a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/linkedresources"
    i94c464ae9becaf7479cf2935acbe8a764ff170c65f1299d8c2f8152ac0ef8388 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/attachments"
    i9ac6ffef48c0d16af625b1cc43c9a117dc8983e6f0457047611dfe061f21a24b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/attachmentsessions"
    ie7794ee0bc0caebd554d994f150a0b208c084c890b7636d17ca1336d19fff68c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/extensions"
    i18fbd397aecd69882835c79225e981ebfe516d7af8083583884fb5e571c16f72 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/checklistitems/item"
    i6918a4356d3d8463417391beae7e08e9b5cea2610411228e7577257cbc0e08e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/extensions/item"
    i6ffb4522698e1ae7e78c00250910de6f003132aa52ac20af42871bdf610bd8fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/linkedresources/item"
    i87ee5bc633819ad2b48229ea64c8ee26fe5d521912a1bf87f7bb378baec3c015 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/attachmentsessions/item"
    ie4fa256cb1392f2daea5ba99c46b598b458d0445c15f594926f1794476e636d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/attachments/item"
)

// TodoTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.todoTaskList entity.
type TodoTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TodoTaskItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoTaskItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TodoTaskItemRequestBuilderGetQueryParameters the tasks in this task list. Read-only. Nullable.
type TodoTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TodoTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TodoTaskItemRequestBuilderGetQueryParameters
}
// TodoTaskItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TodoTaskItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments the attachments property
func (m *TodoTaskItemRequestBuilder) Attachments()(*i94c464ae9becaf7479cf2935acbe8a764ff170c65f1299d8c2f8152ac0ef8388.AttachmentsRequestBuilder) {
    return i94c464ae9becaf7479cf2935acbe8a764ff170c65f1299d8c2f8152ac0ef8388.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.tasks.item.attachments.item collection
func (m *TodoTaskItemRequestBuilder) AttachmentsById(id string)(*ie4fa256cb1392f2daea5ba99c46b598b458d0445c15f594926f1794476e636d6.AttachmentBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachmentBase%2Did"] = id
    }
    return ie4fa256cb1392f2daea5ba99c46b598b458d0445c15f594926f1794476e636d6.NewAttachmentBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttachmentSessions the attachmentSessions property
func (m *TodoTaskItemRequestBuilder) AttachmentSessions()(*i9ac6ffef48c0d16af625b1cc43c9a117dc8983e6f0457047611dfe061f21a24b.AttachmentSessionsRequestBuilder) {
    return i9ac6ffef48c0d16af625b1cc43c9a117dc8983e6f0457047611dfe061f21a24b.NewAttachmentSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentSessionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.tasks.item.attachmentSessions.item collection
func (m *TodoTaskItemRequestBuilder) AttachmentSessionsById(id string)(*i87ee5bc633819ad2b48229ea64c8ee26fe5d521912a1bf87f7bb378baec3c015.AttachmentSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachmentSession%2Did"] = id
    }
    return i87ee5bc633819ad2b48229ea64c8ee26fe5d521912a1bf87f7bb378baec3c015.NewAttachmentSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ChecklistItems the checklistItems property
func (m *TodoTaskItemRequestBuilder) ChecklistItems()(*i3adc90057c795befd8028e54019c8fc8d20d74d0cdd16cef85704caf4cba6c78.ChecklistItemsRequestBuilder) {
    return i3adc90057c795befd8028e54019c8fc8d20d74d0cdd16cef85704caf4cba6c78.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.tasks.item.checklistItems.item collection
func (m *TodoTaskItemRequestBuilder) ChecklistItemsById(id string)(*i18fbd397aecd69882835c79225e981ebfe516d7af8083583884fb5e571c16f72.ChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem%2Did"] = id
    }
    return i18fbd397aecd69882835c79225e981ebfe516d7af8083583884fb5e571c16f72.NewChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTodoTaskItemRequestBuilderInternal instantiates a new TodoTaskItemRequestBuilder and sets the default values.
func NewTodoTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TodoTaskItemRequestBuilder) {
    m := &TodoTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTodoTaskItemRequestBuilder instantiates a new TodoTaskItemRequestBuilder and sets the default values.
func NewTodoTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TodoTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTodoTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property tasks for me
func (m *TodoTaskItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TodoTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TodoTaskItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TodoTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TodoTaskItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TodoTaskable, requestConfiguration *TodoTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TodoTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TodoTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions the extensions property
func (m *TodoTaskItemRequestBuilder) Extensions()(*ie7794ee0bc0caebd554d994f150a0b208c084c890b7636d17ca1336d19fff68c.ExtensionsRequestBuilder) {
    return ie7794ee0bc0caebd554d994f150a0b208c084c890b7636d17ca1336d19fff68c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.tasks.item.extensions.item collection
func (m *TodoTaskItemRequestBuilder) ExtensionsById(id string)(*i6918a4356d3d8463417391beae7e08e9b5cea2610411228e7577257cbc0e08e8.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i6918a4356d3d8463417391beae7e08e9b5cea2610411228e7577257cbc0e08e8.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the tasks in this task list. Read-only. Nullable.
func (m *TodoTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TodoTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TodoTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTodoTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TodoTaskable), nil
}
// LinkedResources the linkedResources property
func (m *TodoTaskItemRequestBuilder) LinkedResources()(*i4b453bade987cdc7d78bf2dbba94d4dbc45a5e405d3622a58a4d26c09b49e08a.LinkedResourcesRequestBuilder) {
    return i4b453bade987cdc7d78bf2dbba94d4dbc45a5e405d3622a58a4d26c09b49e08a.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.tasks.item.linkedResources.item collection
func (m *TodoTaskItemRequestBuilder) LinkedResourcesById(id string)(*i6ffb4522698e1ae7e78c00250910de6f003132aa52ac20af42871bdf610bd8fa.LinkedResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource%2Did"] = id
    }
    return i6ffb4522698e1ae7e78c00250910de6f003132aa52ac20af42871bdf610bd8fa.NewLinkedResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property tasks in me
func (m *TodoTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TodoTaskable, requestConfiguration *TodoTaskItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TodoTaskable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTodoTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TodoTaskable), nil
}
