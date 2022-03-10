package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04c203075f0399dc7cbc1f50b6f38b707b07be14a2fb7d44bb011552ceb712d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/externalcolumns"
    i0d031e5a8bc99797a387d266d9cbab194ba77e02db5f5c44866b22aa56851245 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/drives"
    i163b1106365b137fdb131300f3c0f6b8e2327678b800c3e078233d7d74db385a "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes"
    i1991c7a4b0a20509c544953013d93e8d47ab98806f12da33b247cec0746146b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/drive"
    i37aa3e131565d412164dd7c8aef4debebaa21612c62cd325e8a708b589e601da "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/getapplicablecontenttypesforlistwithlistid"
    i41d40ece57596aff17bbd4e071da917e9696de96733cdb577567079fdc4fd02b "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/columns"
    i46cbb195f6e4660d7b45749753a475a645bf6784bfc9f9adc1d79eccfc9a7b40 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i470ba8ba4f5c4ab0e450363d2c299e7fcf642547542575d37e8ab1b0a273d750 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/pages"
    i513536400845064c1293d5ecd0000bf384a646a6457e9128f86c1d6d415af282 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/operations"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i751a7aab9250bdc979b3359c5e5d3056a650dbc575f271e01cb91720c05b3c48 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote"
    i7991f174b586f44e549c3d8dce6877288c6b0fbc7ba8f49a6e6e98bd4137217d "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/items"
    i7f031c5d4daa82a55f4b2c1717082c661c29d6eee85712b7eaf54f2616b0d9a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists"
    i808417908586dcac70c7b70c188cca33c7915492b55710e4dc2e94823e3765cc "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/getbypathwithpath"
    i827abacdcee64b9285bd46dfe43361a6c33ec040f3cd6bc2bc4729b0ab6b2aed "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/sites"
    i842757c0e3ad7774d9626c19093a7db6bb052d1a0ac574cbe23bc486c5cb9c24 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/analytics"
    ia0f6da1c0dc630de425d1e6b22fa3f2a8cec2b725423ef3c6f168587074ebe5e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore"
    ic4329a5ed1e54000b3f53e26c12bea37b9b3002a2e267004ed8d9bfc75c3f9c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/permissions"
    i0e66ddd1911d4faa407029ec9d0d44af8f3f0dffd6d4ff8e37227f18c6dcb9e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/sites/item"
    i13de7ce6001c6b4559e0a6f9bb09511cb1f9e883a741518534e6b03fba66897a "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/pages/item"
    i3721ddd37978ed5887e741cd4ec530f247da77968194df669e6b5cf53370bed5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item"
    i414042b7d435cf3a5ea7119418d530da6e7fbf15fcc40be074e1d7c9a6996fd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/columns/item"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i7c78829e5ecb7ce49dec3e15a0823e233af9edec6a7211cb55cf0d8ee73c2a68 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/drives/item"
    ic0a30c9fcd69c709d292863ba51b88fe2799046dcce950f9d1bc1a372381eb06 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item"
    ic6fa3aac02c15ce02c1307596eb029012f7b5635177e98447f51c9743b7179bf "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/externalcolumns/item"
    icf0323a037a02f8c72110955832cfede83bf2251b506c77f844efff04e3320d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/items/item"
    ied00bf4ac26d58488e9ee1a357d4bbdf2553de648d04880076d6295dcc23e36f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/operations/item"
    if030207cd960e3ca95a11c646417c461af3fa2b3617d9b9b2af40510bf1e4eb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/permissions/item"
)

// SiteItemRequestBuilder provides operations to manage the collection of site entities.
type SiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SiteItemRequestBuilderDeleteOptions options for Delete
type SiteItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SiteItemRequestBuilderGetOptions options for Get
type SiteItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SiteItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SiteItemRequestBuilderGetQueryParameters get entity from sites by key
type SiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SiteItemRequestBuilderPatchOptions options for Patch
type SiteItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Siteable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SiteItemRequestBuilder) Analytics()(*i842757c0e3ad7774d9626c19093a7db6bb052d1a0ac574cbe23bc486c5cb9c24.AnalyticsRequestBuilder) {
    return i842757c0e3ad7774d9626c19093a7db6bb052d1a0ac574cbe23bc486c5cb9c24.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) Columns()(*i41d40ece57596aff17bbd4e071da917e9696de96733cdb577567079fdc4fd02b.ColumnsRequestBuilder) {
    return i41d40ece57596aff17bbd4e071da917e9696de96733cdb577567079fdc4fd02b.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.columns.item collection
func (m *SiteItemRequestBuilder) ColumnsById(id string)(*i414042b7d435cf3a5ea7119418d530da6e7fbf15fcc40be074e1d7c9a6996fd1.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i414042b7d435cf3a5ea7119418d530da6e7fbf15fcc40be074e1d7c9a6996fd1.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SiteItemRequestBuilder) {
    m := &SiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SiteItemRequestBuilder) ContentTypes()(*i163b1106365b137fdb131300f3c0f6b8e2327678b800c3e078233d7d74db385a.ContentTypesRequestBuilder) {
    return i163b1106365b137fdb131300f3c0f6b8e2327678b800c3e078233d7d74db385a.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item collection
func (m *SiteItemRequestBuilder) ContentTypesById(id string)(*i3721ddd37978ed5887e741cd4ec530f247da77968194df669e6b5cf53370bed5.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i3721ddd37978ed5887e741cd4ec530f247da77968194df669e6b5cf53370bed5.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete entity from sites
func (m *SiteItemRequestBuilder) CreateDeleteRequestInformation(options *SiteItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from sites by key
func (m *SiteItemRequestBuilder) CreateGetRequestInformation(options *SiteItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in sites
func (m *SiteItemRequestBuilder) CreatePatchRequestInformation(options *SiteItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from sites
func (m *SiteItemRequestBuilder) Delete(options *SiteItemRequestBuilderDeleteOptions)(error) {
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
func (m *SiteItemRequestBuilder) Drive()(*i1991c7a4b0a20509c544953013d93e8d47ab98806f12da33b247cec0746146b2.DriveRequestBuilder) {
    return i1991c7a4b0a20509c544953013d93e8d47ab98806f12da33b247cec0746146b2.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) Drives()(*i0d031e5a8bc99797a387d266d9cbab194ba77e02db5f5c44866b22aa56851245.DrivesRequestBuilder) {
    return i0d031e5a8bc99797a387d266d9cbab194ba77e02db5f5c44866b22aa56851245.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.drives.item collection
func (m *SiteItemRequestBuilder) DrivesById(id string)(*i7c78829e5ecb7ce49dec3e15a0823e233af9edec6a7211cb55cf0d8ee73c2a68.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive_id"] = id
    }
    return i7c78829e5ecb7ce49dec3e15a0823e233af9edec6a7211cb55cf0d8ee73c2a68.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) ExternalColumns()(*i04c203075f0399dc7cbc1f50b6f38b707b07be14a2fb7d44bb011552ceb712d6.ExternalColumnsRequestBuilder) {
    return i04c203075f0399dc7cbc1f50b6f38b707b07be14a2fb7d44bb011552ceb712d6.NewExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.externalColumns.item collection
func (m *SiteItemRequestBuilder) ExternalColumnsById(id string)(*ic6fa3aac02c15ce02c1307596eb029012f7b5635177e98447f51c9743b7179bf.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ic6fa3aac02c15ce02c1307596eb029012f7b5635177e98447f51c9743b7179bf.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from sites by key
func (m *SiteItemRequestBuilder) Get(options *SiteItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Siteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateSiteFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Siteable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *SiteItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i46cbb195f6e4660d7b45749753a475a645bf6784bfc9f9adc1d79eccfc9a7b40.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i46cbb195f6e4660d7b45749753a475a645bf6784bfc9f9adc1d79eccfc9a7b40.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// GetApplicableContentTypesForListWithListId provides operations to call the getApplicableContentTypesForList method.
func (m *SiteItemRequestBuilder) GetApplicableContentTypesForListWithListId(listId *string)(*i37aa3e131565d412164dd7c8aef4debebaa21612c62cd325e8a708b589e601da.GetApplicableContentTypesForListWithListIdRequestBuilder) {
    return i37aa3e131565d412164dd7c8aef4debebaa21612c62cd325e8a708b589e601da.NewGetApplicableContentTypesForListWithListIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, listId);
}
// GetByPathWithPath provides operations to call the getByPath method.
func (m *SiteItemRequestBuilder) GetByPathWithPath(path *string)(*i808417908586dcac70c7b70c188cca33c7915492b55710e4dc2e94823e3765cc.GetByPathWithPathRequestBuilder) {
    return i808417908586dcac70c7b70c188cca33c7915492b55710e4dc2e94823e3765cc.NewGetByPathWithPathRequestBuilderInternal(m.pathParameters, m.requestAdapter, path);
}
func (m *SiteItemRequestBuilder) Items()(*i7991f174b586f44e549c3d8dce6877288c6b0fbc7ba8f49a6e6e98bd4137217d.ItemsRequestBuilder) {
    return i7991f174b586f44e549c3d8dce6877288c6b0fbc7ba8f49a6e6e98bd4137217d.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.items.item collection
func (m *SiteItemRequestBuilder) ItemsById(id string)(*icf0323a037a02f8c72110955832cfede83bf2251b506c77f844efff04e3320d5.BaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem_id"] = id
    }
    return icf0323a037a02f8c72110955832cfede83bf2251b506c77f844efff04e3320d5.NewBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) Lists()(*i7f031c5d4daa82a55f4b2c1717082c661c29d6eee85712b7eaf54f2616b0d9a0.ListsRequestBuilder) {
    return i7f031c5d4daa82a55f4b2c1717082c661c29d6eee85712b7eaf54f2616b0d9a0.NewListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item collection
func (m *SiteItemRequestBuilder) ListsById(id string)(*ic0a30c9fcd69c709d292863ba51b88fe2799046dcce950f9d1bc1a372381eb06.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list_id"] = id
    }
    return ic0a30c9fcd69c709d292863ba51b88fe2799046dcce950f9d1bc1a372381eb06.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) Onenote()(*i751a7aab9250bdc979b3359c5e5d3056a650dbc575f271e01cb91720c05b3c48.OnenoteRequestBuilder) {
    return i751a7aab9250bdc979b3359c5e5d3056a650dbc575f271e01cb91720c05b3c48.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) Operations()(*i513536400845064c1293d5ecd0000bf384a646a6457e9128f86c1d6d415af282.OperationsRequestBuilder) {
    return i513536400845064c1293d5ecd0000bf384a646a6457e9128f86c1d6d415af282.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.operations.item collection
func (m *SiteItemRequestBuilder) OperationsById(id string)(*ied00bf4ac26d58488e9ee1a357d4bbdf2553de648d04880076d6295dcc23e36f.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation_id"] = id
    }
    return ied00bf4ac26d58488e9ee1a357d4bbdf2553de648d04880076d6295dcc23e36f.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) Pages()(*i470ba8ba4f5c4ab0e450363d2c299e7fcf642547542575d37e8ab1b0a273d750.PagesRequestBuilder) {
    return i470ba8ba4f5c4ab0e450363d2c299e7fcf642547542575d37e8ab1b0a273d750.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.pages.item collection
func (m *SiteItemRequestBuilder) PagesById(id string)(*i13de7ce6001c6b4559e0a6f9bb09511cb1f9e883a741518534e6b03fba66897a.SitePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sitePage_id"] = id
    }
    return i13de7ce6001c6b4559e0a6f9bb09511cb1f9e883a741518534e6b03fba66897a.NewSitePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in sites
func (m *SiteItemRequestBuilder) Patch(options *SiteItemRequestBuilderPatchOptions)(error) {
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
func (m *SiteItemRequestBuilder) Permissions()(*ic4329a5ed1e54000b3f53e26c12bea37b9b3002a2e267004ed8d9bfc75c3f9c8.PermissionsRequestBuilder) {
    return ic4329a5ed1e54000b3f53e26c12bea37b9b3002a2e267004ed8d9bfc75c3f9c8.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.permissions.item collection
func (m *SiteItemRequestBuilder) PermissionsById(id string)(*if030207cd960e3ca95a11c646417c461af3fa2b3617d9b9b2af40510bf1e4eb9.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return if030207cd960e3ca95a11c646417c461af3fa2b3617d9b9b2af40510bf1e4eb9.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) Sites()(*i827abacdcee64b9285bd46dfe43361a6c33ec040f3cd6bc2bc4729b0ab6b2aed.SitesRequestBuilder) {
    return i827abacdcee64b9285bd46dfe43361a6c33ec040f3cd6bc2bc4729b0ab6b2aed.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.sites.item collection
func (m *SiteItemRequestBuilder) SitesById(id string)(*i0e66ddd1911d4faa407029ec9d0d44af8f3f0dffd6d4ff8e37227f18c6dcb9e5.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site_id1"] = id
    }
    return i0e66ddd1911d4faa407029ec9d0d44af8f3f0dffd6d4ff8e37227f18c6dcb9e5.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SiteItemRequestBuilder) TermStore()(*ia0f6da1c0dc630de425d1e6b22fa3f2a8cec2b725423ef3c6f168587074ebe5e.TermStoreRequestBuilder) {
    return ia0f6da1c0dc630de425d1e6b22fa3f2a8cec2b725423ef3c6f168587074ebe5e.NewTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
