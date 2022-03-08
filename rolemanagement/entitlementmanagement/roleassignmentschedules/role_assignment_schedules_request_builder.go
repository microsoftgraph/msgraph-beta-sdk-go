package roleassignmentschedules

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i02b700ff9bfe2c2e9adf73ddd9baee36ab851b4ed1601020e1a718900b05e877 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedules/filterbycurrentuserwithon"
    i572d617e820eedc64c47dbb92112001310a1fc7b1ea5b64f88edc5ec9ae1b304 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedules/count"
)

// RoleAssignmentSchedulesRequestBuilder provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
type RoleAssignmentSchedulesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleAssignmentSchedulesRequestBuilderGetOptions options for Get
type RoleAssignmentSchedulesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleAssignmentSchedulesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleAssignmentSchedulesRequestBuilderGetQueryParameters get roleAssignmentSchedules from roleManagement
type RoleAssignmentSchedulesRequestBuilderGetQueryParameters struct {
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
// RoleAssignmentSchedulesRequestBuilderPostOptions options for Post
type RoleAssignmentSchedulesRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRoleAssignmentSchedulesRequestBuilderInternal instantiates a new RoleAssignmentSchedulesRequestBuilder and sets the default values.
func NewRoleAssignmentSchedulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleAssignmentSchedulesRequestBuilder) {
    m := &RoleAssignmentSchedulesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentSchedules{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleAssignmentSchedulesRequestBuilder instantiates a new RoleAssignmentSchedulesRequestBuilder and sets the default values.
func NewRoleAssignmentSchedulesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleAssignmentSchedulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleAssignmentSchedulesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RoleAssignmentSchedulesRequestBuilder) Count()(*i572d617e820eedc64c47dbb92112001310a1fc7b1ea5b64f88edc5ec9ae1b304.CountRequestBuilder) {
    return i572d617e820eedc64c47dbb92112001310a1fc7b1ea5b64f88edc5ec9ae1b304.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get roleAssignmentSchedules from roleManagement
func (m *RoleAssignmentSchedulesRequestBuilder) CreateGetRequestInformation(options *RoleAssignmentSchedulesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to roleAssignmentSchedules for roleManagement
func (m *RoleAssignmentSchedulesRequestBuilder) CreatePostRequestInformation(options *RoleAssignmentSchedulesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
func (m *RoleAssignmentSchedulesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*i02b700ff9bfe2c2e9adf73ddd9baee36ab851b4ed1601020e1a718900b05e877.FilterByCurrentUserWithOnRequestBuilder) {
    return i02b700ff9bfe2c2e9adf73ddd9baee36ab851b4ed1601020e1a718900b05e877.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get get roleAssignmentSchedules from roleManagement
func (m *RoleAssignmentSchedulesRequestBuilder) Get(options *RoleAssignmentSchedulesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUnifiedRoleAssignmentScheduleCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleCollectionResponseable), nil
}
// Post create new navigation property to roleAssignmentSchedules for roleManagement
func (m *RoleAssignmentSchedulesRequestBuilder) Post(options *RoleAssignmentSchedulesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleable), nil
}
