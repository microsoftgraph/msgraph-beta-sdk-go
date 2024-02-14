package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ItemTeamDefinitionChannelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionChannelsRequestBuilderGetQueryParameters retrieve the list of channels in this team.
type ItemTeamDefinitionChannelsRequestBuilderGetQueryParameters struct {
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
// ItemTeamDefinitionChannelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionChannelsRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionChannelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllMessages provides operations to call the allMessages method.
// returns a *ItemTeamDefinitionChannelsAllMessagesRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsRequestBuilder) AllMessages()(*ItemTeamDefinitionChannelsAllMessagesRequestBuilder) {
    return NewItemTeamDefinitionChannelsAllMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByChannelId provides operations to manage the channels property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionChannelsChannelItemRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsRequestBuilder) ByChannelId(channelId string)(*ItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if channelId != "" {
        urlTplParams["channel%2Did"] = channelId
    }
    return NewItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamDefinitionChannelsRequestBuilderInternal instantiates a new ItemTeamDefinitionChannelsRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsRequestBuilder) {
    m := &ItemTeamDefinitionChannelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionChannelsRequestBuilder instantiates a new ItemTeamDefinitionChannelsRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamDefinitionChannelsCountRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsRequestBuilder) Count()(*ItemTeamDefinitionChannelsCountRequestBuilder) {
    return NewItemTeamDefinitionChannelsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the list of channels in this team.
// returns a ChannelCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-list?view=graph-rest-1.0
func (m *ItemTeamDefinitionChannelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChannelCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChannelCollectionResponseable), nil
}
// GetAllMessages provides operations to call the getAllMessages method.
// returns a *ItemTeamDefinitionChannelsGetAllMessagesRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsRequestBuilder) GetAllMessages()(*ItemTeamDefinitionChannelsGetAllMessagesRequestBuilder) {
    return NewItemTeamDefinitionChannelsGetAllMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetAllRetainedMessages provides operations to call the getAllRetainedMessages method.
// returns a *ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsRequestBuilder) GetAllRetainedMessages()(*ItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilder) {
    return NewItemTeamDefinitionChannelsGetAllRetainedMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new channel in a team, as specified in the request body. When you create a channel, the maximum length of the channel's displayName is 50 characters. This is the name that appears to the user in Microsoft Teams. You can add a maximum of 200 members when you create a private channel.
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-post?view=graph-rest-1.0
func (m *ItemTeamDefinitionChannelsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ItemTeamDefinitionChannelsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// ToGetRequestInformation retrieve the list of channels in this team.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionChannelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new channel in a team, as specified in the request body. When you create a channel, the maximum length of the channel's displayName is 50 characters. This is the name that appears to the user in Microsoft Teams. You can add a maximum of 200 members when you create a private channel.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionChannelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ItemTeamDefinitionChannelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamDefinitionChannelsRequestBuilder when successful
func (m *ItemTeamDefinitionChannelsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionChannelsRequestBuilder) {
    return NewItemTeamDefinitionChannelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
