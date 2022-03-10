package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i06e4181aced663adc7fbdb3cb4c988a7a6beb6c3785f12ab745919dd85125ab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/versions"
    i0a17cf731c23d6e4727ba8a73d33899fa4ea38fb37e6d735f98309b15a7c2b69 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/thumbnails"
    i1fbe4351522af91b57830e50c006bdb551c7ca3a3787ac321f0078b23d060a08 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/analytics"
    i453dec8d5b2506d83e2b2cc3f9a050eec11c4d706de7ea8bc7fbd8799f987208 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/permissions"
    i5144829ed9ea4cf7511d885c4395be6b0222317174dc7bb6ad18ee407b8168f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/children"
    i51fe5298187f35d39c337b7a44e27500dcf648d6420637b841399d59fdaebc0d "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/activities"
    i71b3e95c86c8f5ef399bf19667e2412ea828fbc57c43baad56914e673acc425c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/content"
    i738d8a11f18a95bc8e755d2094bf701b0a46d070f4f6c65e721127975f050488 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/listitem"
    ie19252fa188ad32a040657628f44a1c37e15da413764610bf6066b5c0ca51438 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/subscriptions"
    i417b4629f76be790e4865ec1d9cdadea49094c31935c0a53a108352ffa82d573 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/subscriptions/item"
    i69aeb96f0523670e5737d9348f72de9928f94935617cab59e3a1089428c456e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/permissions/item"
    i9aca5b5830e5d26d108fef7f0db5b38c4c4f9d803d0a42c7ff00b4ae2076da7b "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/versions/item"
    i9c03fb6e7773a4148610927c9ff52c7be7a85aca9ef7cf1f951a1c1cc320c384 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/activities/item"
    ic0ddbe55cde2c293b7c14edaea407621874d67b1853262b2710698d88cec3d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/thumbnails/item"
    id20ab52caa3b0e55bf6e7ff60caafbbe0e9eab65a2f37164e83ef97b0577dea0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item/children/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i51fe5298187f35d39c337b7a44e27500dcf648d6420637b841399d59fdaebc0d.ActivitiesRequestBuilder) {
    return i51fe5298187f35d39c337b7a44e27500dcf648d6420637b841399d59fdaebc0d.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.special.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i9c03fb6e7773a4148610927c9ff52c7be7a85aca9ef7cf1f951a1c1cc320c384.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i9c03fb6e7773a4148610927c9ff52c7be7a85aca9ef7cf1f951a1c1cc320c384.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i1fbe4351522af91b57830e50c006bdb551c7ca3a3787ac321f0078b23d060a08.AnalyticsRequestBuilder) {
    return i1fbe4351522af91b57830e50c006bdb551c7ca3a3787ac321f0078b23d060a08.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i5144829ed9ea4cf7511d885c4395be6b0222317174dc7bb6ad18ee407b8168f6.ChildrenRequestBuilder) {
    return i5144829ed9ea4cf7511d885c4395be6b0222317174dc7bb6ad18ee407b8168f6.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.special.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*id20ab52caa3b0e55bf6e7ff60caafbbe0e9eab65a2f37164e83ef97b0577dea0.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return id20ab52caa3b0e55bf6e7ff60caafbbe0e9eab65a2f37164e83ef97b0577dea0.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/special/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i71b3e95c86c8f5ef399bf19667e2412ea828fbc57c43baad56914e673acc425c.ContentRequestBuilder) {
    return i71b3e95c86c8f5ef399bf19667e2412ea828fbc57c43baad56914e673acc425c.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property special for drives
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
// CreatePatchRequestInformation update the navigation property special in drives
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
// Delete delete navigation property special for drives
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i738d8a11f18a95bc8e755d2094bf701b0a46d070f4f6c65e721127975f050488.ListItemRequestBuilder) {
    return i738d8a11f18a95bc8e755d2094bf701b0a46d070f4f6c65e721127975f050488.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property special in drives
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i453dec8d5b2506d83e2b2cc3f9a050eec11c4d706de7ea8bc7fbd8799f987208.PermissionsRequestBuilder) {
    return i453dec8d5b2506d83e2b2cc3f9a050eec11c4d706de7ea8bc7fbd8799f987208.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.special.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i69aeb96f0523670e5737d9348f72de9928f94935617cab59e3a1089428c456e4.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i69aeb96f0523670e5737d9348f72de9928f94935617cab59e3a1089428c456e4.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*ie19252fa188ad32a040657628f44a1c37e15da413764610bf6066b5c0ca51438.SubscriptionsRequestBuilder) {
    return ie19252fa188ad32a040657628f44a1c37e15da413764610bf6066b5c0ca51438.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.special.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i417b4629f76be790e4865ec1d9cdadea49094c31935c0a53a108352ffa82d573.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i417b4629f76be790e4865ec1d9cdadea49094c31935c0a53a108352ffa82d573.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i0a17cf731c23d6e4727ba8a73d33899fa4ea38fb37e6d735f98309b15a7c2b69.ThumbnailsRequestBuilder) {
    return i0a17cf731c23d6e4727ba8a73d33899fa4ea38fb37e6d735f98309b15a7c2b69.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.special.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ic0ddbe55cde2c293b7c14edaea407621874d67b1853262b2710698d88cec3d1a.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return ic0ddbe55cde2c293b7c14edaea407621874d67b1853262b2710698d88cec3d1a.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i06e4181aced663adc7fbdb3cb4c988a7a6beb6c3785f12ab745919dd85125ab5.VersionsRequestBuilder) {
    return i06e4181aced663adc7fbdb3cb4c988a7a6beb6c3785f12ab745919dd85125ab5.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.special.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i9aca5b5830e5d26d108fef7f0db5b38c4c4f9d803d0a42c7ff00b4ae2076da7b.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i9aca5b5830e5d26d108fef7f0db5b38c4c4f9d803d0a42c7ff00b4ae2076da7b.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
