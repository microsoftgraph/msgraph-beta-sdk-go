package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder provides operations to call the refresh method.
type EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderInternal instantiates a new RefreshRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) {
    m := &EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceScope/accessPackageResource/refresh", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder instantiates a new RefreshRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Azure AD entitlement management, refresh the accessPackageResource object to fetch the latest details for displayName, description, and resourceType from the origin system. For the AadApplication originSystem, this operation also updates the displayName and description for the accessPackageResourceRole. 
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageresource-refresh?view=graph-rest-1.0
func (m *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation in Azure AD entitlement management, refresh the accessPackageResource object to fetch the latest details for displayName, description, and resourceType from the origin system. For the AadApplication originSystem, this operation also updates the displayName and description for the accessPackageResourceRole. 
func (m *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder) {
    return NewEntitlementManagementAccessPackageResourceRoleScopesItemAccessPackageResourceScopeAccessPackageResourceRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
