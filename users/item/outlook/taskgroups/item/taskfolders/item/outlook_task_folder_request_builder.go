package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9232dc9a3eb1378fa8ab78540bf13f11f1ab9949e5170e28c8c489942e74ac83 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/multivalueextendedproperties"
    ia2c6de18fe93e8e7f9292ac15424fb781c2838b95b4e00f25a8262ffa9865608 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks"
    ib439086047d429f844b496ff8c963c45f26f72f62f1b200dbad602723f92e0dd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/singlevalueextendedproperties"
    i81872ba447d8d450bb84c0c0c18419a62d7689176babe2d34cd58db372f9485d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item"
    i89bc63c8513b2df831f97375088aa608a991f1de6b732207a66f40dd22b4e15a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/multivalueextendedproperties/item"
    i8dfde376fcdc847c160e38c1bb7577866731487084a179701aa2f8b0cfc4db62 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \users\{user-id}\outlook\taskGroups\{outlookTaskGroup-id}\taskFolders\{outlookTaskFolder-id}
type OutlookTaskFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type OutlookTaskFolderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type OutlookTaskFolderRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OutlookTaskFolderRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The collection of task folders in the task group. Read-only. Nullable.
type OutlookTaskFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type OutlookTaskFolderRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new OutlookTaskFolderRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOutlookTaskFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderRequestBuilder) {
    m := &OutlookTaskFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook/taskGroups/{outlookTaskGroup_id}/taskFolders/{outlookTaskFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OutlookTaskFolderRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOutlookTaskFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// The collection of task folders in the task group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) CreateDeleteRequestInformation(options *OutlookTaskFolderRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The collection of task folders in the task group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) CreateGetRequestInformation(options *OutlookTaskFolderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The collection of task folders in the task group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) CreatePatchRequestInformation(options *OutlookTaskFolderRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The collection of task folders in the task group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) Delete(options *OutlookTaskFolderRequestBuilderDeleteOptions)(error) {
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
// The collection of task folders in the task group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) Get(options *OutlookTaskFolderRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookTaskFolder() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder), nil
}
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedProperties()(*i9232dc9a3eb1378fa8ab78540bf13f11f1ab9949e5170e28c8c489942e74ac83.MultiValueExtendedPropertiesRequestBuilder) {
    return i9232dc9a3eb1378fa8ab78540bf13f11f1ab9949e5170e28c8c489942e74ac83.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i89bc63c8513b2df831f97375088aa608a991f1de6b732207a66f40dd22b4e15a.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i89bc63c8513b2df831f97375088aa608a991f1de6b732207a66f40dd22b4e15a.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The collection of task folders in the task group. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) Patch(options *OutlookTaskFolderRequestBuilderPatchOptions)(error) {
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
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedProperties()(*ib439086047d429f844b496ff8c963c45f26f72f62f1b200dbad602723f92e0dd.SingleValueExtendedPropertiesRequestBuilder) {
    return ib439086047d429f844b496ff8c963c45f26f72f62f1b200dbad602723f92e0dd.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8dfde376fcdc847c160e38c1bb7577866731487084a179701aa2f8b0cfc4db62.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i8dfde376fcdc847c160e38c1bb7577866731487084a179701aa2f8b0cfc4db62.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) Tasks()(*ia2c6de18fe93e8e7f9292ac15424fb781c2838b95b4e00f25a8262ffa9865608.TasksRequestBuilder) {
    return ia2c6de18fe93e8e7f9292ac15424fb781c2838b95b4e00f25a8262ffa9865608.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.tasks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskFolderRequestBuilder) TasksById(id string)(*i81872ba447d8d450bb84c0c0c18419a62d7689176babe2d34cd58db372f9485d.OutlookTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return i81872ba447d8d450bb84c0c0c18419a62d7689176babe2d34cd58db372f9485d.NewOutlookTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
