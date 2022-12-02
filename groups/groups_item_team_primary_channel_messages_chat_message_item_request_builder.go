package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.channel entity.
type GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters a collection of all the messages in the channel. A navigation property. Nullable.
type GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters
}
// GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderInternal instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewGroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) {
    m := &GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/primaryChannel/messages/{chatMessage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewGroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property messages for groups
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of all the messages in the channel. A navigation property. Nullable.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property messages in groups
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property messages for groups
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get a collection of all the messages in the channel. A navigation property. Nullable.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable), nil
}
// HostedContents provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) HostedContents()(*GroupsItemTeamPrimaryChannelMessagesItemHostedContentsRequestBuilder) {
    return NewGroupsItemTeamPrimaryChannelMessagesItemHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) HostedContentsById(id string)(*GroupsItemTeamPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent%2Did"] = id
    }
    return NewGroupsItemTeamPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in groups
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable), nil
}
// Replies provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) Replies()(*GroupsItemTeamPrimaryChannelMessagesItemRepliesRequestBuilder) {
    return NewGroupsItemTeamPrimaryChannelMessagesItemRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RepliesById provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) RepliesById(id string)(*GroupsItemTeamPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did1"] = id
    }
    return NewGroupsItemTeamPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SetReaction provides operations to call the setReaction method.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) SetReaction()(*GroupsItemTeamPrimaryChannelMessagesItemSetReactionRequestBuilder) {
    return NewGroupsItemTeamPrimaryChannelMessagesItemSetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftDelete provides operations to call the softDelete method.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) SoftDelete()(*GroupsItemTeamPrimaryChannelMessagesItemSoftDeleteRequestBuilder) {
    return NewGroupsItemTeamPrimaryChannelMessagesItemSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UndoSoftDelete provides operations to call the undoSoftDelete method.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) UndoSoftDelete()(*GroupsItemTeamPrimaryChannelMessagesItemUndoSoftDeleteRequestBuilder) {
    return NewGroupsItemTeamPrimaryChannelMessagesItemUndoSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsetReaction provides operations to call the unsetReaction method.
func (m *GroupsItemTeamPrimaryChannelMessagesChatMessageItemRequestBuilder) UnsetReaction()(*GroupsItemTeamPrimaryChannelMessagesItemUnsetReactionRequestBuilder) {
    return NewGroupsItemTeamPrimaryChannelMessagesItemUnsetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
