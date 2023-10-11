package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder provides operations to call the unassignTag method.
type ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilderInternal instantiates a new MicrosoftGraphManagedTenantsUnassignTagRequestBuilder and sets the default values.
func NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder) {
    m := &ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenantTags/{tenantTag%2Did}/microsoft.graph.managedTenants.unassignTag", pathParameters),
    }
    return m
}
// NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder instantiates a new MicrosoftGraphManagedTenantsUnassignTagRequestBuilder and sets the default values.
func NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Post un-assigns the tenant tag from the specified managed tenants. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-tenanttag-unassigntag?view=graph-rest-1.0
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder) Post(ctx context.Context, body ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBodyable, requestConfiguration *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantTagable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantTagable), nil
}
// ToPostRequestInformation un-assigns the tenant tag from the specified managed tenants. This API is available in the following national cloud deployments.
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBodyable, requestConfiguration *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder) {
    return NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
