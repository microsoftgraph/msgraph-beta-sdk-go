package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i18f0f3d80ebc26c14d0b298a950e35a82ff4632a691418331bc5c8ae9f7e1749 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/attachments"
    i25ba706e121e0ece5ebf98516c94795df9564057f989070249206d4cfb86fd93 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/mentions"
    i38bd3b59f483648cf474b3f78ca224a4030f1ccdc95d8b3eca57afa7733c4cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/createreplyall"
    i56e6bf2f80d7479fcf2b474c3916c0c636102e2f6f6405a2dd2c51ff45359403 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/calendarsharingmessage"
    i97f19227de75a13e8bce2a9f2f0fe826731ce9d14c85ded6e3178da92ad9f8e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/move"
    i98dfc01c9609a1f92d8d40b05cc6915fe1b5f42504125809b5cbadd5bac1be4a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/extensions"
    ia085204d95108fdfba3f30179e84fe9d56d2a41f0196b3faaccac5933897f4d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties"
    ia2c3cee1b082982c4046236c2e1484b1d12b01e7d202b5ad24986f82a79d1367 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/value"
    iaff702774cc3e61c13d7e163970e1629db453df271cbb59488849ed2a3ae0ec7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/createforward"
    ib47eb4bd511bd586fac1efd76227784d7f819cd4b070f8b6ebf528e44d879c62 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/copy"
    ic04bbf54cf6a3374e1844545b8f1fbb2c3e149b51fc1b499324271cb325f7593 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/forward"
    ic14505206271e848bd0b9454c10f984b045b8d95d2098c58abdc8853bee7b8e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/eventmessagerequest"
    ic2eb00d9cef4107c8d86f41b9ef42043e0038c1234cf6e8fb334c17ec24dd276 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/createreply"
    ic57b0d8e565c4fbad1c8bb7680c7f1bebe14d763e641b8aaef0cf7ed6e29b08a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/send"
    idb7f86c95ef4ba5d0a0f2d5e20325af74620e389e5d247f6967baca5d122864a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/reply"
    ide503fc3f46ba166f51a715b9dc0b7e9a83eaf22bc09b1d3634024b436aad454 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties"
    ie673455ab447de7fc89be141da7ff6bb95f9d92378eaefe3914fe3f9e387c4e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/unsubscribe"
    ie9c259a2102e27c149931d444b88050c501e3097576971fd37ec324913adb6b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/replyall"
    i12ca9076148c4a8f67d38f9a192b8c7b772df3687a3cc8863d5e5956f3e8e836 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/attachments/item"
    i402e59797825c6487d460e7a202a15eff3f3d7da3857b91b14cc5ac2341d77cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties/item"
    i8fd7cb6657cd70065c95e991e7c57ff53240dfa4ec292476f12a358369c42f03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties/item"
    ic6ebbe1487ad64bdb56672d7a43c44d420745ba51711b450fa500eb99d41f83b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/extensions/item"
    ie8c7eefceaadf180d29aab9898dd5e226c4bc43a528d2ec1beaac3b40d5d3747 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/mentions/item"
)

// MessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
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
// MessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
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
func (m *MessageItemRequestBuilder) Attachments()(*i18f0f3d80ebc26c14d0b298a950e35a82ff4632a691418331bc5c8ae9f7e1749.AttachmentsRequestBuilder) {
    return i18f0f3d80ebc26c14d0b298a950e35a82ff4632a691418331bc5c8ae9f7e1749.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*i12ca9076148c4a8f67d38f9a192b8c7b772df3687a3cc8863d5e5956f3e8e836.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i12ca9076148c4a8f67d38f9a192b8c7b772df3687a3cc8863d5e5956f3e8e836.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*i56e6bf2f80d7479fcf2b474c3916c0c636102e2f6f6405a2dd2c51ff45359403.CalendarSharingMessageRequestBuilder) {
    return i56e6bf2f80d7479fcf2b474c3916c0c636102e2f6f6405a2dd2c51ff45359403.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder_id}/childFolders/{mailFolder_id1}/messages/{message_id}{?select,expand}";
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
func (m *MessageItemRequestBuilder) Content()(*ia2c3cee1b082982c4046236c2e1484b1d12b01e7d202b5ad24986f82a79d1367.ContentRequestBuilder) {
    return ia2c3cee1b082982c4046236c2e1484b1d12b01e7d202b5ad24986f82a79d1367.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Copy()(*ib47eb4bd511bd586fac1efd76227784d7f819cd4b070f8b6ebf528e44d879c62.CopyRequestBuilder) {
    return ib47eb4bd511bd586fac1efd76227784d7f819cd4b070f8b6ebf528e44d879c62.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageItemRequestBuilder) CreateForward()(*iaff702774cc3e61c13d7e163970e1629db453df271cbb59488849ed2a3ae0ec7.CreateForwardRequestBuilder) {
    return iaff702774cc3e61c13d7e163970e1629db453df271cbb59488849ed2a3ae0ec7.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
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
func (m *MessageItemRequestBuilder) CreateReply()(*ic2eb00d9cef4107c8d86f41b9ef42043e0038c1234cf6e8fb334c17ec24dd276.CreateReplyRequestBuilder) {
    return ic2eb00d9cef4107c8d86f41b9ef42043e0038c1234cf6e8fb334c17ec24dd276.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CreateReplyAll()(*i38bd3b59f483648cf474b3f78ca224a4030f1ccdc95d8b3eca57afa7733c4cdc.CreateReplyAllRequestBuilder) {
    return i38bd3b59f483648cf474b3f78ca224a4030f1ccdc95d8b3eca57afa7733c4cdc.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageItemRequestBuilder) EventMessageRequest()(*ic14505206271e848bd0b9454c10f984b045b8d95d2098c58abdc8853bee7b8e9.EventMessageRequestRequestBuilder) {
    return ic14505206271e848bd0b9454c10f984b045b8d95d2098c58abdc8853bee7b8e9.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Extensions()(*i98dfc01c9609a1f92d8d40b05cc6915fe1b5f42504125809b5cbadd5bac1be4a.ExtensionsRequestBuilder) {
    return i98dfc01c9609a1f92d8d40b05cc6915fe1b5f42504125809b5cbadd5bac1be4a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*ic6ebbe1487ad64bdb56672d7a43c44d420745ba51711b450fa500eb99d41f83b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic6ebbe1487ad64bdb56672d7a43c44d420745ba51711b450fa500eb99d41f83b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Forward()(*ic04bbf54cf6a3374e1844545b8f1fbb2c3e149b51fc1b499324271cb325f7593.ForwardRequestBuilder) {
    return ic04bbf54cf6a3374e1844545b8f1fbb2c3e149b51fc1b499324271cb325f7593.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
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
func (m *MessageItemRequestBuilder) Mentions()(*i25ba706e121e0ece5ebf98516c94795df9564057f989070249206d4cfb86fd93.MentionsRequestBuilder) {
    return i25ba706e121e0ece5ebf98516c94795df9564057f989070249206d4cfb86fd93.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.mentions.item collection
func (m *MessageItemRequestBuilder) MentionsById(id string)(*ie8c7eefceaadf180d29aab9898dd5e226c4bc43a528d2ec1beaac3b40d5d3747.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return ie8c7eefceaadf180d29aab9898dd5e226c4bc43a528d2ec1beaac3b40d5d3747.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Move()(*i97f19227de75a13e8bce2a9f2f0fe826731ce9d14c85ded6e3178da92ad9f8e1.MoveRequestBuilder) {
    return i97f19227de75a13e8bce2a9f2f0fe826731ce9d14c85ded6e3178da92ad9f8e1.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*ia085204d95108fdfba3f30179e84fe9d56d2a41f0196b3faaccac5933897f4d8.MultiValueExtendedPropertiesRequestBuilder) {
    return ia085204d95108fdfba3f30179e84fe9d56d2a41f0196b3faaccac5933897f4d8.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i402e59797825c6487d460e7a202a15eff3f3d7da3857b91b14cc5ac2341d77cb.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i402e59797825c6487d460e7a202a15eff3f3d7da3857b91b14cc5ac2341d77cb.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *MessageItemRequestBuilder) Reply()(*idb7f86c95ef4ba5d0a0f2d5e20325af74620e389e5d247f6967baca5d122864a.ReplyRequestBuilder) {
    return idb7f86c95ef4ba5d0a0f2d5e20325af74620e389e5d247f6967baca5d122864a.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) ReplyAll()(*ie9c259a2102e27c149931d444b88050c501e3097576971fd37ec324913adb6b7.ReplyAllRequestBuilder) {
    return ie9c259a2102e27c149931d444b88050c501e3097576971fd37ec324913adb6b7.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Send()(*ic57b0d8e565c4fbad1c8bb7680c7f1bebe14d763e641b8aaef0cf7ed6e29b08a.SendRequestBuilder) {
    return ic57b0d8e565c4fbad1c8bb7680c7f1bebe14d763e641b8aaef0cf7ed6e29b08a.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*ide503fc3f46ba166f51a715b9dc0b7e9a83eaf22bc09b1d3634024b436aad454.SingleValueExtendedPropertiesRequestBuilder) {
    return ide503fc3f46ba166f51a715b9dc0b7e9a83eaf22bc09b1d3634024b436aad454.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8fd7cb6657cd70065c95e991e7c57ff53240dfa4ec292476f12a358369c42f03.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i8fd7cb6657cd70065c95e991e7c57ff53240dfa4ec292476f12a358369c42f03.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Unsubscribe()(*ie673455ab447de7fc89be141da7ff6bb95f9d92378eaefe3914fe3f9e387c4e3.UnsubscribeRequestBuilder) {
    return ie673455ab447de7fc89be141da7ff6bb95f9d92378eaefe3914fe3f9e387c4e3.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
