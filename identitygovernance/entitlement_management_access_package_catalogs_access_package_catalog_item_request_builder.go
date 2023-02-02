package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder provides operations to manage the accessPackageCatalogs property of the microsoft.graph.entitlementManagement entity.
type EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters a container of access packages.
type EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceRoles provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceRoles()(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackageResourceRolesById provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceRolesById(id string)(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AccessPackageResources provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResources()(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackageResourcesById provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourcesById(id string)(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesAccessPackageResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesAccessPackageResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceScopes()(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackageResourceScopesById provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceScopesById(id string)(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackages()(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AccessPackagesById provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackagesById(id string)(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal instantiates a new AccessPackageCatalogItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, accessPackageCatalogId *string)(*EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) {
    m := &EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageCatalogId != nil {
        urlTplParams["accessPackageCatalog%2Did"] = *accessPackageCatalogId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder instantiates a new AccessPackageCatalogItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CustomAccessPackageWorkflowExtensions provides operations to manage the customAccessPackageWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) CustomAccessPackageWorkflowExtensions()(*EntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CustomAccessPackageWorkflowExtensionsById provides operations to manage the customAccessPackageWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) CustomAccessPackageWorkflowExtensionsById(id string)(*EntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsCustomAccessPackageWorkflowExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsCustomAccessPackageWorkflowExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Delete delete navigation property accessPackageCatalogs for identityGovernance
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a container of access packages.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable), nil
}
// Patch update the navigation property accessPackageCatalogs in identityGovernance
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackageCatalogs for identityGovernance
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation a container of access packages.
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property accessPackageCatalogs in identityGovernance
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
