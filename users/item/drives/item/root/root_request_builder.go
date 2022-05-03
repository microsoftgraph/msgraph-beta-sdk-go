package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/permissions"
    i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/activities"
    i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/analytics"
    i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/subscriptions"
    i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/thumbnails"
    i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/listitem"
    iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/content"
    ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/children"
    id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/versions"
    i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/thumbnails/item"
    i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/activities/item"
    i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/children/item"
    i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/versions/item"
    i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/subscriptions/item"
    ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/permissions/item"
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
func (m *RootRequestBuilder) Activities()(*i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07.ActivitiesRequestBuilder) {
    return i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d.AnalyticsRequestBuilder) {
    return i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85.ChildrenRequestBuilder) {
    return ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
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
func (m *RootRequestBuilder) Content()(*iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9.ContentRequestBuilder) {
    return iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for users
func (m *RootRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for users
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in users
func (m *RootRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in users
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
// DeleteWithResponseHandler delete navigation property root for users
func (m *RootRequestBuilder) DeleteWithResponseHandler(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property root for users
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
func (m *RootRequestBuilder) ListItem()(*i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1.ListItemRequestBuilder) {
    return i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property root in users
func (m *RootRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property root in users
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
func (m *RootRequestBuilder) Permissions()(*i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851.PermissionsRequestBuilder) {
    return i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d.SubscriptionsRequestBuilder) {
    return i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530.ThumbnailsRequestBuilder) {
    return i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b.VersionsRequestBuilder) {
    return id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
