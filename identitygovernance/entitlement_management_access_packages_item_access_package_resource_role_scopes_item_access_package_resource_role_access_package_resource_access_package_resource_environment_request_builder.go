package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters contains the environment information for the resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports $expand.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters
}
// NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal instantiates a new AccessPackageResourceEnvironmentRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    m := &EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceRole/accessPackageResource/accessPackageResourceEnvironment{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder instantiates a new AccessPackageResourceEnvironmentRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the environment information for the resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports $expand.
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable), nil
}
// ToGetRequestInformation contains the environment information for the resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports $expand.
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementManagementAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
