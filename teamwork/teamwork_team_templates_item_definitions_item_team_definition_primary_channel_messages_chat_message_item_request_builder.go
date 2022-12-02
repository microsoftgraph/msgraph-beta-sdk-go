package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.channel entity.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters a collection of all the messages in the channel. A navigation property. Nullable.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetQueryParameters
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderInternal instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/messages/{chatMessage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property messages for teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) HostedContents()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemHostedContentsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) HostedContentsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) Replies()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RepliesById provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) RepliesById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did1"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SetReaction provides operations to call the setReaction method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) SetReaction()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemSetReactionRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemSetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftDelete provides operations to call the softDelete method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) SoftDelete()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemSoftDeleteRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UndoSoftDelete provides operations to call the undoSoftDelete method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) UndoSoftDelete()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemUndoSoftDeleteRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemUndoSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsetReaction provides operations to call the unsetReaction method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) UnsetReaction()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemUnsetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
