package roleeligibilityscheduleinstances

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i49fbc318735f19eacde3c858843d463f156cbb40bb6e7488adcd177086c9e868 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityscheduleinstances/filterbycurrentuserwithon"
)

// RoleEligibilityScheduleInstancesRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleEligibilityScheduleInstances
type RoleEligibilityScheduleInstancesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleEligibilityScheduleInstancesRequestBuilderGetOptions options for Get
type RoleEligibilityScheduleInstancesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleEligibilityScheduleInstancesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleEligibilityScheduleInstancesRequestBuilderGetQueryParameters get roleEligibilityScheduleInstances from roleManagement
type RoleEligibilityScheduleInstancesRequestBuilderGetQueryParameters struct {
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
// RoleEligibilityScheduleInstancesRequestBuilderPostOptions options for Post
type RoleEligibilityScheduleInstancesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleInstance;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRoleEligibilityScheduleInstancesRequestBuilderInternal instantiates a new RoleEligibilityScheduleInstancesRequestBuilder and sets the default values.
func NewRoleEligibilityScheduleInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleEligibilityScheduleInstancesRequestBuilder) {
    m := &RoleEligibilityScheduleInstancesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleEligibilityScheduleInstances{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleEligibilityScheduleInstancesRequestBuilder instantiates a new RoleEligibilityScheduleInstancesRequestBuilder and sets the default values.
func NewRoleEligibilityScheduleInstancesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleEligibilityScheduleInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleEligibilityScheduleInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get roleEligibilityScheduleInstances from roleManagement
func (m *RoleEligibilityScheduleInstancesRequestBuilder) CreateGetRequestInformation(options *RoleEligibilityScheduleInstancesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to roleEligibilityScheduleInstances for roleManagement
func (m *RoleEligibilityScheduleInstancesRequestBuilder) CreatePostRequestInformation(options *RoleEligibilityScheduleInstancesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \roleManagement\directory\roleEligibilityScheduleInstances\microsoft.graph.filterByCurrentUser(on={on})
func (m *RoleEligibilityScheduleInstancesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*i49fbc318735f19eacde3c858843d463f156cbb40bb6e7488adcd177086c9e868.FilterByCurrentUserWithOnRequestBuilder) {
    return i49fbc318735f19eacde3c858843d463f156cbb40bb6e7488adcd177086c9e868.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get get roleEligibilityScheduleInstances from roleManagement
func (m *RoleEligibilityScheduleInstancesRequestBuilder) Get(options *RoleEligibilityScheduleInstancesRequestBuilderGetOptions)(*RoleEligibilityScheduleInstancesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleEligibilityScheduleInstancesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RoleEligibilityScheduleInstancesResponse), nil
}
// Post create new navigation property to roleEligibilityScheduleInstances for roleManagement
func (m *RoleEligibilityScheduleInstancesRequestBuilder) Post(options *RoleEligibilityScheduleInstancesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleInstance, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleEligibilityScheduleInstance() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleInstance), nil
}
