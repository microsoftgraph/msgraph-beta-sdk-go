package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) CompleteMigration()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemCompleteMigrationRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder{
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
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property channels for teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) FilesFolder()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemFilesFolderRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) Members()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) MembersById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) Messages()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) MessagesById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) ProvisionEmail()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemProvisionEmailRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail provides operations to call the removeEmail method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) RemoveEmail()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemRemoveEmailRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) SharedWithTeams()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemSharedWithTeamsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) SharedWithTeamsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) Tabs()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemTabsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) TabsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
