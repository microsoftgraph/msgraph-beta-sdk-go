package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
type CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetQueryParameters get salesInvoices from financials
type CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Cancel()(*CompaniesItemSalesInvoicesItemCancelRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CancelAndSend provides operations to call the cancelAndSend method.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) CancelAndSend()(*CompaniesItemSalesInvoicesItemCancelAndSendRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemCancelAndSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderInternal instantiates a new SalesInvoiceItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) {
    m := &CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder instantiates a new SalesInvoiceItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.salesInvoice entity.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Currency()(*CompaniesItemSalesInvoicesItemCurrencyRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemCurrencyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesInvoice entity.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Customer()(*CompaniesItemSalesInvoicesItemCustomerRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemCustomerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesInvoices from financials
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, requestConfiguration *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) PaymentTerm()(*CompaniesItemSalesInvoicesItemPaymentTermRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemPaymentTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PostAndSend provides operations to call the postAndSend method.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) PostAndSend()(*CompaniesItemSalesInvoicesItemPostAndSendRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemPostAndSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PostPath provides operations to call the post method.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) PostPath()(*CompaniesItemSalesInvoicesItemPostRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemPostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesInvoiceLines provides operations to manage the salesInvoiceLines property of the microsoft.graph.salesInvoice entity.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) SalesInvoiceLines()(*CompaniesItemSalesInvoicesItemSalesInvoiceLinesRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Send provides operations to call the send method.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) Send()(*CompaniesItemSalesInvoicesItemSendRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemSendRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.salesInvoice entity.
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) ShipmentMethod()(*CompaniesItemSalesInvoicesItemShipmentMethodRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemShipmentMethodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get salesInvoices from financials
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceable, requestConfiguration *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder) {
    return NewCompaniesItemSalesInvoicesSalesInvoiceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
