package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Attachments()(*MeMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) AttachmentsById(id string)(*MeMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    m := &MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Content()(*MeMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Copy()(*MeMailFoldersItemChildFoldersItemMessagesItemCopyRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for me
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateForward provides operations to call the createForward method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateForward()(*MeMailFoldersItemChildFoldersItemMessagesItemCreateForwardRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in me
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateReply provides operations to call the createReply method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateReply()(*MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateReplyAll provides operations to call the createReplyAll method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateReplyAll()(*MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for me
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Extensions()(*MeMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ExtensionsById(id string)(*MeMailFoldersItemChildFoldersItemMessagesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemMessagesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Forward()(*MeMailFoldersItemChildFoldersItemMessagesItemForwardRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// Mentions provides operations to manage the mentions property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Mentions()(*MeMailFoldersItemChildFoldersItemMessagesItemMentionsRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById provides operations to manage the mentions property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MentionsById(id string)(*MeMailFoldersItemChildFoldersItemMessagesItemMentionsMentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemMessagesItemMentionsMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move provides operations to call the move method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Move()(*MeMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MultiValueExtendedProperties()(*MeMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*MeMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in me
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// Reply provides operations to call the reply method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Reply()(*MeMailFoldersItemChildFoldersItemMessagesItemReplyRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReplyAll provides operations to call the replyAll method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ReplyAll()(*MeMailFoldersItemChildFoldersItemMessagesItemReplyAllRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Send provides operations to call the send method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Send()(*MeMailFoldersItemChildFoldersItemMessagesItemSendRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) SingleValueExtendedProperties()(*MeMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*MeMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubscribe provides operations to call the unsubscribe method.
func (m *MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Unsubscribe()(*MeMailFoldersItemChildFoldersItemMessagesItemUnsubscribeRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
