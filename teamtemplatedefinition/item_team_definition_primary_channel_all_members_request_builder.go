package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder provides operations to manage the allMembers property of the microsoft.graph.channel entity.
type ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderGetQueryParameters a collection of membership records associated with the channel. It includes both direct and indirect members of shared channels.
type ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderGetQueryParameters struct {
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
// ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Add provides operations to call the add method.
// returns a *ItemTeamDefinitionPrimaryChannelAllMembersAddRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) Add()(*ItemTeamDefinitionPrimaryChannelAllMembersAddRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelAllMembersAddRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByConversationMemberId provides operations to manage the allMembers property of the microsoft.graph.channel entity.
// returns a *ItemTeamDefinitionPrimaryChannelAllMembersConversationMemberItemRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) ByConversationMemberId(conversationMemberId string)(*ItemTeamDefinitionPrimaryChannelAllMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if conversationMemberId != "" {
        urlTplParams["conversationMember%2Did"] = conversationMemberId
    }
    return NewItemTeamDefinitionPrimaryChannelAllMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderInternal instantiates a new ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) {
    m := &ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/allMembers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder instantiates a new ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder and sets the default values.
func NewItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamDefinitionPrimaryChannelAllMembersCountRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) Count()(*ItemTeamDefinitionPrimaryChannelAllMembersCountRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelAllMembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of membership records associated with the channel. It includes both direct and indirect members of shared channels.
// returns a ConversationMemberCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConversationMemberCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberCollectionResponseable), nil
}
// Post create new navigation property to allMembers for teamTemplateDefinition
// returns a ConversationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberable, requestConfiguration *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConversationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberable), nil
}
// Remove provides operations to call the remove method.
// returns a *ItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) Remove()(*ItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelAllMembersRemoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation a collection of membership records associated with the channel. It includes both direct and indirect members of shared channels.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to allMembers for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberable, requestConfiguration *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder when successful
func (m *ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelAllMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
