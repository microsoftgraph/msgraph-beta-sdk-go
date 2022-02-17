package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0f8ca929323429752e21f9acb28f80889948ea9ff72848249fca8c74c368589e "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities"
    i1c4ca68087caab22a544a9dc3fb1c29f6076a9264ac284502cbf04a28444aea7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/subscriptions"
    i2ef5930655c892d16078905ec81246034d704193e69907b59b2f14feab329b60 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id28cbea101cb9521a14cf9640c45800a00477bd62395913072f4e81bf1074012 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/drive"
    idf9a3059d40c90c42068fa8bf9ac9fd67dedf1182980c80d7a384967f9dff8d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items"
    ie43f83ef262031fdb1eb8da2814c4165f5f945383e15adf8746e6a96f3432e23 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/operations"
    if9ff29afb76bbf1ffef5a126d0ba635e14dcd31de06309e9c7c1a69ad2600738 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/columns"
    i327c97fc6449280089d733ba8dee83d9d74247819100cf6629d3355abfea1166 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/operations/item"
    i37771808e5dcfef82cc1e8bc33ad871d583b797abbb1222c5dc6198791d77af2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/subscriptions/item"
    i58ec931c55ab0d0111918ef65feb8b5a53f2f824a19e864696540521faf828f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item"
    i654b96f70860efab65ce04591400b15de588e186e6e7217fc1ca04b9de80627c "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/columns/item"
    i6dd3632bd204e05cac0aa8e9c03666774381075a56cb272b02bc729017380185 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item"
    iaea7c5cfa69bd0ae6b21223f98a830158fc5a2a0297fda11a70a877a663f4f26 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item"
)

// ListRequestBuilder builds and executes requests for operations under \drive\list
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
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListRequestBuilder) Activities()(*i0f8ca929323429752e21f9acb28f80889948ea9ff72848249fca8c74c368589e.ActivitiesRequestBuilder) {
    return i0f8ca929323429752e21f9acb28f80889948ea9ff72848249fca8c74c368589e.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.activities.item collection
func (m *ListRequestBuilder) ActivitiesById(id string)(*iaea7c5cfa69bd0ae6b21223f98a830158fc5a2a0297fda11a70a877a663f4f26.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return iaea7c5cfa69bd0ae6b21223f98a830158fc5a2a0297fda11a70a877a663f4f26.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Columns()(*if9ff29afb76bbf1ffef5a126d0ba635e14dcd31de06309e9c7c1a69ad2600738.ColumnsRequestBuilder) {
    return if9ff29afb76bbf1ffef5a126d0ba635e14dcd31de06309e9c7c1a69ad2600738.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*i654b96f70860efab65ce04591400b15de588e186e6e7217fc1ca04b9de80627c.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i654b96f70860efab65ce04591400b15de588e186e6e7217fc1ca04b9de80627c.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListRequestBuilder instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListRequestBuilder) ContentTypes()(*i2ef5930655c892d16078905ec81246034d704193e69907b59b2f14feab329b60.ContentTypesRequestBuilder) {
    return i2ef5930655c892d16078905ec81246034d704193e69907b59b2f14feab329b60.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i6dd3632bd204e05cac0aa8e9c03666774381075a56cb272b02bc729017380185.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i6dd3632bd204e05cac0aa8e9c03666774381075a56cb272b02bc729017380185.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
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
// CreatePatchRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
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
// Delete for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Delete(options *ListRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListRequestBuilder) Drive()(*id28cbea101cb9521a14cf9640c45800a00477bd62395913072f4e81bf1074012.DriveRequestBuilder) {
    return id28cbea101cb9521a14cf9640c45800a00477bd62395913072f4e81bf1074012.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Get(options *ListRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewList() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List), nil
}
func (m *ListRequestBuilder) Items()(*idf9a3059d40c90c42068fa8bf9ac9fd67dedf1182980c80d7a384967f9dff8d5.ItemsRequestBuilder) {
    return idf9a3059d40c90c42068fa8bf9ac9fd67dedf1182980c80d7a384967f9dff8d5.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*i58ec931c55ab0d0111918ef65feb8b5a53f2f824a19e864696540521faf828f4.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return i58ec931c55ab0d0111918ef65feb8b5a53f2f824a19e864696540521faf828f4.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Operations()(*ie43f83ef262031fdb1eb8da2814c4165f5f945383e15adf8746e6a96f3432e23.OperationsRequestBuilder) {
    return ie43f83ef262031fdb1eb8da2814c4165f5f945383e15adf8746e6a96f3432e23.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*i327c97fc6449280089d733ba8dee83d9d74247819100cf6629d3355abfea1166.RichLongRunningOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation_id"] = id
    }
    return i327c97fc6449280089d733ba8dee83d9d74247819100cf6629d3355abfea1166.NewRichLongRunningOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Patch(options *ListRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListRequestBuilder) Subscriptions()(*i1c4ca68087caab22a544a9dc3fb1c29f6076a9264ac284502cbf04a28444aea7.SubscriptionsRequestBuilder) {
    return i1c4ca68087caab22a544a9dc3fb1c29f6076a9264ac284502cbf04a28444aea7.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*i37771808e5dcfef82cc1e8bc33ad871d583b797abbb1222c5dc6198791d77af2.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i37771808e5dcfef82cc1e8bc33ad871d583b797abbb1222c5dc6198791d77af2.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
