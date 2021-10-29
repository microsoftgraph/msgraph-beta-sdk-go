package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i16212364c1dea56c6119bb5dd239204de69b1c17f96ec621220b9a5614b38bbb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources/item/accesspackageresourceroles"
    i5d262bbbcb43d2b9e1b400baffc9f3a7bb17c73f075a158a7cfe37762aadcaaa "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources/item/accesspackageresourceenvironment"
    ic9285d9a7d6082203e2fbfa2808f33349d166bd92e0e743aeb8ed1d8dd9c8cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources/item/accesspackageresourcescopes"
    i174deb9de981cafe7f9ddccd961b2d261a0bb7bcff06a081c1a3f170c2458a38 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources/item/accesspackageresourcescopes/item"
    i330811327782e70957ad4d3cd9a9f2b9e4097eb63c95c4e677e87ea79cc76b06 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources/item/accesspackageresourceroles/item"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\accessPackageAssignmentPolicies\{accessPackageAssignmentPolicy-id}\accessPackageCatalog\accessPackageResources\{accessPackageResource-id}
type AccessPackageResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AccessPackageResourceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// Read-only. Nullable.
type AccessPackageResourceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*i5d262bbbcb43d2b9e1b400baffc9f3a7bb17c73f075a158a7cfe37762aadcaaa.AccessPackageResourceEnvironmentRequestBuilder) {
    return i5d262bbbcb43d2b9e1b400baffc9f3a7bb17c73f075a158a7cfe37762aadcaaa.NewAccessPackageResourceEnvironmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*i16212364c1dea56c6119bb5dd239204de69b1c17f96ec621220b9a5614b38bbb.AccessPackageResourceRolesRequestBuilder) {
    return i16212364c1dea56c6119bb5dd239204de69b1c17f96ec621220b9a5614b38bbb.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackage.accessPackageAssignmentPolicies.item.accessPackageCatalog.accessPackageResources.item.accessPackageResourceRoles.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRolesById(id string)(*i330811327782e70957ad4d3cd9a9f2b9e4097eb63c95c4e677e87ea79cc76b06.AccessPackageResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole_id"] = id
    }
    return i330811327782e70957ad4d3cd9a9f2b9e4097eb63c95c4e677e87ea79cc76b06.NewAccessPackageResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*ic9285d9a7d6082203e2fbfa2808f33349d166bd92e0e743aeb8ed1d8dd9c8cc4.AccessPackageResourceScopesRequestBuilder) {
    return ic9285d9a7d6082203e2fbfa2808f33349d166bd92e0e743aeb8ed1d8dd9c8cc4.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackage.accessPackageAssignmentPolicies.item.accessPackageCatalog.accessPackageResources.item.accessPackageResourceScopes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopesById(id string)(*i174deb9de981cafe7f9ddccd961b2d261a0bb7bcff06a081c1a3f170c2458a38.AccessPackageResourceScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope_id"] = id
    }
    return i174deb9de981cafe7f9ddccd961b2d261a0bb7bcff06a081c1a3f170c2458a38.NewAccessPackageResourceScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    m := &AccessPackageResourceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}/accessPackageAssignments/{accessPackageAssignment_id}/accessPackage/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy_id}/accessPackageCatalog/accessPackageResources/{accessPackageResource_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageResourceRequestBuilder) CreateGetRequestInformation(options *AccessPackageResourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
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
