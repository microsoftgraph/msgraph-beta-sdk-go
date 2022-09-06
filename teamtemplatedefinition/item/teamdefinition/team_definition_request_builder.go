package teamdefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06669f89aa10d33f3324e6c38c914f53e35ad89608734918f08156f5cc3b707e "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/permissiongrants"
    i0e6f3a7a73cee7f768cb152f7d405f796cdbd7247f0c124cef6b16ba213fb9ab "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/template"
    i309969c86942ebfb44e1c1d4e06638fd50b311b9d33b0f81e4327392e2a13611 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/sendactivitynotification"
    i3599fe11def4426815875877f539ae68f359ff4e094dcb0f3f693d59581a1272 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/archive"
    i488672bee064dab885477f581520812e1c2bca16bf8360e46c811ddf7d5c0c2b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/schedule"
    i4f7d26813ec739800d67f5accfb0d57f10ad2160a848a7d2cd42d83d75d0ad2b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/operations"
    i53f8b1928746e03f49d2d29b7f430fed2cd4a8beb98a4254d32b7313baff0d07 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/incomingchannels"
    i59f0e35559aa0c747f8247a29112665a4c5f7c5986f92cd45b697ea5a40da935 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/photo"
    i5e45a1972dc49830245c85afa85cff5f4192010146f48af8d1243cd8e86c6d74 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/owners"
    i6cfb900b23c46a8fde31fdf87d569e2d5765fa08711edb2caac718640fb9d4fb "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/clone"
    i722f784d97ba6f5ad90ed27a970cf9f406c19786076a5b0f4645e377e11a926c "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/primarychannel"
    i89195ea342b7ec59a51fc947ca2297d4c486fb5481eab7e64d3ae70439e602b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/tags"
    i97bc522cbde108ec1513f40ba6505004b143bdf919d6a10d72a36cc7caf26c68 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/templatedefinition"
    ia761ada2a1af9fc46cb8f2314c4aa620b1015cd14aeb34145a73e553de26355b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/unarchive"
    ib34eb25f9c94bd2eed7209724fe99e44a342ffdb57364329cf799a6562c0b3e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/completemigration"
    ib874097dfdadb36724631faad7d63e2d5631bb46dc68b90f4c4d1fc4f8a75666 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/group"
    ibdc4e79164f9d758d80a3545c73d94dced4b5caa68bff30274ab5d3e6d4d93cc "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/members"
    ibfb7b8aa6fb3463d097e6e244a7f6bfaecf7d2f1d1ce32ed70d436eee13c9fb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/installedapps"
    id851d7cb6c796133b5c2cfa36fe4ee7743eabbd355376217105e58c114c50eab "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels"
    ifd22a8b3a6eba59eaf6f04c7592834eeebe5d9b057b4225ce6bc399a698838bd "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/allchannels"
    i13306a4e10ad8672063e90a5a6ea0d03ccf0912a68449b93d634fba98d794437 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/members/item"
    i2379a8c42f2ea6a6b2e96d0c96399b30e569c44d2c5ef9984e5b7a0094839d5d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/permissiongrants/item"
    i2b58435154b6aa90371d199c399a088325cc21b98c3e054be6ec737c487a27e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/allchannels/item"
    i3454b3adff6fab7441c2d8b209e49775132eb97ac78de47dc3f63f82a7235631 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/owners/item"
    i349df068b3e1cc967c3b750b6c759ed91ecb356f259a98be7da1aabb080ac1c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item"
    i3ccef4dd4ee0b3e5a1c00200f6fff0fac02183e816d553a61b332f87a3f55b5d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/incomingchannels/item"
    i45753fa859ef2f40ff54a91c54bfab46e4e6a2a4e91fb4811822f753c92f8033 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/operations/item"
    i64e892be62d748d1d2f4bcf268b9dc7dbf36f2355af48f13cdfb7fee9be10ed7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/tags/item"
    iaf9d47d526208e5deb90ced05087e6917acd8e544e0e82aba5078dc282a8e98a "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/installedapps/item"
)

// TeamDefinitionRequestBuilder provides operations to manage the teamDefinition property of the microsoft.graph.teamTemplateDefinition entity.
type TeamDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamDefinitionRequestBuilderGetQueryParameters get teamDefinition from teamTemplateDefinition
type TeamDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamDefinitionRequestBuilderGetQueryParameters
}
// TeamDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels the allChannels property
func (m *TeamDefinitionRequestBuilder) AllChannels()(*ifd22a8b3a6eba59eaf6f04c7592834eeebe5d9b057b4225ce6bc399a698838bd.AllChannelsRequestBuilder) {
    return ifd22a8b3a6eba59eaf6f04c7592834eeebe5d9b057b4225ce6bc399a698838bd.NewAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.allChannels.item collection
func (m *TeamDefinitionRequestBuilder) AllChannelsById(id string)(*i2b58435154b6aa90371d199c399a088325cc21b98c3e054be6ec737c487a27e4.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i2b58435154b6aa90371d199c399a088325cc21b98c3e054be6ec737c487a27e4.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive the archive property
func (m *TeamDefinitionRequestBuilder) Archive()(*i3599fe11def4426815875877f539ae68f359ff4e094dcb0f3f693d59581a1272.ArchiveRequestBuilder) {
    return i3599fe11def4426815875877f539ae68f359ff4e094dcb0f3f693d59581a1272.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels the channels property
func (m *TeamDefinitionRequestBuilder) Channels()(*id851d7cb6c796133b5c2cfa36fe4ee7743eabbd355376217105e58c114c50eab.ChannelsRequestBuilder) {
    return id851d7cb6c796133b5c2cfa36fe4ee7743eabbd355376217105e58c114c50eab.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.channels.item collection
func (m *TeamDefinitionRequestBuilder) ChannelsById(id string)(*i349df068b3e1cc967c3b750b6c759ed91ecb356f259a98be7da1aabb080ac1c8.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i349df068b3e1cc967c3b750b6c759ed91ecb356f259a98be7da1aabb080ac1c8.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone the clone property
func (m *TeamDefinitionRequestBuilder) Clone()(*i6cfb900b23c46a8fde31fdf87d569e2d5765fa08711edb2caac718640fb9d4fb.CloneRequestBuilder) {
    return i6cfb900b23c46a8fde31fdf87d569e2d5765fa08711edb2caac718640fb9d4fb.NewCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration the completeMigration property
func (m *TeamDefinitionRequestBuilder) CompleteMigration()(*ib34eb25f9c94bd2eed7209724fe99e44a342ffdb57364329cf799a6562c0b3e4.CompleteMigrationRequestBuilder) {
    return ib34eb25f9c94bd2eed7209724fe99e44a342ffdb57364329cf799a6562c0b3e4.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamDefinitionRequestBuilderInternal instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamDefinitionRequestBuilder) {
    m := &TeamDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamDefinitionRequestBuilder instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property teamDefinition for teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property teamDefinition for teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *TeamDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get teamDefinition from teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get teamDefinition from teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TeamDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property teamDefinition in teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property teamDefinition in teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property teamDefinition for teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get teamDefinition from teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
func (m *TeamDefinitionRequestBuilder) Group()(*ib874097dfdadb36724631faad7d63e2d5631bb46dc68b90f4c4d1fc4f8a75666.GroupRequestBuilder) {
    return ib874097dfdadb36724631faad7d63e2d5631bb46dc68b90f4c4d1fc4f8a75666.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels the incomingChannels property
func (m *TeamDefinitionRequestBuilder) IncomingChannels()(*i53f8b1928746e03f49d2d29b7f430fed2cd4a8beb98a4254d32b7313baff0d07.IncomingChannelsRequestBuilder) {
    return i53f8b1928746e03f49d2d29b7f430fed2cd4a8beb98a4254d32b7313baff0d07.NewIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.incomingChannels.item collection
func (m *TeamDefinitionRequestBuilder) IncomingChannelsById(id string)(*i3ccef4dd4ee0b3e5a1c00200f6fff0fac02183e816d553a61b332f87a3f55b5d.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i3ccef4dd4ee0b3e5a1c00200f6fff0fac02183e816d553a61b332f87a3f55b5d.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps the installedApps property
func (m *TeamDefinitionRequestBuilder) InstalledApps()(*ibfb7b8aa6fb3463d097e6e244a7f6bfaecf7d2f1d1ce32ed70d436eee13c9fb5.InstalledAppsRequestBuilder) {
    return ibfb7b8aa6fb3463d097e6e244a7f6bfaecf7d2f1d1ce32ed70d436eee13c9fb5.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.installedApps.item collection
func (m *TeamDefinitionRequestBuilder) InstalledAppsById(id string)(*iaf9d47d526208e5deb90ced05087e6917acd8e544e0e82aba5078dc282a8e98a.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return iaf9d47d526208e5deb90ced05087e6917acd8e544e0e82aba5078dc282a8e98a.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *TeamDefinitionRequestBuilder) Members()(*ibdc4e79164f9d758d80a3545c73d94dced4b5caa68bff30274ab5d3e6d4d93cc.MembersRequestBuilder) {
    return ibdc4e79164f9d758d80a3545c73d94dced4b5caa68bff30274ab5d3e6d4d93cc.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.members.item collection
func (m *TeamDefinitionRequestBuilder) MembersById(id string)(*i13306a4e10ad8672063e90a5a6ea0d03ccf0912a68449b93d634fba98d794437.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i13306a4e10ad8672063e90a5a6ea0d03ccf0912a68449b93d634fba98d794437.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *TeamDefinitionRequestBuilder) Operations()(*i4f7d26813ec739800d67f5accfb0d57f10ad2160a848a7d2cd42d83d75d0ad2b.OperationsRequestBuilder) {
    return i4f7d26813ec739800d67f5accfb0d57f10ad2160a848a7d2cd42d83d75d0ad2b.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.operations.item collection
func (m *TeamDefinitionRequestBuilder) OperationsById(id string)(*i45753fa859ef2f40ff54a91c54bfab46e4e6a2a4e91fb4811822f753c92f8033.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return i45753fa859ef2f40ff54a91c54bfab46e4e6a2a4e91fb4811822f753c92f8033.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners the owners property
func (m *TeamDefinitionRequestBuilder) Owners()(*i5e45a1972dc49830245c85afa85cff5f4192010146f48af8d1243cd8e86c6d74.OwnersRequestBuilder) {
    return i5e45a1972dc49830245c85afa85cff5f4192010146f48af8d1243cd8e86c6d74.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.owners.item collection
func (m *TeamDefinitionRequestBuilder) OwnersById(id string)(*i3454b3adff6fab7441c2d8b209e49775132eb97ac78de47dc3f63f82a7235631.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return i3454b3adff6fab7441c2d8b209e49775132eb97ac78de47dc3f63f82a7235631.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property teamDefinition in teamTemplateDefinition
func (m *TeamDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamDefinitionRequestBuilderPatchRequestConfiguration)(error) {
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
func (m *TeamDefinitionRequestBuilder) PermissionGrants()(*i06669f89aa10d33f3324e6c38c914f53e35ad89608734918f08156f5cc3b707e.PermissionGrantsRequestBuilder) {
    return i06669f89aa10d33f3324e6c38c914f53e35ad89608734918f08156f5cc3b707e.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.permissionGrants.item collection
func (m *TeamDefinitionRequestBuilder) PermissionGrantsById(id string)(*i2379a8c42f2ea6a6b2e96d0c96399b30e569c44d2c5ef9984e5b7a0094839d5d.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i2379a8c42f2ea6a6b2e96d0c96399b30e569c44d2c5ef9984e5b7a0094839d5d.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *TeamDefinitionRequestBuilder) Photo()(*i59f0e35559aa0c747f8247a29112665a4c5f7c5986f92cd45b697ea5a40da935.PhotoRequestBuilder) {
    return i59f0e35559aa0c747f8247a29112665a4c5f7c5986f92cd45b697ea5a40da935.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel the primaryChannel property
func (m *TeamDefinitionRequestBuilder) PrimaryChannel()(*i722f784d97ba6f5ad90ed27a970cf9f406c19786076a5b0f4645e377e11a926c.PrimaryChannelRequestBuilder) {
    return i722f784d97ba6f5ad90ed27a970cf9f406c19786076a5b0f4645e377e11a926c.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule the schedule property
func (m *TeamDefinitionRequestBuilder) Schedule()(*i488672bee064dab885477f581520812e1c2bca16bf8360e46c811ddf7d5c0c2b.ScheduleRequestBuilder) {
    return i488672bee064dab885477f581520812e1c2bca16bf8360e46c811ddf7d5c0c2b.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification the sendActivityNotification property
func (m *TeamDefinitionRequestBuilder) SendActivityNotification()(*i309969c86942ebfb44e1c1d4e06638fd50b311b9d33b0f81e4327392e2a13611.SendActivityNotificationRequestBuilder) {
    return i309969c86942ebfb44e1c1d4e06638fd50b311b9d33b0f81e4327392e2a13611.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags the tags property
func (m *TeamDefinitionRequestBuilder) Tags()(*i89195ea342b7ec59a51fc947ca2297d4c486fb5481eab7e64d3ae70439e602b3.TagsRequestBuilder) {
    return i89195ea342b7ec59a51fc947ca2297d4c486fb5481eab7e64d3ae70439e602b3.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.tags.item collection
func (m *TeamDefinitionRequestBuilder) TagsById(id string)(*i64e892be62d748d1d2f4bcf268b9dc7dbf36f2355af48f13cdfb7fee9be10ed7.TeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return i64e892be62d748d1d2f4bcf268b9dc7dbf36f2355af48f13cdfb7fee9be10ed7.NewTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template the template property
func (m *TeamDefinitionRequestBuilder) Template()(*i0e6f3a7a73cee7f768cb152f7d405f796cdbd7247f0c124cef6b16ba213fb9ab.TemplateRequestBuilder) {
    return i0e6f3a7a73cee7f768cb152f7d405f796cdbd7247f0c124cef6b16ba213fb9ab.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplateDefinition the templateDefinition property
func (m *TeamDefinitionRequestBuilder) TemplateDefinition()(*i97bc522cbde108ec1513f40ba6505004b143bdf919d6a10d72a36cc7caf26c68.TemplateDefinitionRequestBuilder) {
    return i97bc522cbde108ec1513f40ba6505004b143bdf919d6a10d72a36cc7caf26c68.NewTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive the unarchive property
func (m *TeamDefinitionRequestBuilder) Unarchive()(*ia761ada2a1af9fc46cb8f2314c4aa620b1015cd14aeb34145a73e553de26355b.UnarchiveRequestBuilder) {
    return ia761ada2a1af9fc46cb8f2314c4aa620b1015cd14aeb34145a73e553de26355b.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
