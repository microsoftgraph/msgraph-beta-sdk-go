package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i581c9040ffa9a41ebb258b7aa9925e3250fe9bb2ce1c3b5c208cfad46525f583 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/sharedwithteams"
    i5950982c49f9b790281e46d917fbc8e2588f887fe004a0880a390604342b9206 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/provisionemail"
    i8757b8026631612ee966156e57d50632cfeac045e534d2126c5163d3a0097bd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/completemigration"
    i92c53b8db605e444c70341770860a5f7d918e589496eb8be91e4920605418725 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname"
    ia1bb0c6a2a0ceaca1655c73634274e170040c2fbcacfb3d1740acc3b766086f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/messages"
    iaf82d6a457d48781c7276b4a97c2a339ebe61e07d40e46b0a8d559917ef53439 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/members"
    ie16264da7ef7dc0b3b387d82b0e496ca42245523fd2e14773f1539a842e5727d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/tabs"
    ief74bb2e9cfd0badb866de8f5f2abc4d8fa0fef866bed8f879fc3669419038e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/filesfolder"
    ifcf2608e32d947c5d3be221d105c318efd200cf378a19dccc1b5f20077a0900f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/removeemail"
    i2e9b24c97e71508a3792941cf07935eb547d1deba8e11c10e5ae1ed6ba62ce06 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/members/item"
    i7a6faf26cbcc42a91f8fc5be95ea8a9e891cf24e5519db5caf3f973431389cd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/sharedwithteams/item"
    i87855f00a209578c65747c399533d921f0254c2b7e7baba787719b2e005fa5c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/tabs/item"
    i9d54d2d9962a5dd3a7c0c61b69a9f264c4d4abc90cf29305c42a197876f3def4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item/messages/item"
)

// ChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type ChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelItemRequestBuilderGetQueryParameters
}
// ChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration the completeMigration property
func (m *ChannelItemRequestBuilder) CompleteMigration()(*i8757b8026631612ee966156e57d50632cfeac045e534d2126c5163d3a0097bd0.CompleteMigrationRequestBuilder) {
    return i8757b8026631612ee966156e57d50632cfeac045e534d2126c5163d3a0097bd0.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/team/channels/{channel%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property channels for users
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property channels for users
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property channels in users
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property channels in users
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property channels for users
func (m *ChannelItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property channels for users
func (m *ChannelItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *ChannelItemRequestBuilder) DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName()(*i92c53b8db605e444c70341770860a5f7d918e589496eb8be91e4920605418725.DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    return i92c53b8db605e444c70341770860a5f7d918e589496eb8be91e4920605418725.NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder the filesFolder property
func (m *ChannelItemRequestBuilder) FilesFolder()(*ief74bb2e9cfd0badb866de8f5f2abc4d8fa0fef866bed8f879fc3669419038e8.FilesFolderRequestBuilder) {
    return ief74bb2e9cfd0badb866de8f5f2abc4d8fa0fef866bed8f879fc3669419038e8.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// Members the members property
func (m *ChannelItemRequestBuilder) Members()(*iaf82d6a457d48781c7276b4a97c2a339ebe61e07d40e46b0a8d559917ef53439.MembersRequestBuilder) {
    return iaf82d6a457d48781c7276b4a97c2a339ebe61e07d40e46b0a8d559917ef53439.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*i2e9b24c97e71508a3792941cf07935eb547d1deba8e11c10e5ae1ed6ba62ce06.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i2e9b24c97e71508a3792941cf07935eb547d1deba8e11c10e5ae1ed6ba62ce06.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *ChannelItemRequestBuilder) Messages()(*ia1bb0c6a2a0ceaca1655c73634274e170040c2fbcacfb3d1740acc3b766086f1.MessagesRequestBuilder) {
    return ia1bb0c6a2a0ceaca1655c73634274e170040c2fbcacfb3d1740acc3b766086f1.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*i9d54d2d9962a5dd3a7c0c61b69a9f264c4d4abc90cf29305c42a197876f3def4.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i9d54d2d9962a5dd3a7c0c61b69a9f264c4d4abc90cf29305c42a197876f3def4.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in users
func (m *ChannelItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property channels in users
func (m *ChannelItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// ProvisionEmail the provisionEmail property
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i5950982c49f9b790281e46d917fbc8e2588f887fe004a0880a390604342b9206.ProvisionEmailRequestBuilder) {
    return i5950982c49f9b790281e46d917fbc8e2588f887fe004a0880a390604342b9206.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *ChannelItemRequestBuilder) RemoveEmail()(*ifcf2608e32d947c5d3be221d105c318efd200cf378a19dccc1b5f20077a0900f.RemoveEmailRequestBuilder) {
    return ifcf2608e32d947c5d3be221d105c318efd200cf378a19dccc1b5f20077a0900f.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams the sharedWithTeams property
func (m *ChannelItemRequestBuilder) SharedWithTeams()(*i581c9040ffa9a41ebb258b7aa9925e3250fe9bb2ce1c3b5c208cfad46525f583.SharedWithTeamsRequestBuilder) {
    return i581c9040ffa9a41ebb258b7aa9925e3250fe9bb2ce1c3b5c208cfad46525f583.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.channels.item.sharedWithTeams.item collection
func (m *ChannelItemRequestBuilder) SharedWithTeamsById(id string)(*i7a6faf26cbcc42a91f8fc5be95ea8a9e891cf24e5519db5caf3f973431389cd7.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return i7a6faf26cbcc42a91f8fc5be95ea8a9e891cf24e5519db5caf3f973431389cd7.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs the tabs property
func (m *ChannelItemRequestBuilder) Tabs()(*ie16264da7ef7dc0b3b387d82b0e496ca42245523fd2e14773f1539a842e5727d.TabsRequestBuilder) {
    return ie16264da7ef7dc0b3b387d82b0e496ca42245523fd2e14773f1539a842e5727d.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i87855f00a209578c65747c399533d921f0254c2b7e7baba787719b2e005fa5c5.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i87855f00a209578c65747c399533d921f0254c2b7e7baba787719b2e005fa5c5.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
