package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2941b092d1da3f50cf26d460dabf5b8dd5821f17aa0a7343cc3136f5456a0457 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item/attachments"
    i50620e709c3f52cd004c28489a85e0f98be4689349900b99c634c678e91fa567 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item/complete"
    i62bc3d9952175eec6126b8681e65cc9ec963a13ac86445775f6888da660916e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item/multivalueextendedproperties"
    i6d43666759dfebc1fcc3061819b4d3409061c5bb53d59b28fe8f3e85c830ec36 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item/singlevalueextendedproperties"
    i0dafce107af728683d398b868f28c4197d94e0a866b6d2b610139604883c700d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item/attachments/item"
    i188db27f897af488ff96a3ecdb45af7f99b7f7b8deaa3c6c1cc2cb07a3cd6ea5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item/singlevalueextendedproperties/item"
    id99bf93f5256a2e6ace43ceeca6d34c8309be0d524ce8576473913eb69dc501d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item/multivalueextendedproperties/item"
)

// OutlookTaskRequestBuilder builds and executes requests for operations under \me\outlook\tasks\{outlookTask-id}
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
// OutlookTaskRequestBuilderGetQueryParameters get tasks from me
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
func (m *OutlookTaskRequestBuilder) Attachments()(*i2941b092d1da3f50cf26d460dabf5b8dd5821f17aa0a7343cc3136f5456a0457.AttachmentsRequestBuilder) {
    return i2941b092d1da3f50cf26d460dabf5b8dd5821f17aa0a7343cc3136f5456a0457.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.tasks.item.attachments.item collection
func (m *OutlookTaskRequestBuilder) AttachmentsById(id string)(*i0dafce107af728683d398b868f28c4197d94e0a866b6d2b610139604883c700d.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i0dafce107af728683d398b868f28c4197d94e0a866b6d2b610139604883c700d.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskRequestBuilder) Complete()(*i50620e709c3f52cd004c28489a85e0f98be4689349900b99c634c678e91fa567.CompleteRequestBuilder) {
    return i50620e709c3f52cd004c28489a85e0f98be4689349900b99c634c678e91fa567.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOutlookTaskRequestBuilderInternal instantiates a new OutlookTaskRequestBuilder and sets the default values.
func NewOutlookTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskRequestBuilder) {
    m := &OutlookTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook/tasks/{outlookTask_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property tasks for me
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
// CreateGetRequestInformation get tasks from me
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
// CreatePatchRequestInformation update the navigation property tasks in me
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
// Delete delete navigation property tasks for me
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
// Get get tasks from me
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
func (m *OutlookTaskRequestBuilder) MultiValueExtendedProperties()(*i62bc3d9952175eec6126b8681e65cc9ec963a13ac86445775f6888da660916e0.MultiValueExtendedPropertiesRequestBuilder) {
    return i62bc3d9952175eec6126b8681e65cc9ec963a13ac86445775f6888da660916e0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.tasks.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskRequestBuilder) MultiValueExtendedPropertiesById(id string)(*id99bf93f5256a2e6ace43ceeca6d34c8309be0d524ce8576473913eb69dc501d.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return id99bf93f5256a2e6ace43ceeca6d34c8309be0d524ce8576473913eb69dc501d.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property tasks in me
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
func (m *OutlookTaskRequestBuilder) SingleValueExtendedProperties()(*i6d43666759dfebc1fcc3061819b4d3409061c5bb53d59b28fe8f3e85c830ec36.SingleValueExtendedPropertiesRequestBuilder) {
    return i6d43666759dfebc1fcc3061819b4d3409061c5bb53d59b28fe8f3e85c830ec36.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.tasks.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i188db27f897af488ff96a3ecdb45af7f99b7f7b8deaa3c6c1cc2cb07a3cd6ea5.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i188db27f897af488ff96a3ecdb45af7f99b7f7b8deaa3c6c1cc2cb07a3cd6ea5.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
