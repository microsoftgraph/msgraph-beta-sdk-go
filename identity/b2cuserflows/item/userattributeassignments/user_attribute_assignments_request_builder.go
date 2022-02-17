package userattributeassignments

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0fb36e749699428077c12ff5812a9bb575834815e4d7e3a86acc87c26c8c32a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userattributeassignments/getorder"
    i5240f41d615de8c637c0c8acdf6dc7282d58b03f61ba028038bdbaa1649ff59f "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userattributeassignments/setorder"
)

// UserAttributeAssignmentsRequestBuilder builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}\userAttributeAssignments
type UserAttributeAssignmentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserAttributeAssignmentsRequestBuilderGetOptions options for Get
type UserAttributeAssignmentsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserAttributeAssignmentsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserAttributeAssignmentsRequestBuilderGetQueryParameters the user attribute assignments included in the user flow.
type UserAttributeAssignmentsRequestBuilderGetQueryParameters struct {
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
// UserAttributeAssignmentsRequestBuilderPostOptions options for Post
type UserAttributeAssignmentsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityUserFlowAttributeAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserAttributeAssignmentsRequestBuilderInternal instantiates a new UserAttributeAssignmentsRequestBuilder and sets the default values.
func NewUserAttributeAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserAttributeAssignmentsRequestBuilder) {
    m := &UserAttributeAssignmentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow_id}/userAttributeAssignments{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserAttributeAssignmentsRequestBuilder instantiates a new UserAttributeAssignmentsRequestBuilder and sets the default values.
func NewUserAttributeAssignmentsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserAttributeAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserAttributeAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the user attribute assignments included in the user flow.
func (m *UserAttributeAssignmentsRequestBuilder) CreateGetRequestInformation(options *UserAttributeAssignmentsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation the user attribute assignments included in the user flow.
func (m *UserAttributeAssignmentsRequestBuilder) CreatePostRequestInformation(options *UserAttributeAssignmentsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the user attribute assignments included in the user flow.
func (m *UserAttributeAssignmentsRequestBuilder) Get(options *UserAttributeAssignmentsRequestBuilderGetOptions)(*UserAttributeAssignmentsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserAttributeAssignmentsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UserAttributeAssignmentsResponse), nil
}
// GetOrder builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}\userAttributeAssignments\microsoft.graph.getOrder()
func (m *UserAttributeAssignmentsRequestBuilder) GetOrder()(*i0fb36e749699428077c12ff5812a9bb575834815e4d7e3a86acc87c26c8c32a8.GetOrderRequestBuilder) {
    return i0fb36e749699428077c12ff5812a9bb575834815e4d7e3a86acc87c26c8c32a8.NewGetOrderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post the user attribute assignments included in the user flow.
func (m *UserAttributeAssignmentsRequestBuilder) Post(options *UserAttributeAssignmentsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityUserFlowAttributeAssignment, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIdentityUserFlowAttributeAssignment() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityUserFlowAttributeAssignment), nil
}
func (m *UserAttributeAssignmentsRequestBuilder) SetOrder()(*i5240f41d615de8c637c0c8acdf6dc7282d58b03f61ba028038bdbaa1649ff59f.SetOrderRequestBuilder) {
    return i5240f41d615de8c637c0c8acdf6dc7282d58b03f61ba028038bdbaa1649ff59f.NewSetOrderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
