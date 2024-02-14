package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder provides operations to manage the presenters property of the microsoft.graph.virtualEventSession entity.
type VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderGetQueryParameters get presenters from solutions
type VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderGetQueryParameters struct {
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
// VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderGetQueryParameters
}
// ByVirtualEventPresenterId provides operations to manage the presenters property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsTownhallsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) ByVirtualEventPresenterId(virtualEventPresenterId string)(*VirtualEventsTownhallsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if virtualEventPresenterId != "" {
        urlTplParams["virtualEventPresenter%2Did"] = virtualEventPresenterId
    }
    return NewVirtualEventsTownhallsItemSessionsItemPresentersVirtualEventPresenterItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderInternal instantiates a new VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) {
    m := &VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/sessions/{virtualEventSession%2Did}/presenters{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder instantiates a new VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualEventsTownhallsItemSessionsItemPresentersCountRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) Count()(*VirtualEventsTownhallsItemSessionsItemPresentersCountRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemPresentersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get presenters from solutions
// returns a VirtualEventPresenterCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventPresenterCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventPresenterCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventPresenterCollectionResponseable), nil
}
// ToGetRequestInformation get presenters from solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
