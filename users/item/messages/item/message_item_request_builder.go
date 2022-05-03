package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i056e6f666cd8140ddec7ddcf1a7900366b0e6af6d9faf45f3a46da610cabfc57 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/eventmessagerequest"
    i2b2060939a0ef1d46f805299a919f2f1acb5c4b8393a3669b6e62c88b0cd21d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/unsubscribe"
    i2c03149e2a2f4f0a33778c3fc372b1c5adf239e90a634545664561317c974ab9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/attachments"
    i2f36980b1005ebc99887b4ad704e1c53024acc466ea8f157d7b2068b15eea5cc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/createforward"
    i4cb3447d6dc3080501e1909902f863f22d9c68e969df7d93ec3d63eab7968f27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/send"
    i4cf366cbeee9f81c540c779d62b211944e2e3503f1361c1a1139b1265c57c869 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/multivalueextendedproperties"
    i931d999adf39b374b26f49ebb1464618c94ed08bd1cdb7225acccee27b1c080d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/reply"
    i9aa296a7568bd3aa8debb9247b58f10846009232cb3afd25c7666ef7204766d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/copy"
    ia0834b4c5190f170531f59f483b526a89731d5a1ffa4962be9d0ca5ca1757caf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/createreply"
    iac36b5af57f64834c4def950ed3d981db5e758b190127f3d98e54d1bdfc0795c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/move"
    ib48ba94d1fb044af23d5ecb6387190df1a8c4c1bc33d60993f88d7b3b7c3f30f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/replyall"
    ibe7f325cbe1ec5e2d66fc8882275565d63ca39eac3637fcdbff9f14dc50cfde5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/mentions"
    ic6de15dbf64e4aab474a1deb971bfb0c896a45a7025134d250ee8434ea00478d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/forward"
    iceab81a5ef304c168de8559bb3112c9e1ec3a8ce4cd374d889e0a91e50a6d957 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/createreplyall"
    id0b86a1c1f3328a66937f969df008a0331f0a3fad808b5b8dd909a32c426eab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/calendarsharingmessage"
    iddc9dbdb21d06b8a1b063bdeb979e0a85ca7160d4d67b055db134c19d315bf2a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/extensions"
    ie94722dbe75a9261a2b8f12426af82f52d5d87a6db61d7a5fbbf3626b6f61e70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/singlevalueextendedproperties"
    if162dca08c4d11c1591fbf1764ccae4d199b1d20b366836ad81039fb45519f61 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/value"
    i0909081a001e274a85adacd7855fad4104566f4a7fa6025579518f53cd144481 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/singlevalueextendedproperties/item"
    i6da3c11b671083e669a9528abeb93c50845a61c10607eaee12efca48168a9571 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/extensions/item"
    i71f4656ff41340d794cffa133f7d9c652e4b62a402e1f0a87c761f16d4c3ac60 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/attachments/item"
    ibb22c71dd88b28423a257b6e8dda3b608e7197c4ac39a2521dabebaa41404686 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/multivalueextendedproperties/item"
    if24f14720e5ede9694b6a001b0205d0b6de138e786e6d8813207ddc75c447e29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/mentions/item"
)

// MessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.user entity.
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MessageItemRequestBuilderGetQueryParameters the messages in a mailbox or folder. Read-only. Nullable.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MessageItemRequestBuilderGetQueryParameters
}
// MessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments the attachments property
func (m *MessageItemRequestBuilder) Attachments()(*i2c03149e2a2f4f0a33778c3fc372b1c5adf239e90a634545664561317c974ab9.AttachmentsRequestBuilder) {
    return i2c03149e2a2f4f0a33778c3fc372b1c5adf239e90a634545664561317c974ab9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*i71f4656ff41340d794cffa133f7d9c652e4b62a402e1f0a87c761f16d4c3ac60.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i71f4656ff41340d794cffa133f7d9c652e4b62a402e1f0a87c761f16d4c3ac60.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarSharingMessage the calendarSharingMessage property
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*id0b86a1c1f3328a66937f969df008a0331f0a3fad808b5b8dd909a32c426eab5.CalendarSharingMessageRequestBuilder) {
    return id0b86a1c1f3328a66937f969df008a0331f0a3fad808b5b8dd909a32c426eab5.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}{?%24select}";
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
func (m *MessageItemRequestBuilder) Content()(*if162dca08c4d11c1591fbf1764ccae4d199b1d20b366836ad81039fb45519f61.ContentRequestBuilder) {
    return if162dca08c4d11c1591fbf1764ccae4d199b1d20b366836ad81039fb45519f61.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *MessageItemRequestBuilder) Copy()(*i9aa296a7568bd3aa8debb9247b58f10846009232cb3afd25c7666ef7204766d4.CopyRequestBuilder) {
    return i9aa296a7568bd3aa8debb9247b58f10846009232cb3afd25c7666ef7204766d4.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateForward the createForward property
func (m *MessageItemRequestBuilder) CreateForward()(*i2f36980b1005ebc99887b4ad704e1c53024acc466ea8f157d7b2068b15eea5cc.CreateForwardRequestBuilder) {
    return i2f36980b1005ebc99887b4ad704e1c53024acc466ea8f157d7b2068b15eea5cc.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property messages in users
func (m *MessageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateReply the createReply property
func (m *MessageItemRequestBuilder) CreateReply()(*ia0834b4c5190f170531f59f483b526a89731d5a1ffa4962be9d0ca5ca1757caf.CreateReplyRequestBuilder) {
    return ia0834b4c5190f170531f59f483b526a89731d5a1ffa4962be9d0ca5ca1757caf.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateReplyAll the createReplyAll property
func (m *MessageItemRequestBuilder) CreateReplyAll()(*iceab81a5ef304c168de8559bb3112c9e1ec3a8ce4cd374d889e0a91e50a6d957.CreateReplyAllRequestBuilder) {
    return iceab81a5ef304c168de8559bb3112c9e1ec3a8ce4cd374d889e0a91e50a6d957.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for users
func (m *MessageItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property messages for users
func (m *MessageItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EventMessageRequest the eventMessageRequest property
func (m *MessageItemRequestBuilder) EventMessageRequest()(*i056e6f666cd8140ddec7ddcf1a7900366b0e6af6d9faf45f3a46da610cabfc57.EventMessageRequestRequestBuilder) {
    return i056e6f666cd8140ddec7ddcf1a7900366b0e6af6d9faf45f3a46da610cabfc57.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *MessageItemRequestBuilder) Extensions()(*iddc9dbdb21d06b8a1b063bdeb979e0a85ca7160d4d67b055db134c19d315bf2a.ExtensionsRequestBuilder) {
    return iddc9dbdb21d06b8a1b063bdeb979e0a85ca7160d4d67b055db134c19d315bf2a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*i6da3c11b671083e669a9528abeb93c50845a61c10607eaee12efca48168a9571.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i6da3c11b671083e669a9528abeb93c50845a61c10607eaee12efca48168a9571.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *MessageItemRequestBuilder) Forward()(*ic6de15dbf64e4aab474a1deb971bfb0c896a45a7025134d250ee8434ea00478d.ForwardRequestBuilder) {
    return ic6de15dbf64e4aab474a1deb971bfb0c896a45a7025134d250ee8434ea00478d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// Mentions the mentions property
func (m *MessageItemRequestBuilder) Mentions()(*ibe7f325cbe1ec5e2d66fc8882275565d63ca39eac3637fcdbff9f14dc50cfde5.MentionsRequestBuilder) {
    return ibe7f325cbe1ec5e2d66fc8882275565d63ca39eac3637fcdbff9f14dc50cfde5.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.mentions.item collection
func (m *MessageItemRequestBuilder) MentionsById(id string)(*if24f14720e5ede9694b6a001b0205d0b6de138e786e6d8813207ddc75c447e29.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return if24f14720e5ede9694b6a001b0205d0b6de138e786e6d8813207ddc75c447e29.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *MessageItemRequestBuilder) Move()(*iac36b5af57f64834c4def950ed3d981db5e758b190127f3d98e54d1bdfc0795c.MoveRequestBuilder) {
    return iac36b5af57f64834c4def950ed3d981db5e758b190127f3d98e54d1bdfc0795c.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*i4cf366cbeee9f81c540c779d62b211944e2e3503f1361c1a1139b1265c57c869.MultiValueExtendedPropertiesRequestBuilder) {
    return i4cf366cbeee9f81c540c779d62b211944e2e3503f1361c1a1139b1265c57c869.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibb22c71dd88b28423a257b6e8dda3b608e7197c4ac39a2521dabebaa41404686.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ibb22c71dd88b28423a257b6e8dda3b608e7197c4ac39a2521dabebaa41404686.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in users
func (m *MessageItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property messages in users
func (m *MessageItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Reply the reply property
func (m *MessageItemRequestBuilder) Reply()(*i931d999adf39b374b26f49ebb1464618c94ed08bd1cdb7225acccee27b1c080d.ReplyRequestBuilder) {
    return i931d999adf39b374b26f49ebb1464618c94ed08bd1cdb7225acccee27b1c080d.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReplyAll the replyAll property
func (m *MessageItemRequestBuilder) ReplyAll()(*ib48ba94d1fb044af23d5ecb6387190df1a8c4c1bc33d60993f88d7b3b7c3f30f.ReplyAllRequestBuilder) {
    return ib48ba94d1fb044af23d5ecb6387190df1a8c4c1bc33d60993f88d7b3b7c3f30f.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Send the send property
func (m *MessageItemRequestBuilder) Send()(*i4cb3447d6dc3080501e1909902f863f22d9c68e969df7d93ec3d63eab7968f27.SendRequestBuilder) {
    return i4cb3447d6dc3080501e1909902f863f22d9c68e969df7d93ec3d63eab7968f27.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*ie94722dbe75a9261a2b8f12426af82f52d5d87a6db61d7a5fbbf3626b6f61e70.SingleValueExtendedPropertiesRequestBuilder) {
    return ie94722dbe75a9261a2b8f12426af82f52d5d87a6db61d7a5fbbf3626b6f61e70.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0909081a001e274a85adacd7855fad4104566f4a7fa6025579518f53cd144481.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i0909081a001e274a85adacd7855fad4104566f4a7fa6025579518f53cd144481.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubscribe the unsubscribe property
func (m *MessageItemRequestBuilder) Unsubscribe()(*i2b2060939a0ef1d46f805299a919f2f1acb5c4b8393a3669b6e62c88b0cd21d0.UnsubscribeRequestBuilder) {
    return i2b2060939a0ef1d46f805299a919f2f1acb5c4b8393a3669b6e62c88b0cd21d0.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
