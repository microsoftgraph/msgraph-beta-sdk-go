package roleassignmentapprovals

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id51c04bdca1b3d3df1fb0cf1a9242cd201c579ca07ebe1d61ada8f6110ca302e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentapprovals/filterbycurrentuserwithon"
)

// RoleAssignmentApprovalsRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleAssignmentApprovals
type RoleAssignmentApprovalsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleAssignmentApprovalsRequestBuilderGetOptions options for Get
type RoleAssignmentApprovalsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleAssignmentApprovalsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleAssignmentApprovalsRequestBuilderGetQueryParameters get roleAssignmentApprovals from roleManagement
type RoleAssignmentApprovalsRequestBuilderGetQueryParameters struct {
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
// RoleAssignmentApprovalsRequestBuilderPostOptions options for Post
type RoleAssignmentApprovalsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Approval;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRoleAssignmentApprovalsRequestBuilderInternal instantiates a new RoleAssignmentApprovalsRequestBuilder and sets the default values.
func NewRoleAssignmentApprovalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleAssignmentApprovalsRequestBuilder) {
    m := &RoleAssignmentApprovalsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleAssignmentApprovals{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleAssignmentApprovalsRequestBuilder instantiates a new RoleAssignmentApprovalsRequestBuilder and sets the default values.
func NewRoleAssignmentApprovalsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleAssignmentApprovalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleAssignmentApprovalsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get roleAssignmentApprovals from roleManagement
func (m *RoleAssignmentApprovalsRequestBuilder) CreateGetRequestInformation(options *RoleAssignmentApprovalsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to roleAssignmentApprovals for roleManagement
func (m *RoleAssignmentApprovalsRequestBuilder) CreatePostRequestInformation(options *RoleAssignmentApprovalsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \roleManagement\directory\roleAssignmentApprovals\microsoft.graph.filterByCurrentUser(on={on})
func (m *RoleAssignmentApprovalsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*id51c04bdca1b3d3df1fb0cf1a9242cd201c579ca07ebe1d61ada8f6110ca302e.FilterByCurrentUserWithOnRequestBuilder) {
    return id51c04bdca1b3d3df1fb0cf1a9242cd201c579ca07ebe1d61ada8f6110ca302e.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get get roleAssignmentApprovals from roleManagement
func (m *RoleAssignmentApprovalsRequestBuilder) Get(options *RoleAssignmentApprovalsRequestBuilderGetOptions)(*RoleAssignmentApprovalsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleAssignmentApprovalsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RoleAssignmentApprovalsResponse), nil
}
// Post create new navigation property to roleAssignmentApprovals for roleManagement
func (m *RoleAssignmentApprovalsRequestBuilder) Post(options *RoleAssignmentApprovalsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Approval, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewApproval() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Approval), nil
}
