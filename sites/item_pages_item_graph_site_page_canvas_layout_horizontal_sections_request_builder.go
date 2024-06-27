package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderGetQueryParameters collection of horizontal sections on the SharePoint page.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderGetQueryParameters struct {
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
// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderGetQueryParameters
}
// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHorizontalSectionId provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
// returns a *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) ByHorizontalSectionId(horizontalSectionId string)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if horizontalSectionId != "" {
        urlTplParams["horizontalSection%2Did"] = horizontalSectionId
    }
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsHorizontalSectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderInternal instantiates a new ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder and sets the default values.
func NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) {
    m := &ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder instantiates a new ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder and sets the default values.
func NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsCountRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) Count()(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsCountRequestBuilder) {
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of horizontal sections on the SharePoint page.
// returns a HorizontalSectionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHorizontalSectionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionCollectionResponseable), nil
}
// Post create new navigation property to horizontalSections for sites
// returns a HorizontalSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionable, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHorizontalSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionable), nil
}
// ToGetRequestInformation collection of horizontal sections on the SharePoint page.
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to horizontalSections for sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionable, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder) {
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
