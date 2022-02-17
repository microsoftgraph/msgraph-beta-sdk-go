package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0c5a6eb9fea66115bd69699c7bf3c14c8c56d6ddaf4d97c8afd618199d935eae "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/base"
    i2730267b865ade42ffde9c118a145ac5cbe3724b0429d2790c809e94a72e89fa "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/copytodefaultcontentlocation"
    i2da8c35e7e97f2ef273098f53bc6e776a22cb7d02b84bfe9900b6a3040d2cc9f "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/associatewithhubsites"
    i3ba10a504af14f129ece317956d8965ac8085afa2820b041d949da7eb57874d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/ispublished"
    i48be76cbf9aaa82f0511c875016a4ee52f28bd44fe4c1ddd7305aecf7b2a5bf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columnlinks"
    i50b805823a2e1e06133a87ccbd948dbdae542783183fd4393f09b7971d5822ab "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/basetypes"
    i5b9a34041146ae573ad4ad3e977514d4015468ef377df8319306938b4570f0c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columnpositions"
    i760311997236edb768f49330f26e7daf027dafc1e3608a33efd7d83824b07a14 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/publish"
    i91d1e71a788d2b90b0e38b793d7946b2b495724509c5446dfad0545167467897 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/unpublish"
    iad7db789331194a827d671c9f4c9c9db1a2a5947624f5e39904deb155bb43cce "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columns"
    ice1877e581ee39a1cf8b872c81041f0bab8f68c1e2db0210dcaa15a5cb798804 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columns/item"
    icfa060aea76f8132231b7d0f12ab0a407b2c00f0d82270c0adce976e1034f6ff "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columnlinks/item"
)

// ContentTypeRequestBuilder builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}
type ContentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeRequestBuilderDeleteOptions options for Delete
type ContentTypeRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeRequestBuilderGetOptions options for Get
type ContentTypeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeRequestBuilderPatchOptions options for Patch
type ContentTypeRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*i2da8c35e7e97f2ef273098f53bc6e776a22cb7d02b84bfe9900b6a3040d2cc9f.AssociateWithHubSitesRequestBuilder) {
    return i2da8c35e7e97f2ef273098f53bc6e776a22cb7d02b84bfe9900b6a3040d2cc9f.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*i0c5a6eb9fea66115bd69699c7bf3c14c8c56d6ddaf4d97c8afd618199d935eae.BaseRequestBuilder) {
    return i0c5a6eb9fea66115bd69699c7bf3c14c8c56d6ddaf4d97c8afd618199d935eae.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*i50b805823a2e1e06133a87ccbd948dbdae542783183fd4393f09b7971d5822ab.BaseTypesRequestBuilder) {
    return i50b805823a2e1e06133a87ccbd948dbdae542783183fd4393f09b7971d5822ab.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*i48be76cbf9aaa82f0511c875016a4ee52f28bd44fe4c1ddd7305aecf7b2a5bf9.ColumnLinksRequestBuilder) {
    return i48be76cbf9aaa82f0511c875016a4ee52f28bd44fe4c1ddd7305aecf7b2a5bf9.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*icfa060aea76f8132231b7d0f12ab0a407b2c00f0d82270c0adce976e1034f6ff.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return icfa060aea76f8132231b7d0f12ab0a407b2c00f0d82270c0adce976e1034f6ff.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*i5b9a34041146ae573ad4ad3e977514d4015468ef377df8319306938b4570f0c6.ColumnPositionsRequestBuilder) {
    return i5b9a34041146ae573ad4ad3e977514d4015468ef377df8319306938b4570f0c6.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*iad7db789331194a827d671c9f4c9c9db1a2a5947624f5e39904deb155bb43cce.ColumnsRequestBuilder) {
    return iad7db789331194a827d671c9f4c9c9db1a2a5947624f5e39904deb155bb43cce.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*ice1877e581ee39a1cf8b872c81041f0bab8f68c1e2db0210dcaa15a5cb798804.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ice1877e581ee39a1cf8b872c81041f0bab8f68c1e2db0210dcaa15a5cb798804.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeRequestBuilderInternal instantiates a new ContentTypeRequestBuilder and sets the default values.
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeRequestBuilder instantiates a new ContentTypeRequestBuilder and sets the default values.
func NewContentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*i2730267b865ade42ffde9c118a145ac5cbe3724b0429d2790c809e94a72e89fa.CopyToDefaultContentLocationRequestBuilder) {
    return i2730267b865ade42ffde9c118a145ac5cbe3724b0429d2790c809e94a72e89fa.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(options *ContentTypeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) CreatePatchRequestInformation(options *ContentTypeRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Delete(options *ContentTypeRequestBuilderDeleteOptions)(error) {
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
// Get the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Get(options *ContentTypeRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContentType() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType), nil
}
// IsPublished builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}\microsoft.graph.isPublished()
func (m *ContentTypeRequestBuilder) IsPublished()(*i3ba10a504af14f129ece317956d8965ac8085afa2820b041d949da7eb57874d0.IsPublishedRequestBuilder) {
    return i3ba10a504af14f129ece317956d8965ac8085afa2820b041d949da7eb57874d0.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Patch(options *ContentTypeRequestBuilderPatchOptions)(error) {
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
func (m *ContentTypeRequestBuilder) Publish()(*i760311997236edb768f49330f26e7daf027dafc1e3608a33efd7d83824b07a14.PublishRequestBuilder) {
    return i760311997236edb768f49330f26e7daf027dafc1e3608a33efd7d83824b07a14.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*i91d1e71a788d2b90b0e38b793d7946b2b495724509c5446dfad0545167467897.UnpublishRequestBuilder) {
    return i91d1e71a788d2b90b0e38b793d7946b2b495724509c5446dfad0545167467897.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
