package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder provides operations to manage the sites property of the microsoft.graph.site entity.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetQueryParameters get a collection of subsites defined for a [site][].
type ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetQueryParameters
}
// NewItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderInternal instantiates a new ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder and sets the default values.
func NewItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) {
    m := &ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/sites{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder instantiates a new ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder and sets the default values.
func NewItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a collection of subsites defined for a [site][].
// returns a SiteCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/site-list-subsites?view=graph-rest-1.0
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteCollectionResponseable), nil
}
// ToGetRequestInformation get a collection of subsites defined for a [site][].
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder when successful
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
