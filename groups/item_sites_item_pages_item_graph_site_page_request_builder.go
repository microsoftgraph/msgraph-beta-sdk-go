package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGraphSitePageRequestBuilder casts the previous resource to sitePage.
type ItemSitesItemPagesItemGraphSitePageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGraphSitePageRequestBuilderGetQueryParameters get the item of type microsoft.graph.baseSitePage as microsoft.graph.sitePage
type ItemSitesItemPagesItemGraphSitePageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemPagesItemGraphSitePageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphSitePageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemGraphSitePageRequestBuilderGetQueryParameters
}
// CanvasLayout provides operations to manage the canvasLayout property of the microsoft.graph.sitePage entity.
func (m *ItemSitesItemPagesItemGraphSitePageRequestBuilder) CanvasLayout()(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageCanvasLayoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemPagesItemGraphSitePageRequestBuilderInternal instantiates a new GraphSitePageRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphSitePageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphSitePageRequestBuilder) {
    m := &ItemSitesItemPagesItemGraphSitePageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGraphSitePageRequestBuilder instantiates a new GraphSitePageRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphSitePageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphSitePageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGraphSitePageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
func (m *ItemSitesItemPagesItemGraphSitePageRequestBuilder) CreatedByUser()(*ItemSitesItemPagesItemGraphSitePageCreatedByUserRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the item of type microsoft.graph.baseSitePage as microsoft.graph.sitePage
func (m *ItemSitesItemPagesItemGraphSitePageRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphSitePageRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SitePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSitePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SitePageable), nil
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
func (m *ItemSitesItemPagesItemGraphSitePageRequestBuilder) LastModifiedByUser()(*ItemSitesItemPagesItemGraphSitePageLastModifiedByUserRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.baseSitePage as microsoft.graph.sitePage
func (m *ItemSitesItemPagesItemGraphSitePageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphSitePageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WebParts provides operations to manage the webParts property of the microsoft.graph.sitePage entity.
func (m *ItemSitesItemPagesItemGraphSitePageRequestBuilder) WebParts()(*ItemSitesItemPagesItemGraphSitePageWebPartsRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageWebPartsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemPagesItemGraphSitePageRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemGraphSitePageRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
