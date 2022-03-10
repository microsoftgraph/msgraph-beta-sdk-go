package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1607991deaac63c8e09452931980c50e39fb5d85b2b680303ed0c1b7aab85da5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/permissions"
    i16fd79586d66666bfa8483655fc5dc83cf4de92d8b1e3c85ebfc086e49d7d133 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/subscriptions"
    i2d2ee0c5a0c48d2f214c4e9a399fa034f460da5d4eda36df2caa9ec57bcf47f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/versions"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i46db24c18b46dfc1bf7543645d33f81fef36caa59eea9ccc2a51be3f75996c43 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/children"
    i5538b7b5579294d6c2d50146537199cc5b579ef148e3dc8e4531242f60cf43f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/content"
    ib8c5dbc73d0798574164df708e94b8fe0c1273073a00188ebf4a54111e239047 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/thumbnails"
    id13fe13effbf03349dc95c9cdc8fb2e0609273c5de93e85529a2c165fdb6ef9d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/activities"
    id308a9af810d367a6b4cf42d4d0a65c152dd4c70e58281180bc95d75eaea988d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/listitem"
    ie456329b50d31622829bcd088711018459fab31ecdc1068c0032297858cfc10b "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/analytics"
    i217e9e659a1a03591667f8725973486bc942c738c7172f79c9db1a3006911441 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/subscriptions/item"
    i8c2c771aa2b3bfa3ff727374d305f1c4d768926645c811be6fdcd32a84f323af "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/activities/item"
    i8ecc6481863e8f912332096d8f95f5df197e0aa4f894065d0c0509374716d4d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/permissions/item"
    ic1ae2326520e74920aae835c86535244b1b52a420d569e924be1009a3314a9c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/versions/item"
    icac371e2ba65d463ec35f422d1a660770e983ae477cea6c374c364966acf94c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/thumbnails/item"
    if61be009bca0eb802ce1050e1b967ba0c27cbb361d993b13a4fc5e39d63398b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item/children/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*id13fe13effbf03349dc95c9cdc8fb2e0609273c5de93e85529a2c165fdb6ef9d.ActivitiesRequestBuilder) {
    return id13fe13effbf03349dc95c9cdc8fb2e0609273c5de93e85529a2c165fdb6ef9d.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.special.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i8c2c771aa2b3bfa3ff727374d305f1c4d768926645c811be6fdcd32a84f323af.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i8c2c771aa2b3bfa3ff727374d305f1c4d768926645c811be6fdcd32a84f323af.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*ie456329b50d31622829bcd088711018459fab31ecdc1068c0032297858cfc10b.AnalyticsRequestBuilder) {
    return ie456329b50d31622829bcd088711018459fab31ecdc1068c0032297858cfc10b.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i46db24c18b46dfc1bf7543645d33f81fef36caa59eea9ccc2a51be3f75996c43.ChildrenRequestBuilder) {
    return i46db24c18b46dfc1bf7543645d33f81fef36caa59eea9ccc2a51be3f75996c43.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.special.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*if61be009bca0eb802ce1050e1b967ba0c27cbb361d993b13a4fc5e39d63398b3.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return if61be009bca0eb802ce1050e1b967ba0c27cbb361d993b13a4fc5e39d63398b3.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/special/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i5538b7b5579294d6c2d50146537199cc5b579ef148e3dc8e4531242f60cf43f1.ContentRequestBuilder) {
    return i5538b7b5579294d6c2d50146537199cc5b579ef148e3dc8e4531242f60cf43f1.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property special for drive
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
// CreatePatchRequestInformation update the navigation property special in drive
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
// Delete delete navigation property special for drive
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
// Get collection of common folders available in OneDrive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*id308a9af810d367a6b4cf42d4d0a65c152dd4c70e58281180bc95d75eaea988d.ListItemRequestBuilder) {
    return id308a9af810d367a6b4cf42d4d0a65c152dd4c70e58281180bc95d75eaea988d.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property special in drive
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i1607991deaac63c8e09452931980c50e39fb5d85b2b680303ed0c1b7aab85da5.PermissionsRequestBuilder) {
    return i1607991deaac63c8e09452931980c50e39fb5d85b2b680303ed0c1b7aab85da5.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.special.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i8ecc6481863e8f912332096d8f95f5df197e0aa4f894065d0c0509374716d4d7.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i8ecc6481863e8f912332096d8f95f5df197e0aa4f894065d0c0509374716d4d7.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i16fd79586d66666bfa8483655fc5dc83cf4de92d8b1e3c85ebfc086e49d7d133.SubscriptionsRequestBuilder) {
    return i16fd79586d66666bfa8483655fc5dc83cf4de92d8b1e3c85ebfc086e49d7d133.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.special.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i217e9e659a1a03591667f8725973486bc942c738c7172f79c9db1a3006911441.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i217e9e659a1a03591667f8725973486bc942c738c7172f79c9db1a3006911441.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*ib8c5dbc73d0798574164df708e94b8fe0c1273073a00188ebf4a54111e239047.ThumbnailsRequestBuilder) {
    return ib8c5dbc73d0798574164df708e94b8fe0c1273073a00188ebf4a54111e239047.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.special.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*icac371e2ba65d463ec35f422d1a660770e983ae477cea6c374c364966acf94c2.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return icac371e2ba65d463ec35f422d1a660770e983ae477cea6c374c364966acf94c2.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i2d2ee0c5a0c48d2f214c4e9a399fa034f460da5d4eda36df2caa9ec57bcf47f1.VersionsRequestBuilder) {
    return i2d2ee0c5a0c48d2f214c4e9a399fa034f460da5d4eda36df2caa9ec57bcf47f1.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.special.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ic1ae2326520e74920aae835c86535244b1b52a420d569e924be1009a3314a9c8.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ic1ae2326520e74920aae835c86535244b1b52a420d569e924be1009a3314a9c8.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
