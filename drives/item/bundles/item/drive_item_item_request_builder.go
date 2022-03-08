package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1d958e64d6eebbf3f0fb4e5a69bab7233eccf733f531b6cc42773691309fe656 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/activities"
    i3df01e03f70e5e6cb622406ba0c7683bbaa6aad7715bbab0cd7ebc48760d0e45 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/analytics"
    i5c0e2caa8d4a52220c4790079120c58840d69f7f18eda16af5953f8412feebcf "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/versions"
    i6bfa2a2a1670df4feeed7bf08a9355ded750728b3804060cb904e70a6c77b6b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/permissions"
    i78f662cc0dbe916ab3d1ebbc2789af002318a53dccca4ae10dfb2594bfa5aece "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/content"
    i8e034313654df94364c270367407349cf1026d620f43f3912a1d9edaae069fab "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/thumbnails"
    i9f05d72b32696f6b92c6f9f4c2420e7f280a8548ac701f15d2a2fd5e250df362 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/listitem"
    icb1b6622cd720f191231210dde1225821f2b7983688a3951e809aa43333186f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/subscriptions"
    if5cb0e0d7efd6ee902a46f01107104b7c63fdbb841518ea1f375158fb5b11e6a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/children"
    i51a0aca5f7c0ec86f28e32c855bf8ff2c6e1103071c107774ed8238f745a6870 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/activities/item"
    i6140129a7318c938f599403d64fe425f5e7cd751a66f8c3c2aff5c6b67f49118 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/subscriptions/item"
    i7463308b64701a8c1643477195afa3d80974526ef967ec78962ad31547f11226 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/children/item"
    ib2c1541e75ec008dcd6170ea115ff105aff06cbd4e9a537a722ca1a8ddcd4d52 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/permissions/item"
    ib68df0728458c6dba50c27f5b9636c00e5b9b86b6a81c124d57152bf33f2817d "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/versions/item"
    ib6cede8e7f7e5cc0c31aabb04cb36893197952992cfb9e1c40779434d3b0d7c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item/thumbnails/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i1d958e64d6eebbf3f0fb4e5a69bab7233eccf733f531b6cc42773691309fe656.ActivitiesRequestBuilder) {
    return i1d958e64d6eebbf3f0fb4e5a69bab7233eccf733f531b6cc42773691309fe656.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.bundles.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i51a0aca5f7c0ec86f28e32c855bf8ff2c6e1103071c107774ed8238f745a6870.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i51a0aca5f7c0ec86f28e32c855bf8ff2c6e1103071c107774ed8238f745a6870.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i3df01e03f70e5e6cb622406ba0c7683bbaa6aad7715bbab0cd7ebc48760d0e45.AnalyticsRequestBuilder) {
    return i3df01e03f70e5e6cb622406ba0c7683bbaa6aad7715bbab0cd7ebc48760d0e45.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*if5cb0e0d7efd6ee902a46f01107104b7c63fdbb841518ea1f375158fb5b11e6a.ChildrenRequestBuilder) {
    return if5cb0e0d7efd6ee902a46f01107104b7c63fdbb841518ea1f375158fb5b11e6a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.bundles.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i7463308b64701a8c1643477195afa3d80974526ef967ec78962ad31547f11226.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i7463308b64701a8c1643477195afa3d80974526ef967ec78962ad31547f11226.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/bundles/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i78f662cc0dbe916ab3d1ebbc2789af002318a53dccca4ae10dfb2594bfa5aece.ContentRequestBuilder) {
    return i78f662cc0dbe916ab3d1ebbc2789af002318a53dccca4ae10dfb2594bfa5aece.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property bundles for drives
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
// CreatePatchRequestInformation update the navigation property bundles in drives
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
// Delete delete navigation property bundles for drives
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i9f05d72b32696f6b92c6f9f4c2420e7f280a8548ac701f15d2a2fd5e250df362.ListItemRequestBuilder) {
    return i9f05d72b32696f6b92c6f9f4c2420e7f280a8548ac701f15d2a2fd5e250df362.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bundles in drives
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i6bfa2a2a1670df4feeed7bf08a9355ded750728b3804060cb904e70a6c77b6b7.PermissionsRequestBuilder) {
    return i6bfa2a2a1670df4feeed7bf08a9355ded750728b3804060cb904e70a6c77b6b7.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.bundles.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ib2c1541e75ec008dcd6170ea115ff105aff06cbd4e9a537a722ca1a8ddcd4d52.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ib2c1541e75ec008dcd6170ea115ff105aff06cbd4e9a537a722ca1a8ddcd4d52.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*icb1b6622cd720f191231210dde1225821f2b7983688a3951e809aa43333186f3.SubscriptionsRequestBuilder) {
    return icb1b6622cd720f191231210dde1225821f2b7983688a3951e809aa43333186f3.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.bundles.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i6140129a7318c938f599403d64fe425f5e7cd751a66f8c3c2aff5c6b67f49118.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i6140129a7318c938f599403d64fe425f5e7cd751a66f8c3c2aff5c6b67f49118.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i8e034313654df94364c270367407349cf1026d620f43f3912a1d9edaae069fab.ThumbnailsRequestBuilder) {
    return i8e034313654df94364c270367407349cf1026d620f43f3912a1d9edaae069fab.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.bundles.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ib6cede8e7f7e5cc0c31aabb04cb36893197952992cfb9e1c40779434d3b0d7c6.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return ib6cede8e7f7e5cc0c31aabb04cb36893197952992cfb9e1c40779434d3b0d7c6.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i5c0e2caa8d4a52220c4790079120c58840d69f7f18eda16af5953f8412feebcf.VersionsRequestBuilder) {
    return i5c0e2caa8d4a52220c4790079120c58840d69f7f18eda16af5953f8412feebcf.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.bundles.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ib68df0728458c6dba50c27f5b9636c00e5b9b86b6a81c124d57152bf33f2817d.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ib68df0728458c6dba50c27f5b9636c00e5b9b86b6a81c124d57152bf33f2817d.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
