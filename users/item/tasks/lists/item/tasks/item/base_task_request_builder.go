package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i63585e4f9db5ce04afe272d39ba9e0f814d15a9ad69564f035b57992fdc60d0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/linkedresources"
    i7eb8a811f65d6bab93f6143e18dd050750b3187b5222bba2363767dea08f400e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/move"
    i9f90a854cb36ba230b51f79e03cf8c485003d9c20892f3aaaf9f9ca6e2311052 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/checklistitems"
    ic3517bb76f041a0e9a13acef1cef85787191fc80023465b6551e0527d3bb7dcf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/parentlist"
    if2aa69e4c699585c66a9474fde03947bf3cd63c88d2a207d476ff72016bb92ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/extensions"
    i400d2cff755de6f49886cdc0854849d644b5a10323083dc756fb70061fd9b716 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/extensions/item"
    i43baf14bdfa52348b0bb48c3d3ece5a20da9f14a5275a095129997e3eaebe805 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/checklistitems/item"
    ic96ef0290e7cd1ba8655c202ac321602c963f0f41deb8e52900dee755247715e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item/linkedresources/item"
)

// BaseTaskRequestBuilder builds and executes requests for operations under \users\{user-id}\tasks\lists\{baseTaskList-id}\tasks\{baseTask-id}
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
func (m *BaseTaskRequestBuilder) ChecklistItems()(*i9f90a854cb36ba230b51f79e03cf8c485003d9c20892f3aaaf9f9ca6e2311052.ChecklistItemsRequestBuilder) {
    return i9f90a854cb36ba230b51f79e03cf8c485003d9c20892f3aaaf9f9ca6e2311052.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.lists.item.tasks.item.checklistItems.item collection
func (m *BaseTaskRequestBuilder) ChecklistItemsById(id string)(*i43baf14bdfa52348b0bb48c3d3ece5a20da9f14a5275a095129997e3eaebe805.ChecklistItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem_id"] = id
    }
    return i43baf14bdfa52348b0bb48c3d3ece5a20da9f14a5275a095129997e3eaebe805.NewChecklistItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskRequestBuilderInternal instantiates a new BaseTaskRequestBuilder and sets the default values.
func NewBaseTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskRequestBuilder) {
    m := &BaseTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/tasks/lists/{baseTaskList_id}/tasks/{baseTask_id}{?select,expand}";
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
func (m *BaseTaskRequestBuilder) Extensions()(*if2aa69e4c699585c66a9474fde03947bf3cd63c88d2a207d476ff72016bb92ed.ExtensionsRequestBuilder) {
    return if2aa69e4c699585c66a9474fde03947bf3cd63c88d2a207d476ff72016bb92ed.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.lists.item.tasks.item.extensions.item collection
func (m *BaseTaskRequestBuilder) ExtensionsById(id string)(*i400d2cff755de6f49886cdc0854849d644b5a10323083dc756fb70061fd9b716.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i400d2cff755de6f49886cdc0854849d644b5a10323083dc756fb70061fd9b716.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *BaseTaskRequestBuilder) LinkedResources()(*i63585e4f9db5ce04afe272d39ba9e0f814d15a9ad69564f035b57992fdc60d0e.LinkedResourcesRequestBuilder) {
    return i63585e4f9db5ce04afe272d39ba9e0f814d15a9ad69564f035b57992fdc60d0e.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.lists.item.tasks.item.linkedResources.item collection
func (m *BaseTaskRequestBuilder) LinkedResourcesById(id string)(*ic96ef0290e7cd1ba8655c202ac321602c963f0f41deb8e52900dee755247715e.LinkedResource_v2RequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_v2_id"] = id
    }
    return ic96ef0290e7cd1ba8655c202ac321602c963f0f41deb8e52900dee755247715e.NewLinkedResource_v2RequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BaseTaskRequestBuilder) Move()(*i7eb8a811f65d6bab93f6143e18dd050750b3187b5222bba2363767dea08f400e.MoveRequestBuilder) {
    return i7eb8a811f65d6bab93f6143e18dd050750b3187b5222bba2363767dea08f400e.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTaskRequestBuilder) ParentList()(*ic3517bb76f041a0e9a13acef1cef85787191fc80023465b6551e0527d3bb7dcf.ParentListRequestBuilder) {
    return ic3517bb76f041a0e9a13acef1cef85787191fc80023465b6551e0527d3bb7dcf.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
