package primarychannel

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i02fac5bded055b1adca5d6d29c12cb7e57d7bdd49adadfada05e5495fa253fbb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/tabs"
    i123b714ca88ab0949ff156392eee87cf864e73abfdf3932739bdbb7bf6dee147 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/provisionemail"
    i2cb64c7feea67ee6582a345caca1bcdfb2a1d38e2bbdb2d3cea443d51dc7e004 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/messages"
    i4d798133101926112680212ca2c43d2eecac62637d765f5d673c73d468076587 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/filesfolder"
    i5608b27f3fe17f702e3afca624e0a52c6cd4c244ef3a0fa8b70f94b12c09736f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/sharedwithteams"
    i630518fea862c1cbe6fa515187ec5f07495ead144ae1e606f379ffd0fa9c4d6c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname"
    i8b912bce189a9a2123cdb51c5e064b07bf7b402fb3f29b2142013b5bf62ed0a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/removeemail"
    icd45740e504bea46da158b603975e20767322cfcc44ed4be2b7cf35bf48c4c39 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/completemigration"
    id06349f37f8224c543de1bf4a94b2c580a16638c6a6485955fc06191bbb1458e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/members"
    i061a95c7dceeca0edf96fa6065225afa5fed9223bc97d451bb900337c61f56f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/messages/item"
    ia0799eb029da592f6cd5557bbb1f1c880a4e6f339acd44057ad51cd71509ab3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/sharedwithteams/item"
    ia911f8872a4d20a1fed82afc3006820ddad1fc77d06eb0842072e244b22816a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/members/item"
    ief96b8d36ab7ce80404fcb52a1f5b70c665948caadd60c3bb686006b33cbe11c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel/tabs/item"
)

// PrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type PrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrimaryChannelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrimaryChannelRequestBuilderGetQueryParameters the general channel for the team.
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrimaryChannelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrimaryChannelRequestBuilderGetQueryParameters
}
// PrimaryChannelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration the completeMigration property
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*icd45740e504bea46da158b603975e20767322cfcc44ed4be2b7cf35bf48c4c39.CompleteMigrationRequestBuilder) {
    return icd45740e504bea46da158b603975e20767322cfcc44ed4be2b7cf35bf48c4c39.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/team/primaryChannel{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property primaryChannel for me
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property primaryChannel for me
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *PrimaryChannelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PrimaryChannelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property primaryChannel in me
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property primaryChannel in me
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *PrimaryChannelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property primaryChannel for me
func (m *PrimaryChannelRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property primaryChannel for me
func (m *PrimaryChannelRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *PrimaryChannelRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *PrimaryChannelRequestBuilder) DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName()(*i630518fea862c1cbe6fa515187ec5f07495ead144ae1e606f379ffd0fa9c4d6c.DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    return i630518fea862c1cbe6fa515187ec5f07495ead144ae1e606f379ffd0fa9c4d6c.NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder the filesFolder property
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*i4d798133101926112680212ca2c43d2eecac62637d765f5d673c73d468076587.FilesFolderRequestBuilder) {
    return i4d798133101926112680212ca2c43d2eecac62637d765f5d673c73d468076587.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the general channel for the team.
func (m *PrimaryChannelRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the general channel for the team.
func (m *PrimaryChannelRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *PrimaryChannelRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *PrimaryChannelRequestBuilder) Members()(*id06349f37f8224c543de1bf4a94b2c580a16638c6a6485955fc06191bbb1458e.MembersRequestBuilder) {
    return id06349f37f8224c543de1bf4a94b2c580a16638c6a6485955fc06191bbb1458e.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.primaryChannel.members.item collection
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*ia911f8872a4d20a1fed82afc3006820ddad1fc77d06eb0842072e244b22816a2.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return ia911f8872a4d20a1fed82afc3006820ddad1fc77d06eb0842072e244b22816a2.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *PrimaryChannelRequestBuilder) Messages()(*i2cb64c7feea67ee6582a345caca1bcdfb2a1d38e2bbdb2d3cea443d51dc7e004.MessagesRequestBuilder) {
    return i2cb64c7feea67ee6582a345caca1bcdfb2a1d38e2bbdb2d3cea443d51dc7e004.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.primaryChannel.messages.item collection
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*i061a95c7dceeca0edf96fa6065225afa5fed9223bc97d451bb900337c61f56f1.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i061a95c7dceeca0edf96fa6065225afa5fed9223bc97d451bb900337c61f56f1.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property primaryChannel in me
func (m *PrimaryChannelRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property primaryChannel in me
func (m *PrimaryChannelRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *PrimaryChannelRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *PrimaryChannelRequestBuilder) ProvisionEmail()(*i123b714ca88ab0949ff156392eee87cf864e73abfdf3932739bdbb7bf6dee147.ProvisionEmailRequestBuilder) {
    return i123b714ca88ab0949ff156392eee87cf864e73abfdf3932739bdbb7bf6dee147.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *PrimaryChannelRequestBuilder) RemoveEmail()(*i8b912bce189a9a2123cdb51c5e064b07bf7b402fb3f29b2142013b5bf62ed0a1.RemoveEmailRequestBuilder) {
    return i8b912bce189a9a2123cdb51c5e064b07bf7b402fb3f29b2142013b5bf62ed0a1.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams the sharedWithTeams property
func (m *PrimaryChannelRequestBuilder) SharedWithTeams()(*i5608b27f3fe17f702e3afca624e0a52c6cd4c244ef3a0fa8b70f94b12c09736f.SharedWithTeamsRequestBuilder) {
    return i5608b27f3fe17f702e3afca624e0a52c6cd4c244ef3a0fa8b70f94b12c09736f.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.primaryChannel.sharedWithTeams.item collection
func (m *PrimaryChannelRequestBuilder) SharedWithTeamsById(id string)(*ia0799eb029da592f6cd5557bbb1f1c880a4e6f339acd44057ad51cd71509ab3a.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return ia0799eb029da592f6cd5557bbb1f1c880a4e6f339acd44057ad51cd71509ab3a.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs the tabs property
func (m *PrimaryChannelRequestBuilder) Tabs()(*i02fac5bded055b1adca5d6d29c12cb7e57d7bdd49adadfada05e5495fa253fbb.TabsRequestBuilder) {
    return i02fac5bded055b1adca5d6d29c12cb7e57d7bdd49adadfada05e5495fa253fbb.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.primaryChannel.tabs.item collection
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*ief96b8d36ab7ce80404fcb52a1f5b70c665948caadd60c3bb686006b33cbe11c.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return ief96b8d36ab7ce80404fcb52a1f5b70c665948caadd60c3bb686006b33cbe11c.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
