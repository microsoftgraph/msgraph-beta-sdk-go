package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8eb6bf0d18006185acc8705e9517cd7f3058b3858cb6b19eaadda0b9488faac8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/singlevalueextendedproperties"
    idccc52741431a4fb690c45f0baefc5521653f1e2e0aad764d9c00094d0afa348 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/multivalueextendedproperties"
    ief425dbd2220862119be08f8de6f5dab765b8c5937368a6c9bed2e16564b73b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts"
    i68130945dcdf75efc2a83f8f4751c565bf7cd1538effbb02d575bd0403ba736b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/contacts/item"
    i93f2db7c39c0ef162bdbf28ffa71f251301cac645a0f820be27fcf05d619c3ee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/singlevalueextendedproperties/item"
    ia97727f0c654d03b46e0952beed1930ebf4fb90131c3546d8112cac09314057b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item/multivalueextendedproperties/item"
)

// ContactFolderItemRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.contactFolder entity.
type ContactFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactFolderItemRequestBuilderDeleteOptions options for Delete
type ContactFolderItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactFolderItemRequestBuilderGetOptions options for Get
type ContactFolderItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContactFolderItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactFolderItemRequestBuilderGetQueryParameters the collection of child folders in the folder. Navigation property. Read-only. Nullable.
type ContactFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContactFolderItemRequestBuilderPatchOptions options for Patch
type ContactFolderItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolderable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewContactFolderItemRequestBuilderInternal instantiates a new ContactFolderItemRequestBuilder and sets the default values.
func NewContactFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderItemRequestBuilder) {
    m := &ContactFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/contactFolders/{contactFolder_id}/childFolders/{contactFolder_id1}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactFolderItemRequestBuilder instantiates a new ContactFolderItemRequestBuilder and sets the default values.
func NewContactFolderItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContactFolderItemRequestBuilder) Contacts()(*ief425dbd2220862119be08f8de6f5dab765b8c5937368a6c9bed2e16564b73b6.ContactsRequestBuilder) {
    return ief425dbd2220862119be08f8de6f5dab765b8c5937368a6c9bed2e16564b73b6.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contactFolders.item.childFolders.item.contacts.item collection
func (m *ContactFolderItemRequestBuilder) ContactsById(id string)(*i68130945dcdf75efc2a83f8f4751c565bf7cd1538effbb02d575bd0403ba736b.ContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return i68130945dcdf75efc2a83f8f4751c565bf7cd1538effbb02d575bd0403ba736b.NewContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property childFolders for users
func (m *ContactFolderItemRequestBuilder) CreateDeleteRequestInformation(options *ContactFolderItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) CreateGetRequestInformation(options *ContactFolderItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property childFolders in users
func (m *ContactFolderItemRequestBuilder) CreatePatchRequestInformation(options *ContactFolderItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property childFolders for users
func (m *ContactFolderItemRequestBuilder) Delete(options *ContactFolderItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of child folders in the folder. Navigation property. Read-only. Nullable.
func (m *ContactFolderItemRequestBuilder) Get(options *ContactFolderItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContactFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolderable), nil
}
func (m *ContactFolderItemRequestBuilder) MultiValueExtendedProperties()(*idccc52741431a4fb690c45f0baefc5521653f1e2e0aad764d9c00094d0afa348.MultiValueExtendedPropertiesRequestBuilder) {
    return idccc52741431a4fb690c45f0baefc5521653f1e2e0aad764d9c00094d0afa348.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contactFolders.item.childFolders.item.multiValueExtendedProperties.item collection
func (m *ContactFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ia97727f0c654d03b46e0952beed1930ebf4fb90131c3546d8112cac09314057b.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ia97727f0c654d03b46e0952beed1930ebf4fb90131c3546d8112cac09314057b.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property childFolders in users
func (m *ContactFolderItemRequestBuilder) Patch(options *ContactFolderItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContactFolderItemRequestBuilder) SingleValueExtendedProperties()(*i8eb6bf0d18006185acc8705e9517cd7f3058b3858cb6b19eaadda0b9488faac8.SingleValueExtendedPropertiesRequestBuilder) {
    return i8eb6bf0d18006185acc8705e9517cd7f3058b3858cb6b19eaadda0b9488faac8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contactFolders.item.childFolders.item.singleValueExtendedProperties.item collection
func (m *ContactFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i93f2db7c39c0ef162bdbf28ffa71f251301cac645a0f820be27fcf05d619c3ee.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i93f2db7c39c0ef162bdbf28ffa71f251301cac645a0f820be27fcf05d619c3ee.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
