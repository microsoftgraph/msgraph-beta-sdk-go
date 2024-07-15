package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeConnectorsItemSyncRequestBuilder provides operations to call the sync method.
type ExchangeConnectorsItemSyncRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeConnectorsItemSyncRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeConnectorsItemSyncRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExchangeConnectorsItemSyncRequestBuilderInternal instantiates a new ExchangeConnectorsItemSyncRequestBuilder and sets the default values.
func NewExchangeConnectorsItemSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeConnectorsItemSyncRequestBuilder) {
    m := &ExchangeConnectorsItemSyncRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/exchangeConnectors/{deviceManagementExchangeConnector%2Did}/sync", pathParameters),
    }
    return m
}
// NewExchangeConnectorsItemSyncRequestBuilder instantiates a new ExchangeConnectorsItemSyncRequestBuilder and sets the default values.
func NewExchangeConnectorsItemSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeConnectorsItemSyncRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeConnectorsItemSyncRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action sync
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeConnectorsItemSyncRequestBuilder) Post(ctx context.Context, body ExchangeConnectorsItemSyncPostRequestBodyable, requestConfiguration *ExchangeConnectorsItemSyncRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action sync
// returns a *RequestInformation when successful
func (m *ExchangeConnectorsItemSyncRequestBuilder) ToPostRequestInformation(ctx context.Context, body ExchangeConnectorsItemSyncPostRequestBodyable, requestConfiguration *ExchangeConnectorsItemSyncRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeConnectorsItemSyncRequestBuilder when successful
func (m *ExchangeConnectorsItemSyncRequestBuilder) WithUrl(rawUrl string)(*ExchangeConnectorsItemSyncRequestBuilder) {
    return NewExchangeConnectorsItemSyncRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
