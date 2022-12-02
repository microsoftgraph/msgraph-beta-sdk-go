package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder provides operations to manage the accessPackageResourceEnvironment property of the microsoft.graph.accessPackageResource entity.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetQueryParameters
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal instantiates a new AccessPackageResourceEnvironmentRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceScope/accessPackageResource/accessPackageResourceEnvironment{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder instantiates a new AccessPackageResourceEnvironmentRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceEnvironmentable), nil
}
