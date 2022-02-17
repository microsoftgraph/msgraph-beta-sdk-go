package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1632b075081c4de18cdd5edc21cc2e7c9063df5f034bfbc388b91e7fd5d658b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts/item/photo"
    ia3c5754e7952201425ec172d10fc498d319c2035a4c55a190f620d0b4b01c485 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts/item/extensions"
    iae9fbb43cc098f4472a316a89e2740bc47e3ccb5ed104f1bdce799439ae5989f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts/item/multivalueextendedproperties"
    ieadc5611b4fe59f70e072dc8fe5da6f8e8138438cbd6b1e9318fe044a90a6f93 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts/item/singlevalueextendedproperties"
    i9b23336ac9fedb6dbfffaab293dfcc15d1cc50e7d10e9f75e9c8c6409c23f56d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts/item/multivalueextendedproperties/item"
    ib7c132abee6c3abbd5a110ac29557982e4780286ba5fd21013aa8ad33178a4d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts/item/singlevalueextendedproperties/item"
    ic33583d21393dd756d0bcc2a0b4d4b73ed3614112219fff63a50309df352674d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts/item/extensions/item"
)

// ContactRequestBuilder builds and executes requests for operations under \me\contacts\{contact-id}
type ContactRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactRequestBuilderDeleteOptions options for Delete
type ContactRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactRequestBuilderGetOptions options for Get
type ContactRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContactRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactRequestBuilderGetQueryParameters the user's contacts. Read-only. Nullable.
type ContactRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// ContactRequestBuilderPatchOptions options for Patch
type ContactRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contact;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewContactRequestBuilderInternal instantiates a new ContactRequestBuilder and sets the default values.
func NewContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactRequestBuilder) {
    m := &ContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contacts/{contact_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactRequestBuilder instantiates a new ContactRequestBuilder and sets the default values.
func NewContactRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the user's contacts. Read-only. Nullable.
func (m *ContactRequestBuilder) CreateDeleteRequestInformation(options *ContactRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the user's contacts. Read-only. Nullable.
func (m *ContactRequestBuilder) CreateGetRequestInformation(options *ContactRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the user's contacts. Read-only. Nullable.
func (m *ContactRequestBuilder) CreatePatchRequestInformation(options *ContactRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the user's contacts. Read-only. Nullable.
func (m *ContactRequestBuilder) Delete(options *ContactRequestBuilderDeleteOptions)(error) {
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
func (m *ContactRequestBuilder) Extensions()(*ia3c5754e7952201425ec172d10fc498d319c2035a4c55a190f620d0b4b01c485.ExtensionsRequestBuilder) {
    return ia3c5754e7952201425ec172d10fc498d319c2035a4c55a190f620d0b4b01c485.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contacts.item.extensions.item collection
func (m *ContactRequestBuilder) ExtensionsById(id string)(*ic33583d21393dd756d0bcc2a0b4d4b73ed3614112219fff63a50309df352674d.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic33583d21393dd756d0bcc2a0b4d4b73ed3614112219fff63a50309df352674d.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the user's contacts. Read-only. Nullable.
func (m *ContactRequestBuilder) Get(options *ContactRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contact, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContact() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contact), nil
}
func (m *ContactRequestBuilder) MultiValueExtendedProperties()(*iae9fbb43cc098f4472a316a89e2740bc47e3ccb5ed104f1bdce799439ae5989f.MultiValueExtendedPropertiesRequestBuilder) {
    return iae9fbb43cc098f4472a316a89e2740bc47e3ccb5ed104f1bdce799439ae5989f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contacts.item.multiValueExtendedProperties.item collection
func (m *ContactRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9b23336ac9fedb6dbfffaab293dfcc15d1cc50e7d10e9f75e9c8c6409c23f56d.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i9b23336ac9fedb6dbfffaab293dfcc15d1cc50e7d10e9f75e9c8c6409c23f56d.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's contacts. Read-only. Nullable.
func (m *ContactRequestBuilder) Patch(options *ContactRequestBuilderPatchOptions)(error) {
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
func (m *ContactRequestBuilder) Photo()(*i1632b075081c4de18cdd5edc21cc2e7c9063df5f034bfbc388b91e7fd5d658b2.PhotoRequestBuilder) {
    return i1632b075081c4de18cdd5edc21cc2e7c9063df5f034bfbc388b91e7fd5d658b2.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactRequestBuilder) SingleValueExtendedProperties()(*ieadc5611b4fe59f70e072dc8fe5da6f8e8138438cbd6b1e9318fe044a90a6f93.SingleValueExtendedPropertiesRequestBuilder) {
    return ieadc5611b4fe59f70e072dc8fe5da6f8e8138438cbd6b1e9318fe044a90a6f93.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contacts.item.singleValueExtendedProperties.item collection
func (m *ContactRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib7c132abee6c3abbd5a110ac29557982e4780286ba5fd21013aa8ad33178a4d0.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib7c132abee6c3abbd5a110ac29557982e4780286ba5fd21013aa8ad33178a4d0.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
