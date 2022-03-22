package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i4d901644c467d6da1064f9218d2a51302b7e04affba7010edb0f37beacafb773 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/subscriptions"
    i53109b7297c409f46286c656c3579c1bb0aff65ea692d18c6998e4eca7e7cae1 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items"
    ib3ee4819f4c735d8a63e343645326309cf0357281fcf420c964b4b39d6bdf9bb "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/drive"
    ic1183cba85b0a8992427e9c7d45ad198b411eb949fdcf8bf5fa204e3fc691a5f "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/columns"
    icf939f29457e323541cf147d31bf6cecf5cb130c9d7c8639ec566b8175d1d5b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes"
    iec4aa5d59e58579afdbecfba73d6bd2372d1b99f1d8b3ee465c20b9750d2cabd "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/operations"
    if45c6d4c98602b6c309b1994f425724fef00fd5a63d747a5169844dea93fc758 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities"
    i55bc9c6f863d9539e65f9d6af8a6bfd60d606e59691d0de7981bc7a5fd1dad6c "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/operations/item"
    ib48ab6e9d0b3b2e8132b00fed04b5ed61b4e70dfa330c315456e7b72365c013e "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item"
    icc49cea450b5f3038e85b1a338cc21ce61681050bc96578d052ebb2a0cf235e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item"
    id75517f929096ab464028a948416a44d3bdb482418ed4adcbaa37d952f45ac81 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/subscriptions/item"
    ie6a239cd17154be4c0edbc83f7fe1c284635d98323d32d5dcce6210dfd425f2b "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/columns/item"
)

// ListRequestBuilder provides operations to manage the list property of the microsoft.graph.sharedDriveItem entity.
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
// ListRequestBuilderGetQueryParameters used to access the underlying list
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
func (m *ListRequestBuilder) Activities()(*if45c6d4c98602b6c309b1994f425724fef00fd5a63d747a5169844dea93fc758.ActivitiesRequestBuilder) {
    return if45c6d4c98602b6c309b1994f425724fef00fd5a63d747a5169844dea93fc758.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) Columns()(*ic1183cba85b0a8992427e9c7d45ad198b411eb949fdcf8bf5fa204e3fc691a5f.ColumnsRequestBuilder) {
    return ic1183cba85b0a8992427e9c7d45ad198b411eb949fdcf8bf5fa204e3fc691a5f.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*ie6a239cd17154be4c0edbc83f7fe1c284635d98323d32d5dcce6210dfd425f2b.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ie6a239cd17154be4c0edbc83f7fe1c284635d98323d32d5dcce6210dfd425f2b.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list{?select,expand}";
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
func (m *ListRequestBuilder) ContentTypes()(*icf939f29457e323541cf147d31bf6cecf5cb130c9d7c8639ec566b8175d1d5b1.ContentTypesRequestBuilder) {
    return icf939f29457e323541cf147d31bf6cecf5cb130c9d7c8639ec566b8175d1d5b1.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*icc49cea450b5f3038e85b1a338cc21ce61681050bc96578d052ebb2a0cf235e4.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return icc49cea450b5f3038e85b1a338cc21ce61681050bc96578d052ebb2a0cf235e4.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for shares
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
// CreateGetRequestInformation used to access the underlying list
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
// CreatePatchRequestInformation update the navigation property list in shares
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
// Delete delete navigation property list for shares
func (m *ListRequestBuilder) Delete(options *ListRequestBuilderDeleteOptions)(error) {
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
func (m *ListRequestBuilder) Drive()(*ib3ee4819f4c735d8a63e343645326309cf0357281fcf420c964b4b39d6bdf9bb.DriveRequestBuilder) {
    return ib3ee4819f4c735d8a63e343645326309cf0357281fcf420c964b4b39d6bdf9bb.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get used to access the underlying list
func (m *ListRequestBuilder) Get(options *ListRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Listable, error) {
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
func (m *ListRequestBuilder) Items()(*i53109b7297c409f46286c656c3579c1bb0aff65ea692d18c6998e4eca7e7cae1.ItemsRequestBuilder) {
    return i53109b7297c409f46286c656c3579c1bb0aff65ea692d18c6998e4eca7e7cae1.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*ib48ab6e9d0b3b2e8132b00fed04b5ed61b4e70dfa330c315456e7b72365c013e.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return ib48ab6e9d0b3b2e8132b00fed04b5ed61b4e70dfa330c315456e7b72365c013e.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Operations()(*iec4aa5d59e58579afdbecfba73d6bd2372d1b99f1d8b3ee465c20b9750d2cabd.OperationsRequestBuilder) {
    return iec4aa5d59e58579afdbecfba73d6bd2372d1b99f1d8b3ee465c20b9750d2cabd.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*i55bc9c6f863d9539e65f9d6af8a6bfd60d606e59691d0de7981bc7a5fd1dad6c.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation_id"] = id
    }
    return i55bc9c6f863d9539e65f9d6af8a6bfd60d606e59691d0de7981bc7a5fd1dad6c.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in shares
func (m *ListRequestBuilder) Patch(options *ListRequestBuilderPatchOptions)(error) {
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
func (m *ListRequestBuilder) Subscriptions()(*i4d901644c467d6da1064f9218d2a51302b7e04affba7010edb0f37beacafb773.SubscriptionsRequestBuilder) {
    return i4d901644c467d6da1064f9218d2a51302b7e04affba7010edb0f37beacafb773.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*id75517f929096ab464028a948416a44d3bdb482418ed4adcbaa37d952f45ac81.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return id75517f929096ab464028a948416a44d3bdb482418ed4adcbaa37d952f45ac81.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
