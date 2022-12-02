package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.deletedTeam entity.
type TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderGetQueryParameters the channels those are either shared with this deleted team or created in this deleted team.
type TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderGetQueryParameters
}
// TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) CompleteMigration()(*TeamworkDeletedTeamsItemChannelsItemCompleteMigrationRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewTeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) {
    m := &TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewTeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property channels for teamwork
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the channels those are either shared with this deleted team or created in this deleted team.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*TeamworkDeletedTeamsItemChannelsItemDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) FilesFolder()(*TeamworkDeletedTeamsItemChannelsItemFilesFolderRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the channels those are either shared with this deleted team or created in this deleted team.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) Members()(*TeamworkDeletedTeamsItemChannelsItemMembersRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) MembersById(id string)(*TeamworkDeletedTeamsItemChannelsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamworkDeletedTeamsItemChannelsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) Messages()(*TeamworkDeletedTeamsItemChannelsItemMessagesRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) MessagesById(id string)(*TeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewTeamworkDeletedTeamsItemChannelsItemMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in teamwork
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) ProvisionEmail()(*TeamworkDeletedTeamsItemChannelsItemProvisionEmailRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail provides operations to call the removeEmail method.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) RemoveEmail()(*TeamworkDeletedTeamsItemChannelsItemRemoveEmailRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) SharedWithTeams()(*TeamworkDeletedTeamsItemChannelsItemSharedWithTeamsRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) SharedWithTeamsById(id string)(*TeamworkDeletedTeamsItemChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return NewTeamworkDeletedTeamsItemChannelsItemSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) Tabs()(*TeamworkDeletedTeamsItemChannelsItemTabsRequestBuilder) {
    return NewTeamworkDeletedTeamsItemChannelsItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamworkDeletedTeamsItemChannelsChannelItemRequestBuilder) TabsById(id string)(*TeamworkDeletedTeamsItemChannelsItemTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewTeamworkDeletedTeamsItemChannelsItemTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
