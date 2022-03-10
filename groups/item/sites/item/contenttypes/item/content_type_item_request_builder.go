package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i01d884c5b597bb4bbb21ef6ddb793472b4e5cfac11eee17fa06a7c93e7af3961 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/unpublish"
    i0dbe0e79f8d2132817cfd4fd9c0ba48a22c1436a654ba3b515c9d78ea04153bf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columnpositions"
    i15b535ee8c2ffa35c2be78c2e83049b8f87132a9d6c9095165aece4073c7e68b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/basetypes"
    i5072c1863bff5f51d828c097f140a9067fffcdd636bce59d31fd3cd41b86a7f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/base"
    i53102ef79a47ebbae7839e80aaa6802806794d37343b777c078fcbaad0a195c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columns"
    iaf87d69143fb1b049d1e8643342885a68173ad08feea74278b1460b57f935ec4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columnlinks"
    iba673aaa67a5f7556ea399c86c8c32b04c721c33b18e487c30de92eea73d517b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/publish"
    ibdbf84d3479e495735427bb0dcdf1fe8f04d07858d5d6a3a0f34ef4cd0ea5c17 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/copytodefaultcontentlocation"
    ic426e586ad0e202686674873cbdea7a8018e5b74cf6d66de10ed9573184cdd97 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/associatewithhubsites"
    ifa15559168b540626f112cbef09640515aab7cf666d590c01974f1056022fd94 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/ispublished"
    i0c5d02e0200c8a34f0952dda3d6bed322ad1bf731eaa3dc3c37e94f37cc44a12 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columnpositions/item"
    i70ae4aef64d211900e5691042f2ac667a4a38c969183ce5df5980865df763382 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columnlinks/item"
    i71496ba62ae8bd6cc9c2a77d0a541c08ebf273eea9ecf422f1040c706f9de5ad "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columns/item"
    if47acb49fe28f5ed2e47bb6c2bd5cb847ac5f03b74ad8571b28a23def60220c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/basetypes/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.site entity.
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
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types defined for this site.
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*ic426e586ad0e202686674873cbdea7a8018e5b74cf6d66de10ed9573184cdd97.AssociateWithHubSitesRequestBuilder) {
    return ic426e586ad0e202686674873cbdea7a8018e5b74cf6d66de10ed9573184cdd97.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i5072c1863bff5f51d828c097f140a9067fffcdd636bce59d31fd3cd41b86a7f7.BaseRequestBuilder) {
    return i5072c1863bff5f51d828c097f140a9067fffcdd636bce59d31fd3cd41b86a7f7.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i15b535ee8c2ffa35c2be78c2e83049b8f87132a9d6c9095165aece4073c7e68b.BaseTypesRequestBuilder) {
    return i15b535ee8c2ffa35c2be78c2e83049b8f87132a9d6c9095165aece4073c7e68b.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*if47acb49fe28f5ed2e47bb6c2bd5cb847ac5f03b74ad8571b28a23def60220c4.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return if47acb49fe28f5ed2e47bb6c2bd5cb847ac5f03b74ad8571b28a23def60220c4.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*iaf87d69143fb1b049d1e8643342885a68173ad08feea74278b1460b57f935ec4.ColumnLinksRequestBuilder) {
    return iaf87d69143fb1b049d1e8643342885a68173ad08feea74278b1460b57f935ec4.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i70ae4aef64d211900e5691042f2ac667a4a38c969183ce5df5980865df763382.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i70ae4aef64d211900e5691042f2ac667a4a38c969183ce5df5980865df763382.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i0dbe0e79f8d2132817cfd4fd9c0ba48a22c1436a654ba3b515c9d78ea04153bf.ColumnPositionsRequestBuilder) {
    return i0dbe0e79f8d2132817cfd4fd9c0ba48a22c1436a654ba3b515c9d78ea04153bf.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i0c5d02e0200c8a34f0952dda3d6bed322ad1bf731eaa3dc3c37e94f37cc44a12.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i0c5d02e0200c8a34f0952dda3d6bed322ad1bf731eaa3dc3c37e94f37cc44a12.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*i53102ef79a47ebbae7839e80aaa6802806794d37343b777c078fcbaad0a195c1.ColumnsRequestBuilder) {
    return i53102ef79a47ebbae7839e80aaa6802806794d37343b777c078fcbaad0a195c1.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i71496ba62ae8bd6cc9c2a77d0a541c08ebf273eea9ecf422f1040c706f9de5ad.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i71496ba62ae8bd6cc9c2a77d0a541c08ebf273eea9ecf422f1040c706f9de5ad.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ibdbf84d3479e495735427bb0dcdf1fe8f04d07858d5d6a3a0f34ef4cd0ea5c17.CopyToDefaultContentLocationRequestBuilder) {
    return ibdbf84d3479e495735427bb0dcdf1fe8f04d07858d5d6a3a0f34ef4cd0ea5c17.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreateGetRequestInformation the collection of content types defined for this site.
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
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types defined for this site.
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*ifa15559168b540626f112cbef09640515aab7cf666d590c01974f1056022fd94.IsPublishedRequestBuilder) {
    return ifa15559168b540626f112cbef09640515aab7cf666d590c01974f1056022fd94.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in groups
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
func (m *ContentTypeItemRequestBuilder) Publish()(*iba673aaa67a5f7556ea399c86c8c32b04c721c33b18e487c30de92eea73d517b.PublishRequestBuilder) {
    return iba673aaa67a5f7556ea399c86c8c32b04c721c33b18e487c30de92eea73d517b.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i01d884c5b597bb4bbb21ef6ddb793472b4e5cfac11eee17fa06a7c93e7af3961.UnpublishRequestBuilder) {
    return i01d884c5b597bb4bbb21ef6ddb793472b4e5cfac11eee17fa06a7c93e7af3961.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
