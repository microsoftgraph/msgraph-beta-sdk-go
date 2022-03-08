package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2af3c28b70aacceae8f982c5b3e17864773fa98654db360bd07d59072fd66c4f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/listitem"
    i30f0846163cdd87bf0315ac5344e6d3bf71c562afa5ae5468efe7e6e8e98f7ff "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/thumbnails"
    i43caa4d1bf49a16d841b5b24375b6ffa18f4d2ebb88ea04089a76aaf22cc32fb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/permissions"
    i4436d92b2e5984b0d8f21ece2f67020a55022821b6bd38ea3e5b42cfc831a63b "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/children"
    i843744030c5821bffacf3b7af4c6a52794f268da2a8714a2a6dc924857e81802 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/activities"
    ia396c3fd9bd63787f4e1030b240c83b0a0701642dbb433b670057d046376ec53 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/versions"
    ia7b6c90e7dedda18db74b28c4028e81a345577d20eadd779b0a1d4e5166a8d5f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/content"
    ic51979757b6a284d9b896edf4893598e4fd0b4ba0231b187ea97a4dd59a9fa83 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/analytics"
    ie5d8e79a829644772083dcff9875ae2d8a86e7f44fdad904f14b14fe0c5b6ffb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/subscriptions"
    i0a84ea84f7d89bafc9741ba8a6086ecaacc72d960e1aaa89d44b8e69649871ac "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/versions/item"
    i5aef8e96eace87b1dc6b97d468f32a42c718b61f1bba9809f6d0d676b6c1a4c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/thumbnails/item"
    i83d741b36025f559c6170b10b0c152c41aa0a879fc087598f556ec1847c2f594 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/permissions/item"
    i89ddebbe4bf7ab8f7e6f6115e78a9b5844ef97545cff53b91712c5067150c8c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/activities/item"
    ibd47e6b1ec26b3e70d04416f404a2c68ea6d7706ac82dd97443b0d6a1cf07b2c "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/children/item"
    ied463b2e0b219d7f6344437b588af0edd3b89b7a851c8259525f81e185c847eb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item/subscriptions/item"
)

// DriveItemItemRequestBuilder provides operations to manage the bundles property of the microsoft.graph.drive entity.
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
// DriveItemItemRequestBuilderGetQueryParameters collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
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
func (m *DriveItemItemRequestBuilder) Activities()(*i843744030c5821bffacf3b7af4c6a52794f268da2a8714a2a6dc924857e81802.ActivitiesRequestBuilder) {
    return i843744030c5821bffacf3b7af4c6a52794f268da2a8714a2a6dc924857e81802.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i89ddebbe4bf7ab8f7e6f6115e78a9b5844ef97545cff53b91712c5067150c8c7.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i89ddebbe4bf7ab8f7e6f6115e78a9b5844ef97545cff53b91712c5067150c8c7.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*ic51979757b6a284d9b896edf4893598e4fd0b4ba0231b187ea97a4dd59a9fa83.AnalyticsRequestBuilder) {
    return ic51979757b6a284d9b896edf4893598e4fd0b4ba0231b187ea97a4dd59a9fa83.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i4436d92b2e5984b0d8f21ece2f67020a55022821b6bd38ea3e5b42cfc831a63b.ChildrenRequestBuilder) {
    return i4436d92b2e5984b0d8f21ece2f67020a55022821b6bd38ea3e5b42cfc831a63b.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*ibd47e6b1ec26b3e70d04416f404a2c68ea6d7706ac82dd97443b0d6a1cf07b2c.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return ibd47e6b1ec26b3e70d04416f404a2c68ea6d7706ac82dd97443b0d6a1cf07b2c.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/bundles/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*ia7b6c90e7dedda18db74b28c4028e81a345577d20eadd779b0a1d4e5166a8d5f.ContentRequestBuilder) {
    return ia7b6c90e7dedda18db74b28c4028e81a345577d20eadd779b0a1d4e5166a8d5f.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property bundles for drive
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
// CreateGetRequestInformation collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
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
// CreatePatchRequestInformation update the navigation property bundles in drive
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
// Delete delete navigation property bundles for drive
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable), nil
}
func (m *DriveItemItemRequestBuilder) ListItem()(*i2af3c28b70aacceae8f982c5b3e17864773fa98654db360bd07d59072fd66c4f.ListItemRequestBuilder) {
    return i2af3c28b70aacceae8f982c5b3e17864773fa98654db360bd07d59072fd66c4f.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bundles in drive
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveItemItemRequestBuilder) Permissions()(*i43caa4d1bf49a16d841b5b24375b6ffa18f4d2ebb88ea04089a76aaf22cc32fb.PermissionsRequestBuilder) {
    return i43caa4d1bf49a16d841b5b24375b6ffa18f4d2ebb88ea04089a76aaf22cc32fb.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i83d741b36025f559c6170b10b0c152c41aa0a879fc087598f556ec1847c2f594.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i83d741b36025f559c6170b10b0c152c41aa0a879fc087598f556ec1847c2f594.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*ie5d8e79a829644772083dcff9875ae2d8a86e7f44fdad904f14b14fe0c5b6ffb.SubscriptionsRequestBuilder) {
    return ie5d8e79a829644772083dcff9875ae2d8a86e7f44fdad904f14b14fe0c5b6ffb.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*ied463b2e0b219d7f6344437b588af0edd3b89b7a851c8259525f81e185c847eb.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ied463b2e0b219d7f6344437b588af0edd3b89b7a851c8259525f81e185c847eb.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i30f0846163cdd87bf0315ac5344e6d3bf71c562afa5ae5468efe7e6e8e98f7ff.ThumbnailsRequestBuilder) {
    return i30f0846163cdd87bf0315ac5344e6d3bf71c562afa5ae5468efe7e6e8e98f7ff.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i5aef8e96eace87b1dc6b97d468f32a42c718b61f1bba9809f6d0d676b6c1a4c0.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i5aef8e96eace87b1dc6b97d468f32a42c718b61f1bba9809f6d0d676b6c1a4c0.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*ia396c3fd9bd63787f4e1030b240c83b0a0701642dbb433b670057d046376ec53.VersionsRequestBuilder) {
    return ia396c3fd9bd63787f4e1030b240c83b0a0701642dbb433b670057d046376ec53.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i0a84ea84f7d89bafc9741ba8a6086ecaacc72d960e1aaa89d44b8e69649871ac.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i0a84ea84f7d89bafc9741ba8a6086ecaacc72d960e1aaa89d44b8e69649871ac.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
