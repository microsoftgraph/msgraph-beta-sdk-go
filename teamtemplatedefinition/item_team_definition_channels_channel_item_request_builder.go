package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionChannelsChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ItemTeamDefinitionChannelsChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type ItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionChannelsChannelItemRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamDefinitionChannelsChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    m := &ItemTeamDefinitionChannelsChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamDefinitionChannelsChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewItemTeamDefinitionChannelsChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property channels for teamTemplateDefinition
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) FilesFolder()(*ItemTeamDefinitionChannelsItemFilesFolderRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the collection of channels and messages associated with the team.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) Members()(*ItemTeamDefinitionChannelsItemMembersRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) MembersById(id string)(*ItemTeamDefinitionChannelsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemTeamDefinitionChannelsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) Messages()(*ItemTeamDefinitionChannelsItemMessagesRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) MessagesById(id string)(*ItemTeamDefinitionChannelsItemMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewItemTeamDefinitionChannelsItemMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphCompleteMigration provides operations to call the completeMigration method.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) MicrosoftGraphCompleteMigration()(*ItemTeamDefinitionChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMicrosoftGraphCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) MicrosoftGraphDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*ItemTeamDefinitionChannelsItemMicrosoftGraphDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMicrosoftGraphDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphProvisionEmail provides operations to call the provisionEmail method.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) MicrosoftGraphProvisionEmail()(*ItemTeamDefinitionChannelsItemMicrosoftGraphProvisionEmailRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMicrosoftGraphProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoveEmail provides operations to call the removeEmail method.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) MicrosoftGraphRemoveEmail()(*ItemTeamDefinitionChannelsItemMicrosoftGraphRemoveEmailRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemMicrosoftGraphRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property channels in teamTemplateDefinition
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) SharedWithTeams()(*ItemTeamDefinitionChannelsItemSharedWithTeamsRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SharedWithTeamsById provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) SharedWithTeamsById(id string)(*ItemTeamDefinitionChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return NewItemTeamDefinitionChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) Tabs()(*ItemTeamDefinitionChannelsItemTabsRequestBuilder) {
    return NewItemTeamDefinitionChannelsItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) TabsById(id string)(*ItemTeamDefinitionChannelsItemTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewItemTeamDefinitionChannelsItemTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property channels for teamTemplateDefinition
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the collection of channels and messages associated with the team.
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionChannelsChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property channels in teamTemplateDefinition
func (m *ItemTeamDefinitionChannelsChannelItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ItemTeamDefinitionChannelsChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
