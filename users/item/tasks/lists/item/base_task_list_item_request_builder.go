package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i15170a666db7f763b065401d814ff76bf95a1f6e9180053455ccffbc92d1ee3a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/extensions"
    ib146adf1178ddf2b46634b441f3b2a7d4bc81f673a129cae90c66783acd3bc28 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks"
    i12318248cc169f54024ab11995a4248b4aa77343bbcd19ddb1d0bc7e075304e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/extensions/item"
    i5c5a5824dbc0bbb7f1e360329f6555426454d6bdc14cbac67bded4003d65880e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/lists/item/tasks/item"
)

// BaseTaskListItemRequestBuilder builds and executes requests for operations under \users\{user-id}\tasks\lists\{baseTaskList-id}
type BaseTaskListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BaseTaskListItemRequestBuilderDeleteOptions options for Delete
type BaseTaskListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTaskListItemRequestBuilderGetOptions options for Get
type BaseTaskListItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BaseTaskListItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTaskListItemRequestBuilderGetQueryParameters the task lists in the users mailbox.
type BaseTaskListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BaseTaskListItemRequestBuilderPatchOptions options for Patch
type BaseTaskListItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskList;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewBaseTaskListItemRequestBuilderInternal instantiates a new BaseTaskListItemRequestBuilder and sets the default values.
func NewBaseTaskListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskListItemRequestBuilder) {
    m := &BaseTaskListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/tasks/lists/{baseTaskList_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseTaskListItemRequestBuilder instantiates a new BaseTaskListItemRequestBuilder and sets the default values.
func NewBaseTaskListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseTaskListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the task lists in the users mailbox.
func (m *BaseTaskListItemRequestBuilder) CreateDeleteRequestInformation(options *BaseTaskListItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *BaseTaskListItemRequestBuilder) CreateGetRequestInformation(options *BaseTaskListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *BaseTaskListItemRequestBuilder) CreatePatchRequestInformation(options *BaseTaskListItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *BaseTaskListItemRequestBuilder) Delete(options *BaseTaskListItemRequestBuilderDeleteOptions)(error) {
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
func (m *BaseTaskListItemRequestBuilder) Extensions()(*i15170a666db7f763b065401d814ff76bf95a1f6e9180053455ccffbc92d1ee3a.ExtensionsRequestBuilder) {
    return i15170a666db7f763b065401d814ff76bf95a1f6e9180053455ccffbc92d1ee3a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.lists.item.extensions.item collection
func (m *BaseTaskListItemRequestBuilder) ExtensionsById(id string)(*i12318248cc169f54024ab11995a4248b4aa77343bbcd19ddb1d0bc7e075304e4.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i12318248cc169f54024ab11995a4248b4aa77343bbcd19ddb1d0bc7e075304e4.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the task lists in the users mailbox.
func (m *BaseTaskListItemRequestBuilder) Get(options *BaseTaskListItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskList, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBaseTaskList() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskList), nil
}
// Patch the task lists in the users mailbox.
func (m *BaseTaskListItemRequestBuilder) Patch(options *BaseTaskListItemRequestBuilderPatchOptions)(error) {
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
func (m *BaseTaskListItemRequestBuilder) Tasks()(*ib146adf1178ddf2b46634b441f3b2a7d4bc81f673a129cae90c66783acd3bc28.TasksRequestBuilder) {
    return ib146adf1178ddf2b46634b441f3b2a7d4bc81f673a129cae90c66783acd3bc28.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.lists.item.tasks.item collection
func (m *BaseTaskListItemRequestBuilder) TasksById(id string)(*i5c5a5824dbc0bbb7f1e360329f6555426454d6bdc14cbac67bded4003d65880e.BaseTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseTask_id"] = id
    }
    return i5c5a5824dbc0bbb7f1e360329f6555426454d6bdc14cbac67bded4003d65880e.NewBaseTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
