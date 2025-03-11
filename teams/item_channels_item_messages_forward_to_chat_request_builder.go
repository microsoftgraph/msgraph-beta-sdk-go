package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemChannelsItemMessagesForwardToChatRequestBuilder provides operations to call the forwardToChat method.
type ItemChannelsItemMessagesForwardToChatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChannelsItemMessagesForwardToChatRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemMessagesForwardToChatRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChannelsItemMessagesForwardToChatRequestBuilderInternal instantiates a new ItemChannelsItemMessagesForwardToChatRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesForwardToChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemMessagesForwardToChatRequestBuilder) {
    m := &ItemChannelsItemMessagesForwardToChatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/channels/{channel%2Did}/messages/forwardToChat", pathParameters),
    }
    return m
}
// NewItemChannelsItemMessagesForwardToChatRequestBuilder instantiates a new ItemChannelsItemMessagesForwardToChatRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesForwardToChatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemMessagesForwardToChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChannelsItemMessagesForwardToChatRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action forwardToChat
// Deprecated: This method is obsolete. Use PostAsForwardToChatPostResponse instead.
// returns a ItemChannelsItemMessagesForwardToChatResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChannelsItemMessagesForwardToChatRequestBuilder) Post(ctx context.Context, body ItemChannelsItemMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemChannelsItemMessagesForwardToChatRequestBuilderPostRequestConfiguration)(ItemChannelsItemMessagesForwardToChatResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemChannelsItemMessagesForwardToChatResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemChannelsItemMessagesForwardToChatResponseable), nil
}
// PostAsForwardToChatPostResponse invoke action forwardToChat
// returns a ItemChannelsItemMessagesForwardToChatPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChannelsItemMessagesForwardToChatRequestBuilder) PostAsForwardToChatPostResponse(ctx context.Context, body ItemChannelsItemMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemChannelsItemMessagesForwardToChatRequestBuilderPostRequestConfiguration)(ItemChannelsItemMessagesForwardToChatPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemChannelsItemMessagesForwardToChatPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemChannelsItemMessagesForwardToChatPostResponseable), nil
}
// ToPostRequestInformation invoke action forwardToChat
// returns a *RequestInformation when successful
func (m *ItemChannelsItemMessagesForwardToChatRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChannelsItemMessagesForwardToChatPostRequestBodyable, requestConfiguration *ItemChannelsItemMessagesForwardToChatRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChannelsItemMessagesForwardToChatRequestBuilder when successful
func (m *ItemChannelsItemMessagesForwardToChatRequestBuilder) WithUrl(rawUrl string)(*ItemChannelsItemMessagesForwardToChatRequestBuilder) {
    return NewItemChannelsItemMessagesForwardToChatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
