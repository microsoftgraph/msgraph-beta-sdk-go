package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsTenantgroupsTenantGroupsRequestBuilder provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsTenantgroupsTenantGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsTenantgroupsTenantGroupsRequestBuilderGetQueryParameters get a list of the tenantGroup objects and their properties.
type ManagedtenantsTenantgroupsTenantGroupsRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsTenantgroupsTenantGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsTenantgroupsTenantGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsTenantgroupsTenantGroupsRequestBuilderGetQueryParameters
}
// ManagedtenantsTenantgroupsTenantGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsTenantgroupsTenantGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTenantGroupId provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsTenantgroupsTenantGroupItemRequestBuilder when successful
func (m *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) ByTenantGroupId(tenantGroupId string)(*ManagedtenantsTenantgroupsTenantGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tenantGroupId != "" {
        urlTplParams["tenantGroup%2Did"] = tenantGroupId
    }
    return NewManagedtenantsTenantgroupsTenantGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsTenantgroupsTenantGroupsRequestBuilderInternal instantiates a new ManagedtenantsTenantgroupsTenantGroupsRequestBuilder and sets the default values.
func NewManagedtenantsTenantgroupsTenantGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) {
    m := &ManagedtenantsTenantgroupsTenantGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenantGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsTenantgroupsTenantGroupsRequestBuilder instantiates a new ManagedtenantsTenantgroupsTenantGroupsRequestBuilder and sets the default values.
func NewManagedtenantsTenantgroupsTenantGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsTenantgroupsTenantGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsTenantgroupsCountRequestBuilder when successful
func (m *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) Count()(*ManagedtenantsTenantgroupsCountRequestBuilder) {
    return NewManagedtenantsTenantgroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the tenantGroup objects and their properties.
// returns a TenantGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-managedtenant-list-tenantgroups?view=graph-rest-beta
func (m *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsTenantgroupsTenantGroupsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupCollectionResponseable), nil
}
// MicrosoftGraphManagedTenantsTenantSearch provides operations to call the tenantSearch method.
// returns a *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder when successful
func (m *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) MicrosoftGraphManagedTenantsTenantSearch()(*ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) {
    return NewManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to tenantGroups for tenantRelationships
// returns a TenantGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable, requestConfiguration *ManagedtenantsTenantgroupsTenantGroupsRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable), nil
}
// ToGetRequestInformation get a list of the tenantGroup objects and their properties.
// returns a *RequestInformation when successful
func (m *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsTenantgroupsTenantGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to tenantGroups for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable, requestConfiguration *ManagedtenantsTenantgroupsTenantGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder when successful
func (m *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) {
    return NewManagedtenantsTenantgroupsTenantGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
