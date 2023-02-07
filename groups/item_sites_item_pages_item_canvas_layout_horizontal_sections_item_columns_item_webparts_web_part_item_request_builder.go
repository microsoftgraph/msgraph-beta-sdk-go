package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderGetQueryParameters the collection of WebParts in this column.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderGetQueryParameters
}
// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderInternal instantiates a new WebPartItemRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) {
    m := &ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts/{webPart%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder instantiates a new WebPartItemRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property webparts for groups
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of WebParts in this column.
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWebPartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable), nil
}
// MicrosoftGraphGetPositionOfWebPart provides operations to call the getPositionOfWebPart method.
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) MicrosoftGraphGetPositionOfWebPart()(*ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemMicrosoftGraphGetPositionOfWebPartRequestBuilder) {
    return NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsItemMicrosoftGraphGetPositionOfWebPartRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property webparts in groups
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWebPartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable), nil
}
// ToDeleteRequestInformation delete navigation property webparts for groups
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the collection of WebParts in this column.
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property webparts in groups
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WebPartable, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
