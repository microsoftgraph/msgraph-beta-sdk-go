package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceRole entity.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScopesById provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) AccessPackageResourceScopesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderInternal instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackageResources/{accessPackageResource%2Did}/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
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
