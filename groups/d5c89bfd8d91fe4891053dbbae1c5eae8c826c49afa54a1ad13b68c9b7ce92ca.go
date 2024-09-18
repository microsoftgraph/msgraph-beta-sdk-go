package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
type ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters the collection of WebParts in this column.
type ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters
}
// ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWebPartId provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
// returns a *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) ByWebPartId(webPartId string)(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if webPartId != "" {
        urlTplParams["webPart%2Did"] = webPartId
    }
    return NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal instantiates a new ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder and sets the default values.
func NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) {
    m := &ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pageTemplates/{pageTemplate%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder instantiates a new ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder and sets the default values.
func NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Count()(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) {
    return NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of WebParts in this column.
// returns a WebPartCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWebPartCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartCollectionResponseable), nil
}
// Post create new navigation property to webparts for groups
// returns a WebPartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWebPartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable), nil
}
// ToGetRequestInformation the collection of WebParts in this column.
// returns a *RequestInformation when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to webparts for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) {
    return NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
