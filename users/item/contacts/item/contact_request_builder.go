package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i32630bf5a40120b413c9183bae333e6054a59c736a62597a7a53fdc7bdd321fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item/multivalueextendedproperties"
    i3fef6dd05b3fc212eff69f402926aec45502db994ed5342e766f3e568323e591 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item/singlevalueextendedproperties"
    i46c2fa37031933cdf768d58a91bf34aedb2ac9a66cfaccfd4c59c530260a509b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item/photo"
    ie3d9f38317d56bf10dda6a36e3760ff01c63ef2ad752e3c69e3cec7076f38854 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item/extensions"
    i2701a9bf3922181a04af924cf3fbecae3d77d44a3405910fbbc6aed642912d7d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item/extensions/item"
    i4fc59cfed5e5f3b9672457001a35df6b3a62d567de18acefa1ca1cc5823c1c4c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item/singlevalueextendedproperties/item"
    i7e5eb71905e81193c6dacd18d4b084ea5663e00b4ae0e2ca050049b100d89d56 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item/multivalueextendedproperties/item"
)

// ContactRequestBuilder builds and executes requests for operations under \users\{user-id}\contacts\{contact-id}
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/contacts/{contact_id}{?select}";
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
func (m *ContactRequestBuilder) Extensions()(*ie3d9f38317d56bf10dda6a36e3760ff01c63ef2ad752e3c69e3cec7076f38854.ExtensionsRequestBuilder) {
    return ie3d9f38317d56bf10dda6a36e3760ff01c63ef2ad752e3c69e3cec7076f38854.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contacts.item.extensions.item collection
func (m *ContactRequestBuilder) ExtensionsById(id string)(*i2701a9bf3922181a04af924cf3fbecae3d77d44a3405910fbbc6aed642912d7d.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i2701a9bf3922181a04af924cf3fbecae3d77d44a3405910fbbc6aed642912d7d.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ContactRequestBuilder) MultiValueExtendedProperties()(*i32630bf5a40120b413c9183bae333e6054a59c736a62597a7a53fdc7bdd321fc.MultiValueExtendedPropertiesRequestBuilder) {
    return i32630bf5a40120b413c9183bae333e6054a59c736a62597a7a53fdc7bdd321fc.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contacts.item.multiValueExtendedProperties.item collection
func (m *ContactRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7e5eb71905e81193c6dacd18d4b084ea5663e00b4ae0e2ca050049b100d89d56.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i7e5eb71905e81193c6dacd18d4b084ea5663e00b4ae0e2ca050049b100d89d56.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *ContactRequestBuilder) Photo()(*i46c2fa37031933cdf768d58a91bf34aedb2ac9a66cfaccfd4c59c530260a509b.PhotoRequestBuilder) {
    return i46c2fa37031933cdf768d58a91bf34aedb2ac9a66cfaccfd4c59c530260a509b.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactRequestBuilder) SingleValueExtendedProperties()(*i3fef6dd05b3fc212eff69f402926aec45502db994ed5342e766f3e568323e591.SingleValueExtendedPropertiesRequestBuilder) {
    return i3fef6dd05b3fc212eff69f402926aec45502db994ed5342e766f3e568323e591.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contacts.item.singleValueExtendedProperties.item collection
func (m *ContactRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4fc59cfed5e5f3b9672457001a35df6b3a62d567de18acefa1ca1cc5823c1c4c.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i4fc59cfed5e5f3b9672457001a35df6b3a62d567de18acefa1ca1cc5823c1c4c.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
