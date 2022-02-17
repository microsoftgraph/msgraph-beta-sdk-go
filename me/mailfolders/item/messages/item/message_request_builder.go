package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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

// MessageRequestBuilder builds and executes requests for operations under \me\mailFolders\{mailFolder-id}\messages\{message-id}
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
// MessageRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
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
func (m *MessageRequestBuilder) Attachments()(*ibc3856e7bed2bea8783ad4cf102bbf5295cc9b7cc63847236c2fcb7a7d37d1b9.AttachmentsRequestBuilder) {
    return ibc3856e7bed2bea8783ad4cf102bbf5295cc9b7cc63847236c2fcb7a7d37d1b9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.attachments.item collection
func (m *MessageRequestBuilder) AttachmentsById(id string)(*i671b494222abfe980bb8d667f8d3d2047da5dddc998c20eab8bb67f1d8672de6.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i671b494222abfe980bb8d667f8d3d2047da5dddc998c20eab8bb67f1d8672de6.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) CalendarSharingMessage()(*i0fccf38aca2f46144aa036d0c0fa6ee24e01b5a7a3993c6f341c44efcb5232f7.CalendarSharingMessageRequestBuilder) {
    return i0fccf38aca2f46144aa036d0c0fa6ee24e01b5a7a3993c6f341c44efcb5232f7.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageRequestBuilderInternal instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder_id}/messages/{message_id}{?select,expand}";
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
func (m *MessageRequestBuilder) Content()(*ibd12efd1c58ab297076323c8b68a12f9e5cfa0fe08af58e0c140c589be1fa326.ContentRequestBuilder) {
    return ibd12efd1c58ab297076323c8b68a12f9e5cfa0fe08af58e0c140c589be1fa326.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Copy()(*i72b0f6daa0bcef81a7a85e98d43a239d41374d9126783782e12898c44e092655.CopyRequestBuilder) {
    return i72b0f6daa0bcef81a7a85e98d43a239d41374d9126783782e12898c44e092655.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) CreateForward()(*i05e8b869ffc519fee49d0ac259a1f1eea9becc6ebc1c5a37c2f85196b6ee8ee3.CreateForwardRequestBuilder) {
    return i05e8b869ffc519fee49d0ac259a1f1eea9becc6ebc1c5a37c2f85196b6ee8ee3.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
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
// CreatePatchRequestInformation the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) CreateReply()(*ib049a164b38d6ba9b4a3b7192971ad5e4b114c289241f24f85456ac7681289f7.CreateReplyRequestBuilder) {
    return ib049a164b38d6ba9b4a3b7192971ad5e4b114c289241f24f85456ac7681289f7.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) CreateReplyAll()(*i3bfc69f268f92d3eb51d630e84bcf0e900efda493e2ca9cded7df23eadb96f04.CreateReplyAllRequestBuilder) {
    return i3bfc69f268f92d3eb51d630e84bcf0e900efda493e2ca9cded7df23eadb96f04.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) EventMessageRequest()(*i36ba3a15f9b3d43bd01853e14f11e541cd30846655c1041d1198a48d07a44b2f.EventMessageRequestRequestBuilder) {
    return i36ba3a15f9b3d43bd01853e14f11e541cd30846655c1041d1198a48d07a44b2f.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Extensions()(*i6757d54d5c87cd028537643e1cb39a47e345c91ea192c6abd02e016c98bbc3d3.ExtensionsRequestBuilder) {
    return i6757d54d5c87cd028537643e1cb39a47e345c91ea192c6abd02e016c98bbc3d3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.extensions.item collection
func (m *MessageRequestBuilder) ExtensionsById(id string)(*i9e4259eea1e8905b783d4fc86aede2440ef9a47ea5f2d15815788ac67eba3d69.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i9e4259eea1e8905b783d4fc86aede2440ef9a47ea5f2d15815788ac67eba3d69.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Forward()(*i272e002e7b409e74d18f258c8a21d69615f41c0248b18a4763a968e49aa9f78c.ForwardRequestBuilder) {
    return i272e002e7b409e74d18f258c8a21d69615f41c0248b18a4763a968e49aa9f78c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) Mentions()(*idd4dd625416cd52060f95e8ad108f1f8964fca2bf3f61a6253fbf7068c91d62d.MentionsRequestBuilder) {
    return idd4dd625416cd52060f95e8ad108f1f8964fca2bf3f61a6253fbf7068c91d62d.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.mentions.item collection
func (m *MessageRequestBuilder) MentionsById(id string)(*id8651f068bd1a845ecd012a2bfde14061474c28c5020e8285e3b05b6f3fae1fc.MentionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return id8651f068bd1a845ecd012a2bfde14061474c28c5020e8285e3b05b6f3fae1fc.NewMentionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Move()(*i9e5b9792ff40a6736ff612b9493f83c37c80ff3eda2ef70fc12193c0dca8b5fd.MoveRequestBuilder) {
    return i9e5b9792ff40a6736ff612b9493f83c37c80ff3eda2ef70fc12193c0dca8b5fd.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*i537e97523a2d84f6e4e5d67a6432c411a6c3d123670344981076f6bca0acd27f.MultiValueExtendedPropertiesRequestBuilder) {
    return i537e97523a2d84f6e4e5d67a6432c411a6c3d123670344981076f6bca0acd27f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie7ad4a32eeb50c09e552caa2785b20f0fd2063c38f777f192847be1c617ca319.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie7ad4a32eeb50c09e552caa2785b20f0fd2063c38f777f192847be1c617ca319.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) Reply()(*i17a9af3ebcd424a0f4a7360aa860cfe25fccbd6f75d1f08e88b91b99baffe208.ReplyRequestBuilder) {
    return i17a9af3ebcd424a0f4a7360aa860cfe25fccbd6f75d1f08e88b91b99baffe208.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) ReplyAll()(*i062177ea9220fd8df21c532fe25b898767dfd98ec54bcfa1a849d01a69301aad.ReplyAllRequestBuilder) {
    return i062177ea9220fd8df21c532fe25b898767dfd98ec54bcfa1a849d01a69301aad.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Send()(*i87147cd2592bb7a4f35c0388ca4cd3212b1247c3d08cee9e1e475c3ba0f8ed64.SendRequestBuilder) {
    return i87147cd2592bb7a4f35c0388ca4cd3212b1247c3d08cee9e1e475c3ba0f8ed64.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*i0c991f94222b3272f1430d5abdc8d769cea7a432edd8c195e8ed7efd98dbe501.SingleValueExtendedPropertiesRequestBuilder) {
    return i0c991f94222b3272f1430d5abdc8d769cea7a432edd8c195e8ed7efd98dbe501.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifa2e1a49bbba48dd1867ab7775f7bc0212fe78345f75addbc5353906df09475c.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ifa2e1a49bbba48dd1867ab7775f7bc0212fe78345f75addbc5353906df09475c.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Unsubscribe()(*idcd1232c5c17ea39a25044877c92f642df5dd303671e6efaa847bc930d661f82.UnsubscribeRequestBuilder) {
    return idcd1232c5c17ea39a25044877c92f642df5dd303671e6efaa847bc930d661f82.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
