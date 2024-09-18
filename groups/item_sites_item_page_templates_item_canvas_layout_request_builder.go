package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder provides operations to manage the canvasLayout property of the microsoft.graph.pageTemplate entity.
type ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderGetQueryParameters the layout of the content in a given SharePoint page template, including horizontal sections and vertical sections.
type ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderGetQueryParameters
}
// ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderInternal instantiates a new ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder and sets the default values.
func NewItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) {
    m := &ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pageTemplates/{pageTemplate%2Did}/canvasLayout{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder instantiates a new ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder and sets the default values.
func NewItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property canvasLayout for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the layout of the content in a given SharePoint page template, including horizontal sections and vertical sections.
// returns a CanvasLayoutable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCanvasLayoutFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable), nil
}
// HorizontalSections provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
// returns a *ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsRequestBuilder when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) HorizontalSections()(*ItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsRequestBuilder) {
    return NewItemSitesItemPageTemplatesItemCanvasLayoutHorizontalSectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property canvasLayout in groups
// returns a CanvasLayoutable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCanvasLayoutFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable), nil
}
// ToDeleteRequestInformation delete navigation property canvasLayout for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the layout of the content in a given SharePoint page template, including horizontal sections and vertical sections.
// returns a *RequestInformation when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property canvasLayout in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CanvasLayoutable, requestConfiguration *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// VerticalSection provides operations to manage the verticalSection property of the microsoft.graph.canvasLayout entity.
// returns a *ItemSitesItemPageTemplatesItemCanvasLayoutVerticalSectionRequestBuilder when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) VerticalSection()(*ItemSitesItemPageTemplatesItemCanvasLayoutVerticalSectionRequestBuilder) {
    return NewItemSitesItemPageTemplatesItemCanvasLayoutVerticalSectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder when successful
func (m *ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder) {
    return NewItemSitesItemPageTemplatesItemCanvasLayoutRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
