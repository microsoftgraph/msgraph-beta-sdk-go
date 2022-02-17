package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0823bc33f3568f83828166ea8ced7949b7cb17282576f6ea005ae3af7e33db1e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/base/publish"
    i33978688bfdd69a91ea02ac3c4b8cb660024258077a29b0f5893d0f0ef5a4fbd "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/base/unpublish"
    i7897dedc8aeb8ade5141636a064e24f3d9e5ae8f62ae612ac492d3d69552b01e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/base/copytodefaultcontentlocation"
    icd203716875673a7a17a810e85107d7bc0a241ff72b1844c18f45c8f4fd8733e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/base/ref"
    id20077693ba13ea0db568e9a8b1822f6ff2ccd2d5dd590ffd8dc544ef5b2f7fe "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/base/associatewithhubsites"
    ieedf3e62b3df475dfaae13e358b16d3cc4b8bbd97a73cdf24cb1a2bd6e29525f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/base/ispublished"
)

// BaseRequestBuilder builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\contentTypes\{contentType-id}\base
type BaseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BaseRequestBuilderGetOptions options for Get
type BaseRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BaseRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseRequestBuilderGetQueryParameters parent contentType from which this content type is derived.
type BaseRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *BaseRequestBuilder) AssociateWithHubSites()(*id20077693ba13ea0db568e9a8b1822f6ff2ccd2d5dd590ffd8dc544ef5b2f7fe.AssociateWithHubSitesRequestBuilder) {
    return id20077693ba13ea0db568e9a8b1822f6ff2ccd2d5dd590ffd8dc544ef5b2f7fe.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseRequestBuilderInternal instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}/base{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseRequestBuilder instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*i7897dedc8aeb8ade5141636a064e24f3d9e5ae8f62ae612ac492d3d69552b01e.CopyToDefaultContentLocationRequestBuilder) {
    return i7897dedc8aeb8ade5141636a064e24f3d9e5ae8f62ae612ac492d3d69552b01e.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation parent contentType from which this content type is derived.
func (m *BaseRequestBuilder) CreateGetRequestInformation(options *BaseRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get parent contentType from which this content type is derived.
func (m *BaseRequestBuilder) Get(options *BaseRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, error) {
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
// IsPublished builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*ieedf3e62b3df475dfaae13e358b16d3cc4b8bbd97a73cdf24cb1a2bd6e29525f.IsPublishedRequestBuilder) {
    return ieedf3e62b3df475dfaae13e358b16d3cc4b8bbd97a73cdf24cb1a2bd6e29525f.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*i0823bc33f3568f83828166ea8ced7949b7cb17282576f6ea005ae3af7e33db1e.PublishRequestBuilder) {
    return i0823bc33f3568f83828166ea8ced7949b7cb17282576f6ea005ae3af7e33db1e.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*icd203716875673a7a17a810e85107d7bc0a241ff72b1844c18f45c8f4fd8733e.RefRequestBuilder) {
    return icd203716875673a7a17a810e85107d7bc0a241ff72b1844c18f45c8f4fd8733e.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*i33978688bfdd69a91ea02ac3c4b8cb660024258077a29b0f5893d0f0ef5a4fbd.UnpublishRequestBuilder) {
    return i33978688bfdd69a91ea02ac3c4b8cb660024258077a29b0f5893d0f0ef5a4fbd.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
