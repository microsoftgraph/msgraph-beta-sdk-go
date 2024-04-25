package connections

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ItemGroupsItemMembersRequestBuilder provides operations to manage the members property of the microsoft.graph.externalConnectors.externalGroup entity.
type ItemGroupsItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGroupsItemMembersRequestBuilderGetQueryParameters a member added to an externalGroup. You can add Microsoft Entra users, Microsoft Entra groups, or other externalGroups as members.
type ItemGroupsItemMembersRequestBuilderGetQueryParameters struct {
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
// ItemGroupsItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupsItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGroupsItemMembersRequestBuilderGetQueryParameters
}
// ItemGroupsItemMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupsItemMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByIdentityId provides operations to manage the members property of the microsoft.graph.externalConnectors.externalGroup entity.
// returns a *ItemGroupsItemMembersIdentityItemRequestBuilder when successful
func (m *ItemGroupsItemMembersRequestBuilder) ByIdentityId(identityId string)(*ItemGroupsItemMembersIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityId != "" {
        urlTplParams["identity%2Did"] = identityId
    }
    return NewItemGroupsItemMembersIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemGroupsItemMembersRequestBuilderInternal instantiates a new ItemGroupsItemMembersRequestBuilder and sets the default values.
func NewItemGroupsItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupsItemMembersRequestBuilder) {
    m := &ItemGroupsItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/connections/{externalConnection%2Did}/groups/{externalGroup%2Did}/members{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemGroupsItemMembersRequestBuilder instantiates a new ItemGroupsItemMembersRequestBuilder and sets the default values.
func NewItemGroupsItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupsItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupsItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemGroupsItemMembersCountRequestBuilder when successful
func (m *ItemGroupsItemMembersRequestBuilder) Count()(*ItemGroupsItemMembersCountRequestBuilder) {
    return NewItemGroupsItemMembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a member added to an externalGroup. You can add Microsoft Entra users, Microsoft Entra groups, or other externalGroups as members.
// returns a IdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGroupsItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGroupsItemMembersRequestBuilderGetRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.IdentityCollectionResponseable, error) {
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
// Post create new navigation property to members for connections
// returns a Identityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGroupsItemMembersRequestBuilder) Post(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, requestConfiguration *ItemGroupsItemMembersRequestBuilderPostRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, error) {
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
func (m *ItemGroupsItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGroupsItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to members for connections
// returns a *RequestInformation when successful
func (m *ItemGroupsItemMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, requestConfiguration *ItemGroupsItemMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGroupsItemMembersRequestBuilder when successful
func (m *ItemGroupsItemMembersRequestBuilder) WithUrl(rawUrl string)(*ItemGroupsItemMembersRequestBuilder) {
    return NewItemGroupsItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
