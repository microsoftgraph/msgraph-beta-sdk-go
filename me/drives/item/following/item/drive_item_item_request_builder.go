package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i3ed1b892af177b0d7cdd9b7c6599b08776bf91d3aa86f60616b0be56aaad79fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/children"
    i5eda42b4dc64a218bb9b4e344d1aa34e98fe68ba3281d56179d85733d53a5999 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/versions"
    i7368f1651bc6e7600ec12237b08e2789ddf3742d30bdefeeb72ce73189b57b02 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/analytics"
    i8c085e9d30c43886d1bab77ddcd1e5c12ceb0912ffd1b434c307f085b0cd7db8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/activities"
    i95506f7e7c8681295c9f55e33cf51ae60db6b45350c51f738dc3d5f11dd7facc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/content"
    ia74ba45720dcac1ad0300840d4f3be6b99e8d82823672f551dfde01065301902 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/thumbnails"
    id6a99e2a1fc1d58414072f77a84af3805ab45a60a6df340f75144e1a6af42e36 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/permissions"
    idf465748d996abf478a73c535a5df3e2c86383a2c369fa7df06dfbae45a0e605 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/subscriptions"
    if75f81c885913a19312727d04ad3690434fc7a207d14351e72a48892e773f162 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem"
    i0b5810780844a8c0224c0ea05ddabc000f3e25cbd0d20806ac7e90f7a8f832b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/thumbnails/item"
    i8eb64db92239e210b7acaa271538c1f6933d73c5472dc0b49c2aa8e92c6e3883 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/activities/item"
    i9e8828e8126b189fe1ac71f69b8029824f950d0a85fde0b35d7c16a491d3017c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/subscriptions/item"
    iadfed5d624933ec7f74f0d8afb6320a6e874d23562a16345113e37afb065df6f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/permissions/item"
    ib6f778420c110584c1579e1da25fd778f0bf406feb60fee1dffda5c1b9b059da "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/versions/item"
    if47c54fe5c05e2959608b4777ba730155557c45d489021ad80c0eaa7cdfda17b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/children/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i8c085e9d30c43886d1bab77ddcd1e5c12ceb0912ffd1b434c307f085b0cd7db8.ActivitiesRequestBuilder) {
    return i8c085e9d30c43886d1bab77ddcd1e5c12ceb0912ffd1b434c307f085b0cd7db8.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i8eb64db92239e210b7acaa271538c1f6933d73c5472dc0b49c2aa8e92c6e3883.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i8eb64db92239e210b7acaa271538c1f6933d73c5472dc0b49c2aa8e92c6e3883.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i7368f1651bc6e7600ec12237b08e2789ddf3742d30bdefeeb72ce73189b57b02.AnalyticsRequestBuilder) {
    return i7368f1651bc6e7600ec12237b08e2789ddf3742d30bdefeeb72ce73189b57b02.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i3ed1b892af177b0d7cdd9b7c6599b08776bf91d3aa86f60616b0be56aaad79fc.ChildrenRequestBuilder) {
    return i3ed1b892af177b0d7cdd9b7c6599b08776bf91d3aa86f60616b0be56aaad79fc.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*if47c54fe5c05e2959608b4777ba730155557c45d489021ad80c0eaa7cdfda17b.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return if47c54fe5c05e2959608b4777ba730155557c45d489021ad80c0eaa7cdfda17b.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/following/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i95506f7e7c8681295c9f55e33cf51ae60db6b45350c51f738dc3d5f11dd7facc.ContentRequestBuilder) {
    return i95506f7e7c8681295c9f55e33cf51ae60db6b45350c51f738dc3d5f11dd7facc.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property following for me
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
// CreatePatchRequestInformation update the navigation property following in me
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
// Delete delete navigation property following for me
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
func (m *DriveItemItemRequestBuilder) ListItem()(*if75f81c885913a19312727d04ad3690434fc7a207d14351e72a48892e773f162.ListItemRequestBuilder) {
    return if75f81c885913a19312727d04ad3690434fc7a207d14351e72a48892e773f162.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property following in me
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
func (m *DriveItemItemRequestBuilder) Permissions()(*id6a99e2a1fc1d58414072f77a84af3805ab45a60a6df340f75144e1a6af42e36.PermissionsRequestBuilder) {
    return id6a99e2a1fc1d58414072f77a84af3805ab45a60a6df340f75144e1a6af42e36.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*iadfed5d624933ec7f74f0d8afb6320a6e874d23562a16345113e37afb065df6f.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return iadfed5d624933ec7f74f0d8afb6320a6e874d23562a16345113e37afb065df6f.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*idf465748d996abf478a73c535a5df3e2c86383a2c369fa7df06dfbae45a0e605.SubscriptionsRequestBuilder) {
    return idf465748d996abf478a73c535a5df3e2c86383a2c369fa7df06dfbae45a0e605.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i9e8828e8126b189fe1ac71f69b8029824f950d0a85fde0b35d7c16a491d3017c.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i9e8828e8126b189fe1ac71f69b8029824f950d0a85fde0b35d7c16a491d3017c.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*ia74ba45720dcac1ad0300840d4f3be6b99e8d82823672f551dfde01065301902.ThumbnailsRequestBuilder) {
    return ia74ba45720dcac1ad0300840d4f3be6b99e8d82823672f551dfde01065301902.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i0b5810780844a8c0224c0ea05ddabc000f3e25cbd0d20806ac7e90f7a8f832b0.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i0b5810780844a8c0224c0ea05ddabc000f3e25cbd0d20806ac7e90f7a8f832b0.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i5eda42b4dc64a218bb9b4e344d1aa34e98fe68ba3281d56179d85733d53a5999.VersionsRequestBuilder) {
    return i5eda42b4dc64a218bb9b4e344d1aa34e98fe68ba3281d56179d85733d53a5999.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ib6f778420c110584c1579e1da25fd778f0bf406feb60fee1dffda5c1b9b059da.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ib6f778420c110584c1579e1da25fd778f0bf406feb60fee1dffda5c1b9b059da.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
