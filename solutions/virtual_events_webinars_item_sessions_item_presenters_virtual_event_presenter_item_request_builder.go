package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder provides operations to manage the presenters property of the microsoft.graph.virtualEventSession entity.
type VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderGetQueryParameters get presenters from solutions
type VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderGetQueryParameters
}
// NewVirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderInternal instantiates a new VirtualEventPresenterItemRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder) {
    m := &VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/sessions/{virtualEventSession%2Did}/presenters/{virtualEventPresenter%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder instantiates a new VirtualEventPresenterItemRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get presenters from solutions
func (m *VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventPresenterable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventPresenterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventPresenterable), nil
}
// ProfilePhoto provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder) ProfilePhoto()(*VirtualEventsWebinarsItemSessionsItemPresentersItemProfilePhotoRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemPresentersItemProfilePhotoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get presenters from solutions
func (m *VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
