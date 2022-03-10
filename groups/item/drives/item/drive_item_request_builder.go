package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i13d98e9df6d9f0a887e46f4163ec3412f091bc518bd28c867198ed7720a66e00 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/activities"
    i1f36e30fbd27737753d6d2b962a59d94570fedd554798c06c1cfbcff3d8b70df "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special"
    i60784fc0b1fdae64d7b335f145a399178a55f485b12d11792e1f95c68a381747 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles"
    i73a42c80ea4ccfba4f756cb06fa52be25571d597fac6ef1971a95d404b1db530 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items"
    i7c600d6e0928ef8f611da39a82e01f9cebbe6ace76373046176b372e2aa84cba "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following"
    ic6ceae14ff14359bd7faaa019270e3cd42dae188d042cb31ed65f533041c1cf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/root"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemRequestBuilderDeleteOptions options for Delete
type DriveItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetOptions options for Get
type DriveItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetQueryParameters the group's drives. Read-only.
type DriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemRequestBuilderPatchOptions options for Patch
type DriveItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Driveable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i642024deaaabff3935441b199198c68c9cb02d76de2c9387b74ce0d89bdb33fc.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["driveItem_id"] = id
    }
    return ia6763fb2ced109fa18e0ccfab7a10be35340fc76308a6e07fc024ea767c0b292.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property drives for groups
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the group's drives. Read-only.
func (m *DriveItemRequestBuilder) CreateGetRequestInformation(options *DriveItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property drives in groups
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property drives for groups
func (m *DriveItemRequestBuilder) Delete(options *DriveItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["driveItem_id"] = id
    }
    return idafad3d0687ed4004b17803cc1a7815961b810b22615b441bd4bb7e7c077c487.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the group's drives. Read-only.
func (m *DriveItemRequestBuilder) Get(options *DriveItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDriveFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Driveable), nil
}
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
        urlTplParams["driveItem_id"] = id
    }
    return id6c24d057150b2e3e1f11e25b66025c6dab03eab253ad1dd508f9ef8782be5ab.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) List()(*ie9c28e229768bea5d1f36c1f8ca219c2f6da3eb66bdd31d0ed0909346f83fda4.ListRequestBuilder) {
    return ie9c28e229768bea5d1f36c1f8ca219c2f6da3eb66bdd31d0ed0909346f83fda4.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in groups
func (m *DriveItemRequestBuilder) Patch(options *DriveItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveItemRequestBuilder) Root()(*ic6ceae14ff14359bd7faaa019270e3cd42dae188d042cb31ed65f533041c1cf6.RootRequestBuilder) {
    return ic6ceae14ff14359bd7faaa019270e3cd42dae188d042cb31ed65f533041c1cf6.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["driveItem_id"] = id
    }
    return ib54d6cd22eb60296ff6077535e6b5a7ff7b53c54c32f23ef99de0ee3a87b2246.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
