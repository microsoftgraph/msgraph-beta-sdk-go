package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i63585e4f9db5ce04afe272d39ba9e0f814d15a9ad69564f035b57992fdc60d0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/linkedresources"
    i7eb8a811f65d6bab93f6143e18dd050750b3187b5222bba2363767dea08f400e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/move"
    i9f90a854cb36ba230b51f79e03cf8c485003d9c20892f3aaaf9f9ca6e2311052 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/checklistitems"
    ic3517bb76f041a0e9a13acef1cef85787191fc80023465b6551e0527d3bb7dcf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/parentlist"
    if2aa69e4c699585c66a9474fde03947bf3cd63c88d2a207d476ff72016bb92ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/extensions"
    i400d2cff755de6f49886cdc0854849d644b5a10323083dc756fb70061fd9b716 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/extensions/item"
    i43baf14bdfa52348b0bb48c3d3ece5a20da9f14a5275a095129997e3eaebe805 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/checklistitems/item"
    ic96ef0290e7cd1ba8655c202ac321602c963f0f41deb8e52900dee755247715e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/linkedresources/item"
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
// BaseTaskItemRequestBuilderDeleteOptions options for Delete
type BaseTaskItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// BaseTaskItemRequestBuilderGetOptions options for Get
type BaseTaskItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BaseTaskItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// BaseTaskItemRequestBuilderGetQueryParameters the tasks in this task list. Read-only. Nullable.
type BaseTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BaseTaskItemRequestBuilderPatchOptions options for Patch
type BaseTaskItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChecklistItems the checklistItems property
func (m *BaseTaskItemRequestBuilder) ChecklistItems()(*i9f90a854cb36ba230b51f79e03cf8c485003d9c20892f3aaaf9f9ca6e2311052.ChecklistItemsRequestBuilder) {
    return i9f90a854cb36ba230b51f79e03cf8c485003d9c20892f3aaaf9f9ca6e2311052.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.lists.item.tasks.item.checklistItems.item collection
func (m *BaseTaskItemRequestBuilder) ChecklistItemsById(id string)(*i43baf14bdfa52348b0bb48c3d3ece5a20da9f14a5275a095129997e3eaebe805.ChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem%2Did"] = id
    }
    return i43baf14bdfa52348b0bb48c3d3ece5a20da9f14a5275a095129997e3eaebe805.NewChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskItemRequestBuilderInternal instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    m := &BaseTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/tasks/lists/{baseTaskList%2Did}/tasks/{baseTask%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property tasks for users
func (m *BaseTaskItemRequestBuilder) CreateDeleteRequestInformation(options *BaseTaskItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskItemRequestBuilder) CreateGetRequestInformation(options *BaseTaskItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property tasks in users
func (m *BaseTaskItemRequestBuilder) CreatePatchRequestInformation(options *BaseTaskItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property tasks for users
func (m *BaseTaskItemRequestBuilder) Delete(options *BaseTaskItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
func (m *BaseTaskItemRequestBuilder) Extensions()(*if2aa69e4c699585c66a9474fde03947bf3cd63c88d2a207d476ff72016bb92ed.ExtensionsRequestBuilder) {
    return if2aa69e4c699585c66a9474fde03947bf3cd63c88d2a207d476ff72016bb92ed.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.lists.item.tasks.item.extensions.item collection
func (m *BaseTaskItemRequestBuilder) ExtensionsById(id string)(*i400d2cff755de6f49886cdc0854849d644b5a10323083dc756fb70061fd9b716.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i400d2cff755de6f49886cdc0854849d644b5a10323083dc756fb70061fd9b716.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskItemRequestBuilder) Get(options *BaseTaskItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBaseTaskFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable), nil
}
// LinkedResources the linkedResources property
func (m *BaseTaskItemRequestBuilder) LinkedResources()(*i63585e4f9db5ce04afe272d39ba9e0f814d15a9ad69564f035b57992fdc60d0e.LinkedResourcesRequestBuilder) {
    return i63585e4f9db5ce04afe272d39ba9e0f814d15a9ad69564f035b57992fdc60d0e.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.lists.item.tasks.item.linkedResources.item collection
func (m *BaseTaskItemRequestBuilder) LinkedResourcesById(id string)(*ic96ef0290e7cd1ba8655c202ac321602c963f0f41deb8e52900dee755247715e.LinkedResource_v2ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_v2%2Did"] = id
    }
    return ic96ef0290e7cd1ba8655c202ac321602c963f0f41deb8e52900dee755247715e.NewLinkedResource_v2ItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *BaseTaskItemRequestBuilder) Move()(*i7eb8a811f65d6bab93f6143e18dd050750b3187b5222bba2363767dea08f400e.MoveRequestBuilder) {
    return i7eb8a811f65d6bab93f6143e18dd050750b3187b5222bba2363767dea08f400e.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentList the parentList property
func (m *BaseTaskItemRequestBuilder) ParentList()(*ic3517bb76f041a0e9a13acef1cef85787191fc80023465b6551e0527d3bb7dcf.ParentListRequestBuilder) {
    return ic3517bb76f041a0e9a13acef1cef85787191fc80023465b6551e0527d3bb7dcf.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property tasks in users
func (m *BaseTaskItemRequestBuilder) Patch(options *BaseTaskItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
