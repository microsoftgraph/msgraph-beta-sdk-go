package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder provides operations to count the resources in the collection.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetQueryParameters
}
// NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) {
    m := &ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) {
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
