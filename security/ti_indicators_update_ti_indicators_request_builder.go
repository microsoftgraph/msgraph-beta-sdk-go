package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TiIndicatorsUpdateTiIndicatorsRequestBuilder provides operations to call the updateTiIndicators method.
type TiIndicatorsUpdateTiIndicatorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TiIndicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiIndicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTiIndicatorsUpdateTiIndicatorsRequestBuilderInternal instantiates a new UpdateTiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsUpdateTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsUpdateTiIndicatorsRequestBuilder) {
    m := &TiIndicatorsUpdateTiIndicatorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/tiIndicators/updateTiIndicators", pathParameters),
    }
    return m
}
// NewTiIndicatorsUpdateTiIndicatorsRequestBuilder instantiates a new UpdateTiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsUpdateTiIndicatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsUpdateTiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiIndicatorsUpdateTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update multiple threat intelligence (TI) indicators in one request instead of multiple requests. This API is supported in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsUpdateTiIndicatorsPostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-updatetiindicators?view=graph-rest-1.0
func (m *TiIndicatorsUpdateTiIndicatorsRequestBuilder) Post(ctx context.Context, body TiIndicatorsUpdateTiIndicatorsPostRequestBodyable, requestConfiguration *TiIndicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration)(TiIndicatorsUpdateTiIndicatorsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiIndicatorsUpdateTiIndicatorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiIndicatorsUpdateTiIndicatorsResponseable), nil
}
// PostAsUpdateTiIndicatorsPostResponse update multiple threat intelligence (TI) indicators in one request instead of multiple requests. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-updatetiindicators?view=graph-rest-1.0
func (m *TiIndicatorsUpdateTiIndicatorsRequestBuilder) PostAsUpdateTiIndicatorsPostResponse(ctx context.Context, body TiIndicatorsUpdateTiIndicatorsPostRequestBodyable, requestConfiguration *TiIndicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration)(TiIndicatorsUpdateTiIndicatorsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiIndicatorsUpdateTiIndicatorsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiIndicatorsUpdateTiIndicatorsPostResponseable), nil
}
// ToPostRequestInformation update multiple threat intelligence (TI) indicators in one request instead of multiple requests. This API is supported in the following national cloud deployments.
func (m *TiIndicatorsUpdateTiIndicatorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TiIndicatorsUpdateTiIndicatorsPostRequestBodyable, requestConfiguration *TiIndicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *TiIndicatorsUpdateTiIndicatorsRequestBuilder) WithUrl(rawUrl string)(*TiIndicatorsUpdateTiIndicatorsRequestBuilder) {
    return NewTiIndicatorsUpdateTiIndicatorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
