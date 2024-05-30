package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesProductsRequestBuilder provides operations to manage the products property of the microsoft.graph.adminWindowsUpdates entity.
type WindowsUpdatesProductsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesProductsRequestBuilderGetQueryParameters a collection of Windows products.
type WindowsUpdatesProductsRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesProductsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesProductsRequestBuilderGetQueryParameters
}
// WindowsUpdatesProductsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByProductId provides operations to manage the products property of the microsoft.graph.adminWindowsUpdates entity.
// returns a *WindowsUpdatesProductsProductItemRequestBuilder when successful
func (m *WindowsUpdatesProductsRequestBuilder) ByProductId(productId string)(*WindowsUpdatesProductsProductItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if productId != "" {
        urlTplParams["product%2Did"] = productId
    }
    return NewWindowsUpdatesProductsProductItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsUpdatesProductsRequestBuilderInternal instantiates a new WindowsUpdatesProductsRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsRequestBuilder) {
    m := &WindowsUpdatesProductsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/products{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesProductsRequestBuilder instantiates a new WindowsUpdatesProductsRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesProductsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsUpdatesProductsCountRequestBuilder when successful
func (m *WindowsUpdatesProductsRequestBuilder) Count()(*WindowsUpdatesProductsCountRequestBuilder) {
    return NewWindowsUpdatesProductsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of Windows products.
// returns a ProductCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesProductsRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ProductCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateProductCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ProductCollectionResponseable), nil
}
// MicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogID provides operations to call the findByCatalogId method.
// returns a *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder when successful
func (m *WindowsUpdatesProductsRequestBuilder) MicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogID(catalogID *string)(*WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder) {
    return NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, catalogID)
}
// MicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumber provides operations to call the findByKbNumber method.
// returns a *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder when successful
func (m *WindowsUpdatesProductsRequestBuilder) MicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumber(kbNumber *int32)(*WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder) {
    return NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, kbNumber)
}
// Post create new navigation property to products for admin
// returns a Productable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsRequestBuilder) Post(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable, requestConfiguration *WindowsUpdatesProductsRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateProductFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable), nil
}
// ToGetRequestInformation a collection of Windows products.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to products for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable, requestConfiguration *WindowsUpdatesProductsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsUpdatesProductsRequestBuilder when successful
func (m *WindowsUpdatesProductsRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesProductsRequestBuilder) {
    return NewWindowsUpdatesProductsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
