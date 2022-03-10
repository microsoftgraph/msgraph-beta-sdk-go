package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i2f9a3172fd1b51c1bea27737325db0e817816bac961b3ca6e9d0ad169a4d025b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/permissions"
    i690b0c2e2f8c9b97b29679ae0a3dde1e0e0e7c55bdc83c018c3c6bc7d49cc6ac "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/subscriptions"
    i7a443dc5adf92e91b05e83d537c8235dfa4550d8689b21b436e16f3fc0366532 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/thumbnails"
    i9e57fd00a48a7440ab50524ceb3fcd3f266c40c8bbefafe635a4913b7c88f04c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/children"
    ia4b4d085e98c9fe1d8a854f8e938bad6085d4eed307a0c9c697fca799b050fa1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/analytics"
    iac2eb2db213420a0252171e58f01253492b43f4150ef05f367daaee3c80091ee "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/versions"
    id85876844dd26e54006037fcbfc333e87bd686c7de16e7a85643a44321e1162d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/activities"
    idb34cd78d533a28d4bf3a7e0f6b1fd2d5b3ba9a706d5d871c91d8b81bcc2c475 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/content"
    ie7409f7388bced8a68b1f1e99777259a1a78c08c8d38410379ef58a1f9f1c651 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/listitem"
    i0ed0bed2d8ee8242dbe484c49b5735ec6d13565a80dfdd2b00cb7f898b066dc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/permissions/item"
    i43ea60510c1a10e6587cf98a6f3d5264d67f6423ae1b81bab847fad0b420c3f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/activities/item"
    i473d56912cdd4b50e205ab3f5776ce97566122776219130f4fb6bc97fe8e18fc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/subscriptions/item"
    i5c9ced66843ba5bf6b891588da20c13205ea07a9c58389447e07f32722e68ae3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/thumbnails/item"
    i9f6981643174c32ba54fffe5bcd80c7bb57399beccbcf2e87f51d2416118daa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/children/item"
    iee2957293d232b0fcc850038468a8092f28cb63be48244bfc74073195be31ca8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/bundles/item/versions/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*id85876844dd26e54006037fcbfc333e87bd686c7de16e7a85643a44321e1162d.ActivitiesRequestBuilder) {
    return id85876844dd26e54006037fcbfc333e87bd686c7de16e7a85643a44321e1162d.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.bundles.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i43ea60510c1a10e6587cf98a6f3d5264d67f6423ae1b81bab847fad0b420c3f2.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i43ea60510c1a10e6587cf98a6f3d5264d67f6423ae1b81bab847fad0b420c3f2.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*ia4b4d085e98c9fe1d8a854f8e938bad6085d4eed307a0c9c697fca799b050fa1.AnalyticsRequestBuilder) {
    return ia4b4d085e98c9fe1d8a854f8e938bad6085d4eed307a0c9c697fca799b050fa1.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i9e57fd00a48a7440ab50524ceb3fcd3f266c40c8bbefafe635a4913b7c88f04c.ChildrenRequestBuilder) {
    return i9e57fd00a48a7440ab50524ceb3fcd3f266c40c8bbefafe635a4913b7c88f04c.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.bundles.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i9f6981643174c32ba54fffe5bcd80c7bb57399beccbcf2e87f51d2416118daa9.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i9f6981643174c32ba54fffe5bcd80c7bb57399beccbcf2e87f51d2416118daa9.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}/bundles/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*idb34cd78d533a28d4bf3a7e0f6b1fd2d5b3ba9a706d5d871c91d8b81bcc2c475.ContentRequestBuilder) {
    return idb34cd78d533a28d4bf3a7e0f6b1fd2d5b3ba9a706d5d871c91d8b81bcc2c475.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property bundles for groups
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
// CreatePatchRequestInformation update the navigation property bundles in groups
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
// Delete delete navigation property bundles for groups
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
// Get collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*ie7409f7388bced8a68b1f1e99777259a1a78c08c8d38410379ef58a1f9f1c651.ListItemRequestBuilder) {
    return ie7409f7388bced8a68b1f1e99777259a1a78c08c8d38410379ef58a1f9f1c651.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bundles in groups
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i2f9a3172fd1b51c1bea27737325db0e817816bac961b3ca6e9d0ad169a4d025b.PermissionsRequestBuilder) {
    return i2f9a3172fd1b51c1bea27737325db0e817816bac961b3ca6e9d0ad169a4d025b.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.bundles.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i0ed0bed2d8ee8242dbe484c49b5735ec6d13565a80dfdd2b00cb7f898b066dc5.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i0ed0bed2d8ee8242dbe484c49b5735ec6d13565a80dfdd2b00cb7f898b066dc5.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i690b0c2e2f8c9b97b29679ae0a3dde1e0e0e7c55bdc83c018c3c6bc7d49cc6ac.SubscriptionsRequestBuilder) {
    return i690b0c2e2f8c9b97b29679ae0a3dde1e0e0e7c55bdc83c018c3c6bc7d49cc6ac.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.bundles.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i473d56912cdd4b50e205ab3f5776ce97566122776219130f4fb6bc97fe8e18fc.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i473d56912cdd4b50e205ab3f5776ce97566122776219130f4fb6bc97fe8e18fc.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i7a443dc5adf92e91b05e83d537c8235dfa4550d8689b21b436e16f3fc0366532.ThumbnailsRequestBuilder) {
    return i7a443dc5adf92e91b05e83d537c8235dfa4550d8689b21b436e16f3fc0366532.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.bundles.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i5c9ced66843ba5bf6b891588da20c13205ea07a9c58389447e07f32722e68ae3.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i5c9ced66843ba5bf6b891588da20c13205ea07a9c58389447e07f32722e68ae3.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*iac2eb2db213420a0252171e58f01253492b43f4150ef05f367daaee3c80091ee.VersionsRequestBuilder) {
    return iac2eb2db213420a0252171e58f01253492b43f4150ef05f367daaee3c80091ee.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.bundles.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*iee2957293d232b0fcc850038468a8092f28cb63be48244bfc74073195be31ca8.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return iee2957293d232b0fcc850038468a8092f28cb63be48244bfc74073195be31ca8.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
