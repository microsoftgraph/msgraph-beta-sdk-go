package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6c56256bfadf73ebbba11c1f3c521a9bf2c51d2ce66b674522da1efed2c76164 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/singlevalueextendedproperties"
    iba7aa0b6c0d2610ed857ef28336339eb91eaad62aaf05abfbeed1aca7f3de9ae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/multivalueextendedproperties"
    ida74707faa099d980fd35eab88d0504497b748f800d0980a6e3e4d7fb5536c1f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks"
    i5d6c8aecacb64bc593c78078fa8da8bad018bb85a6e98cbb99d79d9c4e381210 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/multivalueextendedproperties/item"
    ia214c090cdeb5788b48c6cd877b5fe213b14f72374a78702e4950bc827828a83 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/singlevalueextendedproperties/item"
    id767174df22d17dacf803edd5f352b8b2e6ce0d941323ea2dd0f699c8c65ba4f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item"
)

// OutlookTaskFolderRequestBuilder builds and executes requests for operations under \me\outlook\taskGroups\{outlookTaskGroup-id}\taskFolders\{outlookTaskFolder-id}
type OutlookTaskFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OutlookTaskFolderRequestBuilderDeleteOptions options for Delete
type OutlookTaskFolderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookTaskFolderRequestBuilderGetOptions options for Get
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
// OutlookTaskFolderRequestBuilderGetQueryParameters the collection of task folders in the task group. Read-only. Nullable.
type OutlookTaskFolderRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// OutlookTaskFolderRequestBuilderPatchOptions options for Patch
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
// NewOutlookTaskFolderRequestBuilderInternal instantiates a new OutlookTaskFolderRequestBuilder and sets the default values.
func NewOutlookTaskFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderRequestBuilder) {
    m := &OutlookTaskFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook/taskGroups/{outlookTaskGroup_id}/taskFolders/{outlookTaskFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutlookTaskFolderRequestBuilder instantiates a new OutlookTaskFolderRequestBuilder and sets the default values.
func NewOutlookTaskFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the collection of task folders in the task group. Read-only. Nullable.
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
// CreateGetRequestInformation the collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskFolderRequestBuilder) CreateGetRequestInformation(options *OutlookTaskFolderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the collection of task folders in the task group. Read-only. Nullable.
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
// Delete the collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskFolderRequestBuilder) Delete(options *OutlookTaskFolderRequestBuilderDeleteOptions)(error) {
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
// Get the collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskFolderRequestBuilder) Get(options *OutlookTaskFolderRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookTaskFolder() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder), nil
}
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedProperties()(*iba7aa0b6c0d2610ed857ef28336339eb91eaad62aaf05abfbeed1aca7f3de9ae.MultiValueExtendedPropertiesRequestBuilder) {
    return iba7aa0b6c0d2610ed857ef28336339eb91eaad62aaf05abfbeed1aca7f3de9ae.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskGroups.item.taskFolders.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5d6c8aecacb64bc593c78078fa8da8bad018bb85a6e98cbb99d79d9c4e381210.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5d6c8aecacb64bc593c78078fa8da8bad018bb85a6e98cbb99d79d9c4e381210.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the collection of task folders in the task group. Read-only. Nullable.
func (m *OutlookTaskFolderRequestBuilder) Patch(options *OutlookTaskFolderRequestBuilderPatchOptions)(error) {
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
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedProperties()(*i6c56256bfadf73ebbba11c1f3c521a9bf2c51d2ce66b674522da1efed2c76164.SingleValueExtendedPropertiesRequestBuilder) {
    return i6c56256bfadf73ebbba11c1f3c521a9bf2c51d2ce66b674522da1efed2c76164.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskGroups.item.taskFolders.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia214c090cdeb5788b48c6cd877b5fe213b14f72374a78702e4950bc827828a83.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ia214c090cdeb5788b48c6cd877b5fe213b14f72374a78702e4950bc827828a83.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) Tasks()(*ida74707faa099d980fd35eab88d0504497b748f800d0980a6e3e4d7fb5536c1f.TasksRequestBuilder) {
    return ida74707faa099d980fd35eab88d0504497b748f800d0980a6e3e4d7fb5536c1f.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskGroups.item.taskFolders.item.tasks.item collection
func (m *OutlookTaskFolderRequestBuilder) TasksById(id string)(*id767174df22d17dacf803edd5f352b8b2e6ce0d941323ea2dd0f699c8c65ba4f.OutlookTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return id767174df22d17dacf803edd5f352b8b2e6ce0d941323ea2dd0f699c8c65ba4f.NewOutlookTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
