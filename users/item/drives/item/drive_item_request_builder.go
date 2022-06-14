package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2614d5d0103e56ecf04c6c4fd0f822c8bb4169f41348567e3129b7e490bdcb9a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/activities"
    i487f0b8b175da4af103a9145d96acb6aae9893982d8774d3214a8a384cbe2a11 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/sharedwithme"
    i692dc1e6a84b3d3b12e09ff20f40b589907efa560331c264538bc5f6da13ea56 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following"
    i6d6ca41bea7345d62c2fc6e26c5908930f4248e99145e3f29e99138769e35ea0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root"
    i8fbeebc4eccf2cac5b307c6dad40201406826338e23db27d5b7087fb1c511a33 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list"
    i945b7af17e70a5c153ccf5092279ad0695fca3f162777d31443b1146f7af0d9c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/searchwithq"
    i9db81101a247c2f318c6dc980b57fdc41219ff102a02f4565abf58a7bd25bc02 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special"
    i9e8763d19a0b58b31ba068155ad48ccdd03e35f21e9127fbc08402fba4782f21 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items"
    iab65ab63810b85ddf181f3ceb253e41e673d177c55f9296e466f666d335ba752 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles"
    ie9dfbaaf9ea9c76631f653acff17130b359159e7c2a776a1a54cddc7f6167506 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/recent"
    i0b052b85d2b65a8720c50aace6ba162194917e9b48b1ec21fb26a7b992f49f80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item"
    i0b84d05595dd5ab19113571f05d199f7f8beafb687e2428e652e75b880a8ae37 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/items/item"
    i1710ef6dc790d2d92c9e3a08bd60c6151b7b10afd1f76d824029788ceef4b4d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item"
    i60852f04cfab9b91ef97b49ef3045833e87b151e085eaab1766d4f0badec40b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item"
    ief65e8da1020d322385dee0b5f53ad802d59fc0ca77551c7f83c8a51a474e0ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/activities/item"
)

// DriveItemRequestBuilder provides operations to manage the drives property of the microsoft.graph.user entity.
type DriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DriveItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DriveItemRequestBuilderGetQueryParameters a collection of drives available for this user. Read-only.
type DriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DriveItemRequestBuilderGetQueryParameters
}
// DriveItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *DriveItemRequestBuilder) Activities()(*i2614d5d0103e56ecf04c6c4fd0f822c8bb4169f41348567e3129b7e490bdcb9a.ActivitiesRequestBuilder) {
    return i2614d5d0103e56ecf04c6c4fd0f822c8bb4169f41348567e3129b7e490bdcb9a.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.activities.item collection
func (m *DriveItemRequestBuilder) ActivitiesById(id string)(*ief65e8da1020d322385dee0b5f53ad802d59fc0ca77551c7f83c8a51a474e0ab.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ief65e8da1020d322385dee0b5f53ad802d59fc0ca77551c7f83c8a51a474e0ab.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bundles the bundles property
func (m *DriveItemRequestBuilder) Bundles()(*iab65ab63810b85ddf181f3ceb253e41e673d177c55f9296e466f666d335ba752.BundlesRequestBuilder) {
    return iab65ab63810b85ddf181f3ceb253e41e673d177c55f9296e466f666d335ba752.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item collection
func (m *DriveItemRequestBuilder) BundlesById(id string)(*i0b052b85d2b65a8720c50aace6ba162194917e9b48b1ec21fb26a7b992f49f80.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i0b052b85d2b65a8720c50aace6ba162194917e9b48b1ec21fb26a7b992f49f80.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property drives for users
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property drives for users
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of drives available for this user. Read-only.
func (m *DriveItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a collection of drives available for this user. Read-only.
func (m *DriveItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property drives in users
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property drives in users
func (m *DriveItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property drives for users
func (m *DriveItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property drives for users
func (m *DriveItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DriveItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Following the following property
func (m *DriveItemRequestBuilder) Following()(*i692dc1e6a84b3d3b12e09ff20f40b589907efa560331c264538bc5f6da13ea56.FollowingRequestBuilder) {
    return i692dc1e6a84b3d3b12e09ff20f40b589907efa560331c264538bc5f6da13ea56.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item collection
func (m *DriveItemRequestBuilder) FollowingById(id string)(*i60852f04cfab9b91ef97b49ef3045833e87b151e085eaab1766d4f0badec40b4.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i60852f04cfab9b91ef97b49ef3045833e87b151e085eaab1766d4f0badec40b4.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get a collection of drives available for this user. Read-only.
func (m *DriveItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler a collection of drives available for this user. Read-only.
func (m *DriveItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DriveItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable), nil
}
// Items the items property
func (m *DriveItemRequestBuilder) Items()(*i9e8763d19a0b58b31ba068155ad48ccdd03e35f21e9127fbc08402fba4782f21.ItemsRequestBuilder) {
    return i9e8763d19a0b58b31ba068155ad48ccdd03e35f21e9127fbc08402fba4782f21.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.items.item collection
func (m *DriveItemRequestBuilder) ItemsById(id string)(*i0b84d05595dd5ab19113571f05d199f7f8beafb687e2428e652e75b880a8ae37.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i0b84d05595dd5ab19113571f05d199f7f8beafb687e2428e652e75b880a8ae37.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List the list property
func (m *DriveItemRequestBuilder) List()(*i8fbeebc4eccf2cac5b307c6dad40201406826338e23db27d5b7087fb1c511a33.ListRequestBuilder) {
    return i8fbeebc4eccf2cac5b307c6dad40201406826338e23db27d5b7087fb1c511a33.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in users
func (m *DriveItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property drives in users
func (m *DriveItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Recent provides operations to call the recent method.
func (m *DriveItemRequestBuilder) Recent()(*ie9dfbaaf9ea9c76631f653acff17130b359159e7c2a776a1a54cddc7f6167506.RecentRequestBuilder) {
    return ie9dfbaaf9ea9c76631f653acff17130b359159e7c2a776a1a54cddc7f6167506.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root the root property
func (m *DriveItemRequestBuilder) Root()(*i6d6ca41bea7345d62c2fc6e26c5908930f4248e99145e3f29e99138769e35ea0.RootRequestBuilder) {
    return i6d6ca41bea7345d62c2fc6e26c5908930f4248e99145e3f29e99138769e35ea0.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemRequestBuilder) SearchWithQ(q *string)(*i945b7af17e70a5c153ccf5092279ad0695fca3f162777d31443b1146f7af0d9c.SearchWithQRequestBuilder) {
    return i945b7af17e70a5c153ccf5092279ad0695fca3f162777d31443b1146f7af0d9c.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DriveItemRequestBuilder) SharedWithMe()(*i487f0b8b175da4af103a9145d96acb6aae9893982d8774d3214a8a384cbe2a11.SharedWithMeRequestBuilder) {
    return i487f0b8b175da4af103a9145d96acb6aae9893982d8774d3214a8a384cbe2a11.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special the special property
func (m *DriveItemRequestBuilder) Special()(*i9db81101a247c2f318c6dc980b57fdc41219ff102a02f4565abf58a7bd25bc02.SpecialRequestBuilder) {
    return i9db81101a247c2f318c6dc980b57fdc41219ff102a02f4565abf58a7bd25bc02.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.special.item collection
func (m *DriveItemRequestBuilder) SpecialById(id string)(*i1710ef6dc790d2d92c9e3a08bd60c6151b7b10afd1f76d824029788ceef4b4d6.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i1710ef6dc790d2d92c9e3a08bd60c6151b7b10afd1f76d824029788ceef4b4d6.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
