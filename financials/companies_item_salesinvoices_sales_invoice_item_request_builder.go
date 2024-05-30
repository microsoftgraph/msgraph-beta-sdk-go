package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
type CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderGetQueryParameters get salesInvoices from financials
type CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
// returns a *CompaniesItemSalesinvoicesItemCancelRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) Cancel()(*CompaniesItemSalesinvoicesItemCancelRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CancelAndSend provides operations to call the cancelAndSend method.
// returns a *CompaniesItemSalesinvoicesItemCancelandsendCancelAndSendRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) CancelAndSend()(*CompaniesItemSalesinvoicesItemCancelandsendCancelAndSendRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemCancelandsendCancelAndSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderInternal instantiates a new CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) {
    m := &CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder instantiates a new CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.salesInvoice entity.
// returns a *CompaniesItemSalesinvoicesItemCurrencyRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) Currency()(*CompaniesItemSalesinvoicesItemCurrencyRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemCurrencyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesInvoice entity.
// returns a *CompaniesItemSalesinvoicesItemCustomerRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) Customer()(*CompaniesItemSalesinvoicesItemCustomerRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemCustomerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesInvoices from financials
// returns a SalesInvoiceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable), nil
}
// Patch update the navigation property salesInvoices in financials
// returns a SalesInvoiceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, requestConfiguration *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesInvoice entity.
// returns a *CompaniesItemSalesinvoicesItemPaymenttermPaymentTermRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) PaymentTerm()(*CompaniesItemSalesinvoicesItemPaymenttermPaymentTermRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemPaymenttermPaymentTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PostAndSend provides operations to call the postAndSend method.
// returns a *CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) PostAndSend()(*CompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemPostandsendPostAndSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PostPath provides operations to call the post method.
// returns a *CompaniesItemSalesinvoicesItemPostRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) PostPath()(*CompaniesItemSalesinvoicesItemPostRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemPostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoiceLines provides operations to manage the salesInvoiceLines property of the microsoft.graph.salesInvoice entity.
// returns a *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) SalesInvoiceLines()(*CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Send provides operations to call the send method.
// returns a *CompaniesItemSalesinvoicesItemSendRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) Send()(*CompaniesItemSalesinvoicesItemSendRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.salesInvoice entity.
// returns a *CompaniesItemSalesinvoicesItemShipmentmethodShipmentMethodRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) ShipmentMethod()(*CompaniesItemSalesinvoicesItemShipmentmethodShipmentMethodRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemShipmentmethodShipmentMethodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get salesInvoices from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property salesInvoices in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, requestConfiguration *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder) {
    return NewCompaniesItemSalesinvoicesSalesInvoiceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
