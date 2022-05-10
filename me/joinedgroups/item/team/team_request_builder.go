package team

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0bd22d14acd5879034c163b078c3b1b8eb2a9adc63b0e6f5012d9195f6aca876 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/installedapps"
    i0e48002ffeba3712f35f2a803886b2f7f6f22fdb49dcf5e23226a97d472b7b1a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/permissiongrants"
    i0ef26a7a288133f0d75cf96466beded00255791dbbf97db173e8a60516139727 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/template"
    i240464d1510ca1667c1822c4c48b1244eefd38846201acf25582ab4735096048 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule"
    i2a35925faef47df44ed32466b7e2a6f0223e3e8b104ee21733303d304ba89738 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/allchannels"
    i4c16420d301ac519bee3f68145c8514ea14a14d9a18a89a970cf5b1136c199bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/channels"
    i631755ac6a9d8774786b5b39fbe602f04e0c8147c9fbdc914d0d1b24fb11f933 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/incomingchannels"
    i7465d68d051e059b7711de821a243bb5ea0f0c49c7901cc401f652c18ac353dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/primarychannel"
    i8d06ae5ae0b91e2f84f591114ce770d37854247efd970671c30446f38c812fe9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/group"
    i8dde215ba3c450a5fc60672f401cac784d183f1bc4087188b5b67197f373d7a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/members"
    id4403ea28d417365bd54a3ba3299d3e6dd5cf680eb1bf1dd575bba29c6fad10a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/owners"
    id462efba9e3a1914399806c5a8cab938c32345c0ccdf3c9c00080784abcef935 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/photo"
    id536d9e84ae37a007118e9b73947525fbcdabf64f1c4982299c8b7595398afdb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/tags"
    id6e9caecfa277ac9140b410d26f6b9ecb49ab815c4ea7e3e489f271e19da80df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/operations"
    i1330f0c7a93f443197dc9a115acb7eb6ea2b403fc567889eb49b0eccfb4bb3ac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/tags/item"
    i4669bd93c643fcd595168c0bb3585bd433e72b4b28050dc7eeac0dedbcb48421 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/allchannels/item"
    i64e51aa25a9d517d904bf3202a1735b74a2440188bc14ba8ebe6c2c620f6417a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/members/item"
    i9c38250d8fedf219e6ef6120c9ee2dc75706ce354bd28a888550d3ef6741bb2a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/incomingchannels/item"
    i9e7ba4c8752c81af6c246bd26b15acc58966996e67dbe3103233ac0f9eaa5af8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/permissiongrants/item"
    ia0ce47cfb13572ef8c520153596bcc84143aa17b89c92c79195b3517cd028a0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/channels/item"
    iaa417fabeab071c6b00b984389be46506fae574d387a53d5de5b19c5327c9559 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/owners/item"
    iabcc0f6c33954103499b85d655658a336b8c3ad3661b9bc5d9c7c6ba941831df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/installedapps/item"
    ib0964d6aea5f9351dd5ba52b8b744398cb38bc0f351c86ff3f8dc03a8930d31f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/operations/item"
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
func (m *TeamRequestBuilder) AllChannels()(*i2a35925faef47df44ed32466b7e2a6f0223e3e8b104ee21733303d304ba89738.AllChannelsRequestBuilder) {
    return i2a35925faef47df44ed32466b7e2a6f0223e3e8b104ee21733303d304ba89738.NewAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.allChannels.item collection
func (m *TeamRequestBuilder) AllChannelsById(id string)(*i4669bd93c643fcd595168c0bb3585bd433e72b4b28050dc7eeac0dedbcb48421.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i4669bd93c643fcd595168c0bb3585bd433e72b4b28050dc7eeac0dedbcb48421.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Channels the channels property
func (m *TeamRequestBuilder) Channels()(*i4c16420d301ac519bee3f68145c8514ea14a14d9a18a89a970cf5b1136c199bc.ChannelsRequestBuilder) {
    return i4c16420d301ac519bee3f68145c8514ea14a14d9a18a89a970cf5b1136c199bc.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.channels.item collection
func (m *TeamRequestBuilder) ChannelsById(id string)(*ia0ce47cfb13572ef8c520153596bcc84143aa17b89c92c79195b3517cd028a0c.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return ia0ce47cfb13572ef8c520153596bcc84143aa17b89c92c79195b3517cd028a0c.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamRequestBuilder) {
    m := &TeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/team{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property team for me
func (m *TeamRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property team for me
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
// CreatePatchRequestInformation update the navigation property team in me
func (m *TeamRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property team in me
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
// Delete delete navigation property team for me
func (m *TeamRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property team for me
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
func (m *TeamRequestBuilder) Group()(*i8d06ae5ae0b91e2f84f591114ce770d37854247efd970671c30446f38c812fe9.GroupRequestBuilder) {
    return i8d06ae5ae0b91e2f84f591114ce770d37854247efd970671c30446f38c812fe9.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels the incomingChannels property
func (m *TeamRequestBuilder) IncomingChannels()(*i631755ac6a9d8774786b5b39fbe602f04e0c8147c9fbdc914d0d1b24fb11f933.IncomingChannelsRequestBuilder) {
    return i631755ac6a9d8774786b5b39fbe602f04e0c8147c9fbdc914d0d1b24fb11f933.NewIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.incomingChannels.item collection
func (m *TeamRequestBuilder) IncomingChannelsById(id string)(*i9c38250d8fedf219e6ef6120c9ee2dc75706ce354bd28a888550d3ef6741bb2a.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i9c38250d8fedf219e6ef6120c9ee2dc75706ce354bd28a888550d3ef6741bb2a.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps the installedApps property
func (m *TeamRequestBuilder) InstalledApps()(*i0bd22d14acd5879034c163b078c3b1b8eb2a9adc63b0e6f5012d9195f6aca876.InstalledAppsRequestBuilder) {
    return i0bd22d14acd5879034c163b078c3b1b8eb2a9adc63b0e6f5012d9195f6aca876.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.installedApps.item collection
func (m *TeamRequestBuilder) InstalledAppsById(id string)(*iabcc0f6c33954103499b85d655658a336b8c3ad3661b9bc5d9c7c6ba941831df.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return iabcc0f6c33954103499b85d655658a336b8c3ad3661b9bc5d9c7c6ba941831df.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *TeamRequestBuilder) Members()(*i8dde215ba3c450a5fc60672f401cac784d183f1bc4087188b5b67197f373d7a5.MembersRequestBuilder) {
    return i8dde215ba3c450a5fc60672f401cac784d183f1bc4087188b5b67197f373d7a5.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.members.item collection
func (m *TeamRequestBuilder) MembersById(id string)(*i64e51aa25a9d517d904bf3202a1735b74a2440188bc14ba8ebe6c2c620f6417a.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i64e51aa25a9d517d904bf3202a1735b74a2440188bc14ba8ebe6c2c620f6417a.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *TeamRequestBuilder) Operations()(*id6e9caecfa277ac9140b410d26f6b9ecb49ab815c4ea7e3e489f271e19da80df.OperationsRequestBuilder) {
    return id6e9caecfa277ac9140b410d26f6b9ecb49ab815c4ea7e3e489f271e19da80df.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.operations.item collection
func (m *TeamRequestBuilder) OperationsById(id string)(*ib0964d6aea5f9351dd5ba52b8b744398cb38bc0f351c86ff3f8dc03a8930d31f.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return ib0964d6aea5f9351dd5ba52b8b744398cb38bc0f351c86ff3f8dc03a8930d31f.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners the owners property
func (m *TeamRequestBuilder) Owners()(*id4403ea28d417365bd54a3ba3299d3e6dd5cf680eb1bf1dd575bba29c6fad10a.OwnersRequestBuilder) {
    return id4403ea28d417365bd54a3ba3299d3e6dd5cf680eb1bf1dd575bba29c6fad10a.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.owners.item collection
func (m *TeamRequestBuilder) OwnersById(id string)(*iaa417fabeab071c6b00b984389be46506fae574d387a53d5de5b19c5327c9559.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return iaa417fabeab071c6b00b984389be46506fae574d387a53d5de5b19c5327c9559.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property team in me
func (m *TeamRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property team in me
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
func (m *TeamRequestBuilder) PermissionGrants()(*i0e48002ffeba3712f35f2a803886b2f7f6f22fdb49dcf5e23226a97d472b7b1a.PermissionGrantsRequestBuilder) {
    return i0e48002ffeba3712f35f2a803886b2f7f6f22fdb49dcf5e23226a97d472b7b1a.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.permissionGrants.item collection
func (m *TeamRequestBuilder) PermissionGrantsById(id string)(*i9e7ba4c8752c81af6c246bd26b15acc58966996e67dbe3103233ac0f9eaa5af8.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i9e7ba4c8752c81af6c246bd26b15acc58966996e67dbe3103233ac0f9eaa5af8.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *TeamRequestBuilder) Photo()(*id462efba9e3a1914399806c5a8cab938c32345c0ccdf3c9c00080784abcef935.PhotoRequestBuilder) {
    return id462efba9e3a1914399806c5a8cab938c32345c0ccdf3c9c00080784abcef935.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel the primaryChannel property
func (m *TeamRequestBuilder) PrimaryChannel()(*i7465d68d051e059b7711de821a243bb5ea0f0c49c7901cc401f652c18ac353dc.PrimaryChannelRequestBuilder) {
    return i7465d68d051e059b7711de821a243bb5ea0f0c49c7901cc401f652c18ac353dc.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule the schedule property
func (m *TeamRequestBuilder) Schedule()(*i240464d1510ca1667c1822c4c48b1244eefd38846201acf25582ab4735096048.ScheduleRequestBuilder) {
    return i240464d1510ca1667c1822c4c48b1244eefd38846201acf25582ab4735096048.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags the tags property
func (m *TeamRequestBuilder) Tags()(*id536d9e84ae37a007118e9b73947525fbcdabf64f1c4982299c8b7595398afdb.TagsRequestBuilder) {
    return id536d9e84ae37a007118e9b73947525fbcdabf64f1c4982299c8b7595398afdb.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.tags.item collection
func (m *TeamRequestBuilder) TagsById(id string)(*i1330f0c7a93f443197dc9a115acb7eb6ea2b403fc567889eb49b0eccfb4bb3ac.TeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return i1330f0c7a93f443197dc9a115acb7eb6ea2b403fc567889eb49b0eccfb4bb3ac.NewTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template the template property
func (m *TeamRequestBuilder) Template()(*i0ef26a7a288133f0d75cf96466beded00255791dbbf97db173e8a60516139727.TemplateRequestBuilder) {
    return i0ef26a7a288133f0d75cf96466beded00255791dbbf97db173e8a60516139727.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
