package appcatalogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder provides operations to manage the media for the appCatalogs entity.
type TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderInternal instantiates a new TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder and sets the default values.
func NewTeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) {
    m := &TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/appCatalogs/teamsApps/{teamsApp%2Did}/appDefinitions/{teamsAppDefinition%2Did}/outlineIcon/hostedContent/$value", pathParameters),
    }
    return m
}
// NewTeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder instantiates a new TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder and sets the default values.
func NewTeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the contents of the app icon if the icon is hosted within the Teams infrastructure.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamworkhostedcontent-get?view=graph-rest-beta
func (m *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put the contents of the app icon if the icon is hosted within the Teams infrastructure.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToDeleteRequestInformation the contents of the app icon if the icon is hosted within the Teams infrastructure.
// returns a *RequestInformation when successful
func (m *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the contents of the app icon if the icon is hosted within the Teams infrastructure.
// returns a *RequestInformation when successful
func (m *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) WithUrl(rawUrl string)(*TeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsItemOutlineIconHostedContentValueContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
