package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionTeamDefinitionRequestBuilder provides operations to manage the teamDefinition property of the microsoft.graph.teamTemplateDefinition entity.
type ItemTeamdefinitionTeamDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionTeamDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionTeamDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamdefinitionTeamDefinitionRequestBuilderGetQueryParameters collection of channel objects. A channel represents a topic, and therefore a logical isolation of discussion, within a team.
type ItemTeamdefinitionTeamDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamdefinitionTeamDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionTeamDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamdefinitionTeamDefinitionRequestBuilderGetQueryParameters
}
// ItemTeamdefinitionTeamDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionTeamDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionAllchannelsAllChannelsRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) AllChannels()(*ItemTeamdefinitionAllchannelsAllChannelsRequestBuilder) {
    return NewItemTeamdefinitionAllchannelsAllChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Archive provides operations to call the archive method.
// returns a *ItemTeamdefinitionArchiveRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Archive()(*ItemTeamdefinitionArchiveRequestBuilder) {
    return NewItemTeamdefinitionArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionChannelsRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Channels()(*ItemTeamdefinitionChannelsRequestBuilder) {
    return NewItemTeamdefinitionChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Clone provides operations to call the clone method.
// returns a *ItemTeamdefinitionCloneRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Clone()(*ItemTeamdefinitionCloneRequestBuilder) {
    return NewItemTeamdefinitionCloneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompleteMigration provides operations to call the completeMigration method.
// returns a *ItemTeamdefinitionCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) CompleteMigration()(*ItemTeamdefinitionCompletemigrationCompleteMigrationRequestBuilder) {
    return NewItemTeamdefinitionCompletemigrationCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamdefinitionTeamDefinitionRequestBuilderInternal instantiates a new ItemTeamdefinitionTeamDefinitionRequestBuilder and sets the default values.
func NewItemTeamdefinitionTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionTeamDefinitionRequestBuilder) {
    m := &ItemTeamdefinitionTeamDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionTeamDefinitionRequestBuilder instantiates a new ItemTeamdefinitionTeamDefinitionRequestBuilder and sets the default values.
func NewItemTeamdefinitionTeamDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionTeamDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionTeamDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property teamDefinition for teamTemplateDefinition
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamdefinitionTeamDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamdefinitionTeamDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
// returns a *ItemTeamdefinitionGroupRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Group()(*ItemTeamdefinitionGroupRequestBuilder) {
    return NewItemTeamdefinitionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionIncomingchannelsIncomingChannelsRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) IncomingChannels()(*ItemTeamdefinitionIncomingchannelsIncomingChannelsRequestBuilder) {
    return NewItemTeamdefinitionIncomingchannelsIncomingChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionInstalledappsInstalledAppsRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) InstalledApps()(*ItemTeamdefinitionInstalledappsInstalledAppsRequestBuilder) {
    return NewItemTeamdefinitionInstalledappsInstalledAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionMembersRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Members()(*ItemTeamdefinitionMembersRequestBuilder) {
    return NewItemTeamdefinitionMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionOperationsRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Operations()(*ItemTeamdefinitionOperationsRequestBuilder) {
    return NewItemTeamdefinitionOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionOwnersRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Owners()(*ItemTeamdefinitionOwnersRequestBuilder) {
    return NewItemTeamdefinitionOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property teamDefinition in teamTemplateDefinition
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *ItemTeamdefinitionTeamDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
// returns a *ItemTeamdefinitionPermissiongrantsPermissionGrantsRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) PermissionGrants()(*ItemTeamdefinitionPermissiongrantsPermissionGrantsRequestBuilder) {
    return NewItemTeamdefinitionPermissiongrantsPermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionPhotoRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Photo()(*ItemTeamdefinitionPhotoRequestBuilder) {
    return NewItemTeamdefinitionPhotoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) PrimaryChannel()(*ItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) {
    return NewItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionScheduleRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Schedule()(*ItemTeamdefinitionScheduleRequestBuilder) {
    return NewItemTeamdefinitionScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
// returns a *ItemTeamdefinitionSendactivitynotificationSendActivityNotificationRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) SendActivityNotification()(*ItemTeamdefinitionSendactivitynotificationSendActivityNotificationRequestBuilder) {
    return NewItemTeamdefinitionSendactivitynotificationSendActivityNotificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionTagsRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Tags()(*ItemTeamdefinitionTagsRequestBuilder) {
    return NewItemTeamdefinitionTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionTemplateRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Template()(*ItemTeamdefinitionTemplateRequestBuilder) {
    return NewItemTeamdefinitionTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TemplateDefinition provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
// returns a *ItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) TemplateDefinition()(*ItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilder) {
    return NewItemTeamdefinitionTemplatedefinitionTemplateDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property teamDefinition for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamdefinitionTeamDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamdefinitionTeamDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property teamDefinition in teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *ItemTeamdefinitionTeamDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamdefinitionUnarchiveRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) Unarchive()(*ItemTeamdefinitionUnarchiveRequestBuilder) {
    return NewItemTeamdefinitionUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamdefinitionTeamDefinitionRequestBuilder when successful
func (m *ItemTeamdefinitionTeamDefinitionRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionTeamDefinitionRequestBuilder) {
    return NewItemTeamdefinitionTeamDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
