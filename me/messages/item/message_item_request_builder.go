package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i023106483eeaa9e057f5ee0ef9eef33b274da7e759afeda3ae1fc96fdf554f97 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/createforward"
    i25c967d20ec85be402e1a7c6d121a8aa4bd088c230867e82dd4528d524fe2b62 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/copy"
    i25d6ea23f29436bc3a139a769323a75062479cfc6f4911effd52f88f00991149 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/mentions"
    i3ffb0826c99fc33b6dcbba6441448cbc6a4ecc20d5ef53f1cf27d694abb3f34f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/move"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
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

// MessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.user entity.
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MessageItemRequestBuilderDeleteOptions options for Delete
type MessageItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageItemRequestBuilderGetOptions options for Get
type MessageItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MessageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageItemRequestBuilderGetQueryParameters the messages in a mailbox or folder. Read-only. Nullable.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// MessageItemRequestBuilderPatchOptions options for Patch
type MessageItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Messageable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MessageItemRequestBuilder) Attachments()(*ic7f35c75b6e62e54c20816e6b842b525069091b3705f0eda5abb2f06698b2e46.AttachmentsRequestBuilder) {
    return ic7f35c75b6e62e54c20816e6b842b525069091b3705f0eda5abb2f06698b2e46.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*i9dfb9baf1034d4d76f7676658c28a48bb59e8a78be7fe78076e5c63c401df6ee.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i9dfb9baf1034d4d76f7676658c28a48bb59e8a78be7fe78076e5c63c401df6ee.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*i5a5621f64056659fd70ac642a8529ade985d1f0cc1dd08f36ddba20deeda0930.CalendarSharingMessageRequestBuilder) {
    return i5a5621f64056659fd70ac642a8529ade985d1f0cc1dd08f36ddba20deeda0930.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/messages/{message_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessageItemRequestBuilder) Content()(*if6a5d561727ceed03138849e77938ba73f694713a884563fbabefa7f5696d88d.ContentRequestBuilder) {
    return if6a5d561727ceed03138849e77938ba73f694713a884563fbabefa7f5696d88d.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Copy()(*i25c967d20ec85be402e1a7c6d121a8aa4bd088c230867e82dd4528d524fe2b62.CopyRequestBuilder) {
    return i25c967d20ec85be402e1a7c6d121a8aa4bd088c230867e82dd4528d524fe2b62.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for me
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation(options *MessageItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MessageItemRequestBuilder) CreateForward()(*i023106483eeaa9e057f5ee0ef9eef33b274da7e759afeda3ae1fc96fdf554f97.CreateForwardRequestBuilder) {
    return i023106483eeaa9e057f5ee0ef9eef33b274da7e759afeda3ae1fc96fdf554f97.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) CreateGetRequestInformation(options *MessageItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in me
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(options *MessageItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MessageItemRequestBuilder) CreateReply()(*ia0f4e54b2ec34c3c2ef48b17a17fd53c26182a272b8e7a1b90ae44c73540b466.CreateReplyRequestBuilder) {
    return ia0f4e54b2ec34c3c2ef48b17a17fd53c26182a272b8e7a1b90ae44c73540b466.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CreateReplyAll()(*idba2d70212ff30cbe1857bf8f733677622af732ef7e26a5c94efed97901ff29f.CreateReplyAllRequestBuilder) {
    return idba2d70212ff30cbe1857bf8f733677622af732ef7e26a5c94efed97901ff29f.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for me
func (m *MessageItemRequestBuilder) Delete(options *MessageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageItemRequestBuilder) EventMessageRequest()(*iafe3c53978f020c73991af93d711edee6038e959ba51d882849f967994ca66be.EventMessageRequestRequestBuilder) {
    return iafe3c53978f020c73991af93d711edee6038e959ba51d882849f967994ca66be.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Extensions()(*ia02e177b92e83647b153fece558ac26faee0de4a3cf44fb885497858dc7ab92a.ExtensionsRequestBuilder) {
    return ia02e177b92e83647b153fece558ac26faee0de4a3cf44fb885497858dc7ab92a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*i492cbc6c0049fb5261ffb6826b11ec61ae24e4935436c10ef66e6d3adf6d07ed.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i492cbc6c0049fb5261ffb6826b11ec61ae24e4935436c10ef66e6d3adf6d07ed.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Forward()(*iac432dfcc94115868ef80393dd9192f38b599425feb55dd1dfe9a9cd4d0abb7d.ForwardRequestBuilder) {
    return iac432dfcc94115868ef80393dd9192f38b599425feb55dd1dfe9a9cd4d0abb7d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) Get(options *MessageItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMessageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Messageable), nil
}
func (m *MessageItemRequestBuilder) Mentions()(*i25d6ea23f29436bc3a139a769323a75062479cfc6f4911effd52f88f00991149.MentionsRequestBuilder) {
    return i25d6ea23f29436bc3a139a769323a75062479cfc6f4911effd52f88f00991149.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.mentions.item collection
func (m *MessageItemRequestBuilder) MentionsById(id string)(*i340f83e62384fa8614f95e359ba5bc4e913d7996b1fea9840ca4fbb8ab177ec8.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return i340f83e62384fa8614f95e359ba5bc4e913d7996b1fea9840ca4fbb8ab177ec8.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Move()(*i3ffb0826c99fc33b6dcbba6441448cbc6a4ecc20d5ef53f1cf27d694abb3f34f.MoveRequestBuilder) {
    return i3ffb0826c99fc33b6dcbba6441448cbc6a4ecc20d5ef53f1cf27d694abb3f34f.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*iae74de854139dabd76aa9dccb352cc3ca74352ebd7b498cb514e46847b2a6d8c.MultiValueExtendedPropertiesRequestBuilder) {
    return iae74de854139dabd76aa9dccb352cc3ca74352ebd7b498cb514e46847b2a6d8c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5114ea7bc7e0bddf25a0022c76f1aa4ea041eb41634ed63956fb24a1c7417d21.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5114ea7bc7e0bddf25a0022c76f1aa4ea041eb41634ed63956fb24a1c7417d21.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in me
func (m *MessageItemRequestBuilder) Patch(options *MessageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageItemRequestBuilder) Reply()(*iedb429bab606ecbf845de51a90f7df524e6a484ae176125e9b28ce85a07ae391.ReplyRequestBuilder) {
    return iedb429bab606ecbf845de51a90f7df524e6a484ae176125e9b28ce85a07ae391.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) ReplyAll()(*iea011617347218b7a8bd89bf3c13bd05549b1a8b3b69952d3dbf3689a69d17cb.ReplyAllRequestBuilder) {
    return iea011617347218b7a8bd89bf3c13bd05549b1a8b3b69952d3dbf3689a69d17cb.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Send()(*id409f9b5c6726c0f1e4f275d9c3ccf1d7a84bf2ecd174e775a18be8fcb6c6a43.SendRequestBuilder) {
    return id409f9b5c6726c0f1e4f275d9c3ccf1d7a84bf2ecd174e775a18be8fcb6c6a43.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*ia5ef716c57bbcc97067064429a1b0cabc40577038c19d17ec373d9d9d3acb49b.SingleValueExtendedPropertiesRequestBuilder) {
    return ia5ef716c57bbcc97067064429a1b0cabc40577038c19d17ec373d9d9d3acb49b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i71a573e50c8dcf20dad2d6f95ff5ba431744934140f7d3c80ab47bf4e765775a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i71a573e50c8dcf20dad2d6f95ff5ba431744934140f7d3c80ab47bf4e765775a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Unsubscribe()(*i9a5d009080ecc4a3753fb5c03ffb0cba6bf87135dffa76209c8e925c7465820b.UnsubscribeRequestBuilder) {
    return i9a5d009080ecc4a3753fb5c03ffb0cba6bf87135dffa76209c8e925c7465820b.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
