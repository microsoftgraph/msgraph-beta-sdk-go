package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.identityGovernance entity.
type EntitlementManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
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
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementRequestBuilderGetQueryParameters
}
// EntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentApprovals provides operations to manage the accessPackageAssignmentApprovals property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageAssignmentApprovalsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals()(*EntitlementManagementAccessPackageAssignmentApprovalsRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignmentPolicies provides operations to manage the accessPackageAssignmentPolicies property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageAssignmentPoliciesRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentPolicies()(*EntitlementManagementAccessPackageAssignmentPoliciesRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignmentRequests provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageAssignmentRequestsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentRequests()(*EntitlementManagementAccessPackageAssignmentRequestsRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignmentResourceRoles provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageAssignmentResourceRolesRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRoles()(*EntitlementManagementAccessPackageAssignmentResourceRolesRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageAssignments provides operations to manage the accessPackageAssignments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageAssignmentsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageAssignments()(*EntitlementManagementAccessPackageAssignmentsRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageCatalogs provides operations to manage the accessPackageCatalogs property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageCatalogsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageCatalogs()(*EntitlementManagementAccessPackageCatalogsRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceEnvironments provides operations to manage the accessPackageResourceEnvironments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageResourceEnvironmentsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceEnvironments()(*EntitlementManagementAccessPackageResourceEnvironmentsRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourceEnvironmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceRequests provides operations to manage the accessPackageResourceRequests property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageResourceRequestsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRequests()(*EntitlementManagementAccessPackageResourceRequestsRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourceRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceRoleScopes provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageResourceRoleScopesRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageResourceRoleScopes()(*EntitlementManagementAccessPackageResourceRoleScopesRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourceRoleScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResources provides operations to manage the accessPackageResources property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackageResourcesRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackageResources()(*EntitlementManagementAccessPackageResourcesRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAccessPackagesRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AccessPackages()(*EntitlementManagementAccessPackagesRequestBuilder) {
    return NewEntitlementManagementAccessPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentRequests provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementAssignmentRequestsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) AssignmentRequests()(*EntitlementManagementAssignmentRequestsRequestBuilder) {
    return NewEntitlementManagementAssignmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConnectedOrganizations provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementConnectedOrganizationsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) ConnectedOrganizations()(*EntitlementManagementConnectedOrganizationsRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    m := &EntitlementManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property entitlementManagement for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *EntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, error) {
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
func (m *EntitlementManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *EntitlementManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, error) {
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
// returns a *EntitlementManagementSettingsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) Settings()(*EntitlementManagementSettingsRequestBuilder) {
    return NewEntitlementManagementSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subjects provides operations to manage the subjects property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementSubjectsRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) Subjects()(*EntitlementManagementSubjectsRequestBuilder) {
    return NewEntitlementManagementSubjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SubjectsWithObjectId provides operations to manage the subjects property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementManagementSubjectsWithObjectIdRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) SubjectsWithObjectId(objectId *string)(*EntitlementManagementSubjectsWithObjectIdRequestBuilder) {
    return NewEntitlementManagementSubjectsWithObjectIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, objectId)
}
// ToDeleteRequestInformation delete navigation property entitlementManagement for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementManagementRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identityGovernance/entitlementManagement", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get entitlementManagement from identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EntitlementManagementable, requestConfiguration *EntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/identityGovernance/entitlementManagement", m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementManagementRequestBuilder when successful
func (m *EntitlementManagementRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementRequestBuilder) {
    return NewEntitlementManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
