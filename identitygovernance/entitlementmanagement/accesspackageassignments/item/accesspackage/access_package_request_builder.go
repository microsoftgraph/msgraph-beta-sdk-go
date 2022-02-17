package accesspackage

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5026331d2a7bdc9d2c2dd6b2ef26b15ecd7a2fa0d7b6e5e51ee3edaaa5b0cc56 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackagesincompatiblewith"
    i66ab1fc6d4d7116b6184cb8766c3a625bdb710add38a93aba66d42256f89d4be "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackagecatalog"
    i68afe5073e23e4326ed3ec5de58b6805edabc8bb1fc63c11b706f82db3a754eb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/getapplicablepolicyrequirements"
    ib91d9df9d429e0fc35844f475381fd503a15a677bacc1be75db14e120d0c4890 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/incompatiblegroups"
    ic5e7fe471d3697546614564c096841925c880120f29fd2c97f7dc52498716347 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/incompatibleaccesspackages"
    id9e10e3e999f301963019a47665bf8203b99c979a74ebc2075deb09c06bd6550 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageresourcerolescopes"
    iffc8da12dbe75797135b0d9f3888b812dbfdc0eb62a4d878fefa2f91d9ee2faa "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies"
    i1ea1714ed88aff5bab6f57f9269b25a4120b7c0c8abfc486a21600548154257f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageresourcerolescopes/item"
    i1fd5eff195b33e63418a6cfc726dff089cd1b3c93946c7718ce2c44ed8c48c3d "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item"
)

// AccessPackageRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage
type AccessPackageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageRequestBuilderDeleteOptions options for Delete
type AccessPackageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageRequestBuilderGetOptions options for Get
type AccessPackageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageRequestBuilderGetQueryParameters read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
type AccessPackageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageRequestBuilderPatchOptions options for Patch
type AccessPackageRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPolicies()(*iffc8da12dbe75797135b0d9f3888b812dbfdc0eb62a4d878fefa2f91d9ee2faa.AccessPackageAssignmentPoliciesRequestBuilder) {
    return iffc8da12dbe75797135b0d9f3888b812dbfdc0eb62a4d878fefa2f91d9ee2faa.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.accessPackageAssignmentPolicies.item collection
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i1fd5eff195b33e63418a6cfc726dff089cd1b3c93946c7718ce2c44ed8c48c3d.AccessPackageAssignmentPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy_id"] = id
    }
    return i1fd5eff195b33e63418a6cfc726dff089cd1b3c93946c7718ce2c44ed8c48c3d.NewAccessPackageAssignmentPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageCatalog()(*i66ab1fc6d4d7116b6184cb8766c3a625bdb710add38a93aba66d42256f89d4be.AccessPackageCatalogRequestBuilder) {
    return i66ab1fc6d4d7116b6184cb8766c3a625bdb710add38a93aba66d42256f89d4be.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopes()(*id9e10e3e999f301963019a47665bf8203b99c979a74ebc2075deb09c06bd6550.AccessPackageResourceRoleScopesRequestBuilder) {
    return id9e10e3e999f301963019a47665bf8203b99c979a74ebc2075deb09c06bd6550.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.accessPackageResourceRoleScopes.item collection
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i1ea1714ed88aff5bab6f57f9269b25a4120b7c0c8abfc486a21600548154257f.AccessPackageResourceRoleScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope_id"] = id
    }
    return i1ea1714ed88aff5bab6f57f9269b25a4120b7c0c8abfc486a21600548154257f.NewAccessPackageResourceRoleScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackagesIncompatibleWith()(*i5026331d2a7bdc9d2c2dd6b2ef26b15ecd7a2fa0d7b6e5e51ee3edaaa5b0cc56.AccessPackagesIncompatibleWithRequestBuilder) {
    return i5026331d2a7bdc9d2c2dd6b2ef26b15ecd7a2fa0d7b6e5e51ee3edaaa5b0cc56.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageRequestBuilderInternal instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    m := &AccessPackageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackage{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageRequestBuilder instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewAccessPackageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageRequestBuilder) CreateGetRequestInformation(options *AccessPackageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageRequestBuilder) CreatePatchRequestInformation(options *AccessPackageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageRequestBuilder) Delete(options *AccessPackageRequestBuilderDeleteOptions)(error) {
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
// Get read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageRequestBuilder) Get(options *AccessPackageRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackage() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage), nil
}
func (m *AccessPackageRequestBuilder) GetApplicablePolicyRequirements()(*i68afe5073e23e4326ed3ec5de58b6805edabc8bb1fc63c11b706f82db3a754eb.GetApplicablePolicyRequirementsRequestBuilder) {
    return i68afe5073e23e4326ed3ec5de58b6805edabc8bb1fc63c11b706f82db3a754eb.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleAccessPackages()(*ic5e7fe471d3697546614564c096841925c880120f29fd2c97f7dc52498716347.IncompatibleAccessPackagesRequestBuilder) {
    return ic5e7fe471d3697546614564c096841925c880120f29fd2c97f7dc52498716347.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleGroups()(*ib91d9df9d429e0fc35844f475381fd503a15a677bacc1be75db14e120d0c4890.IncompatibleGroupsRequestBuilder) {
    return ib91d9df9d429e0fc35844f475381fd503a15a677bacc1be75db14e120d0c4890.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageRequestBuilder) Patch(options *AccessPackageRequestBuilderPatchOptions)(error) {
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
