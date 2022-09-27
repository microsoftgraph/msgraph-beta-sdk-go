package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i114446cf99a8c2cab55bcbf7f2812842087ff6ca04506525cbbf85ce10ad446f "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/completemigration"
    i1c9c72c64b9c5ca81a6efda3dc7072b3f4cb5b3f98708515c54c1c12dec7505b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/provisionemail"
    i246c4ec01a943d4243f06f13bbfe1a15d8c228da1fc81cf229f9ec621ba80240 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/sharedwithteams"
    i2e5332869b3935b69006e2b5f7d4416c38293be14ea94f95bad97617032c19fb "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/doesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalname"
    i47858a15ced41c8104f1ef7358992df4ba49d645d7708a39273d75a019712385 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/members"
    i4e10bca7b4f870ae0a347ba6e7e1c528dfdb0e44eef9a2d70bcb50d97293a1d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/removeemail"
    i6383d2cbaddc4d7b219cfa9bf48792d718da8865ec877ab5b534465b502783e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/messages"
    i7706307c366a954d91d65bba26a3e94b73712ee892f8681752855e64913dfa20 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/tabs"
    if0ae075362be61b54b851f5e3dcee513aefe3e48fb64b811fee0b92f22fcad37 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/filesfolder"
    i1d05f633e15dab3a5e1b1925f12035ca899bf3f9e81f95defb8461288635b434 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/members/item"
    i27483dcc8990bc921343302a8443441ac849b164d100d2c5f412a17939549bdc "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/tabs/item"
    i85f747a5ec85016eb2aebe8fbeacd53ca61cbdfb50db612a0171c28545ccf49d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/sharedwithteams/item"
    ibb2bf6e047ded610dd544ab8cddad89da97080cd28b501343c7f361845040a3b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item/teamdefinition/channels/item/messages/item"
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
func (m *ChannelItemRequestBuilder) CompleteMigration()(*i114446cf99a8c2cab55bcbf7f2812842087ff6ca04506525cbbf85ce10ad446f.CompleteMigrationRequestBuilder) {
    return i114446cf99a8c2cab55bcbf7f2812842087ff6ca04506525cbbf85ce10ad446f.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property channels for teamTemplateDefinition
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property channels for teamTemplateDefinition
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
// CreatePatchRequestInformation update the navigation property channels in teamTemplateDefinition
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property channels in teamTemplateDefinition
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property channels for teamTemplateDefinition
func (m *ChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *ChannelItemRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*i2e5332869b3935b69006e2b5f7d4416c38293be14ea94f95bad97617032c19fb.DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return i2e5332869b3935b69006e2b5f7d4416c38293be14ea94f95bad97617032c19fb.NewDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder the filesFolder property
func (m *ChannelItemRequestBuilder) FilesFolder()(*if0ae075362be61b54b851f5e3dcee513aefe3e48fb64b811fee0b92f22fcad37.FilesFolderRequestBuilder) {
    return if0ae075362be61b54b851f5e3dcee513aefe3e48fb64b811fee0b92f22fcad37.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Members the members property
func (m *ChannelItemRequestBuilder) Members()(*i47858a15ced41c8104f1ef7358992df4ba49d645d7708a39273d75a019712385.MembersRequestBuilder) {
    return i47858a15ced41c8104f1ef7358992df4ba49d645d7708a39273d75a019712385.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*i1d05f633e15dab3a5e1b1925f12035ca899bf3f9e81f95defb8461288635b434.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i1d05f633e15dab3a5e1b1925f12035ca899bf3f9e81f95defb8461288635b434.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *ChannelItemRequestBuilder) Messages()(*i6383d2cbaddc4d7b219cfa9bf48792d718da8865ec877ab5b534465b502783e5.MessagesRequestBuilder) {
    return i6383d2cbaddc4d7b219cfa9bf48792d718da8865ec877ab5b534465b502783e5.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*ibb2bf6e047ded610dd544ab8cddad89da97080cd28b501343c7f361845040a3b.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return ibb2bf6e047ded610dd544ab8cddad89da97080cd28b501343c7f361845040a3b.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in teamTemplateDefinition
func (m *ChannelItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// ProvisionEmail the provisionEmail property
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i1c9c72c64b9c5ca81a6efda3dc7072b3f4cb5b3f98708515c54c1c12dec7505b.ProvisionEmailRequestBuilder) {
    return i1c9c72c64b9c5ca81a6efda3dc7072b3f4cb5b3f98708515c54c1c12dec7505b.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *ChannelItemRequestBuilder) RemoveEmail()(*i4e10bca7b4f870ae0a347ba6e7e1c528dfdb0e44eef9a2d70bcb50d97293a1d0.RemoveEmailRequestBuilder) {
    return i4e10bca7b4f870ae0a347ba6e7e1c528dfdb0e44eef9a2d70bcb50d97293a1d0.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams the sharedWithTeams property
func (m *ChannelItemRequestBuilder) SharedWithTeams()(*i246c4ec01a943d4243f06f13bbfe1a15d8c228da1fc81cf229f9ec621ba80240.SharedWithTeamsRequestBuilder) {
    return i246c4ec01a943d4243f06f13bbfe1a15d8c228da1fc81cf229f9ec621ba80240.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.channels.item.sharedWithTeams.item collection
func (m *ChannelItemRequestBuilder) SharedWithTeamsById(id string)(*i85f747a5ec85016eb2aebe8fbeacd53ca61cbdfb50db612a0171c28545ccf49d.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return i85f747a5ec85016eb2aebe8fbeacd53ca61cbdfb50db612a0171c28545ccf49d.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs the tabs property
func (m *ChannelItemRequestBuilder) Tabs()(*i7706307c366a954d91d65bba26a3e94b73712ee892f8681752855e64913dfa20.TabsRequestBuilder) {
    return i7706307c366a954d91d65bba26a3e94b73712ee892f8681752855e64913dfa20.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item.teamDefinition.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i27483dcc8990bc921343302a8443441ac849b164d100d2c5f412a17939549bdc.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i27483dcc8990bc921343302a8443441ac849b164d100d2c5f412a17939549bdc.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
