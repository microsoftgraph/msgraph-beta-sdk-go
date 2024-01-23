package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder provides operations to manage the sessions property of the microsoft.graph.virtualEventPresenter entity.
type VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderGetQueryParameters get sessions from solutions
type VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderGetQueryParameters
}
// NewVirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderInternal instantiates a new SessionsWithJoinWebUrlRequestBuilder and sets the default values.
func NewVirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, joinWebUrl *string)(*VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder) {
    m := &VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/presenters/{virtualEventPresenter%2Did}/sessions(joinWebUrl='{joinWebUrl}'){?%24select,%24expand}", pathParameters),
    }
    if joinWebUrl != nil {
        m.BaseRequestBuilder.PathParameters["joinWebUrl"] = *joinWebUrl
    }
    return m
}
// NewVirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder instantiates a new SessionsWithJoinWebUrlRequestBuilder and sets the default values.
func NewVirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get sessions from solutions
func (m *VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable), nil
}
// ToGetRequestInformation get sessions from solutions
func (m *VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder) {
    return NewVirtualEventsEventsItemPresentersItemSessionsWithJoinWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
