package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdminWindowsUpdatesRequestBuilder provides operations to manage the updates property of the microsoft.graph.windowsUpdates.windows entity.
type AdminWindowsUpdatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminWindowsUpdatesRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdminWindowsUpdatesRequestBuilderGetQueryParameters entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
type AdminWindowsUpdatesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdminWindowsUpdatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminWindowsUpdatesRequestBuilderGetQueryParameters
}
// AdminWindowsUpdatesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.windowsUpdates.updates entity.
func (m *AdminWindowsUpdatesRequestBuilder) Catalog()(*AdminWindowsUpdatesCatalogRequestBuilder) {
    return NewAdminWindowsUpdatesCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAdminWindowsUpdatesRequestBuilderInternal instantiates a new UpdatesRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesRequestBuilder) {
    m := &AdminWindowsUpdatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminWindowsUpdatesRequestBuilder instantiates a new UpdatesRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminWindowsUpdatesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property updates for admin
func (m *AdminWindowsUpdatesRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *AdminWindowsUpdatesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property updates in admin
func (m *AdminWindowsUpdatesRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable, requestConfiguration *AdminWindowsUpdatesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property updates for admin
func (m *AdminWindowsUpdatesRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdminWindowsUpdatesRequestBuilderDeleteRequestConfiguration)(error) {
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
// Deployments provides operations to manage the deployments property of the microsoft.graph.windowsUpdates.updates entity.
func (m *AdminWindowsUpdatesRequestBuilder) Deployments()(*AdminWindowsUpdatesDeploymentsRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeploymentsById provides operations to manage the deployments property of the microsoft.graph.windowsUpdates.updates entity.
func (m *AdminWindowsUpdatesRequestBuilder) DeploymentsById(id string)(*AdminWindowsUpdatesDeploymentsDeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deployment%2Did"] = id
    }
    return NewAdminWindowsUpdatesDeploymentsDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *AdminWindowsUpdatesRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminWindowsUpdatesRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable), nil
}
// Patch update the navigation property updates in admin
func (m *AdminWindowsUpdatesRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable, requestConfiguration *AdminWindowsUpdatesRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable), nil
}
// ResourceConnections provides operations to manage the resourceConnections property of the microsoft.graph.windowsUpdates.updates entity.
func (m *AdminWindowsUpdatesRequestBuilder) ResourceConnections()(*AdminWindowsUpdatesResourceConnectionsRequestBuilder) {
    return NewAdminWindowsUpdatesResourceConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceConnectionsById provides operations to manage the resourceConnections property of the microsoft.graph.windowsUpdates.updates entity.
func (m *AdminWindowsUpdatesRequestBuilder) ResourceConnectionsById(id string)(*AdminWindowsUpdatesResourceConnectionsResourceConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceConnection%2Did"] = id
    }
    return NewAdminWindowsUpdatesResourceConnectionsResourceConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdatableAssets provides operations to manage the updatableAssets property of the microsoft.graph.windowsUpdates.updates entity.
func (m *AdminWindowsUpdatesRequestBuilder) UpdatableAssets()(*AdminWindowsUpdatesUpdatableAssetsRequestBuilder) {
    return NewAdminWindowsUpdatesUpdatableAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdatableAssetsById provides operations to manage the updatableAssets property of the microsoft.graph.windowsUpdates.updates entity.
func (m *AdminWindowsUpdatesRequestBuilder) UpdatableAssetsById(id string)(*AdminWindowsUpdatesUpdatableAssetsUpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset%2Did"] = id
    }
    return NewAdminWindowsUpdatesUpdatableAssetsUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
