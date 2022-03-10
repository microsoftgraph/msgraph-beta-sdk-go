package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i126d6d90faf0a2e1ca6085a4ee795d080b8ddc8182416126c037a09df060556b "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/activities"
    i26c0e47deeccd0569817ee9563473af86d3fb135ace1eec2800e531691ab05e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/listitem"
    i582cb9c6bef0aae8daf0be7d7610a8ba399c916babc23092941ddd8bf10789ab "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/subscriptions"
    i75eb803effc4876884c7e0c6664eddcbe2110bbdf4af5809f0a2c16a06b0a9ed "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/content"
    i9027c2ab10c1c928e6718fd372bb0d74cc5234d8b748fb5916e97a54e0599dd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/analytics"
    iacf1b80a8db4e84c17e629183e19f550fa2610657a5eb8636590fdac3db07363 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/versions"
    ib3bd06f8e1e207e080f64d8a6cd8ce5e1327a5ef255058104a55ed0c8473f2a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/permissions"
    id13da4b2bd76059d2d931c13c4257c02f7a09aa908d907205dade1351d5725f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/children"
    iea190aebe9524ac79895ae8839755317346f7aab63e45e4ba717d600e82c10c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/thumbnails"
    i03e85d138955064441e78dc3dd7c270d1ab92cf722a47f57ae4a887084b774cc "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/thumbnails/item"
    i0c3312095e79899c3f4d2d7323c8af4df868a15ee70fafca42d294a45d0de323 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/activities/item"
    i245376ebbaa4395b8ae12d6918e678ed88ef1da4222e64ec6d602618dbf3f8f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/permissions/item"
    i2bd853e718c71da309cf2f322ec345aa9be84109d5f6bdb9668752531eb2abd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/children/item"
    i7005a7f9bb8b1e3b4e12a5fadf6a7ea7d38f29843b058e94512aeeb6e601933c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/versions/item"
    i932a034e5b44572be1c4b4344122fdf04f1e5347a58a48c4450f53b81c5f6f8f "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item/subscriptions/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i126d6d90faf0a2e1ca6085a4ee795d080b8ddc8182416126c037a09df060556b.ActivitiesRequestBuilder) {
    return i126d6d90faf0a2e1ca6085a4ee795d080b8ddc8182416126c037a09df060556b.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.following.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i0c3312095e79899c3f4d2d7323c8af4df868a15ee70fafca42d294a45d0de323.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i0c3312095e79899c3f4d2d7323c8af4df868a15ee70fafca42d294a45d0de323.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i9027c2ab10c1c928e6718fd372bb0d74cc5234d8b748fb5916e97a54e0599dd8.AnalyticsRequestBuilder) {
    return i9027c2ab10c1c928e6718fd372bb0d74cc5234d8b748fb5916e97a54e0599dd8.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*id13da4b2bd76059d2d931c13c4257c02f7a09aa908d907205dade1351d5725f7.ChildrenRequestBuilder) {
    return id13da4b2bd76059d2d931c13c4257c02f7a09aa908d907205dade1351d5725f7.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.following.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i2bd853e718c71da309cf2f322ec345aa9be84109d5f6bdb9668752531eb2abd1.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i2bd853e718c71da309cf2f322ec345aa9be84109d5f6bdb9668752531eb2abd1.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/following/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i75eb803effc4876884c7e0c6664eddcbe2110bbdf4af5809f0a2c16a06b0a9ed.ContentRequestBuilder) {
    return i75eb803effc4876884c7e0c6664eddcbe2110bbdf4af5809f0a2c16a06b0a9ed.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property following for drives
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
// CreatePatchRequestInformation update the navigation property following in drives
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
// Delete delete navigation property following for drives
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
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
// Get the list of items the user is following. Only in OneDrive for Business.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable), nil
}
func (m *DriveItemItemRequestBuilder) ListItem()(*i26c0e47deeccd0569817ee9563473af86d3fb135ace1eec2800e531691ab05e7.ListItemRequestBuilder) {
    return i26c0e47deeccd0569817ee9563473af86d3fb135ace1eec2800e531691ab05e7.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property following in drives
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*ib3bd06f8e1e207e080f64d8a6cd8ce5e1327a5ef255058104a55ed0c8473f2a6.PermissionsRequestBuilder) {
    return ib3bd06f8e1e207e080f64d8a6cd8ce5e1327a5ef255058104a55ed0c8473f2a6.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.following.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i245376ebbaa4395b8ae12d6918e678ed88ef1da4222e64ec6d602618dbf3f8f4.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i245376ebbaa4395b8ae12d6918e678ed88ef1da4222e64ec6d602618dbf3f8f4.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i582cb9c6bef0aae8daf0be7d7610a8ba399c916babc23092941ddd8bf10789ab.SubscriptionsRequestBuilder) {
    return i582cb9c6bef0aae8daf0be7d7610a8ba399c916babc23092941ddd8bf10789ab.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.following.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i932a034e5b44572be1c4b4344122fdf04f1e5347a58a48c4450f53b81c5f6f8f.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i932a034e5b44572be1c4b4344122fdf04f1e5347a58a48c4450f53b81c5f6f8f.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*iea190aebe9524ac79895ae8839755317346f7aab63e45e4ba717d600e82c10c3.ThumbnailsRequestBuilder) {
    return iea190aebe9524ac79895ae8839755317346f7aab63e45e4ba717d600e82c10c3.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.following.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i03e85d138955064441e78dc3dd7c270d1ab92cf722a47f57ae4a887084b774cc.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i03e85d138955064441e78dc3dd7c270d1ab92cf722a47f57ae4a887084b774cc.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*iacf1b80a8db4e84c17e629183e19f550fa2610657a5eb8636590fdac3db07363.VersionsRequestBuilder) {
    return iacf1b80a8db4e84c17e629183e19f550fa2610657a5eb8636590fdac3db07363.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.following.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i7005a7f9bb8b1e3b4e12a5fadf6a7ea7d38f29843b058e94512aeeb6e601933c.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i7005a7f9bb8b1e3b4e12a5fadf6a7ea7d38f29843b058e94512aeeb6e601933c.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
