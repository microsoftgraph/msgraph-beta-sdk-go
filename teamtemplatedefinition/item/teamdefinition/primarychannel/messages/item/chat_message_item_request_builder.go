package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i186bb0602e85f23e12399262c053c6bc56aa80946b3a520147dacceff7bd4855 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel/messages/item/undosoftdelete"
    i4b99b499ae63f60c8123ad6f52c5afce397b1414d259685afb078d47dace9542 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel/messages/item/unsetreaction"
    i61d38a60764c2371b0c1c7f223d9aee07a2d4eda0019b6902d3a3fc079be7392 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel/messages/item/setreaction"
    i6c04788945571ac1d10a9ba2142d40d6a6e42fbcef53a5dd36d46e8bbe9c81e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel/messages/item/hostedcontents"
    ic0e9a2a26218b5f5851049908b03c2e8721b1a32017aa0c09bc15c24596c1e2b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel/messages/item/replies"
    idc176eb9f0dbd718ce19f8e658856d2be83f0bcdb09199c7dbf305da99fbcfc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel/messages/item/softdelete"
    i15a62d65ba9aa9bc3f1a92c3b3b2d22761b07ca6a82c10bab5e6ac483c637607 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel/messages/item/replies/item"
    i6fb33d9f7bc8af482bcb50e3f06f97cd5b11e4c45cac6bd7dfcf5f6ac7cbfdd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel/messages/item/hostedcontents/item"
)

// ChatMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.channel entity.
type ChatMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChatMessageItemRequestBuilderGetQueryParameters a collection of all the messages in the channel. A navigation property. Nullable.
type ChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatMessageItemRequestBuilderGetQueryParameters
}
// ChatMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatMessageItemRequestBuilderInternal instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatMessageItemRequestBuilder) {
    m := &ChatMessageItemRequestBuilder{
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
// NewChatMessageItemRequestBuilder instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property messages for teamTemplateDefinition
func (m *ChatMessageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ChatMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatMessageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ChatMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatMessageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *ChatMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChatMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ChatMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChatMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *ChatMessageItemRequestBuilder) HostedContents()(*i6c04788945571ac1d10a9ba2142d40d6a6e42fbcef53a5dd36d46e8bbe9c81e4.HostedContentsRequestBuilder) {
    return i6c04788945571ac1d10a9ba2142d40d6a6e42fbcef53a5dd36d46e8bbe9c81e4.NewHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById provides operations to manage the hostedContents property of the microsoft.graph.chatMessage entity.
func (m *ChatMessageItemRequestBuilder) HostedContentsById(id string)(*i6fb33d9f7bc8af482bcb50e3f06f97cd5b11e4c45cac6bd7dfcf5f6ac7cbfdd0.ChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent%2Did"] = id
    }
    return i6fb33d9f7bc8af482bcb50e3f06f97cd5b11e4c45cac6bd7dfcf5f6ac7cbfdd0.NewChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in teamTemplateDefinition
func (m *ChatMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, requestConfiguration *ChatMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
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
func (m *ChatMessageItemRequestBuilder) Replies()(*ic0e9a2a26218b5f5851049908b03c2e8721b1a32017aa0c09bc15c24596c1e2b.RepliesRequestBuilder) {
    return ic0e9a2a26218b5f5851049908b03c2e8721b1a32017aa0c09bc15c24596c1e2b.NewRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RepliesById provides operations to manage the replies property of the microsoft.graph.chatMessage entity.
func (m *ChatMessageItemRequestBuilder) RepliesById(id string)(*i15a62d65ba9aa9bc3f1a92c3b3b2d22761b07ca6a82c10bab5e6ac483c637607.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did1"] = id
    }
    return i15a62d65ba9aa9bc3f1a92c3b3b2d22761b07ca6a82c10bab5e6ac483c637607.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SetReaction provides operations to call the setReaction method.
func (m *ChatMessageItemRequestBuilder) SetReaction()(*i61d38a60764c2371b0c1c7f223d9aee07a2d4eda0019b6902d3a3fc079be7392.SetReactionRequestBuilder) {
    return i61d38a60764c2371b0c1c7f223d9aee07a2d4eda0019b6902d3a3fc079be7392.NewSetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftDelete provides operations to call the softDelete method.
func (m *ChatMessageItemRequestBuilder) SoftDelete()(*idc176eb9f0dbd718ce19f8e658856d2be83f0bcdb09199c7dbf305da99fbcfc3.SoftDeleteRequestBuilder) {
    return idc176eb9f0dbd718ce19f8e658856d2be83f0bcdb09199c7dbf305da99fbcfc3.NewSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UndoSoftDelete provides operations to call the undoSoftDelete method.
func (m *ChatMessageItemRequestBuilder) UndoSoftDelete()(*i186bb0602e85f23e12399262c053c6bc56aa80946b3a520147dacceff7bd4855.UndoSoftDeleteRequestBuilder) {
    return i186bb0602e85f23e12399262c053c6bc56aa80946b3a520147dacceff7bd4855.NewUndoSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsetReaction provides operations to call the unsetReaction method.
func (m *ChatMessageItemRequestBuilder) UnsetReaction()(*i4b99b499ae63f60c8123ad6f52c5afce397b1414d259685afb078d47dace9542.UnsetReactionRequestBuilder) {
    return i4b99b499ae63f60c8123ad6f52c5afce397b1414d259685afb078d47dace9542.NewUnsetReactionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
