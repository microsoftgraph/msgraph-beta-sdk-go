package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i02954bc3c177a5c172817e8d08f8f0ababe0cadc5d87b42b0db53ba7fc5f6a48 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/activities"
    i06266d3ee1778a67ea6aa5e6efa5c54100b8ee30dc2f1c72685c2b48800eea16 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/recent"
    i10734e5574d026ee7008fa1a817b773a50eaab1d8d0e4c14bbec11b98889c7b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root"
    i1dd633d0c191de365555cbb935c24a20715a6c574ba0efe83341cfc53c41b0aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/sharedwithme"
    i38620e6bf818a619e37deec08fae87e873d6a9f1841ad70b3e20f4051b72418c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list"
    i3a7460915971e60b05b27c726b8a042546eedba133703a52e040cef3172d8aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items"
    i550cba7ea0323d87e9a845d8a8e607c5fcacfc6414513b4b0e5465503097f5a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following"
    i7d7eb2492b989eaeef6f0dcf519dc8bbf21be1aceb0d0b65df678728db6f9300 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special"
    ica565798ba029e4cabed39d365bc977c9f867673798acc888624e05dd00bf327 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/searchwithq"
    iea353883790e50113ed2641481a361b4706b7899221656df6d941c5a5c0e50ab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles"
    i27c3061950c48cb0d030b04c4820e7394d3adf08b0833fad916517cfb33b5c89 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/activities/item"
    i4717b582f54b9d16c81cfb2753bde0a6c3b5d72f2b3e6358fee228dcd0596f00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item"
    i611d484de79682584613cb6c70270d1ee38deed49323be6a56db4c4d59546e81 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item"
    i7f3652e0df9aa18944938e03a9df7978eda2c160eb890d34ad100d7d368150ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item"
    ie765b99b1e518bbaf92749f18947b0fd27e063cf001932f8d3c2478d4e0cbf00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items/item"
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
func (m *DriveItemRequestBuilder) Activities()(*i02954bc3c177a5c172817e8d08f8f0ababe0cadc5d87b42b0db53ba7fc5f6a48.ActivitiesRequestBuilder) {
    return i02954bc3c177a5c172817e8d08f8f0ababe0cadc5d87b42b0db53ba7fc5f6a48.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.activities.item collection
func (m *DriveItemRequestBuilder) ActivitiesById(id string)(*i27c3061950c48cb0d030b04c4820e7394d3adf08b0833fad916517cfb33b5c89.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i27c3061950c48cb0d030b04c4820e7394d3adf08b0833fad916517cfb33b5c89.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bundles the bundles property
func (m *DriveItemRequestBuilder) Bundles()(*iea353883790e50113ed2641481a361b4706b7899221656df6d941c5a5c0e50ab.BundlesRequestBuilder) {
    return iea353883790e50113ed2641481a361b4706b7899221656df6d941c5a5c0e50ab.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.bundles.item collection
func (m *DriveItemRequestBuilder) BundlesById(id string)(*i7f3652e0df9aa18944938e03a9df7978eda2c160eb890d34ad100d7d368150ff.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i7f3652e0df9aa18944938e03a9df7978eda2c160eb890d34ad100d7d368150ff.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property drives for me
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
// CreateGetRequestInformation a collection of drives available for this user. Read-only.
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
// CreatePatchRequestInformation update the navigation property drives in me
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
// Delete delete navigation property drives for me
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
// Following the following property
func (m *DriveItemRequestBuilder) Following()(*i550cba7ea0323d87e9a845d8a8e607c5fcacfc6414513b4b0e5465503097f5a0.FollowingRequestBuilder) {
    return i550cba7ea0323d87e9a845d8a8e607c5fcacfc6414513b4b0e5465503097f5a0.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item collection
func (m *DriveItemRequestBuilder) FollowingById(id string)(*i611d484de79682584613cb6c70270d1ee38deed49323be6a56db4c4d59546e81.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i611d484de79682584613cb6c70270d1ee38deed49323be6a56db4c4d59546e81.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get a collection of drives available for this user. Read-only.
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
// Items the items property
func (m *DriveItemRequestBuilder) Items()(*i3a7460915971e60b05b27c726b8a042546eedba133703a52e040cef3172d8aa3.ItemsRequestBuilder) {
    return i3a7460915971e60b05b27c726b8a042546eedba133703a52e040cef3172d8aa3.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.items.item collection
func (m *DriveItemRequestBuilder) ItemsById(id string)(*ie765b99b1e518bbaf92749f18947b0fd27e063cf001932f8d3c2478d4e0cbf00.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ie765b99b1e518bbaf92749f18947b0fd27e063cf001932f8d3c2478d4e0cbf00.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List the list property
func (m *DriveItemRequestBuilder) List()(*i38620e6bf818a619e37deec08fae87e873d6a9f1841ad70b3e20f4051b72418c.ListRequestBuilder) {
    return i38620e6bf818a619e37deec08fae87e873d6a9f1841ad70b3e20f4051b72418c.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in me
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
func (m *DriveItemRequestBuilder) Recent()(*i06266d3ee1778a67ea6aa5e6efa5c54100b8ee30dc2f1c72685c2b48800eea16.RecentRequestBuilder) {
    return i06266d3ee1778a67ea6aa5e6efa5c54100b8ee30dc2f1c72685c2b48800eea16.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root the root property
func (m *DriveItemRequestBuilder) Root()(*i10734e5574d026ee7008fa1a817b773a50eaab1d8d0e4c14bbec11b98889c7b6.RootRequestBuilder) {
    return i10734e5574d026ee7008fa1a817b773a50eaab1d8d0e4c14bbec11b98889c7b6.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemRequestBuilder) SearchWithQ(q *string)(*ica565798ba029e4cabed39d365bc977c9f867673798acc888624e05dd00bf327.SearchWithQRequestBuilder) {
    return ica565798ba029e4cabed39d365bc977c9f867673798acc888624e05dd00bf327.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DriveItemRequestBuilder) SharedWithMe()(*i1dd633d0c191de365555cbb935c24a20715a6c574ba0efe83341cfc53c41b0aa.SharedWithMeRequestBuilder) {
    return i1dd633d0c191de365555cbb935c24a20715a6c574ba0efe83341cfc53c41b0aa.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special the special property
func (m *DriveItemRequestBuilder) Special()(*i7d7eb2492b989eaeef6f0dcf519dc8bbf21be1aceb0d0b65df678728db6f9300.SpecialRequestBuilder) {
    return i7d7eb2492b989eaeef6f0dcf519dc8bbf21be1aceb0d0b65df678728db6f9300.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.special.item collection
func (m *DriveItemRequestBuilder) SpecialById(id string)(*i4717b582f54b9d16c81cfb2753bde0a6c3b5d72f2b3e6358fee228dcd0596f00.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i4717b582f54b9d16c81cfb2753bde0a6c3b5d72f2b3e6358fee228dcd0596f00.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
