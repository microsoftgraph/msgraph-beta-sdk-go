package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i08a11edc7565dc89ab6a3aa415589d5e2a5eab6c70e37bc3efd76c51b376effc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/children"
    i3e4abd10d11ac18d3839120caf3dabb3b47b79f70c8060f16cb542ae6f2ea875 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/permissions"
    i423953b40289ef2b307528d1d6f9c324712a0c85f78dc0e62a26840fd7d6d652 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/subscriptions"
    i689c3c6c08f355d2642d7d646547494d02fe15dcad7b574263872bdd1cfb27d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/activities"
    i79baf4e22f92b91854923facbad092b629260b25fc09361b9b373e9799ca63c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/thumbnails"
    i7e107463f1a8e4a62866358f8922333e55259a12b94a9e650b117df071050e62 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/analytics"
    i85e6b0e72a39339364371b81b37074fc7401cfb9d3e641f41e0179d90cbab06e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/listitem"
    ic0e678833416bf40ce9cbd9f84d6a3a08f35a56c3493d96672e8e743d26c68a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/versions"
    if38bb6762a649854de056b93161b9b4916a9ad2f99501d1c6f16b37dd9f8cf3d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/content"
    i2082b114e6d9cb9ec2f926fcfe205d300766ab8a74b84d6efb07f0397aaf683a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/thumbnails/item"
    i48c826eb30be650a8bba064364f5c715a721486b53714112b6b781f72f27724a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/versions/item"
    i5702202950600d844571a1d01cd3070c55fec0a0d3a1ace13a72ca1a4caa5127 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/subscriptions/item"
    i5f16d7365feb7f0a2700d025c9c50d3d316cb7b8c9c708c49987a65c13f1c834 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/children/item"
    i9a1fa10b89bcdd668d837e3c611f40aa2480178a5a3bcfee2779f7ab0c829a97 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/activities/item"
    icf27db966ca72ab106646c0ed1433e953e83bcdd6f5e900db1d70d158d87642f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/following/item/permissions/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i689c3c6c08f355d2642d7d646547494d02fe15dcad7b574263872bdd1cfb27d9.ActivitiesRequestBuilder) {
    return i689c3c6c08f355d2642d7d646547494d02fe15dcad7b574263872bdd1cfb27d9.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.following.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i9a1fa10b89bcdd668d837e3c611f40aa2480178a5a3bcfee2779f7ab0c829a97.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i9a1fa10b89bcdd668d837e3c611f40aa2480178a5a3bcfee2779f7ab0c829a97.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i7e107463f1a8e4a62866358f8922333e55259a12b94a9e650b117df071050e62.AnalyticsRequestBuilder) {
    return i7e107463f1a8e4a62866358f8922333e55259a12b94a9e650b117df071050e62.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i08a11edc7565dc89ab6a3aa415589d5e2a5eab6c70e37bc3efd76c51b376effc.ChildrenRequestBuilder) {
    return i08a11edc7565dc89ab6a3aa415589d5e2a5eab6c70e37bc3efd76c51b376effc.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.following.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i5f16d7365feb7f0a2700d025c9c50d3d316cb7b8c9c708c49987a65c13f1c834.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i5f16d7365feb7f0a2700d025c9c50d3d316cb7b8c9c708c49987a65c13f1c834.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}/following/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*if38bb6762a649854de056b93161b9b4916a9ad2f99501d1c6f16b37dd9f8cf3d.ContentRequestBuilder) {
    return if38bb6762a649854de056b93161b9b4916a9ad2f99501d1c6f16b37dd9f8cf3d.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property following for groups
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
// CreatePatchRequestInformation update the navigation property following in groups
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
// Delete delete navigation property following for groups
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i85e6b0e72a39339364371b81b37074fc7401cfb9d3e641f41e0179d90cbab06e.ListItemRequestBuilder) {
    return i85e6b0e72a39339364371b81b37074fc7401cfb9d3e641f41e0179d90cbab06e.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property following in groups
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i3e4abd10d11ac18d3839120caf3dabb3b47b79f70c8060f16cb542ae6f2ea875.PermissionsRequestBuilder) {
    return i3e4abd10d11ac18d3839120caf3dabb3b47b79f70c8060f16cb542ae6f2ea875.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.following.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*icf27db966ca72ab106646c0ed1433e953e83bcdd6f5e900db1d70d158d87642f.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return icf27db966ca72ab106646c0ed1433e953e83bcdd6f5e900db1d70d158d87642f.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i423953b40289ef2b307528d1d6f9c324712a0c85f78dc0e62a26840fd7d6d652.SubscriptionsRequestBuilder) {
    return i423953b40289ef2b307528d1d6f9c324712a0c85f78dc0e62a26840fd7d6d652.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.following.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i5702202950600d844571a1d01cd3070c55fec0a0d3a1ace13a72ca1a4caa5127.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i5702202950600d844571a1d01cd3070c55fec0a0d3a1ace13a72ca1a4caa5127.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i79baf4e22f92b91854923facbad092b629260b25fc09361b9b373e9799ca63c4.ThumbnailsRequestBuilder) {
    return i79baf4e22f92b91854923facbad092b629260b25fc09361b9b373e9799ca63c4.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.following.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i2082b114e6d9cb9ec2f926fcfe205d300766ab8a74b84d6efb07f0397aaf683a.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i2082b114e6d9cb9ec2f926fcfe205d300766ab8a74b84d6efb07f0397aaf683a.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*ic0e678833416bf40ce9cbd9f84d6a3a08f35a56c3493d96672e8e743d26c68a7.VersionsRequestBuilder) {
    return ic0e678833416bf40ce9cbd9f84d6a3a08f35a56c3493d96672e8e743d26c68a7.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.following.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i48c826eb30be650a8bba064364f5c715a721486b53714112b6b781f72f27724a.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i48c826eb30be650a8bba064364f5c715a721486b53714112b6b781f72f27724a.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
