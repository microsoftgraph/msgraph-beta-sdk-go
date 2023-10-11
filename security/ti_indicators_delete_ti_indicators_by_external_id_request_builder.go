package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder provides operations to call the deleteTiIndicatorsByExternalId method.
type TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderInternal instantiates a new DeleteTiIndicatorsByExternalIdRequestBuilder and sets the default values.
func NewTiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder) {
    m := &TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/tiIndicators/deleteTiIndicatorsByExternalId", pathParameters),
    }
    return m
}
// NewTiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder instantiates a new DeleteTiIndicatorsByExternalIdRequestBuilder and sets the default values.
func NewTiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post delete multiple threat intelligence (TI) indicators in one request instead of multiple requests, when the request contains external IDs instead of IDs. This API is available in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsDeleteTiIndicatorsByExternalIdPostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-deletetiindicatorsbyexternalid?view=graph-rest-1.0
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder) Post(ctx context.Context, body TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBodyable, requestConfiguration *TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration)(TiIndicatorsDeleteTiIndicatorsByExternalIdResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiIndicatorsDeleteTiIndicatorsByExternalIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiIndicatorsDeleteTiIndicatorsByExternalIdResponseable), nil
}
// PostAsDeleteTiIndicatorsByExternalIdPostResponse delete multiple threat intelligence (TI) indicators in one request instead of multiple requests, when the request contains external IDs instead of IDs. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-deletetiindicatorsbyexternalid?view=graph-rest-1.0
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder) PostAsDeleteTiIndicatorsByExternalIdPostResponse(ctx context.Context, body TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBodyable, requestConfiguration *TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration)(TiIndicatorsDeleteTiIndicatorsByExternalIdPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiIndicatorsDeleteTiIndicatorsByExternalIdPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiIndicatorsDeleteTiIndicatorsByExternalIdPostResponseable), nil
}
// ToPostRequestInformation delete multiple threat intelligence (TI) indicators in one request instead of multiple requests, when the request contains external IDs instead of IDs. This API is available in the following national cloud deployments.
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBodyable, requestConfiguration *TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder) WithUrl(rawUrl string)(*TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder) {
    return NewTiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
