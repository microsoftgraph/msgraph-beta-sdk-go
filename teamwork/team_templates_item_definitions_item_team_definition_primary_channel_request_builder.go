package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderGetQueryParameters get the default channel, **General**, of a team.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderGetQueryParameters
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property primaryChannel for teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) FilesFolder()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelFilesFolderRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get the default channel, **General**, of a team.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/team-get-primarychannel?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) Members()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMembersRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) MembersById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) Messages()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) MessagesById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphCompleteMigration provides operations to call the completeMigration method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) MicrosoftGraphCompleteMigration()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMicrosoftGraphCompleteMigrationRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMicrosoftGraphCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) MicrosoftGraphDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMicrosoftGraphDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMicrosoftGraphDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphProvisionEmail provides operations to call the provisionEmail method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) MicrosoftGraphProvisionEmail()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMicrosoftGraphProvisionEmailRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMicrosoftGraphProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoveEmail provides operations to call the removeEmail method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) MicrosoftGraphRemoveEmail()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMicrosoftGraphRemoveEmailRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelMicrosoftGraphRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property primaryChannel in teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) SharedWithTeams()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelSharedWithTeamsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SharedWithTeamsById provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) SharedWithTeamsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) Tabs()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelTabsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) TabsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property primaryChannel for teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get the default channel, **General**, of a team.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property primaryChannel in teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
