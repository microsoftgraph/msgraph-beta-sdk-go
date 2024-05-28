package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedsimactivationcodepoolsItemAssignRequestBuilder provides operations to call the assign method.
type EmbeddedsimactivationcodepoolsItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedsimactivationcodepoolsItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEmbeddedsimactivationcodepoolsItemAssignRequestBuilderInternal instantiates a new EmbeddedsimactivationcodepoolsItemAssignRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsItemAssignRequestBuilder) {
    m := &EmbeddedsimactivationcodepoolsItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}/assign", pathParameters),
    }
    return m
}
// NewEmbeddedsimactivationcodepoolsItemAssignRequestBuilder instantiates a new EmbeddedsimactivationcodepoolsItemAssignRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedsimactivationcodepoolsItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
// Deprecated: This method is obsolete. Use PostAsAssignPostResponse instead.
// returns a EmbeddedsimactivationcodepoolsItemAssignResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemAssignRequestBuilder) Post(ctx context.Context, body EmbeddedsimactivationcodepoolsItemAssignPostRequestBodyable, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignRequestBuilderPostRequestConfiguration)(EmbeddedsimactivationcodepoolsItemAssignResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmbeddedsimactivationcodepoolsItemAssignResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EmbeddedsimactivationcodepoolsItemAssignResponseable), nil
}
// PostAsAssignPostResponse invoke action assign
// returns a EmbeddedsimactivationcodepoolsItemAssignPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemAssignRequestBuilder) PostAsAssignPostResponse(ctx context.Context, body EmbeddedsimactivationcodepoolsItemAssignPostRequestBodyable, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignRequestBuilderPostRequestConfiguration)(EmbeddedsimactivationcodepoolsItemAssignPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmbeddedsimactivationcodepoolsItemAssignPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EmbeddedsimactivationcodepoolsItemAssignPostResponseable), nil
}
// ToPostRequestInformation invoke action assign
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmbeddedsimactivationcodepoolsItemAssignPostRequestBodyable, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmbeddedsimactivationcodepoolsItemAssignRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsItemAssignRequestBuilder) WithUrl(rawUrl string)(*EmbeddedsimactivationcodepoolsItemAssignRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
