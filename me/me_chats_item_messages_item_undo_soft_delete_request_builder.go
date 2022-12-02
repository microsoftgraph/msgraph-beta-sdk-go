package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeChatsItemMessagesItemUndoSoftDeleteRequestBuilder provides operations to call the undoSoftDelete method.
type MeChatsItemMessagesItemUndoSoftDeleteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeChatsItemMessagesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeChatsItemMessagesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeChatsItemMessagesItemUndoSoftDeleteRequestBuilderInternal instantiates a new UndoSoftDeleteRequestBuilder and sets the default values.
func NewMeChatsItemMessagesItemUndoSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeChatsItemMessagesItemUndoSoftDeleteRequestBuilder) {
    m := &MeChatsItemMessagesItemUndoSoftDeleteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat%2Did}/messages/{chatMessage%2Did}/microsoft.graph.undoSoftDelete";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeChatsItemMessagesItemUndoSoftDeleteRequestBuilder instantiates a new UndoSoftDeleteRequestBuilder and sets the default values.
func NewMeChatsItemMessagesItemUndoSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeChatsItemMessagesItemUndoSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeChatsItemMessagesItemUndoSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation undo soft deletion of a single message or a message reply in a channel or a chat.
func (m *MeChatsItemMessagesItemUndoSoftDeleteRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *MeChatsItemMessagesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post undo soft deletion of a single message or a message reply in a channel or a chat.
func (m *MeChatsItemMessagesItemUndoSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *MeChatsItemMessagesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
