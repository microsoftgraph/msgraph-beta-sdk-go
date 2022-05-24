package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/operations"
    i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/permissiongrants"
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
func (m *ChatItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property chats for me
func (m *ChatItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ChatItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChatItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get chats from me
func (m *ChatItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ChatItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
func (m *ChatItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property chats in me
func (m *ChatItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, requestConfiguration *ChatItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property chats for me
func (m *ChatItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property chats for me
func (m *ChatItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ChatItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get chats from me
func (m *ChatItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get chats from me
func (m *ChatItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ChatItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable), nil
}
// InstalledApps the installedApps property
func (m *ChatItemRequestBuilder) InstalledApps()(*ia4a5a30ed707af790b1c2b15936fcf68a4ac1eaee6407d4e5837f9dee33d2717.InstalledAppsRequestBuilder) {
    return ia4a5a30ed707af790b1c2b15936fcf68a4ac1eaee6407d4e5837f9dee33d2717.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.installedApps.item collection
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
// LastMessagePreview the lastMessagePreview property
func (m *ChatItemRequestBuilder) LastMessagePreview()(*idead0084f5300487c412f0afd4cb69c96dab7aa445fdb45feba337598ba6e4a4.LastMessagePreviewRequestBuilder) {
    return idead0084f5300487c412f0afd4cb69c96dab7aa445fdb45feba337598ba6e4a4.NewLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members the members property
func (m *ChatItemRequestBuilder) Members()(*ibcd033263282f46cf08c9ffed250e454bfe93a5f0d41ede7466f3725b012808c.MembersRequestBuilder) {
    return ibcd033263282f46cf08c9ffed250e454bfe93a5f0d41ede7466f3725b012808c.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.members.item collection
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
// Messages the messages property
func (m *ChatItemRequestBuilder) Messages()(*if873fe971f2b52c7a520477d31db735634ccbf62cf832a0dacc1f30a21e89298.MessagesRequestBuilder) {
    return if873fe971f2b52c7a520477d31db735634ccbf62cf832a0dacc1f30a21e89298.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.messages.item collection
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
// Operations the operations property
func (m *ChatItemRequestBuilder) Operations()(*i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35.OperationsRequestBuilder) {
    return i1bcde115488e1362c679009fc0d5cf62474b9a7ba3d051b9bd8214ab3f0c2c35.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.operations.item collection
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
func (m *ChatItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property chats in me
func (m *ChatItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Chatable, requestConfiguration *ChatItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PermissionGrants the permissionGrants property
func (m *ChatItemRequestBuilder) PermissionGrants()(*i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f.PermissionGrantsRequestBuilder) {
    return i38e5134ffce5124e21fd0f668bb73e5568b24566a9e956b4d12afef11bee220f.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.permissionGrants.item collection
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
// PinnedMessages the pinnedMessages property
func (m *ChatItemRequestBuilder) PinnedMessages()(*ib32e501e2ee14b9c518eef16cdd1655db30b5806e5f4be17957c21a66e36c695.PinnedMessagesRequestBuilder) {
    return ib32e501e2ee14b9c518eef16cdd1655db30b5806e5f4be17957c21a66e36c695.NewPinnedMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PinnedMessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.pinnedMessages.item collection
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
// Tabs the tabs property
func (m *ChatItemRequestBuilder) Tabs()(*if345d5c48676b774a18852c60a5c6a74c7235e1f5439793ed859b22d9c6e5d11.TabsRequestBuilder) {
    return if345d5c48676b774a18852c60a5c6a74c7235e1f5439793ed859b22d9c6e5d11.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.tabs.item collection
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
