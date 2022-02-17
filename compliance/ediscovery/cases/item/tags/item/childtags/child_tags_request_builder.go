package childtags

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3c5316294c65d1516282ef1d734bb37439fe53314f0edfadaef345221f47e252 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/tags/item/childtags/ref"
    ifdcd1631502f3f3c71c9fc98000e43ad8b96769caa5b48a321bfb319bda76469 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/tags/item/childtags/ashierarchy"
)

// ChildTagsRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\tags\{tag-id}\childTags
type ChildTagsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChildTagsRequestBuilderGetOptions options for Get
type ChildTagsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChildTagsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChildTagsRequestBuilderGetQueryParameters returns the tags that are a child of a tag.
type ChildTagsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// AsHierarchy builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\tags\{tag-id}\childTags\microsoft.graph.ediscovery.asHierarchy()
func (m *ChildTagsRequestBuilder) AsHierarchy()(*ifdcd1631502f3f3c71c9fc98000e43ad8b96769caa5b48a321bfb319bda76469.AsHierarchyRequestBuilder) {
    return ifdcd1631502f3f3c71c9fc98000e43ad8b96769caa5b48a321bfb319bda76469.NewAsHierarchyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChildTagsRequestBuilderInternal instantiates a new ChildTagsRequestBuilder and sets the default values.
func NewChildTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChildTagsRequestBuilder) {
    m := &ChildTagsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/tags/{tag_id}/childTags{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChildTagsRequestBuilder instantiates a new ChildTagsRequestBuilder and sets the default values.
func NewChildTagsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChildTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChildTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation returns the tags that are a child of a tag.
func (m *ChildTagsRequestBuilder) CreateGetRequestInformation(options *ChildTagsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get returns the tags that are a child of a tag.
func (m *ChildTagsRequestBuilder) Get(options *ChildTagsRequestBuilderGetOptions)(*ChildTagsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChildTagsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ChildTagsResponse), nil
}
func (m *ChildTagsRequestBuilder) Ref()(*i3c5316294c65d1516282ef1d734bb37439fe53314f0edfadaef345221f47e252.RefRequestBuilder) {
    return i3c5316294c65d1516282ef1d734bb37439fe53314f0edfadaef345221f47e252.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
