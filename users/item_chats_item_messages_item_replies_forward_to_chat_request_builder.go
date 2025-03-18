package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder provides operations to call the forwardToChat method.
type ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemMessagesItemRepliesForwardToChatRequestBuilderInternal instantiates a new ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder and sets the default values.
func NewItemChatsItemMessagesItemRepliesForwardToChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder) {
    m := &ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/messages/{chatMessage%2Did}/replies/forwardToChat", pathParameters),
    }
    return m
}
// NewItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder instantiates a new ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder and sets the default values.
func NewItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemMessagesItemRepliesForwardToChatRequestBuilderInternal(urlParams, requestAdapter)
}
// Post forward a chat message, a channel message, or a channel message reply to a chat.
// Deprecated: This method is obsolete. Use PostAsForwardToChatPostResponse instead.
// returns a ItemChatsItemMessagesItemRepliesForwardToChatResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-forwardtochat?view=graph-rest-beta
func (m *ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder) Post(ctx context.Context, body ItemChatsItemMessagesItemRepliesForwardToChatPostRequestBodyable, requestConfiguration *ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration)(ItemChatsItemMessagesItemRepliesForwardToChatResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemChatsItemMessagesItemRepliesForwardToChatResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemChatsItemMessagesItemRepliesForwardToChatResponseable), nil
}
// PostAsForwardToChatPostResponse forward a chat message, a channel message, or a channel message reply to a chat.
// returns a ItemChatsItemMessagesItemRepliesForwardToChatPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-forwardtochat?view=graph-rest-beta
func (m *ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder) PostAsForwardToChatPostResponse(ctx context.Context, body ItemChatsItemMessagesItemRepliesForwardToChatPostRequestBodyable, requestConfiguration *ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration)(ItemChatsItemMessagesItemRepliesForwardToChatPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemChatsItemMessagesItemRepliesForwardToChatPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemChatsItemMessagesItemRepliesForwardToChatPostResponseable), nil
}
// ToPostRequestInformation forward a chat message, a channel message, or a channel message reply to a chat.
// returns a *RequestInformation when successful
func (m *ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemMessagesItemRepliesForwardToChatPostRequestBodyable, requestConfiguration *ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder when successful
func (m *ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder) {
    return NewItemChatsItemMessagesItemRepliesForwardToChatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
