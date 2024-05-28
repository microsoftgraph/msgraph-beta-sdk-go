package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder provides operations to call the submitTiIndicators method.
type TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilderInternal instantiates a new TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder and sets the default values.
func NewTiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder) {
    m := &TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/tiIndicators/submitTiIndicators", pathParameters),
    }
    return m
}
// NewTiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder instantiates a new TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder and sets the default values.
func NewTiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: This method is obsolete. Use PostAsSubmitTiIndicatorsPostResponse instead.
// returns a TiindicatorsSubmittiindicatorsSubmitTiIndicatorsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-submittiindicators?view=graph-rest-beta
func (m *TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder) Post(ctx context.Context, body TiindicatorsSubmittiindicatorsSubmitTiIndicatorsPostRequestBodyable, requestConfiguration *TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration)(TiindicatorsSubmittiindicatorsSubmitTiIndicatorsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiindicatorsSubmittiindicatorsSubmitTiIndicatorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiindicatorsSubmittiindicatorsSubmitTiIndicatorsResponseable), nil
}
// PostAsSubmitTiIndicatorsPostResponse upload multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a TiindicatorsSubmittiindicatorsSubmitTiIndicatorsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-submittiindicators?view=graph-rest-beta
func (m *TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder) PostAsSubmitTiIndicatorsPostResponse(ctx context.Context, body TiindicatorsSubmittiindicatorsSubmitTiIndicatorsPostRequestBodyable, requestConfiguration *TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration)(TiindicatorsSubmittiindicatorsSubmitTiIndicatorsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiindicatorsSubmittiindicatorsSubmitTiIndicatorsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiindicatorsSubmittiindicatorsSubmitTiIndicatorsPostResponseable), nil
}
// ToPostRequestInformation upload multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a *RequestInformation when successful
func (m *TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TiindicatorsSubmittiindicatorsSubmitTiIndicatorsPostRequestBodyable, requestConfiguration *TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a *TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder when successful
func (m *TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder) WithUrl(rawUrl string)(*TiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder) {
    return NewTiindicatorsSubmittiindicatorsSubmitTiIndicatorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
