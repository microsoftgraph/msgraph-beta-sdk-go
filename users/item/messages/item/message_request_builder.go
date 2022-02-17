package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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

// MessageRequestBuilder builds and executes requests for operations under \users\{user-id}\messages\{message-id}
type MessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MessageRequestBuilderDeleteOptions options for Delete
type MessageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageRequestBuilderGetOptions options for Get
type MessageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MessageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageRequestBuilderGetQueryParameters the messages in a mailbox or folder. Read-only. Nullable.
type MessageRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// MessageRequestBuilderPatchOptions options for Patch
type MessageRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Message;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MessageRequestBuilder) Attachments()(*i2c03149e2a2f4f0a33778c3fc372b1c5adf239e90a634545664561317c974ab9.AttachmentsRequestBuilder) {
    return i2c03149e2a2f4f0a33778c3fc372b1c5adf239e90a634545664561317c974ab9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.attachments.item collection
func (m *MessageRequestBuilder) AttachmentsById(id string)(*i71f4656ff41340d794cffa133f7d9c652e4b62a402e1f0a87c761f16d4c3ac60.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i71f4656ff41340d794cffa133f7d9c652e4b62a402e1f0a87c761f16d4c3ac60.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) CalendarSharingMessage()(*id0b86a1c1f3328a66937f969df008a0331f0a3fad808b5b8dd909a32c426eab5.CalendarSharingMessageRequestBuilder) {
    return id0b86a1c1f3328a66937f969df008a0331f0a3fad808b5b8dd909a32c426eab5.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageRequestBuilderInternal instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/messages/{message_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageRequestBuilder instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessageRequestBuilder) Content()(*if162dca08c4d11c1591fbf1764ccae4d199b1d20b366836ad81039fb45519f61.ContentRequestBuilder) {
    return if162dca08c4d11c1591fbf1764ccae4d199b1d20b366836ad81039fb45519f61.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Copy()(*i9aa296a7568bd3aa8debb9247b58f10846009232cb3afd25c7666ef7204766d4.CopyRequestBuilder) {
    return i9aa296a7568bd3aa8debb9247b58f10846009232cb3afd25c7666ef7204766d4.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) CreateDeleteRequestInformation(options *MessageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MessageRequestBuilder) CreateForward()(*i2f36980b1005ebc99887b4ad704e1c53024acc466ea8f157d7b2068b15eea5cc.CreateForwardRequestBuilder) {
    return i2f36980b1005ebc99887b4ad704e1c53024acc466ea8f157d7b2068b15eea5cc.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) CreateGetRequestInformation(options *MessageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) CreatePatchRequestInformation(options *MessageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MessageRequestBuilder) CreateReply()(*ia0834b4c5190f170531f59f483b526a89731d5a1ffa4962be9d0ca5ca1757caf.CreateReplyRequestBuilder) {
    return ia0834b4c5190f170531f59f483b526a89731d5a1ffa4962be9d0ca5ca1757caf.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) CreateReplyAll()(*iceab81a5ef304c168de8559bb3112c9e1ec3a8ce4cd374d889e0a91e50a6d957.CreateReplyAllRequestBuilder) {
    return iceab81a5ef304c168de8559bb3112c9e1ec3a8ce4cd374d889e0a91e50a6d957.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) Delete(options *MessageRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageRequestBuilder) EventMessageRequest()(*i056e6f666cd8140ddec7ddcf1a7900366b0e6af6d9faf45f3a46da610cabfc57.EventMessageRequestRequestBuilder) {
    return i056e6f666cd8140ddec7ddcf1a7900366b0e6af6d9faf45f3a46da610cabfc57.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Extensions()(*iddc9dbdb21d06b8a1b063bdeb979e0a85ca7160d4d67b055db134c19d315bf2a.ExtensionsRequestBuilder) {
    return iddc9dbdb21d06b8a1b063bdeb979e0a85ca7160d4d67b055db134c19d315bf2a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.extensions.item collection
func (m *MessageRequestBuilder) ExtensionsById(id string)(*i6da3c11b671083e669a9528abeb93c50845a61c10607eaee12efca48168a9571.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i6da3c11b671083e669a9528abeb93c50845a61c10607eaee12efca48168a9571.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Forward()(*ic6de15dbf64e4aab474a1deb971bfb0c896a45a7025134d250ee8434ea00478d.ForwardRequestBuilder) {
    return ic6de15dbf64e4aab474a1deb971bfb0c896a45a7025134d250ee8434ea00478d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) Get(options *MessageRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Message, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMessage() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Message), nil
}
func (m *MessageRequestBuilder) Mentions()(*ibe7f325cbe1ec5e2d66fc8882275565d63ca39eac3637fcdbff9f14dc50cfde5.MentionsRequestBuilder) {
    return ibe7f325cbe1ec5e2d66fc8882275565d63ca39eac3637fcdbff9f14dc50cfde5.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.mentions.item collection
func (m *MessageRequestBuilder) MentionsById(id string)(*if24f14720e5ede9694b6a001b0205d0b6de138e786e6d8813207ddc75c447e29.MentionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return if24f14720e5ede9694b6a001b0205d0b6de138e786e6d8813207ddc75c447e29.NewMentionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Move()(*iac36b5af57f64834c4def950ed3d981db5e758b190127f3d98e54d1bdfc0795c.MoveRequestBuilder) {
    return iac36b5af57f64834c4def950ed3d981db5e758b190127f3d98e54d1bdfc0795c.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*i4cf366cbeee9f81c540c779d62b211944e2e3503f1361c1a1139b1265c57c869.MultiValueExtendedPropertiesRequestBuilder) {
    return i4cf366cbeee9f81c540c779d62b211944e2e3503f1361c1a1139b1265c57c869.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibb22c71dd88b28423a257b6e8dda3b608e7197c4ac39a2521dabebaa41404686.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ibb22c71dd88b28423a257b6e8dda3b608e7197c4ac39a2521dabebaa41404686.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) Patch(options *MessageRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageRequestBuilder) Reply()(*i931d999adf39b374b26f49ebb1464618c94ed08bd1cdb7225acccee27b1c080d.ReplyRequestBuilder) {
    return i931d999adf39b374b26f49ebb1464618c94ed08bd1cdb7225acccee27b1c080d.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) ReplyAll()(*ib48ba94d1fb044af23d5ecb6387190df1a8c4c1bc33d60993f88d7b3b7c3f30f.ReplyAllRequestBuilder) {
    return ib48ba94d1fb044af23d5ecb6387190df1a8c4c1bc33d60993f88d7b3b7c3f30f.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Send()(*i4cb3447d6dc3080501e1909902f863f22d9c68e969df7d93ec3d63eab7968f27.SendRequestBuilder) {
    return i4cb3447d6dc3080501e1909902f863f22d9c68e969df7d93ec3d63eab7968f27.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*ie94722dbe75a9261a2b8f12426af82f52d5d87a6db61d7a5fbbf3626b6f61e70.SingleValueExtendedPropertiesRequestBuilder) {
    return ie94722dbe75a9261a2b8f12426af82f52d5d87a6db61d7a5fbbf3626b6f61e70.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0909081a001e274a85adacd7855fad4104566f4a7fa6025579518f53cd144481.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i0909081a001e274a85adacd7855fad4104566f4a7fa6025579518f53cd144481.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Unsubscribe()(*i2b2060939a0ef1d46f805299a919f2f1acb5c4b8393a3669b6e62c88b0cd21d0.UnsubscribeRequestBuilder) {
    return i2b2060939a0ef1d46f805299a919f2f1acb5c4b8393a3669b6e62c88b0cd21d0.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
