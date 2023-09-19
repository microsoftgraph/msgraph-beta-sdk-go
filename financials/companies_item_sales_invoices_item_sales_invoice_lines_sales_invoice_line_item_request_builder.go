package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder provides operations to manage the salesInvoiceLines property of the microsoft.graph.salesInvoice entity.
type CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetQueryParameters get salesInvoiceLines from financials
type CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.salesInvoiceLine entity.
func (m *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) Account()(*CompaniesItemSalesInvoicesItemSalesInvoiceLinesItemAccountRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal instantiates a new SalesInvoiceLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) {
    m := &CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}/salesInvoiceLines/{salesInvoiceLine%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder instantiates a new SalesInvoiceLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get salesInvoiceLines from financials
func (m *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable), nil
}
// Item provides operations to manage the item property of the microsoft.graph.salesInvoiceLine entity.
func (m *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) Item()(*CompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property salesInvoiceLines in financials
func (m *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable, requestConfiguration *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable), nil
}
// ToGetRequestInformation get salesInvoiceLines from financials
func (m *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property salesInvoiceLines in financials
func (m *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineable, requestConfiguration *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder) {
    return NewCompaniesItemSalesInvoicesItemSalesInvoiceLinesSalesInvoiceLineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
