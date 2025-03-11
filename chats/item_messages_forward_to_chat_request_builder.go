package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMessagesForwardToChatRequestBuilder provides operations to call the forwardToChat method.
type ItemMessagesForwardToChatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMessagesForwardToChatRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesForwardToChatRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMessagesForwardToChatRequestBuilderInternal instantiates a new ItemMessagesForwardToChatRequestBuilder and sets the default values.
func NewItemMessagesForwardToChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesForwardToChatRequestBuilder) {
    m := &ItemMessagesForwardToChatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/messages/forwardToChat", pathParameters),
    }
    return m
}
// NewItemMessagesForwardToChatRequestBuilder instantiates a new ItemMessagesForwardToChatRequestBuilder and sets the default values.
func NewItemMessagesForwardToChatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesForwardToChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMessagesForwardToChatRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action forwardToChat
// Deprecated: This method is obsolete. Use PostAsForwardToChatPostResponse instead.
// returns a ItemMessagesForwardToChatResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMessagesForwardToChatRequestBuilder) Post(ctx context.Context, body ItemMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemMessagesForwardToChatRequestBuilderPostRequestConfiguration)(ItemMessagesForwardToChatResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMessagesForwardToChatResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMessagesForwardToChatResponseable), nil
}
// PostAsForwardToChatPostResponse invoke action forwardToChat
// returns a ItemMessagesForwardToChatPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMessagesForwardToChatRequestBuilder) PostAsForwardToChatPostResponse(ctx context.Context, body ItemMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemMessagesForwardToChatRequestBuilderPostRequestConfiguration)(ItemMessagesForwardToChatPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMessagesForwardToChatPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMessagesForwardToChatPostResponseable), nil
}
// ToPostRequestInformation invoke action forwardToChat
// returns a *RequestInformation when successful
func (m *ItemMessagesForwardToChatRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemMessagesForwardToChatRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMessagesForwardToChatRequestBuilder when successful
func (m *ItemMessagesForwardToChatRequestBuilder) WithUrl(rawUrl string)(*ItemMessagesForwardToChatRequestBuilder) {
    return NewItemMessagesForwardToChatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
