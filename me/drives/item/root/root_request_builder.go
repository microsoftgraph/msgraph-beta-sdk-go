package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/subscriptions"
    i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/thumbnails"
    i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/permissions"
    i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/children"
    i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/analytics"
    ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/listitem"
    ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/activities"
    ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/content"
    iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/versions"
    i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/activities/item"
    i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/versions/item"
    i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/permissions/item"
    i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/subscriptions/item"
    id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/thumbnails/item"
    ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/children/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RootRequestBuilderGetQueryParameters
}
// RootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *RootRequestBuilder) Activities()(*ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f.ActivitiesRequestBuilder) {
    return ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9.AnalyticsRequestBuilder) {
    return i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a.ChildrenRequestBuilder) {
    return i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *RootRequestBuilder) Content()(*ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa.ContentRequestBuilder) {
    return ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for me
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for me
func (m *RootRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property root in me
func (m *RootRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in me
func (m *RootRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property root for me
func (m *RootRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property root for me
func (m *RootRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the root folder of the drive. Read-only.
func (m *RootRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RootRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *RootRequestBuilder) ListItem()(*ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c.ListItemRequestBuilder) {
    return ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in me
func (m *RootRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property root in me
func (m *RootRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *RootRequestBuilder) Permissions()(*i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea.PermissionsRequestBuilder) {
    return i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980.SubscriptionsRequestBuilder) {
    return i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25.ThumbnailsRequestBuilder) {
    return i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca.VersionsRequestBuilder) {
    return iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
