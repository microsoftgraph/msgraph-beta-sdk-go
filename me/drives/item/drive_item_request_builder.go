package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i02954bc3c177a5c172817e8d08f8f0ababe0cadc5d87b42b0db53ba7fc5f6a48 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/activities"
    i10734e5574d026ee7008fa1a817b773a50eaab1d8d0e4c14bbec11b98889c7b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root"
    i38620e6bf818a619e37deec08fae87e873d6a9f1841ad70b3e20f4051b72418c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list"
    i3a7460915971e60b05b27c726b8a042546eedba133703a52e040cef3172d8aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/items"
    i550cba7ea0323d87e9a845d8a8e607c5fcacfc6414513b4b0e5465503097f5a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following"
    i7d7eb2492b989eaeef6f0dcf519dc8bbf21be1aceb0d0b65df678728db6f9300 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special"
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
// DriveItemRequestBuilderGetQueryParameters a collection of drives available for this user. Read-only.
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
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i27c3061950c48cb0d030b04c4820e7394d3adf08b0833fad916517cfb33b5c89.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["driveItem_id"] = id
    }
    return i7f3652e0df9aa18944938e03a9df7978eda2c160eb890d34ad100d7d368150ff.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property drives for me
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
// CreateGetRequestInformation a collection of drives available for this user. Read-only.
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
// CreatePatchRequestInformation update the navigation property drives in me
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
// Delete delete navigation property drives for me
func (m *DriveItemRequestBuilder) Delete(options *DriveItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["driveItem_id"] = id
    }
    return i611d484de79682584613cb6c70270d1ee38deed49323be6a56db4c4d59546e81.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get a collection of drives available for this user. Read-only.
func (m *DriveItemRequestBuilder) Get(options *DriveItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDriveFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Driveable), nil
}
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
        urlTplParams["driveItem_id"] = id
    }
    return ie765b99b1e518bbaf92749f18947b0fd27e063cf001932f8d3c2478d4e0cbf00.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) List()(*i38620e6bf818a619e37deec08fae87e873d6a9f1841ad70b3e20f4051b72418c.ListRequestBuilder) {
    return i38620e6bf818a619e37deec08fae87e873d6a9f1841ad70b3e20f4051b72418c.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in me
func (m *DriveItemRequestBuilder) Patch(options *DriveItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveItemRequestBuilder) Root()(*i10734e5574d026ee7008fa1a817b773a50eaab1d8d0e4c14bbec11b98889c7b6.RootRequestBuilder) {
    return i10734e5574d026ee7008fa1a817b773a50eaab1d8d0e4c14bbec11b98889c7b6.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["driveItem_id"] = id
    }
    return i4717b582f54b9d16c81cfb2753bde0a6c3b5d72f2b3e6358fee228dcd0596f00.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
