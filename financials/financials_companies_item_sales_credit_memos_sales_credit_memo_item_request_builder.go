package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
type FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetQueryParameters get salesCreditMemos from financials
type FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderInternal instantiates a new SalesCreditMemoItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) {
    m := &FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder instantiates a new SalesCreditMemoItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesCreditMemos from financials
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesCreditMemos in financials
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, requestConfiguration *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currency provides operations to manage the currency property of the microsoft.graph.salesCreditMemo entity.
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) Currency()(*FinancialsCompaniesItemSalesCreditMemosItemCurrencyRequestBuilder) {
    return NewFinancialsCompaniesItemSalesCreditMemosItemCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesCreditMemo entity.
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) Customer()(*FinancialsCompaniesItemSalesCreditMemosItemCustomerRequestBuilder) {
    return NewFinancialsCompaniesItemSalesCreditMemosItemCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get salesCreditMemos from financials
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable), nil
}
// Patch update the navigation property salesCreditMemos in financials
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, requestConfiguration *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesCreditMemo entity.
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) PaymentTerm()(*FinancialsCompaniesItemSalesCreditMemosItemPaymentTermRequestBuilder) {
    return NewFinancialsCompaniesItemSalesCreditMemosItemPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemoLines provides operations to manage the salesCreditMemoLines property of the microsoft.graph.salesCreditMemo entity.
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) SalesCreditMemoLines()(*FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesRequestBuilder) {
    return NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemoLinesById provides operations to manage the salesCreditMemoLines property of the microsoft.graph.salesCreditMemo entity.
func (m *FinancialsCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) SalesCreditMemoLinesById(id string)(*FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemoLine%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
