package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
type FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetQueryParameters get salesInvoices from financials
type FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Cancel()(*FinancialsCompaniesItemSalesInvoicesItemCancelRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CancelAndSend provides operations to call the cancelAndSend method.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) CancelAndSend()(*FinancialsCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewFinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderInternal instantiates a new SalesInvoiceItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) {
    m := &FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder instantiates a new SalesInvoiceItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesInvoices from financials
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesInvoices in financials
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, requestConfiguration *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currency provides operations to manage the currency property of the microsoft.graph.salesInvoice entity.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Currency()(*FinancialsCompaniesItemSalesInvoicesItemCurrencyRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesInvoice entity.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Customer()(*FinancialsCompaniesItemSalesInvoicesItemCustomerRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get salesInvoices from financials
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable), nil
}
// Patch update the navigation property salesInvoices in financials
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, requestConfiguration *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesInvoice entity.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) PaymentTerm()(*FinancialsCompaniesItemSalesInvoicesItemPaymentTermRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post provides operations to call the post method.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Post()(*FinancialsCompaniesItemSalesInvoicesItemPostRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemPostRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PostAndSend provides operations to call the postAndSend method.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) PostAndSend()(*FinancialsCompaniesItemSalesInvoicesItemPostAndSendRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemPostAndSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoiceLines provides operations to manage the salesInvoiceLines property of the microsoft.graph.salesInvoice entity.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) SalesInvoiceLines()(*FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoiceLinesById provides operations to manage the salesInvoiceLines property of the microsoft.graph.salesInvoice entity.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) SalesInvoiceLinesById(id string)(*FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoiceLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Send provides operations to call the send method.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Send()(*FinancialsCompaniesItemSalesInvoicesItemSendRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.salesInvoice entity.
func (m *FinancialsCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) ShipmentMethod()(*FinancialsCompaniesItemSalesInvoicesItemShipmentMethodRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
