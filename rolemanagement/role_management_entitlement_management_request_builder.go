package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementEntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.roleManagement entity.
type RoleManagementEntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementEntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementEntitlementManagementRequestBuilderGetQueryParameters the RbacApplication for Entitlement Management
type RoleManagementEntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementEntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementEntitlementManagementRequestBuilderGetQueryParameters
}
// RoleManagementEntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRequestBuilder) {
    m := &RoleManagementEntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property entitlementManagement for roleManagement
func (m *RoleManagementEntitlementManagementRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the RbacApplication for Entitlement Management
func (m *RoleManagementEntitlementManagementRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property entitlementManagement in roleManagement
func (m *RoleManagementEntitlementManagementRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *RoleManagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property entitlementManagement for roleManagement
func (m *RoleManagementEntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the RbacApplication for Entitlement Management
func (m *RoleManagementEntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable), nil
}
// Patch update the navigation property entitlementManagement in roleManagement
func (m *RoleManagementEntitlementManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *RoleManagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable), nil
}
// ResourceNamespaces provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) ResourceNamespaces()(*RoleManagementEntitlementManagementResourceNamespacesRequestBuilder) {
    return NewRoleManagementEntitlementManagementResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceNamespacesById provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) ResourceNamespacesById(id string)(*RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentApprovals provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentApprovals()(*RoleManagementEntitlementManagementRoleAssignmentApprovalsRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentApprovalsById provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentApprovalsById(id string)(*RoleManagementEntitlementManagementRoleAssignmentApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleAssignmentApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignments()(*RoleManagementEntitlementManagementRoleAssignmentsRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentsById(id string)(*RoleManagementEntitlementManagementRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleInstances provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentScheduleInstances()(*RoleManagementEntitlementManagementRoleAssignmentScheduleInstancesRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleInstancesById provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentScheduleInstancesById(id string)(*RoleManagementEntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleRequests provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentScheduleRequests()(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleRequestsById provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentScheduleRequestsById(id string)(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleRequest%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentSchedules provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentSchedules()(*RoleManagementEntitlementManagementRoleAssignmentSchedulesRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentSchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentSchedulesById provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleAssignmentSchedulesById(id string)(*RoleManagementEntitlementManagementRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentSchedule%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleDefinitions()(*RoleManagementEntitlementManagementRoleDefinitionsRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleDefinitionsById(id string)(*RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleInstances provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleEligibilityScheduleInstances()(*RoleManagementEntitlementManagementRoleEligibilityScheduleInstancesRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleInstancesById provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleEligibilityScheduleInstancesById(id string)(*RoleManagementEntitlementManagementRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleInstance%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleRequests provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleEligibilityScheduleRequests()(*RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleRequestsById provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleEligibilityScheduleRequestsById(id string)(*RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilitySchedules provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleEligibilitySchedules()(*RoleManagementEntitlementManagementRoleEligibilitySchedulesRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleEligibilitySchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilitySchedulesById provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleEligibilitySchedulesById(id string)(*RoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilitySchedule%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId provides operations to call the roleScheduleInstances method.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId()(*RoleManagementEntitlementManagementRoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId provides operations to call the roleSchedules method.
func (m *RoleManagementEntitlementManagementRequestBuilder) RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId()(*RoleManagementEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveRoleAssignments provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) TransitiveRoleAssignments()(*RoleManagementEntitlementManagementTransitiveRoleAssignmentsRequestBuilder) {
    return NewRoleManagementEntitlementManagementTransitiveRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveRoleAssignmentsById provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementEntitlementManagementRequestBuilder) TransitiveRoleAssignmentsById(id string)(*RoleManagementEntitlementManagementTransitiveRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementTransitiveRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
