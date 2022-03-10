package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i03d3ef7585a04ffb6af1445f4c82e2a3830ad11d34279ecc27458e831073c3e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/listitem"
    i14f2dc5d0fbf74fa69792cff00f5cbb05c8e491ef9d5ec9c8c4d4b229830311b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/activities"
    i5828d42a5b7bc96fb77988943f2bafe518edbc68745bfcc284c71895b03433f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/subscriptions"
    i5addea073d7bd942a21ce268311e1949821336e25c39da2b6c7453cb9d50d9bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/thumbnails"
    i9080bc4aa517b5d65531f488af21d6cd8582dc413c24e81a1a002cf3f738ed9d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/permissions"
    i98c3a04252c41ecddc4f6f22880b144390b95d90ab222719c60ec749d7a101a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/content"
    ia5cc64d29a3cc6f9d5e3fcc5dc0941d9c84391f20a6552a9d8770ac04a586dc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/analytics"
    ie75c3da2009f092dbe2fa2bd59b684f396414096ee20d1f9ddc52b9019185b9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/children"
    if8de1919c5bfd5723f44ae5f7f39e525ca98171284e281b57d90c1416593eeb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/versions"
    i0e36ac8442010f7d4230b14141622b9c10daf5e21e0f5e047c3af7d66f867048 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/children/item"
    iae4363252131b22a43cdaea4832103bed61cffcc7c04a55e0784b2015bfb0053 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/permissions/item"
    ib876094e22727a16b40c61d63196401f141955d237499c3126ba8fb85ddbba2b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/versions/item"
    id1b64a8c282630d242ecda3c6831b919bf9f4a85787f8579aee92d4e11cdc0a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/subscriptions/item"
    ie3c70d0db863eaf2274654766405e6fda64a8beaeed562a467d8fc7a6d600482 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/thumbnails/item"
    ie8e3fe9bc33531bc1b773b9d6b65b3f6f07bfb4739c4329da24dfd79d07b3679 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/special/item/activities/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i14f2dc5d0fbf74fa69792cff00f5cbb05c8e491ef9d5ec9c8c4d4b229830311b.ActivitiesRequestBuilder) {
    return i14f2dc5d0fbf74fa69792cff00f5cbb05c8e491ef9d5ec9c8c4d4b229830311b.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.special.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*ie8e3fe9bc33531bc1b773b9d6b65b3f6f07bfb4739c4329da24dfd79d07b3679.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ie8e3fe9bc33531bc1b773b9d6b65b3f6f07bfb4739c4329da24dfd79d07b3679.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*ia5cc64d29a3cc6f9d5e3fcc5dc0941d9c84391f20a6552a9d8770ac04a586dc2.AnalyticsRequestBuilder) {
    return ia5cc64d29a3cc6f9d5e3fcc5dc0941d9c84391f20a6552a9d8770ac04a586dc2.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*ie75c3da2009f092dbe2fa2bd59b684f396414096ee20d1f9ddc52b9019185b9a.ChildrenRequestBuilder) {
    return ie75c3da2009f092dbe2fa2bd59b684f396414096ee20d1f9ddc52b9019185b9a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.special.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i0e36ac8442010f7d4230b14141622b9c10daf5e21e0f5e047c3af7d66f867048.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i0e36ac8442010f7d4230b14141622b9c10daf5e21e0f5e047c3af7d66f867048.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/special/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i98c3a04252c41ecddc4f6f22880b144390b95d90ab222719c60ec749d7a101a0.ContentRequestBuilder) {
    return i98c3a04252c41ecddc4f6f22880b144390b95d90ab222719c60ec749d7a101a0.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property special for me
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
// CreatePatchRequestInformation update the navigation property special in me
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
// Delete delete navigation property special for me
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i03d3ef7585a04ffb6af1445f4c82e2a3830ad11d34279ecc27458e831073c3e1.ListItemRequestBuilder) {
    return i03d3ef7585a04ffb6af1445f4c82e2a3830ad11d34279ecc27458e831073c3e1.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property special in me
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i9080bc4aa517b5d65531f488af21d6cd8582dc413c24e81a1a002cf3f738ed9d.PermissionsRequestBuilder) {
    return i9080bc4aa517b5d65531f488af21d6cd8582dc413c24e81a1a002cf3f738ed9d.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.special.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*iae4363252131b22a43cdaea4832103bed61cffcc7c04a55e0784b2015bfb0053.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return iae4363252131b22a43cdaea4832103bed61cffcc7c04a55e0784b2015bfb0053.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i5828d42a5b7bc96fb77988943f2bafe518edbc68745bfcc284c71895b03433f5.SubscriptionsRequestBuilder) {
    return i5828d42a5b7bc96fb77988943f2bafe518edbc68745bfcc284c71895b03433f5.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.special.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*id1b64a8c282630d242ecda3c6831b919bf9f4a85787f8579aee92d4e11cdc0a6.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return id1b64a8c282630d242ecda3c6831b919bf9f4a85787f8579aee92d4e11cdc0a6.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i5addea073d7bd942a21ce268311e1949821336e25c39da2b6c7453cb9d50d9bd.ThumbnailsRequestBuilder) {
    return i5addea073d7bd942a21ce268311e1949821336e25c39da2b6c7453cb9d50d9bd.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.special.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ie3c70d0db863eaf2274654766405e6fda64a8beaeed562a467d8fc7a6d600482.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return ie3c70d0db863eaf2274654766405e6fda64a8beaeed562a467d8fc7a6d600482.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*if8de1919c5bfd5723f44ae5f7f39e525ca98171284e281b57d90c1416593eeb2.VersionsRequestBuilder) {
    return if8de1919c5bfd5723f44ae5f7f39e525ca98171284e281b57d90c1416593eeb2.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.special.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ib876094e22727a16b40c61d63196401f141955d237499c3126ba8fb85ddbba2b.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ib876094e22727a16b40c61d63196401f141955d237499c3126ba8fb85ddbba2b.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
