package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i05fcb316fe8de8ec51486cfbfedf66e2f948e1c7839866241ad0615f02835bc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/thumbnails"
    i0ef34d76b49b11e2cc91d52c3918b60c6461347a955086f8a4668b5c1884115d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/permissions"
    i11cbe0e280eff024374239d9ddb8cf77156984055d65eaad8fab807dd66e1aae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/versions"
    i411cc6e09415497fdcaa1efa4ce38754c0f75a2ac6753fbdec316664fbe34c0b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/activities"
    i63c50a017fbc22e2cfab55eef29fc378c5fdd84591137cd766eb414eceb38c8c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/subscriptions"
    i768a53a4eefa41148954f6756c859d3d0616dfbc7edb8e8921f5025f8d1963b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/children"
    i81fbc50e77c099218ed3b0741178d113225c43436f1fe49068bd1c6eb2009f7d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/listitem"
    i92b1041e6f84d8bda4abe49a03fdc4afec1a63b1db3ace7fbd971ffc47d1f410 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/content"
    icee77210cdf073e51a21d353033cb3c2f35776c104510c78c75f15c47eef8850 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/analytics"
    i1ab5999899aa3fd7ee3e46316600a5a25b3837375c5f90f7a333f4800a2d4003 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/thumbnails/item"
    i4789a52454c475f8919b773b9abaaa2e9cf99508170aecc16040d9c5b637d18c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/children/item"
    i7b0211cb3c4f4d84f1885c1116fbf9bba3501d9fd3e04d752fca445e07814a84 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/activities/item"
    i9ea2dbed4bcf95caa8551ea2dabc66a656bfac9511000b4e2df923ee6a69a8a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/subscriptions/item"
    icae928e76fa13710b4c26c6e335a50337fbf054fbd3fb5cdaba5c915062842bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/permissions/item"
    if88c41658dc5d006a9811d058af31eb81aa7982c1099cdf52c84b881e44e8e5d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/special/item/versions/item"
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
func (m *DriveItemItemRequestBuilder) Activities()(*i411cc6e09415497fdcaa1efa4ce38754c0f75a2ac6753fbdec316664fbe34c0b.ActivitiesRequestBuilder) {
    return i411cc6e09415497fdcaa1efa4ce38754c0f75a2ac6753fbdec316664fbe34c0b.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.special.item.activities.item collection
func (m *DriveItemItemRequestBuilder) ActivitiesById(id string)(*i7b0211cb3c4f4d84f1885c1116fbf9bba3501d9fd3e04d752fca445e07814a84.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i7b0211cb3c4f4d84f1885c1116fbf9bba3501d9fd3e04d752fca445e07814a84.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Analytics()(*icee77210cdf073e51a21d353033cb3c2f35776c104510c78c75f15c47eef8850.AnalyticsRequestBuilder) {
    return icee77210cdf073e51a21d353033cb3c2f35776c104510c78c75f15c47eef8850.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i768a53a4eefa41148954f6756c859d3d0616dfbc7edb8e8921f5025f8d1963b0.ChildrenRequestBuilder) {
    return i768a53a4eefa41148954f6756c859d3d0616dfbc7edb8e8921f5025f8d1963b0.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.special.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i4789a52454c475f8919b773b9abaaa2e9cf99508170aecc16040d9c5b637d18c.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i4789a52454c475f8919b773b9abaaa2e9cf99508170aecc16040d9c5b637d18c.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/special/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i92b1041e6f84d8bda4abe49a03fdc4afec1a63b1db3ace7fbd971ffc47d1f410.ContentRequestBuilder) {
    return i92b1041e6f84d8bda4abe49a03fdc4afec1a63b1db3ace7fbd971ffc47d1f410.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property special for users
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
// CreatePatchRequestInformation update the navigation property special in users
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
// Delete delete navigation property special for users
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
// Get collection of common folders available in OneDrive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i81fbc50e77c099218ed3b0741178d113225c43436f1fe49068bd1c6eb2009f7d.ListItemRequestBuilder) {
    return i81fbc50e77c099218ed3b0741178d113225c43436f1fe49068bd1c6eb2009f7d.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property special in users
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i0ef34d76b49b11e2cc91d52c3918b60c6461347a955086f8a4668b5c1884115d.PermissionsRequestBuilder) {
    return i0ef34d76b49b11e2cc91d52c3918b60c6461347a955086f8a4668b5c1884115d.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.special.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*icae928e76fa13710b4c26c6e335a50337fbf054fbd3fb5cdaba5c915062842bc.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return icae928e76fa13710b4c26c6e335a50337fbf054fbd3fb5cdaba5c915062842bc.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i63c50a017fbc22e2cfab55eef29fc378c5fdd84591137cd766eb414eceb38c8c.SubscriptionsRequestBuilder) {
    return i63c50a017fbc22e2cfab55eef29fc378c5fdd84591137cd766eb414eceb38c8c.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.special.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i9ea2dbed4bcf95caa8551ea2dabc66a656bfac9511000b4e2df923ee6a69a8a0.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i9ea2dbed4bcf95caa8551ea2dabc66a656bfac9511000b4e2df923ee6a69a8a0.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i05fcb316fe8de8ec51486cfbfedf66e2f948e1c7839866241ad0615f02835bc7.ThumbnailsRequestBuilder) {
    return i05fcb316fe8de8ec51486cfbfedf66e2f948e1c7839866241ad0615f02835bc7.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.special.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i1ab5999899aa3fd7ee3e46316600a5a25b3837375c5f90f7a333f4800a2d4003.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i1ab5999899aa3fd7ee3e46316600a5a25b3837375c5f90f7a333f4800a2d4003.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i11cbe0e280eff024374239d9ddb8cf77156984055d65eaad8fab807dd66e1aae.VersionsRequestBuilder) {
    return i11cbe0e280eff024374239d9ddb8cf77156984055d65eaad8fab807dd66e1aae.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.special.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*if88c41658dc5d006a9811d058af31eb81aa7982c1099cdf52c84b881e44e8e5d.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return if88c41658dc5d006a9811d058af31eb81aa7982c1099cdf52c84b881e44e8e5d.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
