package roleassignmentscheduleinstances

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia8c4cab071c96e21125cbe53f785c45da39afa5b7fa02d3706f75bbc7425f75e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentscheduleinstances/filterbycurrentuserwithon"
)

// RoleAssignmentScheduleInstancesRequestBuilder builds and executes requests for operations under \roleManagement\entitlementManagement\roleAssignmentScheduleInstances
type RoleAssignmentScheduleInstancesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleAssignmentScheduleInstancesRequestBuilderGetOptions options for Get
type RoleAssignmentScheduleInstancesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleAssignmentScheduleInstancesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleAssignmentScheduleInstancesRequestBuilderGetQueryParameters get roleAssignmentScheduleInstances from roleManagement
type RoleAssignmentScheduleInstancesRequestBuilderGetQueryParameters struct {
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
// RoleAssignmentScheduleInstancesRequestBuilderPostOptions options for Post
type RoleAssignmentScheduleInstancesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleInstance;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRoleAssignmentScheduleInstancesRequestBuilderInternal instantiates a new RoleAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewRoleAssignmentScheduleInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleAssignmentScheduleInstancesRequestBuilder) {
    m := &RoleAssignmentScheduleInstancesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleInstances{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleAssignmentScheduleInstancesRequestBuilder instantiates a new RoleAssignmentScheduleInstancesRequestBuilder and sets the default values.
func NewRoleAssignmentScheduleInstancesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleAssignmentScheduleInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleAssignmentScheduleInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get roleAssignmentScheduleInstances from roleManagement
func (m *RoleAssignmentScheduleInstancesRequestBuilder) CreateGetRequestInformation(options *RoleAssignmentScheduleInstancesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to roleAssignmentScheduleInstances for roleManagement
func (m *RoleAssignmentScheduleInstancesRequestBuilder) CreatePostRequestInformation(options *RoleAssignmentScheduleInstancesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \roleManagement\entitlementManagement\roleAssignmentScheduleInstances\microsoft.graph.filterByCurrentUser(on={on})
func (m *RoleAssignmentScheduleInstancesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*ia8c4cab071c96e21125cbe53f785c45da39afa5b7fa02d3706f75bbc7425f75e.FilterByCurrentUserWithOnRequestBuilder) {
    return ia8c4cab071c96e21125cbe53f785c45da39afa5b7fa02d3706f75bbc7425f75e.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get get roleAssignmentScheduleInstances from roleManagement
func (m *RoleAssignmentScheduleInstancesRequestBuilder) Get(options *RoleAssignmentScheduleInstancesRequestBuilderGetOptions)(*RoleAssignmentScheduleInstancesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleAssignmentScheduleInstancesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RoleAssignmentScheduleInstancesResponse), nil
}
// Post create new navigation property to roleAssignmentScheduleInstances for roleManagement
func (m *RoleAssignmentScheduleInstancesRequestBuilder) Post(options *RoleAssignmentScheduleInstancesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleInstance, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleAssignmentScheduleInstance() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleInstance), nil
}
