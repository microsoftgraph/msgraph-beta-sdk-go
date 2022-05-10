package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09d4dc851d5953d58187c74569e57c45eeadd608ce6d5cda6ab399cde2d1fc68 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/activities"
    i1f3a6e9c86e5b034c8fa45efe2ab068d31ba24e081da4647a4cf7f622bb330d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root"
    i354819977222ec25ccb855b2c65e2a34db15f6cf4aa9ea05b01f0cd9697e3e3b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/bundles"
    i38901b482327dd22606c1b4b91388052e2bf2fc45cb9d926d8f2f8e7d35b85fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/items"
    ia53ba40695c49763c010561b36523e758483b42d2e677bf1f6185d496f24c3f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/following"
    ia83dd16cbcab16b5c7adaabbe5934caf03fa1b23a0441de8c6b962b9aa054f24 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/list"
    ib11ab85aee9d209c297ca809b535a0d82cb670f95bdbdac838c4f09eea952961 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/special"
    i3e2ebdb6db6eb46e901223e1891b32e0bf81d6c6dd0899111b17be6df061afca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/special/item"
    i4c86288525b2d1b5d34556dd506377c0b0e951f12fc731c6a4095614ada62e95 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/bundles/item"
    i7e759696ace6a16bd100c729b8f5e6ec446cb0b0db66ef4433557b5d8af23c15 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/activities/item"
    idfb4790bf624d08616bb3513bffebd055a5dcadbacca095857d461aa5f1a530a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/items/item"
    ifde0edbf6175f3d8478b6b331745f214de6dc67656cbf75cdb2d250f1a2c819f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/following/item"
)

// DriveItemRequestBuilder provides operations to manage the drives property of the microsoft.graph.group entity.
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
// DriveItemRequestBuilderGetQueryParameters the group's drives. Read-only.
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
func (m *DriveItemRequestBuilder) Activities()(*i09d4dc851d5953d58187c74569e57c45eeadd608ce6d5cda6ab399cde2d1fc68.ActivitiesRequestBuilder) {
    return i09d4dc851d5953d58187c74569e57c45eeadd608ce6d5cda6ab399cde2d1fc68.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.activities.item collection
func (m *DriveItemRequestBuilder) ActivitiesById(id string)(*i7e759696ace6a16bd100c729b8f5e6ec446cb0b0db66ef4433557b5d8af23c15.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i7e759696ace6a16bd100c729b8f5e6ec446cb0b0db66ef4433557b5d8af23c15.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bundles the bundles property
func (m *DriveItemRequestBuilder) Bundles()(*i354819977222ec25ccb855b2c65e2a34db15f6cf4aa9ea05b01f0cd9697e3e3b.BundlesRequestBuilder) {
    return i354819977222ec25ccb855b2c65e2a34db15f6cf4aa9ea05b01f0cd9697e3e3b.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.bundles.item collection
func (m *DriveItemRequestBuilder) BundlesById(id string)(*i4c86288525b2d1b5d34556dd506377c0b0e951f12fc731c6a4095614ada62e95.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i4c86288525b2d1b5d34556dd506377c0b0e951f12fc731c6a4095614ada62e95.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/drives/{drive%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation the group's drives. Read-only.
func (m *DriveItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the group's drives. Read-only.
func (m *DriveItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DriveItemRequestBuilder) Following()(*ia53ba40695c49763c010561b36523e758483b42d2e677bf1f6185d496f24c3f3.FollowingRequestBuilder) {
    return ia53ba40695c49763c010561b36523e758483b42d2e677bf1f6185d496f24c3f3.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.following.item collection
func (m *DriveItemRequestBuilder) FollowingById(id string)(*ifde0edbf6175f3d8478b6b331745f214de6dc67656cbf75cdb2d250f1a2c819f.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ifde0edbf6175f3d8478b6b331745f214de6dc67656cbf75cdb2d250f1a2c819f.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the group's drives. Read-only.
func (m *DriveItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the group's drives. Read-only.
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
func (m *DriveItemRequestBuilder) Items()(*i38901b482327dd22606c1b4b91388052e2bf2fc45cb9d926d8f2f8e7d35b85fc.ItemsRequestBuilder) {
    return i38901b482327dd22606c1b4b91388052e2bf2fc45cb9d926d8f2f8e7d35b85fc.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.items.item collection
func (m *DriveItemRequestBuilder) ItemsById(id string)(*idfb4790bf624d08616bb3513bffebd055a5dcadbacca095857d461aa5f1a530a.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return idfb4790bf624d08616bb3513bffebd055a5dcadbacca095857d461aa5f1a530a.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List the list property
func (m *DriveItemRequestBuilder) List()(*ia83dd16cbcab16b5c7adaabbe5934caf03fa1b23a0441de8c6b962b9aa054f24.ListRequestBuilder) {
    return ia83dd16cbcab16b5c7adaabbe5934caf03fa1b23a0441de8c6b962b9aa054f24.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Root the root property
func (m *DriveItemRequestBuilder) Root()(*i1f3a6e9c86e5b034c8fa45efe2ab068d31ba24e081da4647a4cf7f622bb330d1.RootRequestBuilder) {
    return i1f3a6e9c86e5b034c8fa45efe2ab068d31ba24e081da4647a4cf7f622bb330d1.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special the special property
func (m *DriveItemRequestBuilder) Special()(*ib11ab85aee9d209c297ca809b535a0d82cb670f95bdbdac838c4f09eea952961.SpecialRequestBuilder) {
    return ib11ab85aee9d209c297ca809b535a0d82cb670f95bdbdac838c4f09eea952961.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.special.item collection
func (m *DriveItemRequestBuilder) SpecialById(id string)(*i3e2ebdb6db6eb46e901223e1891b32e0bf81d6c6dd0899111b17be6df061afca.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i3e2ebdb6db6eb46e901223e1891b32e0bf81d6c6dd0899111b17be6df061afca.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
