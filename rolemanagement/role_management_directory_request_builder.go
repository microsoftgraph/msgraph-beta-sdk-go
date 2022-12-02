package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementDirectoryRequestBuilder provides operations to manage the directory property of the microsoft.graph.roleManagement entity.
type RoleManagementDirectoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementDirectoryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDirectoryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementDirectoryRequestBuilderGetQueryParameters get directory from roleManagement
type RoleManagementDirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementDirectoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDirectoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementDirectoryRequestBuilderGetQueryParameters
}
// RoleManagementDirectoryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDirectoryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewRoleManagementDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDirectoryRequestBuilder) {
    m := &RoleManagementDirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewRoleManagementDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property directory for roleManagement
func (m *RoleManagementDirectoryRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDirectoryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get directory from roleManagement
func (m *RoleManagementDirectoryRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDirectoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property directory in roleManagement
func (m *RoleManagementDirectoryRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *RoleManagementDirectoryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property directory for roleManagement
func (m *RoleManagementDirectoryRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementDirectoryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get directory from roleManagement
func (m *RoleManagementDirectoryRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementDirectoryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
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
// Patch update the navigation property directory in roleManagement
func (m *RoleManagementDirectoryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *RoleManagementDirectoryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
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
func (m *RoleManagementDirectoryRequestBuilder) ResourceNamespaces()(*RoleManagementDirectoryResourceNamespacesRequestBuilder) {
    return NewRoleManagementDirectoryResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceNamespacesById provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) ResourceNamespacesById(id string)(*RoleManagementDirectoryResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace%2Did"] = id
    }
    return NewRoleManagementDirectoryResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentApprovals provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentApprovals()(*RoleManagementDirectoryRoleAssignmentApprovalsRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentApprovalsById provides operations to manage the roleAssignmentApprovals property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentApprovalsById(id string)(*RoleManagementDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleAssignmentApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignments()(*RoleManagementDirectoryRoleAssignmentsRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentsById(id string)(*RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleInstances provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentScheduleInstances()(*RoleManagementDirectoryRoleAssignmentScheduleInstancesRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleInstancesById provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentScheduleInstancesById(id string)(*RoleManagementDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleRequests provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentScheduleRequests()(*RoleManagementDirectoryRoleAssignmentScheduleRequestsRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleRequestsById provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentScheduleRequestsById(id string)(*RoleManagementDirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleRequest%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentSchedules provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentSchedules()(*RoleManagementDirectoryRoleAssignmentSchedulesRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentSchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentSchedulesById provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleAssignmentSchedulesById(id string)(*RoleManagementDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentSchedule%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleAssignmentSchedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleDefinitions()(*RoleManagementDirectoryRoleDefinitionsRequestBuilder) {
    return NewRoleManagementDirectoryRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleDefinitionsById(id string)(*RoleManagementDirectoryRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleInstances provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleEligibilityScheduleInstances()(*RoleManagementDirectoryRoleEligibilityScheduleInstancesRequestBuilder) {
    return NewRoleManagementDirectoryRoleEligibilityScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleInstancesById provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleEligibilityScheduleInstancesById(id string)(*RoleManagementDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleInstance%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleRequests provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleEligibilityScheduleRequests()(*RoleManagementDirectoryRoleEligibilityScheduleRequestsRequestBuilder) {
    return NewRoleManagementDirectoryRoleEligibilityScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleRequestsById provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleEligibilityScheduleRequestsById(id string)(*RoleManagementDirectoryRoleEligibilityScheduleRequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleEligibilityScheduleRequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilitySchedules provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleEligibilitySchedules()(*RoleManagementDirectoryRoleEligibilitySchedulesRequestBuilder) {
    return NewRoleManagementDirectoryRoleEligibilitySchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilitySchedulesById provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) RoleEligibilitySchedulesById(id string)(*RoleManagementDirectoryRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilitySchedule%2Did"] = id
    }
    return NewRoleManagementDirectoryRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId provides operations to call the roleScheduleInstances method.
func (m *RoleManagementDirectoryRequestBuilder) RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId()(*RoleManagementDirectoryRoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return NewRoleManagementDirectoryRoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId provides operations to call the roleSchedules method.
func (m *RoleManagementDirectoryRequestBuilder) RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId()(*RoleManagementDirectoryRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return NewRoleManagementDirectoryRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveRoleAssignments provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) TransitiveRoleAssignments()(*RoleManagementDirectoryTransitiveRoleAssignmentsRequestBuilder) {
    return NewRoleManagementDirectoryTransitiveRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveRoleAssignmentsById provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
func (m *RoleManagementDirectoryRequestBuilder) TransitiveRoleAssignmentsById(id string)(*RoleManagementDirectoryTransitiveRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return NewRoleManagementDirectoryTransitiveRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
