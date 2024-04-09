package appcatalogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder provides operations to manage the dashboardCards property of the microsoft.graph.teamsAppDefinition entity.
type TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderGetQueryParameters dashboard cards specified in the Teams app manifest.
type TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderGetQueryParameters
}
// TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTeamsAppDashboardCardDefinitionId provides operations to manage the dashboardCards property of the microsoft.graph.teamsAppDefinition entity.
// returns a *TeamsAppsItemAppDefinitionsItemDashboardCardsTeamsAppDashboardCardDefinitionItemRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) ByTeamsAppDashboardCardDefinitionId(teamsAppDashboardCardDefinitionId string)(*TeamsAppsItemAppDefinitionsItemDashboardCardsTeamsAppDashboardCardDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if teamsAppDashboardCardDefinitionId != "" {
        urlTplParams["teamsAppDashboardCardDefinition%2Did"] = teamsAppDashboardCardDefinitionId
    }
    return NewTeamsAppsItemAppDefinitionsItemDashboardCardsTeamsAppDashboardCardDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderInternal instantiates a new TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder and sets the default values.
func NewTeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) {
    m := &TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/appCatalogs/teamsApps/{teamsApp%2Did}/appDefinitions/{teamsAppDefinition%2Did}/dashboardCards{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder instantiates a new TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder and sets the default values.
func NewTeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TeamsAppsItemAppDefinitionsItemDashboardCardsCountRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) Count()(*TeamsAppsItemAppDefinitionsItemDashboardCardsCountRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsItemDashboardCardsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get dashboard cards specified in the Teams app manifest.
// returns a TeamsAppDashboardCardDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDashboardCardDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamsAppDashboardCardDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDashboardCardDefinitionCollectionResponseable), nil
}
// Post create new navigation property to dashboardCards for appCatalogs
// returns a TeamsAppDashboardCardDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDashboardCardDefinitionable, requestConfiguration *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDashboardCardDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamsAppDashboardCardDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDashboardCardDefinitionable), nil
}
// ToGetRequestInformation dashboard cards specified in the Teams app manifest.
// returns a *RequestInformation when successful
func (m *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dashboardCards for appCatalogs
// returns a *RequestInformation when successful
func (m *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamsAppDashboardCardDefinitionable, requestConfiguration *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) WithUrl(rawUrl string)(*TeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsItemDashboardCardsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
