package entitlementmanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0d1afe57eb8a36b53e5129a0b937776c5fcca08608b49a51664fbd3159daa08a "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourceenvironments"
    i288a3763f785b246d1da17397590bd077e283e82acf56e8c177d923272c52724 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments"
    i3d425a2bba157a442f755259375e0a18ff87074e6c0ec53bac6d292a43602107 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations"
    i4457cf48628c9e6d5d1e3963ea8a19dd7f20c98f2fd6035004babfeb7e44fbc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentapprovals"
    i511c33fd684f429fa206d02a10594bbfb5f9acceeb4f93b868ea9062ddbb0c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentpolicies"
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

// EntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.identityGovernance entity.
type EntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementRequestBuilderGetQueryParameters get entitlementManagement from identityGovernance
type EntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementRequestBuilderGetQueryParameters
}
// EntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentApprovals the accessPackageAssignmentApprovals property
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals()(*i4457cf48628c9e6d5d1e3963ea8a19dd7f20c98f2fd6035004babfeb7e44fbc8.AccessPackageAssignmentApprovalsRequestBuilder) {
    return i4457cf48628c9e6d5d1e3963ea8a19dd7f20c98f2fd6035004babfeb7e44fbc8.NewAccessPackageAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentApprovalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentApprovals.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentApprovalsById(id string)(*ifa7e2fb3dcb6477dacae87733273579930d76bee0f987eee04a36237466ad30b.ApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return ifa7e2fb3dcb6477dacae87733273579930d76bee0f987eee04a36237466ad30b.NewApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignmentPolicies the accessPackageAssignmentPolicies property
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentPolicies()(*i511c33fd684f429fa206d02a10594bbfb5f9acceeb4f93b868ea9062ddbb0c0f.AccessPackageAssignmentPoliciesRequestBuilder) {
    return i511c33fd684f429fa206d02a10594bbfb5f9acceeb4f93b868ea9062ddbb0c0f.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentPolicies.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i259297720910df1e18cab41eaf6e93c6e506e2f6ead7e552913518ba4de7df35.AccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return i259297720910df1e18cab41eaf6e93c6e506e2f6ead7e552913518ba4de7df35.NewAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignmentRequests the accessPackageAssignmentRequests property
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentRequests()(*i824703be434261db4e58f4b65123444442b5ebe1eb3fb9b603f8864af564b622.AccessPackageAssignmentRequestsRequestBuilder) {
    return i824703be434261db4e58f4b65123444442b5ebe1eb3fb9b603f8864af564b622.NewAccessPackageAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentRequestsById(id string)(*i76c412e0065634c6fc606ecd0b6e09322c5b7f7b51446f949b2f794f2fc40926.AccessPackageAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest%2Did"] = id
    }
    return i76c412e0065634c6fc606ecd0b6e09322c5b7f7b51446f949b2f794f2fc40926.NewAccessPackageAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignmentResourceRoles the accessPackageAssignmentResourceRoles property
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRoles()(*iba574113ccbfd66c0b718094a7417a26f2be8bc7a95f74bfffd609a4a5312cd1.AccessPackageAssignmentResourceRolesRequestBuilder) {
    return iba574113ccbfd66c0b718094a7417a26f2be8bc7a95f74bfffd609a4a5312cd1.NewAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRolesById(id string)(*i46b0d052826008fb4b9e2de18eaadb28b08fda0c4481c35e9fdf6c08800480a1.AccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentResourceRole%2Did"] = id
    }
    return i46b0d052826008fb4b9e2de18eaadb28b08fda0c4481c35e9fdf6c08800480a1.NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignments the accessPackageAssignments property
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignments()(*i288a3763f785b246d1da17397590bd077e283e82acf56e8c177d923272c52724.AccessPackageAssignmentsRequestBuilder) {
    return i288a3763f785b246d1da17397590bd077e283e82acf56e8c177d923272c52724.NewAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentsById(id string)(*ib0d512b118f91cf1bc458b5feacf403e01f72b051f66b4e84985a41cb52dab99.AccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment%2Did"] = id
    }
    return ib0d512b118f91cf1bc458b5feacf403e01f72b051f66b4e84985a41cb52dab99.NewAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageCatalogs the accessPackageCatalogs property
func (m *EntitlementManagementRequestBuilder) AccessPackageCatalogs()(*i7b6be768761b0aa85e64ce3193b16f7637eb84e57ba59b80caa82554e9087b35.AccessPackageCatalogsRequestBuilder) {
    return i7b6be768761b0aa85e64ce3193b16f7637eb84e57ba59b80caa82554e9087b35.NewAccessPackageCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageCatalogsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageCatalogsById(id string)(*i82310b3d54600cbafcfbb579b511cc82bb29ed7fb49516f06dce0fcc7951bdf0.AccessPackageCatalogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageCatalog%2Did"] = id
    }
    return i82310b3d54600cbafcfbb579b511cc82bb29ed7fb49516f06dce0fcc7951bdf0.NewAccessPackageCatalogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceEnvironments the accessPackageResourceEnvironments property
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceEnvironments()(*i0d1afe57eb8a36b53e5129a0b937776c5fcca08608b49a51664fbd3159daa08a.AccessPackageResourceEnvironmentsRequestBuilder) {
    return i0d1afe57eb8a36b53e5129a0b937776c5fcca08608b49a51664fbd3159daa08a.NewAccessPackageResourceEnvironmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceEnvironmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResourceEnvironments.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceEnvironmentsById(id string)(*id480c6778d8753dfd5aecaa8cfff24f904f1d9cfb89137f57d90578c37490dde.AccessPackageResourceEnvironmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceEnvironment%2Did"] = id
    }
    return id480c6778d8753dfd5aecaa8cfff24f904f1d9cfb89137f57d90578c37490dde.NewAccessPackageResourceEnvironmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceRequests the accessPackageResourceRequests property
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRequests()(*ieba606bbb218a20cbd99e9bc1e380a2e15b31da11df8da17e8bcdbcf42633d86.AccessPackageResourceRequestsRequestBuilder) {
    return ieba606bbb218a20cbd99e9bc1e380a2e15b31da11df8da17e8bcdbcf42633d86.NewAccessPackageResourceRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResourceRequests.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRequestsById(id string)(*if945e55d38fc375c56dbf90cf6e39a7d61a6cd01059d47fb2b23e38f45842f3e.AccessPackageResourceRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRequest%2Did"] = id
    }
    return if945e55d38fc375c56dbf90cf6e39a7d61a6cd01059d47fb2b23e38f45842f3e.NewAccessPackageResourceRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceRoleScopes the accessPackageResourceRoleScopes property
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRoleScopes()(*i5c68f346a5bdbb8fe1dea52fb202c7346657f21eeba8976e031c861e4999a938.AccessPackageResourceRoleScopesRequestBuilder) {
    return i5c68f346a5bdbb8fe1dea52fb202c7346657f21eeba8976e031c861e4999a938.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResourceRoleScopes.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i73885afb0012ccafe2d98c7ad3770430abc8b68787e78c60063bf998dee8be11.AccessPackageResourceRoleScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope%2Did"] = id
    }
    return i73885afb0012ccafe2d98c7ad3770430abc8b68787e78c60063bf998dee8be11.NewAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResources the accessPackageResources property
func (m *EntitlementManagementRequestBuilder) AccessPackageResources()(*if86a04543e04099c6faebfb83f2509f6287141cd5f4ff9c77352a589372a8f17.AccessPackageResourcesRequestBuilder) {
    return if86a04543e04099c6faebfb83f2509f6287141cd5f4ff9c77352a589372a8f17.NewAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageResources.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackageResourcesById(id string)(*ia3f92f11727041477c1ada6bfc08139cadda160d066f173db00cce4f8358bc50.AccessPackageResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource%2Did"] = id
    }
    return ia3f92f11727041477c1ada6bfc08139cadda160d066f173db00cce4f8358bc50.NewAccessPackageResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackages the accessPackages property
func (m *EntitlementManagementRequestBuilder) AccessPackages()(*ib411956083b5fb15c9a26abd5c5030d74fc10163e4c5a0f8d95848eca3093621.AccessPackagesRequestBuilder) {
    return ib411956083b5fb15c9a26abd5c5030d74fc10163e4c5a0f8d95848eca3093621.NewAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item collection
func (m *EntitlementManagementRequestBuilder) AccessPackagesById(id string)(*i832824d2195e5d19b0fbdef0bb45eca79f632cf74d39d2b1535ca2c63e49b230.AccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did"] = id
    }
    return i832824d2195e5d19b0fbdef0bb45eca79f632cf74d39d2b1535ca2c63e49b230.NewAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConnectedOrganizations the connectedOrganizations property
func (m *EntitlementManagementRequestBuilder) ConnectedOrganizations()(*i3d425a2bba157a442f755259375e0a18ff87074e6c0ec53bac6d292a43602107.ConnectedOrganizationsRequestBuilder) {
    return i3d425a2bba157a442f755259375e0a18ff87074e6c0ec53bac6d292a43602107.NewConnectedOrganizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectedOrganizationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item collection
func (m *EntitlementManagementRequestBuilder) ConnectedOrganizationsById(id string)(*i6733f0715c511b6a99e8e8280d52d93704622cda212a0d03f6340a8c98c403e5.ConnectedOrganizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectedOrganization%2Did"] = id
    }
    return i6733f0715c511b6a99e8e8280d52d93704622cda212a0d03f6340a8c98c403e5.NewConnectedOrganizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    m := &EntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property entitlementManagement for identityGovernance
func (m *EntitlementManagementRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property entitlementManagement for identityGovernance
func (m *EntitlementManagementRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformationWithRequestConfiguration get entitlementManagement from identityGovernance
func (m *EntitlementManagementRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get entitlementManagement from identityGovernance
func (m *EntitlementManagementRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property entitlementManagement in identityGovernance
func (m *EntitlementManagementRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property entitlementManagement in identityGovernance
func (m *EntitlementManagementRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *EntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property entitlementManagement for identityGovernance
func (m *EntitlementManagementRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property entitlementManagement for identityGovernance
func (m *EntitlementManagementRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EntitlementManagementRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler get entitlementManagement from identityGovernance
func (m *EntitlementManagementRequestBuilder) GetWithResponseHandler(requestConfiguration *EntitlementManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler get entitlementManagement from identityGovernance
func (m *EntitlementManagementRequestBuilder) GetWithResponseHandler(requestConfiguration *EntitlementManagementRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEntitlementManagementFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable), nil
}
// PatchWithResponseHandler update the navigation property entitlementManagement in identityGovernance
func (m *EntitlementManagementRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *EntitlementManagementRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property entitlementManagement in identityGovernance
func (m *EntitlementManagementRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *EntitlementManagementRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Settings the settings property
func (m *EntitlementManagementRequestBuilder) Settings()(*i8d110cb2de6a3e3af2c3ad8f85fecae6316e989d5010ae9071a62d837f70c4f9.SettingsRequestBuilder) {
    return i8d110cb2de6a3e3af2c3ad8f85fecae6316e989d5010ae9071a62d837f70c4f9.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
