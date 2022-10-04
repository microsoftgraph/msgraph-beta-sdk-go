package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i234f9ceec46f67f5e601f3edb10d34513a18ecd109769d4b121a078d1b5a44d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/dimensions/item/dimensionvalues"
    i8ed42469cc87a553b3cd8749bdd04e937f267d96ef2528e79e78a8822e6f6112 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/dimensions/item/dimensionvalues/item"
)

// DimensionItemRequestBuilder provides operations to manage the dimensions property of the microsoft.graph.company entity.
type DimensionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DimensionItemRequestBuilderGetQueryParameters get dimensions from financials
type DimensionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DimensionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DimensionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DimensionItemRequestBuilderGetQueryParameters
}
// NewDimensionItemRequestBuilderInternal instantiates a new DimensionItemRequestBuilder and sets the default values.
func NewDimensionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DimensionItemRequestBuilder) {
    m := &DimensionItemRequestBuilder{
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
// NewDimensionItemRequestBuilder instantiates a new DimensionItemRequestBuilder and sets the default values.
func NewDimensionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DimensionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDimensionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get dimensions from financials
func (m *DimensionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DimensionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DimensionValues the dimensionValues property
func (m *DimensionItemRequestBuilder) DimensionValues()(*i234f9ceec46f67f5e601f3edb10d34513a18ecd109769d4b121a078d1b5a44d3.DimensionValuesRequestBuilder) {
    return i234f9ceec46f67f5e601f3edb10d34513a18ecd109769d4b121a078d1b5a44d3.NewDimensionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DimensionValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.dimensions.item.dimensionValues.item collection
func (m *DimensionItemRequestBuilder) DimensionValuesById(id string)(*i8ed42469cc87a553b3cd8749bdd04e937f267d96ef2528e79e78a8822e6f6112.DimensionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimensionValue%2Did"] = id
    }
    return i8ed42469cc87a553b3cd8749bdd04e937f267d96ef2528e79e78a8822e6f6112.NewDimensionValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get dimensions from financials
func (m *DimensionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DimensionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Dimensionable, error) {
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
