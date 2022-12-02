package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.identityGovernance entity.
type IdentityGovernanceEntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementRequestBuilderGetQueryParameters get entitlementManagement from identityGovernance
type IdentityGovernanceEntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentApprovals provides operations to manage the accessPackageAssignmentApprovals property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentApprovalsById provides operations to manage the accessPackageAssignmentApprovals property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentApprovalsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignmentPolicies provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentPolicies()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentPoliciesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignmentRequests provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentRequests()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequestsById provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentRequestsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignmentResourceRoles provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRoles()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentResourceRolesById provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRolesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentResourceRole%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRolesAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageAssignments provides operations to manage the accessPackageAssignments property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignments()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentsById provides operations to manage the accessPackageAssignments property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageCatalogs provides operations to manage the accessPackageCatalogs property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageCatalogs()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageCatalogsById provides operations to manage the accessPackageCatalogs property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageCatalogsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageCatalog%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceEnvironments provides operations to manage the accessPackageResourceEnvironments property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageResourceEnvironments()(*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceEnvironmentsById provides operations to manage the accessPackageResourceEnvironments property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageResourceEnvironmentsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceEnvironment%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceEnvironmentsAccessPackageResourceEnvironmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceRequests provides operations to manage the accessPackageResourceRequests property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageResourceRequests()(*IdentityGovernanceEntitlementManagementAccessPackageResourceRequestsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRequestsById provides operations to manage the accessPackageResourceRequests property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageResourceRequestsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageResourceRequestsAccessPackageResourceRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRequest%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceRequestsAccessPackageResourceRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceRoleScopes provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageResourceRoleScopes()(*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResources provides operations to manage the accessPackageResources property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageResources()(*IdentityGovernanceEntitlementManagementAccessPackageResourcesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourcesById provides operations to manage the accessPackageResources property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageResourcesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageResourcesAccessPackageResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourcesAccessPackageResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackages()(*IdentityGovernanceEntitlementManagementAccessPackagesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesById provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackagesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConnectedOrganizations provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) ConnectedOrganizations()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectedOrganizationsById provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) ConnectedOrganizationsById(id string)(*IdentityGovernanceEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectedOrganization%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementRequestBuilder{
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
// NewIdentityGovernanceEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property entitlementManagement for identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entitlementManagement from identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property entitlementManagement in identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property entitlementManagement for identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entitlementManagement from identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEntitlementManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable), nil
}
// Patch update the navigation property entitlementManagement in identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEntitlementManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Settings()(*IdentityGovernanceEntitlementManagementSettingsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Subjects provides operations to manage the subjects property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Subjects()(*IdentityGovernanceEntitlementManagementSubjectsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementSubjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubjectsById provides operations to manage the subjects property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) SubjectsById(id string)(*IdentityGovernanceEntitlementManagementSubjectsAccessPackageSubjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageSubject%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementSubjectsAccessPackageSubjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
