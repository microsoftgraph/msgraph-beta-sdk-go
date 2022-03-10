package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i00cadd5bdea57495f568208d047694716e945ce53c288e67c05708de06827683 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties"
    i375d78994fa81e234d561e0c503e7f3e57a2cdec4348afefa3aad37bf5e7950e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/eventmessagerequest"
    i3a7166d614f37e53779085c27ccebc55d8ba7da54ce9e2467bb4ef84feb2bd4c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createforward"
    i3b845216e66b858f20515650fee13a45e57907540b3fef30c499b388000b593f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createreplyall"
    i412bc4d3ecb81a4bac7a52995cdbd130de5d929ef5eeedc316a33d88fea79866 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/attachments"
    i5a588eedeadfbc918f35890e674b70004e75980729a73df65ebc0508e1748596 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/move"
    i5bfd8f520c8c442c080004e492a7c27b37d2f4ba1c8f714496c29b07f50e81e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/unsubscribe"
    i7e23fb7052b79b47fe07fd4f07db233d6a980543604b40b302c8879839e25102 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/send"
    i83e014dcfbda4d03116a64e01e601cbae72ba7b630422f5e635bfe1232f9a25c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/replyall"
    i9b97ea3ffdf3d74eddb0a6192d3fe8c03634840e6917a447984a672288a61edf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/extensions"
    i9c01bcaebd609b85ba8581e3a4e4c71b4ef27495c83bf3a764ce4dbce9542e1b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/value"
    i9e698067a956139684fc5f333b26f76bdf7eba5aef5f8e685f4a812d8eda85a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/mentions"
    ic696f4baefc6d45a63b0236eb3b86070c596446ad210c49354d50d95abc9446f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/forward"
    ic9f613d805e74cd01c80cdae89dd189e97bed8805284bde26421b192080b975c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/calendarsharingmessage"
    icc3a33f018c9d6725fa2b89cca1d94d14ba987c2fb3824e00cef071a7cc9c4bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/createreply"
    id97b39ca2a01c3891d8b6b2d850bc6e198a70689a4ebb3fd7975c962fd6e4d18 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties"
    iefd49125a49593760cf161984131bff95b4bf49a63a3ab136b070ddf347adb24 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/reply"
    ifa5c40480383bdfb7990bea8312f8d4037d9fde7c2d940a69afe370e15a3d702 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/copy"
    i075faccccc7205b32bfc44721cfd5ca685b307dfe965b3c03f26947de7606a21 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/mentions/item"
    i6ea2e8d28e24b8938d5023a1eab602b1ea6653bc7bddc5811bdd554bd2f897f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/attachments/item"
    i71143bcc60f051851219edfad7faec4f6dc0a6505548d4921c7d1f0df0e6a28a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties/item"
    ic8cbcc35f6513a4c89985dfc6121977fbc90be0e866483fb78cea71a1e7c27c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties/item"
    id611f3d66a9657a64b68f9aff10db2bf2f1d4523dda945f850b6588960d403a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/extensions/item"
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
func (m *MessageItemRequestBuilder) Attachments()(*i412bc4d3ecb81a4bac7a52995cdbd130de5d929ef5eeedc316a33d88fea79866.AttachmentsRequestBuilder) {
    return i412bc4d3ecb81a4bac7a52995cdbd130de5d929ef5eeedc316a33d88fea79866.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*i6ea2e8d28e24b8938d5023a1eab602b1ea6653bc7bddc5811bdd554bd2f897f6.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i6ea2e8d28e24b8938d5023a1eab602b1ea6653bc7bddc5811bdd554bd2f897f6.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CalendarSharingMessage()(*ic9f613d805e74cd01c80cdae89dd189e97bed8805284bde26421b192080b975c.CalendarSharingMessageRequestBuilder) {
    return ic9f613d805e74cd01c80cdae89dd189e97bed8805284bde26421b192080b975c.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/childFolders/{mailFolder_id1}/messages/{message_id}{?select,expand}";
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
func (m *MessageItemRequestBuilder) Content()(*i9c01bcaebd609b85ba8581e3a4e4c71b4ef27495c83bf3a764ce4dbce9542e1b.ContentRequestBuilder) {
    return i9c01bcaebd609b85ba8581e3a4e4c71b4ef27495c83bf3a764ce4dbce9542e1b.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Copy()(*ifa5c40480383bdfb7990bea8312f8d4037d9fde7c2d940a69afe370e15a3d702.CopyRequestBuilder) {
    return ifa5c40480383bdfb7990bea8312f8d4037d9fde7c2d940a69afe370e15a3d702.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for users
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
func (m *MessageItemRequestBuilder) CreateForward()(*i3a7166d614f37e53779085c27ccebc55d8ba7da54ce9e2467bb4ef84feb2bd4c.CreateForwardRequestBuilder) {
    return i3a7166d614f37e53779085c27ccebc55d8ba7da54ce9e2467bb4ef84feb2bd4c.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreatePatchRequestInformation update the navigation property messages in users
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
func (m *MessageItemRequestBuilder) CreateReply()(*icc3a33f018c9d6725fa2b89cca1d94d14ba987c2fb3824e00cef071a7cc9c4bc.CreateReplyRequestBuilder) {
    return icc3a33f018c9d6725fa2b89cca1d94d14ba987c2fb3824e00cef071a7cc9c4bc.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) CreateReplyAll()(*i3b845216e66b858f20515650fee13a45e57907540b3fef30c499b388000b593f.CreateReplyAllRequestBuilder) {
    return i3b845216e66b858f20515650fee13a45e57907540b3fef30c499b388000b593f.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for users
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
func (m *MessageItemRequestBuilder) EventMessageRequest()(*i375d78994fa81e234d561e0c503e7f3e57a2cdec4348afefa3aad37bf5e7950e.EventMessageRequestRequestBuilder) {
    return i375d78994fa81e234d561e0c503e7f3e57a2cdec4348afefa3aad37bf5e7950e.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Extensions()(*i9b97ea3ffdf3d74eddb0a6192d3fe8c03634840e6917a447984a672288a61edf.ExtensionsRequestBuilder) {
    return i9b97ea3ffdf3d74eddb0a6192d3fe8c03634840e6917a447984a672288a61edf.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*id611f3d66a9657a64b68f9aff10db2bf2f1d4523dda945f850b6588960d403a0.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return id611f3d66a9657a64b68f9aff10db2bf2f1d4523dda945f850b6588960d403a0.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Forward()(*ic696f4baefc6d45a63b0236eb3b86070c596446ad210c49354d50d95abc9446f.ForwardRequestBuilder) {
    return ic696f4baefc6d45a63b0236eb3b86070c596446ad210c49354d50d95abc9446f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageItemRequestBuilder) Mentions()(*i9e698067a956139684fc5f333b26f76bdf7eba5aef5f8e685f4a812d8eda85a1.MentionsRequestBuilder) {
    return i9e698067a956139684fc5f333b26f76bdf7eba5aef5f8e685f4a812d8eda85a1.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.mentions.item collection
func (m *MessageItemRequestBuilder) MentionsById(id string)(*i075faccccc7205b32bfc44721cfd5ca685b307dfe965b3c03f26947de7606a21.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return i075faccccc7205b32bfc44721cfd5ca685b307dfe965b3c03f26947de7606a21.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Move()(*i5a588eedeadfbc918f35890e674b70004e75980729a73df65ebc0508e1748596.MoveRequestBuilder) {
    return i5a588eedeadfbc918f35890e674b70004e75980729a73df65ebc0508e1748596.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*i00cadd5bdea57495f568208d047694716e945ce53c288e67c05708de06827683.MultiValueExtendedPropertiesRequestBuilder) {
    return i00cadd5bdea57495f568208d047694716e945ce53c288e67c05708de06827683.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic8cbcc35f6513a4c89985dfc6121977fbc90be0e866483fb78cea71a1e7c27c6.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ic8cbcc35f6513a4c89985dfc6121977fbc90be0e866483fb78cea71a1e7c27c6.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in users
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
func (m *MessageItemRequestBuilder) Reply()(*iefd49125a49593760cf161984131bff95b4bf49a63a3ab136b070ddf347adb24.ReplyRequestBuilder) {
    return iefd49125a49593760cf161984131bff95b4bf49a63a3ab136b070ddf347adb24.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) ReplyAll()(*i83e014dcfbda4d03116a64e01e601cbae72ba7b630422f5e635bfe1232f9a25c.ReplyAllRequestBuilder) {
    return i83e014dcfbda4d03116a64e01e601cbae72ba7b630422f5e635bfe1232f9a25c.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Send()(*i7e23fb7052b79b47fe07fd4f07db233d6a980543604b40b302c8879839e25102.SendRequestBuilder) {
    return i7e23fb7052b79b47fe07fd4f07db233d6a980543604b40b302c8879839e25102.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*id97b39ca2a01c3891d8b6b2d850bc6e198a70689a4ebb3fd7975c962fd6e4d18.SingleValueExtendedPropertiesRequestBuilder) {
    return id97b39ca2a01c3891d8b6b2d850bc6e198a70689a4ebb3fd7975c962fd6e4d18.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i71143bcc60f051851219edfad7faec4f6dc0a6505548d4921c7d1f0df0e6a28a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i71143bcc60f051851219edfad7faec4f6dc0a6505548d4921c7d1f0df0e6a28a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageItemRequestBuilder) Unsubscribe()(*i5bfd8f520c8c442c080004e492a7c27b37d2f4ba1c8f714496c29b07f50e81e8.UnsubscribeRequestBuilder) {
    return i5bfd8f520c8c442c080004e492a7c27b37d2f4ba1c8f714496c29b07f50e81e8.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
