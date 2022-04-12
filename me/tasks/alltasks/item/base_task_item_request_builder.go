package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4d914957eb3084c5bc657dbc0bcea87fd6de701ecb15755c642a6ec8d4991865 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/move"
    i5ff907d6b73cbfca73bba51a106f2a6632bf7114dcb344eff85a8986c70e23f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/checklistitems"
    i63ee81a3dfe76eef9ca33b8a554278668a0c2207db00c252f06339df970ed8a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/extensions"
    ide349d375a635575a72282b91fadf447d228db56856a5a438bddd7d5c0f84140 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/linkedresources"
    ifd0bb0f9f690dbfe8c39f0fc370f4acb50e5071e4d99b8d5c83b43cc562eb0ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/parentlist"
    i6820be8e1566091a81ecbd31899376a78e7c325b99449ad78a8ce0df252e0f9e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/linkedresources/item"
    idf980e1c1f1f1a265b11cc926e016e5a7cf411227fcd5921660f5f9ccf16d5a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/extensions/item"
    iff25132f9ec78dc6599b2dbf74fafad86c86d8fce4ee83662fd3257c45299fb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/checklistitems/item"
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
// BaseTaskItemRequestBuilderGetQueryParameters all tasks in the users mailbox.
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
func (m *BaseTaskItemRequestBuilder) ChecklistItems()(*i5ff907d6b73cbfca73bba51a106f2a6632bf7114dcb344eff85a8986c70e23f2.ChecklistItemsRequestBuilder) {
    return i5ff907d6b73cbfca73bba51a106f2a6632bf7114dcb344eff85a8986c70e23f2.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.alltasks.item.checklistItems.item collection
func (m *BaseTaskItemRequestBuilder) ChecklistItemsById(id string)(*iff25132f9ec78dc6599b2dbf74fafad86c86d8fce4ee83662fd3257c45299fb9.ChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem%2Did"] = id
    }
    return iff25132f9ec78dc6599b2dbf74fafad86c86d8fce4ee83662fd3257c45299fb9.NewChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskItemRequestBuilderInternal instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    m := &BaseTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/tasks/alltasks/{baseTask%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property alltasks for me
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
// CreateGetRequestInformation all tasks in the users mailbox.
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
// CreatePatchRequestInformation update the navigation property alltasks in me
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
// Delete delete navigation property alltasks for me
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
func (m *BaseTaskItemRequestBuilder) Extensions()(*i63ee81a3dfe76eef9ca33b8a554278668a0c2207db00c252f06339df970ed8a8.ExtensionsRequestBuilder) {
    return i63ee81a3dfe76eef9ca33b8a554278668a0c2207db00c252f06339df970ed8a8.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.alltasks.item.extensions.item collection
func (m *BaseTaskItemRequestBuilder) ExtensionsById(id string)(*idf980e1c1f1f1a265b11cc926e016e5a7cf411227fcd5921660f5f9ccf16d5a5.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return idf980e1c1f1f1a265b11cc926e016e5a7cf411227fcd5921660f5f9ccf16d5a5.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get all tasks in the users mailbox.
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
func (m *BaseTaskItemRequestBuilder) LinkedResources()(*ide349d375a635575a72282b91fadf447d228db56856a5a438bddd7d5c0f84140.LinkedResourcesRequestBuilder) {
    return ide349d375a635575a72282b91fadf447d228db56856a5a438bddd7d5c0f84140.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.alltasks.item.linkedResources.item collection
func (m *BaseTaskItemRequestBuilder) LinkedResourcesById(id string)(*i6820be8e1566091a81ecbd31899376a78e7c325b99449ad78a8ce0df252e0f9e.LinkedResource_v2ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_v2%2Did"] = id
    }
    return i6820be8e1566091a81ecbd31899376a78e7c325b99449ad78a8ce0df252e0f9e.NewLinkedResource_v2ItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *BaseTaskItemRequestBuilder) Move()(*i4d914957eb3084c5bc657dbc0bcea87fd6de701ecb15755c642a6ec8d4991865.MoveRequestBuilder) {
    return i4d914957eb3084c5bc657dbc0bcea87fd6de701ecb15755c642a6ec8d4991865.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentList the parentList property
func (m *BaseTaskItemRequestBuilder) ParentList()(*ifd0bb0f9f690dbfe8c39f0fc370f4acb50e5071e4d99b8d5c83b43cc562eb0ed.ParentListRequestBuilder) {
    return ifd0bb0f9f690dbfe8c39f0fc370f4acb50e5071e4d99b8d5c83b43cc562eb0ed.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property alltasks in me
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
