package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder provides operations to call the setReaction method.
type DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderInternal instantiates a new MicrosoftGraphSetReactionRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/microsoft.graph.setReaction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder instantiates a new MicrosoftGraphSetReactionRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReaction
func (m *DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder) Post(ctx context.Context, body DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyable, requestConfiguration *DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderPostRequestConfiguration)(error) {
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
func (m *DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionSetReactionPostRequestBodyable, requestConfiguration *DeletedTeamsItemChannelsItemMessagesItemRepliesItemMicrosoftGraphSetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
