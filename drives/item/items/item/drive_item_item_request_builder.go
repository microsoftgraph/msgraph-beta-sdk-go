package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2192a3e160b84ecba9604063909cf5ffaa935f21179db85377f12faa9bfefe21 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/listitem"
    i2b1899d1328125e0a77e9a9689526a1e68e60f25b055a3d8dc339cd92f8f4c9f "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/thumbnails"
    i2db7da3b0d8ecabb8da814ea2dce35d6747e7abdbe8eb1d9bf691b3a286bb748 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/activities"
    i8104738c0867acea61b09341148ceb8c1ba8e8feba80d4f22961b90cdf61d068 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/versions"
    ib45834c2958675f19eed8470552cd0de3fd39cfc19fa4f66ec2dd86532af6004 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/permissions"
    ibbaf95ebb8b2c19c17c02a61a20827f061096731272353f476d3016f130f94c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/analytics"
    ic3c82b3e3aa4c99dd8bd9720303fa07cd398152ceb551f726c583832e8058410 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/content"
    icf8dc0aeed4d322b794a311185fc330019da0eb089fded259534290fa691efa5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/children"
    id156ba179ca21bc9f80b635a70caebe7575117940bbe10acc18bc0f7eb713a02 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/subscriptions"
    i0db2f2eb073ed9a6c54c1ff2536a0ad7aefaaa4c9d3d2eb7c9602d9278bdfb4a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/versions/item"
    i0e7818f4889eec436f3bc22f1b52293cc43bdf6ac39a914954617708bff2ae3f "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/activities/item"
    i25423134922b3b1531df0f8f7c7fa9ef48f5b4996405c15d4398ee30f2e58ea2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/subscriptions/item"
    i4e42914cc521bccb391e2aac7c82cb87958c98b64eaf603e8dd1d316370819f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/thumbnails/item"
    i78d8c1a174e59b42b1fe5def2e27f453c23ad81f5897c7370b6978e35e91329c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/permissions/item"
    i9024a536564adce7d0824c7a902d876ac948990b9f4cc7b309c7bbaedd90750c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item/children/item"
)

// DriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DriveItemItemRequestBuilderGetQueryParameters
}
// DriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *DriveItemItemRequestBuilder) Activities()(*i2db7da3b0d8ecabb8da814ea2dce35d6747e7abdbe8eb1d9bf691b3a286bb748.ActivitiesRequestBuilder) {
    return i2db7da3b0d8ecabb8da814ea2dce35d6747e7abdbe8eb1d9bf691b3a286bb748.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.items.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i0e7818f4889eec436f3bc22f1b52293cc43bdf6ac39a914954617708bff2ae3f.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i0e7818f4889eec436f3bc22f1b52293cc43bdf6ac39a914954617708bff2ae3f.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*ibbaf95ebb8b2c19c17c02a61a20827f061096731272353f476d3016f130f94c7.AnalyticsRequestBuilder) {
    return ibbaf95ebb8b2c19c17c02a61a20827f061096731272353f476d3016f130f94c7.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*icf8dc0aeed4d322b794a311185fc330019da0eb089fded259534290fa691efa5.ChildrenRequestBuilder) {
    return icf8dc0aeed4d322b794a311185fc330019da0eb089fded259534290fa691efa5.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i9024a536564adce7d0824c7a902d876ac948990b9f4cc7b309c7bbaedd90750c.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return i9024a536564adce7d0824c7a902d876ac948990b9f4cc7b309c7bbaedd90750c.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *DriveItemItemRequestBuilder) Content()(*ic3c82b3e3aa4c99dd8bd9720303fa07cd398152ceb551f726c583832e8058410.ContentRequestBuilder) {
    return ic3c82b3e3aa4c99dd8bd9720303fa07cd398152ceb551f726c583832e8058410.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property items for drives
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property items for drives
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property items in drives
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property items in drives
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property items for drives
func (m *DriveItemItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property items for drives
func (m *DriveItemItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) GetWithResponseHandler(requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) GetWithResponseHandler(requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// ListItem the listItem property
func (m *DriveItemItemRequestBuilder) ListItem()(*i2192a3e160b84ecba9604063909cf5ffaa935f21179db85377f12faa9bfefe21.ListItemRequestBuilder) {
    return i2192a3e160b84ecba9604063909cf5ffaa935f21179db85377f12faa9bfefe21.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property items in drives
func (m *DriveItemItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property items in drives
func (m *DriveItemItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Permissions the permissions property
func (m *DriveItemItemRequestBuilder) Permissions()(*ib45834c2958675f19eed8470552cd0de3fd39cfc19fa4f66ec2dd86532af6004.PermissionsRequestBuilder) {
    return ib45834c2958675f19eed8470552cd0de3fd39cfc19fa4f66ec2dd86532af6004.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i78d8c1a174e59b42b1fe5def2e27f453c23ad81f5897c7370b6978e35e91329c.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i78d8c1a174e59b42b1fe5def2e27f453c23ad81f5897c7370b6978e35e91329c.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*id156ba179ca21bc9f80b635a70caebe7575117940bbe10acc18bc0f7eb713a02.SubscriptionsRequestBuilder) {
    return id156ba179ca21bc9f80b635a70caebe7575117940bbe10acc18bc0f7eb713a02.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i25423134922b3b1531df0f8f7c7fa9ef48f5b4996405c15d4398ee30f2e58ea2.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i25423134922b3b1531df0f8f7c7fa9ef48f5b4996405c15d4398ee30f2e58ea2.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i2b1899d1328125e0a77e9a9689526a1e68e60f25b055a3d8dc339cd92f8f4c9f.ThumbnailsRequestBuilder) {
    return i2b1899d1328125e0a77e9a9689526a1e68e60f25b055a3d8dc339cd92f8f4c9f.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i4e42914cc521bccb391e2aac7c82cb87958c98b64eaf603e8dd1d316370819f7.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i4e42914cc521bccb391e2aac7c82cb87958c98b64eaf603e8dd1d316370819f7.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i8104738c0867acea61b09341148ceb8c1ba8e8feba80d4f22961b90cdf61d068.VersionsRequestBuilder) {
    return i8104738c0867acea61b09341148ceb8c1ba8e8feba80d4f22961b90cdf61d068.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i0db2f2eb073ed9a6c54c1ff2536a0ad7aefaaa4c9d3d2eb7c9602d9278bdfb4a.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i0db2f2eb073ed9a6c54c1ff2536a0ad7aefaaa4c9d3d2eb7c9602d9278bdfb4a.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
