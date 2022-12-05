package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters replies for a specified message. Supports $expand for channel messages.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderGetQueryParameters
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderInternal instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property replies for teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation replies for a specified message. Supports $expand for channel messages.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property replies in teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property replies for teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get replies for a specified message. Supports $expand for channel messages.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) HostedContents()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemHostedContentsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) HostedContentsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemHostedContentsChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property replies in teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
// SetReaction provides operations to call the setReaction method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) SetReaction()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemSetReactionRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemSetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftDelete provides operations to call the softDelete method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) SoftDelete()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UndoSoftDelete provides operations to call the undoSoftDelete method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) UndoSoftDelete()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemUndoSoftDeleteRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemUndoSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsetReaction provides operations to call the unsetReaction method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesChatMessageItemRequestBuilder) UnsetReaction()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemUnsetReactionRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesItemRepliesItemUnsetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
