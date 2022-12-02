package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder provides operations to manage the teamDefinition property of the microsoft.graph.teamTemplateDefinition entity.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetQueryParameters get the properties of the team associated with a teamTemplateDefinition object.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetQueryParameters
}
// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) AllChannels()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) AllChannelsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive provides operations to call the archive method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Archive()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionArchiveRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Channels()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) ChannelsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone provides operations to call the clone method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Clone()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionCloneRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) CompleteMigration()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionCompleteMigrationRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderInternal instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder{
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
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property teamDefinition for teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get the properties of the team associated with a teamTemplateDefinition object.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property teamDefinition for teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties of the team associated with a teamTemplateDefinition object.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
// Group provides operations to manage the group property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Group()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionGroupRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) IncomingChannels()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) IncomingChannelsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) InstalledApps()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionInstalledAppsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) InstalledAppsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Members()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionMembersRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) MembersById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Operations()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionOperationsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) OperationsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Owners()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) OwnersById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionOwnersUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property teamDefinition in teamwork
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) PermissionGrants()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) PermissionGrantsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Photo()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPhotoRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) PrimaryChannel()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Schedule()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) SendActivityNotification()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionSendActivityNotificationRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Tags()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionTagsRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) TagsById(id string)(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Template()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionTemplateRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplateDefinition provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) TemplateDefinition()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionTemplateDefinitionRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive provides operations to call the unarchive method.
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionRequestBuilder) Unarchive()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionUnarchiveRequestBuilder) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
