package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
type IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResource provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceScope entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) AccessPackageResource()(*IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesItemAccessPackageResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal instantiates a new AccessPackageResourceScopeItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResources/{accessPackageResource%2Did}/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource/accessPackageResourceScopes/{accessPackageResourceScope%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder instantiates a new AccessPackageResourceScopeItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageResourceScopes for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageResourceScopes in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackageResourceScopes for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable), nil
}
// Patch update the navigation property accessPackageResourceScopes in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable), nil
}
