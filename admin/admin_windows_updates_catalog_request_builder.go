package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdminWindowsUpdatesCatalogRequestBuilder provides operations to manage the catalog property of the microsoft.graph.windowsUpdates.updates entity.
type AdminWindowsUpdatesCatalogRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminWindowsUpdatesCatalogRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesCatalogRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdminWindowsUpdatesCatalogRequestBuilderGetQueryParameters catalog of content that can be approved for deployment by the deployment service. Read-only.
type AdminWindowsUpdatesCatalogRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdminWindowsUpdatesCatalogRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesCatalogRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminWindowsUpdatesCatalogRequestBuilderGetQueryParameters
}
// AdminWindowsUpdatesCatalogRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesCatalogRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdminWindowsUpdatesCatalogRequestBuilderInternal instantiates a new CatalogRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesCatalogRequestBuilder) {
    m := &AdminWindowsUpdatesCatalogRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/catalog{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminWindowsUpdatesCatalogRequestBuilder instantiates a new CatalogRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesCatalogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesCatalogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminWindowsUpdatesCatalogRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property catalog for admin
func (m *AdminWindowsUpdatesCatalogRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesCatalogRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *AdminWindowsUpdatesCatalogRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesCatalogRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property catalog in admin
func (m *AdminWindowsUpdatesCatalogRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Catalogable, requestConfiguration *AdminWindowsUpdatesCatalogRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property catalog for admin
func (m *AdminWindowsUpdatesCatalogRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdminWindowsUpdatesCatalogRequestBuilderDeleteRequestConfiguration)(error) {
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
// Entries provides operations to manage the entries property of the microsoft.graph.windowsUpdates.catalog entity.
func (m *AdminWindowsUpdatesCatalogRequestBuilder) Entries()(*AdminWindowsUpdatesCatalogEntriesRequestBuilder) {
    return NewAdminWindowsUpdatesCatalogEntriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EntriesById provides operations to manage the entries property of the microsoft.graph.windowsUpdates.catalog entity.
func (m *AdminWindowsUpdatesCatalogRequestBuilder) EntriesById(id string)(*AdminWindowsUpdatesCatalogEntriesCatalogEntryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["catalogEntry%2Did"] = id
    }
    return NewAdminWindowsUpdatesCatalogEntriesCatalogEntryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *AdminWindowsUpdatesCatalogRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminWindowsUpdatesCatalogRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Catalogable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Catalogable), nil
}
// Patch update the navigation property catalog in admin
func (m *AdminWindowsUpdatesCatalogRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Catalogable, requestConfiguration *AdminWindowsUpdatesCatalogRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Catalogable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Catalogable), nil
}
