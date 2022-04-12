package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1bbccbf9426fe67fa6f6e0e82f4dc14a70426e09a2fcc024d5e65737db89046c "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/markchatreadforuser"
    i23589f3f748b89f7a29ae1df54df4f7df1caa88dd2c8baf068f6007161c231c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/members"
    i48acda2de821dcf6abda941549a7467829c8e46e9e93090c8a580c73cf41fd2c "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/tabs"
    i4bec072883218972d9ecf2d62b7b914fff709b6026c787b10def2ac654def45f "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/permissiongrants"
    i5c9a73c476a5addf7da2bf806ecd2a618c858959d6475118492c27ad046fd135 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/markchatunreadforuser"
    i5f9832a2ccbfad05d4a38c197ccbad6908c541610b76ee98871142a004a67381 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/messages"
    i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/lastmessagepreview"
    ia1c4e944487866ec7cc165c592a0c533e7de01646fc63da1a218b9e5193eed39 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/installedapps"
    ia8541a5f81434dfc5a41205ab203298be0f48ebea3f8e26bf486b07a9e94c53a "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/operations"
    ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/sendactivitynotification"
    ice45878899bc758c6f22c47f7c7a49091bd4e926117a30848b2f3324044c4026 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/hideforuser"
    ice67a28769d9dc460d8b8b694cac83721204ddcc51142921e76a4f618e8d6b3b "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/unhideforuser"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatItemRequestBuilderDeleteOptions options for Delete
type ChatItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChatItemRequestBuilderGetOptions options for Get
type ChatItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChatItemRequestBuilderGetQueryParameters get entity from chats by key
type ChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatItemRequestBuilderPatchOptions options for Patch
type ChatItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatItemRequestBuilder) {
    m := &ChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/chats/{chat%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from chats
func (m *ChatItemRequestBuilder) CreateDeleteRequestInformation(options *ChatItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from chats by key
func (m *ChatItemRequestBuilder) CreateGetRequestInformation(options *ChatItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in chats
func (m *ChatItemRequestBuilder) CreatePatchRequestInformation(options *ChatItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
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
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from chats by key
func (m *ChatItemRequestBuilder) Get(options *ChatItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable), nil
}
// HideForUser the hideForUser property
func (m *ChatItemRequestBuilder) HideForUser()(*ice45878899bc758c6f22c47f7c7a49091bd4e926117a30848b2f3324044c4026.HideForUserRequestBuilder) {
    return ice45878899bc758c6f22c47f7c7a49091bd4e926117a30848b2f3324044c4026.NewHideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledApps the installedApps property
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
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return ie61dfb73c398f7dead8cc5dd9db4a47ecb07aade0608fc8c54a8403b80033d16.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LastMessagePreview the lastMessagePreview property
func (m *ChatItemRequestBuilder) LastMessagePreview()(*i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2.LastMessagePreviewRequestBuilder) {
    return i8a2fb47d75dd04219d9d1d15bca3e9e7ee5f68a38e11554619e97a6fb34a94b2.NewLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MarkChatReadForUser the markChatReadForUser property
func (m *ChatItemRequestBuilder) MarkChatReadForUser()(*i1bbccbf9426fe67fa6f6e0e82f4dc14a70426e09a2fcc024d5e65737db89046c.MarkChatReadForUserRequestBuilder) {
    return i1bbccbf9426fe67fa6f6e0e82f4dc14a70426e09a2fcc024d5e65737db89046c.NewMarkChatReadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MarkChatUnreadForUser the markChatUnreadForUser property
func (m *ChatItemRequestBuilder) MarkChatUnreadForUser()(*i5c9a73c476a5addf7da2bf806ecd2a618c858959d6475118492c27ad046fd135.MarkChatUnreadForUserRequestBuilder) {
    return i5c9a73c476a5addf7da2bf806ecd2a618c858959d6475118492c27ad046fd135.NewMarkChatUnreadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members the members property
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
        urlTplParams["conversationMember%2Did"] = id
    }
    return ie0ebf374d51813131d325ba2e354a6f6ad5ee1d5255e9f99cae418db4e0bece9.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
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
        urlTplParams["chatMessage%2Did"] = id
    }
    return i6df2ce3633ff2cfd4dfa5779c3abecf4c08acad5103f1c0967265043d3f020b9.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
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
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return id4f1ff0b95f623eb5dd88b3faaa90d8a8d15dc658762e14d6f20f043f7a41bd6.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in chats
func (m *ChatItemRequestBuilder) Patch(options *ChatItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PermissionGrants the permissionGrants property
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
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return id9d0b7bf8dee04396567860b43345e6dbc5dcecbe79ddf717e30dd368560da73.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendActivityNotification the sendActivityNotification property
func (m *ChatItemRequestBuilder) SendActivityNotification()(*ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7.SendActivityNotificationRequestBuilder) {
    return ib8fd861a1697792849ae9042379690c5fb26911a9afc665ed66bc94e066d53e7.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tabs the tabs property
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
        urlTplParams["teamsTab%2Did"] = id
    }
    return i77b0138be5c0e0924bfa917dfb3aaa7ceeb0f4b1a36c37a0e5a330c8386fded0.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnhideForUser the unhideForUser property
func (m *ChatItemRequestBuilder) UnhideForUser()(*ice67a28769d9dc460d8b8b694cac83721204ddcc51142921e76a4f618e8d6b3b.UnhideForUserRequestBuilder) {
    return ice67a28769d9dc460d8b8b694cac83721204ddcc51142921e76a4f618e8d6b3b.NewUnhideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
