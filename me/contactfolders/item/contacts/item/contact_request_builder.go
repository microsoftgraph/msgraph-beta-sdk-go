package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i52980e71a95c6997485ce47d4102d33335b4794977006887b239ae086c62b76f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts/item/extensions"
    ic02b356a36f84c8086585441504304808a8fde8f1d5dc63713c5f0ce93fe773a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts/item/photo"
    id135be1faebc3a476da0b46a5e60e4f8e1fe7ebe3e75aa38be52a41ecf079cd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts/item/multivalueextendedproperties"
    id821620ff4f9ab75705b5ca05b8636b0a4916f03566ffaed2287df45e9c77b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts/item/singlevalueextendedproperties"
    iaf512c028c7972dbb22a68b938957b64ed1aa9c89e8bfc77ed0aa39d7f9cce46 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts/item/multivalueextendedproperties/item"
    id26ee7cb435738682451f095aaaa69f8c9e03886d56355d147fef1acef738b9b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts/item/singlevalueextendedproperties/item"
    iffcd3d1e1cf26dabcd109b4514bbd3f3c5eeffd74b5b1db80bb25d65a776d3bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts/item/extensions/item"
)

// ContactRequestBuilder builds and executes requests for operations under \me\contactFolders\{contactFolder-id}\contacts\{contact-id}
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
// ContactRequestBuilderGetQueryParameters the contacts in the folder. Navigation property. Read-only. Nullable.
type ContactRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
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
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder_id}/contacts/{contact_id}{?select,expand}";
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
// CreateDeleteRequestInformation the contacts in the folder. Navigation property. Read-only. Nullable.
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
// CreateGetRequestInformation the contacts in the folder. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation the contacts in the folder. Navigation property. Read-only. Nullable.
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
// Delete the contacts in the folder. Navigation property. Read-only. Nullable.
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
func (m *ContactRequestBuilder) Extensions()(*i52980e71a95c6997485ce47d4102d33335b4794977006887b239ae086c62b76f.ExtensionsRequestBuilder) {
    return i52980e71a95c6997485ce47d4102d33335b4794977006887b239ae086c62b76f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.contacts.item.extensions.item collection
func (m *ContactRequestBuilder) ExtensionsById(id string)(*iffcd3d1e1cf26dabcd109b4514bbd3f3c5eeffd74b5b1db80bb25d65a776d3bd.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iffcd3d1e1cf26dabcd109b4514bbd3f3c5eeffd74b5b1db80bb25d65a776d3bd.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the contacts in the folder. Navigation property. Read-only. Nullable.
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
func (m *ContactRequestBuilder) MultiValueExtendedProperties()(*id135be1faebc3a476da0b46a5e60e4f8e1fe7ebe3e75aa38be52a41ecf079cd5.MultiValueExtendedPropertiesRequestBuilder) {
    return id135be1faebc3a476da0b46a5e60e4f8e1fe7ebe3e75aa38be52a41ecf079cd5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.contacts.item.multiValueExtendedProperties.item collection
func (m *ContactRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iaf512c028c7972dbb22a68b938957b64ed1aa9c89e8bfc77ed0aa39d7f9cce46.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iaf512c028c7972dbb22a68b938957b64ed1aa9c89e8bfc77ed0aa39d7f9cce46.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the contacts in the folder. Navigation property. Read-only. Nullable.
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
func (m *ContactRequestBuilder) Photo()(*ic02b356a36f84c8086585441504304808a8fde8f1d5dc63713c5f0ce93fe773a.PhotoRequestBuilder) {
    return ic02b356a36f84c8086585441504304808a8fde8f1d5dc63713c5f0ce93fe773a.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactRequestBuilder) SingleValueExtendedProperties()(*id821620ff4f9ab75705b5ca05b8636b0a4916f03566ffaed2287df45e9c77b4e.SingleValueExtendedPropertiesRequestBuilder) {
    return id821620ff4f9ab75705b5ca05b8636b0a4916f03566ffaed2287df45e9c77b4e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.contacts.item.singleValueExtendedProperties.item collection
func (m *ContactRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id26ee7cb435738682451f095aaaa69f8c9e03886d56355d147fef1acef738b9b.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return id26ee7cb435738682451f095aaaa69f8c9e03886d56355d147fef1acef738b9b.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
