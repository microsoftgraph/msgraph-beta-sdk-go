package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1bbccbf9426fe67fa6f6e0e82f4dc14a70426e09a2fcc024d5e65737db89046c "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/markchatreadforuser"
    i23589f3f748b89f7a29ae1df54df4f7df1caa88dd2c8baf068f6007161c231c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/members"
    i48acda2de821dcf6abda941549a7467829c8e46e9e93090c8a580c73cf41fd2c "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/tabs"
    i4bec072883218972d9ecf2d62b7b914fff709b6026c787b10def2ac654def45f "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/permissiongrants"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5c9a73c476a5addf7da2bf806ecd2a618c858959d6475118492c27ad046fd135 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/markchatunreadforuser"
    i5f9832a2ccbfad05d4a38c197ccbad6908c541610b76ee98871142a004a67381 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/messages"
    i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/lastmessagepreview"
    ia1c4e944487866ec7cc165c592a0c533e7de01646fc63da1a218b9e5193eed39 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/installedapps"
    ia8541a5f81434dfc5a41205ab203298be0f48ebea3f8e26bf486b07a9e94c53a "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/operations"
    ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/sendactivitynotification"
    ice45878899bc758c6f22c47f7c7a49091bd4e926117a30848b2f3324044c4026 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/hideforuser"
    ice67a28769d9dc460d8b8b694cac83721204ddcc51142921e76a4f618e8d6b3b "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/unhideforuser"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i6df2ce3633ff2cfd4dfa5779c3abecf4c08acad5103f1c0967265043d3f020b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/messages/item"
    i77b0138be5c0e0924bfa917dfb3aaa7ceeb0f4b1a36c37a0e5a330c8386fded0 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/tabs/item"
    id4f1ff0b95f623eb5dd88b3faaa90d8a8d15dc658762e14d6f20f043f7a41bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/operations/item"
    id9d0b7bf8dee04396567860b43345e6dbc5dcecbe79ddf717e30dd368560da73 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/permissiongrants/item"
    ie0ebf374d51813131d325ba2e354a6f6ad5ee1d5255e9f99cae418db4e0bece9 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/members/item"
    ie61dfb73c398f7dead8cc5dd9db4a47ecb07aade0608fc8c54a8403b80033d16 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/installedapps/item"
)

// ChatItemRequestBuilder provides operations to manage the collection of chat entities.
type ChatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChatItemRequestBuilderDeleteOptions options for Delete
type ChatItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatItemRequestBuilderGetOptions options for Get
type ChatItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChatItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatItemRequestBuilderGetQueryParameters get entity from chats by key
type ChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ChatItemRequestBuilderPatchOptions options for Patch
type ChatItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chatable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatItemRequestBuilder) {
    m := &ChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/chats/{chat_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from chats
func (m *ChatItemRequestBuilder) CreateDeleteRequestInformation(options *ChatItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from chats by key
func (m *ChatItemRequestBuilder) CreateGetRequestInformation(options *ChatItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in chats
func (m *ChatItemRequestBuilder) CreatePatchRequestInformation(options *ChatItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from chats
func (m *ChatItemRequestBuilder) Delete(options *ChatItemRequestBuilderDeleteOptions)(error) {
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
// Get get entity from chats by key
func (m *ChatItemRequestBuilder) Get(options *ChatItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateChatFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Chatable), nil
}
func (m *ChatItemRequestBuilder) HideForUser()(*ice45878899bc758c6f22c47f7c7a49091bd4e926117a30848b2f3324044c4026.HideForUserRequestBuilder) {
    return ice45878899bc758c6f22c47f7c7a49091bd4e926117a30848b2f3324044c4026.NewHideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) InstalledApps()(*ia1c4e944487866ec7cc165c592a0c533e7de01646fc63da1a218b9e5193eed39.InstalledAppsRequestBuilder) {
    return ia1c4e944487866ec7cc165c592a0c533e7de01646fc63da1a218b9e5193eed39.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.chats.item.installedApps.item collection
func (m *ChatItemRequestBuilder) InstalledAppsById(id string)(*ie61dfb73c398f7dead8cc5dd9db4a47ecb07aade0608fc8c54a8403b80033d16.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return ie61dfb73c398f7dead8cc5dd9db4a47ecb07aade0608fc8c54a8403b80033d16.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) LastMessagePreview()(*i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2.LastMessagePreviewRequestBuilder) {
    return i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2.NewLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) MarkChatReadForUser()(*i1bbccbf9426fe67fa6f6e0e82f4dc14a70426e09a2fcc024d5e65737db89046c.MarkChatReadForUserRequestBuilder) {
    return i1bbccbf9426fe67fa6f6e0e82f4dc14a70426e09a2fcc024d5e65737db89046c.NewMarkChatReadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) MarkChatUnreadForUser()(*i5c9a73c476a5addf7da2bf806ecd2a618c858959d6475118492c27ad046fd135.MarkChatUnreadForUserRequestBuilder) {
    return i5c9a73c476a5addf7da2bf806ecd2a618c858959d6475118492c27ad046fd135.NewMarkChatUnreadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Members()(*i23589f3f748b89f7a29ae1df54df4f7df1caa88dd2c8baf068f6007161c231c0.MembersRequestBuilder) {
    return i23589f3f748b89f7a29ae1df54df4f7df1caa88dd2c8baf068f6007161c231c0.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.chats.item.members.item collection
func (m *ChatItemRequestBuilder) MembersById(id string)(*ie0ebf374d51813131d325ba2e354a6f6ad5ee1d5255e9f99cae418db4e0bece9.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ie0ebf374d51813131d325ba2e354a6f6ad5ee1d5255e9f99cae418db4e0bece9.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Messages()(*i5f9832a2ccbfad05d4a38c197ccbad6908c541610b76ee98871142a004a67381.MessagesRequestBuilder) {
    return i5f9832a2ccbfad05d4a38c197ccbad6908c541610b76ee98871142a004a67381.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.chats.item.messages.item collection
func (m *ChatItemRequestBuilder) MessagesById(id string)(*i6df2ce3633ff2cfd4dfa5779c3abecf4c08acad5103f1c0967265043d3f020b9.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i6df2ce3633ff2cfd4dfa5779c3abecf4c08acad5103f1c0967265043d3f020b9.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Operations()(*ia8541a5f81434dfc5a41205ab203298be0f48ebea3f8e26bf486b07a9e94c53a.OperationsRequestBuilder) {
    return ia8541a5f81434dfc5a41205ab203298be0f48ebea3f8e26bf486b07a9e94c53a.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.chats.item.operations.item collection
func (m *ChatItemRequestBuilder) OperationsById(id string)(*id4f1ff0b95f623eb5dd88b3faaa90d8a8d15dc658762e14d6f20f043f7a41bd6.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation_id"] = id
    }
    return id4f1ff0b95f623eb5dd88b3faaa90d8a8d15dc658762e14d6f20f043f7a41bd6.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in chats
func (m *ChatItemRequestBuilder) Patch(options *ChatItemRequestBuilderPatchOptions)(error) {
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
func (m *ChatItemRequestBuilder) PermissionGrants()(*i4bec072883218972d9ecf2d62b7b914fff709b6026c787b10def2ac654def45f.PermissionGrantsRequestBuilder) {
    return i4bec072883218972d9ecf2d62b7b914fff709b6026c787b10def2ac654def45f.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.chats.item.permissionGrants.item collection
func (m *ChatItemRequestBuilder) PermissionGrantsById(id string)(*id9d0b7bf8dee04396567860b43345e6dbc5dcecbe79ddf717e30dd368560da73.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant_id"] = id
    }
    return id9d0b7bf8dee04396567860b43345e6dbc5dcecbe79ddf717e30dd368560da73.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) SendActivityNotification()(*ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7.SendActivityNotificationRequestBuilder) {
    return ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Tabs()(*i48acda2de821dcf6abda941549a7467829c8e46e9e93090c8a580c73cf41fd2c.TabsRequestBuilder) {
    return i48acda2de821dcf6abda941549a7467829c8e46e9e93090c8a580c73cf41fd2c.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.chats.item.tabs.item collection
func (m *ChatItemRequestBuilder) TabsById(id string)(*i77b0138be5c0e0924bfa917dfb3aaa7ceeb0f4b1a36c37a0e5a330c8386fded0.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i77b0138be5c0e0924bfa917dfb3aaa7ceeb0f4b1a36c37a0e5a330c8386fded0.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) UnhideForUser()(*ice67a28769d9dc460d8b8b694cac83721204ddcc51142921e76a4f618e8d6b3b.UnhideForUserRequestBuilder) {
    return ice67a28769d9dc460d8b8b694cac83721204ddcc51142921e76a4f618e8d6b3b.NewUnhideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
