package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder provides operations to manage the siteLists property of the microsoft.graph.internetExplorerMode entity.
type AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetQueryParameters a collection of site lists to support Internet Explorer mode.
type AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetQueryParameters
}
// AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal instantiates a new BrowserSiteListItemRequestBuilder and sets the default values.
func NewAdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) {
    m := &AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder instantiates a new BrowserSiteListItemRequestBuilder and sets the default values.
func NewAdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property siteLists for admin
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of site lists to support Internet Explorer mode.
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property siteLists in admin
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, requestConfiguration *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property siteLists for admin
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of site lists to support Internet Explorer mode.
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBrowserSiteListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable), nil
}
// Patch update the navigation property siteLists in admin
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, requestConfiguration *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBrowserSiteListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable), nil
}
// Publish provides operations to call the publish method.
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Publish()(*AdminEdgeInternetExplorerModeSiteListsItemPublishRequestBuilder) {
    return NewAdminEdgeInternetExplorerModeSiteListsItemPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedCookies provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) SharedCookies()(*AdminEdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) {
    return NewAdminEdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedCookiesById provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) SharedCookiesById(id string)(*AdminEdgeInternetExplorerModeSiteListsItemSharedCookiesBrowserSharedCookieItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["browserSharedCookie%2Did"] = id
    }
    return NewAdminEdgeInternetExplorerModeSiteListsItemSharedCookiesBrowserSharedCookieItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) Sites()(*AdminEdgeInternetExplorerModeSiteListsItemSitesRequestBuilder) {
    return NewAdminEdgeInternetExplorerModeSiteListsItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
func (m *AdminEdgeInternetExplorerModeSiteListsBrowserSiteListItemRequestBuilder) SitesById(id string)(*AdminEdgeInternetExplorerModeSiteListsItemSitesBrowserSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["browserSite%2Did"] = id
    }
    return NewAdminEdgeInternetExplorerModeSiteListsItemSitesBrowserSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
