package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1f85445d7b444008169eaf9327d51601f1bd336ea6914bd28ad55e38055756dc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/children"
    i3024b2f2227c6e3745be71888c36c5f9177a6d0af76447dfde3c08334d0096f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/content"
    i5314f02a2cce55a2a6d7270bad8ea7bfd4c5a6a7e8cecc087a3d5d56070d4f58 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/versions"
    i645235fcee1fba5343c927aafcbde61f92f4bff639152c10e38ee8051b7ff665 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/thumbnails"
    i65d8a0619d5f103c25e9381e1e4b63f5813f1dc7411ab1e85b959ccd0bb4ec88 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/activities"
    i79a4d22c5f98f107d515f155aeb20473db4468d0690036f40d1333bf07b788e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem"
    i830b9b8f511775afb279623c558ebbbdc3aaaebe8886a0179c44fd05545b348e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/analytics"
    icca8a6c42f25adbe90b397ee92929d1e5988c719fe2892db2b496bb136504903 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/subscriptions"
    id2daeeee03b1b1f52891c3dd722517cfc03bd31b6c4aa6df7ac1d23b5cd6da2c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/permissions"
    i0962ce6ad456bc20ef85b5fbf650f35dbd3f6fc95cc9eed69bfd84132a5a8b93 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/permissions/item"
    i3eee4a59967d03ee4e62df2454d5cd526ec255059e3a5b90c76e3542bcd8fc76 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/activities/item"
    i665f843442a3271f860ec40675873e4a5d9487c20fe3b607978e481ab35281f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/children/item"
    i8e00377406e1d18c76b85d6bea246e5a8c838a4bbfb7ddb80d2811b8ba979197 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/subscriptions/item"
    iec8444bfe140eb58322a98cfc4e3136836b081c08d175ae6c2e4e56a49f66920 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/versions/item"
    if438d71550c9b3c8e4c16cefd7d2888338ecaa58626bb8b5b0d66feff04336df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/thumbnails/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i65d8a0619d5f103c25e9381e1e4b63f5813f1dc7411ab1e85b959ccd0bb4ec88.ActivitiesRequestBuilder) {
    return i65d8a0619d5f103c25e9381e1e4b63f5813f1dc7411ab1e85b959ccd0bb4ec88.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i3eee4a59967d03ee4e62df2454d5cd526ec255059e3a5b90c76e3542bcd8fc76.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i3eee4a59967d03ee4e62df2454d5cd526ec255059e3a5b90c76e3542bcd8fc76.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i830b9b8f511775afb279623c558ebbbdc3aaaebe8886a0179c44fd05545b348e.AnalyticsRequestBuilder) {
    return i830b9b8f511775afb279623c558ebbbdc3aaaebe8886a0179c44fd05545b348e.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i1f85445d7b444008169eaf9327d51601f1bd336ea6914bd28ad55e38055756dc.ChildrenRequestBuilder) {
    return i1f85445d7b444008169eaf9327d51601f1bd336ea6914bd28ad55e38055756dc.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i665f843442a3271f860ec40675873e4a5d9487c20fe3b607978e481ab35281f6.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i665f843442a3271f860ec40675873e4a5d9487c20fe3b607978e481ab35281f6.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/bundles/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i3024b2f2227c6e3745be71888c36c5f9177a6d0af76447dfde3c08334d0096f7.ContentRequestBuilder) {
    return i3024b2f2227c6e3745be71888c36c5f9177a6d0af76447dfde3c08334d0096f7.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property bundles for users
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
// CreatePatchRequestInformation update the navigation property bundles in users
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
// Delete delete navigation property bundles for users
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
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
// Get collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable, error) {
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i79a4d22c5f98f107d515f155aeb20473db4468d0690036f40d1333bf07b788e8.ListItemRequestBuilder) {
    return i79a4d22c5f98f107d515f155aeb20473db4468d0690036f40d1333bf07b788e8.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bundles in users
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*id2daeeee03b1b1f52891c3dd722517cfc03bd31b6c4aa6df7ac1d23b5cd6da2c.PermissionsRequestBuilder) {
    return id2daeeee03b1b1f52891c3dd722517cfc03bd31b6c4aa6df7ac1d23b5cd6da2c.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i0962ce6ad456bc20ef85b5fbf650f35dbd3f6fc95cc9eed69bfd84132a5a8b93.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i0962ce6ad456bc20ef85b5fbf650f35dbd3f6fc95cc9eed69bfd84132a5a8b93.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*icca8a6c42f25adbe90b397ee92929d1e5988c719fe2892db2b496bb136504903.SubscriptionsRequestBuilder) {
    return icca8a6c42f25adbe90b397ee92929d1e5988c719fe2892db2b496bb136504903.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i8e00377406e1d18c76b85d6bea246e5a8c838a4bbfb7ddb80d2811b8ba979197.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i8e00377406e1d18c76b85d6bea246e5a8c838a4bbfb7ddb80d2811b8ba979197.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i645235fcee1fba5343c927aafcbde61f92f4bff639152c10e38ee8051b7ff665.ThumbnailsRequestBuilder) {
    return i645235fcee1fba5343c927aafcbde61f92f4bff639152c10e38ee8051b7ff665.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*if438d71550c9b3c8e4c16cefd7d2888338ecaa58626bb8b5b0d66feff04336df.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return if438d71550c9b3c8e4c16cefd7d2888338ecaa58626bb8b5b0d66feff04336df.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i5314f02a2cce55a2a6d7270bad8ea7bfd4c5a6a7e8cecc087a3d5d56070d4f58.VersionsRequestBuilder) {
    return i5314f02a2cce55a2a6d7270bad8ea7bfd4c5a6a7e8cecc087a3d5d56070d4f58.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*iec8444bfe140eb58322a98cfc4e3136836b081c08d175ae6c2e4e56a49f66920.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return iec8444bfe140eb58322a98cfc4e3136836b081c08d175ae6c2e4e56a49f66920.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
