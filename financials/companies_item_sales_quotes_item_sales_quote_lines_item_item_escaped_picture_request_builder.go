package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder provides operations to manage the picture property of the microsoft.graph.item entity.
type CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderGetQueryParameters get picture from financials
type CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderGetQueryParameters struct {
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
// CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderGetQueryParameters
}
// CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPictureId provides operations to manage the picture property of the microsoft.graph.item entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPicturePictureItemRequestBuilder when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) ByPictureId(pictureId string)(*CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPicturePictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if pictureId != "" {
        urlTplParams["picture%2Did"] = pictureId
    }
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPicturePictureItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByPictureIdGuid provides operations to manage the picture property of the microsoft.graph.item entity.
// returns a *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPicturePictureItemRequestBuilder when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) ByPictureIdGuid(pictureId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPicturePictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["picture%2Did"] = pictureId.String()
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPicturePictureItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderInternal instantiates a new CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder and sets the default values.
func NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) {
    m := &CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesQuotes/{salesQuote%2Did}/salesQuoteLines/{salesQuoteLine%2Did}/item/picture{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder instantiates a new CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder and sets the default values.
func NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureCountRequestBuilder when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) Count()(*CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureCountRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get picture from financials
// returns a PictureCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PictureCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePictureCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PictureCollectionResponseable), nil
}
// Post create new navigation property to picture for financials
// returns a Pictureable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Pictureable, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Pictureable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePictureFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Pictureable), nil
}
// ToGetRequestInformation get picture from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to picture for financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Pictureable, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedPictureRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
