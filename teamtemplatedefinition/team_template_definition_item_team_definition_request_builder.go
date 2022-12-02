package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplateDefinitionItemTeamDefinitionRequestBuilder provides operations to manage the teamDefinition property of the microsoft.graph.teamTemplateDefinition entity.
type TeamTemplateDefinitionItemTeamDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplateDefinitionItemTeamDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplateDefinitionItemTeamDefinitionRequestBuilderGetQueryParameters get the properties of the team associated with a teamTemplateDefinition object.
type TeamTemplateDefinitionItemTeamDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplateDefinitionItemTeamDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplateDefinitionItemTeamDefinitionRequestBuilderGetQueryParameters
}
// TeamTemplateDefinitionItemTeamDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) AllChannels()(*TeamTemplateDefinitionItemTeamDefinitionAllChannelsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) AllChannelsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive provides operations to call the archive method.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Archive()(*TeamTemplateDefinitionItemTeamDefinitionArchiveRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Channels()(*TeamTemplateDefinitionItemTeamDefinitionChannelsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) ChannelsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone provides operations to call the clone method.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Clone()(*TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) CompleteMigration()(*TeamTemplateDefinitionItemTeamDefinitionCompleteMigrationRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamTemplateDefinitionItemTeamDefinitionRequestBuilderInternal instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) {
    m := &TeamTemplateDefinitionItemTeamDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamTemplateDefinitionItemTeamDefinitionRequestBuilder instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplateDefinitionItemTeamDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property teamDefinition for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property teamDefinition in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property teamDefinition for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Group()(*TeamTemplateDefinitionItemTeamDefinitionGroupRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) IncomingChannels()(*TeamTemplateDefinitionItemTeamDefinitionIncomingChannelsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) IncomingChannelsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) InstalledApps()(*TeamTemplateDefinitionItemTeamDefinitionInstalledAppsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) InstalledAppsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Members()(*TeamTemplateDefinitionItemTeamDefinitionMembersRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) MembersById(id string)(*TeamTemplateDefinitionItemTeamDefinitionMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Operations()(*TeamTemplateDefinitionItemTeamDefinitionOperationsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) OperationsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Owners()(*TeamTemplateDefinitionItemTeamDefinitionOwnersRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) OwnersById(id string)(*TeamTemplateDefinitionItemTeamDefinitionOwnersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionOwnersUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property teamDefinition in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) PermissionGrants()(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) PermissionGrantsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Photo()(*TeamTemplateDefinitionItemTeamDefinitionPhotoRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) PrimaryChannel()(*TeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Schedule()(*TeamTemplateDefinitionItemTeamDefinitionScheduleRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) SendActivityNotification()(*TeamTemplateDefinitionItemTeamDefinitionSendActivityNotificationRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Tags()(*TeamTemplateDefinitionItemTeamDefinitionTagsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) TagsById(id string)(*TeamTemplateDefinitionItemTeamDefinitionTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewTeamTemplateDefinitionItemTeamDefinitionTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Template()(*TeamTemplateDefinitionItemTeamDefinitionTemplateRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplateDefinition provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) TemplateDefinition()(*TeamTemplateDefinitionItemTeamDefinitionTemplateDefinitionRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive provides operations to call the unarchive method.
func (m *TeamTemplateDefinitionItemTeamDefinitionRequestBuilder) Unarchive()(*TeamTemplateDefinitionItemTeamDefinitionUnarchiveRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
