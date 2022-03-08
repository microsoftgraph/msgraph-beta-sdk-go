package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0b07c5dc9fe1626f9f4538802c6d2503f44649bbce2abf1cdd3434d275e4d245 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes"
    i0da7cc6a659d17723767fb74f5241600495eac2b59b18630144f4e423f4759fe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/activities"
    iac853d10b6031957d1ad5d9e724c33018835d350736634c64b13458da21cd244 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/operations"
    ib21059ad8e9c7a50e12a7873696a669c9c8d6fe19e93558b2028f9a1ba858054 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/subscriptions"
    ibbef1fcee6a65e2ff0e41429b96ab9fba78712dfbd6979f758df51da2e6d0530 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/columns"
    ic1f8f80d30ba31bbeb7a0cb49b60889f514fbfe71b663edc153b8c89cf95e076 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/drive"
    ie83f05f3acccac621a64f5ea3a942ddbc686d1c58e4876bb65d0882dd77f0e29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items"
    i380f172145f2ac6a0882edbe65678687803aa5e1aadd56e82ca69d451f4a8946 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item"
    ibd46df32871447a12dd32dccbb92699ef0fa89ba47873fc3f1feb818300159c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/columns/item"
    ic145fd2ba7a9f49c1d068c64c8d0020a1f2b01a0bff310cf265690a936a92e40 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/activities/item"
    icbff86f1fe58ab2199799a9d9886ea90f1bd4b42f542de5eb3063a1f866849c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/subscriptions/item"
    ie597e3ec1515a5dc9491a6b53efc44bc794dbeaa9ac98c79c055643313191ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/operations/item"
    ife5b2b324259a801f94101e3426db3c970756094c9e6f8088049817f101b0774 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/items/item"
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
func (m *ListRequestBuilder) Activities()(*i0da7cc6a659d17723767fb74f5241600495eac2b59b18630144f4e423f4759fe.ActivitiesRequestBuilder) {
    return i0da7cc6a659d17723767fb74f5241600495eac2b59b18630144f4e423f4759fe.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.activities.item collection
func (m *ListRequestBuilder) ActivitiesById(id string)(*ic145fd2ba7a9f49c1d068c64c8d0020a1f2b01a0bff310cf265690a936a92e40.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ic145fd2ba7a9f49c1d068c64c8d0020a1f2b01a0bff310cf265690a936a92e40.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Columns()(*ibbef1fcee6a65e2ff0e41429b96ab9fba78712dfbd6979f758df51da2e6d0530.ColumnsRequestBuilder) {
    return ibbef1fcee6a65e2ff0e41429b96ab9fba78712dfbd6979f758df51da2e6d0530.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*ibd46df32871447a12dd32dccbb92699ef0fa89ba47873fc3f1feb818300159c9.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ibd46df32871447a12dd32dccbb92699ef0fa89ba47873fc3f1feb818300159c9.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/list{?select,expand}";
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
func (m *ListRequestBuilder) ContentTypes()(*i0b07c5dc9fe1626f9f4538802c6d2503f44649bbce2abf1cdd3434d275e4d245.ContentTypesRequestBuilder) {
    return i0b07c5dc9fe1626f9f4538802c6d2503f44649bbce2abf1cdd3434d275e4d245.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i380f172145f2ac6a0882edbe65678687803aa5e1aadd56e82ca69d451f4a8946.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i380f172145f2ac6a0882edbe65678687803aa5e1aadd56e82ca69d451f4a8946.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for users
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
// CreatePatchRequestInformation update the navigation property list in users
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
// Delete delete navigation property list for users
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
func (m *ListRequestBuilder) Drive()(*ic1f8f80d30ba31bbeb7a0cb49b60889f514fbfe71b663edc153b8c89cf95e076.DriveRequestBuilder) {
    return ic1f8f80d30ba31bbeb7a0cb49b60889f514fbfe71b663edc153b8c89cf95e076.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListRequestBuilder) Items()(*ie83f05f3acccac621a64f5ea3a942ddbc686d1c58e4876bb65d0882dd77f0e29.ItemsRequestBuilder) {
    return ie83f05f3acccac621a64f5ea3a942ddbc686d1c58e4876bb65d0882dd77f0e29.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*ife5b2b324259a801f94101e3426db3c970756094c9e6f8088049817f101b0774.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return ife5b2b324259a801f94101e3426db3c970756094c9e6f8088049817f101b0774.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Operations()(*iac853d10b6031957d1ad5d9e724c33018835d350736634c64b13458da21cd244.OperationsRequestBuilder) {
    return iac853d10b6031957d1ad5d9e724c33018835d350736634c64b13458da21cd244.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*ie597e3ec1515a5dc9491a6b53efc44bc794dbeaa9ac98c79c055643313191ea1.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation_id"] = id
    }
    return ie597e3ec1515a5dc9491a6b53efc44bc794dbeaa9ac98c79c055643313191ea1.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in users
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
func (m *ListRequestBuilder) Subscriptions()(*ib21059ad8e9c7a50e12a7873696a669c9c8d6fe19e93558b2028f9a1ba858054.SubscriptionsRequestBuilder) {
    return ib21059ad8e9c7a50e12a7873696a669c9c8d6fe19e93558b2028f9a1ba858054.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*icbff86f1fe58ab2199799a9d9886ea90f1bd4b42f542de5eb3063a1f866849c2.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return icbff86f1fe58ab2199799a9d9886ea90f1bd4b42f542de5eb3063a1f866849c2.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
