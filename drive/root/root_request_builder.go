package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09a126f4908f567485fc32bb158a1a4b20211b117ad08b9bfc7d26eb9f5451eb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/content"
    i184aef01445a07bacf7eb34262da7f276c260c40192ac823afb65cc4a7633f1a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/analytics"
    i2b695595caf897cadb50aebc5a0f557a1cc4ad81052a1cac0c20769137b6c449 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/permissions"
    ia37b7adeb36356a67b48f2ea7350e35f5c25696a3527d14f974953d783db6caf "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/activities"
    iad9096207b89af0d541fff142c32b69ad17329661f64af59e5ec79321f8a805e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/subscriptions"
    ib695e935397298b0896bd5bea3f2317818e631b3dea7090cab84dbbadfddade9 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/listitem"
    ibe4c9dbbb4f55f58fde5a1c1a98e87d374d3f80083bfeb6dc4bfcad2b85c8247 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/children"
    ibf04211e5706a8fa2d6037c385f02c3728741c1bda1f14d4fdb076dc797038e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/thumbnails"
    iea718b4670a92f7fa645e61304ac21dd5e1759f54b3afa9d4ea12f9bbd3632be "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/versions"
    i3a6d4d1d922aca6539049acf1be4e4ceb74e33d1e4ef5fdbacc65fdfc3aa17a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/children/item"
    i81e065e135c5e13716fa269e31607dd05236975d7f5c49a9a00d5f54c15273e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/permissions/item"
    i8e59f342fa199bace9f2a29bd494ef5b5eef2ba7d4ffe16ffe63cc6571047eff "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/thumbnails/item"
    iafd51b75b6c2dfedbe569009d68f20b2a000ae6cbaf10b0895d4ce579d42bbe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/activities/item"
    ib6b7d756f8f85b8dd43fd8babecc7cacaa631bf86762453c1fa457a19e388b6b "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/versions/item"
    ifb9212e266c7381fa46df1045f3a949ee177f2cb7ddd4cd50628cbe0b09f0fa6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root/subscriptions/item"
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
// RootRequestBuilderDeleteOptions options for Delete
type RootRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// RootRequestBuilderGetOptions options for Get
type RootRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RootRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RootRequestBuilderPatchOptions options for Patch
type RootRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Activities the activities property
func (m *RootRequestBuilder) Activities()(*ia37b7adeb36356a67b48f2ea7350e35f5c25696a3527d14f974953d783db6caf.ActivitiesRequestBuilder) {
    return ia37b7adeb36356a67b48f2ea7350e35f5c25696a3527d14f974953d783db6caf.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*iafd51b75b6c2dfedbe569009d68f20b2a000ae6cbaf10b0895d4ce579d42bbe5.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return iafd51b75b6c2dfedbe569009d68f20b2a000ae6cbaf10b0895d4ce579d42bbe5.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i184aef01445a07bacf7eb34262da7f276c260c40192ac823afb65cc4a7633f1a.AnalyticsRequestBuilder) {
    return i184aef01445a07bacf7eb34262da7f276c260c40192ac823afb65cc4a7633f1a.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*ibe4c9dbbb4f55f58fde5a1c1a98e87d374d3f80083bfeb6dc4bfcad2b85c8247.ChildrenRequestBuilder) {
    return ibe4c9dbbb4f55f58fde5a1c1a98e87d374d3f80083bfeb6dc4bfcad2b85c8247.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i3a6d4d1d922aca6539049acf1be4e4ceb74e33d1e4ef5fdbacc65fdfc3aa17a4.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i3a6d4d1d922aca6539049acf1be4e4ceb74e33d1e4ef5fdbacc65fdfc3aa17a4.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/root{?%24select,%24expand}";
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
func (m *RootRequestBuilder) Content()(*i09a126f4908f567485fc32bb158a1a4b20211b117ad08b9bfc7d26eb9f5451eb.ContentRequestBuilder) {
    return i09a126f4908f567485fc32bb158a1a4b20211b117ad08b9bfc7d26eb9f5451eb.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for drive
func (m *RootRequestBuilder) CreateDeleteRequestInformation(options *RootRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation(options *RootRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property root in drive
func (m *RootRequestBuilder) CreatePatchRequestInformation(options *RootRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property root for drive
func (m *RootRequestBuilder) Delete(options *RootRequestBuilderDeleteOptions)(error) {
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
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get(options *RootRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*ib695e935397298b0896bd5bea3f2317818e631b3dea7090cab84dbbadfddade9.ListItemRequestBuilder) {
    return ib695e935397298b0896bd5bea3f2317818e631b3dea7090cab84dbbadfddade9.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in drive
func (m *RootRequestBuilder) Patch(options *RootRequestBuilderPatchOptions)(error) {
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
// Permissions the permissions property
func (m *RootRequestBuilder) Permissions()(*i2b695595caf897cadb50aebc5a0f557a1cc4ad81052a1cac0c20769137b6c449.PermissionsRequestBuilder) {
    return i2b695595caf897cadb50aebc5a0f557a1cc4ad81052a1cac0c20769137b6c449.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i81e065e135c5e13716fa269e31607dd05236975d7f5c49a9a00d5f54c15273e8.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i81e065e135c5e13716fa269e31607dd05236975d7f5c49a9a00d5f54c15273e8.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*iad9096207b89af0d541fff142c32b69ad17329661f64af59e5ec79321f8a805e.SubscriptionsRequestBuilder) {
    return iad9096207b89af0d541fff142c32b69ad17329661f64af59e5ec79321f8a805e.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*ifb9212e266c7381fa46df1045f3a949ee177f2cb7ddd4cd50628cbe0b09f0fa6.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ifb9212e266c7381fa46df1045f3a949ee177f2cb7ddd4cd50628cbe0b09f0fa6.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*ibf04211e5706a8fa2d6037c385f02c3728741c1bda1f14d4fdb076dc797038e0.ThumbnailsRequestBuilder) {
    return ibf04211e5706a8fa2d6037c385f02c3728741c1bda1f14d4fdb076dc797038e0.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i8e59f342fa199bace9f2a29bd494ef5b5eef2ba7d4ffe16ffe63cc6571047eff.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i8e59f342fa199bace9f2a29bd494ef5b5eef2ba7d4ffe16ffe63cc6571047eff.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*iea718b4670a92f7fa645e61304ac21dd5e1759f54b3afa9d4ea12f9bbd3632be.VersionsRequestBuilder) {
    return iea718b4670a92f7fa645e61304ac21dd5e1759f54b3afa9d4ea12f9bbd3632be.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*ib6b7d756f8f85b8dd43fd8babecc7cacaa631bf86762453c1fa457a19e388b6b.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return ib6b7d756f8f85b8dd43fd8babecc7cacaa631bf86762453c1fa457a19e388b6b.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
