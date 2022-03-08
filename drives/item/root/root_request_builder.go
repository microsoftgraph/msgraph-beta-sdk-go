package root

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/children"
    i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/subscriptions"
    i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/thumbnails"
    i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/permissions"
    i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/content"
    i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/versions"
    i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem"
    i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/activities"
    ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/analytics"
    i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/versions/item"
    i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/children/item"
    i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/thumbnails/item"
    i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/subscriptions/item"
    ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/permissions/item"
    ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/activities/item"
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
func (m *RootRequestBuilder) Activities()(*i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309.ActivitiesRequestBuilder) {
    return i9cd042c9e46438c7bce87f358e11b3cca0d1ddc296c2c5c1db9328d1d7d50309.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ie50aea2d159f88c27cbd91bde6c1de53b048faf70798c48a700e279c1e5aff78.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Analytics()(*ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1.AnalyticsRequestBuilder) {
    return ib9cdff2592b0acec4aa284ce6ebe0a2cf32a5af41f645119e427abf571a30ea1.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RootRequestBuilder) Children()(*i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba.ChildrenRequestBuilder) {
    return i0cfb21b89f2be6662549c26a3797133b25184b3dfb5b75f7fcc39e668f1f52ba.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i4f1c46079242ae874121e4eaab792681034ad778ef6010021ce2f297e80d9e7e.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/root{?select,expand}";
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
func (m *RootRequestBuilder) Content()(*i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7.ContentRequestBuilder) {
    return i75f74a9f470214e89920f233558fd6a17cfa92ca6bf22dfcb4f98c21a1f009c7.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for drives
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
// CreatePatchRequestInformation update the navigation property root in drives
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
// Delete delete navigation property root for drives
func (m *RootRequestBuilder) Delete(options *RootRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
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
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable), nil
}
func (m *RootRequestBuilder) ListItem()(*i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d.ListItemRequestBuilder) {
    return i8556f9208dc03b23bb1904471ffd2bdfe79b4f5c2b4dfbfcb65bb05f00abab2d.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in drives
func (m *RootRequestBuilder) Patch(options *RootRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *RootRequestBuilder) Permissions()(*i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254.PermissionsRequestBuilder) {
    return i55c2129817d1116256382e18ea19914df03e1a7f25fde5c2e13a0c87abbcd254.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ic63da4ae39cf08296b5759a146dd61e14200041401e8eecb075724a9f1b76d19.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Subscriptions()(*i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967.SubscriptionsRequestBuilder) {
    return i36eabd21e28fa780f76254e70543251347fc0637b588934f3852063c63424967.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i7c142e142c53d798fd36169ec23884ef699397425bef56513abf167c332f6077.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Thumbnails()(*i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1.ThumbnailsRequestBuilder) {
    return i538aec6ec1a5dbcbc75dda7f47d60743e51c6eb965f1999aa7e8fa1b049b4fb1.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i57df13ad7fef0bc108a1f0f224b2282dbbdcc2899d051c8a902e93cbaa841679.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Versions()(*i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490.VersionsRequestBuilder) {
    return i844a09e2a848a0518b30d6c9486b0fc21ed3d53251643cfcfff453511ccef490.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i0ad9c3bbc7ea3c357a65b8cc5f691784c52ce76f23b1eeb54bebe4e2e38fe6f2.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
