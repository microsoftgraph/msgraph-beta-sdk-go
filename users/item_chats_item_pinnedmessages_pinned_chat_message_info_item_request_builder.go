package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
type ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderGetQueryParameters a collection of all the pinned messages in the chat. Nullable.
type ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderGetQueryParameters
}
// ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderInternal instantiates a new ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder and sets the default values.
func NewItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) {
    m := &ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/pinnedMessages/{pinnedChatMessageInfo%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder instantiates a new ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder and sets the default values.
func NewItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property pinnedMessages for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get a collection of all the pinned messages in the chat. Nullable.
// returns a PinnedChatMessageInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PinnedChatMessageInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePinnedChatMessageInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PinnedChatMessageInfoable), nil
}
// Message provides operations to manage the message property of the microsoft.graph.pinnedChatMessageInfo entity.
// returns a *ItemChatsItemPinnedmessagesItemMessageRequestBuilder when successful
func (m *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) Message()(*ItemChatsItemPinnedmessagesItemMessageRequestBuilder) {
    return NewItemChatsItemPinnedmessagesItemMessageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property pinnedMessages in users
// returns a PinnedChatMessageInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PinnedChatMessageInfoable, requestConfiguration *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PinnedChatMessageInfoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePinnedChatMessageInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PinnedChatMessageInfoable), nil
}
// ToDeleteRequestInformation delete navigation property pinnedMessages for users
// returns a *RequestInformation when successful
func (m *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of all the pinned messages in the chat. Nullable.
// returns a *RequestInformation when successful
func (m *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property pinnedMessages in users
// returns a *RequestInformation when successful
func (m *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PinnedChatMessageInfoable, requestConfiguration *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder when successful
func (m *ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder) {
    return NewItemChatsItemPinnedmessagesPinnedChatMessageInfoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
