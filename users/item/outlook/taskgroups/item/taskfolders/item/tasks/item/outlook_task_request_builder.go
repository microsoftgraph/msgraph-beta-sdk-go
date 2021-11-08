package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0718f2c7a8c775532496a98d185ae4cfcb0a71392db198afe0ebb5deddd3bfdd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/attachments"
    i0e3cd5b3643da0403d06ec109f55ebdd821409ca2b64e8929398963347543678 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/complete"
    i2488a196ac751025aa25b85acf5ef7ee5e6c2a37ea695ecd83a9367d0fbc5c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/singlevalueextendedproperties"
    iec0dcd96a66f87e472d58355850a2262b0d0a17eb7b7dc13ff86168db555b568 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/multivalueextendedproperties"
    i175900f9235a31a08f8f7870a3a233657ed75f27ea0f9c33a35029ae11dadd54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/attachments/item"
    ie3373c186191b55aac268824553eab67e16cbf405be52509be0da3cd61c6f601 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/multivalueextendedproperties/item"
    if8a609bdb585ca1f49d83f9d345d43bf56ed08b184a513e1f1cd22d928dad3ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/taskgroups/item/taskfolders/item/tasks/item/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \users\{user-id}\outlook\taskGroups\{outlookTaskGroup-id}\taskFolders\{outlookTaskFolder-id}\tasks\{outlookTask-id}
type OutlookTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type OutlookTaskRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// The tasks in this task folder. Read-only. Nullable.
type OutlookTaskRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
func (m *OutlookTaskRequestBuilder) Attachments()(*i0718f2c7a8c775532496a98d185ae4cfcb0a71392db198afe0ebb5deddd3bfdd.AttachmentsRequestBuilder) {
    return i0718f2c7a8c775532496a98d185ae4cfcb0a71392db198afe0ebb5deddd3bfdd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.users.item.outlook.taskGroups.item.taskFolders.item.tasks.item.attachments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskRequestBuilder) AttachmentsById(id string)(*i175900f9235a31a08f8f7870a3a233657ed75f27ea0f9c33a35029ae11dadd54.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i175900f9235a31a08f8f7870a3a233657ed75f27ea0f9c33a35029ae11dadd54.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) Complete()(*i0e3cd5b3643da0403d06ec109f55ebdd821409ca2b64e8929398963347543678.CompleteRequestBuilder) {
    return i0e3cd5b3643da0403d06ec109f55ebdd821409ca2b64e8929398963347543678.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new OutlookTaskRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOutlookTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    m := &OutlookTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook/taskGroups/{outlookTaskGroup_id}/taskFolders/{outlookTaskFolder_id}/tasks/{outlookTask_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OutlookTaskRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOutlookTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// The tasks in this task folder. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// The tasks in this task folder. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskRequestBuilder) CreateGetRequestInformation(options *OutlookTaskRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The tasks in this task folder. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// The tasks in this task folder. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskRequestBuilder) Delete(options *OutlookTaskRequestBuilderDeleteOptions)(error) {
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
// The tasks in this task folder. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskRequestBuilder) Get(options *OutlookTaskRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookTask() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTask), nil
}
func (m *OutlookTaskRequestBuilder) MultiValueExtendedProperties()(*iec0dcd96a66f87e472d58355850a2262b0d0a17eb7b7dc13ff86168db555b568.MultiValueExtendedPropertiesRequestBuilder) {
    return iec0dcd96a66f87e472d58355850a2262b0d0a17eb7b7dc13ff86168db555b568.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.users.item.outlook.taskGroups.item.taskFolders.item.tasks.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie3373c186191b55aac268824553eab67e16cbf405be52509be0da3cd61c6f601.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie3373c186191b55aac268824553eab67e16cbf405be52509be0da3cd61c6f601.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The tasks in this task folder. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskRequestBuilder) Patch(options *OutlookTaskRequestBuilderPatchOptions)(error) {
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
func (m *OutlookTaskRequestBuilder) SingleValueExtendedProperties()(*i2488a196ac751025aa25b85acf5ef7ee5e6c2a37ea695ecd83a9367d0fbc5c85.SingleValueExtendedPropertiesRequestBuilder) {
    return i2488a196ac751025aa25b85acf5ef7ee5e6c2a37ea695ecd83a9367d0fbc5c85.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.users.item.outlook.taskGroups.item.taskFolders.item.tasks.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if8a609bdb585ca1f49d83f9d345d43bf56ed08b184a513e1f1cd22d928dad3ed.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if8a609bdb585ca1f49d83f9d345d43bf56ed08b184a513e1f1cd22d928dad3ed.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
