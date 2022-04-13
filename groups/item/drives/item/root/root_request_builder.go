package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0038b1d7c66ec87b720bb6f335bf260d01e756277ebb6253a9675ba71f125eeb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/thumbnails"
    i3ee022e362928301f95e98536126ae08be259ddb3d74fcc35882ec3d6f209cc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/content"
    i59e8bbfbd3f60bae567a8408523b60e2da6b7989b7e672dd97a3c767fd417f23 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/subscriptions"
    i5ce3ec157cfddb136d0d349fb3cecea590a20e228fd29509b7b1497c1f04a2ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/listitem"
    i6ec1534d2a31fed146899b9d06252a57f3682a9d852217105bf67ef0be0d3e10 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/children"
    i82154dd857d1f4147284ffd2c005f1cb11a2e6dba3e0d970db99c39f04c4ad3d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/activities"
    i8e06844303b316afa4bb273c79b3822aa8583f64316c2c08e3e3fa351d53c86e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/versions"
    id97719380bf38dd33a04dc42f078677496b6a3e8aab5f3da8752ffd389d429e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/analytics"
    idb035671a49671c9c3860c6e20e520df498ee814eaa53d85eba7bc03cdc024ed "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/permissions"
    i34c79bd133f2c52bf82779507f18904f43c5de372db9b004de5a4f83505809bf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/children/item"
    i532e3abc360ac6f1fa4f890bd2c94e559c2d68325145fc45dba0b117dcc9b68e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/versions/item"
    i7d0a17f40953bcd7090a60e2809adcb724cfcd3db5d588342f80400d3b778d2c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/subscriptions/item"
    i83b0f9503a085d5076598d259da4b52da18ce21d577b876440a3993997518ab9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/permissions/item"
    ib24f9a2a2766fa9430ff8f21a5a0491c4493d94f0d2bef85b623278ad14433ef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/thumbnails/item"
    ie651bb323681d0ad0e4318a76c4128cd5d1eac694c815fdcf3388b964126ff0d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root/activities/item"
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
func (m *RootRequestBuilder) Activities()(*i82154dd857d1f4147284ffd2c005f1cb11a2e6dba3e0d970db99c39f04c4ad3d.ActivitiesRequestBuilder) {
    return i82154dd857d1f4147284ffd2c005f1cb11a2e6dba3e0d970db99c39f04c4ad3d.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*ie651bb323681d0ad0e4318a76c4128cd5d1eac694c815fdcf3388b964126ff0d.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ie651bb323681d0ad0e4318a76c4128cd5d1eac694c815fdcf3388b964126ff0d.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*id97719380bf38dd33a04dc42f078677496b6a3e8aab5f3da8752ffd389d429e8.AnalyticsRequestBuilder) {
    return id97719380bf38dd33a04dc42f078677496b6a3e8aab5f3da8752ffd389d429e8.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i6ec1534d2a31fed146899b9d06252a57f3682a9d852217105bf67ef0be0d3e10.ChildrenRequestBuilder) {
    return i6ec1534d2a31fed146899b9d06252a57f3682a9d852217105bf67ef0be0d3e10.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i34c79bd133f2c52bf82779507f18904f43c5de372db9b004de5a4f83505809bf.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i34c79bd133f2c52bf82779507f18904f43c5de372db9b004de5a4f83505809bf.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
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
func (m *RootRequestBuilder) Content()(*i3ee022e362928301f95e98536126ae08be259ddb3d74fcc35882ec3d6f209cc7.ContentRequestBuilder) {
    return i3ee022e362928301f95e98536126ae08be259ddb3d74fcc35882ec3d6f209cc7.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for groups
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
// CreatePatchRequestInformation update the navigation property root in groups
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
// Delete delete navigation property root for groups
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
func (m *RootRequestBuilder) ListItem()(*i5ce3ec157cfddb136d0d349fb3cecea590a20e228fd29509b7b1497c1f04a2ea.ListItemRequestBuilder) {
    return i5ce3ec157cfddb136d0d349fb3cecea590a20e228fd29509b7b1497c1f04a2ea.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in groups
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
func (m *RootRequestBuilder) Permissions()(*idb035671a49671c9c3860c6e20e520df498ee814eaa53d85eba7bc03cdc024ed.PermissionsRequestBuilder) {
    return idb035671a49671c9c3860c6e20e520df498ee814eaa53d85eba7bc03cdc024ed.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i83b0f9503a085d5076598d259da4b52da18ce21d577b876440a3993997518ab9.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i83b0f9503a085d5076598d259da4b52da18ce21d577b876440a3993997518ab9.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i59e8bbfbd3f60bae567a8408523b60e2da6b7989b7e672dd97a3c767fd417f23.SubscriptionsRequestBuilder) {
    return i59e8bbfbd3f60bae567a8408523b60e2da6b7989b7e672dd97a3c767fd417f23.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i7d0a17f40953bcd7090a60e2809adcb724cfcd3db5d588342f80400d3b778d2c.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i7d0a17f40953bcd7090a60e2809adcb724cfcd3db5d588342f80400d3b778d2c.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i0038b1d7c66ec87b720bb6f335bf260d01e756277ebb6253a9675ba71f125eeb.ThumbnailsRequestBuilder) {
    return i0038b1d7c66ec87b720bb6f335bf260d01e756277ebb6253a9675ba71f125eeb.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*ib24f9a2a2766fa9430ff8f21a5a0491c4493d94f0d2bef85b623278ad14433ef.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return ib24f9a2a2766fa9430ff8f21a5a0491c4493d94f0d2bef85b623278ad14433ef.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*i8e06844303b316afa4bb273c79b3822aa8583f64316c2c08e3e3fa351d53c86e.VersionsRequestBuilder) {
    return i8e06844303b316afa4bb273c79b3822aa8583f64316c2c08e3e3fa351d53c86e.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i532e3abc360ac6f1fa4f890bd2c94e559c2d68325145fc45dba0b117dcc9b68e.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i532e3abc360ac6f1fa4f890bd2c94e559c2d68325145fc45dba0b117dcc9b68e.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
