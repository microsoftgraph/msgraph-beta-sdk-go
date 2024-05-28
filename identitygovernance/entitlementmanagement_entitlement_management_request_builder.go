package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementEntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.identityGovernance entity.
type EntitlementmanagementEntitlementManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters get entitlementManagement from identityGovernance
type EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters
}
// EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentApprovals provides operations to manage the accessPackageAssignmentApprovals property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageassignmentapprovalsAccessPackageAssignmentApprovalsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals()(*EntitlementmanagementAccesspackageassignmentapprovalsAccessPackageAssignmentApprovalsRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentapprovalsAccessPackageAssignmentApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignmentPolicies provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageassignmentpoliciesAccessPackageAssignmentPoliciesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageAssignmentPolicies()(*EntitlementmanagementAccesspackageassignmentpoliciesAccessPackageAssignmentPoliciesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentpoliciesAccessPackageAssignmentPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignmentRequests provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageassignmentrequestsAccessPackageAssignmentRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageAssignmentRequests()(*EntitlementmanagementAccesspackageassignmentrequestsAccessPackageAssignmentRequestsRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentrequestsAccessPackageAssignmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignmentResourceRoles provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageassignmentresourcerolesAccessPackageAssignmentResourceRolesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRoles()(*EntitlementmanagementAccesspackageassignmentresourcerolesAccessPackageAssignmentResourceRolesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentresourcerolesAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignments provides operations to manage the accessPackageAssignments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageassignmentsAccessPackageAssignmentsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageAssignments()(*EntitlementmanagementAccesspackageassignmentsAccessPackageAssignmentsRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsAccessPackageAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageCatalogs provides operations to manage the accessPackageCatalogs property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackagecatalogsAccessPackageCatalogsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageCatalogs()(*EntitlementmanagementAccesspackagecatalogsAccessPackageCatalogsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsAccessPackageCatalogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceEnvironments provides operations to manage the accessPackageResourceEnvironments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageResourceEnvironments()(*EntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourceenvironmentsAccessPackageResourceEnvironmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceRequests provides operations to manage the accessPackageResourceRequests property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageresourcerequestsAccessPackageResourceRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageResourceRequests()(*EntitlementmanagementAccesspackageresourcerequestsAccessPackageResourceRequestsRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourcerequestsAccessPackageResourceRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceRoleScopes provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageresourcerolescopesAccessPackageResourceRoleScopesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageResourceRoleScopes()(*EntitlementmanagementAccesspackageresourcerolescopesAccessPackageResourceRoleScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourcerolescopesAccessPackageResourceRoleScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResources provides operations to manage the accessPackageResources property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageresourcesAccessPackageResourcesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageResources()(*EntitlementmanagementAccesspackageresourcesAccessPackageResourcesRequestBuilder) {
    return NewEntitlementmanagementAccesspackageresourcesAccessPackageResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackagesAccessPackagesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackages()(*EntitlementmanagementAccesspackagesAccessPackagesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesAccessPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentRequests provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAssignmentrequestsAssignmentRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AssignmentRequests()(*EntitlementmanagementAssignmentrequestsAssignmentRequestsRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsAssignmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConnectedOrganizations provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ConnectedOrganizations()(*EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) {
    return NewEntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementEntitlementManagementRequestBuilderInternal instantiates a new EntitlementmanagementEntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    m := &EntitlementmanagementEntitlementManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementEntitlementManagementRequestBuilder instantiates a new EntitlementmanagementEntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementmanagementEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property entitlementManagement for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entitlementManagement from identityGovernance
// returns a EntitlementManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEntitlementManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable), nil
}
// Patch update the navigation property entitlementManagement in identityGovernance
// returns a EntitlementManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEntitlementManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementSettingsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Settings()(*EntitlementmanagementSettingsRequestBuilder) {
    return NewEntitlementmanagementSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subjects provides operations to manage the subjects property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementSubjectsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Subjects()(*EntitlementmanagementSubjectsRequestBuilder) {
    return NewEntitlementmanagementSubjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SubjectsWithObjectId provides operations to manage the subjects property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementSubjectswithobjectidSubjectsWithObjectIdRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) SubjectsWithObjectId(objectId *string)(*EntitlementmanagementSubjectswithobjectidSubjectsWithObjectIdRequestBuilder) {
    return NewEntitlementmanagementSubjectswithobjectidSubjectsWithObjectIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, objectId)
}
// ToDeleteRequestInformation delete navigation property entitlementManagement for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get entitlementManagement from identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property entitlementManagement in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementEntitlementManagementRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    return NewEntitlementmanagementEntitlementManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
