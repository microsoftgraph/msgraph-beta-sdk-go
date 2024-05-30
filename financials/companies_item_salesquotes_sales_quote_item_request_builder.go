package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesquotesSalesQuoteItemRequestBuilder provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
type CompaniesItemSalesquotesSalesQuoteItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesquotesSalesQuoteItemRequestBuilderGetQueryParameters get salesQuotes from financials
type CompaniesItemSalesquotesSalesQuoteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesquotesSalesQuoteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesquotesSalesQuoteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesquotesSalesQuoteItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesquotesSalesQuoteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesquotesSalesQuoteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalesquotesSalesQuoteItemRequestBuilderInternal instantiates a new CompaniesItemSalesquotesSalesQuoteItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesquotesSalesQuoteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) {
    m := &CompaniesItemSalesquotesSalesQuoteItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesQuotes/{salesQuote%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesquotesSalesQuoteItemRequestBuilder instantiates a new CompaniesItemSalesquotesSalesQuoteItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesquotesSalesQuoteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesquotesSalesQuoteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.salesQuote entity.
// returns a *CompaniesItemSalesquotesItemCurrencyRequestBuilder when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) Currency()(*CompaniesItemSalesquotesItemCurrencyRequestBuilder) {
    return NewCompaniesItemSalesquotesItemCurrencyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesQuote entity.
// returns a *CompaniesItemSalesquotesItemCustomerRequestBuilder when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) Customer()(*CompaniesItemSalesquotesItemCustomerRequestBuilder) {
    return NewCompaniesItemSalesquotesItemCustomerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesQuotes from financials
// returns a SalesQuoteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesquotesSalesQuoteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable), nil
}
// MakeInvoice provides operations to call the makeInvoice method.
// returns a *CompaniesItemSalesquotesItemMakeinvoiceMakeInvoiceRequestBuilder when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) MakeInvoice()(*CompaniesItemSalesquotesItemMakeinvoiceMakeInvoiceRequestBuilder) {
    return NewCompaniesItemSalesquotesItemMakeinvoiceMakeInvoiceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property salesQuotes in financials
// returns a SalesQuoteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, requestConfiguration *CompaniesItemSalesquotesSalesQuoteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesQuote entity.
// returns a *CompaniesItemSalesquotesItemPaymenttermPaymentTermRequestBuilder when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) PaymentTerm()(*CompaniesItemSalesquotesItemPaymenttermPaymentTermRequestBuilder) {
    return NewCompaniesItemSalesquotesItemPaymenttermPaymentTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
// returns a *CompaniesItemSalesquotesItemSalesquotelinesSalesQuoteLinesRequestBuilder when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) SalesQuoteLines()(*CompaniesItemSalesquotesItemSalesquotelinesSalesQuoteLinesRequestBuilder) {
    return NewCompaniesItemSalesquotesItemSalesquotelinesSalesQuoteLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Send provides operations to call the send method.
// returns a *CompaniesItemSalesquotesItemSendRequestBuilder when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) Send()(*CompaniesItemSalesquotesItemSendRequestBuilder) {
    return NewCompaniesItemSalesquotesItemSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.salesQuote entity.
// returns a *CompaniesItemSalesquotesItemShipmentmethodShipmentMethodRequestBuilder when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) ShipmentMethod()(*CompaniesItemSalesquotesItemShipmentmethodShipmentMethodRequestBuilder) {
    return NewCompaniesItemSalesquotesItemShipmentmethodShipmentMethodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get salesQuotes from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesquotesSalesQuoteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property salesQuotes in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, requestConfiguration *CompaniesItemSalesquotesSalesQuoteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder when successful
func (m *CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesquotesSalesQuoteItemRequestBuilder) {
    return NewCompaniesItemSalesquotesSalesQuoteItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
