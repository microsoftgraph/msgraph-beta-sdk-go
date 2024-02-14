package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder provides operations to manage the accessPackageCatalogs property of the microsoft.graph.entitlementManagement entity.
type EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters retrieve the properties and relationships of an accessPackageCatalog object.
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
// AccessPackageCustomWorkflowExtensions provides operations to manage the accessPackageCustomWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementManagementAccessPackageCatalogsItemAccessPackageCustomWorkflowExtensionsRequestBuilder when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageCustomWorkflowExtensions()(*EntitlementManagementAccessPackageCatalogsItemAccessPackageCustomWorkflowExtensionsRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageCustomWorkflowExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceRoles provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesRequestBuilder when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceRoles()(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResources provides operations to manage the accessPackageResources property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesRequestBuilder when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResources()(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackageResourceScopes provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesRequestBuilder when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackageResourceScopes()(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementManagementAccessPackageCatalogsItemAccessPackagesRequestBuilder when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackages()(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal instantiates a new EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) {
    m := &EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder instantiates a new EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomAccessPackageWorkflowExtensions provides operations to manage the customAccessPackageWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsRequestBuilder when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) CustomAccessPackageWorkflowExtensions()(*EntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemCustomAccessPackageWorkflowExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete an accessPackageCatalog.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackagecatalog-delete?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of an accessPackageCatalog object.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageCatalogable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackagecatalog-get?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable), nil
}
// Patch update an existing accessPackageCatalog object to change one or more of its properties, such as the display name or description.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageCatalogable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackagecatalog-update?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable), nil
}
// ToDeleteRequestInformation delete an accessPackageCatalog.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of an accessPackageCatalog object.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update an existing accessPackageCatalog object to change one or more of its properties, such as the display name or description.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageCatalogable, requestConfiguration *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}", m.BaseRequestBuilder.PathParameters)
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder when successful
func (m *EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsAccessPackageCatalogItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
