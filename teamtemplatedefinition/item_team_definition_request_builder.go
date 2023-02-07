package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionRequestBuilder provides operations to manage the teamDefinition property of the microsoft.graph.teamTemplateDefinition entity.
type ItemTeamDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamDefinitionRequestBuilderGetQueryParameters get the properties of the team associated with a teamTemplateDefinition object.
type ItemTeamDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) AllChannels()(*ItemTeamDefinitionAllChannelsRequestBuilder) {
    return NewItemTeamDefinitionAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) AllChannelsById(id string)(*ItemTeamDefinitionAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemTeamDefinitionAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) Channels()(*ItemTeamDefinitionChannelsRequestBuilder) {
    return NewItemTeamDefinitionChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) ChannelsById(id string)(*ItemTeamDefinitionChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemTeamDefinitionChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemTeamDefinitionRequestBuilderInternal instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewItemTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionRequestBuilder) {
    m := &ItemTeamDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamDefinitionRequestBuilder instantiates a new TeamDefinitionRequestBuilder and sets the default values.
func NewItemTeamDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property teamDefinition for teamTemplateDefinition
func (m *ItemTeamDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemTeamDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
func (m *ItemTeamDefinitionRequestBuilder) Group()(*ItemTeamDefinitionGroupRequestBuilder) {
    return NewItemTeamDefinitionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) IncomingChannels()(*ItemTeamDefinitionIncomingChannelsRequestBuilder) {
    return NewItemTeamDefinitionIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) IncomingChannelsById(id string)(*ItemTeamDefinitionIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemTeamDefinitionIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) InstalledApps()(*ItemTeamDefinitionInstalledAppsRequestBuilder) {
    return NewItemTeamDefinitionInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) InstalledAppsById(id string)(*ItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) Members()(*ItemTeamDefinitionMembersRequestBuilder) {
    return NewItemTeamDefinitionMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) MembersById(id string)(*ItemTeamDefinitionMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemTeamDefinitionMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphArchive provides operations to call the archive method.
func (m *ItemTeamDefinitionRequestBuilder) MicrosoftGraphArchive()(*ItemTeamDefinitionMicrosoftGraphArchiveRequestBuilder) {
    return NewItemTeamDefinitionMicrosoftGraphArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphClone provides operations to call the clone method.
func (m *ItemTeamDefinitionRequestBuilder) MicrosoftGraphClone()(*ItemTeamDefinitionMicrosoftGraphCloneRequestBuilder) {
    return NewItemTeamDefinitionMicrosoftGraphCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCompleteMigration provides operations to call the completeMigration method.
func (m *ItemTeamDefinitionRequestBuilder) MicrosoftGraphCompleteMigration()(*ItemTeamDefinitionMicrosoftGraphCompleteMigrationRequestBuilder) {
    return NewItemTeamDefinitionMicrosoftGraphCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendActivityNotification provides operations to call the sendActivityNotification method.
func (m *ItemTeamDefinitionRequestBuilder) MicrosoftGraphSendActivityNotification()(*ItemTeamDefinitionMicrosoftGraphSendActivityNotificationRequestBuilder) {
    return NewItemTeamDefinitionMicrosoftGraphSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnarchive provides operations to call the unarchive method.
func (m *ItemTeamDefinitionRequestBuilder) MicrosoftGraphUnarchive()(*ItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilder) {
    return NewItemTeamDefinitionMicrosoftGraphUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) Operations()(*ItemTeamDefinitionOperationsRequestBuilder) {
    return NewItemTeamDefinitionOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) OperationsById(id string)(*ItemTeamDefinitionOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewItemTeamDefinitionOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) Owners()(*ItemTeamDefinitionOwnersRequestBuilder) {
    return NewItemTeamDefinitionOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OwnersById provides operations to manage the owners property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) OwnersById(id string)(*ItemTeamDefinitionOwnersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewItemTeamDefinitionOwnersUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property teamDefinition in teamTemplateDefinition
func (m *ItemTeamDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *ItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
func (m *ItemTeamDefinitionRequestBuilder) PermissionGrants()(*ItemTeamDefinitionPermissionGrantsRequestBuilder) {
    return NewItemTeamDefinitionPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PermissionGrantsById provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) PermissionGrantsById(id string)(*ItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return NewItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) Photo()(*ItemTeamDefinitionPhotoRequestBuilder) {
    return NewItemTeamDefinitionPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) PrimaryChannel()(*ItemTeamDefinitionPrimaryChannelRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) Schedule()(*ItemTeamDefinitionScheduleRequestBuilder) {
    return NewItemTeamDefinitionScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) Tags()(*ItemTeamDefinitionTagsRequestBuilder) {
    return NewItemTeamDefinitionTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) TagsById(id string)(*ItemTeamDefinitionTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewItemTeamDefinitionTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) Template()(*ItemTeamDefinitionTemplateRequestBuilder) {
    return NewItemTeamDefinitionTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TemplateDefinition provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
func (m *ItemTeamDefinitionRequestBuilder) TemplateDefinition()(*ItemTeamDefinitionTemplateDefinitionRequestBuilder) {
    return NewItemTeamDefinitionTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property teamDefinition for teamTemplateDefinition
func (m *ItemTeamDefinitionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemTeamDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property teamDefinition in teamTemplateDefinition
func (m *ItemTeamDefinitionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *ItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
