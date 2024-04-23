package appcatalogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder provides operations to manage the appDefinitions property of the microsoft.graph.teamsApp entity.
type TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderGetQueryParameters the details for each version of the app.
type TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderGetQueryParameters
}
// TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bot provides operations to manage the bot property of the microsoft.graph.teamsAppDefinition entity.
// returns a *TeamsAppsItemAppDefinitionsItemBotRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) Bot()(*TeamsAppsItemAppDefinitionsItemBotRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsItemBotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColorIcon provides operations to manage the colorIcon property of the microsoft.graph.teamsAppDefinition entity.
// returns a *TeamsAppsItemAppDefinitionsItemColorIconRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) ColorIcon()(*TeamsAppsItemAppDefinitionsItemColorIconRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsItemColorIconRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderInternal instantiates a new TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder and sets the default values.
func NewTeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) {
    m := &TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/appCatalogs/teamsApps/{teamsApp%2Did}/appDefinitions/{teamsAppDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder instantiates a new TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder and sets the default values.
func NewTeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DashboardCards provides operations to manage the dashboardCards property of the microsoft.graph.teamsAppDefinition entity.
// returns a *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) DashboardCards()(*TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property appDefinitions for appCatalogs
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the details for each version of the app.
// returns a TeamsAppDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamsAppDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable), nil
}
// OutlineIcon provides operations to manage the outlineIcon property of the microsoft.graph.teamsAppDefinition entity.
// returns a *TeamsAppsItemAppDefinitionsItemOutlineIconRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) OutlineIcon()(*TeamsAppsItemAppDefinitionsItemOutlineIconRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsItemOutlineIconRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property appDefinitions in appCatalogs
// returns a TeamsAppDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, requestConfiguration *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamsAppDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable), nil
}
// ToDeleteRequestInformation delete navigation property appDefinitions for appCatalogs
// returns a *RequestInformation when successful
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the details for each version of the app.
// returns a *RequestInformation when successful
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property appDefinitions in appCatalogs
// returns a *RequestInformation when successful
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, requestConfiguration *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*TeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsTeamsAppDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
