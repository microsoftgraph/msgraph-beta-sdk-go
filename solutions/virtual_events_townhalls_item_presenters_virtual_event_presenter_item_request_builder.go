package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder provides operations to manage the presenters property of the microsoft.graph.virtualEvent entity.
type VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderGetQueryParameters read the properties and relationships of a virtualEventPresenter object.
type VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderGetQueryParameters
}
// VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderInternal instantiates a new VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) {
    m := &VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/presenters/{virtualEventPresenter%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder instantiates a new VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a virtualEventPresenter from a virtual event. Currently the supported virtual event type is virtualEventTownhall.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventpresenter-delete?view=graph-rest-beta
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a virtualEventPresenter object.
// returns a VirtualEventPresenterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventpresenter-get?view=graph-rest-beta
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventPresenterable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property presenters in solutions
// returns a VirtualEventPresenterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventPresenterable, requestConfiguration *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventPresenterable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Sessions provides operations to manage the sessions property of the microsoft.graph.virtualEventPresenter entity.
// returns a *VirtualEventsTownhallsItemPresentersItemSessionsRequestBuilder when successful
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) Sessions()(*VirtualEventsTownhallsItemPresentersItemSessionsRequestBuilder) {
    return NewVirtualEventsTownhallsItemPresentersItemSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SessionsWithJoinWebUrl provides operations to manage the sessions property of the microsoft.graph.virtualEventPresenter entity.
// returns a *VirtualEventsTownhallsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder when successful
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) SessionsWithJoinWebUrl(joinWebUrl *string)(*VirtualEventsTownhallsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder) {
    return NewVirtualEventsTownhallsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, joinWebUrl)
}
// ToDeleteRequestInformation delete a virtualEventPresenter from a virtual event. Currently the supported virtual event type is virtualEventTownhall.
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a virtualEventPresenter object.
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property presenters in solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventPresenterable, requestConfiguration *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder when successful
func (m *VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder) {
    return NewVirtualEventsTownhallsItemPresentersVirtualEventPresenterItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
