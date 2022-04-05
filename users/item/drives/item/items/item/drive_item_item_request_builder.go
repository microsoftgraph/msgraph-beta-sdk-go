package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03fda11e671eb36bee13e25e3e9b7862084dc785a14f093a05e62df80681796f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/content"
    i23f743e3e59d3542f75db30e3b39223f6f1c3f485139a340010e208a5d7f562b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/versions"
    i5efcbe4b461c812355323caa9dc418905255f803cf1cf6260dcc892a8b9ed564 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/children"
    i7021ce989d3be4c224e26d0dbaa6136089065fa4c7c08b6f8d132f6d6588c963 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/analytics"
    i914d35933988d4207bb270b7495822ddb2cdff917df7876985c635fc1775a412 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/thumbnails"
    ia3bb93bc9befb9c2bcebd5daca5f25eb48c4c0564a67fc189f5a53eea0bdf60c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/listitem"
    idbe8275856d8fedd2fa248418909c5b36129b3c0b7154d2c961811b3cf35e829 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/activities"
    idf0b80a60e82ebe95e49f8a72b4fdaff6d5011f039b985bdaa81f830d809be12 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/permissions"
    if15a9ba7729fdebc3c57a8ca44847546466efb149c4df186ff7d8cf6395e1c35 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/subscriptions"
    i4842cc54fffeac38d470c14c899c0d0e0fddefe6bc8020310bfa6c3a7477bc1f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/subscriptions/item"
    i7310b5803cdc10efcb7c5025f842e7b6c7175de4f87b0ff20422795e648d9a26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/activities/item"
    i9a2829926ca071ddf884f456479b8ccc48890a1b9bcfaec84bf1efe1d866494d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/versions/item"
    i9d26c5ed3e1dc52c55976335bc31cd3e38d3224ac6417512d4910d03c1776584 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/thumbnails/item"
    ide9d84689ef4976f2fb708443e6cf9f693bf4dc8a4301d5ab6caf47d2aec7e07 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/permissions/item"
    iefaf12d1fdb776bb7078a229349c10b5181934746d31a87a8236cbed30664968 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item/children/item"
)

// DriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemItemRequestBuilderDeleteOptions options for Delete
type DriveItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DriveItemItemRequestBuilderGetOptions options for Get
type DriveItemItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DriveItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemItemRequestBuilderPatchOptions options for Patch
type DriveItemItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Activities the activities property
func (m *DriveItemItemRequestBuilder) Activities()(*idbe8275856d8fedd2fa248418909c5b36129b3c0b7154d2c961811b3cf35e829.ActivitiesRequestBuilder) {
    return idbe8275856d8fedd2fa248418909c5b36129b3c0b7154d2c961811b3cf35e829.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i7310b5803cdc10efcb7c5025f842e7b6c7175de4f87b0ff20422795e648d9a26.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i7310b5803cdc10efcb7c5025f842e7b6c7175de4f87b0ff20422795e648d9a26.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *DriveItemItemRequestBuilder) Analytics()(*i7021ce989d3be4c224e26d0dbaa6136089065fa4c7c08b6f8d132f6d6588c963.AnalyticsRequestBuilder) {
    return i7021ce989d3be4c224e26d0dbaa6136089065fa4c7c08b6f8d132f6d6588c963.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *DriveItemItemRequestBuilder) Children()(*i5efcbe4b461c812355323caa9dc418905255f803cf1cf6260dcc892a8b9ed564.ChildrenRequestBuilder) {
    return i5efcbe4b461c812355323caa9dc418905255f803cf1cf6260dcc892a8b9ed564.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*iefaf12d1fdb776bb7078a229349c10b5181934746d31a87a8236cbed30664968.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return iefaf12d1fdb776bb7078a229349c10b5181934746d31a87a8236cbed30664968.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/items/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i03fda11e671eb36bee13e25e3e9b7862084dc785a14f093a05e62df80681796f.ContentRequestBuilder) {
    return i03fda11e671eb36bee13e25e3e9b7862084dc785a14f093a05e62df80681796f.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for users
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation(options *DriveItemItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property items in users
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property items for users
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
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
// Get all items contained in the drive. Read-only. Nullable.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *DriveItemItemRequestBuilder) ListItem()(*ia3bb93bc9befb9c2bcebd5daca5f25eb48c4c0564a67fc189f5a53eea0bdf60c.ListItemRequestBuilder) {
    return ia3bb93bc9befb9c2bcebd5daca5f25eb48c4c0564a67fc189f5a53eea0bdf60c.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in users
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*idf0b80a60e82ebe95e49f8a72b4fdaff6d5011f039b985bdaa81f830d809be12.PermissionsRequestBuilder) {
    return idf0b80a60e82ebe95e49f8a72b4fdaff6d5011f039b985bdaa81f830d809be12.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ide9d84689ef4976f2fb708443e6cf9f693bf4dc8a4301d5ab6caf47d2aec7e07.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ide9d84689ef4976f2fb708443e6cf9f693bf4dc8a4301d5ab6caf47d2aec7e07.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *DriveItemItemRequestBuilder) Subscriptions()(*if15a9ba7729fdebc3c57a8ca44847546466efb149c4df186ff7d8cf6395e1c35.SubscriptionsRequestBuilder) {
    return if15a9ba7729fdebc3c57a8ca44847546466efb149c4df186ff7d8cf6395e1c35.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i4842cc54fffeac38d470c14c899c0d0e0fddefe6bc8020310bfa6c3a7477bc1f.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i4842cc54fffeac38d470c14c899c0d0e0fddefe6bc8020310bfa6c3a7477bc1f.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i914d35933988d4207bb270b7495822ddb2cdff917df7876985c635fc1775a412.ThumbnailsRequestBuilder) {
    return i914d35933988d4207bb270b7495822ddb2cdff917df7876985c635fc1775a412.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i9d26c5ed3e1dc52c55976335bc31cd3e38d3224ac6417512d4910d03c1776584.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i9d26c5ed3e1dc52c55976335bc31cd3e38d3224ac6417512d4910d03c1776584.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *DriveItemItemRequestBuilder) Versions()(*i23f743e3e59d3542f75db30e3b39223f6f1c3f485139a340010e208a5d7f562b.VersionsRequestBuilder) {
    return i23f743e3e59d3542f75db30e3b39223f6f1c3f485139a340010e208a5d7f562b.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i9a2829926ca071ddf884f456479b8ccc48890a1b9bcfaec84bf1efe1d866494d.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i9a2829926ca071ddf884f456479b8ccc48890a1b9bcfaec84bf1efe1d866494d.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
