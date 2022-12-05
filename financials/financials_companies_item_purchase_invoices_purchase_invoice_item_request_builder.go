package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
type FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderGetQueryParameters get purchaseInvoices from financials
type FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderInternal instantiates a new PurchaseInvoiceItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) {
    m := &FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoices/{purchaseInvoice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder instantiates a new PurchaseInvoiceItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get purchaseInvoices from financials
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property purchaseInvoices in financials
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, requestConfiguration *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currency provides operations to manage the currency property of the microsoft.graph.purchaseInvoice entity.
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) Currency()(*FinancialsCompaniesItemPurchaseInvoicesItemCurrencyRequestBuilder) {
    return NewFinancialsCompaniesItemPurchaseInvoicesItemCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get purchaseInvoices from financials
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable), nil
}
// Patch update the navigation property purchaseInvoices in financials
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, requestConfiguration *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceable), nil
}
// Post provides operations to call the post method.
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) Post()(*FinancialsCompaniesItemPurchaseInvoicesItemPostRequestBuilder) {
    return NewFinancialsCompaniesItemPurchaseInvoicesItemPostRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoiceLines provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.purchaseInvoice entity.
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) PurchaseInvoiceLines()(*FinancialsCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesRequestBuilder) {
    return NewFinancialsCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoiceLinesById provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.purchaseInvoice entity.
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) PurchaseInvoiceLinesById(id string)(*FinancialsCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoiceLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Vendor_escaped provides operations to manage the vendor property of the microsoft.graph.purchaseInvoice entity.
func (m *FinancialsCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) Vendor_escaped()(*FinancialsCompaniesItemPurchaseInvoicesItemVendorRequestBuilder) {
    return NewFinancialsCompaniesItemPurchaseInvoicesItemVendorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
