package parentlist

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i92da41fea3da5d7717eb0093ced514fbd82f90ad48e1fd94d7eeedb63c207e29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/parentlist/tasks"
    ida9b39dc6cd04d5bf4195eaacbdb7bf00f1ffd3c3c02c2dc7f72b63b1dc06267 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/parentlist/extensions"
    i079d8c9aea62bc833b74baaa08031d0566d72d72e29ad60e90bb055794cc6033 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/parentlist/tasks/item"
    i72d816769d471cd143ef29e1e1c317ab444402e0363fc2ec57a2546cad901cef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks/alltasks/item/parentlist/extensions/item"
)

// ParentListRequestBuilder builds and executes requests for operations under \users\{user-id}\tasks\alltasks\{baseTask-id}\parentList
type ParentListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParentListRequestBuilderDeleteOptions options for Delete
type ParentListRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentListRequestBuilderGetOptions options for Get
type ParentListRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentListRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentListRequestBuilderGetQueryParameters the list which contains the task.
type ParentListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ParentListRequestBuilderPatchOptions options for Patch
type ParentListRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskList;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewParentListRequestBuilderInternal instantiates a new ParentListRequestBuilder and sets the default values.
func NewParentListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentListRequestBuilder) {
    m := &ParentListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/tasks/alltasks/{baseTask_id}/parentList{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParentListRequestBuilder instantiates a new ParentListRequestBuilder and sets the default values.
func NewParentListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentListRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list which contains the task.
func (m *ParentListRequestBuilder) CreateDeleteRequestInformation(options *ParentListRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list which contains the task.
func (m *ParentListRequestBuilder) CreateGetRequestInformation(options *ParentListRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list which contains the task.
func (m *ParentListRequestBuilder) CreatePatchRequestInformation(options *ParentListRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list which contains the task.
func (m *ParentListRequestBuilder) Delete(options *ParentListRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ParentListRequestBuilder) Extensions()(*ida9b39dc6cd04d5bf4195eaacbdb7bf00f1ffd3c3c02c2dc7f72b63b1dc06267.ExtensionsRequestBuilder) {
    return ida9b39dc6cd04d5bf4195eaacbdb7bf00f1ffd3c3c02c2dc7f72b63b1dc06267.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.alltasks.item.parentList.extensions.item collection
func (m *ParentListRequestBuilder) ExtensionsById(id string)(*i72d816769d471cd143ef29e1e1c317ab444402e0363fc2ec57a2546cad901cef.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i72d816769d471cd143ef29e1e1c317ab444402e0363fc2ec57a2546cad901cef.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list which contains the task.
func (m *ParentListRequestBuilder) Get(options *ParentListRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskList, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBaseTaskList() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTaskList), nil
}
// Patch the list which contains the task.
func (m *ParentListRequestBuilder) Patch(options *ParentListRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ParentListRequestBuilder) Tasks()(*i92da41fea3da5d7717eb0093ced514fbd82f90ad48e1fd94d7eeedb63c207e29.TasksRequestBuilder) {
    return i92da41fea3da5d7717eb0093ced514fbd82f90ad48e1fd94d7eeedb63c207e29.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.tasks.alltasks.item.parentList.tasks.item collection
func (m *ParentListRequestBuilder) TasksById(id string)(*i079d8c9aea62bc833b74baaa08031d0566d72d72e29ad60e90bb055794cc6033.BaseTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseTask_id1"] = id
    }
    return i079d8c9aea62bc833b74baaa08031d0566d72d72e29ad60e90bb055794cc6033.NewBaseTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
