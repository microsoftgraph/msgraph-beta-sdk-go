package accesspackageresource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8296d4c043d5a60bb34fc8c6e864a861a1c50021f62866dfc96530575ca74cfd "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourceroles/item/accesspackageresource/accesspackageresourceroles"
    ia2b95ffb42ba3ef177aa3f3bd737ad90d3712852a8e6f8ebf3847227ef254dd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourceroles/item/accesspackageresource/accesspackageresourceenvironment"
    ic601d226f06525215c0b7289cb52776ba413ee7db0e1162ec58262128f2c0e63 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourceroles/item/accesspackageresource/accesspackageresourcescopes"
    iaa2d19b22298b622821f0ae9ec357d445070144bf4fa0b2717168860b96db23f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourceroles/item/accesspackageresource/accesspackageresourceroles/item"
    icd640ac26d82d7df7815b7a96e130ec2f3a860805729ed7884b80a071a9c1e5b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourceroles/item/accesspackageresource/accesspackageresourcescopes/item"
)

// accessPackageResourceRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentRequests\{accessPackageAssignmentRequest-id}\accessPackageAssignment\accessPackageAssignmentPolicy\accessPackageCatalog\accessPackageResourceRoles\{accessPackageResourceRole-id}\accessPackageResource
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
// accessPackageResourceRequestBuilderGetQueryParameters read-only. Nullable.
type AccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
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
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*ia2b95ffb42ba3ef177aa3f3bd737ad90d3712852a8e6f8ebf3847227ef254dd3.AccessPackageResourceEnvironmentRequestBuilder) {
    return ia2b95ffb42ba3ef177aa3f3bd737ad90d3712852a8e6f8ebf3847227ef254dd3.NewAccessPackageResourceEnvironmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*i8296d4c043d5a60bb34fc8c6e864a861a1c50021f62866dfc96530575ca74cfd.AccessPackageResourceRolesRequestBuilder) {
    return i8296d4c043d5a60bb34fc8c6e864a861a1c50021f62866dfc96530575ca74cfd.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item.accessPackageAssignment.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackageResourceRoles.item.accessPackageResource.accessPackageResourceRoles.item collection
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRolesById(id string)(*iaa2d19b22298b622821f0ae9ec357d445070144bf4fa0b2717168860b96db23f.AccessPackageResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole_id1"] = id
    }
    return iaa2d19b22298b622821f0ae9ec357d445070144bf4fa0b2717168860b96db23f.NewAccessPackageResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*ic601d226f06525215c0b7289cb52776ba413ee7db0e1162ec58262128f2c0e63.AccessPackageResourceScopesRequestBuilder) {
    return ic601d226f06525215c0b7289cb52776ba413ee7db0e1162ec58262128f2c0e63.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item.accessPackageAssignment.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackageResourceRoles.item.accessPackageResource.accessPackageResourceScopes.item collection
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopesById(id string)(*icd640ac26d82d7df7815b7a96e130ec2f3a860805729ed7884b80a071a9c1e5b.AccessPackageResourceScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope_id"] = id
    }
    return icd640ac26d82d7df7815b7a96e130ec2f3a860805729ed7884b80a071a9c1e5b.NewAccessPackageResourceScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageResourceRequestBuilderInternal instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    m := &AccessPackageResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/{accessPackageAssignmentRequest_id}/accessPackageAssignment/accessPackageAssignmentPolicy/accessPackageCatalog/accessPackageResourceRoles/{accessPackageResourceRole_id}/accessPackageResource{?select,expand}";
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
// CreateDeleteRequestInformation read-only. Nullable.
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
// CreateGetRequestInformation read-only. Nullable.
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
// CreatePatchRequestInformation read-only. Nullable.
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
// Delete read-only. Nullable.
func (m *AccessPackageResourceRequestBuilder) Delete(options *AccessPackageResourceRequestBuilderDeleteOptions)(error) {
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
// Get read-only. Nullable.
func (m *AccessPackageResourceRequestBuilder) Get(options *AccessPackageResourceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageResource, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageResource() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageResource), nil
}
// Patch read-only. Nullable.
func (m *AccessPackageResourceRequestBuilder) Patch(options *AccessPackageResourceRequestBuilderPatchOptions)(error) {
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
