package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters
}
// TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) CompleteMigration()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemCompleteMigrationRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    m := &TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder{
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
// NewTeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property channels for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property channels for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) FilesFolder()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemFilesFolderRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) Members()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemMembersRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) MembersById(id string)(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) Messages()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemMessagesRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) MessagesById(id string)(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) ProvisionEmail()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemProvisionEmailRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail provides operations to call the removeEmail method.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) RemoveEmail()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemRemoveEmailRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) SharedWithTeams()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemSharedWithTeamsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) SharedWithTeamsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) Tabs()(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemTabsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) TabsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionChannelsItemTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsItemTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
