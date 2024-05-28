package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AlertsUpdatealertsUpdateAlertsRequestBuilder provides operations to call the updateAlerts method.
type AlertsUpdatealertsUpdateAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AlertsUpdatealertsUpdateAlertsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AlertsUpdatealertsUpdateAlertsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAlertsUpdatealertsUpdateAlertsRequestBuilderInternal instantiates a new AlertsUpdatealertsUpdateAlertsRequestBuilder and sets the default values.
func NewAlertsUpdatealertsUpdateAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertsUpdatealertsUpdateAlertsRequestBuilder) {
    m := &AlertsUpdatealertsUpdateAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/alerts/updateAlerts", pathParameters),
    }
    return m
}
// NewAlertsUpdatealertsUpdateAlertsRequestBuilder instantiates a new AlertsUpdatealertsUpdateAlertsRequestBuilder and sets the default values.
func NewAlertsUpdatealertsUpdateAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertsUpdatealertsUpdateAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAlertsUpdatealertsUpdateAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update multiple alerts in one request instead of multiple requests.
// Deprecated: This method is obsolete. Use PostAsUpdateAlertsPostResponse instead.
// returns a AlertsUpdatealertsUpdateAlertsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/alert-updatealerts?view=graph-rest-beta
func (m *AlertsUpdatealertsUpdateAlertsRequestBuilder) Post(ctx context.Context, body AlertsUpdatealertsUpdateAlertsPostRequestBodyable, requestConfiguration *AlertsUpdatealertsUpdateAlertsRequestBuilderPostRequestConfiguration)(AlertsUpdatealertsUpdateAlertsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertsUpdatealertsUpdateAlertsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertsUpdatealertsUpdateAlertsResponseable), nil
}
// PostAsUpdateAlertsPostResponse update multiple alerts in one request instead of multiple requests.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a AlertsUpdatealertsUpdateAlertsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/alert-updatealerts?view=graph-rest-beta
func (m *AlertsUpdatealertsUpdateAlertsRequestBuilder) PostAsUpdateAlertsPostResponse(ctx context.Context, body AlertsUpdatealertsUpdateAlertsPostRequestBodyable, requestConfiguration *AlertsUpdatealertsUpdateAlertsRequestBuilderPostRequestConfiguration)(AlertsUpdatealertsUpdateAlertsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertsUpdatealertsUpdateAlertsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertsUpdatealertsUpdateAlertsPostResponseable), nil
}
// ToPostRequestInformation update multiple alerts in one request instead of multiple requests.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a *RequestInformation when successful
func (m *AlertsUpdatealertsUpdateAlertsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AlertsUpdatealertsUpdateAlertsPostRequestBodyable, requestConfiguration *AlertsUpdatealertsUpdateAlertsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AlertsUpdatealertsUpdateAlertsRequestBuilder when successful
func (m *AlertsUpdatealertsUpdateAlertsRequestBuilder) WithUrl(rawUrl string)(*AlertsUpdatealertsUpdateAlertsRequestBuilder) {
    return NewAlertsUpdatealertsUpdateAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
