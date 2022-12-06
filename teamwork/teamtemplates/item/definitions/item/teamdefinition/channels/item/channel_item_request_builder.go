package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i23174b3966c959cccc521016991f4508c59ba0534dda1012d148b9d87673c186 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/provisionemail"
    i712135ff983486cfbbc461434d717ad2dbe42d7e8613c02bf86f98c44a0b525d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/doesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalname"
    i9772eaf653c8fd1b17f0402a6660ada51cd5c80f976522ab7dafc8bbcde5712c "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/completemigration"
    ia7e113d3213e8f915822d1c8e3871e25f0553d4362585d72d8a3a93cb22babe1 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/filesfolder"
    ib24051c1cc505006dd08fda88ab0d714f4e4cd9066301c77737d3d6b572855f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/tabs"
    ib689a754e6dff4c3362bced348dd60e3a74b5007a109d785be6e48d98e0042c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/sharedwithteams"
    ibf1b64bd2f23ff6e7e56d1b6b8d14bc5937ef20a010edaf4433886a1b079ddf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/removeemail"
    ic817f22c185bd64fd1315db689a5e8c8e2bce12c7d4ea65a9d6297d771e7d2de "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/messages"
    id22588e70b9478071be674565dab0719df395c53280e7edb0f67eebbbfc912da "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/members"
    i0a5770a887de57448bd268bdd2f4e76e6c929e9a32dfc76be088bc1676694760 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/messages/item"
    i0cd62cc9f0b74a794db9a714f2a173dff8756f0338439f90005eabcadb37a740 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/members/item"
    ic64ca1dab181016af55324f82328d18b10b1c27a96de4997a9118c409c1ff2dd "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/tabs/item"
    ie366236adcc5dfa7add489d9f6240f103d718ca81566322fb9863da17f6c2087 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/channels/item/sharedwithteams/item"
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
// CompleteMigration provides operations to call the completeMigration method.
func (m *ChannelItemRequestBuilder) CompleteMigration()(*i9772eaf653c8fd1b17f0402a6660ada51cd5c80f976522ab7dafc8bbcde5712c.CompleteMigrationRequestBuilder) {
    return i9772eaf653c8fd1b17f0402a6660ada51cd5c80f976522ab7dafc8bbcde5712c.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property channels for teamwork
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ChannelItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property channels in teamwork
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property channels for teamwork
func (m *ChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *ChannelItemRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*i712135ff983486cfbbc461434d717ad2dbe42d7e8613c02bf86f98c44a0b525d.DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return i712135ff983486cfbbc461434d717ad2dbe42d7e8613c02bf86f98c44a0b525d.NewDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) FilesFolder()(*ia7e113d3213e8f915822d1c8e3871e25f0553d4362585d72d8a3a93cb22babe1.FilesFolderRequestBuilder) {
    return ia7e113d3213e8f915822d1c8e3871e25f0553d4362585d72d8a3a93cb22babe1.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) Members()(*id22588e70b9478071be674565dab0719df395c53280e7edb0f67eebbbfc912da.MembersRequestBuilder) {
    return id22588e70b9478071be674565dab0719df395c53280e7edb0f67eebbbfc912da.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) MembersById(id string)(*i0cd62cc9f0b74a794db9a714f2a173dff8756f0338439f90005eabcadb37a740.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i0cd62cc9f0b74a794db9a714f2a173dff8756f0338439f90005eabcadb37a740.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) Messages()(*ic817f22c185bd64fd1315db689a5e8c8e2bce12c7d4ea65a9d6297d771e7d2de.MessagesRequestBuilder) {
    return ic817f22c185bd64fd1315db689a5e8c8e2bce12c7d4ea65a9d6297d771e7d2de.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*i0a5770a887de57448bd268bdd2f4e76e6c929e9a32dfc76be088bc1676694760.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i0a5770a887de57448bd268bdd2f4e76e6c929e9a32dfc76be088bc1676694760.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in teamwork
func (m *ChannelItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// ProvisionEmail provides operations to call the provisionEmail method.
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i23174b3966c959cccc521016991f4508c59ba0534dda1012d148b9d87673c186.ProvisionEmailRequestBuilder) {
    return i23174b3966c959cccc521016991f4508c59ba0534dda1012d148b9d87673c186.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail provides operations to call the removeEmail method.
func (m *ChannelItemRequestBuilder) RemoveEmail()(*ibf1b64bd2f23ff6e7e56d1b6b8d14bc5937ef20a010edaf4433886a1b079ddf5.RemoveEmailRequestBuilder) {
    return ibf1b64bd2f23ff6e7e56d1b6b8d14bc5937ef20a010edaf4433886a1b079ddf5.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) SharedWithTeams()(*ib689a754e6dff4c3362bced348dd60e3a74b5007a109d785be6e48d98e0042c1.SharedWithTeamsRequestBuilder) {
    return ib689a754e6dff4c3362bced348dd60e3a74b5007a109d785be6e48d98e0042c1.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) SharedWithTeamsById(id string)(*ie366236adcc5dfa7add489d9f6240f103d718ca81566322fb9863da17f6c2087.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return ie366236adcc5dfa7add489d9f6240f103d718ca81566322fb9863da17f6c2087.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) Tabs()(*ib24051c1cc505006dd08fda88ab0d714f4e4cd9066301c77737d3d6b572855f3.TabsRequestBuilder) {
    return ib24051c1cc505006dd08fda88ab0d714f4e4cd9066301c77737d3d6b572855f3.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *ChannelItemRequestBuilder) TabsById(id string)(*ic64ca1dab181016af55324f82328d18b10b1c27a96de4997a9118c409c1ff2dd.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return ic64ca1dab181016af55324f82328d18b10b1c27a96de4997a9118c409c1ff2dd.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
