package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i34a7b474b81463b5a935d0c7abd8b14c118747cb2fb1cb87675c568c5b71b226 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item/singlevalueextendedproperties"
    i498203f2ea7d074a7874dfbedbae4da296a7aeed0fcef3b887a372a4bc41c094 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item/multivalueextendedproperties"
    i9d43990f76626f62d3520057fea9989c2b0fc0042707cb5da52862ddeec426d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item/attachments"
    ie004f5afb37d458d2e1011ba48d8618ebeb7cdb02e08f17f3f2ff1dc96ecc08b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item/complete"
    i444f8debaf035e0504d2a87ab804e4412ba23628297ccfb6dc62cb3ed324fcdb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item/attachments/item"
    ib4938b3e5bff9c974a9922cd155691a232752eca5e852fa2fe6c3091307d358c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item/multivalueextendedproperties/item"
    ie942dbe22794319c0c885d9bd37df549542f1e6ee4f5f30f284b1eee6536f5f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item/taskfolders/item/tasks/item/singlevalueextendedproperties/item"
)

// OutlookTaskRequestBuilder builds and executes requests for operations under \me\outlook\taskGroups\{outlookTaskGroup-id}\taskFolders\{outlookTaskFolder-id}\tasks\{outlookTask-id}
type OutlookTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OutlookTaskRequestBuilderDeleteOptions options for Delete
type OutlookTaskRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookTaskRequestBuilderGetOptions options for Get
type OutlookTaskRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OutlookTaskRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookTaskRequestBuilderGetQueryParameters the tasks in this task folder. Read-only. Nullable.
type OutlookTaskRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// OutlookTaskRequestBuilderPatchOptions options for Patch
type OutlookTaskRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OutlookTaskRequestBuilder) Attachments()(*i9d43990f76626f62d3520057fea9989c2b0fc0042707cb5da52862ddeec426d9.AttachmentsRequestBuilder) {
    return i9d43990f76626f62d3520057fea9989c2b0fc0042707cb5da52862ddeec426d9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskGroups.item.taskFolders.item.tasks.item.attachments.item collection
func (m *OutlookTaskRequestBuilder) AttachmentsById(id string)(*i444f8debaf035e0504d2a87ab804e4412ba23628297ccfb6dc62cb3ed324fcdb.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i444f8debaf035e0504d2a87ab804e4412ba23628297ccfb6dc62cb3ed324fcdb.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) Complete()(*ie004f5afb37d458d2e1011ba48d8618ebeb7cdb02e08f17f3f2ff1dc96ecc08b.CompleteRequestBuilder) {
    return ie004f5afb37d458d2e1011ba48d8618ebeb7cdb02e08f17f3f2ff1dc96ecc08b.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOutlookTaskRequestBuilderInternal instantiates a new OutlookTaskRequestBuilder and sets the default values.
func NewOutlookTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    m := &OutlookTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook/taskGroups/{outlookTaskGroup_id}/taskFolders/{outlookTaskFolder_id}/tasks/{outlookTask_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutlookTaskRequestBuilder instantiates a new OutlookTaskRequestBuilder and sets the default values.
func NewOutlookTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskRequestBuilder) CreateDeleteRequestInformation(options *OutlookTaskRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskRequestBuilder) CreateGetRequestInformation(options *OutlookTaskRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskRequestBuilder) CreatePatchRequestInformation(options *OutlookTaskRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskRequestBuilder) Delete(options *OutlookTaskRequestBuilderDeleteOptions)(error) {
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
// Get the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskRequestBuilder) Get(options *OutlookTaskRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookTask() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask), nil
}
func (m *OutlookTaskRequestBuilder) MultiValueExtendedProperties()(*i498203f2ea7d074a7874dfbedbae4da296a7aeed0fcef3b887a372a4bc41c094.MultiValueExtendedPropertiesRequestBuilder) {
    return i498203f2ea7d074a7874dfbedbae4da296a7aeed0fcef3b887a372a4bc41c094.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskGroups.item.taskFolders.item.tasks.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib4938b3e5bff9c974a9922cd155691a232752eca5e852fa2fe6c3091307d358c.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib4938b3e5bff9c974a9922cd155691a232752eca5e852fa2fe6c3091307d358c.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskRequestBuilder) Patch(options *OutlookTaskRequestBuilderPatchOptions)(error) {
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
func (m *OutlookTaskRequestBuilder) SingleValueExtendedProperties()(*i34a7b474b81463b5a935d0c7abd8b14c118747cb2fb1cb87675c568c5b71b226.SingleValueExtendedPropertiesRequestBuilder) {
    return i34a7b474b81463b5a935d0c7abd8b14c118747cb2fb1cb87675c568c5b71b226.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskGroups.item.taskFolders.item.tasks.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie942dbe22794319c0c885d9bd37df549542f1e6ee4f5f30f284b1eee6536f5f6.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ie942dbe22794319c0c885d9bd37df549542f1e6ee4f5f30f284b1eee6536f5f6.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
