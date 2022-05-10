package team

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06d235680d564a3ef57a123529fb839bc78e3d783b194964232f13427d844c71 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/permissiongrants"
    i098a6064a8215b73409a768cbccc320cf24d0e03bb66fd6ff13dea00dc4c414a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule"
    i136f797f85a9f1e6634e5466e71b785c456c2c1eee86c35a2b13d02ce441eab8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/photo"
    i22e0bf3bf0b411437eeed11a4d4964d039826e45fa3b7d2978d58f77ef6d414a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel"
    i2d7c8edc35a718e5c4f64b1a631d454b58f3633bb6b45da7d2b00b5696121cf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/group"
    i3ad9c42d0b35437b41e00867b5229d07541db72e766a62e96824083895df6239 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/tags"
    i41c2fb11f10b987b8a05f187a9c4e4a504845ed53417c45415411a9fd2859fb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/installedapps"
    i5a44e9fa0f41db9f12554a40b2361f881ec68728354eb0eaf00991ba89feb889 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/allchannels"
    i5d1c107e7e1a9d8c559c04c61948425587739f96067ccd77b10ae829994e099f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/template"
    i68f975f75a8c70230e847e815c9956c54465b239a3a098f22d1b7b4db24aaaff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/incomingchannels"
    i7f037032cb5c2440ea470ea00f61b05e2d204ea8d03bb699f6f42d05d4e815cb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/operations"
    ibf5449fd0a50931554dd2426c9124154c1abe82f10512e5e5ccf3b89c4cd2205 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/owners"
    id1d3dcc4007d1efae7e27722e3e927990a64dc9e23112a6e64730ba1806f81f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels"
    ifa790b1527164940bf26fc19476f63bf969f985c52463b401346f01e3a200442 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/members"
    i114956b0565480a9f0e3ece19ff7a6e9b9139109673c060fcfe4320248517137 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/tags/item"
    i3791bcab4fb82a3187acc5ccf447d89dede419fc866690f59c7cbfb00e8856f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/members/item"
    i4c2da1205261a5cdb81be517761c49a129028dcf9b638e023dcba729ed4fee39 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/installedapps/item"
    i54fd8b7b29c3233b404042d76c929d46892d8e6bf3f63b1662cae7d736041bbc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/permissiongrants/item"
    ia614ecc024ed0f1e5aeb6ace4794d20a7b4245518ae60e71f57611979cf500a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/operations/item"
    id53c6629ea6adea0f01bd755c74daf59eb13ca71515cdb627063efdda71b1b00 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/channels/item"
    idc5981e6c9ba563058e9df0cab2d09f2943dcf07f193b2af58e884fac3f25504 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/incomingchannels/item"
    ieb17ba6164c7e120127d94a189d59c168683b600843f0b407607bd6f5513b088 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/owners/item"
    iebd396df2605fb456fb8fd43cdcf23816ccd5ee5aebaa2f1f58cae63eb3cac77 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/allchannels/item"
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
func (m *TeamRequestBuilder) AllChannels()(*i5a44e9fa0f41db9f12554a40b2361f881ec68728354eb0eaf00991ba89feb889.AllChannelsRequestBuilder) {
    return i5a44e9fa0f41db9f12554a40b2361f881ec68728354eb0eaf00991ba89feb889.NewAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.allChannels.item collection
func (m *TeamRequestBuilder) AllChannelsById(id string)(*iebd396df2605fb456fb8fd43cdcf23816ccd5ee5aebaa2f1f58cae63eb3cac77.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return iebd396df2605fb456fb8fd43cdcf23816ccd5ee5aebaa2f1f58cae63eb3cac77.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Channels the channels property
func (m *TeamRequestBuilder) Channels()(*id1d3dcc4007d1efae7e27722e3e927990a64dc9e23112a6e64730ba1806f81f6.ChannelsRequestBuilder) {
    return id1d3dcc4007d1efae7e27722e3e927990a64dc9e23112a6e64730ba1806f81f6.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.channels.item collection
func (m *TeamRequestBuilder) ChannelsById(id string)(*id53c6629ea6adea0f01bd755c74daf59eb13ca71515cdb627063efdda71b1b00.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return id53c6629ea6adea0f01bd755c74daf59eb13ca71515cdb627063efdda71b1b00.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamRequestBuilder) {
    m := &TeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/team{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property team for users
func (m *TeamRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property team for users
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
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property team in users
func (m *TeamRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property team in users
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
// Delete delete navigation property team for users
func (m *TeamRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property team for users
func (m *TeamRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *TeamRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get the team associated with this group.
func (m *TeamRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the team associated with this group.
func (m *TeamRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *TeamRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
func (m *TeamRequestBuilder) Group()(*i2d7c8edc35a718e5c4f64b1a631d454b58f3633bb6b45da7d2b00b5696121cf0.GroupRequestBuilder) {
    return i2d7c8edc35a718e5c4f64b1a631d454b58f3633bb6b45da7d2b00b5696121cf0.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels the incomingChannels property
func (m *TeamRequestBuilder) IncomingChannels()(*i68f975f75a8c70230e847e815c9956c54465b239a3a098f22d1b7b4db24aaaff.IncomingChannelsRequestBuilder) {
    return i68f975f75a8c70230e847e815c9956c54465b239a3a098f22d1b7b4db24aaaff.NewIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.incomingChannels.item collection
func (m *TeamRequestBuilder) IncomingChannelsById(id string)(*idc5981e6c9ba563058e9df0cab2d09f2943dcf07f193b2af58e884fac3f25504.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return idc5981e6c9ba563058e9df0cab2d09f2943dcf07f193b2af58e884fac3f25504.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps the installedApps property
func (m *TeamRequestBuilder) InstalledApps()(*i41c2fb11f10b987b8a05f187a9c4e4a504845ed53417c45415411a9fd2859fb3.InstalledAppsRequestBuilder) {
    return i41c2fb11f10b987b8a05f187a9c4e4a504845ed53417c45415411a9fd2859fb3.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.installedApps.item collection
func (m *TeamRequestBuilder) InstalledAppsById(id string)(*i4c2da1205261a5cdb81be517761c49a129028dcf9b638e023dcba729ed4fee39.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return i4c2da1205261a5cdb81be517761c49a129028dcf9b638e023dcba729ed4fee39.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *TeamRequestBuilder) Members()(*ifa790b1527164940bf26fc19476f63bf969f985c52463b401346f01e3a200442.MembersRequestBuilder) {
    return ifa790b1527164940bf26fc19476f63bf969f985c52463b401346f01e3a200442.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.members.item collection
func (m *TeamRequestBuilder) MembersById(id string)(*i3791bcab4fb82a3187acc5ccf447d89dede419fc866690f59c7cbfb00e8856f2.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i3791bcab4fb82a3187acc5ccf447d89dede419fc866690f59c7cbfb00e8856f2.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *TeamRequestBuilder) Operations()(*i7f037032cb5c2440ea470ea00f61b05e2d204ea8d03bb699f6f42d05d4e815cb.OperationsRequestBuilder) {
    return i7f037032cb5c2440ea470ea00f61b05e2d204ea8d03bb699f6f42d05d4e815cb.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.operations.item collection
func (m *TeamRequestBuilder) OperationsById(id string)(*ia614ecc024ed0f1e5aeb6ace4794d20a7b4245518ae60e71f57611979cf500a7.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return ia614ecc024ed0f1e5aeb6ace4794d20a7b4245518ae60e71f57611979cf500a7.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners the owners property
func (m *TeamRequestBuilder) Owners()(*ibf5449fd0a50931554dd2426c9124154c1abe82f10512e5e5ccf3b89c4cd2205.OwnersRequestBuilder) {
    return ibf5449fd0a50931554dd2426c9124154c1abe82f10512e5e5ccf3b89c4cd2205.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.owners.item collection
func (m *TeamRequestBuilder) OwnersById(id string)(*ieb17ba6164c7e120127d94a189d59c168683b600843f0b407607bd6f5513b088.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did1"] = id
    }
    return ieb17ba6164c7e120127d94a189d59c168683b600843f0b407607bd6f5513b088.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property team in users
func (m *TeamRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property team in users
func (m *TeamRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *TeamRequestBuilder) PermissionGrants()(*i06d235680d564a3ef57a123529fb839bc78e3d783b194964232f13427d844c71.PermissionGrantsRequestBuilder) {
    return i06d235680d564a3ef57a123529fb839bc78e3d783b194964232f13427d844c71.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.permissionGrants.item collection
func (m *TeamRequestBuilder) PermissionGrantsById(id string)(*i54fd8b7b29c3233b404042d76c929d46892d8e6bf3f63b1662cae7d736041bbc.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i54fd8b7b29c3233b404042d76c929d46892d8e6bf3f63b1662cae7d736041bbc.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *TeamRequestBuilder) Photo()(*i136f797f85a9f1e6634e5466e71b785c456c2c1eee86c35a2b13d02ce441eab8.PhotoRequestBuilder) {
    return i136f797f85a9f1e6634e5466e71b785c456c2c1eee86c35a2b13d02ce441eab8.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel the primaryChannel property
func (m *TeamRequestBuilder) PrimaryChannel()(*i22e0bf3bf0b411437eeed11a4d4964d039826e45fa3b7d2978d58f77ef6d414a.PrimaryChannelRequestBuilder) {
    return i22e0bf3bf0b411437eeed11a4d4964d039826e45fa3b7d2978d58f77ef6d414a.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule the schedule property
func (m *TeamRequestBuilder) Schedule()(*i098a6064a8215b73409a768cbccc320cf24d0e03bb66fd6ff13dea00dc4c414a.ScheduleRequestBuilder) {
    return i098a6064a8215b73409a768cbccc320cf24d0e03bb66fd6ff13dea00dc4c414a.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags the tags property
func (m *TeamRequestBuilder) Tags()(*i3ad9c42d0b35437b41e00867b5229d07541db72e766a62e96824083895df6239.TagsRequestBuilder) {
    return i3ad9c42d0b35437b41e00867b5229d07541db72e766a62e96824083895df6239.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.tags.item collection
func (m *TeamRequestBuilder) TagsById(id string)(*i114956b0565480a9f0e3ece19ff7a6e9b9139109673c060fcfe4320248517137.TeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return i114956b0565480a9f0e3ece19ff7a6e9b9139109673c060fcfe4320248517137.NewTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template the template property
func (m *TeamRequestBuilder) Template()(*i5d1c107e7e1a9d8c559c04c61948425587739f96067ccd77b10ae829994e099f.TemplateRequestBuilder) {
    return i5d1c107e7e1a9d8c559c04c61948425587739f96067ccd77b10ae829994e099f.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
