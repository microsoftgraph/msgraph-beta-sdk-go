package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i28e55a47e788f2facc1660de24a086958a7918679925a5ce71516e5d9f73e4b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/cancel"
    i2cef0677a41a89f3b35ba8a969822c8bd74d3f2b806471c99641128a74c6d983 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/targetschedule"
    i42b41d4fd35db87fa4e585e05c89566c3db5f09fd0ba9febd2c8f4f83f4f3171 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/directoryscope"
    icbc656be4360057e1ea1d730a350a6babae196e7449d637ea202e68cf814697c "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/principal"
    id474be08b30a83047268aeb159d98ccab694e6d09f2eb1077120c6d6454d819e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/appscope"
    idf7b5f031542e8138350d419b142bb2e4921b774d4b2174b040680b2cc9094b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item/roledefinition"
)

// UnifiedRoleEligibilityScheduleRequestRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleEligibilityScheduleRequests\{unifiedRoleEligibilityScheduleRequest-id}
type UnifiedRoleEligibilityScheduleRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleEligibilityScheduleRequestRequestBuilderDeleteOptions options for Delete
type UnifiedRoleEligibilityScheduleRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleEligibilityScheduleRequestRequestBuilderGetOptions options for Get
type UnifiedRoleEligibilityScheduleRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleEligibilityScheduleRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleEligibilityScheduleRequestRequestBuilderGetQueryParameters get roleEligibilityScheduleRequests from roleManagement
type UnifiedRoleEligibilityScheduleRequestRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleEligibilityScheduleRequestRequestBuilderPatchOptions options for Patch
type UnifiedRoleEligibilityScheduleRequestRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) AppScope()(*id474be08b30a83047268aeb159d98ccab694e6d09f2eb1077120c6d6454d819e.AppScopeRequestBuilder) {
    return id474be08b30a83047268aeb159d98ccab694e6d09f2eb1077120c6d6454d819e.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) Cancel()(*i28e55a47e788f2facc1660de24a086958a7918679925a5ce71516e5d9f73e4b4.CancelRequestBuilder) {
    return i28e55a47e788f2facc1660de24a086958a7918679925a5ce71516e5d9f73e4b4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleEligibilityScheduleRequestRequestBuilderInternal instantiates a new UnifiedRoleEligibilityScheduleRequestRequestBuilder and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleEligibilityScheduleRequestRequestBuilder) {
    m := &UnifiedRoleEligibilityScheduleRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleEligibilityScheduleRequestRequestBuilder instantiates a new UnifiedRoleEligibilityScheduleRequestRequestBuilder and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleEligibilityScheduleRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleEligibilityScheduleRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleEligibilityScheduleRequests for roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleEligibilityScheduleRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation get roleEligibilityScheduleRequests from roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleEligibilityScheduleRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleEligibilityScheduleRequests in roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleEligibilityScheduleRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete delete navigation property roleEligibilityScheduleRequests for roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) Delete(options *UnifiedRoleEligibilityScheduleRequestRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) DirectoryScope()(*i42b41d4fd35db87fa4e585e05c89566c3db5f09fd0ba9febd2c8f4f83f4f3171.DirectoryScopeRequestBuilder) {
    return i42b41d4fd35db87fa4e585e05c89566c3db5f09fd0ba9febd2c8f4f83f4f3171.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleEligibilityScheduleRequests from roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) Get(options *UnifiedRoleEligibilityScheduleRequestRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleEligibilityScheduleRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleRequest), nil
}
// Patch update the navigation property roleEligibilityScheduleRequests in roleManagement
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) Patch(options *UnifiedRoleEligibilityScheduleRequestRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) Principal()(*icbc656be4360057e1ea1d730a350a6babae196e7449d637ea202e68cf814697c.PrincipalRequestBuilder) {
    return icbc656be4360057e1ea1d730a350a6babae196e7449d637ea202e68cf814697c.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) RoleDefinition()(*idf7b5f031542e8138350d419b142bb2e4921b774d4b2174b040680b2cc9094b8.RoleDefinitionRequestBuilder) {
    return idf7b5f031542e8138350d419b142bb2e4921b774d4b2174b040680b2cc9094b8.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) TargetSchedule()(*i2cef0677a41a89f3b35ba8a969822c8bd74d3f2b806471c99641128a74c6d983.TargetScheduleRequestBuilder) {
    return i2cef0677a41a89f3b35ba8a969822c8bd74d3f2b806471c99641128a74c6d983.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
