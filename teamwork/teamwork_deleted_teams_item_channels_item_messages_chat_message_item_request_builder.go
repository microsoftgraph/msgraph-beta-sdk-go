package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.channel entity.
type TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderGetQueryParameters a collection of all the messages in the channel. A navigation property. Nullable.
type TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderGetQueryParameters
}
// TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderInternal instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewTeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) {
    m := &TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewTeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property messages for teamwork
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in teamwork
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property messages for teamwork
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) HostedContents()(*TeamworkDeletedTeamsItemChannelsItemMessagesItemHostedContentsRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesItemHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) HostedContentsById(id string)(*TeamworkDeletedTeamsItemChannelsItemMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent%2Did"] = id
    }
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in teamwork
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) Replies()(*TeamworkDeletedTeamsItemChannelsItemMessagesItemRepliesRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesItemRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RepliesById provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) RepliesById(id string)(*TeamworkDeletedTeamsItemChannelsItemMessagesItemRepliesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did1"] = id
    }
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesItemRepliesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SetReaction provides operations to call the setReaction method.
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) SetReaction()(*TeamworkDeletedTeamsItemChannelsItemMessagesItemSetReactionRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesItemSetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftDelete provides operations to call the softDelete method.
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) SoftDelete()(*TeamworkDeletedTeamsItemChannelsItemMessagesItemSoftDeleteRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesItemSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UndoSoftDelete provides operations to call the undoSoftDelete method.
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) UndoSoftDelete()(*TeamworkDeletedTeamsItemChannelsItemMessagesItemUndoSoftDeleteRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesItemUndoSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsetReaction provides operations to call the unsetReaction method.
func (m *TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) UnsetReaction()(*TeamworkDeletedTeamsItemChannelsItemMessagesItemUnsetReactionRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesItemUnsetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
