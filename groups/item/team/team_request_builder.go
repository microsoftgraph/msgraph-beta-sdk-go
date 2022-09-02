package team

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1bc1304b8111c00df1a1189b84612743ebd4140d3975e0163827f3227dbb3e26 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/owners"
    i40ac59694f013e4b5b1e15575a5f6e0e93c3dab17a07b99aa7bc71b7067d946b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/template"
    i40ad1e813da21e893b26dc980943dedf8dfc449b6a62bf0f23277a7fec206bdf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/clone"
    i4e5568f1c0c9475041cb74f67b12738556910403e2b8b6271342b9609cffb136 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/tags"
    i5fb3ff9b9b79d5dea8cd859e96b4dad4364d2ff120dd48a9b17e17c4655607ac "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/group"
    i72506239d97b0f613e771f51ee2e92a24d89d70f0bbe0f64b4b61ac678b9be53 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/operations"
    i76785ecf032a4612f92def7ae6282a42550c8c1c10ef0f267f85ddc39f32c731 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/sendactivitynotification"
    i76c04af53818aa81ae8d528cbf5ff1061e9881ae0aa40fc1b2a42c6f9c295f58 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/installedapps"
    i7c89f46be075ee198d292af9b1c0f6efd812425fa184c3a9e5ad7438750815f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/photo"
    ia1969c0d9ca1ba48c1d20424d2ddc86eaabf08bdbe12e5b37baef9f6bca9a7ab "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/completemigration"
    ib337fa6eb6ff059ecd97b0b2773310f5a1b513bc20687bd49d35845e3b85761b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel"
    ib5a346ad53a955553527a92fb840c9f85a98d661eacead83d75c66aa0f460b9e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/schedule"
    ib8cfafd2f6ca25503668caed9b0cd67ee91541ffff70e7906c57fa89871b1a20 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/allchannels"
    ibdfa5b0e7aaa94eac928479c155fbb787d756b4bc59311677196676a36ce0197 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/archive"
    ibf8b99dc30819d603bffb578233c10cffceb8078bd72e4af9a92dc5c1985fc6a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/unarchive"
    id961e2ba048867154192d00a449dada9288b6e199b160e9e8eeb3d0645964ea5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/incomingchannels"
    id9e50d3b0a842bb327abf9b574823f3c70464594cfd1aa87bb2f9a8a79f0a550 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels"
    ie0e01504fbdc67467c7eb28dc3a399495eedf258b673445ef02aa1d07d21f356 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/members"
    ifc8ec29a2b557d07fde7a3f16a15bb2eff8cfb5c5b04346762d3a931d4f9553b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants"
    i4ec70ed3a09b5467fad127726e19cca6112801b6feb60254f18e5a2533c3ee56 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/members/item"
    i561e7ee1a103538797b93465871642e260b5e8db2b292a02eee8e68e2a67345a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/installedapps/item"
    i76e835f69addd3b76fa4f4c624ee19dc988f6341fdf4ca5ba96721cbd87ec805 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/allchannels/item"
    i9a8b892ec6341ca828b87b1cde48342764942138cc5b79f3266c017f0fc03ffd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/operations/item"
    ib6cc8679c1d8b1fdbaf622c34a88a722a821e1146db4542c60501545741765e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants/item"
    ic7aad6617468ac8ab7994cbe9e03741fe79693d41944df1e32f15797dcdc1000 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/tags/item"
    id2d1db801a84d6c9a3a750acf0225d5e09a69c9865390c7bab93fa541cbc6125 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/owners/item"
    idf9596a657cf549a24aca9ce70a9474252cf1784a579a04033ef97eb87e63a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item"
    ifc1883ca07fdc0cb053dd4f20f86b918549ab8d93ac7e207d922aa8370ee771f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/incomingchannels/item"
)

// TeamRequestBuilder provides operations to manage the team property of the microsoft.graph.group entity.
type TeamRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamRequestBuilderGetQueryParameters the team associated with this group.
type TeamRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamRequestBuilderGetQueryParameters
}
// TeamRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels the allChannels property
func (m *TeamRequestBuilder) AllChannels()(*ib8cfafd2f6ca25503668caed9b0cd67ee91541ffff70e7906c57fa89871b1a20.AllChannelsRequestBuilder) {
    return ib8cfafd2f6ca25503668caed9b0cd67ee91541ffff70e7906c57fa89871b1a20.NewAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.allChannels.item collection
func (m *TeamRequestBuilder) AllChannelsById(id string)(*i76e835f69addd3b76fa4f4c624ee19dc988f6341fdf4ca5ba96721cbd87ec805.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i76e835f69addd3b76fa4f4c624ee19dc988f6341fdf4ca5ba96721cbd87ec805.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive the archive property
func (m *TeamRequestBuilder) Archive()(*ibdfa5b0e7aaa94eac928479c155fbb787d756b4bc59311677196676a36ce0197.ArchiveRequestBuilder) {
    return ibdfa5b0e7aaa94eac928479c155fbb787d756b4bc59311677196676a36ce0197.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels the channels property
func (m *TeamRequestBuilder) Channels()(*id9e50d3b0a842bb327abf9b574823f3c70464594cfd1aa87bb2f9a8a79f0a550.ChannelsRequestBuilder) {
    return id9e50d3b0a842bb327abf9b574823f3c70464594cfd1aa87bb2f9a8a79f0a550.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.channels.item collection
func (m *TeamRequestBuilder) ChannelsById(id string)(*idf9596a657cf549a24aca9ce70a9474252cf1784a579a04033ef97eb87e63a89.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return idf9596a657cf549a24aca9ce70a9474252cf1784a579a04033ef97eb87e63a89.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone the clone property
func (m *TeamRequestBuilder) Clone()(*i40ad1e813da21e893b26dc980943dedf8dfc449b6a62bf0f23277a7fec206bdf.CloneRequestBuilder) {
    return i40ad1e813da21e893b26dc980943dedf8dfc449b6a62bf0f23277a7fec206bdf.NewCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration the completeMigration property
func (m *TeamRequestBuilder) CompleteMigration()(*ia1969c0d9ca1ba48c1d20424d2ddc86eaabf08bdbe12e5b37baef9f6bca9a7ab.CompleteMigrationRequestBuilder) {
    return ia1969c0d9ca1ba48c1d20424d2ddc86eaabf08bdbe12e5b37baef9f6bca9a7ab.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamRequestBuilder) {
    m := &TeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamRequestBuilder instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property team for groups
func (m *TeamRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property team for groups
func (m *TeamRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *TeamRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the team associated with this group.
func (m *TeamRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the team associated with this group.
func (m *TeamRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TeamRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property team in groups
func (m *TeamRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property team in groups
func (m *TeamRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property team for groups
func (m *TeamRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get the team associated with this group.
func (m *TeamRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable), nil
}
// Group the group property
func (m *TeamRequestBuilder) Group()(*i5fb3ff9b9b79d5dea8cd859e96b4dad4364d2ff120dd48a9b17e17c4655607ac.GroupRequestBuilder) {
    return i5fb3ff9b9b79d5dea8cd859e96b4dad4364d2ff120dd48a9b17e17c4655607ac.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels the incomingChannels property
func (m *TeamRequestBuilder) IncomingChannels()(*id961e2ba048867154192d00a449dada9288b6e199b160e9e8eeb3d0645964ea5.IncomingChannelsRequestBuilder) {
    return id961e2ba048867154192d00a449dada9288b6e199b160e9e8eeb3d0645964ea5.NewIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.incomingChannels.item collection
func (m *TeamRequestBuilder) IncomingChannelsById(id string)(*ifc1883ca07fdc0cb053dd4f20f86b918549ab8d93ac7e207d922aa8370ee771f.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return ifc1883ca07fdc0cb053dd4f20f86b918549ab8d93ac7e207d922aa8370ee771f.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps the installedApps property
func (m *TeamRequestBuilder) InstalledApps()(*i76c04af53818aa81ae8d528cbf5ff1061e9881ae0aa40fc1b2a42c6f9c295f58.InstalledAppsRequestBuilder) {
    return i76c04af53818aa81ae8d528cbf5ff1061e9881ae0aa40fc1b2a42c6f9c295f58.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.installedApps.item collection
func (m *TeamRequestBuilder) InstalledAppsById(id string)(*i561e7ee1a103538797b93465871642e260b5e8db2b292a02eee8e68e2a67345a.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return i561e7ee1a103538797b93465871642e260b5e8db2b292a02eee8e68e2a67345a.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *TeamRequestBuilder) Members()(*ie0e01504fbdc67467c7eb28dc3a399495eedf258b673445ef02aa1d07d21f356.MembersRequestBuilder) {
    return ie0e01504fbdc67467c7eb28dc3a399495eedf258b673445ef02aa1d07d21f356.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.members.item collection
func (m *TeamRequestBuilder) MembersById(id string)(*i4ec70ed3a09b5467fad127726e19cca6112801b6feb60254f18e5a2533c3ee56.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i4ec70ed3a09b5467fad127726e19cca6112801b6feb60254f18e5a2533c3ee56.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *TeamRequestBuilder) Operations()(*i72506239d97b0f613e771f51ee2e92a24d89d70f0bbe0f64b4b61ac678b9be53.OperationsRequestBuilder) {
    return i72506239d97b0f613e771f51ee2e92a24d89d70f0bbe0f64b4b61ac678b9be53.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.operations.item collection
func (m *TeamRequestBuilder) OperationsById(id string)(*i9a8b892ec6341ca828b87b1cde48342764942138cc5b79f3266c017f0fc03ffd.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return i9a8b892ec6341ca828b87b1cde48342764942138cc5b79f3266c017f0fc03ffd.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners the owners property
func (m *TeamRequestBuilder) Owners()(*i1bc1304b8111c00df1a1189b84612743ebd4140d3975e0163827f3227dbb3e26.OwnersRequestBuilder) {
    return i1bc1304b8111c00df1a1189b84612743ebd4140d3975e0163827f3227dbb3e26.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.owners.item collection
func (m *TeamRequestBuilder) OwnersById(id string)(*id2d1db801a84d6c9a3a750acf0225d5e09a69c9865390c7bab93fa541cbc6125.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return id2d1db801a84d6c9a3a750acf0225d5e09a69c9865390c7bab93fa541cbc6125.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property team in groups
func (m *TeamRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// PermissionGrants the permissionGrants property
func (m *TeamRequestBuilder) PermissionGrants()(*ifc8ec29a2b557d07fde7a3f16a15bb2eff8cfb5c5b04346762d3a931d4f9553b.PermissionGrantsRequestBuilder) {
    return ifc8ec29a2b557d07fde7a3f16a15bb2eff8cfb5c5b04346762d3a931d4f9553b.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.permissionGrants.item collection
func (m *TeamRequestBuilder) PermissionGrantsById(id string)(*ib6cc8679c1d8b1fdbaf622c34a88a722a821e1146db4542c60501545741765e0.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return ib6cc8679c1d8b1fdbaf622c34a88a722a821e1146db4542c60501545741765e0.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *TeamRequestBuilder) Photo()(*i7c89f46be075ee198d292af9b1c0f6efd812425fa184c3a9e5ad7438750815f0.PhotoRequestBuilder) {
    return i7c89f46be075ee198d292af9b1c0f6efd812425fa184c3a9e5ad7438750815f0.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel the primaryChannel property
func (m *TeamRequestBuilder) PrimaryChannel()(*ib337fa6eb6ff059ecd97b0b2773310f5a1b513bc20687bd49d35845e3b85761b.PrimaryChannelRequestBuilder) {
    return ib337fa6eb6ff059ecd97b0b2773310f5a1b513bc20687bd49d35845e3b85761b.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule the schedule property
func (m *TeamRequestBuilder) Schedule()(*ib5a346ad53a955553527a92fb840c9f85a98d661eacead83d75c66aa0f460b9e.ScheduleRequestBuilder) {
    return ib5a346ad53a955553527a92fb840c9f85a98d661eacead83d75c66aa0f460b9e.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification the sendActivityNotification property
func (m *TeamRequestBuilder) SendActivityNotification()(*i76785ecf032a4612f92def7ae6282a42550c8c1c10ef0f267f85ddc39f32c731.SendActivityNotificationRequestBuilder) {
    return i76785ecf032a4612f92def7ae6282a42550c8c1c10ef0f267f85ddc39f32c731.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags the tags property
func (m *TeamRequestBuilder) Tags()(*i4e5568f1c0c9475041cb74f67b12738556910403e2b8b6271342b9609cffb136.TagsRequestBuilder) {
    return i4e5568f1c0c9475041cb74f67b12738556910403e2b8b6271342b9609cffb136.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.tags.item collection
func (m *TeamRequestBuilder) TagsById(id string)(*ic7aad6617468ac8ab7994cbe9e03741fe79693d41944df1e32f15797dcdc1000.TeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return ic7aad6617468ac8ab7994cbe9e03741fe79693d41944df1e32f15797dcdc1000.NewTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template the template property
func (m *TeamRequestBuilder) Template()(*i40ac59694f013e4b5b1e15575a5f6e0e93c3dab17a07b99aa7bc71b7067d946b.TemplateRequestBuilder) {
    return i40ac59694f013e4b5b1e15575a5f6e0e93c3dab17a07b99aa7bc71b7067d946b.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive the unarchive property
func (m *TeamRequestBuilder) Unarchive()(*ibf8b99dc30819d603bffb578233c10cffceb8078bd72e4af9a92dc5c1985fc6a.UnarchiveRequestBuilder) {
    return ibf8b99dc30819d603bffb578233c10cffceb8078bd72e4af9a92dc5c1985fc6a.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
