package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder provides operations to manage the salesCreditMemoLines property of the microsoft.graph.salesCreditMemo entity.
type FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetQueryParameters get salesCreditMemoLines from financials
type FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.salesCreditMemoLine entity.
func (m *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) Account()(*FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder) {
    return NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal instantiates a new SalesCreditMemoLineItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) {
    m := &FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}/salesCreditMemoLines/{salesCreditMemoLine%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder instantiates a new SalesCreditMemoLineItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesCreditMemoLines from financials
func (m *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesCreditMemoLines in financials
func (m *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, requestConfiguration *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get salesCreditMemoLines from financials
func (m *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable), nil
}
// Item provides operations to manage the item property of the microsoft.graph.salesCreditMemoLine entity.
func (m *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) Item()(*FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemRequestBuilder) {
    return NewFinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesCreditMemoLines in financials
func (m *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, requestConfiguration *FinancialsCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable), nil
}
