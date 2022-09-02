package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i13d98e9df6d9f0a887e46f4163ec3412f091bc518bd28c867198ed7720a66e00 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/activities"
    i1f36e30fbd27737753d6d2b962a59d94570fedd554798c06c1cfbcff3d8b70df "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special"
    i590a88bcee7fefe7ec9501aa51256a2fd93e9d3df9b1bb8f967bccf4471acf95 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/sharedwithme"
    i60784fc0b1fdae64d7b335f145a399178a55f485b12d11792e1f95c68a381747 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles"
    i73a42c80ea4ccfba4f756cb06fa52be25571d597fac6ef1971a95d404b1db530 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items"
    i7c600d6e0928ef8f611da39a82e01f9cebbe6ace76373046176b372e2aa84cba "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following"
    ia24a6d4516ff09182b1823691ec0812dc4a67166b9f0acfcb0d3c76bef46cbe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/searchwithq"
    ic6ceae14ff14359bd7faaa019270e3cd42dae188d042cb31ed65f533041c1cf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root"
    ida16b0d9e01c0d97027ab594642165ef048ea8fc002522904106fefcc7f43e07 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/recent"
    ie9c28e229768bea5d1f36c1f8ca219c2f6da3eb66bdd31d0ed0909346f83fda4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list"
    i642024deaaabff3935441b199198c68c9cb02d76de2c9387b74ce0d89bdb33fc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/activities/item"
    ia6763fb2ced109fa18e0ccfab7a10be35340fc76308a6e07fc024ea767c0b292 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item"
    ib54d6cd22eb60296ff6077535e6b5a7ff7b53c54c32f23ef99de0ee3a87b2246 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item"
    id6c24d057150b2e3e1f11e25b66025c6dab03eab253ad1dd508f9ef8782be5ab "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item"
    idafad3d0687ed4004b17803cc1a7815961b810b22615b441bd4bb7e7c077c487 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item"
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
func (m *DriveItemRequestBuilder) Activities()(*i13d98e9df6d9f0a887e46f4163ec3412f091bc518bd28c867198ed7720a66e00.ActivitiesRequestBuilder) {
    return i13d98e9df6d9f0a887e46f4163ec3412f091bc518bd28c867198ed7720a66e00.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.activities.item collection
func (m *DriveItemRequestBuilder) ActivitiesById(id string)(*i642024deaaabff3935441b199198c68c9cb02d76de2c9387b74ce0d89bdb33fc.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i642024deaaabff3935441b199198c68c9cb02d76de2c9387b74ce0d89bdb33fc.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bundles the bundles property
func (m *DriveItemRequestBuilder) Bundles()(*i60784fc0b1fdae64d7b335f145a399178a55f485b12d11792e1f95c68a381747.BundlesRequestBuilder) {
    return i60784fc0b1fdae64d7b335f145a399178a55f485b12d11792e1f95c68a381747.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.bundles.item collection
func (m *DriveItemRequestBuilder) BundlesById(id string)(*ia6763fb2ced109fa18e0ccfab7a10be35340fc76308a6e07fc024ea767c0b292.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ia6763fb2ced109fa18e0ccfab7a10be35340fc76308a6e07fc024ea767c0b292.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property drives for groups
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property drives for groups
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
// CreatePatchRequestInformation update the navigation property drives in groups
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property drives in groups
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
// Delete delete navigation property drives for groups
func (m *DriveItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DriveItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Following the following property
func (m *DriveItemRequestBuilder) Following()(*i7c600d6e0928ef8f611da39a82e01f9cebbe6ace76373046176b372e2aa84cba.FollowingRequestBuilder) {
    return i7c600d6e0928ef8f611da39a82e01f9cebbe6ace76373046176b372e2aa84cba.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.following.item collection
func (m *DriveItemRequestBuilder) FollowingById(id string)(*idafad3d0687ed4004b17803cc1a7815961b810b22615b441bd4bb7e7c077c487.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return idafad3d0687ed4004b17803cc1a7815961b810b22615b441bd4bb7e7c077c487.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the group's drives. Read-only.
func (m *DriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DriveItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Items the items property
func (m *DriveItemRequestBuilder) Items()(*i73a42c80ea4ccfba4f756cb06fa52be25571d597fac6ef1971a95d404b1db530.ItemsRequestBuilder) {
    return i73a42c80ea4ccfba4f756cb06fa52be25571d597fac6ef1971a95d404b1db530.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item collection
func (m *DriveItemRequestBuilder) ItemsById(id string)(*id6c24d057150b2e3e1f11e25b66025c6dab03eab253ad1dd508f9ef8782be5ab.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return id6c24d057150b2e3e1f11e25b66025c6dab03eab253ad1dd508f9ef8782be5ab.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List the list property
func (m *DriveItemRequestBuilder) List()(*ie9c28e229768bea5d1f36c1f8ca219c2f6da3eb66bdd31d0ed0909346f83fda4.ListRequestBuilder) {
    return ie9c28e229768bea5d1f36c1f8ca219c2f6da3eb66bdd31d0ed0909346f83fda4.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in groups
func (m *DriveItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Recent provides operations to call the recent method.
func (m *DriveItemRequestBuilder) Recent()(*ida16b0d9e01c0d97027ab594642165ef048ea8fc002522904106fefcc7f43e07.RecentRequestBuilder) {
    return ida16b0d9e01c0d97027ab594642165ef048ea8fc002522904106fefcc7f43e07.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root the root property
func (m *DriveItemRequestBuilder) Root()(*ic6ceae14ff14359bd7faaa019270e3cd42dae188d042cb31ed65f533041c1cf6.RootRequestBuilder) {
    return ic6ceae14ff14359bd7faaa019270e3cd42dae188d042cb31ed65f533041c1cf6.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemRequestBuilder) SearchWithQ(q *string)(*ia24a6d4516ff09182b1823691ec0812dc4a67166b9f0acfcb0d3c76bef46cbe7.SearchWithQRequestBuilder) {
    return ia24a6d4516ff09182b1823691ec0812dc4a67166b9f0acfcb0d3c76bef46cbe7.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DriveItemRequestBuilder) SharedWithMe()(*i590a88bcee7fefe7ec9501aa51256a2fd93e9d3df9b1bb8f967bccf4471acf95.SharedWithMeRequestBuilder) {
    return i590a88bcee7fefe7ec9501aa51256a2fd93e9d3df9b1bb8f967bccf4471acf95.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special the special property
func (m *DriveItemRequestBuilder) Special()(*i1f36e30fbd27737753d6d2b962a59d94570fedd554798c06c1cfbcff3d8b70df.SpecialRequestBuilder) {
    return i1f36e30fbd27737753d6d2b962a59d94570fedd554798c06c1cfbcff3d8b70df.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.special.item collection
func (m *DriveItemRequestBuilder) SpecialById(id string)(*ib54d6cd22eb60296ff6077535e6b5a7ff7b53c54c32f23ef99de0ee3a87b2246.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ib54d6cd22eb60296ff6077535e6b5a7ff7b53c54c32f23ef99de0ee3a87b2246.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
