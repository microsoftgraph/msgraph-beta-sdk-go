package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder provides operations to manage the accessPackageCatalogs property of the microsoft.graph.entitlementManagement entity.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters a container of access packages.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceRoles provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceRoles()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRolesById provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceRolesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResources provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResources()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourcesById provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourcesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesAccessPackageResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesAccessPackageResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceScopes()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScopesById provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceScopesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackages()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesById provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackagesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal instantiates a new AccessPackageCatalogItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder instantiates a new AccessPackageCatalogItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageCatalogs for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a container of access packages.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageCatalogs in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CustomAccessPackageWorkflowExtensions provides operations to manage the customAccessPackageWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) CustomAccessPackageWorkflowExtensions()(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomAccessPackageWorkflowExtensionsById provides operations to manage the customAccessPackageWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) CustomAccessPackageWorkflowExtensionsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsCustomAccessPackageWorkflowExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customAccessPackageWorkflowExtension%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsCustomAccessPackageWorkflowExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property accessPackageCatalogs for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a container of access packages.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable), nil
}
// Patch update the navigation property accessPackageCatalogs in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable), nil
}
