package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6cd0dd643a8853c5e36c664f3266e8bc15e43dfd7da99ab8ec908ba5a2e2846c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/unpublish"
    i71a00748ec1b3f1916f08f1378ac75be568e315e8c45671703ce0a998a523535 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/basetypes"
    i8b97072a604d4b89bbdcd18ca9b769ea9b5f892512bf132888c5ceb0d874996d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columnpositions"
    i8f46b19091d8fc007c8204da63e2fecf63e3e12658f88e306d5afdce4ccf783c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/associatewithhubsites"
    ia5eb9dbcf222805d1a39da295590fdc1cdff5f38e71fbe02b0cbde0ac79badf8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/ispublished"
    ic33bb574f58874bf5559129e9d7f446f758bfc0baa4cc59067b93d1138da4756 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columnlinks"
    ic77a6b884e9bdc94d36c4085a09bbcf0676b01c87b3bc44e07282c5c2f0de7d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/publish"
    id33ae53bb415f03b269f46aa3901afb4aa844cd415a3baca9ba770f3f695c744 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    ida787a0eecc9cbdceedc5fa328132abca55da74c1475c3cd90d33f515a955a58 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/base"
    idb904522da356683cdbc2ac70d390d7692262d927ea7a3752d44744d20e4c435 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columns"
    i5c3b82c0971f80566e2f27fb191b318b1e3605714d556179aab3bf7a231cc5e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columnlinks/item"
    i881e5cfd2288828ecac5b4440b76a4eba0683212edb64e32af0744d8d4e2c0e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columns/item"
    iab2728a4c36bbe5d4b7f8a74e0d7c74f5bb30814b9618b8c70571b67c6bd0c75 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/basetypes/item"
    ib8b83821cbc0899cc9656c33cc49eb4adc13be44554b10215afcaaf8eae6327a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columnpositions/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeItemRequestBuilderPatchOptions options for Patch
type ContentTypeItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i8f46b19091d8fc007c8204da63e2fecf63e3e12658f88e306d5afdce4ccf783c.AssociateWithHubSitesRequestBuilder) {
    return i8f46b19091d8fc007c8204da63e2fecf63e3e12658f88e306d5afdce4ccf783c.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*ida787a0eecc9cbdceedc5fa328132abca55da74c1475c3cd90d33f515a955a58.BaseRequestBuilder) {
    return ida787a0eecc9cbdceedc5fa328132abca55da74c1475c3cd90d33f515a955a58.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i71a00748ec1b3f1916f08f1378ac75be568e315e8c45671703ce0a998a523535.BaseTypesRequestBuilder) {
    return i71a00748ec1b3f1916f08f1378ac75be568e315e8c45671703ce0a998a523535.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*iab2728a4c36bbe5d4b7f8a74e0d7c74f5bb30814b9618b8c70571b67c6bd0c75.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return iab2728a4c36bbe5d4b7f8a74e0d7c74f5bb30814b9618b8c70571b67c6bd0c75.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*ic33bb574f58874bf5559129e9d7f446f758bfc0baa4cc59067b93d1138da4756.ColumnLinksRequestBuilder) {
    return ic33bb574f58874bf5559129e9d7f446f758bfc0baa4cc59067b93d1138da4756.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i5c3b82c0971f80566e2f27fb191b318b1e3605714d556179aab3bf7a231cc5e3.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i5c3b82c0971f80566e2f27fb191b318b1e3605714d556179aab3bf7a231cc5e3.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i8b97072a604d4b89bbdcd18ca9b769ea9b5f892512bf132888c5ceb0d874996d.ColumnPositionsRequestBuilder) {
    return i8b97072a604d4b89bbdcd18ca9b769ea9b5f892512bf132888c5ceb0d874996d.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ib8b83821cbc0899cc9656c33cc49eb4adc13be44554b10215afcaaf8eae6327a.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ib8b83821cbc0899cc9656c33cc49eb4adc13be44554b10215afcaaf8eae6327a.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*idb904522da356683cdbc2ac70d390d7692262d927ea7a3752d44744d20e4c435.ColumnsRequestBuilder) {
    return idb904522da356683cdbc2ac70d390d7692262d927ea7a3752d44744d20e4c435.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i881e5cfd2288828ecac5b4440b76a4eba0683212edb64e32af0744d8d4e2c0e1.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i881e5cfd2288828ecac5b4440b76a4eba0683212edb64e32af0744d8d4e2c0e1.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*id33ae53bb415f03b269f46aa3901afb4aa844cd415a3baca9ba770f3f695c744.CopyToDefaultContentLocationRequestBuilder) {
    return id33ae53bb415f03b269f46aa3901afb4aa844cd415a3baca9ba770f3f695c744.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*ia5eb9dbcf222805d1a39da295590fdc1cdff5f38e71fbe02b0cbde0ac79badf8.IsPublishedRequestBuilder) {
    return ia5eb9dbcf222805d1a39da295590fdc1cdff5f38e71fbe02b0cbde0ac79badf8.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
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
func (m *ContentTypeItemRequestBuilder) Publish()(*ic77a6b884e9bdc94d36c4085a09bbcf0676b01c87b3bc44e07282c5c2f0de7d1.PublishRequestBuilder) {
    return ic77a6b884e9bdc94d36c4085a09bbcf0676b01c87b3bc44e07282c5c2f0de7d1.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i6cd0dd643a8853c5e36c664f3266e8bc15e43dfd7da99ab8ec908ba5a2e2846c.UnpublishRequestBuilder) {
    return i6cd0dd643a8853c5e36c664f3266e8bc15e43dfd7da99ab8ec908ba5a2e2846c.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
