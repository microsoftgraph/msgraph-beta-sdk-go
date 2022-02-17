package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/sharedwithme"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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

// DriveRequestBuilder builds and executes requests for operations under \drives\{drive-id}
type DriveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveRequestBuilderDeleteOptions options for Delete
type DriveRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveRequestBuilderGetOptions options for Get
type DriveRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveRequestBuilderGetQueryParameters get entity from drives by key
type DriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveRequestBuilderPatchOptions options for Patch
type DriveRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveRequestBuilder) Activities()(*i57502ad1c70adde13f4f4191c9621344f90ed303c506bcfac62cc6e0f211be2f.ActivitiesRequestBuilder) {
    return i57502ad1c70adde13f4f4191c9621344f90ed303c506bcfac62cc6e0f211be2f.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.activities.item collection
func (m *DriveRequestBuilder) ActivitiesById(id string)(*ib301dcbbaf0fbf1ec225026dfd49aed454924f5c763fba68c72d04e8b817701e.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ib301dcbbaf0fbf1ec225026dfd49aed454924f5c763fba68c72d04e8b817701e.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) Bundles()(*ie5ad7fd3da5b255316aa44176c2ac2f76115483b551789e79d599fe1a09af718.BundlesRequestBuilder) {
    return ie5ad7fd3da5b255316aa44176c2ac2f76115483b551789e79d599fe1a09af718.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.bundles.item collection
func (m *DriveRequestBuilder) BundlesById(id string)(*ifcd10669d080989622673b9afdc2d4fc13b76538e8a7179f0887b9fb078584da.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return ifcd10669d080989622673b9afdc2d4fc13b76538e8a7179f0887b9fb078584da.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveRequestBuilderInternal instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    m := &DriveRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveRequestBuilder instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from drives
func (m *DriveRequestBuilder) CreateDeleteRequestInformation(options *DriveRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from drives by key
func (m *DriveRequestBuilder) CreateGetRequestInformation(options *DriveRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in drives
func (m *DriveRequestBuilder) CreatePatchRequestInformation(options *DriveRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from drives
func (m *DriveRequestBuilder) Delete(options *DriveRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveRequestBuilder) Following()(*i766cace85617fc30182532069c7098a07d81041ac72cd410a073786ca7d45111.FollowingRequestBuilder) {
    return i766cace85617fc30182532069c7098a07d81041ac72cd410a073786ca7d45111.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.following.item collection
func (m *DriveRequestBuilder) FollowingById(id string)(*i2ad598655381501616b067740dbd6e99ee1ac52e664f55c88fbd4f908cb5ad17.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i2ad598655381501616b067740dbd6e99ee1ac52e664f55c88fbd4f908cb5ad17.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from drives by key
func (m *DriveRequestBuilder) Get(options *DriveRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDrive() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive), nil
}
func (m *DriveRequestBuilder) Items()(*ia54161ca01b96b4ea12e25e22325b7cb0b3cdd9294192cfbb560f72ea0db851c.ItemsRequestBuilder) {
    return ia54161ca01b96b4ea12e25e22325b7cb0b3cdd9294192cfbb560f72ea0db851c.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.items.item collection
func (m *DriveRequestBuilder) ItemsById(id string)(*i8f091105e59b53a3e7170d648bf47128b9c71e68ca777e009b742a1fec950621.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i8f091105e59b53a3e7170d648bf47128b9c71e68ca777e009b742a1fec950621.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) List()(*id58c10ecda198f10fa52e1b05c30148577f6b2f028c79657172b0133231e02af.ListRequestBuilder) {
    return id58c10ecda198f10fa52e1b05c30148577f6b2f028c79657172b0133231e02af.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in drives
func (m *DriveRequestBuilder) Patch(options *DriveRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Recent builds and executes requests for operations under \drives\{drive-id}\microsoft.graph.recent()
func (m *DriveRequestBuilder) Recent()(*i7da819773ad4a87ff5f9f3d30f0fb8b5d88885b8321f8d2383deef979e936cfe.RecentRequestBuilder) {
    return i7da819773ad4a87ff5f9f3d30f0fb8b5d88885b8321f8d2383deef979e936cfe.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Root()(*i6215746464d214fce2dc1b69b48ffe1487b12ff3c49cd51dca35bca5dbb61512.RootRequestBuilder) {
    return i6215746464d214fce2dc1b69b48ffe1487b12ff3c49cd51dca35bca5dbb61512.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ builds and executes requests for operations under \drives\{drive-id}\microsoft.graph.search(q='{q}')
func (m *DriveRequestBuilder) SearchWithQ(q *string)(*ib66a732acc47ab5473e84f72cd54f49f7de1f8a9430fe5f4ba9571cb77231778.SearchWithQRequestBuilder) {
    return ib66a732acc47ab5473e84f72cd54f49f7de1f8a9430fe5f4ba9571cb77231778.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe builds and executes requests for operations under \drives\{drive-id}\microsoft.graph.sharedWithMe()
func (m *DriveRequestBuilder) SharedWithMe()(*i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680.SharedWithMeRequestBuilder) {
    return i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Special()(*i80ad3ea091c56b67759525855e8ed8174db7bd7b55988fb87998e2088ca25d80.SpecialRequestBuilder) {
    return i80ad3ea091c56b67759525855e8ed8174db7bd7b55988fb87998e2088ca25d80.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.special.item collection
func (m *DriveRequestBuilder) SpecialById(id string)(*i30b8e702569d51cb08732e656d751cb44abd73ef6cd97f739a42e0e3b1c6e4c4.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i30b8e702569d51cb08732e656d751cb44abd73ef6cd97f739a42e0e3b1c6e4c4.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
