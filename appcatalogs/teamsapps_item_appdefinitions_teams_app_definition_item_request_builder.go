package appcatalogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder provides operations to manage the appDefinitions property of the microsoft.graph.teamsApp entity.
type TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderGetQueryParameters the details for each version of the app.
type TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderGetQueryParameters
}
// TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Bot provides operations to manage the bot property of the microsoft.graph.teamsAppDefinition entity.
// returns a *TeamsappsItemAppdefinitionsItemBotRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) Bot()(*TeamsappsItemAppdefinitionsItemBotRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsItemBotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColorIcon provides operations to manage the colorIcon property of the microsoft.graph.teamsAppDefinition entity.
// returns a *TeamsappsItemAppdefinitionsItemColoriconColorIconRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) ColorIcon()(*TeamsappsItemAppdefinitionsItemColoriconColorIconRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsItemColoriconColorIconRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderInternal instantiates a new TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder and sets the default values.
func NewTeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) {
    m := &TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/appCatalogs/teamsApps/{teamsApp%2Did}/appDefinitions/{teamsAppDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder instantiates a new TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder and sets the default values.
func NewTeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DashboardCards provides operations to manage the dashboardCards property of the microsoft.graph.teamsAppDefinition entity.
// returns a *TeamsappsItemAppdefinitionsItemDashboardcardsDashboardCardsRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) DashboardCards()(*TeamsappsItemAppdefinitionsItemDashboardcardsDashboardCardsRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsItemDashboardcardsDashboardCardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property appDefinitions for appCatalogs
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, error) {
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
// returns a *TeamsappsItemAppdefinitionsItemOutlineiconOutlineIconRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) OutlineIcon()(*TeamsappsItemAppdefinitionsItemOutlineiconOutlineIconRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsItemOutlineiconOutlineIconRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch publish an app to the Microsoft Teams app catalog.Specifically, this API publishes the app to your organization's catalog (the tenant app catalog);the created resource has a distributionMethod property value of organization. The requiresReview property allows any user to submit an app for review by an administrator. Admins can approve or reject these apps via this API or the Microsoft Teams admin center.
// returns a TeamsAppDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamsapp-publish?view=graph-rest-beta
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, requestConfiguration *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, error) {
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
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation publish an app to the Microsoft Teams app catalog.Specifically, this API publishes the app to your organization's catalog (the tenant app catalog);the created resource has a distributionMethod property value of organization. The requiresReview property allows any user to submit an app for review by an administrator. Admins can approve or reject these apps via this API or the Microsoft Teams admin center.
// returns a *RequestInformation when successful
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDefinitionable, requestConfiguration *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
