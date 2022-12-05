package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
type FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetQueryParameters get salesQuotes from financials
type FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderInternal instantiates a new SalesQuoteItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) {
    m := &FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesQuotes/{salesQuote%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder instantiates a new SalesQuoteItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesQuotes from financials
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property salesQuotes in financials
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, requestConfiguration *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Currency provides operations to manage the currency property of the microsoft.graph.salesQuote entity.
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Currency()(*FinancialsCompaniesItemSalesQuotesItemCurrencyRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuotesItemCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesQuote entity.
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Customer()(*FinancialsCompaniesItemSalesQuotesItemCustomerRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuotesItemCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get salesQuotes from financials
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable), nil
}
// MakeInvoice provides operations to call the makeInvoice method.
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) MakeInvoice()(*FinancialsCompaniesItemSalesQuotesItemMakeInvoiceRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuotesItemMakeInvoiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesQuotes in financials
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, requestConfiguration *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesQuote entity.
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) PaymentTerm()(*FinancialsCompaniesItemSalesQuotesItemPaymentTermRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuotesItemPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) SalesQuoteLines()(*FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLinesById provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) SalesQuoteLinesById(id string)(*FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuoteLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Send provides operations to call the send method.
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Send()(*FinancialsCompaniesItemSalesQuotesItemSendRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuotesItemSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.salesQuote entity.
func (m *FinancialsCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) ShipmentMethod()(*FinancialsCompaniesItemSalesQuotesItemShipmentMethodRequestBuilder) {
    return NewFinancialsCompaniesItemSalesQuotesItemShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
