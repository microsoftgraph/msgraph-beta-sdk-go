package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/children"
    i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/subscriptions"
    i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/thumbnails"
    i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/permissions"
    i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/content"
    i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/versions"
    i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem"
    i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/activities"
    ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/analytics"
    i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/versions/item"
    i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/children/item"
    i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/thumbnails/item"
    i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/subscriptions/item"
    ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/permissions/item"
    ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/activities/item"
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
func (m *RootRequestBuilder) Activities()(*i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309.ActivitiesRequestBuilder) {
    return i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1.AnalyticsRequestBuilder) {
    return ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba.ChildrenRequestBuilder) {
    return i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/root{?%24select,%24expand}";
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
func (m *RootRequestBuilder) Content()(*i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7.ContentRequestBuilder) {
    return i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for drives
func (m *RootRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for drives
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
// CreateGetRequestInformationWithRequestConfiguration the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in drives
func (m *RootRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in drives
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
// DeleteWithResponseHandler delete navigation property root for drives
func (m *RootRequestBuilder) DeleteWithResponseHandler(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property root for drives
func (m *RootRequestBuilder) DeleteWithResponseHandler(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler the root folder of the drive. Read-only.
func (m *RootRequestBuilder) GetWithResponseHandler(requestConfiguration *RootRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the root folder of the drive. Read-only.
func (m *RootRequestBuilder) GetWithResponseHandler(requestConfiguration *RootRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *RootRequestBuilder) ListItem()(*i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d.ListItemRequestBuilder) {
    return i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property root in drives
func (m *RootRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property root in drives
func (m *RootRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *RootRequestBuilder) Permissions()(*i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254.PermissionsRequestBuilder) {
    return i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967.SubscriptionsRequestBuilder) {
    return i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1.ThumbnailsRequestBuilder) {
    return i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490.VersionsRequestBuilder) {
    return i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
