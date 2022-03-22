package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities"
    i22f86b00dbed60c4ae66f22a96d25d718083d7ee61dda8addf35f98b8a81a481 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/operations"
    i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes"
    i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/columns"
    i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/drive"
    ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items"
    iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/subscriptions"
    i3b9fc89ed20f5bc6ee2709cbea9da1d72a73e1a237c50904ede0b8286e254701 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/operations/item"
    i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/subscriptions/item"
    i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/columns/item"
    i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item"
    i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item"
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
func (m *ListItemRequestBuilder) Activities()(*i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3.ActivitiesRequestBuilder) {
    return i03a20b5598fd1386b590353696f1734de0fd085772d233356819dfacb48398d3.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Columns()(*i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462.ColumnsRequestBuilder) {
    return i56769840e4943569ca7939554b4fb3706451d9fbe4e7f5691fd1542fe0b67462.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.columns.item collection
func (m *ListItemRequestBuilder) ColumnsById(id string)(*i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i84a6362366aa96a7cf24a63ec6fc8d616b93a84df4ef10a2dba84a08f8fae66c.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}{?select,expand}";
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
func (m *ListItemRequestBuilder) ContentTypes()(*i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31.ContentTypesRequestBuilder) {
    return i5445d6f5807fa1ccb62abcf3fcf2e362a866146d9d1678500a88c96beb6b4f31.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.contentTypes.item collection
func (m *ListItemRequestBuilder) ContentTypesById(id string)(*i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i853471746fafe0cd69c71c0d080ce487ac5cbe90e44c1a8e2f9f2b94a267d3bb.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property lists for sites
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
// CreatePatchRequestInformation update the navigation property lists in sites
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
// Delete delete navigation property lists for sites
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
func (m *ListItemRequestBuilder) Drive()(*i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f.DriveRequestBuilder) {
    return i8857dde88d0f1b1bcab36ac86f53cdf1cf5c8b4e4ce3497f1447b17884b4432f.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListItemRequestBuilder) Items()(*ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558.ItemsRequestBuilder) {
    return ib8242e278d08beecc8437de78cdbd8c65c699b45068fd7140d7cabf2c945f558.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.items.item collection
func (m *ListItemRequestBuilder) ItemsById(id string)(*i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return i877059848b5a1f38ddd1b25cd60b076bc0b0068b6a845b9627d81740e7eadef6.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Operations()(*i22f86b00dbed60c4ae66f22a96d25d718083d7ee61dda8addf35f98b8a81a481.OperationsRequestBuilder) {
    return i22f86b00dbed60c4ae66f22a96d25d718083d7ee61dda8addf35f98b8a81a481.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.operations.item collection
func (m *ListItemRequestBuilder) OperationsById(id string)(*i3b9fc89ed20f5bc6ee2709cbea9da1d72a73e1a237c50904ede0b8286e254701.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation_id"] = id
    }
    return i3b9fc89ed20f5bc6ee2709cbea9da1d72a73e1a237c50904ede0b8286e254701.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property lists in sites
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
func (m *ListItemRequestBuilder) Subscriptions()(*iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40.SubscriptionsRequestBuilder) {
    return iff76d0e8c7b3b7caff43b18695744b55ff20d7bdd54fb63bb6332cc1ae912a40.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.subscriptions.item collection
func (m *ListItemRequestBuilder) SubscriptionsById(id string)(*i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i574bb51b011ae57be83ae4c1378131adfd141ea2f28ea8d6fb084a7c2a316f9e.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
