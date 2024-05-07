package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TiIndicatorsSubmitTiIndicatorsRequestBuilder provides operations to call the submitTiIndicators method.
type TiIndicatorsSubmitTiIndicatorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTiIndicatorsSubmitTiIndicatorsRequestBuilderInternal instantiates a new TiIndicatorsSubmitTiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsSubmitTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsSubmitTiIndicatorsRequestBuilder) {
    m := &TiIndicatorsSubmitTiIndicatorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/tiIndicators/submitTiIndicators", pathParameters),
    }
    return m
}
// NewTiIndicatorsSubmitTiIndicatorsRequestBuilder instantiates a new TiIndicatorsSubmitTiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsSubmitTiIndicatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsSubmitTiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiIndicatorsSubmitTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: This method is obsolete. Use PostAsSubmitTiIndicatorsPostResponse instead.
// returns a TiIndicatorsSubmitTiIndicatorsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-submittiindicators?view=graph-rest-beta
func (m *TiIndicatorsSubmitTiIndicatorsRequestBuilder) Post(ctx context.Context, body TiIndicatorsSubmitTiIndicatorsPostRequestBodyable, requestConfiguration *TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration)(TiIndicatorsSubmitTiIndicatorsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiIndicatorsSubmitTiIndicatorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiIndicatorsSubmitTiIndicatorsResponseable), nil
}
// PostAsSubmitTiIndicatorsPostResponse upload multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a TiIndicatorsSubmitTiIndicatorsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-submittiindicators?view=graph-rest-beta
func (m *TiIndicatorsSubmitTiIndicatorsRequestBuilder) PostAsSubmitTiIndicatorsPostResponse(ctx context.Context, body TiIndicatorsSubmitTiIndicatorsPostRequestBodyable, requestConfiguration *TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration)(TiIndicatorsSubmitTiIndicatorsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiIndicatorsSubmitTiIndicatorsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiIndicatorsSubmitTiIndicatorsPostResponseable), nil
}
// ToPostRequestInformation upload multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a *RequestInformation when successful
func (m *TiIndicatorsSubmitTiIndicatorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TiIndicatorsSubmitTiIndicatorsPostRequestBodyable, requestConfiguration *TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TiIndicatorsSubmitTiIndicatorsRequestBuilder when successful
func (m *TiIndicatorsSubmitTiIndicatorsRequestBuilder) WithUrl(rawUrl string)(*TiIndicatorsSubmitTiIndicatorsRequestBuilder) {
    return NewTiIndicatorsSubmitTiIndicatorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
