package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// BaseTaskRequestBuilder builds and executes requests for operations under \me\tasks\lists\{baseTaskList-id}\tasks\{baseTask-id}
type BaseTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BaseTaskRequestBuilderDeleteOptions options for Delete
type BaseTaskRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTaskRequestBuilderGetOptions options for Get
type BaseTaskRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BaseTaskRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTaskRequestBuilderGetQueryParameters the tasks in this task list. Read-only. Nullable.
type BaseTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BaseTaskRequestBuilderPatchOptions options for Patch
type BaseTaskRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTask;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *BaseTaskRequestBuilder) ChecklistItems()(*i7bfd138746038534bad1aac3dce035ec42cf5d70a13a1413d2f2e46892f38d09.ChecklistItemsRequestBuilder) {
    return i7bfd138746038534bad1aac3dce035ec42cf5d70a13a1413d2f2e46892f38d09.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.lists.item.tasks.item.checklistItems.item collection
func (m *BaseTaskRequestBuilder) ChecklistItemsById(id string)(*i6044c1d2eb01b0e0622e3b6c3849a33775aab29f024a15554528173596ac8f6b.ChecklistItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem_id"] = id
    }
    return i6044c1d2eb01b0e0622e3b6c3849a33775aab29f024a15554528173596ac8f6b.NewChecklistItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskRequestBuilderInternal instantiates a new BaseTaskRequestBuilder and sets the default values.
func NewBaseTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskRequestBuilder) {
    m := &BaseTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/tasks/lists/{baseTaskList_id}/tasks/{baseTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseTaskRequestBuilder instantiates a new BaseTaskRequestBuilder and sets the default values.
func NewBaseTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskRequestBuilder) CreateDeleteRequestInformation(options *BaseTaskRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *BaseTaskRequestBuilder) CreateGetRequestInformation(options *BaseTaskRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskRequestBuilder) CreatePatchRequestInformation(options *BaseTaskRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskRequestBuilder) Delete(options *BaseTaskRequestBuilderDeleteOptions)(error) {
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
func (m *BaseTaskRequestBuilder) Extensions()(*ia9d0c01958ba0799d3542155fd11213ed561dc63cc667c395497892e9e77f1d3.ExtensionsRequestBuilder) {
    return ia9d0c01958ba0799d3542155fd11213ed561dc63cc667c395497892e9e77f1d3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.lists.item.tasks.item.extensions.item collection
func (m *BaseTaskRequestBuilder) ExtensionsById(id string)(*i49f37fa0ca8dbd0779122380b5f0f7e39b7ff61cbc48e2b361a00f669235f391.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i49f37fa0ca8dbd0779122380b5f0f7e39b7ff61cbc48e2b361a00f669235f391.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskRequestBuilder) Get(options *BaseTaskRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTask, error) {
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
func (m *BaseTaskRequestBuilder) LinkedResources()(*i275fea45cd066f9262919593c03c0b1e1e7e71a727b5af64101dd64bc6e46b89.LinkedResourcesRequestBuilder) {
    return i275fea45cd066f9262919593c03c0b1e1e7e71a727b5af64101dd64bc6e46b89.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.lists.item.tasks.item.linkedResources.item collection
func (m *BaseTaskRequestBuilder) LinkedResourcesById(id string)(*i34e9ef0b7eef0d6704c833332ffa8d16c019e313b1def9f065f50d25a8f8068a.LinkedResource_v2RequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_v2_id"] = id
    }
    return i34e9ef0b7eef0d6704c833332ffa8d16c019e313b1def9f065f50d25a8f8068a.NewLinkedResource_v2RequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BaseTaskRequestBuilder) Move()(*ic7aa8eb0e8aefbfa6858e7f2fc465b90d431c4e0f73bb00cf8b56f796639bf40.MoveRequestBuilder) {
    return ic7aa8eb0e8aefbfa6858e7f2fc465b90d431c4e0f73bb00cf8b56f796639bf40.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTaskRequestBuilder) ParentList()(*i7260b89120c013953545f8577655f38687e6049054459b4979fc66317449b7d4.ParentListRequestBuilder) {
    return i7260b89120c013953545f8577655f38687e6049054459b4979fc66317449b7d4.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the tasks in this task list. Read-only. Nullable.
func (m *BaseTaskRequestBuilder) Patch(options *BaseTaskRequestBuilderPatchOptions)(error) {
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
