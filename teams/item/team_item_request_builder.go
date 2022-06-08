package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i048dd3eda626b5c64ad8acc7d562355d6648e78b61cb9150e852218748576fca "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/installedapps"
    i0acc5a25a7c609872781f3db26798489e7383966bcf6ccb86a382799c15443b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/operations"
    i279367125768e25f45766783a5a7e5b8365fc99cf8e0872f52ac607992a020d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/clone"
    i390feffd708cca0f7b6a5195858181d312c7117b9de188d29741208793e80c6c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/completemigration"
    i39e1c08574131f72876adb29d6160e8debd9e80a5d56f6e5667bb4466e655518 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/archive"
    i3c28f7755196c773e7492789789df058eacb0aa39df51194978fa414b36be92c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/tags"
    i3e307cff6710e4b62c7cd2d770514afca50fee73703ce6144ba602807a55f52d "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels"
    i4560e2d0c3ef25c4962765280004a00c65a0b596f685f368c272fff7174d3a9e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule"
    i565462b4460b597336ced0fd92db4b81564041b705b332120e10962a0807ca79 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/owners"
    i6356c4272d892a6b1753bbed0e3b3f265f9d5630d86393549775b40b26fad406 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/group"
    i7c7aa71fa334e1234ebb2953457bb8a87886a5cc192ada64a17391c65db699d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/template"
    ia3d84312b941dc5c3aaa9a4a37bc36866b779989abc4a4a4acd33b1e65368bcd "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/members"
    ia42fa4d97821d26378d9f7bee3d9af6325ce61b826e31042ce499fd11eed580f "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/allchannels"
    ib3d27e8197f598d7bf72f02ae854d33875306efb682c4984f7c88f1856aadd6d "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/unarchive"
    ib426e286e70724fac9bfa5fc5f23505db892b09653a2e444dddddd561a3ded1e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/photo"
    ibd507eb28e5ea81639c3f3209c96a828f71ad54ba353088076f57e6cb7e164f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/incomingchannels"
    icf541587ec6c05dc98eb683d6e0738a0265ce6e56b0857fe95939ba5cf9e730f "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel"
    iea95a35c2aecd362bfafe8dae609401aa9fd68c91737848de33cd6ecad21f51b "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/sendactivitynotification"
    iecf1ce1fbbc38cbe05b8bcd58d24b218680aef89fbf95142ed150a70e3ef887a "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/permissiongrants"
    i15093721720dbb1ad7a9823d9c98bb0b1b6afb6839ccc9b99d7f9f369fea596c "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/members/item"
    i1abe3d3d937fd27f8c8c35d38ad858733332abc3e4fbd568ec7960fd205c0b87 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/allchannels/item"
    i75af39801867d39724f8d806372cf754dc48d3407d87f336daf7647dbed1347f "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/installedapps/item"
    i7b61bd390262c74ea942911c717d521b8d8938c0bc5a556ea5c1aa8faf593708 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/tags/item"
    i816caa6537457ad2108f6d14fc5f83f30b010c0f3ebaa1f99a81c44641b9d1e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/channels/item"
    ib1fa9341ce0e195b37a0fcac629719d660cd28ed736c3abf54fe54102868c36d "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/operations/item"
    iba418f5dd9a4de62eac90abccb0188ef53c41d0492c29cceac1855283b8a7c57 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/incomingchannels/item"
    iba461bd7c348411b42a81c1813125516bec205337ca79a73172439185e9cd1c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/permissiongrants/item"
    ifc8558814710963aeed18cad6c4a77609daafa95e19508c1be52d9732d135fc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/owners/item"
)

// TeamItemRequestBuilder provides operations to manage the collection of team entities.
type TeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamItemRequestBuilderGetQueryParameters retrieve the properties and relationships of the specified team.
type TeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamItemRequestBuilderGetQueryParameters
}
// TeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels the allChannels property
func (m *TeamItemRequestBuilder) AllChannels()(*ia42fa4d97821d26378d9f7bee3d9af6325ce61b826e31042ce499fd11eed580f.AllChannelsRequestBuilder) {
    return ia42fa4d97821d26378d9f7bee3d9af6325ce61b826e31042ce499fd11eed580f.NewAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.allChannels.item collection
func (m *TeamItemRequestBuilder) AllChannelsById(id string)(*i1abe3d3d937fd27f8c8c35d38ad858733332abc3e4fbd568ec7960fd205c0b87.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i1abe3d3d937fd27f8c8c35d38ad858733332abc3e4fbd568ec7960fd205c0b87.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive the archive property
func (m *TeamItemRequestBuilder) Archive()(*i39e1c08574131f72876adb29d6160e8debd9e80a5d56f6e5667bb4466e655518.ArchiveRequestBuilder) {
    return i39e1c08574131f72876adb29d6160e8debd9e80a5d56f6e5667bb4466e655518.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels the channels property
func (m *TeamItemRequestBuilder) Channels()(*i3e307cff6710e4b62c7cd2d770514afca50fee73703ce6144ba602807a55f52d.ChannelsRequestBuilder) {
    return i3e307cff6710e4b62c7cd2d770514afca50fee73703ce6144ba602807a55f52d.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.channels.item collection
func (m *TeamItemRequestBuilder) ChannelsById(id string)(*i816caa6537457ad2108f6d14fc5f83f30b010c0f3ebaa1f99a81c44641b9d1e8.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i816caa6537457ad2108f6d14fc5f83f30b010c0f3ebaa1f99a81c44641b9d1e8.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone the clone property
func (m *TeamItemRequestBuilder) Clone()(*i279367125768e25f45766783a5a7e5b8365fc99cf8e0872f52ac607992a020d7.CloneRequestBuilder) {
    return i279367125768e25f45766783a5a7e5b8365fc99cf8e0872f52ac607992a020d7.NewCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration the completeMigration property
func (m *TeamItemRequestBuilder) CompleteMigration()(*i390feffd708cca0f7b6a5195858181d312c7117b9de188d29741208793e80c6c.CompleteMigrationRequestBuilder) {
    return i390feffd708cca0f7b6a5195858181d312c7117b9de188d29741208793e80c6c.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamItemRequestBuilder) {
    m := &TeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from teams
func (m *TeamItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete entity from teams
func (m *TeamItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the properties and relationships of the specified team.
func (m *TeamItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve the properties and relationships of the specified team.
func (m *TeamItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of the specified team.
func (m *TeamItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the properties of the specified team.
func (m *TeamItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from teams
func (m *TeamItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete entity from teams
func (m *TeamItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get retrieve the properties and relationships of the specified team.
func (m *TeamItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler retrieve the properties and relationships of the specified team.
func (m *TeamItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable), nil
}
// Group the group property
func (m *TeamItemRequestBuilder) Group()(*i6356c4272d892a6b1753bbed0e3b3f265f9d5630d86393549775b40b26fad406.GroupRequestBuilder) {
    return i6356c4272d892a6b1753bbed0e3b3f265f9d5630d86393549775b40b26fad406.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels the incomingChannels property
func (m *TeamItemRequestBuilder) IncomingChannels()(*ibd507eb28e5ea81639c3f3209c96a828f71ad54ba353088076f57e6cb7e164f7.IncomingChannelsRequestBuilder) {
    return ibd507eb28e5ea81639c3f3209c96a828f71ad54ba353088076f57e6cb7e164f7.NewIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.incomingChannels.item collection
func (m *TeamItemRequestBuilder) IncomingChannelsById(id string)(*iba418f5dd9a4de62eac90abccb0188ef53c41d0492c29cceac1855283b8a7c57.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return iba418f5dd9a4de62eac90abccb0188ef53c41d0492c29cceac1855283b8a7c57.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps the installedApps property
func (m *TeamItemRequestBuilder) InstalledApps()(*i048dd3eda626b5c64ad8acc7d562355d6648e78b61cb9150e852218748576fca.InstalledAppsRequestBuilder) {
    return i048dd3eda626b5c64ad8acc7d562355d6648e78b61cb9150e852218748576fca.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.installedApps.item collection
func (m *TeamItemRequestBuilder) InstalledAppsById(id string)(*i75af39801867d39724f8d806372cf754dc48d3407d87f336daf7647dbed1347f.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return i75af39801867d39724f8d806372cf754dc48d3407d87f336daf7647dbed1347f.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *TeamItemRequestBuilder) Members()(*ia3d84312b941dc5c3aaa9a4a37bc36866b779989abc4a4a4acd33b1e65368bcd.MembersRequestBuilder) {
    return ia3d84312b941dc5c3aaa9a4a37bc36866b779989abc4a4a4acd33b1e65368bcd.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.members.item collection
func (m *TeamItemRequestBuilder) MembersById(id string)(*i15093721720dbb1ad7a9823d9c98bb0b1b6afb6839ccc9b99d7f9f369fea596c.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i15093721720dbb1ad7a9823d9c98bb0b1b6afb6839ccc9b99d7f9f369fea596c.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *TeamItemRequestBuilder) Operations()(*i0acc5a25a7c609872781f3db26798489e7383966bcf6ccb86a382799c15443b7.OperationsRequestBuilder) {
    return i0acc5a25a7c609872781f3db26798489e7383966bcf6ccb86a382799c15443b7.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.operations.item collection
func (m *TeamItemRequestBuilder) OperationsById(id string)(*ib1fa9341ce0e195b37a0fcac629719d660cd28ed736c3abf54fe54102868c36d.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return ib1fa9341ce0e195b37a0fcac629719d660cd28ed736c3abf54fe54102868c36d.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners the owners property
func (m *TeamItemRequestBuilder) Owners()(*i565462b4460b597336ced0fd92db4b81564041b705b332120e10962a0807ca79.OwnersRequestBuilder) {
    return i565462b4460b597336ced0fd92db4b81564041b705b332120e10962a0807ca79.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.owners.item collection
func (m *TeamItemRequestBuilder) OwnersById(id string)(*ifc8558814710963aeed18cad6c4a77609daafa95e19508c1be52d9732d135fc7.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return ifc8558814710963aeed18cad6c4a77609daafa95e19508c1be52d9732d135fc7.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of the specified team.
func (m *TeamItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the properties of the specified team.
func (m *TeamItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *TeamItemRequestBuilder) PermissionGrants()(*iecf1ce1fbbc38cbe05b8bcd58d24b218680aef89fbf95142ed150a70e3ef887a.PermissionGrantsRequestBuilder) {
    return iecf1ce1fbbc38cbe05b8bcd58d24b218680aef89fbf95142ed150a70e3ef887a.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.permissionGrants.item collection
func (m *TeamItemRequestBuilder) PermissionGrantsById(id string)(*iba461bd7c348411b42a81c1813125516bec205337ca79a73172439185e9cd1c4.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return iba461bd7c348411b42a81c1813125516bec205337ca79a73172439185e9cd1c4.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *TeamItemRequestBuilder) Photo()(*ib426e286e70724fac9bfa5fc5f23505db892b09653a2e444dddddd561a3ded1e.PhotoRequestBuilder) {
    return ib426e286e70724fac9bfa5fc5f23505db892b09653a2e444dddddd561a3ded1e.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel the primaryChannel property
func (m *TeamItemRequestBuilder) PrimaryChannel()(*icf541587ec6c05dc98eb683d6e0738a0265ce6e56b0857fe95939ba5cf9e730f.PrimaryChannelRequestBuilder) {
    return icf541587ec6c05dc98eb683d6e0738a0265ce6e56b0857fe95939ba5cf9e730f.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule the schedule property
func (m *TeamItemRequestBuilder) Schedule()(*i4560e2d0c3ef25c4962765280004a00c65a0b596f685f368c272fff7174d3a9e.ScheduleRequestBuilder) {
    return i4560e2d0c3ef25c4962765280004a00c65a0b596f685f368c272fff7174d3a9e.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification the sendActivityNotification property
func (m *TeamItemRequestBuilder) SendActivityNotification()(*iea95a35c2aecd362bfafe8dae609401aa9fd68c91737848de33cd6ecad21f51b.SendActivityNotificationRequestBuilder) {
    return iea95a35c2aecd362bfafe8dae609401aa9fd68c91737848de33cd6ecad21f51b.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags the tags property
func (m *TeamItemRequestBuilder) Tags()(*i3c28f7755196c773e7492789789df058eacb0aa39df51194978fa414b36be92c.TagsRequestBuilder) {
    return i3c28f7755196c773e7492789789df058eacb0aa39df51194978fa414b36be92c.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.tags.item collection
func (m *TeamItemRequestBuilder) TagsById(id string)(*i7b61bd390262c74ea942911c717d521b8d8938c0bc5a556ea5c1aa8faf593708.TeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return i7b61bd390262c74ea942911c717d521b8d8938c0bc5a556ea5c1aa8faf593708.NewTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template the template property
func (m *TeamItemRequestBuilder) Template()(*i7c7aa71fa334e1234ebb2953457bb8a87886a5cc192ada64a17391c65db699d5.TemplateRequestBuilder) {
    return i7c7aa71fa334e1234ebb2953457bb8a87886a5cc192ada64a17391c65db699d5.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive the unarchive property
func (m *TeamItemRequestBuilder) Unarchive()(*ib3d27e8197f598d7bf72f02ae854d33875306efb682c4984f7c88f1856aadd6d.UnarchiveRequestBuilder) {
    return ib3d27e8197f598d7bf72f02ae854d33875306efb682c4984f7c88f1856aadd6d.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
