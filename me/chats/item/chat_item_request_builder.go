package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0bafb1debfff9bf875c891125572344c403e42dae31d9f4889de4d6f1774a3f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/markchatunreadforuser"
    i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/operations"
    i2b860dec50e5faa8f825fb8103a2b27cedc04224865564cb96149aa6bb97d7d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/sendactivitynotification"
    i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/permissiongrants"
    i75244b45b923891db6e9b0b23e8f43de1095113f335d4507e9c0b54ebcd0c21e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/markchatreadforuser"
    i81f518d742ed9b6a68e2df0007e45cc345d66defa6db59a94afc19f8c01870d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/hideforuser"
    i8791640889cb2a0ec1b53c3db460f3f7c82671cef2a627f7fc2649c53e0bb14e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/unhideforuser"
    ia4a5a30ed707af790b1c2b15936fcf68a4ac1eaee6407d4e5837f9dee33d2717 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/installedapps"
    ib32e501e2ee14b9c518eef16cdd1655db30b5806e5f4be17957c21a66e36c695 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/pinnedmessages"
    ibcd033263282f46cf08c9ffed250e454bfe93a5f0d41ede7466f3725b012808c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/members"
    idead0084f5300487c412f0afd4cb69c96dab7aa445fdb45feba337598ba6e4a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/lastmessagepreview"
    if345d5c48676b774a18852c60a5c6a74c7235e1f5439793ed859b22d9c6e5d11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/tabs"
    if873fe971f2b52c7a520477d31db735634ccbf62cf832a0dacc1f30a21e89298 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages"
    i1608545c1b63cdff025064ae1096d8c51ba0eae3a6adeffca8ec30d3f85d8859 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/tabs/item"
    i1da914fe8033359b6525954501f6aeb6fc0ad41d92f250fea01dc364798e325a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages/item"
    i7cae322e7e19ac03ee3835d6259dde3546087d9bc77605923cd230b57741a237 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/installedapps/item"
    i9ac05382512bc3e53d355e9c8817600d551e3fbb197180b2be3832f4e968befa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/permissiongrants/item"
    ib7358fe28b26b0df886473cc1c746ecee1f749db40a9123335298a7e8ff8c07b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/pinnedmessages/item"
    ibcd7e2b55750c3cf3853ab13d9a8350fa0813cc9dafb723b8df641e30a377aed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/members/item"
    iee8e28251b196a5ca3014f3cb541ead07ac22409864668375196a8ca8542d1a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/operations/item"
)

// ChatItemRequestBuilder provides operations to manage the chats property of the microsoft.graph.user entity.
type ChatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChatItemRequestBuilderGetQueryParameters get chats from me
type ChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatItemRequestBuilderGetQueryParameters
}
// ChatItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatItemRequestBuilder) {
    m := &ChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property chats for me
func (m *ChatItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ChatItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get chats from me
func (m *ChatItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ChatItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property chats in me
func (m *ChatItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, requestConfiguration *ChatItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property chats for me
func (m *ChatItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChatItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get chats from me
func (m *ChatItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChatItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable), nil
}
// HideForUser provides operations to call the hideForUser method.
func (m *ChatItemRequestBuilder) HideForUser()(*i81f518d742ed9b6a68e2df0007e45cc345d66defa6db59a94afc19f8c01870d9.HideForUserRequestBuilder) {
    return i81f518d742ed9b6a68e2df0007e45cc345d66defa6db59a94afc19f8c01870d9.NewHideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) InstalledApps()(*ia4a5a30ed707af790b1c2b15936fcf68a4ac1eaee6407d4e5837f9dee33d2717.InstalledAppsRequestBuilder) {
    return ia4a5a30ed707af790b1c2b15936fcf68a4ac1eaee6407d4e5837f9dee33d2717.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) InstalledAppsById(id string)(*i7cae322e7e19ac03ee3835d6259dde3546087d9bc77605923cd230b57741a237.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return i7cae322e7e19ac03ee3835d6259dde3546087d9bc77605923cd230b57741a237.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LastMessagePreview provides operations to manage the lastMessagePreview property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) LastMessagePreview()(*idead0084f5300487c412f0afd4cb69c96dab7aa445fdb45feba337598ba6e4a4.LastMessagePreviewRequestBuilder) {
    return idead0084f5300487c412f0afd4cb69c96dab7aa445fdb45feba337598ba6e4a4.NewLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MarkChatReadForUser provides operations to call the markChatReadForUser method.
func (m *ChatItemRequestBuilder) MarkChatReadForUser()(*i75244b45b923891db6e9b0b23e8f43de1095113f335d4507e9c0b54ebcd0c21e.MarkChatReadForUserRequestBuilder) {
    return i75244b45b923891db6e9b0b23e8f43de1095113f335d4507e9c0b54ebcd0c21e.NewMarkChatReadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MarkChatUnreadForUser provides operations to call the markChatUnreadForUser method.
func (m *ChatItemRequestBuilder) MarkChatUnreadForUser()(*i0bafb1debfff9bf875c891125572344c403e42dae31d9f4889de4d6f1774a3f1.MarkChatUnreadForUserRequestBuilder) {
    return i0bafb1debfff9bf875c891125572344c403e42dae31d9f4889de4d6f1774a3f1.NewMarkChatUnreadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) Members()(*ibcd033263282f46cf08c9ffed250e454bfe93a5f0d41ede7466f3725b012808c.MembersRequestBuilder) {
    return ibcd033263282f46cf08c9ffed250e454bfe93a5f0d41ede7466f3725b012808c.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) MembersById(id string)(*ibcd7e2b55750c3cf3853ab13d9a8350fa0813cc9dafb723b8df641e30a377aed.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return ibcd7e2b55750c3cf3853ab13d9a8350fa0813cc9dafb723b8df641e30a377aed.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) Messages()(*if873fe971f2b52c7a520477d31db735634ccbf62cf832a0dacc1f30a21e89298.MessagesRequestBuilder) {
    return if873fe971f2b52c7a520477d31db735634ccbf62cf832a0dacc1f30a21e89298.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) MessagesById(id string)(*i1da914fe8033359b6525954501f6aeb6fc0ad41d92f250fea01dc364798e325a.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i1da914fe8033359b6525954501f6aeb6fc0ad41d92f250fea01dc364798e325a.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) Operations()(*i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35.OperationsRequestBuilder) {
    return i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) OperationsById(id string)(*iee8e28251b196a5ca3014f3cb541ead07ac22409864668375196a8ca8542d1a9.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return iee8e28251b196a5ca3014f3cb541ead07ac22409864668375196a8ca8542d1a9.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property chats in me
func (m *ChatItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, requestConfiguration *ChatItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable), nil
}
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) PermissionGrants()(*i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f.PermissionGrantsRequestBuilder) {
    return i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the permissionGrants property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) PermissionGrantsById(id string)(*i9ac05382512bc3e53d355e9c8817600d551e3fbb197180b2be3832f4e968befa.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i9ac05382512bc3e53d355e9c8817600d551e3fbb197180b2be3832f4e968befa.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PinnedMessages provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) PinnedMessages()(*ib32e501e2ee14b9c518eef16cdd1655db30b5806e5f4be17957c21a66e36c695.PinnedMessagesRequestBuilder) {
    return ib32e501e2ee14b9c518eef16cdd1655db30b5806e5f4be17957c21a66e36c695.NewPinnedMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PinnedMessagesById provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) PinnedMessagesById(id string)(*ib7358fe28b26b0df886473cc1c746ecee1f749db40a9123335298a7e8ff8c07b.PinnedChatMessageInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["pinnedChatMessageInfo%2Did"] = id
    }
    return ib7358fe28b26b0df886473cc1c746ecee1f749db40a9123335298a7e8ff8c07b.NewPinnedChatMessageInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *ChatItemRequestBuilder) SendActivityNotification()(*i2b860dec50e5faa8f825fb8103a2b27cedc04224865564cb96149aa6bb97d7d5.SendActivityNotificationRequestBuilder) {
    return i2b860dec50e5faa8f825fb8103a2b27cedc04224865564cb96149aa6bb97d7d5.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) Tabs()(*if345d5c48676b774a18852c60a5c6a74c7235e1f5439793ed859b22d9c6e5d11.TabsRequestBuilder) {
    return if345d5c48676b774a18852c60a5c6a74c7235e1f5439793ed859b22d9c6e5d11.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ChatItemRequestBuilder) TabsById(id string)(*i1608545c1b63cdff025064ae1096d8c51ba0eae3a6adeffca8ec30d3f85d8859.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i1608545c1b63cdff025064ae1096d8c51ba0eae3a6adeffca8ec30d3f85d8859.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnhideForUser provides operations to call the unhideForUser method.
func (m *ChatItemRequestBuilder) UnhideForUser()(*i8791640889cb2a0ec1b53c3db460f3f7c82671cef2a627f7fc2649c53e0bb14e.UnhideForUserRequestBuilder) {
    return i8791640889cb2a0ec1b53c3db460f3f7c82671cef2a627f7fc2649c53e0bb14e.NewUnhideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
