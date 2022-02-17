package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1805cd70634bc7e8d887e29d88417f987d5c7a4621c3b7c6766152f522306051 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base/copytodefaultcontentlocation"
    i25a02cca8ea3cf1e8f58173d57ca4bac3df04843add8bb130a61e038258e71fe "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base/publish"
    i36e927991c051b16eb9cf03e5cd437775ab1b66963e5fc1a646ae3e7963e9116 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base/ref"
    i7a19def03e13141505cc556c2a35ef3aa6adf81925bd7f41c77c24e04035046a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base/associatewithhubsites"
    i8647a40ee6012651b3137b5fac10fdbf5eec0b653bf4c25ffef1c524b6ced9b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base/unpublish"
    i9a6959f01ec77007244c7d6c6b76a7beddcda5b60a9a3677d97e1ba5f870b17f "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base/ispublished"
)

// BaseRequestBuilder builds and executes requests for operations under \drive\list\contentTypes\{contentType-id}\base
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
func (m *BaseRequestBuilder) AssociateWithHubSites()(*i7a19def03e13141505cc556c2a35ef3aa6adf81925bd7f41c77c24e04035046a.AssociateWithHubSitesRequestBuilder) {
    return i7a19def03e13141505cc556c2a35ef3aa6adf81925bd7f41c77c24e04035046a.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseRequestBuilderInternal instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/contentTypes/{contentType_id}/base{?select,expand}";
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
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*i1805cd70634bc7e8d887e29d88417f987d5c7a4621c3b7c6766152f522306051.CopyToDefaultContentLocationRequestBuilder) {
    return i1805cd70634bc7e8d887e29d88417f987d5c7a4621c3b7c6766152f522306051.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// IsPublished builds and executes requests for operations under \drive\list\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*i9a6959f01ec77007244c7d6c6b76a7beddcda5b60a9a3677d97e1ba5f870b17f.IsPublishedRequestBuilder) {
    return i9a6959f01ec77007244c7d6c6b76a7beddcda5b60a9a3677d97e1ba5f870b17f.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*i25a02cca8ea3cf1e8f58173d57ca4bac3df04843add8bb130a61e038258e71fe.PublishRequestBuilder) {
    return i25a02cca8ea3cf1e8f58173d57ca4bac3df04843add8bb130a61e038258e71fe.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*i36e927991c051b16eb9cf03e5cd437775ab1b66963e5fc1a646ae3e7963e9116.RefRequestBuilder) {
    return i36e927991c051b16eb9cf03e5cd437775ab1b66963e5fc1a646ae3e7963e9116.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*i8647a40ee6012651b3137b5fac10fdbf5eec0b653bf4c25ffef1c524b6ced9b8.UnpublishRequestBuilder) {
    return i8647a40ee6012651b3137b5fac10fdbf5eec0b653bf4c25ffef1c524b6ced9b8.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
