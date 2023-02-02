package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder provides operations to call the reply method.
type ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilderInternal instantiates a new ReplyRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder) {
    m := &ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/microsoft.graph.reply";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder instantiates a new ReplyRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add an attachment when creating a group post.  This operation limits the size of the attachment you can add to under 3 MB. An attachment can be one of the following types: All these types of attachment resources are derived from the attachmentresource. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/post-post-attachments?view=graph-rest-1.0
func (m *ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder) Post(ctx context.Context, body ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyPostRequestBodyable, requestConfiguration *ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation add an attachment when creating a group post.  This operation limits the size of the attachment you can add to under 3 MB. An attachment can be one of the following types: All these types of attachment resources are derived from the attachmentresource. 
func (m *ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyPostRequestBodyable, requestConfiguration *ItemConversationsItemThreadsItemMicrosoftGraphReplyReplyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
