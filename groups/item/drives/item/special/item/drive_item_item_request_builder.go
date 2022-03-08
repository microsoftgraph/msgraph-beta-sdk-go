package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i21325c8490d1ffc94e0a74ca140d8298ae98dd90bc5e1dd8cdf1badf795afc01 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/versions"
    i5c89caa7c2b4b37281dbd0ebb8539116c410a7586a56b648be3bc9de3089f22b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/children"
    i71245db7a9258dc83dbfc24a07f22580787d792bf9a997a3022568fd589dfc76 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/thumbnails"
    i7d59061d7023b497c78edfec6fb908d09cd4f60b5794309d7df51ba37ac15021 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/permissions"
    i80ec37695fb0c6f7287762d2bf0c7a3545998a2887535ba30c462ead7db36f09 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/analytics"
    ic27fefa8188c28f1f6c00853a2a3f045e2f3cb587be40a5fb34d7b5773e43981 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/subscriptions"
    ic7e118719124077284075c90431075df070724d0459318fd78deaa2dbf9f85d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/activities"
    ie54d09be1c3b1e08ba33cd5ccc436e9c44538573455789114ba553d33d223852 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/content"
    ifa33a4f5768800519a117d1d4e38ad3495b78aa0a34adffd7e39802c62d5d25d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/listitem"
    i0169fa2dbea4f1d319dc112a590580d58a8f89a7dfcb6d544ef707657d3ed151 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/subscriptions/item"
    i14253fddcd09d57888b82eb6ebdbd3415a71cdc4b1899d6bd8d44db1e4b5b369 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/versions/item"
    i9c0263d0fbf38ec20c1b0af3962d9b110dc2cc5b1fd31b81d2fd2356f143bb2c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/activities/item"
    id2f4a6d2b80079fd24ba030c0d78cc281a5d8c10e2735da02070d29c7946ad13 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/children/item"
    ida9aca8074ea5a4ebca2697924853ebaedb13bac0afd14d700c396cc453071de "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/thumbnails/item"
    ie5590b1f22b74bdc87a0c03ac32e766a07438a3d715b5b69e380aa7105270390 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/special/item/permissions/item"
)

// DriveItemItemRequestBuilder provides operations to manage the special property of the microsoft.graph.drive entity.
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
// DriveItemItemRequestBuilderGetQueryParameters collection of common folders available in OneDrive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) Activities()(*ic7e118719124077284075c90431075df070724d0459318fd78deaa2dbf9f85d3.ActivitiesRequestBuilder) {
    return ic7e118719124077284075c90431075df070724d0459318fd78deaa2dbf9f85d3.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.special.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i9c0263d0fbf38ec20c1b0af3962d9b110dc2cc5b1fd31b81d2fd2356f143bb2c.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i9c0263d0fbf38ec20c1b0af3962d9b110dc2cc5b1fd31b81d2fd2356f143bb2c.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i80ec37695fb0c6f7287762d2bf0c7a3545998a2887535ba30c462ead7db36f09.AnalyticsRequestBuilder) {
    return i80ec37695fb0c6f7287762d2bf0c7a3545998a2887535ba30c462ead7db36f09.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i5c89caa7c2b4b37281dbd0ebb8539116c410a7586a56b648be3bc9de3089f22b.ChildrenRequestBuilder) {
    return i5c89caa7c2b4b37281dbd0ebb8539116c410a7586a56b648be3bc9de3089f22b.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.special.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*id2f4a6d2b80079fd24ba030c0d78cc281a5d8c10e2735da02070d29c7946ad13.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return id2f4a6d2b80079fd24ba030c0d78cc281a5d8c10e2735da02070d29c7946ad13.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}/special/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*ie54d09be1c3b1e08ba33cd5ccc436e9c44538573455789114ba553d33d223852.ContentRequestBuilder) {
    return ie54d09be1c3b1e08ba33cd5ccc436e9c44538573455789114ba553d33d223852.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property special for groups
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
// CreateGetRequestInformation collection of common folders available in OneDrive. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property special in groups
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
// Delete delete navigation property special for groups
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
// Get collection of common folders available in OneDrive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*ifa33a4f5768800519a117d1d4e38ad3495b78aa0a34adffd7e39802c62d5d25d.ListItemRequestBuilder) {
    return ifa33a4f5768800519a117d1d4e38ad3495b78aa0a34adffd7e39802c62d5d25d.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property special in groups
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i7d59061d7023b497c78edfec6fb908d09cd4f60b5794309d7df51ba37ac15021.PermissionsRequestBuilder) {
    return i7d59061d7023b497c78edfec6fb908d09cd4f60b5794309d7df51ba37ac15021.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.special.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ie5590b1f22b74bdc87a0c03ac32e766a07438a3d715b5b69e380aa7105270390.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ie5590b1f22b74bdc87a0c03ac32e766a07438a3d715b5b69e380aa7105270390.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*ic27fefa8188c28f1f6c00853a2a3f045e2f3cb587be40a5fb34d7b5773e43981.SubscriptionsRequestBuilder) {
    return ic27fefa8188c28f1f6c00853a2a3f045e2f3cb587be40a5fb34d7b5773e43981.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.special.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i0169fa2dbea4f1d319dc112a590580d58a8f89a7dfcb6d544ef707657d3ed151.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i0169fa2dbea4f1d319dc112a590580d58a8f89a7dfcb6d544ef707657d3ed151.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i71245db7a9258dc83dbfc24a07f22580787d792bf9a997a3022568fd589dfc76.ThumbnailsRequestBuilder) {
    return i71245db7a9258dc83dbfc24a07f22580787d792bf9a997a3022568fd589dfc76.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.special.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ida9aca8074ea5a4ebca2697924853ebaedb13bac0afd14d700c396cc453071de.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return ida9aca8074ea5a4ebca2697924853ebaedb13bac0afd14d700c396cc453071de.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i21325c8490d1ffc94e0a74ca140d8298ae98dd90bc5e1dd8cdf1badf795afc01.VersionsRequestBuilder) {
    return i21325c8490d1ffc94e0a74ca140d8298ae98dd90bc5e1dd8cdf1badf795afc01.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.special.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i14253fddcd09d57888b82eb6ebdbd3415a71cdc4b1899d6bd8d44db1e4b5b369.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i14253fddcd09d57888b82eb6ebdbd3415a71cdc4b1899d6bd8d44db1e4b5b369.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
