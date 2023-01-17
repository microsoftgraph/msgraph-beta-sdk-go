package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
type CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetQueryParameters get salesQuotes from financials
type CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderInternal instantiates a new SalesQuoteItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) {
    m := &CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder{
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
// NewCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder instantiates a new SalesQuoteItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesQuotesSalesQuoteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesQuotesSalesQuoteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.salesQuote entity.
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Currency()(*CompaniesItemSalesQuotesItemCurrencyRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesQuote entity.
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Customer()(*CompaniesItemSalesQuotesItemCustomerRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get salesQuotes from financials
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable), nil
}
// MakeInvoice provides operations to call the makeInvoice method.
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) MakeInvoice()(*CompaniesItemSalesQuotesItemMakeInvoiceRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemMakeInvoiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesQuotes in financials
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, requestConfiguration *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesQuote entity.
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) PaymentTerm()(*CompaniesItemSalesQuotesItemPaymentTermRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) SalesQuoteLines()(*CompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLinesById provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) SalesQuoteLinesById(id string)(*CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuoteLine%2Did"] = id
    }
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Send provides operations to call the send method.
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) Send()(*CompaniesItemSalesQuotesItemSendRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.salesQuote entity.
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) ShipmentMethod()(*CompaniesItemSalesQuotesItemShipmentMethodRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToGetRequestInformation get salesQuotes from financials
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property salesQuotes in financials
func (m *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteable, requestConfiguration *CompaniesItemSalesQuotesSalesQuoteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
