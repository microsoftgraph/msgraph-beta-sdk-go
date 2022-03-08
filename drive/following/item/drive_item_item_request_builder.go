package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3614bfa7f99c5c5c8eecb937ba18fda11f9747ff21700776b7aea636be001db1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/analytics"
    i422a4373e0965913ac0d82ab650e181c787d5d5131256f6462568769f17f6c92 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/activities"
    i809dd175543aa0a9823826d014f6b8d94f6514baa2169920a99a50c7afa23950 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/content"
    ia755cc25b36de3eeeac98374e053ac7fcf5c9c1981ede9a92dd0b95b215f0502 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/thumbnails"
    ib01178bc92e683f57d8f3e77b447fbeac86d028abf51027611381654dfd03ff0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/listitem"
    ib896fbc7c56ea69172b48ecc34a2699cece3c5bad2498b9ccbd3a1aca70b2ef5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/versions"
    id76da30a23d51b4ef51a6953565c5287e71aa4f131cb90e756222ec0948c9db7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/permissions"
    ie469c6c69191a556eb9ad6fa33d1d17cbeefeaf0622961603a3f5a4857286202 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/subscriptions"
    iecd3c38bae28d3b228692a4940fde3e51a1b7829e65bbc8607587466a939a8fc "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/children"
    i18ae31e0ae755a6cf36d0cbc2445b1c5bd23cf8589cff3eec27920210aa7941b "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/activities/item"
    i85ac2d6fe304bb0bb92d6e0d03c52eadce5cc1d4c10cb90d0f308808179d8f73 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/permissions/item"
    i8df977978de8b74f8ba0e9c09db0a4a8d170ceeb11b4f3f87fd61fc2e49c71e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/children/item"
    ibe27173f65043006836ea6465f64d2b271368568bb9df7b0ed63e286f1fde1d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/versions/item"
    id7474a56b708ff13105f23cf34755e9da6acd5a7d9e2271bd0d086be157e555a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/thumbnails/item"
    ieb4683e2829c923848f4d219d9af391a1900de63a6aa18724db02681650c09c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item/subscriptions/item"
)

// DriveItemItemRequestBuilder provides operations to manage the following property of the microsoft.graph.drive entity.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemItemRequestBuilderDeleteOptions options for Delete
type DriveItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemItemRequestBuilderGetOptions options for Get
type DriveItemItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemItemRequestBuilderGetQueryParameters the list of items the user is following. Only in OneDrive for Business.
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemItemRequestBuilderPatchOptions options for Patch
type DriveItemItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveItemItemRequestBuilder) Activities()(*i422a4373e0965913ac0d82ab650e181c787d5d5131256f6462568769f17f6c92.ActivitiesRequestBuilder) {
    return i422a4373e0965913ac0d82ab650e181c787d5d5131256f6462568769f17f6c92.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.following.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i18ae31e0ae755a6cf36d0cbc2445b1c5bd23cf8589cff3eec27920210aa7941b.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i18ae31e0ae755a6cf36d0cbc2445b1c5bd23cf8589cff3eec27920210aa7941b.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i3614bfa7f99c5c5c8eecb937ba18fda11f9747ff21700776b7aea636be001db1.AnalyticsRequestBuilder) {
    return i3614bfa7f99c5c5c8eecb937ba18fda11f9747ff21700776b7aea636be001db1.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*iecd3c38bae28d3b228692a4940fde3e51a1b7829e65bbc8607587466a939a8fc.ChildrenRequestBuilder) {
    return iecd3c38bae28d3b228692a4940fde3e51a1b7829e65bbc8607587466a939a8fc.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.following.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i8df977978de8b74f8ba0e9c09db0a4a8d170ceeb11b4f3f87fd61fc2e49c71e4.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i8df977978de8b74f8ba0e9c09db0a4a8d170ceeb11b4f3f87fd61fc2e49c71e4.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/following/{driveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DriveItemItemRequestBuilder) Content()(*i809dd175543aa0a9823826d014f6b8d94f6514baa2169920a99a50c7afa23950.ContentRequestBuilder) {
    return i809dd175543aa0a9823826d014f6b8d94f6514baa2169920a99a50c7afa23950.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property following for drive
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of items the user is following. Only in OneDrive for Business.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation(options *DriveItemItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property following in drive
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property following for drive
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
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
// Get the list of items the user is following. Only in OneDrive for Business.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable), nil
}
func (m *DriveItemItemRequestBuilder) ListItem()(*ib01178bc92e683f57d8f3e77b447fbeac86d028abf51027611381654dfd03ff0.ListItemRequestBuilder) {
    return ib01178bc92e683f57d8f3e77b447fbeac86d028abf51027611381654dfd03ff0.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property following in drive
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*id76da30a23d51b4ef51a6953565c5287e71aa4f131cb90e756222ec0948c9db7.PermissionsRequestBuilder) {
    return id76da30a23d51b4ef51a6953565c5287e71aa4f131cb90e756222ec0948c9db7.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.following.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i85ac2d6fe304bb0bb92d6e0d03c52eadce5cc1d4c10cb90d0f308808179d8f73.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i85ac2d6fe304bb0bb92d6e0d03c52eadce5cc1d4c10cb90d0f308808179d8f73.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*ie469c6c69191a556eb9ad6fa33d1d17cbeefeaf0622961603a3f5a4857286202.SubscriptionsRequestBuilder) {
    return ie469c6c69191a556eb9ad6fa33d1d17cbeefeaf0622961603a3f5a4857286202.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.following.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*ieb4683e2829c923848f4d219d9af391a1900de63a6aa18724db02681650c09c5.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ieb4683e2829c923848f4d219d9af391a1900de63a6aa18724db02681650c09c5.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*ia755cc25b36de3eeeac98374e053ac7fcf5c9c1981ede9a92dd0b95b215f0502.ThumbnailsRequestBuilder) {
    return ia755cc25b36de3eeeac98374e053ac7fcf5c9c1981ede9a92dd0b95b215f0502.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.following.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*id7474a56b708ff13105f23cf34755e9da6acd5a7d9e2271bd0d086be157e555a.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return id7474a56b708ff13105f23cf34755e9da6acd5a7d9e2271bd0d086be157e555a.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*ib896fbc7c56ea69172b48ecc34a2699cece3c5bad2498b9ccbd3a1aca70b2ef5.VersionsRequestBuilder) {
    return ib896fbc7c56ea69172b48ecc34a2699cece3c5bad2498b9ccbd3a1aca70b2ef5.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.following.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ibe27173f65043006836ea6465f64d2b271368568bb9df7b0ed63e286f1fde1d6.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ibe27173f65043006836ea6465f64d2b271368568bb9df7b0ed63e286f1fde1d6.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
