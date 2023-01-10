package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder provides operations to manage the accessPackageResourceRoleScopes property of the microsoft.graph.accessPackage entity.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderGetQueryParameters get accessPackageResourceRoleScopes from identityGovernance
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceRole provides operations to manage the accessPackageResourceRole property of the microsoft.graph.accessPackageResourceRoleScope entity.
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) AccessPackageResourceRole()(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScope provides operations to manage the accessPackageResourceScope property of the microsoft.graph.accessPackageResourceRoleScope entity.
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) AccessPackageResourceScope()(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderInternal instantiates a new AccessPackageResourceRoleScopeItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder instantiates a new AccessPackageResourceRoleScopeItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResourceRoleScopes for identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get get accessPackageResourceRoleScopes from identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable), nil
}
// Patch update the navigation property accessPackageResourceRoleScopes in identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackageResourceRoleScopes for identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get accessPackageResourceRoleScopes from identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the navigation property accessPackageResourceRoleScopes in identityGovernance
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleScopeable, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesAccessPackageResourceRoleScopeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
