package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i072ddba6d43643cb4e0a492cfc6003149b0f11a1bbb94fcbbdcf71f0770d542a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items"
    i18a4f9c70e9242b75dfae66cc04b68ef7851da26b5ea1fbcef0a5f8835a49e7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/activities"
    i1f28b50836f57b5d58bbaf6536fcecef381f989c4e9cb3ea42add732311ae102 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/subscriptions"
    i261ceb17d8b68c77c74467f613cdf5532cd5e5ea820c84d683a39d39c13c4ebe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/columns"
    ib0080b16f4fd1f64bad0fc40faa8191f83bc5af05b27dabb5054539619462021 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/drive"
    ic1ed073c992d668b0fff838202aaefd614de5739048c54fa321f3c0cf2913074 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/operations"
    icf07d83821e028c339f3d32b2cc131377b2d116e69b7fd99e7db53d773436c38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes"
    i0170c0ef9fee9daf3d24d9e190f3c34c7df524fa1ab8cd49794a581d983747d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/activities/item"
    i40668099437af2cc778494c140ef105574b3d11b369844ae96e6da4e56ed7dbb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/columns/item"
    i69531fc3cfd71461eb64582da70e123676f0660112f5d4521198ea126071edb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/operations/item"
    i8d0b9819c6c73c68982134065361fc5eb557175052e7f6f62cb6465cea9691ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item"
    i9afbcf26a085bae33bec61cc4b40f70ee21e783407b414e87aa6707a9555ac2a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item"
    ibdce946586319f7ecb5c54b6cd60e9e1e6b2a7b464a4783217eb2d2c5fbc3414 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/subscriptions/item"
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
func (m *ListRequestBuilder) Activities()(*i18a4f9c70e9242b75dfae66cc04b68ef7851da26b5ea1fbcef0a5f8835a49e7e.ActivitiesRequestBuilder) {
    return i18a4f9c70e9242b75dfae66cc04b68ef7851da26b5ea1fbcef0a5f8835a49e7e.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.activities.item collection
func (m *ListRequestBuilder) ActivitiesById(id string)(*i0170c0ef9fee9daf3d24d9e190f3c34c7df524fa1ab8cd49794a581d983747d1.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i0170c0ef9fee9daf3d24d9e190f3c34c7df524fa1ab8cd49794a581d983747d1.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Columns()(*i261ceb17d8b68c77c74467f613cdf5532cd5e5ea820c84d683a39d39c13c4ebe.ColumnsRequestBuilder) {
    return i261ceb17d8b68c77c74467f613cdf5532cd5e5ea820c84d683a39d39c13c4ebe.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*i40668099437af2cc778494c140ef105574b3d11b369844ae96e6da4e56ed7dbb.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i40668099437af2cc778494c140ef105574b3d11b369844ae96e6da4e56ed7dbb.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/list{?select,expand}";
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
func (m *ListRequestBuilder) ContentTypes()(*icf07d83821e028c339f3d32b2cc131377b2d116e69b7fd99e7db53d773436c38.ContentTypesRequestBuilder) {
    return icf07d83821e028c339f3d32b2cc131377b2d116e69b7fd99e7db53d773436c38.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i9afbcf26a085bae33bec61cc4b40f70ee21e783407b414e87aa6707a9555ac2a.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i9afbcf26a085bae33bec61cc4b40f70ee21e783407b414e87aa6707a9555ac2a.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for me
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
// CreatePatchRequestInformation update the navigation property list in me
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
// Delete delete navigation property list for me
func (m *ListRequestBuilder) Delete(options *ListRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListRequestBuilder) Drive()(*ib0080b16f4fd1f64bad0fc40faa8191f83bc5af05b27dabb5054539619462021.DriveRequestBuilder) {
    return ib0080b16f4fd1f64bad0fc40faa8191f83bc5af05b27dabb5054539619462021.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Get(options *ListRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateListFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable), nil
}
func (m *ListRequestBuilder) Items()(*i072ddba6d43643cb4e0a492cfc6003149b0f11a1bbb94fcbbdcf71f0770d542a.ItemsRequestBuilder) {
    return i072ddba6d43643cb4e0a492cfc6003149b0f11a1bbb94fcbbdcf71f0770d542a.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*i8d0b9819c6c73c68982134065361fc5eb557175052e7f6f62cb6465cea9691ad.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return i8d0b9819c6c73c68982134065361fc5eb557175052e7f6f62cb6465cea9691ad.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Operations()(*ic1ed073c992d668b0fff838202aaefd614de5739048c54fa321f3c0cf2913074.OperationsRequestBuilder) {
    return ic1ed073c992d668b0fff838202aaefd614de5739048c54fa321f3c0cf2913074.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*i69531fc3cfd71461eb64582da70e123676f0660112f5d4521198ea126071edb4.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation_id"] = id
    }
    return i69531fc3cfd71461eb64582da70e123676f0660112f5d4521198ea126071edb4.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in me
func (m *ListRequestBuilder) Patch(options *ListRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListRequestBuilder) Subscriptions()(*i1f28b50836f57b5d58bbaf6536fcecef381f989c4e9cb3ea42add732311ae102.SubscriptionsRequestBuilder) {
    return i1f28b50836f57b5d58bbaf6536fcecef381f989c4e9cb3ea42add732311ae102.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*ibdce946586319f7ecb5c54b6cd60e9e1e6b2a7b464a4783217eb2d2c5fbc3414.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ibdce946586319f7ecb5c54b6cd60e9e1e6b2a7b464a4783217eb2d2c5fbc3414.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
