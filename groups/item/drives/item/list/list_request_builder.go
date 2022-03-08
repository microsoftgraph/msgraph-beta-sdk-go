package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1271f32d5f7d3b0586f8de93a5b0535ce223cbfd2b48e6f411f71c399da18327 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/columns"
    i43e0b5ce13128879f06e2b07e0ccd9f8dc1dafb84ac39aa8ef8923e069dc6a8b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes"
    i7bfd5cc8c38cfc387360207e15abdf3e46f2fcf5d71bc51b310a210e16307965 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/subscriptions"
    i923c78037243955d3854651774d2282943abed938d26a8809587cc4b58252199 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/operations"
    ia7fc35da0de2bd686f2aa5a8a1fdfc88ac3a5bbd57126d8a1bd0644eceb39752 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/activities"
    ic55f9fa554d8a522a28b3fbd75b161bce003442994fcc35c64c623a12261a280 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/drive"
    idc21fe63ea818594f7f29929ecb8355791aa9fdd52478809de7eff177ca4c915 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/items"
    i649cdddb50a802c43dc3af1a06a94775bd06e4c17653045c20a5077da03ca95d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/columns/item"
    i998c00f0ec1e2b40c5518299a0453c926896385255ff91e4bf3d2b7234a467a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item"
    ia7e7318eb558bf8bfc7b2e73993fafd19e3292f1d7b06c2efa414823b5a839db "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/subscriptions/item"
    iae2fd07cb5857ac5712f84db0ed1eab2184e2ab848fee6addbad50e8b44d0302 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/operations/item"
    ib48fcf8c239e744cfc9fe4f1cd3da6fd77f080fc2f12a3108a7c772bb44e64bd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/activities/item"
    iba31399dc8f34e424c8ebc1f280f4f9e3e58432f001227720cab229ef435deaa "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/items/item"
)

// ListRequestBuilder provides operations to manage the list property of the microsoft.graph.drive entity.
type ListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListRequestBuilderDeleteOptions options for Delete
type ListRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListRequestBuilderGetOptions options for Get
type ListRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListRequestBuilderGetQueryParameters for drives in SharePoint, the underlying document library list. Read-only. Nullable.
type ListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListRequestBuilderPatchOptions options for Patch
type ListRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListRequestBuilder) Activities()(*ia7fc35da0de2bd686f2aa5a8a1fdfc88ac3a5bbd57126d8a1bd0644eceb39752.ActivitiesRequestBuilder) {
    return ia7fc35da0de2bd686f2aa5a8a1fdfc88ac3a5bbd57126d8a1bd0644eceb39752.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.activities.item collection
func (m *ListRequestBuilder) ActivitiesById(id string)(*ib48fcf8c239e744cfc9fe4f1cd3da6fd77f080fc2f12a3108a7c772bb44e64bd.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ib48fcf8c239e744cfc9fe4f1cd3da6fd77f080fc2f12a3108a7c772bb44e64bd.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Columns()(*i1271f32d5f7d3b0586f8de93a5b0535ce223cbfd2b48e6f411f71c399da18327.ColumnsRequestBuilder) {
    return i1271f32d5f7d3b0586f8de93a5b0535ce223cbfd2b48e6f411f71c399da18327.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*i649cdddb50a802c43dc3af1a06a94775bd06e4c17653045c20a5077da03ca95d.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i649cdddb50a802c43dc3af1a06a94775bd06e4c17653045c20a5077da03ca95d.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}/list{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListRequestBuilder instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListRequestBuilder) ContentTypes()(*i43e0b5ce13128879f06e2b07e0ccd9f8dc1dafb84ac39aa8ef8923e069dc6a8b.ContentTypesRequestBuilder) {
    return i43e0b5ce13128879f06e2b07e0ccd9f8dc1dafb84ac39aa8ef8923e069dc6a8b.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i998c00f0ec1e2b40c5518299a0453c926896385255ff91e4bf3d2b7234a467a8.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i998c00f0ec1e2b40c5518299a0453c926896385255ff91e4bf3d2b7234a467a8.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for groups
func (m *ListRequestBuilder) CreateDeleteRequestInformation(options *ListRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) CreateGetRequestInformation(options *ListRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property list in groups
func (m *ListRequestBuilder) CreatePatchRequestInformation(options *ListRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property list for groups
func (m *ListRequestBuilder) Delete(options *ListRequestBuilderDeleteOptions)(error) {
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
func (m *ListRequestBuilder) Drive()(*ic55f9fa554d8a522a28b3fbd75b161bce003442994fcc35c64c623a12261a280.DriveRequestBuilder) {
    return ic55f9fa554d8a522a28b3fbd75b161bce003442994fcc35c64c623a12261a280.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Get(options *ListRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateListFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable), nil
}
func (m *ListRequestBuilder) Items()(*idc21fe63ea818594f7f29929ecb8355791aa9fdd52478809de7eff177ca4c915.ItemsRequestBuilder) {
    return idc21fe63ea818594f7f29929ecb8355791aa9fdd52478809de7eff177ca4c915.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*iba31399dc8f34e424c8ebc1f280f4f9e3e58432f001227720cab229ef435deaa.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return iba31399dc8f34e424c8ebc1f280f4f9e3e58432f001227720cab229ef435deaa.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Operations()(*i923c78037243955d3854651774d2282943abed938d26a8809587cc4b58252199.OperationsRequestBuilder) {
    return i923c78037243955d3854651774d2282943abed938d26a8809587cc4b58252199.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*iae2fd07cb5857ac5712f84db0ed1eab2184e2ab848fee6addbad50e8b44d0302.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation_id"] = id
    }
    return iae2fd07cb5857ac5712f84db0ed1eab2184e2ab848fee6addbad50e8b44d0302.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in groups
func (m *ListRequestBuilder) Patch(options *ListRequestBuilderPatchOptions)(error) {
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
func (m *ListRequestBuilder) Subscriptions()(*i7bfd5cc8c38cfc387360207e15abdf3e46f2fcf5d71bc51b310a210e16307965.SubscriptionsRequestBuilder) {
    return i7bfd5cc8c38cfc387360207e15abdf3e46f2fcf5d71bc51b310a210e16307965.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*ia7e7318eb558bf8bfc7b2e73993fafd19e3292f1d7b06c2efa414823b5a839db.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ia7e7318eb558bf8bfc7b2e73993fafd19e3292f1d7b06c2efa414823b5a839db.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
