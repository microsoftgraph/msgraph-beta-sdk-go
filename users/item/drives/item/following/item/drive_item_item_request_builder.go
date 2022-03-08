package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i250757ab11c7812509a51c9b493b53cd482a62df2de58cf587bf23efffd92e3a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/children"
    i2b124f40c66d4ad62cf2fec53ed5b5c5a0f7ac61d57920dcfa1fdf74f12cde96 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/activities"
    i5b1f031b640ea2281c9362cd8082a0af9981f17dcfa35f70a808632b0829d163 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/analytics"
    i5fc52960312d3658e9878fec07ef10fa59c577acefd43d96b21aa1b84732aa26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/subscriptions"
    i600975f3534d1ea4241cb6c791247f46593656198a707bbcebe6fe3bffcbb09c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/permissions"
    i91320585bcc095fb2de73064da160983c10e008c8636d74fbfa7eabb080e2f17 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/content"
    ib25f66d01014ea99569893cd772a349bb6ed502ca5f79043b92b7038dad21b16 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/thumbnails"
    ie641a8a15d2c7b2fb3c0e8c6b704eb3a4067227db3812d39038dbde75c71832a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/versions"
    iee48a5d4c8d90be7a26dec692c5db72099ed157dcec87fc370e76c6d060218b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem"
    i2a59fa1ef01cdd89020e91ee3a36c442e8b0e5ebbc70360d7b2fa1d1e130990c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/thumbnails/item"
    i534b450563a8e4c4464b761e767155c87f01b0cbf0631a847aa9963fbc37688b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/subscriptions/item"
    i9286effe7e06a1d7d8558089a8bfab151e9320f926daaf8eec9214c2450d8a0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/versions/item"
    iaf57a27b8c569c217bbf62c0dd97e220e61c94ae87e0ba7da939975ba64068ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/children/item"
    iafcaeb335a32fbc07bd158fc44c49edc5f8df815824b7bdfe3dfb501b9b3e141 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/activities/item"
    ie1c09b499780b5a2486b5092d3a39e70dff6e10ddc3497fbc29d44abeacb918b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/permissions/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i2b124f40c66d4ad62cf2fec53ed5b5c5a0f7ac61d57920dcfa1fdf74f12cde96.ActivitiesRequestBuilder) {
    return i2b124f40c66d4ad62cf2fec53ed5b5c5a0f7ac61d57920dcfa1fdf74f12cde96.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*iafcaeb335a32fbc07bd158fc44c49edc5f8df815824b7bdfe3dfb501b9b3e141.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return iafcaeb335a32fbc07bd158fc44c49edc5f8df815824b7bdfe3dfb501b9b3e141.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i5b1f031b640ea2281c9362cd8082a0af9981f17dcfa35f70a808632b0829d163.AnalyticsRequestBuilder) {
    return i5b1f031b640ea2281c9362cd8082a0af9981f17dcfa35f70a808632b0829d163.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i250757ab11c7812509a51c9b493b53cd482a62df2de58cf587bf23efffd92e3a.ChildrenRequestBuilder) {
    return i250757ab11c7812509a51c9b493b53cd482a62df2de58cf587bf23efffd92e3a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*iaf57a27b8c569c217bbf62c0dd97e220e61c94ae87e0ba7da939975ba64068ed.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return iaf57a27b8c569c217bbf62c0dd97e220e61c94ae87e0ba7da939975ba64068ed.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/following/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i91320585bcc095fb2de73064da160983c10e008c8636d74fbfa7eabb080e2f17.ContentRequestBuilder) {
    return i91320585bcc095fb2de73064da160983c10e008c8636d74fbfa7eabb080e2f17.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property following for users
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
// CreatePatchRequestInformation update the navigation property following in users
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
// Delete delete navigation property following for users
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
// Get the list of items the user is following. Only in OneDrive for Business.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*iee48a5d4c8d90be7a26dec692c5db72099ed157dcec87fc370e76c6d060218b8.ListItemRequestBuilder) {
    return iee48a5d4c8d90be7a26dec692c5db72099ed157dcec87fc370e76c6d060218b8.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property following in users
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i600975f3534d1ea4241cb6c791247f46593656198a707bbcebe6fe3bffcbb09c.PermissionsRequestBuilder) {
    return i600975f3534d1ea4241cb6c791247f46593656198a707bbcebe6fe3bffcbb09c.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ie1c09b499780b5a2486b5092d3a39e70dff6e10ddc3497fbc29d44abeacb918b.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ie1c09b499780b5a2486b5092d3a39e70dff6e10ddc3497fbc29d44abeacb918b.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i5fc52960312d3658e9878fec07ef10fa59c577acefd43d96b21aa1b84732aa26.SubscriptionsRequestBuilder) {
    return i5fc52960312d3658e9878fec07ef10fa59c577acefd43d96b21aa1b84732aa26.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i534b450563a8e4c4464b761e767155c87f01b0cbf0631a847aa9963fbc37688b.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i534b450563a8e4c4464b761e767155c87f01b0cbf0631a847aa9963fbc37688b.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*ib25f66d01014ea99569893cd772a349bb6ed502ca5f79043b92b7038dad21b16.ThumbnailsRequestBuilder) {
    return ib25f66d01014ea99569893cd772a349bb6ed502ca5f79043b92b7038dad21b16.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i2a59fa1ef01cdd89020e91ee3a36c442e8b0e5ebbc70360d7b2fa1d1e130990c.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i2a59fa1ef01cdd89020e91ee3a36c442e8b0e5ebbc70360d7b2fa1d1e130990c.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*ie641a8a15d2c7b2fb3c0e8c6b704eb3a4067227db3812d39038dbde75c71832a.VersionsRequestBuilder) {
    return ie641a8a15d2c7b2fb3c0e8c6b704eb3a4067227db3812d39038dbde75c71832a.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i9286effe7e06a1d7d8558089a8bfab151e9320f926daaf8eec9214c2450d8a0e.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i9286effe7e06a1d7d8558089a8bfab151e9320f926daaf8eec9214c2450d8a0e.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
