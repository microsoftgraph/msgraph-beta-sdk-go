package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i217423aac1c727efdab2e5c54952470b0921f5b13d098bdf8b94800ab37d04ff "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/versions"
    i2960dd668f9d968328a1273c3905b10cd1454ba2be4d2d2518d741e39d34cb1a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/permissions"
    i465b2509f766b56413a989b9b1093c9dd885bd5ad9db93359e73becb56b4c8ed "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/subscriptions"
    i81e3ed87008521c326b3ac3c0fc74dbb4bb9dde9ae30e3d99226e6c7ba4eb8f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/thumbnails"
    i92430ec89e7e973775f796d4ac8a8d9780b1bc0d5ad60e7cd57a383e6eef3d5d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/children"
    ia251bf4684ce6c652f9eb3c25934c42fe972ed1dd2719fe5854f1afbc301666a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/listitem"
    iaeabdcfc742f6374216f8726dbe42c3ec126484f5bc5c2b32acc5da0596d09ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/content"
    idba818df722d655b17638c47b441ec5821743a8896803f37a5aab7f7e02a3c0a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/activities"
    idce2f25461c814df9cbddafba89d9ae2f557d73761223fcda23c80b4f5f3ee97 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/analytics"
    i461ab09785c4611bbf76605b4487a9f56eecb53b5d26bfb47eb14a0eadbe1bb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/children/item"
    i46b6ddeff5b488527d3295386cc098d2c87ca02641d9736e959f2cf64b66f45f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/versions/item"
    i55d3cce86d6128683aaeda467695b6cc98199af11165ef0493911461655f52da "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/subscriptions/item"
    i730e2e95d5c36abb56a8162182bfcf1bdb4f3d68c9721d0c03062451b2f30c28 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/thumbnails/item"
    i8b6ef88519f11f71a2ec7d5241d5e3086e28ed71a8d6b76a60af4f6ffdb80e54 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/activities/item"
    id992d48f7a3335ca926bd5d1d61ee50b229fce67be95588ae2b9110f96ad30d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/items/item/permissions/item"
)

// DriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
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
// DriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) Activities()(*idba818df722d655b17638c47b441ec5821743a8896803f37a5aab7f7e02a3c0a.ActivitiesRequestBuilder) {
    return idba818df722d655b17638c47b441ec5821743a8896803f37a5aab7f7e02a3c0a.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i8b6ef88519f11f71a2ec7d5241d5e3086e28ed71a8d6b76a60af4f6ffdb80e54.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i8b6ef88519f11f71a2ec7d5241d5e3086e28ed71a8d6b76a60af4f6ffdb80e54.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*idce2f25461c814df9cbddafba89d9ae2f557d73761223fcda23c80b4f5f3ee97.AnalyticsRequestBuilder) {
    return idce2f25461c814df9cbddafba89d9ae2f557d73761223fcda23c80b4f5f3ee97.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i92430ec89e7e973775f796d4ac8a8d9780b1bc0d5ad60e7cd57a383e6eef3d5d.ChildrenRequestBuilder) {
    return i92430ec89e7e973775f796d4ac8a8d9780b1bc0d5ad60e7cd57a383e6eef3d5d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i461ab09785c4611bbf76605b4487a9f56eecb53b5d26bfb47eb14a0eadbe1bb0.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i461ab09785c4611bbf76605b4487a9f56eecb53b5d26bfb47eb14a0eadbe1bb0.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}/items/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*iaeabdcfc742f6374216f8726dbe42c3ec126484f5bc5c2b32acc5da0596d09ea.ContentRequestBuilder) {
    return iaeabdcfc742f6374216f8726dbe42c3ec126484f5bc5c2b32acc5da0596d09ea.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for groups
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
// CreateGetRequestInformation all items contained in the drive. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property items in groups
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
// Delete delete navigation property items for groups
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
// Get all items contained in the drive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*ia251bf4684ce6c652f9eb3c25934c42fe972ed1dd2719fe5854f1afbc301666a.ListItemRequestBuilder) {
    return ia251bf4684ce6c652f9eb3c25934c42fe972ed1dd2719fe5854f1afbc301666a.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in groups
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i2960dd668f9d968328a1273c3905b10cd1454ba2be4d2d2518d741e39d34cb1a.PermissionsRequestBuilder) {
    return i2960dd668f9d968328a1273c3905b10cd1454ba2be4d2d2518d741e39d34cb1a.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*id992d48f7a3335ca926bd5d1d61ee50b229fce67be95588ae2b9110f96ad30d8.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return id992d48f7a3335ca926bd5d1d61ee50b229fce67be95588ae2b9110f96ad30d8.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i465b2509f766b56413a989b9b1093c9dd885bd5ad9db93359e73becb56b4c8ed.SubscriptionsRequestBuilder) {
    return i465b2509f766b56413a989b9b1093c9dd885bd5ad9db93359e73becb56b4c8ed.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i55d3cce86d6128683aaeda467695b6cc98199af11165ef0493911461655f52da.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i55d3cce86d6128683aaeda467695b6cc98199af11165ef0493911461655f52da.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i81e3ed87008521c326b3ac3c0fc74dbb4bb9dde9ae30e3d99226e6c7ba4eb8f5.ThumbnailsRequestBuilder) {
    return i81e3ed87008521c326b3ac3c0fc74dbb4bb9dde9ae30e3d99226e6c7ba4eb8f5.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i730e2e95d5c36abb56a8162182bfcf1bdb4f3d68c9721d0c03062451b2f30c28.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i730e2e95d5c36abb56a8162182bfcf1bdb4f3d68c9721d0c03062451b2f30c28.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i217423aac1c727efdab2e5c54952470b0921f5b13d098bdf8b94800ab37d04ff.VersionsRequestBuilder) {
    return i217423aac1c727efdab2e5c54952470b0921f5b13d098bdf8b94800ab37d04ff.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.items.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i46b6ddeff5b488527d3295386cc098d2c87ca02641d9736e959f2cf64b66f45f.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i46b6ddeff5b488527d3295386cc098d2c87ca02641d9736e959f2cf64b66f45f.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
