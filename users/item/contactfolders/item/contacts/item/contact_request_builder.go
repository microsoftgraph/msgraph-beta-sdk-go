package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6f50d26878fd185dc0091888927cefa6d22be26e266781698fdd9e57d070de0c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item/photo"
    i89bb2a9305a78884b329be4830d7eb5f76a915b7db5307a4731bf60c5e8da8b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item/extensions"
    ie7b563dbeffa1f957d2c0512ba297dddec30375379bb001f6739318d0a52731b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item/multivalueextendedproperties"
    if0c6fb799e92d58a2210d46d5d2fcbcb99038ffc0ada12112b85ab35811dcb64 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item/singlevalueextendedproperties"
    i146c0887b9e7dd2662837f2ab05cee53be10c6a519941cb4f01e6044c8b94b2e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item/extensions/item"
    ib1e2ef514044b750e503c517726ae877cf24fced5b603805ddfbb2d6eae61348 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item/singlevalueextendedproperties/item"
    ib2a2411005f9939b0a566be28f8ffacea27e7e31b9875fdbe29afbf5a61a1d09 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item/multivalueextendedproperties/item"
)

// ContactRequestBuilder builds and executes requests for operations under \users\{user-id}\contactFolders\{contactFolder-id}\contacts\{contact-id}
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/contactFolders/{contactFolder_id}/contacts/{contact_id}{?select,expand}";
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
func (m *ContactRequestBuilder) Extensions()(*i89bb2a9305a78884b329be4830d7eb5f76a915b7db5307a4731bf60c5e8da8b8.ExtensionsRequestBuilder) {
    return i89bb2a9305a78884b329be4830d7eb5f76a915b7db5307a4731bf60c5e8da8b8.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contactFolders.item.contacts.item.extensions.item collection
func (m *ContactRequestBuilder) ExtensionsById(id string)(*i146c0887b9e7dd2662837f2ab05cee53be10c6a519941cb4f01e6044c8b94b2e.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i146c0887b9e7dd2662837f2ab05cee53be10c6a519941cb4f01e6044c8b94b2e.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ContactRequestBuilder) MultiValueExtendedProperties()(*ie7b563dbeffa1f957d2c0512ba297dddec30375379bb001f6739318d0a52731b.MultiValueExtendedPropertiesRequestBuilder) {
    return ie7b563dbeffa1f957d2c0512ba297dddec30375379bb001f6739318d0a52731b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contactFolders.item.contacts.item.multiValueExtendedProperties.item collection
func (m *ContactRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib2a2411005f9939b0a566be28f8ffacea27e7e31b9875fdbe29afbf5a61a1d09.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib2a2411005f9939b0a566be28f8ffacea27e7e31b9875fdbe29afbf5a61a1d09.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ContactRequestBuilder) Photo()(*i6f50d26878fd185dc0091888927cefa6d22be26e266781698fdd9e57d070de0c.PhotoRequestBuilder) {
    return i6f50d26878fd185dc0091888927cefa6d22be26e266781698fdd9e57d070de0c.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactRequestBuilder) SingleValueExtendedProperties()(*if0c6fb799e92d58a2210d46d5d2fcbcb99038ffc0ada12112b85ab35811dcb64.SingleValueExtendedPropertiesRequestBuilder) {
    return if0c6fb799e92d58a2210d46d5d2fcbcb99038ffc0ada12112b85ab35811dcb64.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contactFolders.item.contacts.item.singleValueExtendedProperties.item collection
func (m *ContactRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib1e2ef514044b750e503c517726ae877cf24fced5b603805ddfbb2d6eae61348.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib1e2ef514044b750e503c517726ae877cf24fced5b603805ddfbb2d6eae61348.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
