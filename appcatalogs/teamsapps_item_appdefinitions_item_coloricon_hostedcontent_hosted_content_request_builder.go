package appcatalogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder provides operations to manage the hostedContent property of the microsoft.graph.teamsAppIcon entity.
type TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderGetQueryParameters retrieve the hosted content in an app's icon.
type TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderGetQueryParameters
}
// TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderInternal instantiates a new TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder and sets the default values.
func NewTeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) {
    m := &TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/appCatalogs/teamsApps/{teamsApp%2Did}/appDefinitions/{teamsAppDefinition%2Did}/colorIcon/hostedContent{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder instantiates a new TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder and sets the default values.
func NewTeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the appCatalogs entity.
// returns a *TeamsappsItemAppdefinitionsItemColoriconHostedcontentValueContentRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) Content()(*TeamsappsItemAppdefinitionsItemColoriconHostedcontentValueContentRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsItemColoriconHostedcontentValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property hostedContent for appCatalogs
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the hosted content in an app's icon.
// returns a TeamworkHostedContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamworkhostedcontent-get?view=graph-rest-beta
func (m *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkHostedContentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkHostedContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkHostedContentable), nil
}
// Patch update the navigation property hostedContent in appCatalogs
// returns a TeamworkHostedContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkHostedContentable, requestConfiguration *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkHostedContentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkHostedContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkHostedContentable), nil
}
// ToDeleteRequestInformation delete navigation property hostedContent for appCatalogs
// returns a *RequestInformation when successful
func (m *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the hosted content in an app's icon.
// returns a *RequestInformation when successful
func (m *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hostedContent in appCatalogs
// returns a *RequestInformation when successful
func (m *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkHostedContentable, requestConfiguration *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) WithUrl(rawUrl string)(*TeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsItemColoriconHostedcontentHostedContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
