package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ConnectionsItemGroupsItemMembersRequestBuilder provides operations to manage the members property of the microsoft.graph.externalConnectors.externalGroup entity.
type ConnectionsItemGroupsItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectionsItemGroupsItemMembersRequestBuilderGetQueryParameters a member added to an externalGroup. You can add Microsoft Entra users, Microsoft Entra groups, or other externalGroups as members.
type ConnectionsItemGroupsItemMembersRequestBuilderGetQueryParameters struct {
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
// ConnectionsItemGroupsItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsItemGroupsItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectionsItemGroupsItemMembersRequestBuilderGetQueryParameters
}
// ConnectionsItemGroupsItemMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsItemGroupsItemMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByIdentityId provides operations to manage the members property of the microsoft.graph.externalConnectors.externalGroup entity.
// returns a *ConnectionsItemGroupsItemMembersIdentityItemRequestBuilder when successful
func (m *ConnectionsItemGroupsItemMembersRequestBuilder) ByIdentityId(identityId string)(*ConnectionsItemGroupsItemMembersIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityId != "" {
        urlTplParams["identity%2Did"] = identityId
    }
    return NewConnectionsItemGroupsItemMembersIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConnectionsItemGroupsItemMembersRequestBuilderInternal instantiates a new ConnectionsItemGroupsItemMembersRequestBuilder and sets the default values.
func NewConnectionsItemGroupsItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemGroupsItemMembersRequestBuilder) {
    m := &ConnectionsItemGroupsItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/connections/{externalConnection%2Did}/groups/{externalGroup%2Did}/members{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConnectionsItemGroupsItemMembersRequestBuilder instantiates a new ConnectionsItemGroupsItemMembersRequestBuilder and sets the default values.
func NewConnectionsItemGroupsItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemGroupsItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectionsItemGroupsItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConnectionsItemGroupsItemMembersCountRequestBuilder when successful
func (m *ConnectionsItemGroupsItemMembersRequestBuilder) Count()(*ConnectionsItemGroupsItemMembersCountRequestBuilder) {
    return NewConnectionsItemGroupsItemMembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a member added to an externalGroup. You can add Microsoft Entra users, Microsoft Entra groups, or other externalGroups as members.
// returns a IdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectionsItemGroupsItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectionsItemGroupsItemMembersRequestBuilderGetRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.IdentityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateIdentityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.IdentityCollectionResponseable), nil
}
// Post create a new externalGroupMember object.
// returns a Identityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/externalconnectors-externalgroup-post-members?view=graph-rest-1.0
func (m *ConnectionsItemGroupsItemMembersRequestBuilder) Post(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, requestConfiguration *ConnectionsItemGroupsItemMembersRequestBuilderPostRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable), nil
}
// ToGetRequestInformation a member added to an externalGroup. You can add Microsoft Entra users, Microsoft Entra groups, or other externalGroups as members.
// returns a *RequestInformation when successful
func (m *ConnectionsItemGroupsItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectionsItemGroupsItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new externalGroupMember object.
// returns a *RequestInformation when successful
func (m *ConnectionsItemGroupsItemMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, requestConfiguration *ConnectionsItemGroupsItemMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConnectionsItemGroupsItemMembersRequestBuilder when successful
func (m *ConnectionsItemGroupsItemMembersRequestBuilder) WithUrl(rawUrl string)(*ConnectionsItemGroupsItemMembersRequestBuilder) {
    return NewConnectionsItemGroupsItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
