package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder provides operations to call the unsetReaction method.
type ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilderInternal instantiates a new ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder) {
    m := &ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/unsetReaction", pathParameters),
    }
    return m
}
// NewItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder instantiates a new ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action unsetReaction
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder) Post(ctx context.Context, body ItemChannelsItemMessagesItemUnsetreactionUnsetReactionPostRequestBodyable, requestConfiguration *ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action unsetReaction
// returns a *RequestInformation when successful
func (m *ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChannelsItemMessagesItemUnsetreactionUnsetReactionPostRequestBodyable, requestConfiguration *ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder when successful
func (m *ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder) WithUrl(rawUrl string)(*ItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder) {
    return NewItemChannelsItemMessagesItemUnsetreactionUnsetReactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
