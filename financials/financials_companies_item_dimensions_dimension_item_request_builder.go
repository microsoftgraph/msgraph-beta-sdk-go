package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder provides operations to manage the dimensions property of the microsoft.graph.company entity.
type FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemDimensionsDimensionItemRequestBuilderGetQueryParameters get dimensions from financials
type FinancialsCompaniesItemDimensionsDimensionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemDimensionsDimensionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemDimensionsDimensionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemDimensionsDimensionItemRequestBuilderGetQueryParameters
}
// NewFinancialsCompaniesItemDimensionsDimensionItemRequestBuilderInternal instantiates a new DimensionItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemDimensionsDimensionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder) {
    m := &FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/dimensions/{dimension%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemDimensionsDimensionItemRequestBuilder instantiates a new DimensionItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemDimensionsDimensionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemDimensionsDimensionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get dimensions from financials
func (m *FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemDimensionsDimensionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DimensionValues provides operations to manage the dimensionValues property of the microsoft.graph.dimension entity.
func (m *FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder) DimensionValues()(*FinancialsCompaniesItemDimensionsItemDimensionValuesRequestBuilder) {
    return NewFinancialsCompaniesItemDimensionsItemDimensionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DimensionValuesById provides operations to manage the dimensionValues property of the microsoft.graph.dimension entity.
func (m *FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder) DimensionValuesById(id string)(*FinancialsCompaniesItemDimensionsItemDimensionValuesDimensionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimensionValue%2Did"] = id
    }
    return NewFinancialsCompaniesItemDimensionsItemDimensionValuesDimensionValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get dimensions from financials
func (m *FinancialsCompaniesItemDimensionsDimensionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemDimensionsDimensionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Dimensionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDimensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Dimensionable), nil
}
