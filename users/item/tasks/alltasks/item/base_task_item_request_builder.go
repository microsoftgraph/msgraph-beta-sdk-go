package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i74a32fabbfbde9205732b623e943f248ae1993d1dd3bd7ca652fe67e4b8aaf82 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/parentlist"
    i8f577f5802f3e0e15c989cddb25554bb5b50f43e79e58f4ffc257d28e636a543 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/linkedresources"
    i9c1618f121ae55105f68bdb7b5d0f4357960354111636ecc523efaeed30cf6d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/checklistitems"
    ib8c99b5c2cdc8e87f18e14d0083feae23e011907d486bb6148cf82988d5f0b41 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/move"
    if5e5723fea0ffe4b997fea3274e0e2616624b553461f0e7384262e884134ab9d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/extensions"
    i12411144c67e21275a25f893f2d9baafcc6a3008328f4ae687d8e4ece105c91d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/checklistitems/item"
    i8aa5d29987d083cb64fee493022b81d976f48df901531ff3abc4b55339d9ae52 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/extensions/item"
    ic1dc19df48522a0c8b54b952e537056a62a5f6c330e50338e3f20f5e81e3a64f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/linkedresources/item"
)

// BaseTaskItemRequestBuilder provides operations to manage the alltasks property of the microsoft.graph.tasks entity.
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
// BaseTaskItemRequestBuilderGetQueryParameters all tasks in the users mailbox.
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
// ChecklistItems the checklistItems property
func (m *BaseTaskItemRequestBuilder) ChecklistItems()(*i9c1618f121ae55105f68bdb7b5d0f4357960354111636ecc523efaeed30cf6d3.ChecklistItemsRequestBuilder) {
    return i9c1618f121ae55105f68bdb7b5d0f4357960354111636ecc523efaeed30cf6d3.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.alltasks.item.checklistItems.item collection
func (m *BaseTaskItemRequestBuilder) ChecklistItemsById(id string)(*i12411144c67e21275a25f893f2d9baafcc6a3008328f4ae687d8e4ece105c91d.ChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem%2Did"] = id
    }
    return i12411144c67e21275a25f893f2d9baafcc6a3008328f4ae687d8e4ece105c91d.NewChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskItemRequestBuilderInternal instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    m := &BaseTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/tasks/alltasks/{baseTask%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property alltasks for users
func (m *BaseTaskItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property alltasks for users
func (m *BaseTaskItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *BaseTaskItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *BaseTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property alltasks in users
func (m *BaseTaskItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property alltasks in users
func (m *BaseTaskItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable, requestConfiguration *BaseTaskItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property alltasks for users
func (m *BaseTaskItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BaseTaskItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *BaseTaskItemRequestBuilder) Extensions()(*if5e5723fea0ffe4b997fea3274e0e2616624b553461f0e7384262e884134ab9d.ExtensionsRequestBuilder) {
    return if5e5723fea0ffe4b997fea3274e0e2616624b553461f0e7384262e884134ab9d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.alltasks.item.extensions.item collection
func (m *BaseTaskItemRequestBuilder) ExtensionsById(id string)(*i8aa5d29987d083cb64fee493022b81d976f48df901531ff3abc4b55339d9ae52.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8aa5d29987d083cb64fee493022b81d976f48df901531ff3abc4b55339d9ae52.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BaseTaskItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// LinkedResources the linkedResources property
func (m *BaseTaskItemRequestBuilder) LinkedResources()(*i8f577f5802f3e0e15c989cddb25554bb5b50f43e79e58f4ffc257d28e636a543.LinkedResourcesRequestBuilder) {
    return i8f577f5802f3e0e15c989cddb25554bb5b50f43e79e58f4ffc257d28e636a543.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.alltasks.item.linkedResources.item collection
func (m *BaseTaskItemRequestBuilder) LinkedResourcesById(id string)(*ic1dc19df48522a0c8b54b952e537056a62a5f6c330e50338e3f20f5e81e3a64f.LinkedResource_v2ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_v2%2Did"] = id
    }
    return ic1dc19df48522a0c8b54b952e537056a62a5f6c330e50338e3f20f5e81e3a64f.NewLinkedResource_v2ItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *BaseTaskItemRequestBuilder) Move()(*ib8c99b5c2cdc8e87f18e14d0083feae23e011907d486bb6148cf82988d5f0b41.MoveRequestBuilder) {
    return ib8c99b5c2cdc8e87f18e14d0083feae23e011907d486bb6148cf82988d5f0b41.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentList the parentList property
func (m *BaseTaskItemRequestBuilder) ParentList()(*i74a32fabbfbde9205732b623e943f248ae1993d1dd3bd7ca652fe67e4b8aaf82.ParentListRequestBuilder) {
    return i74a32fabbfbde9205732b623e943f248ae1993d1dd3bd7ca652fe67e4b8aaf82.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property alltasks in users
func (m *BaseTaskItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseTaskable, requestConfiguration *BaseTaskItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
