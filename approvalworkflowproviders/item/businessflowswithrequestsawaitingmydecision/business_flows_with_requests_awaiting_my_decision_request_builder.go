package businessflowswithrequestsawaitingmydecision

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \approvalWorkflowProviders\{approvalWorkflowProvider-id}\businessFlowsWithRequestsAwaitingMyDecision
type BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
type BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetQueryParameters struct {
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
type BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BusinessFlow;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    m := &BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/approvalWorkflowProviders/{approvalWorkflowProvider_id}/businessFlowsWithRequestsAwaitingMyDecision{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
// Parameters:
//  - options : Options for the request
func (m *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) CreateGetRequestInformation(options *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Create new navigation property to businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
// Parameters:
//  - options : Options for the request
func (m *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) CreatePostRequestInformation(options *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
// Parameters:
//  - options : Options for the request
func (m *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) Get(options *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderGetOptions)(*BusinessFlowsWithRequestsAwaitingMyDecisionResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBusinessFlowsWithRequestsAwaitingMyDecisionResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*BusinessFlowsWithRequestsAwaitingMyDecisionResponse), nil
}
// Create new navigation property to businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
// Parameters:
//  - options : Options for the request
func (m *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) Post(options *BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BusinessFlow, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBusinessFlow() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BusinessFlow), nil
}
