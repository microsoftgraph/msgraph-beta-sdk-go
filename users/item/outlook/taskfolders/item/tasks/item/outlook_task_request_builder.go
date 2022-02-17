package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i07134841e1432d33390c1388213ee45e6595d43d6e0eb485c3d75c5572f84b26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/singlevalueextendedproperties"
    i30d7001f23621738f5c0ebf4de402aa0940898dabde9c286e7bde7c8fe3a247a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/attachments"
    ibeff8c7ce77c8ec152f96297bac9743c951b7257a49b9309cd240df9aa287826 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/multivalueextendedproperties"
    ieeee0fbd5dbbdd405ef962f5cf1518d5b11621a62b9593c79e43559e54f18acc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/complete"
    i42d4fa154ff58194c41d7a068d2ef0ea649f669575467c50bac921280cece148 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/singlevalueextendedproperties/item"
    i8c19f39e583cd9a1b9f36e6c8fe3732bbb945cb45363ddd749683c9976564002 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/multivalueextendedproperties/item"
    ic6491e7b1ea237fd1a5fc05790050c780572b7f37e967ddffd4225e1dc62ce51 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskfolders/item/tasks/item/attachments/item"
)

// OutlookTaskRequestBuilder builds and executes requests for operations under \users\{user-id}\outlook\taskFolders\{outlookTaskFolder-id}\tasks\{outlookTask-id}
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
func (m *OutlookTaskRequestBuilder) Attachments()(*i30d7001f23621738f5c0ebf4de402aa0940898dabde9c286e7bde7c8fe3a247a.AttachmentsRequestBuilder) {
    return i30d7001f23621738f5c0ebf4de402aa0940898dabde9c286e7bde7c8fe3a247a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskFolders.item.tasks.item.attachments.item collection
func (m *OutlookTaskRequestBuilder) AttachmentsById(id string)(*ic6491e7b1ea237fd1a5fc05790050c780572b7f37e967ddffd4225e1dc62ce51.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ic6491e7b1ea237fd1a5fc05790050c780572b7f37e967ddffd4225e1dc62ce51.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) Complete()(*ieeee0fbd5dbbdd405ef962f5cf1518d5b11621a62b9593c79e43559e54f18acc.CompleteRequestBuilder) {
    return ieeee0fbd5dbbdd405ef962f5cf1518d5b11621a62b9593c79e43559e54f18acc.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOutlookTaskRequestBuilderInternal instantiates a new OutlookTaskRequestBuilder and sets the default values.
func NewOutlookTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    m := &OutlookTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook/taskFolders/{outlookTaskFolder_id}/tasks/{outlookTask_id}{?select}";
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
func (m *OutlookTaskRequestBuilder) MultiValueExtendedProperties()(*ibeff8c7ce77c8ec152f96297bac9743c951b7257a49b9309cd240df9aa287826.MultiValueExtendedPropertiesRequestBuilder) {
    return ibeff8c7ce77c8ec152f96297bac9743c951b7257a49b9309cd240df9aa287826.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskFolders.item.tasks.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i8c19f39e583cd9a1b9f36e6c8fe3732bbb945cb45363ddd749683c9976564002.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i8c19f39e583cd9a1b9f36e6c8fe3732bbb945cb45363ddd749683c9976564002.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *OutlookTaskRequestBuilder) SingleValueExtendedProperties()(*i07134841e1432d33390c1388213ee45e6595d43d6e0eb485c3d75c5572f84b26.SingleValueExtendedPropertiesRequestBuilder) {
    return i07134841e1432d33390c1388213ee45e6595d43d6e0eb485c3d75c5572f84b26.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.taskFolders.item.tasks.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i42d4fa154ff58194c41d7a068d2ef0ea649f669575467c50bac921280cece148.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i42d4fa154ff58194c41d7a068d2ef0ea649f669575467c50bac921280cece148.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
