package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder provides operations to call the assignTag method.
type ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilderInternal instantiates a new MicrosoftGraphManagedTenantsAssignTagRequestBuilder and sets the default values.
func NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder) {
    m := &ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenantTags/{tenantTag%2Did}/microsoft.graph.managedTenants.assignTag";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder instantiates a new MicrosoftGraphManagedTenantsAssignTagRequestBuilder and sets the default values.
func NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Post assign the tenant tag to the specified managed tenants.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/managedtenants-tenanttag-assigntag?view=graph-rest-1.0
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder) Post(ctx context.Context, body ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagAssignTagPostRequestBodyable, requestConfiguration *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantTagable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantTagable), nil
}
// ToPostRequestInformation assign the tenant tag to the specified managed tenants.
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagAssignTagPostRequestBodyable, requestConfiguration *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsAssignTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
