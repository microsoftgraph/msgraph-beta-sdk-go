package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i023106483eeaa9e057f5ee0ef9eef33b274da7e759afeda3ae1fc96fdf554f97 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/createforward"
    i25c967d20ec85be402e1a7c6d121a8aa4bd088c230867e82dd4528d524fe2b62 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/copy"
    i25d6ea23f29436bc3a139a769323a75062479cfc6f4911effd52f88f00991149 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/mentions"
    i3ffb0826c99fc33b6dcbba6441448cbc6a4ecc20d5ef53f1cf27d694abb3f34f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/move"
    i5a5621f64056659fd70ac642a8529ade985d1f0cc1dd08f36ddba20deeda0930 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/calendarsharingmessage"
    i9a5d009080ecc4a3753fb5c03ffb0cba6bf87135dffa76209c8e925c7465820b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/unsubscribe"
    ia02e177b92e83647b153fece558ac26faee0de4a3cf44fb885497858dc7ab92a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/extensions"
    ia0f4e54b2ec34c3c2ef48b17a17fd53c26182a272b8e7a1b90ae44c73540b466 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/createreply"
    ia5ef716c57bbcc97067064429a1b0cabc40577038c19d17ec373d9d9d3acb49b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/singlevalueextendedproperties"
    iac432dfcc94115868ef80393dd9192f38b599425feb55dd1dfe9a9cd4d0abb7d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/forward"
    iae74de854139dabd76aa9dccb352cc3ca74352ebd7b498cb514e46847b2a6d8c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/multivalueextendedproperties"
    iafe3c53978f020c73991af93d711edee6038e959ba51d882849f967994ca66be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/eventmessagerequest"
    ic7f35c75b6e62e54c20816e6b842b525069091b3705f0eda5abb2f06698b2e46 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/attachments"
    id409f9b5c6726c0f1e4f275d9c3ccf1d7a84bf2ecd174e775a18be8fcb6c6a43 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/send"
    idba2d70212ff30cbe1857bf8f733677622af732ef7e26a5c94efed97901ff29f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/createreplyall"
    iea011617347218b7a8bd89bf3c13bd05549b1a8b3b69952d3dbf3689a69d17cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/replyall"
    iedb429bab606ecbf845de51a90f7df524e6a484ae176125e9b28ce85a07ae391 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/reply"
    if6a5d561727ceed03138849e77938ba73f694713a884563fbabefa7f5696d88d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/value"
    i340f83e62384fa8614f95e359ba5bc4e913d7996b1fea9840ca4fbb8ab177ec8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/mentions/item"
    i492cbc6c0049fb5261ffb6826b11ec61ae24e4935436c10ef66e6d3adf6d07ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/extensions/item"
    i5114ea7bc7e0bddf25a0022c76f1aa4ea041eb41634ed63956fb24a1c7417d21 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/multivalueextendedproperties/item"
    i71a573e50c8dcf20dad2d6f95ff5ba431744934140f7d3c80ab47bf4e765775a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/singlevalueextendedproperties/item"
    i9dfb9baf1034d4d76f7676658c28a48bb59e8a78be7fe78076e5c63c401df6ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/attachments/item"
)

// MessageRequestBuilder builds and executes requests for operations under \me\messages\{message-id}
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
func (m *MessageRequestBuilder) Attachments()(*ic7f35c75b6e62e54c20816e6b842b525069091b3705f0eda5abb2f06698b2e46.AttachmentsRequestBuilder) {
    return ic7f35c75b6e62e54c20816e6b842b525069091b3705f0eda5abb2f06698b2e46.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.attachments.item collection
func (m *MessageRequestBuilder) AttachmentsById(id string)(*i9dfb9baf1034d4d76f7676658c28a48bb59e8a78be7fe78076e5c63c401df6ee.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i9dfb9baf1034d4d76f7676658c28a48bb59e8a78be7fe78076e5c63c401df6ee.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) CalendarSharingMessage()(*i5a5621f64056659fd70ac642a8529ade985d1f0cc1dd08f36ddba20deeda0930.CalendarSharingMessageRequestBuilder) {
    return i5a5621f64056659fd70ac642a8529ade985d1f0cc1dd08f36ddba20deeda0930.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageRequestBuilderInternal instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/messages/{message_id}{?select}";
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
func (m *MessageRequestBuilder) Content()(*if6a5d561727ceed03138849e77938ba73f694713a884563fbabefa7f5696d88d.ContentRequestBuilder) {
    return if6a5d561727ceed03138849e77938ba73f694713a884563fbabefa7f5696d88d.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Copy()(*i25c967d20ec85be402e1a7c6d121a8aa4bd088c230867e82dd4528d524fe2b62.CopyRequestBuilder) {
    return i25c967d20ec85be402e1a7c6d121a8aa4bd088c230867e82dd4528d524fe2b62.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageRequestBuilder) CreateForward()(*i023106483eeaa9e057f5ee0ef9eef33b274da7e759afeda3ae1fc96fdf554f97.CreateForwardRequestBuilder) {
    return i023106483eeaa9e057f5ee0ef9eef33b274da7e759afeda3ae1fc96fdf554f97.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageRequestBuilder) CreateReply()(*ia0f4e54b2ec34c3c2ef48b17a17fd53c26182a272b8e7a1b90ae44c73540b466.CreateReplyRequestBuilder) {
    return ia0f4e54b2ec34c3c2ef48b17a17fd53c26182a272b8e7a1b90ae44c73540b466.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) CreateReplyAll()(*idba2d70212ff30cbe1857bf8f733677622af732ef7e26a5c94efed97901ff29f.CreateReplyAllRequestBuilder) {
    return idba2d70212ff30cbe1857bf8f733677622af732ef7e26a5c94efed97901ff29f.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageRequestBuilder) EventMessageRequest()(*iafe3c53978f020c73991af93d711edee6038e959ba51d882849f967994ca66be.EventMessageRequestRequestBuilder) {
    return iafe3c53978f020c73991af93d711edee6038e959ba51d882849f967994ca66be.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Extensions()(*ia02e177b92e83647b153fece558ac26faee0de4a3cf44fb885497858dc7ab92a.ExtensionsRequestBuilder) {
    return ia02e177b92e83647b153fece558ac26faee0de4a3cf44fb885497858dc7ab92a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.extensions.item collection
func (m *MessageRequestBuilder) ExtensionsById(id string)(*i492cbc6c0049fb5261ffb6826b11ec61ae24e4935436c10ef66e6d3adf6d07ed.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i492cbc6c0049fb5261ffb6826b11ec61ae24e4935436c10ef66e6d3adf6d07ed.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Forward()(*iac432dfcc94115868ef80393dd9192f38b599425feb55dd1dfe9a9cd4d0abb7d.ForwardRequestBuilder) {
    return iac432dfcc94115868ef80393dd9192f38b599425feb55dd1dfe9a9cd4d0abb7d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageRequestBuilder) Mentions()(*i25d6ea23f29436bc3a139a769323a75062479cfc6f4911effd52f88f00991149.MentionsRequestBuilder) {
    return i25d6ea23f29436bc3a139a769323a75062479cfc6f4911effd52f88f00991149.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.mentions.item collection
func (m *MessageRequestBuilder) MentionsById(id string)(*i340f83e62384fa8614f95e359ba5bc4e913d7996b1fea9840ca4fbb8ab177ec8.MentionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return i340f83e62384fa8614f95e359ba5bc4e913d7996b1fea9840ca4fbb8ab177ec8.NewMentionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Move()(*i3ffb0826c99fc33b6dcbba6441448cbc6a4ecc20d5ef53f1cf27d694abb3f34f.MoveRequestBuilder) {
    return i3ffb0826c99fc33b6dcbba6441448cbc6a4ecc20d5ef53f1cf27d694abb3f34f.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*iae74de854139dabd76aa9dccb352cc3ca74352ebd7b498cb514e46847b2a6d8c.MultiValueExtendedPropertiesRequestBuilder) {
    return iae74de854139dabd76aa9dccb352cc3ca74352ebd7b498cb514e46847b2a6d8c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.multiValueExtendedProperties.item collection
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5114ea7bc7e0bddf25a0022c76f1aa4ea041eb41634ed63956fb24a1c7417d21.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5114ea7bc7e0bddf25a0022c76f1aa4ea041eb41634ed63956fb24a1c7417d21.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *MessageRequestBuilder) Reply()(*iedb429bab606ecbf845de51a90f7df524e6a484ae176125e9b28ce85a07ae391.ReplyRequestBuilder) {
    return iedb429bab606ecbf845de51a90f7df524e6a484ae176125e9b28ce85a07ae391.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) ReplyAll()(*iea011617347218b7a8bd89bf3c13bd05549b1a8b3b69952d3dbf3689a69d17cb.ReplyAllRequestBuilder) {
    return iea011617347218b7a8bd89bf3c13bd05549b1a8b3b69952d3dbf3689a69d17cb.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Send()(*id409f9b5c6726c0f1e4f275d9c3ccf1d7a84bf2ecd174e775a18be8fcb6c6a43.SendRequestBuilder) {
    return id409f9b5c6726c0f1e4f275d9c3ccf1d7a84bf2ecd174e775a18be8fcb6c6a43.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*ia5ef716c57bbcc97067064429a1b0cabc40577038c19d17ec373d9d9d3acb49b.SingleValueExtendedPropertiesRequestBuilder) {
    return ia5ef716c57bbcc97067064429a1b0cabc40577038c19d17ec373d9d9d3acb49b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.singleValueExtendedProperties.item collection
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i71a573e50c8dcf20dad2d6f95ff5ba431744934140f7d3c80ab47bf4e765775a.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i71a573e50c8dcf20dad2d6f95ff5ba431744934140f7d3c80ab47bf4e765775a.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Unsubscribe()(*i9a5d009080ecc4a3753fb5c03ffb0cba6bf87135dffa76209c8e925c7465820b.UnsubscribeRequestBuilder) {
    return i9a5d009080ecc4a3753fb5c03ffb0cba6bf87135dffa76209c8e925c7465820b.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
