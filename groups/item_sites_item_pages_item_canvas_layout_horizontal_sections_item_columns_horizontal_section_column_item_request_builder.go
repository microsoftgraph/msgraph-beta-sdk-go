package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters the set of vertical columns in this section.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters
}
// ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal instantiates a new HorizontalSectionColumnItemRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, horizontalSectionColumnId *string)(*ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    m := &ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if horizontalSectionColumnId != nil {
        urlTplParams["horizontalSectionColumn%2Did"] = *horizontalSectionColumnId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder instantiates a new HorizontalSectionColumnItemRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property columns for groups
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the set of vertical columns in this section.
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable), nil
}
// Patch update the navigation property columns in groups
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable), nil
}
// ToDeleteRequestInformation delete navigation property columns for groups
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the set of vertical columns in this section.
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property columns in groups
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HorizontalSectionColumnable, requestConfiguration *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Webparts provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Webparts()(*ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) {
    return NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WebpartsById provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
func (m *ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) WebpartsById(id string)(*ItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
