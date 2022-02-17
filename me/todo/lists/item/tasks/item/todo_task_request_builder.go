package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4b453bade987cdc7d78bf2dbba94d4dbc45a5e405d3622a58a4d26c09b49e08a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/linkedresources"
    ie7794ee0bc0caebd554d994f150a0b208c084c890b7636d17ca1336d19fff68c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/extensions"
    i6918a4356d3d8463417391beae7e08e9b5cea2610411228e7577257cbc0e08e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/extensions/item"
    i6ffb4522698e1ae7e78c00250910de6f003132aa52ac20af42871bdf610bd8fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item/linkedresources/item"
)

// TodoTaskRequestBuilder builds and executes requests for operations under \me\todo\lists\{todoTaskList-id}\tasks\{todoTask-id}
type TodoTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TodoTaskRequestBuilderDeleteOptions options for Delete
type TodoTaskRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TodoTaskRequestBuilderGetOptions options for Get
type TodoTaskRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TodoTaskRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TodoTaskRequestBuilderGetQueryParameters the tasks in this task list. Read-only. Nullable.
type TodoTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TodoTaskRequestBuilderPatchOptions options for Patch
type TodoTaskRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TodoTask;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTodoTaskRequestBuilderInternal instantiates a new TodoTaskRequestBuilder and sets the default values.
func NewTodoTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TodoTaskRequestBuilder) {
    m := &TodoTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/todo/lists/{todoTaskList_id}/tasks/{todoTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTodoTaskRequestBuilder instantiates a new TodoTaskRequestBuilder and sets the default values.
func NewTodoTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TodoTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTodoTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the tasks in this task list. Read-only. Nullable.
func (m *TodoTaskRequestBuilder) CreateDeleteRequestInformation(options *TodoTaskRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TodoTaskRequestBuilder) CreateGetRequestInformation(options *TodoTaskRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TodoTaskRequestBuilder) CreatePatchRequestInformation(options *TodoTaskRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TodoTaskRequestBuilder) Delete(options *TodoTaskRequestBuilderDeleteOptions)(error) {
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
func (m *TodoTaskRequestBuilder) Extensions()(*ie7794ee0bc0caebd554d994f150a0b208c084c890b7636d17ca1336d19fff68c.ExtensionsRequestBuilder) {
    return ie7794ee0bc0caebd554d994f150a0b208c084c890b7636d17ca1336d19fff68c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.tasks.item.extensions.item collection
func (m *TodoTaskRequestBuilder) ExtensionsById(id string)(*i6918a4356d3d8463417391beae7e08e9b5cea2610411228e7577257cbc0e08e8.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i6918a4356d3d8463417391beae7e08e9b5cea2610411228e7577257cbc0e08e8.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the tasks in this task list. Read-only. Nullable.
func (m *TodoTaskRequestBuilder) Get(options *TodoTaskRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TodoTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTodoTask() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TodoTask), nil
}
func (m *TodoTaskRequestBuilder) LinkedResources()(*i4b453bade987cdc7d78bf2dbba94d4dbc45a5e405d3622a58a4d26c09b49e08a.LinkedResourcesRequestBuilder) {
    return i4b453bade987cdc7d78bf2dbba94d4dbc45a5e405d3622a58a4d26c09b49e08a.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.tasks.item.linkedResources.item collection
func (m *TodoTaskRequestBuilder) LinkedResourcesById(id string)(*i6ffb4522698e1ae7e78c00250910de6f003132aa52ac20af42871bdf610bd8fa.LinkedResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_id"] = id
    }
    return i6ffb4522698e1ae7e78c00250910de6f003132aa52ac20af42871bdf610bd8fa.NewLinkedResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the tasks in this task list. Read-only. Nullable.
func (m *TodoTaskRequestBuilder) Patch(options *TodoTaskRequestBuilderPatchOptions)(error) {
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
