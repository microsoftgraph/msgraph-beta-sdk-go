package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder provides operations to manage the siteLists property of the microsoft.graph.internetExplorerMode entity.
type EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetQueryParameters a collection of site lists to support Internet Explorer mode.
type EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetQueryParameters
}
// EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal instantiates a new EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) {
    m := &EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder instantiates a new EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property siteLists for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of site lists to support Internet Explorer mode.
// returns a BrowserSiteListable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBrowserSiteListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable), nil
}
// Patch update the navigation property siteLists in admin
// returns a BrowserSiteListable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, requestConfiguration *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBrowserSiteListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable), nil
}
// Publish provides operations to call the publish method.
// returns a *EdgeInternetExplorerModeSiteListsItemPublishRequestBuilder when successful
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Publish()(*EdgeInternetExplorerModeSiteListsItemPublishRequestBuilder) {
    return NewEdgeInternetExplorerModeSiteListsItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedCookies provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
// returns a *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder when successful
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) SharedCookies()(*EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) {
    return NewEdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
// returns a *EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder when successful
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Sites()(*EdgeInternetExplorerModeSiteListsItemSitesRequestBuilder) {
    return NewEdgeInternetExplorerModeSiteListsItemSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property siteLists for admin
// returns a *RequestInformation when successful
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of site lists to support Internet Explorer mode.
// returns a *RequestInformation when successful
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property siteLists in admin
// returns a *RequestInformation when successful
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, requestConfiguration *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder when successful
func (m *EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) WithUrl(rawUrl string)(*EdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) {
    return NewEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
