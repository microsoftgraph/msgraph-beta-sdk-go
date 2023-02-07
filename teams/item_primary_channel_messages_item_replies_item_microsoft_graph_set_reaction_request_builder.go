package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder provides operations to call the setReaction method.
type ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderInternal instantiates a new MicrosoftGraphSetReactionRequestBuilder and sets the default values.
func NewItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder) {
    m := &ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/microsoft.graph.setReaction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder instantiates a new MicrosoftGraphSetReactionRequestBuilder and sets the default values.
func NewItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReaction
func (m *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder) Post(ctx context.Context, body ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyable, requestConfiguration *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action setReaction
func (m *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyable, requestConfiguration *ItemPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
