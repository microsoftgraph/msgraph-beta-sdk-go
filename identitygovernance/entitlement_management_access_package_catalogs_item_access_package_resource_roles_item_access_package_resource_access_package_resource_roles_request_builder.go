package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessPackageResourceRoleId1 provides operations to manage the accessPackageResourceRoles property of the microsoft.graph.accessPackageResource entity.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) ByAccessPackageResourceRoleId1(accessPackageResourceRoleId1 string)(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceRoleId1 != "" {
        urlTplParams["accessPackageResourceRole%2Did1"] = accessPackageResourceRoleId1
    }
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesAccessPackageResourceRoleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderInternal instantiates a new AccessPackageResourceRolesRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) {
    m := &EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource/accessPackageResourceRoles{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder instantiates a new AccessPackageResourceRolesRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) Count()(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesCountRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceRoleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleCollectionResponseable), nil
}
// Post create new navigation property to accessPackageResourceRoles for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation read-only. Nullable. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to accessPackageResourceRoles for identityGovernance
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceRoleable, requestConfiguration *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceRolesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
