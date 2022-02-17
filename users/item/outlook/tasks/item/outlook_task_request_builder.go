package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4b9d4654b2d0be177a461cbf7bc3120b7d72197137ebdad70610309e863012de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item/multivalueextendedproperties"
    i5f9147d736494d2f3513c196dabc85442955ac41d807005c7a3f2575c751bfb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item/attachments"
    i6421a31b795ecc630a1fd0f306d37a2948ceb176d58fee3cacbb70bdda3be196 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item/singlevalueextendedproperties"
    i877daa06d4ce14c4194ae0939ce1a6fcb72c37e83096ec8d4e05bcd37e6108db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item/complete"
    i217dd05f0714bded4927301b5d1219654d224978a3d8c839f6c12da84106b4db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item/attachments/item"
    i512f019a5c0b6347bfc14ade805130177db816ece605df9b57db754d6e5b7816 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item/singlevalueextendedproperties/item"
    i682e26936704da1fcbc018d796323c7a7bfa541b0cba81f5d501f2e178dabe65 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook/tasks/item/multivalueextendedproperties/item"
)

// OutlookTaskRequestBuilder builds and executes requests for operations under \users\{user-id}\outlook\tasks\{outlookTask-id}
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
// OutlookTaskRequestBuilderGetQueryParameters get tasks from users
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
func (m *OutlookTaskRequestBuilder) Attachments()(*i5f9147d736494d2f3513c196dabc85442955ac41d807005c7a3f2575c751bfb7.AttachmentsRequestBuilder) {
    return i5f9147d736494d2f3513c196dabc85442955ac41d807005c7a3f2575c751bfb7.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.tasks.item.attachments.item collection
func (m *OutlookTaskRequestBuilder) AttachmentsById(id string)(*i217dd05f0714bded4927301b5d1219654d224978a3d8c839f6c12da84106b4db.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i217dd05f0714bded4927301b5d1219654d224978a3d8c839f6c12da84106b4db.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) Complete()(*i877daa06d4ce14c4194ae0939ce1a6fcb72c37e83096ec8d4e05bcd37e6108db.CompleteRequestBuilder) {
    return i877daa06d4ce14c4194ae0939ce1a6fcb72c37e83096ec8d4e05bcd37e6108db.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOutlookTaskRequestBuilderInternal instantiates a new OutlookTaskRequestBuilder and sets the default values.
func NewOutlookTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    m := &OutlookTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook/tasks/{outlookTask_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property tasks for users
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
// CreateGetRequestInformation get tasks from users
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
// CreatePatchRequestInformation update the navigation property tasks in users
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
// Delete delete navigation property tasks for users
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
// Get get tasks from users
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
func (m *OutlookTaskRequestBuilder) MultiValueExtendedProperties()(*i4b9d4654b2d0be177a461cbf7bc3120b7d72197137ebdad70610309e863012de.MultiValueExtendedPropertiesRequestBuilder) {
    return i4b9d4654b2d0be177a461cbf7bc3120b7d72197137ebdad70610309e863012de.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.tasks.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i682e26936704da1fcbc018d796323c7a7bfa541b0cba81f5d501f2e178dabe65.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i682e26936704da1fcbc018d796323c7a7bfa541b0cba81f5d501f2e178dabe65.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property tasks in users
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
func (m *OutlookTaskRequestBuilder) SingleValueExtendedProperties()(*i6421a31b795ecc630a1fd0f306d37a2948ceb176d58fee3cacbb70bdda3be196.SingleValueExtendedPropertiesRequestBuilder) {
    return i6421a31b795ecc630a1fd0f306d37a2948ceb176d58fee3cacbb70bdda3be196.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.outlook.tasks.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i512f019a5c0b6347bfc14ade805130177db816ece605df9b57db754d6e5b7816.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i512f019a5c0b6347bfc14ade805130177db816ece605df9b57db754d6e5b7816.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
