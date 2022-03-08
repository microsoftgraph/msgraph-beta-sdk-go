package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0104cc5dce3ca8302e928e80f230d775873a3021040176857a4c258f114daab0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/columns"
    i0e007aa34fb69d97a6a17377c2beb98ad1f83ea0f5f76f4244ab35f40e17aa9e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/columnlinks"
    i613135f6e35dc125aded5209f74f252eab2df3914640eb972c0236d26856d446 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/basetypes"
    i87193cad230c8c8a2c44359a9d21042c2c7eeb8eb08fa74838d9d743f801c3ce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/columnpositions"
    i927ba2c9713209c9fc28b7d9e24e7cc5d67137c8b8d246f4797570f984067918 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/associatewithhubsites"
    ib9afe2cf8200e0cd74ec213f116ca05ebf9091b6d031c71c5e25a8e6fa821c7d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/unpublish"
    ibe7fed46eb1799fbb78d69ddd8bbd5ac1ecfad010c2c7a826acb006a0bfb567e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/base"
    icafe670302445f1b94b16a31e4bdca78fe27c112bca2a0bea1767e9a21da6065 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/ispublished"
    ieced29e4d1896317b15d69be7bf7e922f80e030172d06581d8a46a35df4a9899 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    iffeb97e1a0dcfd6eb8d6c33225a0f5ec18c5a84e861a603d2e283ca78967e0eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/publish"
    i5dfb6cd368a070137133f8078341dd40ac157e4746bdbf82ddf06aee9493892a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/basetypes/item"
    ib6d7cdf52baa00e51444899d63a570a21112f1abee6e0a81d5e7a9740efdc536 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/columnpositions/item"
    ibd2cdfb026621ca2b9f293b70ca5fe16db470895209f5812930fb496772cf8b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/columns/item"
    if2e98d2db82e87e0cebdd19414f2a0f03788d2f8d1df8cd9d0b17abceb16c14e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item/columnlinks/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i927ba2c9713209c9fc28b7d9e24e7cc5d67137c8b8d246f4797570f984067918.AssociateWithHubSitesRequestBuilder) {
    return i927ba2c9713209c9fc28b7d9e24e7cc5d67137c8b8d246f4797570f984067918.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*ibe7fed46eb1799fbb78d69ddd8bbd5ac1ecfad010c2c7a826acb006a0bfb567e.BaseRequestBuilder) {
    return ibe7fed46eb1799fbb78d69ddd8bbd5ac1ecfad010c2c7a826acb006a0bfb567e.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i613135f6e35dc125aded5209f74f252eab2df3914640eb972c0236d26856d446.BaseTypesRequestBuilder) {
    return i613135f6e35dc125aded5209f74f252eab2df3914640eb972c0236d26856d446.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i5dfb6cd368a070137133f8078341dd40ac157e4746bdbf82ddf06aee9493892a.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i5dfb6cd368a070137133f8078341dd40ac157e4746bdbf82ddf06aee9493892a.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i0e007aa34fb69d97a6a17377c2beb98ad1f83ea0f5f76f4244ab35f40e17aa9e.ColumnLinksRequestBuilder) {
    return i0e007aa34fb69d97a6a17377c2beb98ad1f83ea0f5f76f4244ab35f40e17aa9e.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*if2e98d2db82e87e0cebdd19414f2a0f03788d2f8d1df8cd9d0b17abceb16c14e.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return if2e98d2db82e87e0cebdd19414f2a0f03788d2f8d1df8cd9d0b17abceb16c14e.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i87193cad230c8c8a2c44359a9d21042c2c7eeb8eb08fa74838d9d743f801c3ce.ColumnPositionsRequestBuilder) {
    return i87193cad230c8c8a2c44359a9d21042c2c7eeb8eb08fa74838d9d743f801c3ce.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ib6d7cdf52baa00e51444899d63a570a21112f1abee6e0a81d5e7a9740efdc536.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ib6d7cdf52baa00e51444899d63a570a21112f1abee6e0a81d5e7a9740efdc536.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*i0104cc5dce3ca8302e928e80f230d775873a3021040176857a4c258f114daab0.ColumnsRequestBuilder) {
    return i0104cc5dce3ca8302e928e80f230d775873a3021040176857a4c258f114daab0.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*ibd2cdfb026621ca2b9f293b70ca5fe16db470895209f5812930fb496772cf8b8.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ibd2cdfb026621ca2b9f293b70ca5fe16db470895209f5812930fb496772cf8b8.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ieced29e4d1896317b15d69be7bf7e922f80e030172d06581d8a46a35df4a9899.CopyToDefaultContentLocationRequestBuilder) {
    return ieced29e4d1896317b15d69be7bf7e922f80e030172d06581d8a46a35df4a9899.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for me
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
// CreatePatchRequestInformation update the navigation property contentTypes in me
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
// Delete delete navigation property contentTypes for me
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*icafe670302445f1b94b16a31e4bdca78fe27c112bca2a0bea1767e9a21da6065.IsPublishedRequestBuilder) {
    return icafe670302445f1b94b16a31e4bdca78fe27c112bca2a0bea1767e9a21da6065.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in me
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
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
func (m *ContentTypeItemRequestBuilder) Publish()(*iffeb97e1a0dcfd6eb8d6c33225a0f5ec18c5a84e861a603d2e283ca78967e0eb.PublishRequestBuilder) {
    return iffeb97e1a0dcfd6eb8d6c33225a0f5ec18c5a84e861a603d2e283ca78967e0eb.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*ib9afe2cf8200e0cd74ec213f116ca05ebf9091b6d031c71c5e25a8e6fa821c7d.UnpublishRequestBuilder) {
    return ib9afe2cf8200e0cd74ec213f116ca05ebf9091b6d031c71c5e25a8e6fa821c7d.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
