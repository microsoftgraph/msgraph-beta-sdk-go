package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.channel entity.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters a collection of all the messages in the channel. A navigation property. Nullable.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters
}
// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderInternal instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) {
    m := &TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/messages/{chatMessage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property messages for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property messages for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) HostedContents()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemHostedContentsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) HostedContentsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) Replies()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemRepliesRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RepliesById provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) RepliesById(id string)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did1"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SetReaction provides operations to call the setReaction method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) SetReaction()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemSetReactionRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemSetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftDelete provides operations to call the softDelete method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) SoftDelete()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemSoftDeleteRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UndoSoftDelete provides operations to call the undoSoftDelete method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) UndoSoftDelete()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUndoSoftDeleteRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUndoSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsetReaction provides operations to call the unsetReaction method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) UnsetReaction()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
