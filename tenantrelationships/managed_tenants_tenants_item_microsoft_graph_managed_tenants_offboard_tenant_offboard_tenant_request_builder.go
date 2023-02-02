package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder provides operations to call the offboardTenant method.
type ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilderInternal instantiates a new OffboardTenantRequestBuilder and sets the default values.
func NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder) {
    m := &ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenants/{tenant%2Did}/microsoft.graph.managedTenants.offboardTenant";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder instantiates a new OffboardTenantRequestBuilder and sets the default values.
func NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilderInternal(urlParams, requestAdapter)
}
// Post carries out the appropriate procedures to remove a managed tenant from the multi-tenant management platform. No relationships, such as commerce and delegate administrative privileges, will be impacted. The only change made by invoking this action is the tenant will be deprovisioned from the multi-tenant management platform.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/managedtenants-tenant-offboardtenant?view=graph-rest-1.0
func (m *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder) Post(ctx context.Context, requestConfiguration *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.Tenantable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.Tenantable), nil
}
// ToPostRequestInformation carries out the appropriate procedures to remove a managed tenant from the multi-tenant management platform. No relationships, such as commerce and delegate administrative privileges, will be impacted. The only change made by invoking this action is the tenant will be deprovisioned from the multi-tenant management platform.
func (m *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsTenantsItemMicrosoftGraphManagedTenantsOffboardTenantOffboardTenantRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
