package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/singlevalueextendedproperties"
    ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/multivalueextendedproperties"
    id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders"
    if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts"
    i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/multivalueextendedproperties/item"
    i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/singlevalueextendedproperties/item"
    ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item"
    icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item"
)

// Builds and executes requests for operations under \users\{user-id}\contactFolders\{contactFolder-id}
type ContactFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ContactFolderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// The user's contacts folders. Read-only. Nullable.
type ContactFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
func (m *ContactFolderRequestBuilder) ChildFolders()(*id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7.ChildFoldersRequestBuilder) {
    return id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.users.item.contactFolders.item.childFolders.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) ChildFoldersById(id string)(*ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder_id1"] = id
    }
    return ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ContactFolderRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    m := &ContactFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/contactFolders/{contactFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ContactFolderRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContactFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContactFolderRequestBuilder) Contacts()(*if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd.ContactsRequestBuilder) {
    return if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.users.item.contactFolders.item.contacts.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) ContactsById(id string)(*icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) CreateGetRequestInformation(options *ContactFolderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) Delete(options *ContactFolderRequestBuilderDeleteOptions)(error) {
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
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) Get(options *ContactFolderRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContactFolder() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder), nil
}
func (m *ContactFolderRequestBuilder) MultiValueExtendedProperties()(*ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e.MultiValueExtendedPropertiesRequestBuilder) {
    return ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.users.item.contactFolders.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The user's contacts folders. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *ContactFolderRequestBuilder) Patch(options *ContactFolderRequestBuilderPatchOptions)(error) {
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
func (m *ContactFolderRequestBuilder) SingleValueExtendedProperties()(*i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6.SingleValueExtendedPropertiesRequestBuilder) {
    return i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.users.item.contactFolders.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContactFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
