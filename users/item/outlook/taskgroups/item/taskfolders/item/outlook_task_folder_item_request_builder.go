package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i9232dc9a3eb1378fa8ab78540bf13f11f1ab9949e5170e28c8c489942e74ac83 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/multivalueextendedproperties"
    ia2c6de18fe93e8e7f9292ac15424fb781c2838b95b4e00f25a8262ffa9865608 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks"
    ib439086047d429f844b496ff8c963c45f26f72f62f1b200dbad602723f92e0dd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/singlevalueextendedproperties"
    i81872ba447d8d450bb84c0c0c18419a62d7689176babe2d34cd58db372f9485d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item"
    i89bc63c8513b2df831f97375088aa608a991f1de6b732207a66f40dd22b4e15a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/multivalueextendedproperties/item"
    i8dfde376fcdc847c160e38c1bb7577866731487084a179701aa2f8b0cfc4db62 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/singlevalueextendedproperties/item"
)

// OutlookTaskFolderItemRequestBuilder provides operations to manage the taskFolders property of the microsoft.graph.outlookTaskGroup entity.
type OutlookTaskFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OutlookTaskFolderItemRequestBuilderDeleteOptions options for Delete
type OutlookTaskFolderItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookTaskFolderItemRequestBuilderGetOptions options for Get
type OutlookTaskFolderItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OutlookTaskFolderItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookTaskFolderItemRequestBuilderGetQueryParameters the collection of task folders in the task group. Read-only. Nullable.
type OutlookTaskFolderItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// OutlookTaskFolderItemRequestBuilderPatchOptions options for Patch
type OutlookTaskFolderItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolderable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOutlookTaskFolderItemRequestBuilderInternal instantiates a new OutlookTaskFolderItemRequestBuilder and sets the default values.
func NewOutlookTaskFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderItemRequestBuilder) {
    m := &OutlookTaskFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook/taskGroups/{outlookTaskGroup_id}/taskFolders/{outlookTaskFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutlookTaskFolderItemRequestBuilder instantiates a new OutlookTaskFolderItemRequestBuilder and sets the default values.
func NewOutlookTaskFolderItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property taskFolders for users
func (m *OutlookTaskFolderItemRequestBuilder) CreateDeleteRequestInformation(options *OutlookTaskFolderItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskFolderItemRequestBuilder) CreateGetRequestInformation(options *OutlookTaskFolderItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property taskFolders in users
func (m *OutlookTaskFolderItemRequestBuilder) CreatePatchRequestInformation(options *OutlookTaskFolderItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property taskFolders for users
func (m *OutlookTaskFolderItemRequestBuilder) Delete(options *OutlookTaskFolderItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskFolderItemRequestBuilder) Get(options *OutlookTaskFolderItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOutlookTaskFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolderable), nil
}
func (m *OutlookTaskFolderItemRequestBuilder) MultiValueExtendedProperties()(*i9232dc9a3eb1378fa8ab78540bf13f11f1ab9949e5170e28c8c489942e74ac83.MultiValueExtendedPropertiesRequestBuilder) {
    return i9232dc9a3eb1378fa8ab78540bf13f11f1ab9949e5170e28c8c489942e74ac83.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i89bc63c8513b2df831f97375088aa608a991f1de6b732207a66f40dd22b4e15a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i89bc63c8513b2df831f97375088aa608a991f1de6b732207a66f40dd22b4e15a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property taskFolders in users
func (m *OutlookTaskFolderItemRequestBuilder) Patch(options *OutlookTaskFolderItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *OutlookTaskFolderItemRequestBuilder) SingleValueExtendedProperties()(*ib439086047d429f844b496ff8c963c45f26f72f62f1b200dbad602723f92e0dd.SingleValueExtendedPropertiesRequestBuilder) {
    return ib439086047d429f844b496ff8c963c45f26f72f62f1b200dbad602723f92e0dd.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8dfde376fcdc847c160e38c1bb7577866731487084a179701aa2f8b0cfc4db62.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i8dfde376fcdc847c160e38c1bb7577866731487084a179701aa2f8b0cfc4db62.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskFolderItemRequestBuilder) Tasks()(*ia2c6de18fe93e8e7f9292ac15424fb781c2838b95b4e00f25a8262ffa9865608.TasksRequestBuilder) {
    return ia2c6de18fe93e8e7f9292ac15424fb781c2838b95b4e00f25a8262ffa9865608.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskGroups.item.taskFolders.item.tasks.item collection
func (m *OutlookTaskFolderItemRequestBuilder) TasksById(id string)(*i81872ba447d8d450bb84c0c0c18419a62d7689176babe2d34cd58db372f9485d.OutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return i81872ba447d8d450bb84c0c0c18419a62d7689176babe2d34cd58db372f9485d.NewOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
