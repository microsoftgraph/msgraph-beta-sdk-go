package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i05e8b869ffc519fee49d0ac259a1f1eea9becc6ebc1c5a37c2f85196b6ee8ee3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/createforward"
    i062177ea9220fd8df21c532fe25b898767dfd98ec54bcfa1a849d01a69301aad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/replyall"
    i0c991f94222b3272f1430d5abdc8d769cea7a432edd8c195e8ed7efd98dbe501 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/singlevalueextendedproperties"
    i0fccf38aca2f46144aa036d0c0fa6ee24e01b5a7a3993c6f341c44efcb5232f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/calendarsharingmessage"
    i17a9af3ebcd424a0f4a7360aa860cfe25fccbd6f75d1f08e88b91b99baffe208 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/reply"
    i272e002e7b409e74d18f258c8a21d69615f41c0248b18a4763a968e49aa9f78c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/forward"
    i36ba3a15f9b3d43bd01853e14f11e541cd30846655c1041d1198a48d07a44b2f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/eventmessagerequest"
    i3bfc69f268f92d3eb51d630e84bcf0e900efda493e2ca9cded7df23eadb96f04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/createreplyall"
    i537e97523a2d84f6e4e5d67a6432c411a6c3d123670344981076f6bca0acd27f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/multivalueextendedproperties"
    i6757d54d5c87cd028537643e1cb39a47e345c91ea192c6abd02e016c98bbc3d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/extensions"
    i72b0f6daa0bcef81a7a85e98d43a239d41374d9126783782e12898c44e092655 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/copy"
    i87147cd2592bb7a4f35c0388ca4cd3212b1247c3d08cee9e1e475c3ba0f8ed64 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/send"
    i9e5b9792ff40a6736ff612b9493f83c37c80ff3eda2ef70fc12193c0dca8b5fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/move"
    ib049a164b38d6ba9b4a3b7192971ad5e4b114c289241f24f85456ac7681289f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/createreply"
    ibc3856e7bed2bea8783ad4cf102bbf5295cc9b7cc63847236c2fcb7a7d37d1b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/attachments"
    ibd12efd1c58ab297076323c8b68a12f9e5cfa0fe08af58e0c140c589be1fa326 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/value"
    idcd1232c5c17ea39a25044877c92f642df5dd303671e6efaa847bc930d661f82 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/unsubscribe"
    idd4dd625416cd52060f95e8ad108f1f8964fca2bf3f61a6253fbf7068c91d62d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/mentions"
    i671b494222abfe980bb8d667f8d3d2047da5dddc998c20eab8bb67f1d8672de6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/attachments/item"
    i9e4259eea1e8905b783d4fc86aede2440ef9a47ea5f2d15815788ac67eba3d69 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/extensions/item"
    id8651f068bd1a845ecd012a2bfde14061474c28c5020e8285e3b05b6f3fae1fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/mentions/item"
    ie7ad4a32eeb50c09e552caa2785b20f0fd2063c38f777f192847be1c617ca319 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/multivalueextendedproperties/item"
    ifa2e1a49bbba48dd1867ab7775f7bc0212fe78345f75addbc5353906df09475c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/singlevalueextendedproperties/item"
)

// MessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessageItemRequestBuilderDeleteOptions options for Delete
type MessageItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// MessageItemRequestBuilderGetOptions options for Get
type MessageItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MessageItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// MessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MessageItemRequestBuilderPatchOptions options for Patch
type MessageItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Attachments the attachments property
func (m *MessageItemRequestBuilder) Attachments()(*ibc3856e7bed2bea8783ad4cf102bbf5295cc9b7cc63847236c2fcb7a7d37d1b9.AttachmentsRequestBuilder) {
    return ibc3856e7bed2bea8783ad4cf102bbf5295cc9b7cc63847236c2fcb7a7d37d1b9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*i671b494222abfe980bb8d667f8d3d2047da5dddc998c20eab8bb67f1d8672de6.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i671b494222abfe980bb8d667f8d3d2047da5dddc998c20eab8bb67f1d8672de6.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarSharingMessage the calendarSharingMessage property
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*i0fccf38aca2f46144aa036d0c0fa6ee24e01b5a7a3993c6f341c44efcb5232f7.CalendarSharingMessageRequestBuilder) {
    return i0fccf38aca2f46144aa036d0c0fa6ee24e01b5a7a3993c6f341c44efcb5232f7.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}/messages/{message%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the Content property
func (m *MessageItemRequestBuilder) Content()(*ibd12efd1c58ab297076323c8b68a12f9e5cfa0fe08af58e0c140c589be1fa326.ContentRequestBuilder) {
    return ibd12efd1c58ab297076323c8b68a12f9e5cfa0fe08af58e0c140c589be1fa326.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *MessageItemRequestBuilder) Copy()(*i72b0f6daa0bcef81a7a85e98d43a239d41374d9126783782e12898c44e092655.CopyRequestBuilder) {
    return i72b0f6daa0bcef81a7a85e98d43a239d41374d9126783782e12898c44e092655.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for me
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation(options *MessageItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateForward the createForward property
func (m *MessageItemRequestBuilder) CreateForward()(*i05e8b869ffc519fee49d0ac259a1f1eea9becc6ebc1c5a37c2f85196b6ee8ee3.CreateForwardRequestBuilder) {
    return i05e8b869ffc519fee49d0ac259a1f1eea9becc6ebc1c5a37c2f85196b6ee8ee3.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) CreateGetRequestInformation(options *MessageItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property messages in me
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(options *MessageItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateReply the createReply property
func (m *MessageItemRequestBuilder) CreateReply()(*ib049a164b38d6ba9b4a3b7192971ad5e4b114c289241f24f85456ac7681289f7.CreateReplyRequestBuilder) {
    return ib049a164b38d6ba9b4a3b7192971ad5e4b114c289241f24f85456ac7681289f7.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateReplyAll the createReplyAll property
func (m *MessageItemRequestBuilder) CreateReplyAll()(*i3bfc69f268f92d3eb51d630e84bcf0e900efda493e2ca9cded7df23eadb96f04.CreateReplyAllRequestBuilder) {
    return i3bfc69f268f92d3eb51d630e84bcf0e900efda493e2ca9cded7df23eadb96f04.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for me
func (m *MessageItemRequestBuilder) Delete(options *MessageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EventMessageRequest the eventMessageRequest property
func (m *MessageItemRequestBuilder) EventMessageRequest()(*i36ba3a15f9b3d43bd01853e14f11e541cd30846655c1041d1198a48d07a44b2f.EventMessageRequestRequestBuilder) {
    return i36ba3a15f9b3d43bd01853e14f11e541cd30846655c1041d1198a48d07a44b2f.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *MessageItemRequestBuilder) Extensions()(*i6757d54d5c87cd028537643e1cb39a47e345c91ea192c6abd02e016c98bbc3d3.ExtensionsRequestBuilder) {
    return i6757d54d5c87cd028537643e1cb39a47e345c91ea192c6abd02e016c98bbc3d3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*i9e4259eea1e8905b783d4fc86aede2440ef9a47ea5f2d15815788ac67eba3d69.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9e4259eea1e8905b783d4fc86aede2440ef9a47ea5f2d15815788ac67eba3d69.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *MessageItemRequestBuilder) Forward()(*i272e002e7b409e74d18f258c8a21d69615f41c0248b18a4763a968e49aa9f78c.ForwardRequestBuilder) {
    return i272e002e7b409e74d18f258c8a21d69615f41c0248b18a4763a968e49aa9f78c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) Get(options *MessageItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// Mentions the mentions property
func (m *MessageItemRequestBuilder) Mentions()(*idd4dd625416cd52060f95e8ad108f1f8964fca2bf3f61a6253fbf7068c91d62d.MentionsRequestBuilder) {
    return idd4dd625416cd52060f95e8ad108f1f8964fca2bf3f61a6253fbf7068c91d62d.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.mentions.item collection
func (m *MessageItemRequestBuilder) MentionsById(id string)(*id8651f068bd1a845ecd012a2bfde14061474c28c5020e8285e3b05b6f3fae1fc.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return id8651f068bd1a845ecd012a2bfde14061474c28c5020e8285e3b05b6f3fae1fc.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *MessageItemRequestBuilder) Move()(*i9e5b9792ff40a6736ff612b9493f83c37c80ff3eda2ef70fc12193c0dca8b5fd.MoveRequestBuilder) {
    return i9e5b9792ff40a6736ff612b9493f83c37c80ff3eda2ef70fc12193c0dca8b5fd.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*i537e97523a2d84f6e4e5d67a6432c411a6c3d123670344981076f6bca0acd27f.MultiValueExtendedPropertiesRequestBuilder) {
    return i537e97523a2d84f6e4e5d67a6432c411a6c3d123670344981076f6bca0acd27f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie7ad4a32eeb50c09e552caa2785b20f0fd2063c38f777f192847be1c617ca319.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ie7ad4a32eeb50c09e552caa2785b20f0fd2063c38f777f192847be1c617ca319.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in me
func (m *MessageItemRequestBuilder) Patch(options *MessageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Reply the reply property
func (m *MessageItemRequestBuilder) Reply()(*i17a9af3ebcd424a0f4a7360aa860cfe25fccbd6f75d1f08e88b91b99baffe208.ReplyRequestBuilder) {
    return i17a9af3ebcd424a0f4a7360aa860cfe25fccbd6f75d1f08e88b91b99baffe208.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReplyAll the replyAll property
func (m *MessageItemRequestBuilder) ReplyAll()(*i062177ea9220fd8df21c532fe25b898767dfd98ec54bcfa1a849d01a69301aad.ReplyAllRequestBuilder) {
    return i062177ea9220fd8df21c532fe25b898767dfd98ec54bcfa1a849d01a69301aad.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Send the send property
func (m *MessageItemRequestBuilder) Send()(*i87147cd2592bb7a4f35c0388ca4cd3212b1247c3d08cee9e1e475c3ba0f8ed64.SendRequestBuilder) {
    return i87147cd2592bb7a4f35c0388ca4cd3212b1247c3d08cee9e1e475c3ba0f8ed64.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*i0c991f94222b3272f1430d5abdc8d769cea7a432edd8c195e8ed7efd98dbe501.SingleValueExtendedPropertiesRequestBuilder) {
    return i0c991f94222b3272f1430d5abdc8d769cea7a432edd8c195e8ed7efd98dbe501.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifa2e1a49bbba48dd1867ab7775f7bc0212fe78345f75addbc5353906df09475c.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ifa2e1a49bbba48dd1867ab7775f7bc0212fe78345f75addbc5353906df09475c.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubscribe the unsubscribe property
func (m *MessageItemRequestBuilder) Unsubscribe()(*idcd1232c5c17ea39a25044877c92f642df5dd303671e6efaa847bc930d661f82.UnsubscribeRequestBuilder) {
    return idcd1232c5c17ea39a25044877c92f642df5dd303671e6efaa847bc930d661f82.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
