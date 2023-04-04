package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceRole entity.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScopesById provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) AccessPackageResourceScopesById(id string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope%2Did"] = id
    }
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderInternal instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceScope/accessPackageResource/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResource for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get get accessPackageResource from identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
// Patch update the navigation property accessPackageResource in identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get accessPackageResource from identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property accessPackageResource in identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
