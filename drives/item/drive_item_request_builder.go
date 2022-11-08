package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/sharedwithme"
    i57502ad1c70adde13f4f4191c9621344f90ed303c506bcfac62cc6e0f211be2f "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities"
    i6215746464d214fce2dc1b69b48ffe1487b12ff3c49cd51dca35bca5dbb61512 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root"
    i766cace85617fc30182532069c7098a07d81041ac72cd410a073786ca7d45111 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following"
    i7da819773ad4a87ff5f9f3d30f0fb8b5d88885b8321f8d2383deef979e936cfe "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/recent"
    i80ad3ea091c56b67759525855e8ed8174db7bd7b55988fb87998e2088ca25d80 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special"
    ia54161ca01b96b4ea12e25e22325b7cb0b3cdd9294192cfbb560f72ea0db851c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items"
    ib66a732acc47ab5473e84f72cd54f49f7de1f8a9430fe5f4ba9571cb77231778 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/searchwithq"
    id58c10ecda198f10fa52e1b05c30148577f6b2f028c79657172b0133231e02af "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list"
    ie5ad7fd3da5b255316aa44176c2ac2f76115483b551789e79d599fe1a09af718 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles"
    i2ad598655381501616b067740dbd6e99ee1ac52e664f55c88fbd4f908cb5ad17 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item"
    i30b8e702569d51cb08732e656d751cb44abd73ef6cd97f739a42e0e3b1c6e4c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item"
    i8f091105e59b53a3e7170d648bf47128b9c71e68ca777e009b742a1fec950621 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item"
    ib301dcbbaf0fbf1ec225026dfd49aed454924f5c763fba68c72d04e8b817701e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item"
    ifcd10669d080989622673b9afdc2d4fc13b76538e8a7179f0887b9fb078584da "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item"
)

// DriveItemRequestBuilder provides operations to manage the collection of drive entities.
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
// DriveItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a Drive resource. A Drive is the top-level container for a file system, such as OneDrive or SharePoint document libraries.
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
// Activities provides operations to manage the activities property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) Activities()(*i57502ad1c70adde13f4f4191c9621344f90ed303c506bcfac62cc6e0f211be2f.ActivitiesRequestBuilder) {
    return i57502ad1c70adde13f4f4191c9621344f90ed303c506bcfac62cc6e0f211be2f.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) ActivitiesById(id string)(*ib301dcbbaf0fbf1ec225026dfd49aed454924f5c763fba68c72d04e8b817701e.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return ib301dcbbaf0fbf1ec225026dfd49aed454924f5c763fba68c72d04e8b817701e.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bundles provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) Bundles()(*ie5ad7fd3da5b255316aa44176c2ac2f76115483b551789e79d599fe1a09af718.BundlesRequestBuilder) {
    return ie5ad7fd3da5b255316aa44176c2ac2f76115483b551789e79d599fe1a09af718.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) BundlesById(id string)(*ifcd10669d080989622673b9afdc2d4fc13b76538e8a7179f0887b9fb078584da.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ifcd10669d080989622673b9afdc2d4fc13b76538e8a7179f0887b9fb078584da.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete entity from drives
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DriveItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the properties and relationships of a Drive resource. A Drive is the top-level container for a file system, such as OneDrive or SharePoint document libraries.
func (m *DriveItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in drives
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete entity from drives
func (m *DriveItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DriveItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Following provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) Following()(*i766cace85617fc30182532069c7098a07d81041ac72cd410a073786ca7d45111.FollowingRequestBuilder) {
    return i766cace85617fc30182532069c7098a07d81041ac72cd410a073786ca7d45111.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) FollowingById(id string)(*i2ad598655381501616b067740dbd6e99ee1ac52e664f55c88fbd4f908cb5ad17.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i2ad598655381501616b067740dbd6e99ee1ac52e664f55c88fbd4f908cb5ad17.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a Drive resource. A Drive is the top-level container for a file system, such as OneDrive or SharePoint document libraries.
func (m *DriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DriveItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) Items()(*ia54161ca01b96b4ea12e25e22325b7cb0b3cdd9294192cfbb560f72ea0db851c.ItemsRequestBuilder) {
    return ia54161ca01b96b4ea12e25e22325b7cb0b3cdd9294192cfbb560f72ea0db851c.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) ItemsById(id string)(*i8f091105e59b53a3e7170d648bf47128b9c71e68ca777e009b742a1fec950621.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i8f091105e59b53a3e7170d648bf47128b9c71e68ca777e009b742a1fec950621.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List provides operations to manage the list property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) List()(*id58c10ecda198f10fa52e1b05c30148577f6b2f028c79657172b0133231e02af.ListRequestBuilder) {
    return id58c10ecda198f10fa52e1b05c30148577f6b2f028c79657172b0133231e02af.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in drives
func (m *DriveItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable), nil
}
// Recent provides operations to call the recent method.
func (m *DriveItemRequestBuilder) Recent()(*i7da819773ad4a87ff5f9f3d30f0fb8b5d88885b8321f8d2383deef979e936cfe.RecentRequestBuilder) {
    return i7da819773ad4a87ff5f9f3d30f0fb8b5d88885b8321f8d2383deef979e936cfe.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root provides operations to manage the root property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) Root()(*i6215746464d214fce2dc1b69b48ffe1487b12ff3c49cd51dca35bca5dbb61512.RootRequestBuilder) {
    return i6215746464d214fce2dc1b69b48ffe1487b12ff3c49cd51dca35bca5dbb61512.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemRequestBuilder) SearchWithQ(q *string)(*ib66a732acc47ab5473e84f72cd54f49f7de1f8a9430fe5f4ba9571cb77231778.SearchWithQRequestBuilder) {
    return ib66a732acc47ab5473e84f72cd54f49f7de1f8a9430fe5f4ba9571cb77231778.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DriveItemRequestBuilder) SharedWithMe()(*i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680.SharedWithMeRequestBuilder) {
    return i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) Special()(*i80ad3ea091c56b67759525855e8ed8174db7bd7b55988fb87998e2088ca25d80.SpecialRequestBuilder) {
    return i80ad3ea091c56b67759525855e8ed8174db7bd7b55988fb87998e2088ca25d80.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *DriveItemRequestBuilder) SpecialById(id string)(*i30b8e702569d51cb08732e656d751cb44abd73ef6cd97f739a42e0e3b1c6e4c4.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i30b8e702569d51cb08732e656d751cb44abd73ef6cd97f739a42e0e3b1c6e4c4.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
