package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesOrdersSalesOrderItemRequestBuilder provides operations to manage the salesOrders property of the microsoft.graph.company entity.
type CompaniesItemSalesOrdersSalesOrderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CompaniesItemSalesOrdersSalesOrderItemRequestBuilderGetQueryParameters get salesOrders from financials
type CompaniesItemSalesOrdersSalesOrderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesOrdersSalesOrderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesOrdersSalesOrderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesOrdersSalesOrderItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesOrdersSalesOrderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesOrdersSalesOrderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalesOrdersSalesOrderItemRequestBuilderInternal instantiates a new SalesOrderItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesOrdersSalesOrderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) {
    m := &CompaniesItemSalesOrdersSalesOrderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesOrders/{salesOrder%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCompaniesItemSalesOrdersSalesOrderItemRequestBuilder instantiates a new SalesOrderItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesOrdersSalesOrderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesOrdersSalesOrderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesOrders from financials
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesOrdersSalesOrderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesOrders in financials
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable, requestConfiguration *CompaniesItemSalesOrdersSalesOrderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currency provides operations to manage the currency property of the microsoft.graph.salesOrder entity.
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) Currency()(*CompaniesItemSalesOrdersItemCurrencyRequestBuilder) {
    return NewCompaniesItemSalesOrdersItemCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesOrder entity.
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) Customer()(*CompaniesItemSalesOrdersItemCustomerRequestBuilder) {
    return NewCompaniesItemSalesOrdersItemCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get salesOrders from financials
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesOrdersSalesOrderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesOrderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable), nil
}
// Patch update the navigation property salesOrders in financials
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable, requestConfiguration *CompaniesItemSalesOrdersSalesOrderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesOrderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesOrder entity.
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) PaymentTerm()(*CompaniesItemSalesOrdersItemPaymentTermRequestBuilder) {
    return NewCompaniesItemSalesOrdersItemPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrderLines provides operations to manage the salesOrderLines property of the microsoft.graph.salesOrder entity.
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) SalesOrderLines()(*CompaniesItemSalesOrdersItemSalesOrderLinesRequestBuilder) {
    return NewCompaniesItemSalesOrdersItemSalesOrderLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrderLinesById provides operations to manage the salesOrderLines property of the microsoft.graph.salesOrder entity.
func (m *CompaniesItemSalesOrdersSalesOrderItemRequestBuilder) SalesOrderLinesById(id string)(*CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrderLine%2Did"] = id
    }
    return NewCompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
