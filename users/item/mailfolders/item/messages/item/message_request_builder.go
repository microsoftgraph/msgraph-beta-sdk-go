package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1388a5b7afebd5c1d7ad7238e86dfa05a923b4d8e35dc4265272c52df5b7ab94 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/move"
    i3c0595e344f6094e852e3be4813c86bfd5e14109889db8d73b847fbb96d50e80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/createreply"
    i58c0a5d3e1252e77598ba1ea581032b7c028f21319ac81bacba8d36aaae8d7c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/forward"
    i779693a4ed700a6242a8750d06dd0c544dee12d3a8a544cfcfd6a28155e1ed26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/extensions"
    i80e9f30d67f4b0e96cf2e956e67ef07179485a391069b5616f4763663728fad5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/copy"
    i884f837848c6f0c6423949ee4d2944c09197989edfb9edbb71b2db93ccab29ba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/multivalueextendedproperties"
    i8d70e83d3fc5a5d914de83826ade66aa9b91ec1bd829919d84e5e87c18949a9e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/singlevalueextendedproperties"
    i98928d52ac1ce3e3977d623196257b1893e884de30d59da419c72f7a6d5748fb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/eventmessagerequest"
    ia41ad1f56fdb6d2218a17ab1c1042e82ae808d6e3db1a17d6944323659bb604a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/value"
    iaa839673fe6539572b711a4256f3b133ae2b587343b870550dd45c8417642743 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/calendarsharingmessage"
    iaaf573b089683872c986a56013f03c41df00859342d210ce296d5f9be8c32fa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/send"
    ib48014d797abbb97d18c2fa3c3e893092e5b77db3345b6160da6647afde9d16d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/createreplyall"
    icca11b370ecbb3da8483d6ba09a0c88648c6a354801f5c04b32d16f550a2c1b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/replyall"
    id21d4bcffeb94fbdc758f4d604a94503e400e91d464e2ea5b2878b43f841f829 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/mentions"
    id9eb0bcd2e9ba6811167e83aa7fba832e7bdcccbf038736c6356977466d233b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/reply"
    idb1ae05cc149928c525695427be582f6e0764cb0ec6bd45bd0e9b145823d8389 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/createforward"
    ie25dbd4775466b7a05037c48ef048e5a9634efd19de463e40384afaa61eb7bcc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/attachments"
    ife8a43d0f9a712fa0dd58ac70aab173d38980c2d93aa4229d9760c4c0907335b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/unsubscribe"
    i41df3c0d8cbd67958e95ec7d8e5e8caeba6577f510da216d6243755d4d47e60b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/extensions/item"
    i758243302df29cce2fe9041a95e536b614e1bce0fa181ec51735b45c28450aac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/attachments/item"
    i8914b25cf327deb26ccf2e46fc7a562348efa5e5885764ce98f231989db617ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/multivalueextendedproperties/item"
    i9617bb1755b63307a9fbb614af96ed336375ccfbf08d3da6a1dd89acbde80d5b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/singlevalueextendedproperties/item"
    idb97761d489a095537daad3218e84b76808b5d8fc51cba00b233be4c0e0f9e9f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/mentions/item"
)

// MessageRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}
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
func (m *MessageRequestBuilder) Attachments()(*ie25dbd4775466b7a05037c48ef048e5a9634efd19de463e40384afaa61eb7bcc.AttachmentsRequestBuilder) {
    return ie25dbd4775466b7a05037c48ef048e5a9634efd19de463e40384afaa61eb7bcc.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.attachments.item collection
func (m *MessageRequestBuilder) AttachmentsById(id string)(*i758243302df29cce2fe9041a95e536b614e1bce0fa181ec51735b45c28450aac.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i758243302df29cce2fe9041a95e536b614e1bce0fa181ec51735b45c28450aac.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) CalendarSharingMessage()(*iaa839673fe6539572b711a4256f3b133ae2b587343b870550dd45c8417642743.CalendarSharingMessageRequestBuilder) {
    return iaa839673fe6539572b711a4256f3b133ae2b587343b870550dd45c8417642743.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessageRequestBuilderInternal instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}{?select,expand}";
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
func (m *MessageRequestBuilder) Content()(*ia41ad1f56fdb6d2218a17ab1c1042e82ae808d6e3db1a17d6944323659bb604a.ContentRequestBuilder) {
    return ia41ad1f56fdb6d2218a17ab1c1042e82ae808d6e3db1a17d6944323659bb604a.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Copy()(*i80e9f30d67f4b0e96cf2e956e67ef07179485a391069b5616f4763663728fad5.CopyRequestBuilder) {
    return i80e9f30d67f4b0e96cf2e956e67ef07179485a391069b5616f4763663728fad5.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageRequestBuilder) CreateForward()(*idb1ae05cc149928c525695427be582f6e0764cb0ec6bd45bd0e9b145823d8389.CreateForwardRequestBuilder) {
    return idb1ae05cc149928c525695427be582f6e0764cb0ec6bd45bd0e9b145823d8389.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageRequestBuilder) CreateReply()(*i3c0595e344f6094e852e3be4813c86bfd5e14109889db8d73b847fbb96d50e80.CreateReplyRequestBuilder) {
    return i3c0595e344f6094e852e3be4813c86bfd5e14109889db8d73b847fbb96d50e80.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) CreateReplyAll()(*ib48014d797abbb97d18c2fa3c3e893092e5b77db3345b6160da6647afde9d16d.CreateReplyAllRequestBuilder) {
    return ib48014d797abbb97d18c2fa3c3e893092e5b77db3345b6160da6647afde9d16d.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageRequestBuilder) EventMessageRequest()(*i98928d52ac1ce3e3977d623196257b1893e884de30d59da419c72f7a6d5748fb.EventMessageRequestRequestBuilder) {
    return i98928d52ac1ce3e3977d623196257b1893e884de30d59da419c72f7a6d5748fb.NewEventMessageRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Extensions()(*i779693a4ed700a6242a8750d06dd0c544dee12d3a8a544cfcfd6a28155e1ed26.ExtensionsRequestBuilder) {
    return i779693a4ed700a6242a8750d06dd0c544dee12d3a8a544cfcfd6a28155e1ed26.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.extensions.item collection
func (m *MessageRequestBuilder) ExtensionsById(id string)(*i41df3c0d8cbd67958e95ec7d8e5e8caeba6577f510da216d6243755d4d47e60b.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i41df3c0d8cbd67958e95ec7d8e5e8caeba6577f510da216d6243755d4d47e60b.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Forward()(*i58c0a5d3e1252e77598ba1ea581032b7c028f21319ac81bacba8d36aaae8d7c9.ForwardRequestBuilder) {
    return i58c0a5d3e1252e77598ba1ea581032b7c028f21319ac81bacba8d36aaae8d7c9.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *MessageRequestBuilder) Mentions()(*id21d4bcffeb94fbdc758f4d604a94503e400e91d464e2ea5b2878b43f841f829.MentionsRequestBuilder) {
    return id21d4bcffeb94fbdc758f4d604a94503e400e91d464e2ea5b2878b43f841f829.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.mentions.item collection
func (m *MessageRequestBuilder) MentionsById(id string)(*idb97761d489a095537daad3218e84b76808b5d8fc51cba00b233be4c0e0f9e9f.MentionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return idb97761d489a095537daad3218e84b76808b5d8fc51cba00b233be4c0e0f9e9f.NewMentionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Move()(*i1388a5b7afebd5c1d7ad7238e86dfa05a923b4d8e35dc4265272c52df5b7ab94.MoveRequestBuilder) {
    return i1388a5b7afebd5c1d7ad7238e86dfa05a923b4d8e35dc4265272c52df5b7ab94.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*i884f837848c6f0c6423949ee4d2944c09197989edfb9edbb71b2db93ccab29ba.MultiValueExtendedPropertiesRequestBuilder) {
    return i884f837848c6f0c6423949ee4d2944c09197989edfb9edbb71b2db93ccab29ba.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i8914b25cf327deb26ccf2e46fc7a562348efa5e5885764ce98f231989db617ab.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i8914b25cf327deb26ccf2e46fc7a562348efa5e5885764ce98f231989db617ab.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *MessageRequestBuilder) Reply()(*id9eb0bcd2e9ba6811167e83aa7fba832e7bdcccbf038736c6356977466d233b9.ReplyRequestBuilder) {
    return id9eb0bcd2e9ba6811167e83aa7fba832e7bdcccbf038736c6356977466d233b9.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) ReplyAll()(*icca11b370ecbb3da8483d6ba09a0c88648c6a354801f5c04b32d16f550a2c1b9.ReplyAllRequestBuilder) {
    return icca11b370ecbb3da8483d6ba09a0c88648c6a354801f5c04b32d16f550a2c1b9.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) Send()(*iaaf573b089683872c986a56013f03c41df00859342d210ce296d5f9be8c32fa3.SendRequestBuilder) {
    return iaaf573b089683872c986a56013f03c41df00859342d210ce296d5f9be8c32fa3.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*i8d70e83d3fc5a5d914de83826ade66aa9b91ec1bd829919d84e5e87c18949a9e.SingleValueExtendedPropertiesRequestBuilder) {
    return i8d70e83d3fc5a5d914de83826ade66aa9b91ec1bd829919d84e5e87c18949a9e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9617bb1755b63307a9fbb614af96ed336375ccfbf08d3da6a1dd89acbde80d5b.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9617bb1755b63307a9fbb614af96ed336375ccfbf08d3da6a1dd89acbde80d5b.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MessageRequestBuilder) Unsubscribe()(*ife8a43d0f9a712fa0dd58ac70aab173d38980c2d93aa4229d9760c4c0907335b.UnsubscribeRequestBuilder) {
    return ife8a43d0f9a712fa0dd58ac70aab173d38980c2d93aa4229d9760c4c0907335b.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
