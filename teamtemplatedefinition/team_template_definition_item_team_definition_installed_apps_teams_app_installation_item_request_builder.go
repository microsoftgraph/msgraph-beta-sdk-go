package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder provides operations to manage the installedApps property of the microsoft.graph.team entity.
type TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderGetQueryParameters the apps installed in this team.
type TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderGetQueryParameters
}
// TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderInternal instantiates a new TeamsAppInstallationItemRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    m := &TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/installedApps/{teamsAppInstallation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder instantiates a new TeamsAppInstallationItemRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property installedApps for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the apps installed in this team.
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property installedApps in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppInstallationable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property installedApps for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the apps installed in this team.
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppInstallationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamsAppInstallationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppInstallationable), nil
}
// Patch update the navigation property installedApps in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppInstallationable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppInstallationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamsAppInstallationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppInstallationable), nil
}
// TeamsApp provides operations to manage the teamsApp property of the microsoft.graph.teamsAppInstallation entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) TeamsApp()(*TeamTemplateDefinitionItemTeamDefinitionInstalledAppsItemTeamsAppRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsItemTeamsAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamsAppDefinition provides operations to manage the teamsAppDefinition property of the microsoft.graph.teamsAppInstallation entity.
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) TeamsAppDefinition()(*TeamTemplateDefinitionItemTeamDefinitionInstalledAppsItemTeamsAppDefinitionRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsItemTeamsAppDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Upgrade provides operations to call the upgrade method.
func (m *TeamTemplateDefinitionItemTeamDefinitionInstalledAppsTeamsAppInstallationItemRequestBuilder) Upgrade()(*TeamTemplateDefinitionItemTeamDefinitionInstalledAppsItemUpgradeRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionInstalledAppsItemUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
