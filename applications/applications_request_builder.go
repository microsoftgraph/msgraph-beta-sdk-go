package applications

import (
    i15d96fb11649ee1a0fc985a251372398ff2507b395416a72de9596e35ddf1812 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/getbyids"
    i19b6cc6577bf7cf252227ad0abef9dc37b552128c2c64907d4ba012be3e583ab "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/delta"
    i7845412e91bb6607a75c0cb9f9ed7cd86d8a84db7fec8ac1aa8c4dadc4902c92 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/validateproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    idec80e72882c64f96a88602c8bb00ec1b82bd733af8fa4accef9e3a5ecbe74d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/getuserownedobjects"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ApplicationsRequestBuilder builds and executes requests for operations under \applications
type ApplicationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ApplicationsRequestBuilderGetOptions options for Get
type ApplicationsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ApplicationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ApplicationsRequestBuilderGetQueryParameters get entities from applications
type ApplicationsRequestBuilderGetQueryParameters struct {
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
// ApplicationsRequestBuilderPostOptions options for Post
type ApplicationsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Application;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewApplicationsRequestBuilderInternal instantiates a new ApplicationsRequestBuilder and sets the default values.
func NewApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplicationsRequestBuilder) {
    m := &ApplicationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationsRequestBuilder instantiates a new ApplicationsRequestBuilder and sets the default values.
func NewApplicationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplicationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get entities from applications
func (m *ApplicationsRequestBuilder) CreateGetRequestInformation(options *ApplicationsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to applications
func (m *ApplicationsRequestBuilder) CreatePostRequestInformation(options *ApplicationsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Delta builds and executes requests for operations under \applications\microsoft.graph.delta()
func (m *ApplicationsRequestBuilder) Delta()(*i19b6cc6577bf7cf252227ad0abef9dc37b552128c2c64907d4ba012be3e583ab.DeltaRequestBuilder) {
    return i19b6cc6577bf7cf252227ad0abef9dc37b552128c2c64907d4ba012be3e583ab.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from applications
func (m *ApplicationsRequestBuilder) Get(options *ApplicationsRequestBuilderGetOptions)(*ApplicationsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApplicationsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ApplicationsResponse), nil
}
func (m *ApplicationsRequestBuilder) GetByIds()(*i15d96fb11649ee1a0fc985a251372398ff2507b395416a72de9596e35ddf1812.GetByIdsRequestBuilder) {
    return i15d96fb11649ee1a0fc985a251372398ff2507b395416a72de9596e35ddf1812.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationsRequestBuilder) GetUserOwnedObjects()(*idec80e72882c64f96a88602c8bb00ec1b82bd733af8fa4accef9e3a5ecbe74d1.GetUserOwnedObjectsRequestBuilder) {
    return idec80e72882c64f96a88602c8bb00ec1b82bd733af8fa4accef9e3a5ecbe74d1.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to applications
func (m *ApplicationsRequestBuilder) Post(options *ApplicationsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Application, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewApplication() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Application), nil
}
func (m *ApplicationsRequestBuilder) ValidateProperties()(*i7845412e91bb6607a75c0cb9f9ed7cd86d8a84db7fec8ac1aa8c4dadc4902c92.ValidatePropertiesRequestBuilder) {
    return i7845412e91bb6607a75c0cb9f9ed7cd86d8a84db7fec8ac1aa8c4dadc4902c92.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
