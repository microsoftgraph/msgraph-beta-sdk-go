package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResource provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceRole entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) AccessPackageResource()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal instantiates a new AccessPackageResourceRoleItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackage/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceScope/accessPackageResource/accessPackageResourceRoles/{accessPackageResourceRole%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder instantiates a new AccessPackageResourceRoleItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageResourceRoles for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable. Supports $expand.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageResourceRoles in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackageResourceRoles for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Nullable. Supports $expand.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable), nil
}
// Patch update the navigation property accessPackageResourceRoles in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable), nil
}
