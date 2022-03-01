package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i74a32fabbfbde9205732b623e943f248ae1993d1dd3bd7ca652fe67e4b8aaf82 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/parentlist"
    i8f577f5802f3e0e15c989cddb25554bb5b50f43e79e58f4ffc257d28e636a543 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/linkedresources"
    i9c1618f121ae55105f68bdb7b5d0f4357960354111636ecc523efaeed30cf6d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/checklistitems"
    ib8c99b5c2cdc8e87f18e14d0083feae23e011907d486bb6148cf82988d5f0b41 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/move"
    if5e5723fea0ffe4b997fea3274e0e2616624b553461f0e7384262e884134ab9d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/extensions"
    i12411144c67e21275a25f893f2d9baafcc6a3008328f4ae687d8e4ece105c91d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/checklistitems/item"
    i8aa5d29987d083cb64fee493022b81d976f48df901531ff3abc4b55339d9ae52 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/extensions/item"
    ic1dc19df48522a0c8b54b952e537056a62a5f6c330e50338e3f20f5e81e3a64f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/linkedresources/item"
)

// BaseTaskItemRequestBuilder builds and executes requests for operations under \users\{user-id}\tasks\alltasks\{baseTask-id}
type BaseTaskItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BaseTaskItemRequestBuilderDeleteOptions options for Delete
type BaseTaskItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTaskItemRequestBuilderGetOptions options for Get
type BaseTaskItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BaseTaskItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTaskItemRequestBuilderGetQueryParameters all tasks in the users mailbox.
type BaseTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BaseTaskItemRequestBuilderPatchOptions options for Patch
type BaseTaskItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTask;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["checklistItem_id"] = id
    }
    return i12411144c67e21275a25f893f2d9baafcc6a3008328f4ae687d8e4ece105c91d.NewChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskItemRequestBuilderInternal instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    m := &BaseTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/tasks/alltasks/{baseTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseTaskItemRequestBuilder instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) CreateDeleteRequestInformation(options *BaseTaskItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) CreateGetRequestInformation(options *BaseTaskItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) CreatePatchRequestInformation(options *BaseTaskItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) Delete(options *BaseTaskItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["extension_id"] = id
    }
    return i8aa5d29987d083cb64fee493022b81d976f48df901531ff3abc4b55339d9ae52.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) Get(options *BaseTaskItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBaseTask() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTask), nil
}
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
        urlTplParams["linkedResource_v2_id"] = id
    }
    return ic1dc19df48522a0c8b54b952e537056a62a5f6c330e50338e3f20f5e81e3a64f.NewLinkedResource_v2ItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BaseTaskItemRequestBuilder) Move()(*ib8c99b5c2cdc8e87f18e14d0083feae23e011907d486bb6148cf82988d5f0b41.MoveRequestBuilder) {
    return ib8c99b5c2cdc8e87f18e14d0083feae23e011907d486bb6148cf82988d5f0b41.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTaskItemRequestBuilder) ParentList()(*i74a32fabbfbde9205732b623e943f248ae1993d1dd3bd7ca652fe67e4b8aaf82.ParentListRequestBuilder) {
    return i74a32fabbfbde9205732b623e943f248ae1993d1dd3bd7ca652fe67e4b8aaf82.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch all tasks in the users mailbox.
func (m *BaseTaskItemRequestBuilder) Patch(options *BaseTaskItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
