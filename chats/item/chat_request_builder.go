package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i23589f3f748b89f7a29ae1df54df4f7df1caa88dd2c8baf068f6007161c231c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/members"
    i48acda2de821dcf6abda941549a7467829c8e46e9e93090c8a580c73cf41fd2c "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/tabs"
    i4bec072883218972d9ecf2d62b7b914fff709b6026c787b10def2ac654def45f "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/permissiongrants"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5f9832a2ccbfad05d4a38c197ccbad6908c541610b76ee98871142a004a67381 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/messages"
    i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/lastmessagepreview"
    ia1c4e944487866ec7cc165c592a0c533e7de01646fc63da1a218b9e5193eed39 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/installedapps"
    ia8541a5f81434dfc5a41205ab203298be0f48ebea3f8e26bf486b07a9e94c53a "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/operations"
    ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/sendactivitynotification"
    i6df2ce3633ff2cfd4dfa5779c3abecf4c08acad5103f1c0967265043d3f020b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/messages/item"
    i77b0138be5c0e0924bfa917dfb3aaa7ceeb0f4b1a36c37a0e5a330c8386fded0 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/tabs/item"
    id4f1ff0b95f623eb5dd88b3faaa90d8a8d15dc658762e14d6f20f043f7a41bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/operations/item"
    id9d0b7bf8dee04396567860b43345e6dbc5dcecbe79ddf717e30dd368560da73 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/permissiongrants/item"
    ie0ebf374d51813131d325ba2e354a6f6ad5ee1d5255e9f99cae418db4e0bece9 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/members/item"
    ie61dfb73c398f7dead8cc5dd9db4a47ecb07aade0608fc8c54a8403b80033d16 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/installedapps/item"
)

type ChatRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ChatRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatRequestBuilder) {
    m := &ChatRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/chats/{chat_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewChatRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ChatRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ChatRequestBuilder) CreateGetRequestInformation(q func (value *ChatRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ChatRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ChatRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chat, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ChatRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChatRequestBuilder) Get(q func (value *ChatRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chat, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewChat() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chat), nil
}
func (m *ChatRequestBuilder) InstalledApps()(*ia1c4e944487866ec7cc165c592a0c533e7de01646fc63da1a218b9e5193eed39.InstalledAppsRequestBuilder) {
    return ia1c4e944487866ec7cc165c592a0c533e7de01646fc63da1a218b9e5193eed39.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) InstalledAppsById(id string)(*ie61dfb73c398f7dead8cc5dd9db4a47ecb07aade0608fc8c54a8403b80033d16.TeamsAppInstallationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return ie61dfb73c398f7dead8cc5dd9db4a47ecb07aade0608fc8c54a8403b80033d16.NewTeamsAppInstallationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatRequestBuilder) LastMessagePreview()(*i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2.LastMessagePreviewRequestBuilder) {
    return i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2.NewLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) Members()(*i23589f3f748b89f7a29ae1df54df4f7df1caa88dd2c8baf068f6007161c231c0.MembersRequestBuilder) {
    return i23589f3f748b89f7a29ae1df54df4f7df1caa88dd2c8baf068f6007161c231c0.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) MembersById(id string)(*ie0ebf374d51813131d325ba2e354a6f6ad5ee1d5255e9f99cae418db4e0bece9.ConversationMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ie0ebf374d51813131d325ba2e354a6f6ad5ee1d5255e9f99cae418db4e0bece9.NewConversationMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatRequestBuilder) Messages()(*i5f9832a2ccbfad05d4a38c197ccbad6908c541610b76ee98871142a004a67381.MessagesRequestBuilder) {
    return i5f9832a2ccbfad05d4a38c197ccbad6908c541610b76ee98871142a004a67381.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) MessagesById(id string)(*i6df2ce3633ff2cfd4dfa5779c3abecf4c08acad5103f1c0967265043d3f020b9.ChatMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i6df2ce3633ff2cfd4dfa5779c3abecf4c08acad5103f1c0967265043d3f020b9.NewChatMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatRequestBuilder) Operations()(*ia8541a5f81434dfc5a41205ab203298be0f48ebea3f8e26bf486b07a9e94c53a.OperationsRequestBuilder) {
    return ia8541a5f81434dfc5a41205ab203298be0f48ebea3f8e26bf486b07a9e94c53a.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) OperationsById(id string)(*id4f1ff0b95f623eb5dd88b3faaa90d8a8d15dc658762e14d6f20f043f7a41bd6.TeamsAsyncOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return id4f1ff0b95f623eb5dd88b3faaa90d8a8d15dc658762e14d6f20f043f7a41bd6.NewTeamsAsyncOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chat, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChatRequestBuilder) PermissionGrants()(*i4bec072883218972d9ecf2d62b7b914fff709b6026c787b10def2ac654def45f.PermissionGrantsRequestBuilder) {
    return i4bec072883218972d9ecf2d62b7b914fff709b6026c787b10def2ac654def45f.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) PermissionGrantsById(id string)(*id9d0b7bf8dee04396567860b43345e6dbc5dcecbe79ddf717e30dd368560da73.ResourceSpecificPermissionGrantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant_id"] = id
    }
    return id9d0b7bf8dee04396567860b43345e6dbc5dcecbe79ddf717e30dd368560da73.NewResourceSpecificPermissionGrantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatRequestBuilder) SendActivityNotification()(*ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7.SendActivityNotificationRequestBuilder) {
    return ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) Tabs()(*i48acda2de821dcf6abda941549a7467829c8e46e9e93090c8a580c73cf41fd2c.TabsRequestBuilder) {
    return i48acda2de821dcf6abda941549a7467829c8e46e9e93090c8a580c73cf41fd2c.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) TabsById(id string)(*i77b0138be5c0e0924bfa917dfb3aaa7ceeb0f4b1a36c37a0e5a330c8386fded0.TeamsTabRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i77b0138be5c0e0924bfa917dfb3aaa7ceeb0f4b1a36c37a0e5a330c8386fded0.NewTeamsTabRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
