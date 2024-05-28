package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderGetQueryParameters the general channel for the team.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderGetQueryParameters
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Archive provides operations to call the archive method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelArchiveRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) Archive()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelArchiveRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompleteMigration provides operations to call the completeMigration method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) CompleteMigration()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelCompletemigrationCompleteMigrationRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelCompletemigrationCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderInternal instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property primaryChannel for teamwork
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderDeleteRequestConfiguration)(error) {
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
// DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName provides operations to call the doesUserHaveAccess method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelFilesfolderFilesFolderRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) FilesFolder()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelFilesfolderFilesFolderRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelFilesfolderFilesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the general channel for the team.
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.channel entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) Members()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMessagesRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) Messages()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMessagesRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property primaryChannel in teamwork
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// ProvisionEmail provides operations to call the provisionEmail method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) ProvisionEmail()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelProvisionemailProvisionEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveEmail provides operations to call the removeEmail method.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelRemoveemailRemoveEmailRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) RemoveEmail()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelRemoveemailRemoveEmailRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelRemoveemailRemoveEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsSharedWithTeamsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) SharedWithTeams()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsSharedWithTeamsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsSharedWithTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelTabsRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) Tabs()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelTabsRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelTabsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property primaryChannel for teamwork
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the general channel for the team.
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property primaryChannel in teamwork
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelUnarchiveRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) Unarchive()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelUnarchiveRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) WithUrl(rawUrl string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelPrimaryChannelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
