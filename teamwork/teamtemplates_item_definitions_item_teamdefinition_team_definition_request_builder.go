package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder provides operations to manage the teamDefinition property of the microsoft.graph.teamTemplateDefinition entity.
type TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderGetQueryParameters collection of channel objects. A channel represents a topic, and therefore a logical isolation of discussion, within a team.
type TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderGetQueryParameters
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionAllchannelsAllChannelsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) AllChannels()(*TeamtemplatesItemDefinitionsItemTeamdefinitionAllchannelsAllChannelsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionAllchannelsAllChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Archive provides operations to call the archive method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionArchiveRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Archive()(*TeamtemplatesItemDefinitionsItemTeamdefinitionArchiveRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Channels()(*TeamtemplatesItemDefinitionsItemTeamdefinitionChannelsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Clone provides operations to call the clone method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionCloneRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Clone()(*TeamtemplatesItemDefinitionsItemTeamdefinitionCloneRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionCloneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompleteMigration provides operations to call the completeMigration method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) CompleteMigration()(*TeamtemplatesItemDefinitionsItemTeamdefinitionCompletemigrationCompleteMigrationRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionCompletemigrationCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderInternal instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property teamDefinition for teamwork
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of channel objects. A channel represents a topic, and therefore a logical isolation of discussion, within a team.
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionGroupRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Group()(*TeamtemplatesItemDefinitionsItemTeamdefinitionGroupRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionIncomingchannelsIncomingChannelsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) IncomingChannels()(*TeamtemplatesItemDefinitionsItemTeamdefinitionIncomingchannelsIncomingChannelsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionIncomingchannelsIncomingChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionInstalledappsInstalledAppsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) InstalledApps()(*TeamtemplatesItemDefinitionsItemTeamdefinitionInstalledappsInstalledAppsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionInstalledappsInstalledAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionMembersRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Members()(*TeamtemplatesItemDefinitionsItemTeamdefinitionMembersRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionOperationsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Operations()(*TeamtemplatesItemDefinitionsItemTeamdefinitionOperationsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionOwnersRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Owners()(*TeamtemplatesItemDefinitionsItemTeamdefinitionOwnersRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property teamDefinition in teamwork
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable), nil
}
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPermissiongrantsPermissionGrantsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) PermissionGrants()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPermissiongrantsPermissionGrantsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPermissiongrantsPermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPhotoRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Photo()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPhotoRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPhotoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) PrimaryChannel()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionScheduleRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Schedule()(*TeamtemplatesItemDefinitionsItemTeamdefinitionScheduleRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionSendactivitynotificationSendActivityNotificationRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) SendActivityNotification()(*TeamtemplatesItemDefinitionsItemTeamdefinitionSendactivitynotificationSendActivityNotificationRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionSendactivitynotificationSendActivityNotificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionTagsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Tags()(*TeamtemplatesItemDefinitionsItemTeamdefinitionTagsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplateRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Template()(*TeamtemplatesItemDefinitionsItemTeamdefinitionTemplateRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TemplateDefinition provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) TemplateDefinition()(*TeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property teamDefinition for teamwork
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of channel objects. A channel represents a topic, and therefore a logical isolation of discussion, within a team.
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property teamDefinition in teamwork
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Unarchive provides operations to call the unarchive method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionUnarchiveRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) Unarchive()(*TeamtemplatesItemDefinitionsItemTeamdefinitionUnarchiveRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) WithUrl(rawUrl string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionTeamDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
