package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetQueryParameters
}
// UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Attachments()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.message entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) AttachmentsById(id string)(*UsersItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewUsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    m := &UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewUsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Content()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Copy()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemCopyRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for users
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateForward()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemCreateForwardRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in users
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateReply()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateReplyAll provides operations to call the createReplyAll method.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) CreateReplyAll()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for users
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Extensions()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.message entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ExtensionsById(id string)(*UsersItemMailFoldersItemChildFoldersItemMessagesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Forward()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemForwardRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
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
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Mentions()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemMentionsRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById provides operations to manage the mentions property of the microsoft.graph.message entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MentionsById(id string)(*UsersItemMailFoldersItemChildFoldersItemMessagesItemMentionsMentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemMentionsMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move provides operations to call the move method.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Move()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.message entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in users
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
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
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Reply()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemReplyRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReplyAll provides operations to call the replyAll method.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) ReplyAll()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemReplyAllRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Send provides operations to call the send method.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Send()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemSendRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.message entity.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubscribe provides operations to call the unsubscribe method.
func (m *UsersItemMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) Unsubscribe()(*UsersItemMailFoldersItemChildFoldersItemMessagesItemUnsubscribeRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersItemMessagesItemUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
