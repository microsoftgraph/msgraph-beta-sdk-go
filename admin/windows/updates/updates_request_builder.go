package updates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/catalog"
    i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments"
    id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets"
    i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item"
    ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item"
)

// UpdatesRequestBuilder provides operations to manage the updates property of the microsoft.graph.windowsUpdates.windows entity.
type UpdatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UpdatesRequestBuilderDeleteOptions options for Delete
type UpdatesRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// UpdatesRequestBuilderGetOptions options for Get
type UpdatesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UpdatesRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// UpdatesRequestBuilderGetQueryParameters entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
type UpdatesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UpdatesRequestBuilderPatchOptions options for Patch
type UpdatesRequestBuilderPatchOptions struct {
    // 
    Body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Catalog the catalog property
func (m *UpdatesRequestBuilder) Catalog()(*i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080.CatalogRequestBuilder) {
    return i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080.NewCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUpdatesRequestBuilderInternal instantiates a new UpdatesRequestBuilder and sets the default values.
func NewUpdatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdatesRequestBuilder) {
    m := &UpdatesRequestBuilder{
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
// NewUpdatesRequestBuilder instantiates a new UpdatesRequestBuilder and sets the default values.
func NewUpdatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property updates for admin
func (m *UpdatesRequestBuilder) CreateDeleteRequestInformation(options *UpdatesRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *UpdatesRequestBuilder) CreateGetRequestInformation(options *UpdatesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property updates in admin
func (m *UpdatesRequestBuilder) CreatePatchRequestInformation(options *UpdatesRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property updates for admin
func (m *UpdatesRequestBuilder) Delete(options *UpdatesRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Deployments the deployments property
func (m *UpdatesRequestBuilder) Deployments()(*i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669.DeploymentsRequestBuilder) {
    return i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669.NewDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeploymentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.admin.windows.updates.deployments.item collection
func (m *UpdatesRequestBuilder) DeploymentsById(id string)(*ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b.DeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deployment%2Did"] = id
    }
    return ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b.NewDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *UpdatesRequestBuilder) Get(options *UpdatesRequestBuilderGetOptions)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatesFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Updatesable), nil
}
// Patch update the navigation property updates in admin
func (m *UpdatesRequestBuilder) Patch(options *UpdatesRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UpdatableAssets the updatableAssets property
func (m *UpdatesRequestBuilder) UpdatableAssets()(*id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863.UpdatableAssetsRequestBuilder) {
    return id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863.NewUpdatableAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdatableAssetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.admin.windows.updates.updatableAssets.item collection
func (m *UpdatesRequestBuilder) UpdatableAssetsById(id string)(*i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b.UpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset%2Did"] = id
    }
    return i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b.NewUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
