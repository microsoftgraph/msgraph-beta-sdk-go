package appcatalogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder provides operations to manage the media for the appCatalogs entity.
type TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderInternal instantiates a new TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder and sets the default values.
func NewTeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) {
    m := &TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/appCatalogs/teamsApps/{teamsApp%2Did}/appDefinitions/{teamsAppDefinition%2Did}/colorIcon/hostedContent/$value", pathParameters),
    }
    return m
}
// NewTeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder instantiates a new TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder and sets the default values.
func NewTeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the contents of the app icon if the icon is hosted within the Teams infrastructure.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the contents of the app icon if the icon is hosted within the Teams infrastructure.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderPutRequestConfiguration)([]byte, error) {
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
func (m *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the contents of the app icon if the icon is hosted within the Teams infrastructure.
// returns a *RequestInformation when successful
func (m *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder when successful
func (m *TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) WithUrl(rawUrl string)(*TeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder) {
    return NewTeamsAppsItemAppDefinitionsItemColorIconHostedContentValueContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
