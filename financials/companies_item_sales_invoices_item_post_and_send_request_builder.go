package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder provides operations to call the postAndSend method.
type CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesInvoicesItemPostAndSendRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesInvoicesItemPostAndSendRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalesInvoicesItemPostAndSendRequestBuilderInternal instantiates a new CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder and sets the default values.
func NewCompaniesItemSalesInvoicesItemPostAndSendRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder) {
    m := &CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}/postAndSend", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesInvoicesItemPostAndSendRequestBuilder instantiates a new CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder and sets the default values.
func NewCompaniesItemSalesInvoicesItemPostAndSendRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesInvoicesItemPostAndSendRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action postAndSend
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder) Post(ctx context.Context, requestConfiguration *CompaniesItemSalesInvoicesItemPostAndSendRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action postAndSend
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesInvoicesItemPostAndSendRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder when successful
func (m *CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemPostAndSendRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
