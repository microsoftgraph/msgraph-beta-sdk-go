package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters contains the environment information for the resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports $expand.
type EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters
}
// NewEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal instantiates a new EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    m := &EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResources/{accessPackageResource%2Did}/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource/accessPackageResourceEnvironment{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder instantiates a new EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the environment information for the resource. This environment can be set using either the @odata.bind annotation or the environment's originId. Supports $expand.
// returns a AccessPackageResourceEnvironmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder when successful
func (m *EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourcesItemAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
