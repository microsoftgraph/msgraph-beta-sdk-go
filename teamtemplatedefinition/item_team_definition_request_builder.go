package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionRequestBuilder provides operations to manage the teamDefinition property of the microsoft.graph.teamTemplateDefinition entity.
type ItemTeamDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// returns a *ItemTeamDefinitionAllChannelsRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) AllChannels()(*ItemTeamDefinitionAllChannelsRequestBuilder) {
    return NewItemTeamDefinitionAllChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Archive provides operations to call the archive method.
// returns a *ItemTeamDefinitionArchiveRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Archive()(*ItemTeamDefinitionArchiveRequestBuilder) {
    return NewItemTeamDefinitionArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionChannelsRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Channels()(*ItemTeamDefinitionChannelsRequestBuilder) {
    return NewItemTeamDefinitionChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Clone provides operations to call the clone method.
// returns a *ItemTeamDefinitionCloneRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Clone()(*ItemTeamDefinitionCloneRequestBuilder) {
    return NewItemTeamDefinitionCloneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompleteMigration provides operations to call the completeMigration method.
// returns a *ItemTeamDefinitionCompleteMigrationRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) CompleteMigration()(*ItemTeamDefinitionCompleteMigrationRequestBuilder) {
    return NewItemTeamDefinitionCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamDefinitionRequestBuilderInternal instantiates a new ItemTeamDefinitionRequestBuilder and sets the default values.
func NewItemTeamDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionRequestBuilder) {
    m := &ItemTeamDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionRequestBuilder instantiates a new ItemTeamDefinitionRequestBuilder and sets the default values.
func NewItemTeamDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property teamDefinition for teamTemplateDefinition
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties of the team associated with a teamTemplateDefinition object.
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamtemplatedefinition-get-teamdefinition?view=graph-rest-1.0
func (m *ItemTeamDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
// returns a *ItemTeamDefinitionGroupRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Group()(*ItemTeamDefinitionGroupRequestBuilder) {
    return NewItemTeamDefinitionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionIncomingChannelsRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) IncomingChannels()(*ItemTeamDefinitionIncomingChannelsRequestBuilder) {
    return NewItemTeamDefinitionIncomingChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionInstalledAppsRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) InstalledApps()(*ItemTeamDefinitionInstalledAppsRequestBuilder) {
    return NewItemTeamDefinitionInstalledAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionMembersRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Members()(*ItemTeamDefinitionMembersRequestBuilder) {
    return NewItemTeamDefinitionMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionOperationsRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Operations()(*ItemTeamDefinitionOperationsRequestBuilder) {
    return NewItemTeamDefinitionOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Owners provides operations to manage the owners property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionOwnersRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Owners()(*ItemTeamDefinitionOwnersRequestBuilder) {
    return NewItemTeamDefinitionOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property teamDefinition in teamTemplateDefinition
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *ItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, error) {
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
// returns a *ItemTeamDefinitionPermissionGrantsRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) PermissionGrants()(*ItemTeamDefinitionPermissionGrantsRequestBuilder) {
    return NewItemTeamDefinitionPermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionPhotoRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Photo()(*ItemTeamDefinitionPhotoRequestBuilder) {
    return NewItemTeamDefinitionPhotoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionPrimaryChannelRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) PrimaryChannel()(*ItemTeamDefinitionPrimaryChannelRequestBuilder) {
    return NewItemTeamDefinitionPrimaryChannelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionScheduleRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Schedule()(*ItemTeamDefinitionScheduleRequestBuilder) {
    return NewItemTeamDefinitionScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
// returns a *ItemTeamDefinitionSendActivityNotificationRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) SendActivityNotification()(*ItemTeamDefinitionSendActivityNotificationRequestBuilder) {
    return NewItemTeamDefinitionSendActivityNotificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionTagsRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Tags()(*ItemTeamDefinitionTagsRequestBuilder) {
    return NewItemTeamDefinitionTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionTemplateRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Template()(*ItemTeamDefinitionTemplateRequestBuilder) {
    return NewItemTeamDefinitionTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TemplateDefinition provides operations to manage the templateDefinition property of the microsoft.graph.team entity.
// returns a *ItemTeamDefinitionTemplateDefinitionRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) TemplateDefinition()(*ItemTeamDefinitionTemplateDefinitionRequestBuilder) {
    return NewItemTeamDefinitionTemplateDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property teamDefinition for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties of the team associated with a teamTemplateDefinition object.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemTeamDefinitionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Teamable, requestConfiguration *ItemTeamDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamDefinitionUnarchiveRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) Unarchive()(*ItemTeamDefinitionUnarchiveRequestBuilder) {
    return NewItemTeamDefinitionUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamDefinitionRequestBuilder when successful
func (m *ItemTeamDefinitionRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionRequestBuilder) {
    return NewItemTeamDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
