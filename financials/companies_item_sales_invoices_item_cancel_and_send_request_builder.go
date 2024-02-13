package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder provides operations to call the cancelAndSend method.
type CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderInternal instantiates a new CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder and sets the default values.
func NewCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder) {
    m := &CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}/cancelAndSend", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder instantiates a new CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder and sets the default values.
func NewCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action cancelAndSend
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder) Post(ctx context.Context, requestConfiguration *CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action cancelAndSend
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder when successful
func (m *CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
