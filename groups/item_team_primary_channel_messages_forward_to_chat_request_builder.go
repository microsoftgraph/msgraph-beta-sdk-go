package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder provides operations to call the forwardToChat method.
type ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamPrimaryChannelMessagesForwardToChatRequestBuilderInternal instantiates a new ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelMessagesForwardToChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder) {
    m := &ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel/messages/forwardToChat", pathParameters),
    }
    return m
}
// NewItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder instantiates a new ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimaryChannelMessagesForwardToChatRequestBuilderInternal(urlParams, requestAdapter)
}
// Post forward a chat message, a channel message, or a channel message reply to a chat.
// Deprecated: This method is obsolete. Use PostAsForwardToChatPostResponse instead.
// returns a ItemTeamPrimaryChannelMessagesForwardToChatResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-forwardtochat?view=graph-rest-beta
func (m *ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder) Post(ctx context.Context, body ItemTeamPrimaryChannelMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilderPostRequestConfiguration)(ItemTeamPrimaryChannelMessagesForwardToChatResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamPrimaryChannelMessagesForwardToChatResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamPrimaryChannelMessagesForwardToChatResponseable), nil
}
// PostAsForwardToChatPostResponse forward a chat message, a channel message, or a channel message reply to a chat.
// returns a ItemTeamPrimaryChannelMessagesForwardToChatPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-forwardtochat?view=graph-rest-beta
func (m *ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder) PostAsForwardToChatPostResponse(ctx context.Context, body ItemTeamPrimaryChannelMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilderPostRequestConfiguration)(ItemTeamPrimaryChannelMessagesForwardToChatPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamPrimaryChannelMessagesForwardToChatPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamPrimaryChannelMessagesForwardToChatPostResponseable), nil
}
// ToPostRequestInformation forward a chat message, a channel message, or a channel message reply to a chat.
// returns a *RequestInformation when successful
func (m *ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamPrimaryChannelMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder when successful
func (m *ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder) WithUrl(rawUrl string)(*ItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder) {
    return NewItemTeamPrimaryChannelMessagesForwardToChatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
