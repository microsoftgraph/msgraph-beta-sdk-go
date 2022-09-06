package teamdefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01d3aa843ae55b558a2b07553f07f64752ef3df3702ad3d7b8793ffb00396a65 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels"
    i02827de57fc02ff264f5b78291aed174984d1ae923d83cad8769b4029940370d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/archive"
    i05e01a7cd1bdf5d9551e8b84bdd1a356ddb3c0d149efa3714e485a9f1d4af3c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule"
    i0a90aaedc4ec0357018e9c2c0cab288bb6c455b33774c505e84776dc302ae93b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/completemigration"
    i25b9b2ebc0ed8d31a86396cbdbd79766d0ca6dc180ea2b668b808395bf6c8123 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/photo"
    i2a7f69c3331e3c6294d302c3ad63b0a41d71b202349cff9ed9a8d8ecd034b44a "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/clone"
    i30cd0e56846ce0bc314e66260e3ff08cecb87b28b88b48a48644621a804f7dc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/unarchive"
    i376d33ab52594655903baee2658a8eb0bf72bb812582ddc80179eb8c7d21b5b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/operations"
    i8737cc3ee7299342142e7b5ad2bd5bb05f79251a72340e78dec8f8375e92c91d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/templatedefinition"
    i9a85f226e9945a36b36dd1100ab1ac9605fd15848e758e3d548dd310910a4956 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/installedapps"
    ia890ddcd803e7ea023c874df73637cb89437966a6dbdedf95b71ef5e9f846efb "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/template"
    iaddb69d8d018d1f468b174e8ccf6d25ad974f32bc2cdfd002572d4137eedee31 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/members"
    ib7c606576e4e69504542006e2c028ab719af041c2e7098893d6b02cbd478aa1a "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/allchannels"
    id2b7ebf24efdc65e80ff4b484b3fd25479ad08bd49549a9e2c06dbc09108e511 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/tags"
    id38e0ba36d910f19e4fc968aa8b080b63f2d11d6468dd7081171ef13731130d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/primarychannel"
    idd660ed291c330794dad096d9900331057bb0e375c05b85770ed2ca4160ae59b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/owners"
    idd8eb4f3af20f81bb6e8ca724655fadb0c315db26afc70374c50396c30fc3426 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/permissiongrants"
    if079ea5274977b63dafe8ecfd81ee0f8d821b54c97860ddc6871e92826aa2cd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/incomingchannels"
    if242908686e9e1f4e68815352ea8307d79622cd4b14f82a6502a5f6a79538135 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/sendactivitynotification"
    ifc067f4e589cefa2364cf2fdea7c0d912114e35cd455211e039c99099fff6cd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/group"
    i3889d2f0c3a1510a1a66b9ba6c2b0933c7a3bcaf53bbba2ddd3b239f035d7f28 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/allchannels/item"
    i437166336ffe89fceeb3cefa21796bb4f65c33c32a109635ef3d5c3b454575b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/permissiongrants/item"
    i6d22d87adbe6cf435cd0da761d8283b2dc15f49542c5286f8b9f9ec36928db0b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/operations/item"
    i7e4d4d36436389cdf78b45b9f02b8cd1a5ef85de40c5216c53ee78522c28a378 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/owners/item"
    ia68eea7f7d3ae5be8bc92fbc161d990e62e9bf7bf3abc49e71e57b22d7caa906 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/incomingchannels/item"
    id02a445d39ef4c7c31a05602b15e6ad28934f29331a3445cba2bbfad8a762c93 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/members/item"
    id8449ae82a8cb950e7e986f802d77a2bb5517fecdb97252c8b2d852a8b8d6be8 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/installedapps/item"
    ie8cb256548eb8dad15b6ef4c0cab173de139c711fcaa31cfed2fe7f96b01f6da "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item"
    if60812fddb6d72c19abd42431bc16adc7721cdf0863a1508844f5af1ce97d9f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/tags/item"
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
// TeamDefinitionRequestBuilderGetQueryParameters get teamDefinition from teamwork
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
func (m *TeamDefinitionRequestBuilder) AllChannels()(*ib7c606576e4e69504542006e2c028ab719af041c2e7098893d6b02cbd478aa1a.AllChannelsRequestBuilder) {
    return ib7c606576e4e69504542006e2c028ab719af041c2e7098893d6b02cbd478aa1a.NewAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.allChannels.item collection
func (m *TeamDefinitionRequestBuilder) AllChannelsById(id string)(*i3889d2f0c3a1510a1a66b9ba6c2b0933c7a3bcaf53bbba2ddd3b239f035d7f28.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return i3889d2f0c3a1510a1a66b9ba6c2b0933c7a3bcaf53bbba2ddd3b239f035d7f28.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive the archive property
func (m *TeamDefinitionRequestBuilder) Archive()(*i02827de57fc02ff264f5b78291aed174984d1ae923d83cad8769b4029940370d.ArchiveRequestBuilder) {
    return i02827de57fc02ff264f5b78291aed174984d1ae923d83cad8769b4029940370d.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels the channels property
func (m *TeamDefinitionRequestBuilder) Channels()(*i01d3aa843ae55b558a2b07553f07f64752ef3df3702ad3d7b8793ffb00396a65.ChannelsRequestBuilder) {
    return i01d3aa843ae55b558a2b07553f07f64752ef3df3702ad3d7b8793ffb00396a65.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.channels.item collection
func (m *TeamDefinitionRequestBuilder) ChannelsById(id string)(*ie8cb256548eb8dad15b6ef4c0cab173de139c711fcaa31cfed2fe7f96b01f6da.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return ie8cb256548eb8dad15b6ef4c0cab173de139c711fcaa31cfed2fe7f96b01f6da.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone the clone property
func (m *TeamDefinitionRequestBuilder) Clone()(*i2a7f69c3331e3c6294d302c3ad63b0a41d71b202349cff9ed9a8d8ecd034b44a.CloneRequestBuilder) {
    return i2a7f69c3331e3c6294d302c3ad63b0a41d71b202349cff9ed9a8d8ecd034b44a.NewCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration the completeMigration property
func (m *TeamDefinitionRequestBuilder) CompleteMigration()(*i0a90aaedc4ec0357018e9c2c0cab288bb6c455b33774c505e84776dc302ae93b.CompleteMigrationRequestBuilder) {
    return i0a90aaedc4ec0357018e9c2c0cab288bb6c455b33774c505e84776dc302ae93b.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamDefinitionRequestBuilderInternal instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamDefinitionRequestBuilder) {
    m := &TeamDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property teamDefinition for teamwork
func (m *TeamDefinitionRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property teamDefinition for teamwork
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
// CreateGetRequestInformation get teamDefinition from teamwork
func (m *TeamDefinitionRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get teamDefinition from teamwork
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
// CreatePatchRequestInformation update the navigation property teamDefinition in teamwork
func (m *TeamDefinitionRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property teamDefinition in teamwork
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
// Delete delete navigation property teamDefinition for teamwork
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
// Get get teamDefinition from teamwork
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
func (m *TeamDefinitionRequestBuilder) Group()(*ifc067f4e589cefa2364cf2fdea7c0d912114e35cd455211e039c99099fff6cd7.GroupRequestBuilder) {
    return ifc067f4e589cefa2364cf2fdea7c0d912114e35cd455211e039c99099fff6cd7.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels the incomingChannels property
func (m *TeamDefinitionRequestBuilder) IncomingChannels()(*if079ea5274977b63dafe8ecfd81ee0f8d821b54c97860ddc6871e92826aa2cd2.IncomingChannelsRequestBuilder) {
    return if079ea5274977b63dafe8ecfd81ee0f8d821b54c97860ddc6871e92826aa2cd2.NewIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.incomingChannels.item collection
func (m *TeamDefinitionRequestBuilder) IncomingChannelsById(id string)(*ia68eea7f7d3ae5be8bc92fbc161d990e62e9bf7bf3abc49e71e57b22d7caa906.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return ia68eea7f7d3ae5be8bc92fbc161d990e62e9bf7bf3abc49e71e57b22d7caa906.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps the installedApps property
func (m *TeamDefinitionRequestBuilder) InstalledApps()(*i9a85f226e9945a36b36dd1100ab1ac9605fd15848e758e3d548dd310910a4956.InstalledAppsRequestBuilder) {
    return i9a85f226e9945a36b36dd1100ab1ac9605fd15848e758e3d548dd310910a4956.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.installedApps.item collection
func (m *TeamDefinitionRequestBuilder) InstalledAppsById(id string)(*id8449ae82a8cb950e7e986f802d77a2bb5517fecdb97252c8b2d852a8b8d6be8.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return id8449ae82a8cb950e7e986f802d77a2bb5517fecdb97252c8b2d852a8b8d6be8.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *TeamDefinitionRequestBuilder) Members()(*iaddb69d8d018d1f468b174e8ccf6d25ad974f32bc2cdfd002572d4137eedee31.MembersRequestBuilder) {
    return iaddb69d8d018d1f468b174e8ccf6d25ad974f32bc2cdfd002572d4137eedee31.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.members.item collection
func (m *TeamDefinitionRequestBuilder) MembersById(id string)(*id02a445d39ef4c7c31a05602b15e6ad28934f29331a3445cba2bbfad8a762c93.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return id02a445d39ef4c7c31a05602b15e6ad28934f29331a3445cba2bbfad8a762c93.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *TeamDefinitionRequestBuilder) Operations()(*i376d33ab52594655903baee2658a8eb0bf72bb812582ddc80179eb8c7d21b5b8.OperationsRequestBuilder) {
    return i376d33ab52594655903baee2658a8eb0bf72bb812582ddc80179eb8c7d21b5b8.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.operations.item collection
func (m *TeamDefinitionRequestBuilder) OperationsById(id string)(*i6d22d87adbe6cf435cd0da761d8283b2dc15f49542c5286f8b9f9ec36928db0b.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return i6d22d87adbe6cf435cd0da761d8283b2dc15f49542c5286f8b9f9ec36928db0b.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners the owners property
func (m *TeamDefinitionRequestBuilder) Owners()(*idd660ed291c330794dad096d9900331057bb0e375c05b85770ed2ca4160ae59b.OwnersRequestBuilder) {
    return idd660ed291c330794dad096d9900331057bb0e375c05b85770ed2ca4160ae59b.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.owners.item collection
func (m *TeamDefinitionRequestBuilder) OwnersById(id string)(*i7e4d4d36436389cdf78b45b9f02b8cd1a5ef85de40c5216c53ee78522c28a378.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return i7e4d4d36436389cdf78b45b9f02b8cd1a5ef85de40c5216c53ee78522c28a378.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property teamDefinition in teamwork
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
func (m *TeamDefinitionRequestBuilder) PermissionGrants()(*idd8eb4f3af20f81bb6e8ca724655fadb0c315db26afc70374c50396c30fc3426.PermissionGrantsRequestBuilder) {
    return idd8eb4f3af20f81bb6e8ca724655fadb0c315db26afc70374c50396c30fc3426.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.permissionGrants.item collection
func (m *TeamDefinitionRequestBuilder) PermissionGrantsById(id string)(*i437166336ffe89fceeb3cefa21796bb4f65c33c32a109635ef3d5c3b454575b7.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i437166336ffe89fceeb3cefa21796bb4f65c33c32a109635ef3d5c3b454575b7.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *TeamDefinitionRequestBuilder) Photo()(*i25b9b2ebc0ed8d31a86396cbdbd79766d0ca6dc180ea2b668b808395bf6c8123.PhotoRequestBuilder) {
    return i25b9b2ebc0ed8d31a86396cbdbd79766d0ca6dc180ea2b668b808395bf6c8123.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel the primaryChannel property
func (m *TeamDefinitionRequestBuilder) PrimaryChannel()(*id38e0ba36d910f19e4fc968aa8b080b63f2d11d6468dd7081171ef13731130d0.PrimaryChannelRequestBuilder) {
    return id38e0ba36d910f19e4fc968aa8b080b63f2d11d6468dd7081171ef13731130d0.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule the schedule property
func (m *TeamDefinitionRequestBuilder) Schedule()(*i05e01a7cd1bdf5d9551e8b84bdd1a356ddb3c0d149efa3714e485a9f1d4af3c9.ScheduleRequestBuilder) {
    return i05e01a7cd1bdf5d9551e8b84bdd1a356ddb3c0d149efa3714e485a9f1d4af3c9.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification the sendActivityNotification property
func (m *TeamDefinitionRequestBuilder) SendActivityNotification()(*if242908686e9e1f4e68815352ea8307d79622cd4b14f82a6502a5f6a79538135.SendActivityNotificationRequestBuilder) {
    return if242908686e9e1f4e68815352ea8307d79622cd4b14f82a6502a5f6a79538135.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags the tags property
func (m *TeamDefinitionRequestBuilder) Tags()(*id2b7ebf24efdc65e80ff4b484b3fd25479ad08bd49549a9e2c06dbc09108e511.TagsRequestBuilder) {
    return id2b7ebf24efdc65e80ff4b484b3fd25479ad08bd49549a9e2c06dbc09108e511.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.teamTemplates.item.definitions.item.teamDefinition.tags.item collection
func (m *TeamDefinitionRequestBuilder) TagsById(id string)(*if60812fddb6d72c19abd42431bc16adc7721cdf0863a1508844f5af1ce97d9f5.TeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return if60812fddb6d72c19abd42431bc16adc7721cdf0863a1508844f5af1ce97d9f5.NewTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template the template property
func (m *TeamDefinitionRequestBuilder) Template()(*ia890ddcd803e7ea023c874df73637cb89437966a6dbdedf95b71ef5e9f846efb.TemplateRequestBuilder) {
    return ia890ddcd803e7ea023c874df73637cb89437966a6dbdedf95b71ef5e9f846efb.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplateDefinition the templateDefinition property
func (m *TeamDefinitionRequestBuilder) TemplateDefinition()(*i8737cc3ee7299342142e7b5ad2bd5bb05f79251a72340e78dec8f8375e92c91d.TemplateDefinitionRequestBuilder) {
    return i8737cc3ee7299342142e7b5ad2bd5bb05f79251a72340e78dec8f8375e92c91d.NewTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive the unarchive property
func (m *TeamDefinitionRequestBuilder) Unarchive()(*i30cd0e56846ce0bc314e66260e3ff08cecb87b28b88b48a48644621a804f7dc0.UnarchiveRequestBuilder) {
    return i30cd0e56846ce0bc314e66260e3ff08cecb87b28b88b48a48644621a804f7dc0.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
