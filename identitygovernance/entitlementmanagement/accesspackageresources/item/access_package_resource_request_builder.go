package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i387fe96c7d96fcb8613570b9322185d08fc785b11085836ec36d6c5b2995c7dd "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresources/item/accesspackageresourceroles"
    i62985a1c8685af456ee76558f9b1f934adb62015f17f6211f96c5aeae4a498fa "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresources/item/accesspackageresourcescopes"
    i67c48ce950164a1a9454de27c6454ee7048e9c15fdf4cdf6fb367bbafae09d09 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresources/item/accesspackageresourceenvironment"
    i1781f468d4b78f1cdf416e6bb87cd4d58004a003d38f2c564b28e09ecb238c33 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresources/item/accesspackageresourcescopes/item"
    id4e464e5928895f8007864f311a44cf0a17d078df1fd85477aa8a89b4982d3a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresources/item/accesspackageresourceroles/item"
)

// AccessPackageResourceRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageResources\{accessPackageResource-id}
type AccessPackageResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageResourceRequestBuilderDeleteOptions options for Delete
type AccessPackageResourceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageResourceRequestBuilderGetOptions options for Get
type AccessPackageResourceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageResourceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageResourceRequestBuilderGetQueryParameters a reference to a resource associated with an access package catalog.
type AccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageResourceRequestBuilderPatchOptions options for Patch
type AccessPackageResourceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageResource;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*i67c48ce950164a1a9454de27c6454ee7048e9c15fdf4cdf6fb367bbafae09d09.AccessPackageResourceEnvironmentRequestBuilder) {
    return i67c48ce950164a1a9454de27c6454ee7048e9c15fdf4cdf6fb367bbafae09d09.NewAccessPackageResourceEnvironmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*i387fe96c7d96fcb8613570b9322185d08fc785b11085836ec36d6c5b2995c7dd.AccessPackageResourceRolesRequestBuilder) {
    return i387fe96c7d96fcb8613570b9322185d08fc785b11085836ec36d6c5b2995c7dd.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResources.item.accessPackageResourceRoles.item collection
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRolesById(id string)(*id4e464e5928895f8007864f311a44cf0a17d078df1fd85477aa8a89b4982d3a5.AccessPackageResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole_id"] = id
    }
    return id4e464e5928895f8007864f311a44cf0a17d078df1fd85477aa8a89b4982d3a5.NewAccessPackageResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*i62985a1c8685af456ee76558f9b1f934adb62015f17f6211f96c5aeae4a498fa.AccessPackageResourceScopesRequestBuilder) {
    return i62985a1c8685af456ee76558f9b1f934adb62015f17f6211f96c5aeae4a498fa.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResources.item.accessPackageResourceScopes.item collection
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopesById(id string)(*i1781f468d4b78f1cdf416e6bb87cd4d58004a003d38f2c564b28e09ecb238c33.AccessPackageResourceScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope_id"] = id
    }
    return i1781f468d4b78f1cdf416e6bb87cd4d58004a003d38f2c564b28e09ecb238c33.NewAccessPackageResourceScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageResourceRequestBuilderInternal instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    m := &AccessPackageResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResources/{accessPackageResource_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageResourceRequestBuilder instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a reference to a resource associated with an access package catalog.
func (m *AccessPackageResourceRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageResourceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a reference to a resource associated with an access package catalog.
func (m *AccessPackageResourceRequestBuilder) CreateGetRequestInformation(options *AccessPackageResourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a reference to a resource associated with an access package catalog.
func (m *AccessPackageResourceRequestBuilder) CreatePatchRequestInformation(options *AccessPackageResourceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a reference to a resource associated with an access package catalog.
func (m *AccessPackageResourceRequestBuilder) Delete(options *AccessPackageResourceRequestBuilderDeleteOptions)(error) {
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
// Get a reference to a resource associated with an access package catalog.
func (m *AccessPackageResourceRequestBuilder) Get(options *AccessPackageResourceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageResource, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageResource() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageResource), nil
}
// Patch a reference to a resource associated with an access package catalog.
func (m *AccessPackageResourceRequestBuilder) Patch(options *AccessPackageResourceRequestBuilderPatchOptions)(error) {
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
