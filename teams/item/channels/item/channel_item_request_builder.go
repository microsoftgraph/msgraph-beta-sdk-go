package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0d3c3bc7ffecb79ef691006b53d99f225238fd156f8da125998e2e94646d7c59 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/provisionemail"
    i283c83b4db244dc7addb2109863ff2121bafc34f89ee6520035ae988c40fb61c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/members"
    i300d5e508c18b0ed5f72f56f6891156b94d9d7bf57dd4cc9236e99e0001c9d5c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/messages"
    i4a0b032148812dd7a6c626731fabeed49b27af755d0f64099a2dad2fe6c463aa "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/tabs"
    i51df0a7d609a404a78b14a8e779ec650531a014c820b75fd5b65e7f1e59fea68 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname"
    i994a1cf7413771c48c58a47b1147a08c8a55fac05a93d165b4fbc06c86315ba1 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/sharedwithteams"
    iaafc2d33c19679581f0ebcc23a4dab0de60433248b7a19e9e790de2ce25dc870 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/filesfolder"
    ibf8343010639ecf1c854bc7d9a70c5124b6d60a75bc3f54375936c1134686c04 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/completemigration"
    ieb3d065db28b358f95647ad7d763fd50c2d96693f776604a02d2f842345cd3be "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/removeemail"
    i07457da9b011f4b8114261d34ff0b729122cabd296529aef5ede38f0d0315e79 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/members/item"
    i1bd9944ba343d9424b1062a0ad9b76c2b12d0a87d3e6d3102a162e6d77945133 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/messages/item"
    i3a7360365f0eaa7cd8e05737099a71b90b500f72f542f94aae165e26cc79e72c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/sharedwithteams/item"
    i668b6c7964e71a13c9358c7e270fbce278a1629d6b040a230ff54a0593a3353a "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item/tabs/item"
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
func (m *ChannelItemRequestBuilder) CompleteMigration()(*ibf8343010639ecf1c854bc7d9a70c5124b6d60a75bc3f54375936c1134686c04.CompleteMigrationRequestBuilder) {
    return ibf8343010639ecf1c854bc7d9a70c5124b6d60a75bc3f54375936c1134686c04.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}/channels/{channel%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property channels for teams
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property channels for teams
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
// CreateGetRequestInformationWithRequestConfiguration the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property channels in teams
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property channels in teams
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
// DeleteWithResponseHandler delete navigation property channels for teams
func (m *ChannelItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property channels for teams
func (m *ChannelItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *ChannelItemRequestBuilder) DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName()(*i51df0a7d609a404a78b14a8e779ec650531a014c820b75fd5b65e7f1e59fea68.DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    return i51df0a7d609a404a78b14a8e779ec650531a014c820b75fd5b65e7f1e59fea68.NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder the filesFolder property
func (m *ChannelItemRequestBuilder) FilesFolder()(*iaafc2d33c19679581f0ebcc23a4dab0de60433248b7a19e9e790de2ce25dc870.FilesFolderRequestBuilder) {
    return iaafc2d33c19679581f0ebcc23a4dab0de60433248b7a19e9e790de2ce25dc870.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *ChannelItemRequestBuilder) Members()(*i283c83b4db244dc7addb2109863ff2121bafc34f89ee6520035ae988c40fb61c.MembersRequestBuilder) {
    return i283c83b4db244dc7addb2109863ff2121bafc34f89ee6520035ae988c40fb61c.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*i07457da9b011f4b8114261d34ff0b729122cabd296529aef5ede38f0d0315e79.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i07457da9b011f4b8114261d34ff0b729122cabd296529aef5ede38f0d0315e79.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *ChannelItemRequestBuilder) Messages()(*i300d5e508c18b0ed5f72f56f6891156b94d9d7bf57dd4cc9236e99e0001c9d5c.MessagesRequestBuilder) {
    return i300d5e508c18b0ed5f72f56f6891156b94d9d7bf57dd4cc9236e99e0001c9d5c.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*i1bd9944ba343d9424b1062a0ad9b76c2b12d0a87d3e6d3102a162e6d77945133.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i1bd9944ba343d9424b1062a0ad9b76c2b12d0a87d3e6d3102a162e6d77945133.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property channels in teams
func (m *ChannelItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property channels in teams
func (m *ChannelItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i0d3c3bc7ffecb79ef691006b53d99f225238fd156f8da125998e2e94646d7c59.ProvisionEmailRequestBuilder) {
    return i0d3c3bc7ffecb79ef691006b53d99f225238fd156f8da125998e2e94646d7c59.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *ChannelItemRequestBuilder) RemoveEmail()(*ieb3d065db28b358f95647ad7d763fd50c2d96693f776604a02d2f842345cd3be.RemoveEmailRequestBuilder) {
    return ieb3d065db28b358f95647ad7d763fd50c2d96693f776604a02d2f842345cd3be.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams the sharedWithTeams property
func (m *ChannelItemRequestBuilder) SharedWithTeams()(*i994a1cf7413771c48c58a47b1147a08c8a55fac05a93d165b4fbc06c86315ba1.SharedWithTeamsRequestBuilder) {
    return i994a1cf7413771c48c58a47b1147a08c8a55fac05a93d165b4fbc06c86315ba1.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item.sharedWithTeams.item collection
func (m *ChannelItemRequestBuilder) SharedWithTeamsById(id string)(*i3a7360365f0eaa7cd8e05737099a71b90b500f72f542f94aae165e26cc79e72c.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return i3a7360365f0eaa7cd8e05737099a71b90b500f72f542f94aae165e26cc79e72c.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs the tabs property
func (m *ChannelItemRequestBuilder) Tabs()(*i4a0b032148812dd7a6c626731fabeed49b27af755d0f64099a2dad2fe6c463aa.TabsRequestBuilder) {
    return i4a0b032148812dd7a6c626731fabeed49b27af755d0f64099a2dad2fe6c463aa.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i668b6c7964e71a13c9358c7e270fbce278a1629d6b040a230ff54a0593a3353a.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i668b6c7964e71a13c9358c7e270fbce278a1629d6b040a230ff54a0593a3353a.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
