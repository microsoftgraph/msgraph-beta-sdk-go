package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
// BaseTaskItemRequestBuilderGetQueryParameters the tasks in this task list. Read-only. Nullable.
type BaseTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BaseTaskItemRequestBuilderPatchOptions options for Patch
type BaseTaskItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *BaseTaskItemRequestBuilder) ChecklistItems()(*i7bfd138746038534bad1aac3dce035ec42cf5d70a13a1413d2f2e46892f38d09.ChecklistItemsRequestBuilder) {
    return i7bfd138746038534bad1aac3dce035ec42cf5d70a13a1413d2f2e46892f38d09.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.lists.item.tasks.item.checklistItems.item collection
func (m *BaseTaskItemRequestBuilder) ChecklistItemsById(id string)(*i6044c1d2eb01b0e0622e3b6c3849a33775aab29f024a15554528173596ac8f6b.ChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem_id"] = id
    }
    return i6044c1d2eb01b0e0622e3b6c3849a33775aab29f024a15554528173596ac8f6b.NewChecklistItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskItemRequestBuilderInternal instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    m := &BaseTaskItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/tasks/lists/{baseTaskList_id}/tasks/{baseTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseTaskItemRequestBuilder instantiates a new BaseTaskItemRequestBuilder and sets the default values.
func NewBaseTaskItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property tasks for me
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
// CreateGetRequestInformation the tasks in this task list. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property tasks in me
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
// Delete delete navigation property tasks for me
func (m *BaseTaskItemRequestBuilder) Delete(options *BaseTaskItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *BaseTaskItemRequestBuilder) Extensions()(*ia9d0c01958ba0799d3542155fd11213ed561dc63cc667c395497892e9e77f1d3.ExtensionsRequestBuilder) {
    return ia9d0c01958ba0799d3542155fd11213ed561dc63cc667c395497892e9e77f1d3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.lists.item.tasks.item.extensions.item collection
func (m *BaseTaskItemRequestBuilder) ExtensionsById(id string)(*i49f37fa0ca8dbd0779122380b5f0f7e39b7ff61cbc48e2b361a00f669235f391.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i49f37fa0ca8dbd0779122380b5f0f7e39b7ff61cbc48e2b361a00f669235f391.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskItemRequestBuilder) Get(options *BaseTaskItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateBaseTaskFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskable), nil
}
func (m *BaseTaskItemRequestBuilder) LinkedResources()(*i275fea45cd066f9262919593c03c0b1e1e7e71a727b5af64101dd64bc6e46b89.LinkedResourcesRequestBuilder) {
    return i275fea45cd066f9262919593c03c0b1e1e7e71a727b5af64101dd64bc6e46b89.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.lists.item.tasks.item.linkedResources.item collection
func (m *BaseTaskItemRequestBuilder) LinkedResourcesById(id string)(*i34e9ef0b7eef0d6704c833332ffa8d16c019e313b1def9f065f50d25a8f8068a.LinkedResource_v2ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_v2_id"] = id
    }
    return i34e9ef0b7eef0d6704c833332ffa8d16c019e313b1def9f065f50d25a8f8068a.NewLinkedResource_v2ItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BaseTaskItemRequestBuilder) Move()(*ic7aa8eb0e8aefbfa6858e7f2fc465b90d431c4e0f73bb00cf8b56f796639bf40.MoveRequestBuilder) {
    return ic7aa8eb0e8aefbfa6858e7f2fc465b90d431c4e0f73bb00cf8b56f796639bf40.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTaskItemRequestBuilder) ParentList()(*i7260b89120c013953545f8577655f38687e6049054459b4979fc66317449b7d4.ParentListRequestBuilder) {
    return i7260b89120c013953545f8577655f38687e6049054459b4979fc66317449b7d4.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property tasks in me
func (m *BaseTaskItemRequestBuilder) Patch(options *BaseTaskItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
