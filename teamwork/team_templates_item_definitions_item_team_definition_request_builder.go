package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder provides operations to manage the teamDefinition property of the microsoft.graph.teamTemplateDefinition entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetQueryParameters get the properties of the team associated with a teamTemplateDefinition object.
type TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetQueryParameters
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) AllChannels()(*TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) AllChannelsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive provides operations to call the archive method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Archive()(*TeamTemplatesItemDefinitionsItemTeamDefinitionArchiveRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Channels()(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) ChannelsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone provides operations to call the clone method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Clone()(*TeamTemplatesItemDefinitionsItemTeamDefinitionCloneRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) CompleteMigration()(*TeamTemplatesItemDefinitionsItemTeamDefinitionCompleteMigrationRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderInternal instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder{
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
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property teamDefinition for teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties of the team associated with a teamTemplateDefinition object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/teamtemplatedefinition-get-teamdefinition?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Group()(*TeamTemplatesItemDefinitionsItemTeamDefinitionGroupRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) IncomingChannels()(*TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) IncomingChannelsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) InstalledApps()(*TeamTemplatesItemDefinitionsItemTeamDefinitionInstalledAppsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) InstalledAppsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Members()(*TeamTemplatesItemDefinitionsItemTeamDefinitionMembersRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) MembersById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Operations()(*TeamTemplatesItemDefinitionsItemTeamDefinitionOperationsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) OperationsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Owners()(*TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) OwnersById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property teamDefinition in teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable), nil
}
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) PermissionGrants()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) PermissionGrantsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Photo()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPhotoRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) PrimaryChannel()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Schedule()(*TeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) SendActivityNotification()(*TeamTemplatesItemDefinitionsItemTeamDefinitionSendActivityNotificationRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Tags()(*TeamTemplatesItemDefinitionsItemTeamDefinitionTagsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) TagsById(id string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Template()(*TeamTemplatesItemDefinitionsItemTeamDefinitionTemplateRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplateDefinition provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) TemplateDefinition()(*TeamTemplatesItemDefinitionsItemTeamDefinitionTemplateDefinitionRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToDeleteRequestInformation delete navigation property teamDefinition for teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get the properties of the team associated with a teamTemplateDefinition object.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property teamDefinition in teamwork
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Unarchive provides operations to call the unarchive method.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Unarchive()(*TeamTemplatesItemDefinitionsItemTeamDefinitionUnarchiveRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
