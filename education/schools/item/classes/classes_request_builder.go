package classes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i7ba770d75fbe90866aa3c61aae2e2b5d45a0abe947caad7545aab97ff401b277 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/classes/ref"
    i80b3aae696c74387947d776e9440e66c848887fd9ba21e0472cd458b6bc539be "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/classes/delta"
)

// ClassesRequestBuilder builds and executes requests for operations under \education\schools\{educationSchool-id}\classes
type ClassesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ClassesRequestBuilderGetOptions options for Get
type ClassesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ClassesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ClassesRequestBuilderGetQueryParameters classes taught at the school. Nullable.
type ClassesRequestBuilderGetQueryParameters struct {
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
// NewClassesRequestBuilderInternal instantiates a new ClassesRequestBuilder and sets the default values.
func NewClassesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ClassesRequestBuilder) {
    m := &ClassesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool_id}/classes{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewClassesRequestBuilder instantiates a new ClassesRequestBuilder and sets the default values.
func NewClassesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ClassesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation classes taught at the school. Nullable.
func (m *ClassesRequestBuilder) CreateGetRequestInformation(options *ClassesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delta builds and executes requests for operations under \education\schools\{educationSchool-id}\classes\microsoft.graph.delta()
func (m *ClassesRequestBuilder) Delta()(*i80b3aae696c74387947d776e9440e66c848887fd9ba21e0472cd458b6bc539be.DeltaRequestBuilder) {
    return i80b3aae696c74387947d776e9440e66c848887fd9ba21e0472cd458b6bc539be.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get classes taught at the school. Nullable.
func (m *ClassesRequestBuilder) Get(options *ClassesRequestBuilderGetOptions)(*ClassesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ClassesResponse), nil
}
func (m *ClassesRequestBuilder) Ref()(*i7ba770d75fbe90866aa3c61aae2e2b5d45a0abe947caad7545aab97ff401b277.RefRequestBuilder) {
    return i7ba770d75fbe90866aa3c61aae2e2b5d45a0abe947caad7545aab97ff401b277.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
