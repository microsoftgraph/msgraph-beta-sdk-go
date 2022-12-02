package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceScope entity.
type IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoles provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRolesById provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) AccessPackageResourceRolesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderInternal instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceRole/accessPackageResource/accessPackageResourceScopes/{accessPackageResourceScope%2Did}/accessPackageResource{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get accessPackageResource from identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageResource in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackageResource for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get accessPackageResource from identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
// Patch update the navigation property accessPackageResource in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
