package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder provides operations to call the setReaction method.
type ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilderInternal instantiates a new SetReactionRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder) {
    m := &ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/microsoft.graph.setReaction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder instantiates a new SetReactionRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReaction
func (m *ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder) Post(ctx context.Context, body ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyable, requestConfiguration *ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyable, requestConfiguration *ItemTeamPrimaryChannelMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
