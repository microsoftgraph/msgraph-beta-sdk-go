package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder provides operations to call the offboardTenant method.
type ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilderInternal instantiates a new MicrosoftGraphManagedTenantsOffboardTenantRequestBuilder and sets the default values.
func NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder) {
    m := &ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenants/{tenant%2Did}/microsoft.graph.managedTenants.offboardTenant", pathParameters),
    }
    return m
}
// NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder instantiates a new MicrosoftGraphManagedTenantsOffboardTenantRequestBuilder and sets the default values.
func NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action offboardTenant
func (m *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder) Post(ctx context.Context, requestConfiguration *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.Tenantable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.Tenantable), nil
}
// ToPostRequestInformation invoke action offboardTenant
func (m *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
