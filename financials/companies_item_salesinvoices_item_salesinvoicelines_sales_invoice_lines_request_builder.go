package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder provides operations to manage the salesInvoiceLines property of the microsoft.graph.salesInvoice entity.
type CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderGetQueryParameters get salesInvoiceLines from financials
type CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderGetQueryParameters
}
// BySalesInvoiceLineId provides operations to manage the salesInvoiceLines property of the microsoft.graph.salesInvoice entity.
// returns a *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLineItemRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) BySalesInvoiceLineId(salesInvoiceLineId string)(*CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if salesInvoiceLineId != "" {
        urlTplParams["salesInvoiceLine%2Did"] = salesInvoiceLineId
    }
    return NewCompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderInternal instantiates a new CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder and sets the default values.
func NewCompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) {
    m := &CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}/salesInvoiceLines{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder instantiates a new CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder and sets the default values.
func NewCompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemSalesinvoicesItemSalesinvoicelinesCountRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) Count()(*CompaniesItemSalesinvoicesItemSalesinvoicelinesCountRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemSalesinvoicelinesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesInvoiceLines from financials
// returns a SalesInvoiceLineCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesInvoiceLineCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesInvoiceLineCollectionResponseable), nil
}
// ToGetRequestInformation get salesInvoiceLines from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder when successful
func (m *CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder) {
    return NewCompaniesItemSalesinvoicesItemSalesinvoicelinesSalesInvoiceLinesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
