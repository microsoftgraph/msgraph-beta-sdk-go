package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder provides operations to call the updateTiIndicators method.
type TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilderInternal instantiates a new TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder and sets the default values.
func NewTiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder) {
    m := &TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/tiIndicators/updateTiIndicators", pathParameters),
    }
    return m
}
// NewTiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder instantiates a new TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder and sets the default values.
func NewTiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: This method is obsolete. Use PostAsUpdateTiIndicatorsPostResponse instead.
// returns a TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-updatetiindicators?view=graph-rest-beta
func (m *TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder) Post(ctx context.Context, body TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsPostRequestBodyable, requestConfiguration *TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration)(TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiindicatorsUpdatetiindicatorsUpdateTiIndicatorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsResponseable), nil
}
// PostAsUpdateTiIndicatorsPostResponse update multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tiindicator-updatetiindicators?view=graph-rest-beta
func (m *TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder) PostAsUpdateTiIndicatorsPostResponse(ctx context.Context, body TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsPostRequestBodyable, requestConfiguration *TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration)(TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTiindicatorsUpdatetiindicatorsUpdateTiIndicatorsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsPostResponseable), nil
}
// ToPostRequestInformation update multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a *RequestInformation when successful
func (m *TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsPostRequestBodyable, requestConfiguration *TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder when successful
func (m *TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder) WithUrl(rawUrl string)(*TiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder) {
    return NewTiindicatorsUpdatetiindicatorsUpdateTiIndicatorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
