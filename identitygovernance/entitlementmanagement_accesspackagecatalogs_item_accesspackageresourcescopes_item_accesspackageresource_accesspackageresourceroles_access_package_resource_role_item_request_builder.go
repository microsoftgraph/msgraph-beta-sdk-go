package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResource provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceRole entity.
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) AccessPackageResource()(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesItemAccesspackageresourceAccessPackageResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) {
    m := &EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackageResourceScopes/{accessPackageResourceScope%2Did}/accessPackageResource/accessPackageResourceRoles/{accessPackageResourceRole%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder instantiates a new EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property accessPackageResourceRoles for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageResourceRoleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable), nil
}
// Patch update the navigation property accessPackageResourceRoles in identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageResourceRoleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackageResourceRoles for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property accessPackageResourceRoles in identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder) {
    return NewEntitlementmanagementAccesspackagecatalogsItemAccesspackageresourcescopesItemAccesspackageresourceAccesspackageresourcerolesAccessPackageResourceRoleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
