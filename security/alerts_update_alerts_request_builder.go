package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AlertsUpdateAlertsRequestBuilder provides operations to call the updateAlerts method.
type AlertsUpdateAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AlertsUpdateAlertsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AlertsUpdateAlertsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAlertsUpdateAlertsRequestBuilderInternal instantiates a new UpdateAlertsRequestBuilder and sets the default values.
func NewAlertsUpdateAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertsUpdateAlertsRequestBuilder) {
    m := &AlertsUpdateAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/alerts/updateAlerts", pathParameters),
    }
    return m
}
// NewAlertsUpdateAlertsRequestBuilder instantiates a new UpdateAlertsRequestBuilder and sets the default values.
func NewAlertsUpdateAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertsUpdateAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAlertsUpdateAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update multiple alerts in one request instead of multiple requests.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/alert-updatealerts?view=graph-rest-1.0
func (m *AlertsUpdateAlertsRequestBuilder) Post(ctx context.Context, body AlertsUpdateAlertsPostRequestBodyable, requestConfiguration *AlertsUpdateAlertsRequestBuilderPostRequestConfiguration)(AlertsUpdateAlertsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertsUpdateAlertsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertsUpdateAlertsResponseable), nil
}
// ToPostRequestInformation update multiple alerts in one request instead of multiple requests.
func (m *AlertsUpdateAlertsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AlertsUpdateAlertsPostRequestBodyable, requestConfiguration *AlertsUpdateAlertsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
