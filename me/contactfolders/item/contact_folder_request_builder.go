package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0bd690d9d85efbda6faca0e423d291baec2293fff52018cbce884dfc92dac3b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/multivalueextendedproperties"
    i6d2115b7dcfd0e4fae6be95d8385d21b635d0ba40344a7e074e24f205edc4e6d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/singlevalueextendedproperties"
    i6e6c11ddaba3735902a0749fc9a96c5bf6fd2336d1b9b5dc988603d3008aa911 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts"
    ia19bf58f393474881f984c54f8f265ef6abe094f707d04b486a0de9f8734db29 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/childfolders"
    i83a77b7e0a5e0641ad16e76e439386e06c6d7348d963e863a4321878f9cc6e1f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/multivalueextendedproperties/item"
    i9711c9898edb32536801206e2e6cfd7381cb9160f0fa2612326197fe9ddb7763 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/childfolders/item"
    iae3f6d356bd9c88fe93e11b66cccca7469ebee9f3bb71e062fb3a078da46a965 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/singlevalueextendedproperties/item"
    ic5c395d29801c827fcd63e433f534d10d09684514c44971ebd9d9776f28dac7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item/contacts/item"
)

// ContactFolderRequestBuilder builds and executes requests for operations under \me\contactFolders\{contactFolder-id}
type ContactFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactFolderRequestBuilderDeleteOptions options for Delete
type ContactFolderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactFolderRequestBuilderGetOptions options for Get
type ContactFolderRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContactFolderRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactFolderRequestBuilderGetQueryParameters the user's contacts folders. Read-only. Nullable.
type ContactFolderRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// ContactFolderRequestBuilderPatchOptions options for Patch
type ContactFolderRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContactFolderRequestBuilder) ChildFolders()(*ia19bf58f393474881f984c54f8f265ef6abe094f707d04b486a0de9f8734db29.ChildFoldersRequestBuilder) {
    return ia19bf58f393474881f984c54f8f265ef6abe094f707d04b486a0de9f8734db29.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.childFolders.item collection
func (m *ContactFolderRequestBuilder) ChildFoldersById(id string)(*i9711c9898edb32536801206e2e6cfd7381cb9160f0fa2612326197fe9ddb7763.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder_id1"] = id
    }
    return i9711c9898edb32536801206e2e6cfd7381cb9160f0fa2612326197fe9ddb7763.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContactFolderRequestBuilderInternal instantiates a new ContactFolderRequestBuilder and sets the default values.
func NewContactFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    m := &ContactFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/contactFolders/{contactFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactFolderRequestBuilder instantiates a new ContactFolderRequestBuilder and sets the default values.
func NewContactFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContactFolderRequestBuilder) Contacts()(*i6e6c11ddaba3735902a0749fc9a96c5bf6fd2336d1b9b5dc988603d3008aa911.ContactsRequestBuilder) {
    return i6e6c11ddaba3735902a0749fc9a96c5bf6fd2336d1b9b5dc988603d3008aa911.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.contacts.item collection
func (m *ContactFolderRequestBuilder) ContactsById(id string)(*ic5c395d29801c827fcd63e433f534d10d09684514c44971ebd9d9776f28dac7e.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return ic5c395d29801c827fcd63e433f534d10d09684514c44971ebd9d9776f28dac7e.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation the user's contacts folders. Read-only. Nullable.
func (m *ContactFolderRequestBuilder) CreateDeleteRequestInformation(options *ContactFolderRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the user's contacts folders. Read-only. Nullable.
func (m *ContactFolderRequestBuilder) CreateGetRequestInformation(options *ContactFolderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the user's contacts folders. Read-only. Nullable.
func (m *ContactFolderRequestBuilder) CreatePatchRequestInformation(options *ContactFolderRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the user's contacts folders. Read-only. Nullable.
func (m *ContactFolderRequestBuilder) Delete(options *ContactFolderRequestBuilderDeleteOptions)(error) {
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
// Get the user's contacts folders. Read-only. Nullable.
func (m *ContactFolderRequestBuilder) Get(options *ContactFolderRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContactFolder() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder), nil
}
func (m *ContactFolderRequestBuilder) MultiValueExtendedProperties()(*i0bd690d9d85efbda6faca0e423d291baec2293fff52018cbce884dfc92dac3b1.MultiValueExtendedPropertiesRequestBuilder) {
    return i0bd690d9d85efbda6faca0e423d291baec2293fff52018cbce884dfc92dac3b1.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.multiValueExtendedProperties.item collection
func (m *ContactFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i83a77b7e0a5e0641ad16e76e439386e06c6d7348d963e863a4321878f9cc6e1f.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i83a77b7e0a5e0641ad16e76e439386e06c6d7348d963e863a4321878f9cc6e1f.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's contacts folders. Read-only. Nullable.
func (m *ContactFolderRequestBuilder) Patch(options *ContactFolderRequestBuilderPatchOptions)(error) {
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
func (m *ContactFolderRequestBuilder) SingleValueExtendedProperties()(*i6d2115b7dcfd0e4fae6be95d8385d21b635d0ba40344a7e074e24f205edc4e6d.SingleValueExtendedPropertiesRequestBuilder) {
    return i6d2115b7dcfd0e4fae6be95d8385d21b635d0ba40344a7e074e24f205edc4e6d.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item.singleValueExtendedProperties.item collection
func (m *ContactFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iae3f6d356bd9c88fe93e11b66cccca7469ebee9f3bb71e062fb3a078da46a965.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return iae3f6d356bd9c88fe93e11b66cccca7469ebee9f3bb71e062fb3a078da46a965.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
