package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1ac26956b54b39830544e91adb7cfdf23c1dc6cee0646f106221409688641e4c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks"
    i1f646178fb04de83e8b723f8eaeaa22bc6b3c977b750cfe10aa98efedf9e4600 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/extensions"
    i489860ecb78ff309c4d67b8b0534a961907f9dfdf71f16d7fec53e2f13eb01d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/extensions/item"
    i6b6d487d9d6b2c29a8a4f8d7500f19acb13dc63261b2fed73efaa1426d376d45 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo/lists/item/tasks/item"
)

// TodoTaskListRequestBuilder builds and executes requests for operations under \me\todo\lists\{todoTaskList-id}
type TodoTaskListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TodoTaskListRequestBuilderDeleteOptions options for Delete
type TodoTaskListRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TodoTaskListRequestBuilderGetOptions options for Get
type TodoTaskListRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TodoTaskListRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TodoTaskListRequestBuilderGetQueryParameters the task lists in the users mailbox.
type TodoTaskListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TodoTaskListRequestBuilderPatchOptions options for Patch
type TodoTaskListRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TodoTaskList;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTodoTaskListRequestBuilderInternal instantiates a new TodoTaskListRequestBuilder and sets the default values.
func NewTodoTaskListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TodoTaskListRequestBuilder) {
    m := &TodoTaskListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/todo/lists/{todoTaskList_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTodoTaskListRequestBuilder instantiates a new TodoTaskListRequestBuilder and sets the default values.
func NewTodoTaskListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TodoTaskListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTodoTaskListRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the task lists in the users mailbox.
func (m *TodoTaskListRequestBuilder) CreateDeleteRequestInformation(options *TodoTaskListRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the task lists in the users mailbox.
func (m *TodoTaskListRequestBuilder) CreateGetRequestInformation(options *TodoTaskListRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the task lists in the users mailbox.
func (m *TodoTaskListRequestBuilder) CreatePatchRequestInformation(options *TodoTaskListRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the task lists in the users mailbox.
func (m *TodoTaskListRequestBuilder) Delete(options *TodoTaskListRequestBuilderDeleteOptions)(error) {
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
func (m *TodoTaskListRequestBuilder) Extensions()(*i1f646178fb04de83e8b723f8eaeaa22bc6b3c977b750cfe10aa98efedf9e4600.ExtensionsRequestBuilder) {
    return i1f646178fb04de83e8b723f8eaeaa22bc6b3c977b750cfe10aa98efedf9e4600.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.extensions.item collection
func (m *TodoTaskListRequestBuilder) ExtensionsById(id string)(*i489860ecb78ff309c4d67b8b0534a961907f9dfdf71f16d7fec53e2f13eb01d3.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i489860ecb78ff309c4d67b8b0534a961907f9dfdf71f16d7fec53e2f13eb01d3.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the task lists in the users mailbox.
func (m *TodoTaskListRequestBuilder) Get(options *TodoTaskListRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TodoTaskList, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTodoTaskList() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TodoTaskList), nil
}
// Patch the task lists in the users mailbox.
func (m *TodoTaskListRequestBuilder) Patch(options *TodoTaskListRequestBuilderPatchOptions)(error) {
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
func (m *TodoTaskListRequestBuilder) Tasks()(*i1ac26956b54b39830544e91adb7cfdf23c1dc6cee0646f106221409688641e4c.TasksRequestBuilder) {
    return i1ac26956b54b39830544e91adb7cfdf23c1dc6cee0646f106221409688641e4c.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.todo.lists.item.tasks.item collection
func (m *TodoTaskListRequestBuilder) TasksById(id string)(*i6b6d487d9d6b2c29a8a4f8d7500f19acb13dc63261b2fed73efaa1426d376d45.TodoTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["todoTask_id"] = id
    }
    return i6b6d487d9d6b2c29a8a4f8d7500f19acb13dc63261b2fed73efaa1426d376d45.NewTodoTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
