package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
type CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderGetQueryParameters get purchaseInvoices from financials
type CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderGetQueryParameters
}
// CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderInternal instantiates a new CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) {
    m := &CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoices/{purchaseInvoice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder instantiates a new CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.purchaseInvoice entity.
// returns a *CompaniesItemPurchaseinvoicesItemCurrencyRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) Currency()(*CompaniesItemPurchaseinvoicesItemCurrencyRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesItemCurrencyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get purchaseInvoices from financials
// returns a PurchaseInvoiceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable), nil
}
// Patch update the navigation property purchaseInvoices in financials
// returns a PurchaseInvoiceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, requestConfiguration *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable), nil
}
// PostPath provides operations to call the post method.
// returns a *CompaniesItemPurchaseinvoicesItemPostRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) PostPath()(*CompaniesItemPurchaseinvoicesItemPostRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesItemPostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PurchaseInvoiceLines provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.purchaseInvoice entity.
// returns a *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLinesRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) PurchaseInvoiceLines()(*CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLinesRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get purchaseInvoices from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property purchaseInvoices in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, requestConfiguration *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// VendorEscaped provides operations to manage the vendor property of the microsoft.graph.purchaseInvoice entity.
// returns a *CompaniesItemPurchaseinvoicesItemVendorRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) VendorEscaped()(*CompaniesItemPurchaseinvoicesItemVendorRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesItemVendorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
