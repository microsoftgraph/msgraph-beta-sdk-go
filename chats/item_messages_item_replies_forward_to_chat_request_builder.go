package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMessagesItemRepliesForwardToChatRequestBuilder provides operations to call the forwardToChat method.
type ItemMessagesItemRepliesForwardToChatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMessagesItemRepliesForwardToChatRequestBuilderInternal instantiates a new ItemMessagesItemRepliesForwardToChatRequestBuilder and sets the default values.
func NewItemMessagesItemRepliesForwardToChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesItemRepliesForwardToChatRequestBuilder) {
    m := &ItemMessagesItemRepliesForwardToChatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/messages/{chatMessage%2Did}/replies/forwardToChat", pathParameters),
    }
    return m
}
// NewItemMessagesItemRepliesForwardToChatRequestBuilder instantiates a new ItemMessagesItemRepliesForwardToChatRequestBuilder and sets the default values.
func NewItemMessagesItemRepliesForwardToChatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesItemRepliesForwardToChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMessagesItemRepliesForwardToChatRequestBuilderInternal(urlParams, requestAdapter)
}
// Post forward a chat message, a channel message, or a channel message reply to a chat.
// Deprecated: This method is obsolete. Use PostAsForwardToChatPostResponse instead.
// returns a ItemMessagesItemRepliesForwardToChatResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-forwardtochat?view=graph-rest-beta
func (m *ItemMessagesItemRepliesForwardToChatRequestBuilder) Post(ctx context.Context, body ItemMessagesItemRepliesForwardToChatPostRequestBodyable, requestConfiguration *ItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration)(ItemMessagesItemRepliesForwardToChatResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMessagesItemRepliesForwardToChatResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMessagesItemRepliesForwardToChatResponseable), nil
}
// PostAsForwardToChatPostResponse forward a chat message, a channel message, or a channel message reply to a chat.
// returns a ItemMessagesItemRepliesForwardToChatPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-forwardtochat?view=graph-rest-beta
func (m *ItemMessagesItemRepliesForwardToChatRequestBuilder) PostAsForwardToChatPostResponse(ctx context.Context, body ItemMessagesItemRepliesForwardToChatPostRequestBodyable, requestConfiguration *ItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration)(ItemMessagesItemRepliesForwardToChatPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMessagesItemRepliesForwardToChatPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMessagesItemRepliesForwardToChatPostResponseable), nil
}
// ToPostRequestInformation forward a chat message, a channel message, or a channel message reply to a chat.
// returns a *RequestInformation when successful
func (m *ItemMessagesItemRepliesForwardToChatRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMessagesItemRepliesForwardToChatPostRequestBodyable, requestConfiguration *ItemMessagesItemRepliesForwardToChatRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMessagesItemRepliesForwardToChatRequestBuilder when successful
func (m *ItemMessagesItemRepliesForwardToChatRequestBuilder) WithUrl(rawUrl string)(*ItemMessagesItemRepliesForwardToChatRequestBuilder) {
    return NewItemMessagesItemRepliesForwardToChatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
