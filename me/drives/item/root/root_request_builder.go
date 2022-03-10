package root

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/subscriptions"
    i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/thumbnails"
    i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/permissions"
    i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/children"
    i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/analytics"
    ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/listitem"
    ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/activities"
    ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/content"
    iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/versions"
    i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/activities/item"
    i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/versions/item"
    i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/permissions/item"
    i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/subscriptions/item"
    id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/thumbnails/item"
    ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/root/children/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RootRequestBuilderDeleteOptions options for Delete
type RootRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RootRequestBuilderGetOptions options for Get
type RootRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RootRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RootRequestBuilderPatchOptions options for Patch
type RootRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RootRequestBuilder) Activities()(*ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f.ActivitiesRequestBuilder) {
    return ia783022e23d4a61455349d260caeaf7f827f148e0b229a292c2448b110933f0f.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i048cac4ee0d8d62b3c2b23511eb9186c6e2fdede9968406b07e88caf4df27ef6.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Analytics()(*i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9.AnalyticsRequestBuilder) {
    return i9c6e8a7fe1c9cc5b7e0ea6c10c1af6f88e7214ebfec50e64d9cecd6189cf62a9.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RootRequestBuilder) Children()(*i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a.ChildrenRequestBuilder) {
    return i73d2513943183cb86cb44c927bf9ec731b074a759959852f6226167de5219d1a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return ie83cf697fa749a07f5a8e20f45c1986602567c1c297e5c1127d6e513588956a7.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/root{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RootRequestBuilder) Content()(*ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa.ContentRequestBuilder) {
    return ifcde4c248d0fb16b7a3592bf94897136073d710d9aad05ea6da584f65e2631fa.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for me
func (m *RootRequestBuilder) CreateDeleteRequestInformation(options *RootRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation(options *RootRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property root in me
func (m *RootRequestBuilder) CreatePatchRequestInformation(options *RootRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property root for me
func (m *RootRequestBuilder) Delete(options *RootRequestBuilderDeleteOptions)(error) {
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
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get(options *RootRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable, error) {
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
func (m *RootRequestBuilder) ListItem()(*ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c.ListItemRequestBuilder) {
    return ia25ae513c8a06c69f8e9a2af04dc1fdb56a445d26c487329094eb799760f235c.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in me
func (m *RootRequestBuilder) Patch(options *RootRequestBuilderPatchOptions)(error) {
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
func (m *RootRequestBuilder) Permissions()(*i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea.PermissionsRequestBuilder) {
    return i67e4eb3c6d101d79ba06ed09a3844590642bdfdffbdee1ba9f29c99aff2eddea.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i794a610a5528f5f6d3c429be3e0f48213381e2de9f7f3c9f949ebf5497110518.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Subscriptions()(*i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980.SubscriptionsRequestBuilder) {
    return i04bf9604718e0ae60b87e876a89e06624e62df01683c4f5c9d5888f0d85fa980.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i7e84b7d2d7188c19f7cdd6dc65d61b7bb5b964d0cca98d42f9cb67fe0ccedb77.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Thumbnails()(*i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25.ThumbnailsRequestBuilder) {
    return i05036fda23d42f66b0c931bf9c67aa77ca384d14e8dbe2c3f399234ec0e22f25.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return id78a84f29c1764444a930c32f704f8900d71328fe66c44a8c731acd2089749ae.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Versions()(*iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca.VersionsRequestBuilder) {
    return iff014ae58e15af2c05d5f6a5712f20fbb33698c83bed3dfbac294a0e5ad0c6ca.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i6246ed490a551373dc53d9eb146e31953d1de9b86dc8fddc38ff9fd7cac5c8f4.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
