package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i8bdf948a1899289d97159a4e0bc8d9f6971c249cbe6dad1d8519e95720de7ddf "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/edge/internetexplorermode/sitelists/item/sites"
    ic5b0e27b4a498b80902d32b80cfbfe51a21bbea8849d456ad1aaf6aafae1ca5f "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/edge/internetexplorermode/sitelists/item/publish"
    id5dc1b15e1b280bf465aba1b6f34139ff9db7661f72b678d5bdcc2a110ed5fb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/edge/internetexplorermode/sitelists/item/sharedcookies"
    i24a0f6c84c9c10fcefc55f2ca85910f20b0d94b867fefdf9f48ac5da0abc2069 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/edge/internetexplorermode/sitelists/item/sites/item"
    i8c767c123063c2a496691c2e9d9a0654fce29c52fc7bbb0aa0517b37eb2a298d "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/edge/internetexplorermode/sitelists/item/sharedcookies/item"
)

// BrowserSiteListItemRequestBuilder provides operations to manage the siteLists property of the microsoft.graph.internetExplorerMode entity.
type BrowserSiteListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BrowserSiteListItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BrowserSiteListItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BrowserSiteListItemRequestBuilderGetQueryParameters a collection of site lists to support Internet Explorer mode.
type BrowserSiteListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BrowserSiteListItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BrowserSiteListItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BrowserSiteListItemRequestBuilderGetQueryParameters
}
// BrowserSiteListItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BrowserSiteListItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBrowserSiteListItemRequestBuilderInternal instantiates a new BrowserSiteListItemRequestBuilder and sets the default values.
func NewBrowserSiteListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BrowserSiteListItemRequestBuilder) {
    m := &BrowserSiteListItemRequestBuilder{
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
// NewBrowserSiteListItemRequestBuilder instantiates a new BrowserSiteListItemRequestBuilder and sets the default values.
func NewBrowserSiteListItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BrowserSiteListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBrowserSiteListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property siteLists for admin
func (m *BrowserSiteListItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *BrowserSiteListItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *BrowserSiteListItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *BrowserSiteListItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *BrowserSiteListItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, requestConfiguration *BrowserSiteListItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *BrowserSiteListItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BrowserSiteListItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *BrowserSiteListItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BrowserSiteListItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, error) {
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
func (m *BrowserSiteListItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, requestConfiguration *BrowserSiteListItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BrowserSiteListable, error) {
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
func (m *BrowserSiteListItemRequestBuilder) Publish()(*ic5b0e27b4a498b80902d32b80cfbfe51a21bbea8849d456ad1aaf6aafae1ca5f.PublishRequestBuilder) {
    return ic5b0e27b4a498b80902d32b80cfbfe51a21bbea8849d456ad1aaf6aafae1ca5f.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedCookies provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
func (m *BrowserSiteListItemRequestBuilder) SharedCookies()(*id5dc1b15e1b280bf465aba1b6f34139ff9db7661f72b678d5bdcc2a110ed5fb2.SharedCookiesRequestBuilder) {
    return id5dc1b15e1b280bf465aba1b6f34139ff9db7661f72b678d5bdcc2a110ed5fb2.NewSharedCookiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedCookiesById provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
func (m *BrowserSiteListItemRequestBuilder) SharedCookiesById(id string)(*i8c767c123063c2a496691c2e9d9a0654fce29c52fc7bbb0aa0517b37eb2a298d.BrowserSharedCookieItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["browserSharedCookie%2Did"] = id
    }
    return i8c767c123063c2a496691c2e9d9a0654fce29c52fc7bbb0aa0517b37eb2a298d.NewBrowserSharedCookieItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
func (m *BrowserSiteListItemRequestBuilder) Sites()(*i8bdf948a1899289d97159a4e0bc8d9f6971c249cbe6dad1d8519e95720de7ddf.SitesRequestBuilder) {
    return i8bdf948a1899289d97159a4e0bc8d9f6971c249cbe6dad1d8519e95720de7ddf.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
func (m *BrowserSiteListItemRequestBuilder) SitesById(id string)(*i24a0f6c84c9c10fcefc55f2ca85910f20b0d94b867fefdf9f48ac5da0abc2069.BrowserSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["browserSite%2Did"] = id
    }
    return i24a0f6c84c9c10fcefc55f2ca85910f20b0d94b867fefdf9f48ac5da0abc2069.NewBrowserSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
