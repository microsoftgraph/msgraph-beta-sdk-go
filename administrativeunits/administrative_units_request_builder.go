package administrativeunits

import (
    i277b79a3b313ac8255d78ae90b5e9c683fe411d78cfeb5e1e83d31d38c270e05 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/getbyids"
    ia5a3db4673512570eb5641390e08a082576bca30d7028287114d44432947b385 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/getuserownedobjects"
    ic5b5428ae35afae2a32cefcf30c64f6ab1fbe0e8909656a8a35a1a32357418eb "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/validateproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if16d334c721262585664df32b0567644d1ab8a2a5c2829d1f520c68c93e0c424 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/delta"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \administrativeUnits
type AdministrativeUnitsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type AdministrativeUnitsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AdministrativeUnitsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entities from administrativeUnits
type AdministrativeUnitsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Options for Post
type AdministrativeUnitsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new AdministrativeUnitsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAdministrativeUnitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitsRequestBuilder) {
    m := &AdministrativeUnitsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/administrativeUnits{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AdministrativeUnitsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAdministrativeUnitsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get entities from administrativeUnits
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitsRequestBuilder) CreateGetRequestInformation(options *AdministrativeUnitsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Add new entity to administrativeUnits
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitsRequestBuilder) CreatePostRequestInformation(options *AdministrativeUnitsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Builds and executes requests for operations under \administrativeUnits\microsoft.graph.delta()
func (m *AdministrativeUnitsRequestBuilder) Delta()(*if16d334c721262585664df32b0567644d1ab8a2a5c2829d1f520c68c93e0c424.DeltaRequestBuilder) {
    return if16d334c721262585664df32b0567644d1ab8a2a5c2829d1f520c68c93e0c424.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get entities from administrativeUnits
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitsRequestBuilder) Get(options *AdministrativeUnitsRequestBuilderGetOptions)(*AdministrativeUnitsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdministrativeUnitsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*AdministrativeUnitsResponse), nil
}
func (m *AdministrativeUnitsRequestBuilder) GetByIds()(*i277b79a3b313ac8255d78ae90b5e9c683fe411d78cfeb5e1e83d31d38c270e05.GetByIdsRequestBuilder) {
    return i277b79a3b313ac8255d78ae90b5e9c683fe411d78cfeb5e1e83d31d38c270e05.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitsRequestBuilder) GetUserOwnedObjects()(*ia5a3db4673512570eb5641390e08a082576bca30d7028287114d44432947b385.GetUserOwnedObjectsRequestBuilder) {
    return ia5a3db4673512570eb5641390e08a082576bca30d7028287114d44432947b385.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Add new entity to administrativeUnits
// Parameters:
//  - options : Options for the request
func (m *AdministrativeUnitsRequestBuilder) Post(options *AdministrativeUnitsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAdministrativeUnit() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit), nil
}
func (m *AdministrativeUnitsRequestBuilder) ValidateProperties()(*ic5b5428ae35afae2a32cefcf30c64f6ab1fbe0e8909656a8a35a1a32357418eb.ValidatePropertiesRequestBuilder) {
    return ic5b5428ae35afae2a32cefcf30c64f6ab1fbe0e8909656a8a35a1a32357418eb.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
