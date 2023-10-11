package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder provides operations to call the undoSoftDelete method.
type ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilderInternal instantiates a new UndoSoftDeleteRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder) {
    m := &ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/messages/{chatMessage%2Did}/undoSoftDelete", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder instantiates a new UndoSoftDeleteRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post undo soft deletion of a single chatMessage or a chat message reply in a channel or a chat. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-undosoftdelete?view=graph-rest-1.0
func (m *ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation undo soft deletion of a single chatMessage or a chat message reply in a channel or a chat. This API is available in the following national cloud deployments.
func (m *ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMessagesItemUndoSoftDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
