package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderGetQueryParameters get the default channel, **General**, of a team.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderGetQueryParameters
}
// TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) CompleteMigration()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelCompleteMigrationRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) {
    m := &TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property primaryChannel for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get the default channel, **General**, of a team.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property primaryChannel in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property primaryChannel for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) FilesFolder()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelFilesFolderRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the default channel, **General**, of a team.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) Members()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMembersRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) MembersById(id string)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) Messages()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) MessagesById(id string)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelMessagesChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property primaryChannel in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) ProvisionEmail()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelProvisionEmailRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail provides operations to call the removeEmail method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) RemoveEmail()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRemoveEmailRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) SharedWithTeams()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelSharedWithTeamsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) SharedWithTeamsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelSharedWithTeamsSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) Tabs()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelTabsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) TabsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelTabsTeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelTabsTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
