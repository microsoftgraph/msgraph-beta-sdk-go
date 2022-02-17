package roleeligibilityschedulerequests

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id2a1798ae235595af9a431da257553ca3777ff6c2d2a287f1ea055513c0d8949 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/filterbycurrentuserwithon"
)

// RoleEligibilityScheduleRequestsRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleEligibilityScheduleRequests
type RoleEligibilityScheduleRequestsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleEligibilityScheduleRequestsRequestBuilderGetOptions options for Get
type RoleEligibilityScheduleRequestsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters get roleEligibilityScheduleRequests from roleManagement
type RoleEligibilityScheduleRequestsRequestBuilderGetQueryParameters struct {
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
// RoleEligibilityScheduleRequestsRequestBuilderPostOptions options for Post
type RoleEligibilityScheduleRequestsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRoleEligibilityScheduleRequestsRequestBuilderInternal instantiates a new RoleEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewRoleEligibilityScheduleRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleEligibilityScheduleRequestsRequestBuilder) {
    m := &RoleEligibilityScheduleRequestsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleEligibilityScheduleRequestsRequestBuilder instantiates a new RoleEligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewRoleEligibilityScheduleRequestsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleEligibilityScheduleRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleEligibilityScheduleRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get roleEligibilityScheduleRequests from roleManagement
func (m *RoleEligibilityScheduleRequestsRequestBuilder) CreateGetRequestInformation(options *RoleEligibilityScheduleRequestsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to roleEligibilityScheduleRequests for roleManagement
func (m *RoleEligibilityScheduleRequestsRequestBuilder) CreatePostRequestInformation(options *RoleEligibilityScheduleRequestsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \roleManagement\directory\roleEligibilityScheduleRequests\microsoft.graph.filterByCurrentUser(on={on})
func (m *RoleEligibilityScheduleRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*id2a1798ae235595af9a431da257553ca3777ff6c2d2a287f1ea055513c0d8949.FilterByCurrentUserWithOnRequestBuilder) {
    return id2a1798ae235595af9a431da257553ca3777ff6c2d2a287f1ea055513c0d8949.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get get roleEligibilityScheduleRequests from roleManagement
func (m *RoleEligibilityScheduleRequestsRequestBuilder) Get(options *RoleEligibilityScheduleRequestsRequestBuilderGetOptions)(*RoleEligibilityScheduleRequestsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleEligibilityScheduleRequestsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RoleEligibilityScheduleRequestsResponse), nil
}
// Post create new navigation property to roleEligibilityScheduleRequests for roleManagement
func (m *RoleEligibilityScheduleRequestsRequestBuilder) Post(options *RoleEligibilityScheduleRequestsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleRequest, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleEligibilityScheduleRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleRequest), nil
}
