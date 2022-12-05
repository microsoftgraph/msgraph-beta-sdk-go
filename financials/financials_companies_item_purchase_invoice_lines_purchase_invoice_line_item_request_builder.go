package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
type FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters get purchaseInvoiceLines from financials
type FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.purchaseInvoiceLine entity.
func (m *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) Account()(*FinancialsCompaniesItemPurchaseInvoiceLinesItemAccountRequestBuilder) {
    return NewFinancialsCompaniesItemPurchaseInvoiceLinesItemAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewFinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal instantiates a new PurchaseInvoiceLineItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) {
    m := &FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoiceLines/{purchaseInvoiceLine%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder instantiates a new PurchaseInvoiceLineItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get purchaseInvoiceLines from financials
func (m *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property purchaseInvoiceLines in financials
func (m *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, requestConfiguration *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get purchaseInvoiceLines from financials
func (m *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable), nil
}
// Item provides operations to manage the item property of the microsoft.graph.purchaseInvoiceLine entity.
func (m *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) Item()(*FinancialsCompaniesItemPurchaseInvoiceLinesItemItemRequestBuilder) {
    return NewFinancialsCompaniesItemPurchaseInvoiceLinesItemItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property purchaseInvoiceLines in financials
func (m *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, requestConfiguration *FinancialsCompaniesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable), nil
}
