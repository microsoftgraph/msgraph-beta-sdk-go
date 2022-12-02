package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder provides operations to manage the salesInvoiceLines property of the microsoft.graph.salesInvoice entity.
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetQueryParameters get salesInvoiceLines from financials
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.salesInvoiceLine entity.
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) Account()(*FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemAccountRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal instantiates a new SalesInvoiceLineItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) {
    m := &FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}/salesInvoiceLines/{salesInvoiceLine%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder instantiates a new SalesInvoiceLineItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesInvoiceLines from financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesInvoiceLines in financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get salesInvoiceLines from financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable), nil
}
// Item provides operations to manage the item property of the microsoft.graph.salesInvoiceLine entity.
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) Item()(*FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemRequestBuilder) {
    return NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesInvoiceLines in financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable), nil
}
