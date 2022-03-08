package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3bceff7bf53c4651563a7ee2cce8430d8d13d0fcccd27ce5176f3543757e0291 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/listitem"
    i824aea247105b9dc671dcd0d2ad0a831502fe321120faa322f1cef3c028abf99 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/subscriptions"
    i9ec6171826960700d13aac7ba32c640e71d062196972597202bebe7e07054680 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/content"
    ia0ba6020c9cda5671979d16bb2d4f926d1dfed3871ea96416db6bc040c020e33 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/activities"
    ib0af710f53af596b7dfe6633f416c37d418ece5314efc5c6d18cbcf80c6780ef "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/versions"
    ic41a6d845c8bfbf203188a872368dc869fb2f2444adefec4dfc6dedaf3693c09 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/children"
    ie0e1b9432f4a8e3b7dd8a26a7c1da79a1955f117617ee102e411f4e2089cc188 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/thumbnails"
    ie2f47471443a1a4b1c019c7eb64a03cc71dd86e3e4a4e7e7b883c4dadf89474c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/analytics"
    if3628debdab307765e2880cf97f79b058f28f4ab90fe427137f843491c01aa03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/permissions"
    i25423c0624ac8be0bbd2bf1e88c1c32bc17a1ef1df918229501c1d2d6e3f43da "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/thumbnails/item"
    i30cb8a4c63463072556c3ec5780c6ba9d82fec92c93213a97ee1e91c459936d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/permissions/item"
    i4af66d73aea0edb26c159836225709f121cbd61adb74b6085e0861ec68b8929a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/children/item"
    i77af1cb75ce5a024ec6da6864089dbcda5ddfbc6b5351e06d34ef4c091439069 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/activities/item"
    i9cc605d01af7691e74fee27aa6bb99c5fbf92de61cbd7d1150e7f28e71257c9d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/versions/item"
    ic587bae841abfb9a6d9b620f286bab4ead170ceab52847955a7ffa485566235e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/bundles/item/subscriptions/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*ia0ba6020c9cda5671979d16bb2d4f926d1dfed3871ea96416db6bc040c020e33.ActivitiesRequestBuilder) {
    return ia0ba6020c9cda5671979d16bb2d4f926d1dfed3871ea96416db6bc040c020e33.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.bundles.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i77af1cb75ce5a024ec6da6864089dbcda5ddfbc6b5351e06d34ef4c091439069.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i77af1cb75ce5a024ec6da6864089dbcda5ddfbc6b5351e06d34ef4c091439069.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*ie2f47471443a1a4b1c019c7eb64a03cc71dd86e3e4a4e7e7b883c4dadf89474c.AnalyticsRequestBuilder) {
    return ie2f47471443a1a4b1c019c7eb64a03cc71dd86e3e4a4e7e7b883c4dadf89474c.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*ic41a6d845c8bfbf203188a872368dc869fb2f2444adefec4dfc6dedaf3693c09.ChildrenRequestBuilder) {
    return ic41a6d845c8bfbf203188a872368dc869fb2f2444adefec4dfc6dedaf3693c09.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.bundles.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i4af66d73aea0edb26c159836225709f121cbd61adb74b6085e0861ec68b8929a.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i4af66d73aea0edb26c159836225709f121cbd61adb74b6085e0861ec68b8929a.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/bundles/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i9ec6171826960700d13aac7ba32c640e71d062196972597202bebe7e07054680.ContentRequestBuilder) {
    return i9ec6171826960700d13aac7ba32c640e71d062196972597202bebe7e07054680.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property bundles for me
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
// CreatePatchRequestInformation update the navigation property bundles in me
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
// Delete delete navigation property bundles for me
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i3bceff7bf53c4651563a7ee2cce8430d8d13d0fcccd27ce5176f3543757e0291.ListItemRequestBuilder) {
    return i3bceff7bf53c4651563a7ee2cce8430d8d13d0fcccd27ce5176f3543757e0291.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bundles in me
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
func (m *DriveItemItemRequestBuilder) Permissions()(*if3628debdab307765e2880cf97f79b058f28f4ab90fe427137f843491c01aa03.PermissionsRequestBuilder) {
    return if3628debdab307765e2880cf97f79b058f28f4ab90fe427137f843491c01aa03.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.bundles.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i30cb8a4c63463072556c3ec5780c6ba9d82fec92c93213a97ee1e91c459936d1.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i30cb8a4c63463072556c3ec5780c6ba9d82fec92c93213a97ee1e91c459936d1.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i824aea247105b9dc671dcd0d2ad0a831502fe321120faa322f1cef3c028abf99.SubscriptionsRequestBuilder) {
    return i824aea247105b9dc671dcd0d2ad0a831502fe321120faa322f1cef3c028abf99.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.bundles.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*ic587bae841abfb9a6d9b620f286bab4ead170ceab52847955a7ffa485566235e.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ic587bae841abfb9a6d9b620f286bab4ead170ceab52847955a7ffa485566235e.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*ie0e1b9432f4a8e3b7dd8a26a7c1da79a1955f117617ee102e411f4e2089cc188.ThumbnailsRequestBuilder) {
    return ie0e1b9432f4a8e3b7dd8a26a7c1da79a1955f117617ee102e411f4e2089cc188.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.bundles.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i25423c0624ac8be0bbd2bf1e88c1c32bc17a1ef1df918229501c1d2d6e3f43da.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i25423c0624ac8be0bbd2bf1e88c1c32bc17a1ef1df918229501c1d2d6e3f43da.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*ib0af710f53af596b7dfe6633f416c37d418ece5314efc5c6d18cbcf80c6780ef.VersionsRequestBuilder) {
    return ib0af710f53af596b7dfe6633f416c37d418ece5314efc5c6d18cbcf80c6780ef.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.bundles.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i9cc605d01af7691e74fee27aa6bb99c5fbf92de61cbd7d1150e7f28e71257c9d.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i9cc605d01af7691e74fee27aa6bb99c5fbf92de61cbd7d1150e7f28e71257c9d.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
