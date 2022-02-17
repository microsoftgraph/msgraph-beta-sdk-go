package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1ad3f3772b94a4fe9eb4a735934236e8263a9a5e8004b1d911e619b2bfbfe3d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/directoryscope"
    i7fbfdd8acfde58648f2e9802cefa0c4426e7a1189e1de1969339baa0c9777c91 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/principal"
    i8625b9843234464c4fef92338104b70660fbcd1c0a46aa2315dcd2d32e2a8289 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/cancel"
    ia2562bf876dabdf8b8f2d1964328f124aa93d0bdee7008197fc224aa17175c37 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/appscope"
    icc542c1f51bf387e280de865ad6853cc67cc8659b96b1ba75c7efcf4ca21486e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/roledefinition"
    ifa3dab5fb247ec11de2a5930325058c34ca3959c5ac61b988f2682b273c02557 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item/targetschedule"
)

// UnifiedRoleEligibilityScheduleRequestRequestBuilder builds and executes requests for operations under \roleManagement\entitlementManagement\roleEligibilityScheduleRequests\{unifiedRoleEligibilityScheduleRequest-id}
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
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) AppScope()(*ia2562bf876dabdf8b8f2d1964328f124aa93d0bdee7008197fc224aa17175c37.AppScopeRequestBuilder) {
    return ia2562bf876dabdf8b8f2d1964328f124aa93d0bdee7008197fc224aa17175c37.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) Cancel()(*i8625b9843234464c4fef92338104b70660fbcd1c0a46aa2315dcd2d32e2a8289.CancelRequestBuilder) {
    return i8625b9843234464c4fef92338104b70660fbcd1c0a46aa2315dcd2d32e2a8289.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleEligibilityScheduleRequestRequestBuilderInternal instantiates a new UnifiedRoleEligibilityScheduleRequestRequestBuilder and sets the default values.
func NewUnifiedRoleEligibilityScheduleRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleEligibilityScheduleRequestRequestBuilder) {
    m := &UnifiedRoleEligibilityScheduleRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest_id}{?select,expand}";
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
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) DirectoryScope()(*i1ad3f3772b94a4fe9eb4a735934236e8263a9a5e8004b1d911e619b2bfbfe3d9.DirectoryScopeRequestBuilder) {
    return i1ad3f3772b94a4fe9eb4a735934236e8263a9a5e8004b1d911e619b2bfbfe3d9.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) Principal()(*i7fbfdd8acfde58648f2e9802cefa0c4426e7a1189e1de1969339baa0c9777c91.PrincipalRequestBuilder) {
    return i7fbfdd8acfde58648f2e9802cefa0c4426e7a1189e1de1969339baa0c9777c91.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) RoleDefinition()(*icc542c1f51bf387e280de865ad6853cc67cc8659b96b1ba75c7efcf4ca21486e.RoleDefinitionRequestBuilder) {
    return icc542c1f51bf387e280de865ad6853cc67cc8659b96b1ba75c7efcf4ca21486e.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleEligibilityScheduleRequestRequestBuilder) TargetSchedule()(*ifa3dab5fb247ec11de2a5930325058c34ca3959c5ac61b988f2682b273c02557.TargetScheduleRequestBuilder) {
    return ifa3dab5fb247ec11de2a5930325058c34ca3959c5ac61b988f2682b273c02557.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
