package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder provides operations to call the softDelete method.
type ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderInternal instantiates a new MicrosoftGraphSoftDeleteRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder) {
    m := &ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/microsoft.graph.softDelete";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder instantiates a new MicrosoftGraphSoftDeleteRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post delete a single message or a message reply in a channel or a chat.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/chatmessage-softdelete?view=graph-rest-1.0
func (m *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation delete a single message or a message reply in a channel or a chat.
func (m *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsItemMessagesItemRepliesItemMicrosoftGraphSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
