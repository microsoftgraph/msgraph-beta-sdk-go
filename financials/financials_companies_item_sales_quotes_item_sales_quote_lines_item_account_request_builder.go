package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder provides operations to manage the account property of the microsoft.graph.salesQuoteLine entity.
type FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderGetQueryParameters get account from financials
type FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderGetQueryParameters
}
// NewFinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderInternal instantiates a new AccountRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder) {
    m := &FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesQuotes/{salesQuote%2Did}/salesQuoteLines/{salesQuoteLine%2Did}/account{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder instantiates a new AccountRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get account from financials
func (m *FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get account from financials
func (m *FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Accountable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccountFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Accountable), nil
}
