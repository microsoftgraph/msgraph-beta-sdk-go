package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i199f4fe84420b18a43deaf0667685d67c47afaf49a79254dcc4a317b5afca849 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes"
    i68ae183c09cc78b27dc0a97dc7978e0b869692de627d8e525f322762c2995563 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/activities"
    i7226c2ae89f801e4b64e6b0edd09181cf93ea776caaaeb9088097745135b7c70 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/drive"
    i8550117bc8dba52f0401a4c1cd2a6c36bd55a7d49e62e678bd768efc96140c4d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/operations"
    i8bcaa2dbdfb8a394ff13d6435d39b59476db4815609b2e8151ab695ec81a0265 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/columns"
    i97017b616da5986cf55b57864c53d1b56efbacdd6545a614ade2f0456a8c8e3e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/subscriptions"
    ia21f197c95ba8505493d74f1ef1d48a5400edd8511cb962781f10325527d4d30 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items"
    i781ff56458c20fa1b5ea82ecc4814d2f31e569406c20158e1b3ad4ed52d395bd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/columns/item"
    i8f66e52499614b6eebc4d2df4d0864612a4704ec85bded58c0834e237a19e822 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item"
    i932d70cc9aea88447f5a87d9f6a7d12fcbd584063696775f85bdbcb52661f2ff "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/operations/item"
    ib6346b85256456286c5c40862b2612677f66fb89b5675f3cdb91584973429793 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/subscriptions/item"
    ieb63f6ff8a52c6339b6fd4c8844e2b1f41f1144c131720fadba69f9d07938a2b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/items/item"
)

// ListItemRequestBuilder provides operations to manage the lists property of the microsoft.graph.site entity.
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetQueryParameters the collection of lists under this site.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemRequestBuilder) Activities()(*i68ae183c09cc78b27dc0a97dc7978e0b869692de627d8e525f322762c2995563.ActivitiesRequestBuilder) {
    return i68ae183c09cc78b27dc0a97dc7978e0b869692de627d8e525f322762c2995563.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Columns()(*i8bcaa2dbdfb8a394ff13d6435d39b59476db4815609b2e8151ab695ec81a0265.ColumnsRequestBuilder) {
    return i8bcaa2dbdfb8a394ff13d6435d39b59476db4815609b2e8151ab695ec81a0265.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.columns.item collection
func (m *ListItemRequestBuilder) ColumnsById(id string)(*i781ff56458c20fa1b5ea82ecc4814d2f31e569406c20158e1b3ad4ed52d395bd.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i781ff56458c20fa1b5ea82ecc4814d2f31e569406c20158e1b3ad4ed52d395bd.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/lists/{list_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListItemRequestBuilder) ContentTypes()(*i199f4fe84420b18a43deaf0667685d67c47afaf49a79254dcc4a317b5afca849.ContentTypesRequestBuilder) {
    return i199f4fe84420b18a43deaf0667685d67c47afaf49a79254dcc4a317b5afca849.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item collection
func (m *ListItemRequestBuilder) ContentTypesById(id string)(*i8f66e52499614b6eebc4d2df4d0864612a4704ec85bded58c0834e237a19e822.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i8f66e52499614b6eebc4d2df4d0864612a4704ec85bded58c0834e237a19e822.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property lists for groups
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of lists under this site.
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property lists in groups
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property lists for groups
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
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
func (m *ListItemRequestBuilder) Drive()(*i7226c2ae89f801e4b64e6b0edd09181cf93ea776caaaeb9088097745135b7c70.DriveRequestBuilder) {
    return i7226c2ae89f801e4b64e6b0edd09181cf93ea776caaaeb9088097745135b7c70.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of lists under this site.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateListFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable), nil
}
func (m *ListItemRequestBuilder) Items()(*ia21f197c95ba8505493d74f1ef1d48a5400edd8511cb962781f10325527d4d30.ItemsRequestBuilder) {
    return ia21f197c95ba8505493d74f1ef1d48a5400edd8511cb962781f10325527d4d30.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.items.item collection
func (m *ListItemRequestBuilder) ItemsById(id string)(*ieb63f6ff8a52c6339b6fd4c8844e2b1f41f1144c131720fadba69f9d07938a2b.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return ieb63f6ff8a52c6339b6fd4c8844e2b1f41f1144c131720fadba69f9d07938a2b.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Operations()(*i8550117bc8dba52f0401a4c1cd2a6c36bd55a7d49e62e678bd768efc96140c4d.OperationsRequestBuilder) {
    return i8550117bc8dba52f0401a4c1cd2a6c36bd55a7d49e62e678bd768efc96140c4d.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.operations.item collection
func (m *ListItemRequestBuilder) OperationsById(id string)(*i932d70cc9aea88447f5a87d9f6a7d12fcbd584063696775f85bdbcb52661f2ff.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation_id"] = id
    }
    return i932d70cc9aea88447f5a87d9f6a7d12fcbd584063696775f85bdbcb52661f2ff.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property lists in groups
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
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
func (m *ListItemRequestBuilder) Subscriptions()(*i97017b616da5986cf55b57864c53d1b56efbacdd6545a614ade2f0456a8c8e3e.SubscriptionsRequestBuilder) {
    return i97017b616da5986cf55b57864c53d1b56efbacdd6545a614ade2f0456a8c8e3e.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.subscriptions.item collection
func (m *ListItemRequestBuilder) SubscriptionsById(id string)(*ib6346b85256456286c5c40862b2612677f66fb89b5675f3cdb91584973429793.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ib6346b85256456286c5c40862b2612677f66fb89b5675f3cdb91584973429793.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
