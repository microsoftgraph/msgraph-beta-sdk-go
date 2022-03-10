package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/publish"
    i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/copytodefaultcontentlocation"
    i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/associatewithhubsites"
    i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnlinks"
    i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/unpublish"
    i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base"
    i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes"
    ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/ispublished"
    icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columns"
    ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnpositions"
    i47b1c489561ff1669cae57a9dd84a767bbed7f82b77c0cdc7389dc73950c6be7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes/item"
    i49c22f7846eb325805a3ce05072fb08dfdd09048ded457ecb9796b7a866ec2dd "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnpositions/item"
    ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columns/item"
    if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnlinks/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a.AssociateWithHubSitesRequestBuilder) {
    return i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641.BaseRequestBuilder) {
    return i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d.BaseTypesRequestBuilder) {
    return i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i47b1c489561ff1669cae57a9dd84a767bbed7f82b77c0cdc7389dc73950c6be7.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i47b1c489561ff1669cae57a9dd84a767bbed7f82b77c0cdc7389dc73950c6be7.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d.ColumnLinksRequestBuilder) {
    return i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90.ColumnPositionsRequestBuilder) {
    return ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i49c22f7846eb325805a3ce05072fb08dfdd09048ded457ecb9796b7a866ec2dd.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i49c22f7846eb325805a3ce05072fb08dfdd09048ded457ecb9796b7a866ec2dd.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc.ColumnsRequestBuilder) {
    return icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1.CopyToDefaultContentLocationRequestBuilder) {
    return i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for drive
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
// CreatePatchRequestInformation update the navigation property contentTypes in drive
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
// Delete delete navigation property contentTypes for drive
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc.IsPublishedRequestBuilder) {
    return ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in drive
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
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
func (m *ContentTypeItemRequestBuilder) Publish()(*i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3.PublishRequestBuilder) {
    return i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb.UnpublishRequestBuilder) {
    return i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
