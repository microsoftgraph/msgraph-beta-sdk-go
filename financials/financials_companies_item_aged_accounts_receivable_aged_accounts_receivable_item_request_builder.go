package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
type FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetQueryParameters get agedAccountsReceivable from financials
type FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetQueryParameters
}
// NewFinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderInternal instantiates a new AgedAccountsReceivableItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) {
    m := &FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/agedAccountsReceivable/{agedAccountsReceivable%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder instantiates a new AgedAccountsReceivableItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get agedAccountsReceivable from financials
func (m *FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get agedAccountsReceivable from financials
func (m *FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgedAccountsReceivableable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgedAccountsReceivableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgedAccountsReceivableable), nil
}
