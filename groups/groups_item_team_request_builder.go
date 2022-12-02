package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemTeamRequestBuilder provides operations to manage the team property of the microsoft.graph.group entity.
type GroupsItemTeamRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemTeamRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemTeamRequestBuilderGetQueryParameters the team associated with this group.
type GroupsItemTeamRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemTeamRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemTeamRequestBuilderGetQueryParameters
}
// GroupsItemTeamRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) AllChannels()(*GroupsItemTeamAllChannelsRequestBuilder) {
    return NewGroupsItemTeamAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) AllChannelsById(id string)(*GroupsItemTeamAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewGroupsItemTeamAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive provides operations to call the archive method.
func (m *GroupsItemTeamRequestBuilder) Archive()(*GroupsItemTeamArchiveRequestBuilder) {
    return NewGroupsItemTeamArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) Channels()(*GroupsItemTeamChannelsRequestBuilder) {
    return NewGroupsItemTeamChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) ChannelsById(id string)(*GroupsItemTeamChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewGroupsItemTeamChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone provides operations to call the clone method.
func (m *GroupsItemTeamRequestBuilder) Clone()(*GroupsItemTeamCloneRequestBuilder) {
    return NewGroupsItemTeamCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *GroupsItemTeamRequestBuilder) CompleteMigration()(*GroupsItemTeamCompleteMigrationRequestBuilder) {
    return NewGroupsItemTeamCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsItemTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewGroupsItemTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemTeamRequestBuilder) {
    m := &GroupsItemTeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemTeamRequestBuilder instantiates a new TeamRequestBuilder and sets the default values.
func NewGroupsItemTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemTeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property team for groups
func (m *GroupsItemTeamRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemTeamRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the team associated with this group.
func (m *GroupsItemTeamRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemTeamRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation create a new team from a group. In order to create a team, the group must have a least one owner. If the group was created less than 15 minutes ago, it's possible for the Create team call to fail with a 404 error code due to replication delays.The recommended pattern is to retry the Create team call three times, with a 10 second delay between calls.
func (m *GroupsItemTeamRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *GroupsItemTeamRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property team for groups
func (m *GroupsItemTeamRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemTeamRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the team associated with this group.
func (m *GroupsItemTeamRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemTeamRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
func (m *GroupsItemTeamRequestBuilder) Group()(*GroupsItemTeamGroupRequestBuilder) {
    return NewGroupsItemTeamGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) IncomingChannels()(*GroupsItemTeamIncomingChannelsRequestBuilder) {
    return NewGroupsItemTeamIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) IncomingChannelsById(id string)(*GroupsItemTeamIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewGroupsItemTeamIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) InstalledApps()(*GroupsItemTeamInstalledAppsRequestBuilder) {
    return NewGroupsItemTeamInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) InstalledAppsById(id string)(*GroupsItemTeamInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewGroupsItemTeamInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) Members()(*GroupsItemTeamMembersRequestBuilder) {
    return NewGroupsItemTeamMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) MembersById(id string)(*GroupsItemTeamMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewGroupsItemTeamMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) Operations()(*GroupsItemTeamOperationsRequestBuilder) {
    return NewGroupsItemTeamOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) OperationsById(id string)(*GroupsItemTeamOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewGroupsItemTeamOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) Owners()(*GroupsItemTeamOwnersRequestBuilder) {
    return NewGroupsItemTeamOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) OwnersById(id string)(*GroupsItemTeamOwnersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewGroupsItemTeamOwnersUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch create a new team from a group. In order to create a team, the group must have a least one owner. If the group was created less than 15 minutes ago, it's possible for the Create team call to fail with a 404 error code due to replication delays.The recommended pattern is to retry the Create team call three times, with a 10 second delay between calls.
func (m *GroupsItemTeamRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *GroupsItemTeamRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
func (m *GroupsItemTeamRequestBuilder) PermissionGrants()(*GroupsItemTeamPermissionGrantsRequestBuilder) {
    return NewGroupsItemTeamPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) PermissionGrantsById(id string)(*GroupsItemTeamPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return NewGroupsItemTeamPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) Photo()(*GroupsItemTeamPhotoRequestBuilder) {
    return NewGroupsItemTeamPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) PrimaryChannel()(*GroupsItemTeamPrimaryChannelRequestBuilder) {
    return NewGroupsItemTeamPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) Schedule()(*GroupsItemTeamScheduleRequestBuilder) {
    return NewGroupsItemTeamScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *GroupsItemTeamRequestBuilder) SendActivityNotification()(*GroupsItemTeamSendActivityNotificationRequestBuilder) {
    return NewGroupsItemTeamSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) Tags()(*GroupsItemTeamTagsRequestBuilder) {
    return NewGroupsItemTeamTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) TagsById(id string)(*GroupsItemTeamTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewGroupsItemTeamTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) Template()(*GroupsItemTeamTemplateRequestBuilder) {
    return NewGroupsItemTeamTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplateDefinition provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
func (m *GroupsItemTeamRequestBuilder) TemplateDefinition()(*GroupsItemTeamTemplateDefinitionRequestBuilder) {
    return NewGroupsItemTeamTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive provides operations to call the unarchive method.
func (m *GroupsItemTeamRequestBuilder) Unarchive()(*GroupsItemTeamUnarchiveRequestBuilder) {
    return NewGroupsItemTeamUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
