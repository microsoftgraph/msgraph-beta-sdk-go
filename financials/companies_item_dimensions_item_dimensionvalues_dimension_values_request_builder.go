package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder provides operations to manage the dimensionValues property of the microsoft.graph.dimension entity.
type CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderGetQueryParameters get dimensionValues from financials
type CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderGetQueryParameters struct {
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
// CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderGetQueryParameters
}
// ByDimensionValueId provides operations to manage the dimensionValues property of the microsoft.graph.dimension entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CompaniesItemDimensionsItemDimensionvaluesDimensionValueItemRequestBuilder when successful
func (m *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) ByDimensionValueId(dimensionValueId string)(*CompaniesItemDimensionsItemDimensionvaluesDimensionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if dimensionValueId != "" {
        urlTplParams["dimensionValue%2Did"] = dimensionValueId
    }
    return NewCompaniesItemDimensionsItemDimensionvaluesDimensionValueItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByDimensionValueIdGuid provides operations to manage the dimensionValues property of the microsoft.graph.dimension entity.
// returns a *CompaniesItemDimensionsItemDimensionvaluesDimensionValueItemRequestBuilder when successful
func (m *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) ByDimensionValueIdGuid(dimensionValueId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CompaniesItemDimensionsItemDimensionvaluesDimensionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["dimensionValue%2Did"] = dimensionValueId.String()
    return NewCompaniesItemDimensionsItemDimensionvaluesDimensionValueItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderInternal instantiates a new CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder and sets the default values.
func NewCompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) {
    m := &CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/dimensions/{dimension%2Did}/dimensionValues{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder instantiates a new CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder and sets the default values.
func NewCompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemDimensionsItemDimensionvaluesCountRequestBuilder when successful
func (m *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) Count()(*CompaniesItemDimensionsItemDimensionvaluesCountRequestBuilder) {
    return NewCompaniesItemDimensionsItemDimensionvaluesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get dimensionValues from financials
// returns a DimensionValueCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DimensionValueCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDimensionValueCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DimensionValueCollectionResponseable), nil
}
// ToGetRequestInformation get dimensionValues from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder when successful
func (m *CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder) {
    return NewCompaniesItemDimensionsItemDimensionvaluesDimensionValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
