package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i37ff40e2aa7be9eca0148887400d7b085d8f3de83bf87492286dc92f5a7ac5eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/analytics"
    i39765dfa1c3eafd523949b6b410f599c0e93968da8276a68b09d6869ce0a7fc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/thumbnails"
    i43351f457508aecc9acc29067e433332ebc46c9941ad1ac6e1a5f33f8fcb1e5d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/children"
    i52fdb84788de6f4295e35bd2351366c8c5b985d4d378bc90530d073b39223a0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/subscriptions"
    i5b0a99c5b67d423f45d52d7551bcbc2389478eaa8636e90fc04fde1312b36353 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/content"
    i5f7f4ee0090cc95d75d7c261819f4cc11af67ee4cf962b90e283ca54b09a795b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/permissions"
    i606f41c5fea6c70fd34ccde609068e6bd09206e372d836cc7a3cd47d661d3be8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/listitem"
    i6f3eef764ea1ae21b6822debfa6f6a704cd6435b66bcebeb13a675ec3815bd92 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/versions"
    if07779cafb32bcb954e8baa2fcdb6283357a647442af0a49aaa2643ddf7a2125 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/activities"
    i73ce730b8b40e7e732bae3aa850e526852f4bbcfac2c62b18d895fb8c40cb028 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/children/item"
    i8b02ac8460c4300fef83eeacf0e3a90fa6ad7e02a0ad5f87838abb6d5fb2ce82 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/subscriptions/item"
    ia7e12ef2cde2e1cee53c4ab5522f78d82906572a22510a1b2961ef668cc7a611 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/versions/item"
    ida01d8a4830793c46212039d68a6eefd4dfcee0c4bbfeb15920c2e1eeba59306 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/activities/item"
    ie01afcc24bcb8de9f524e134f72f107f5d9980714d4b3f83b8b0a3d66372a5c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/permissions/item"
    ie8883816406be65444fb3f248da99471dc749778a92112ecb2d943acde1ab3bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item/thumbnails/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*if07779cafb32bcb954e8baa2fcdb6283357a647442af0a49aaa2643ddf7a2125.ActivitiesRequestBuilder) {
    return if07779cafb32bcb954e8baa2fcdb6283357a647442af0a49aaa2643ddf7a2125.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*ida01d8a4830793c46212039d68a6eefd4dfcee0c4bbfeb15920c2e1eeba59306.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ida01d8a4830793c46212039d68a6eefd4dfcee0c4bbfeb15920c2e1eeba59306.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*i37ff40e2aa7be9eca0148887400d7b085d8f3de83bf87492286dc92f5a7ac5eb.AnalyticsRequestBuilder) {
    return i37ff40e2aa7be9eca0148887400d7b085d8f3de83bf87492286dc92f5a7ac5eb.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*i43351f457508aecc9acc29067e433332ebc46c9941ad1ac6e1a5f33f8fcb1e5d.ChildrenRequestBuilder) {
    return i43351f457508aecc9acc29067e433332ebc46c9941ad1ac6e1a5f33f8fcb1e5d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i73ce730b8b40e7e732bae3aa850e526852f4bbcfac2c62b18d895fb8c40cb028.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return i73ce730b8b40e7e732bae3aa850e526852f4bbcfac2c62b18d895fb8c40cb028.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i5b0a99c5b67d423f45d52d7551bcbc2389478eaa8636e90fc04fde1312b36353.ContentRequestBuilder) {
    return i5b0a99c5b67d423f45d52d7551bcbc2389478eaa8636e90fc04fde1312b36353.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for me
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property items for me
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
// CreateGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property items in me
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property items in me
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
// Delete delete navigation property items for me
func (m *DriveItemItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property items for me
func (m *DriveItemItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DriveItemItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DriveItemItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i606f41c5fea6c70fd34ccde609068e6bd09206e372d836cc7a3cd47d661d3be8.ListItemRequestBuilder) {
    return i606f41c5fea6c70fd34ccde609068e6bd09206e372d836cc7a3cd47d661d3be8.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in me
func (m *DriveItemItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property items in me
func (m *DriveItemItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DriveItemItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i5f7f4ee0090cc95d75d7c261819f4cc11af67ee4cf962b90e283ca54b09a795b.PermissionsRequestBuilder) {
    return i5f7f4ee0090cc95d75d7c261819f4cc11af67ee4cf962b90e283ca54b09a795b.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ie01afcc24bcb8de9f524e134f72f107f5d9980714d4b3f83b8b0a3d66372a5c4.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return ie01afcc24bcb8de9f524e134f72f107f5d9980714d4b3f83b8b0a3d66372a5c4.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i52fdb84788de6f4295e35bd2351366c8c5b985d4d378bc90530d073b39223a0c.SubscriptionsRequestBuilder) {
    return i52fdb84788de6f4295e35bd2351366c8c5b985d4d378bc90530d073b39223a0c.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i8b02ac8460c4300fef83eeacf0e3a90fa6ad7e02a0ad5f87838abb6d5fb2ce82.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i8b02ac8460c4300fef83eeacf0e3a90fa6ad7e02a0ad5f87838abb6d5fb2ce82.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i39765dfa1c3eafd523949b6b410f599c0e93968da8276a68b09d6869ce0a7fc8.ThumbnailsRequestBuilder) {
    return i39765dfa1c3eafd523949b6b410f599c0e93968da8276a68b09d6869ce0a7fc8.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ie8883816406be65444fb3f248da99471dc749778a92112ecb2d943acde1ab3bc.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return ie8883816406be65444fb3f248da99471dc749778a92112ecb2d943acde1ab3bc.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i6f3eef764ea1ae21b6822debfa6f6a704cd6435b66bcebeb13a675ec3815bd92.VersionsRequestBuilder) {
    return i6f3eef764ea1ae21b6822debfa6f6a704cd6435b66bcebeb13a675ec3815bd92.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ia7e12ef2cde2e1cee53c4ab5522f78d82906572a22510a1b2961ef668cc7a611.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return ia7e12ef2cde2e1cee53c4ab5522f78d82906572a22510a1b2961ef668cc7a611.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
