package entitlementmanagement

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0d1afe57eb8a36b53e5129a0b937776c5fcca08608b49a51664fbd3159daa08a "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourceenvironments"
    i288a3763f785b246d1da17397590bd077e283e82acf56e8c177d923272c52724 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments"
    i3d425a2bba157a442f755259375e0a18ff87074e6c0ec53bac6d292a43602107 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations"
    i4457cf48628c9e6d5d1e3963ea8a19dd7f20c98f2fd6035004babfeb7e44fbc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentapprovals"
    i511c33fd684f429fa206d02a10594bbfb5f9acceeb4f93b868ea9062ddbb0c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentpolicies"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5c68f346a5bdbb8fe1dea52fb202c7346657f21eeba8976e031c861e4999a938 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourcerolescopes"
    i7b6be768761b0aa85e64ce3193b16f7637eb84e57ba59b80caa82554e9087b35 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs"
    i824703be434261db4e58f4b65123444442b5ebe1eb3fb9b603f8864af564b622 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests"
    i8d110cb2de6a3e3af2c3ad8f85fecae6316e989d5010ae9071a62d837f70c4f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/settings"
    ib411956083b5fb15c9a26abd5c5030d74fc10163e4c5a0f8d95848eca3093621 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages"
    iba574113ccbfd66c0b718094a7417a26f2be8bc7a95f74bfffd609a4a5312cd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles"
    ieba606bbb218a20cbd99e9bc1e380a2e15b31da11df8da17e8bcdbcf42633d86 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourcerequests"
    if86a04543e04099c6faebfb83f2509f6287141cd5f4ff9c77352a589372a8f17 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresources"
    i259297720910df1e18cab41eaf6e93c6e506e2f6ead7e552913518ba4de7df35 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentpolicies/item"
    i46b0d052826008fb4b9e2de18eaadb28b08fda0c4481c35e9fdf6c08800480a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item"
    i6733f0715c511b6a99e8e8280d52d93704622cda212a0d03f6340a8c98c403e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item"
    i73885afb0012ccafe2d98c7ad3770430abc8b68787e78c60063bf998dee8be11 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourcerolescopes/item"
    i76c412e0065634c6fc606ecd0b6e09322c5b7f7b51446f949b2f794f2fc40926 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item"
    i82310b3d54600cbafcfbb579b511cc82bb29ed7fb49516f06dce0fcc7951bdf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item"
    i832824d2195e5d19b0fbdef0bb45eca79f632cf74d39d2b1535ca2c63e49b230 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item"
    ia3f92f11727041477c1ada6bfc08139cadda160d066f173db00cce4f8358bc50 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresources/item"
    ib0d512b118f91cf1bc458b5feacf403e01f72b051f66b4e84985a41cb52dab99 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item"
    id480c6778d8753dfd5aecaa8cfff24f904f1d9cfb89137f57d90578c37490dde "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourceenvironments/item"
    if945e55d38fc375c56dbf90cf6e39a7d61a6cd01059d47fb2b23e38f45842f3e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourcerequests/item"
    ifa7e2fb3dcb6477dacae87733273579930d76bee0f987eee04a36237466ad30b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentapprovals/item"
)

// EntitlementManagementRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement
type EntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EntitlementManagementRequestBuilderDeleteOptions options for Delete
type EntitlementManagementRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EntitlementManagementRequestBuilderGetOptions options for Get
type EntitlementManagementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EntitlementManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EntitlementManagementRequestBuilderGetQueryParameters get entitlementManagement from identityGovernance
type EntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EntitlementManagementRequestBuilderPatchOptions options for Patch
type EntitlementManagementRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EntitlementManagement;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals()(*i4457cf48628c9e6d5d1e3963ea8a19dd7f20c98f2fd6035004babfeb7e44fbc8.AccessPackageAssignmentApprovalsRequestBuilder) {
    return i4457cf48628c9e6d5d1e3963ea8a19dd7f20c98f2fd6035004babfeb7e44fbc8.NewAccessPackageAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentApprovalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentApprovals.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentApprovalsById(id string)(*ifa7e2fb3dcb6477dacae87733273579930d76bee0f987eee04a36237466ad30b.ApprovalRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval_id"] = id
    }
    return ifa7e2fb3dcb6477dacae87733273579930d76bee0f987eee04a36237466ad30b.NewApprovalRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentPolicies()(*i511c33fd684f429fa206d02a10594bbfb5f9acceeb4f93b868ea9062ddbb0c0f.AccessPackageAssignmentPoliciesRequestBuilder) {
    return i511c33fd684f429fa206d02a10594bbfb5f9acceeb4f93b868ea9062ddbb0c0f.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentPolicies.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i259297720910df1e18cab41eaf6e93c6e506e2f6ead7e552913518ba4de7df35.AccessPackageAssignmentPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy_id"] = id
    }
    return i259297720910df1e18cab41eaf6e93c6e506e2f6ead7e552913518ba4de7df35.NewAccessPackageAssignmentPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentRequests()(*i824703be434261db4e58f4b65123444442b5ebe1eb3fb9b603f8864af564b622.AccessPackageAssignmentRequestsRequestBuilder) {
    return i824703be434261db4e58f4b65123444442b5ebe1eb3fb9b603f8864af564b622.NewAccessPackageAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentRequestsById(id string)(*i76c412e0065634c6fc606ecd0b6e09322c5b7f7b51446f949b2f794f2fc40926.AccessPackageAssignmentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest_id"] = id
    }
    return i76c412e0065634c6fc606ecd0b6e09322c5b7f7b51446f949b2f794f2fc40926.NewAccessPackageAssignmentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRoles()(*iba574113ccbfd66c0b718094a7417a26f2be8bc7a95f74bfffd609a4a5312cd1.AccessPackageAssignmentResourceRolesRequestBuilder) {
    return iba574113ccbfd66c0b718094a7417a26f2be8bc7a95f74bfffd609a4a5312cd1.NewAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRolesById(id string)(*i46b0d052826008fb4b9e2de18eaadb28b08fda0c4481c35e9fdf6c08800480a1.AccessPackageAssignmentResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentResourceRole_id"] = id
    }
    return i46b0d052826008fb4b9e2de18eaadb28b08fda0c4481c35e9fdf6c08800480a1.NewAccessPackageAssignmentResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignments()(*i288a3763f785b246d1da17397590bd077e283e82acf56e8c177d923272c52724.AccessPackageAssignmentsRequestBuilder) {
    return i288a3763f785b246d1da17397590bd077e283e82acf56e8c177d923272c52724.NewAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentsById(id string)(*ib0d512b118f91cf1bc458b5feacf403e01f72b051f66b4e84985a41cb52dab99.AccessPackageAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment_id"] = id
    }
    return ib0d512b118f91cf1bc458b5feacf403e01f72b051f66b4e84985a41cb52dab99.NewAccessPackageAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageCatalogs()(*i7b6be768761b0aa85e64ce3193b16f7637eb84e57ba59b80caa82554e9087b35.AccessPackageCatalogsRequestBuilder) {
    return i7b6be768761b0aa85e64ce3193b16f7637eb84e57ba59b80caa82554e9087b35.NewAccessPackageCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageCatalogsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageCatalogsById(id string)(*i82310b3d54600cbafcfbb579b511cc82bb29ed7fb49516f06dce0fcc7951bdf0.AccessPackageCatalogRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageCatalog_id"] = id
    }
    return i82310b3d54600cbafcfbb579b511cc82bb29ed7fb49516f06dce0fcc7951bdf0.NewAccessPackageCatalogRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceEnvironments()(*i0d1afe57eb8a36b53e5129a0b937776c5fcca08608b49a51664fbd3159daa08a.AccessPackageResourceEnvironmentsRequestBuilder) {
    return i0d1afe57eb8a36b53e5129a0b937776c5fcca08608b49a51664fbd3159daa08a.NewAccessPackageResourceEnvironmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceEnvironmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResourceEnvironments.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceEnvironmentsById(id string)(*id480c6778d8753dfd5aecaa8cfff24f904f1d9cfb89137f57d90578c37490dde.AccessPackageResourceEnvironmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceEnvironment_id"] = id
    }
    return id480c6778d8753dfd5aecaa8cfff24f904f1d9cfb89137f57d90578c37490dde.NewAccessPackageResourceEnvironmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRequests()(*ieba606bbb218a20cbd99e9bc1e380a2e15b31da11df8da17e8bcdbcf42633d86.AccessPackageResourceRequestsRequestBuilder) {
    return ieba606bbb218a20cbd99e9bc1e380a2e15b31da11df8da17e8bcdbcf42633d86.NewAccessPackageResourceRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResourceRequests.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRequestsById(id string)(*if945e55d38fc375c56dbf90cf6e39a7d61a6cd01059d47fb2b23e38f45842f3e.AccessPackageResourceRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRequest_id"] = id
    }
    return if945e55d38fc375c56dbf90cf6e39a7d61a6cd01059d47fb2b23e38f45842f3e.NewAccessPackageResourceRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRoleScopes()(*i5c68f346a5bdbb8fe1dea52fb202c7346657f21eeba8976e031c861e4999a938.AccessPackageResourceRoleScopesRequestBuilder) {
    return i5c68f346a5bdbb8fe1dea52fb202c7346657f21eeba8976e031c861e4999a938.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResourceRoleScopes.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i73885afb0012ccafe2d98c7ad3770430abc8b68787e78c60063bf998dee8be11.AccessPackageResourceRoleScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope_id"] = id
    }
    return i73885afb0012ccafe2d98c7ad3770430abc8b68787e78c60063bf998dee8be11.NewAccessPackageResourceRoleScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackageResources()(*if86a04543e04099c6faebfb83f2509f6287141cd5f4ff9c77352a589372a8f17.AccessPackageResourcesRequestBuilder) {
    return if86a04543e04099c6faebfb83f2509f6287141cd5f4ff9c77352a589372a8f17.NewAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResources.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageResourcesById(id string)(*ia3f92f11727041477c1ada6bfc08139cadda160d066f173db00cce4f8358bc50.AccessPackageResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource_id"] = id
    }
    return ia3f92f11727041477c1ada6bfc08139cadda160d066f173db00cce4f8358bc50.NewAccessPackageResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) AccessPackages()(*ib411956083b5fb15c9a26abd5c5030d74fc10163e4c5a0f8d95848eca3093621.AccessPackagesRequestBuilder) {
    return ib411956083b5fb15c9a26abd5c5030d74fc10163e4c5a0f8d95848eca3093621.NewAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackagesById(id string)(*i832824d2195e5d19b0fbdef0bb45eca79f632cf74d39d2b1535ca2c63e49b230.AccessPackageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage_id"] = id
    }
    return i832824d2195e5d19b0fbdef0bb45eca79f632cf74d39d2b1535ca2c63e49b230.NewAccessPackageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) ConnectedOrganizations()(*i3d425a2bba157a442f755259375e0a18ff87074e6c0ec53bac6d292a43602107.ConnectedOrganizationsRequestBuilder) {
    return i3d425a2bba157a442f755259375e0a18ff87074e6c0ec53bac6d292a43602107.NewConnectedOrganizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectedOrganizationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item collection
func (m *EntitlementManagementRequestBuilder) ConnectedOrganizationsById(id string)(*i6733f0715c511b6a99e8e8280d52d93704622cda212a0d03f6340a8c98c403e5.ConnectedOrganizationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectedOrganization_id"] = id
    }
    return i6733f0715c511b6a99e8e8280d52d93704622cda212a0d03f6340a8c98c403e5.NewConnectedOrganizationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    m := &EntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property entitlementManagement for identityGovernance
func (m *EntitlementManagementRequestBuilder) CreateDeleteRequestInformation(options *EntitlementManagementRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entitlementManagement from identityGovernance
func (m *EntitlementManagementRequestBuilder) CreateGetRequestInformation(options *EntitlementManagementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property entitlementManagement in identityGovernance
func (m *EntitlementManagementRequestBuilder) CreatePatchRequestInformation(options *EntitlementManagementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property entitlementManagement for identityGovernance
func (m *EntitlementManagementRequestBuilder) Delete(options *EntitlementManagementRequestBuilderDeleteOptions)(error) {
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
// Get get entitlementManagement from identityGovernance
func (m *EntitlementManagementRequestBuilder) Get(options *EntitlementManagementRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EntitlementManagement, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntitlementManagement() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EntitlementManagement), nil
}
// Patch update the navigation property entitlementManagement in identityGovernance
func (m *EntitlementManagementRequestBuilder) Patch(options *EntitlementManagementRequestBuilderPatchOptions)(error) {
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
func (m *EntitlementManagementRequestBuilder) Settings()(*i8d110cb2de6a3e3af2c3ad8f85fecae6316e989d5010ae9071a62d837f70c4f9.SettingsRequestBuilder) {
    return i8d110cb2de6a3e3af2c3ad8f85fecae6316e989d5010ae9071a62d837f70c4f9.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
