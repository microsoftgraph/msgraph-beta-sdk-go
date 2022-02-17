package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/targetschedule"
    i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/appscope"
    i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/principal"
    i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/activatedusing"
    iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/roledefinition"
    ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/directoryscope"
    id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/cancel"
)

// UnifiedRoleAssignmentScheduleRequestRequestBuilder builds and executes requests for operations under \roleManagement\entitlementManagement\roleAssignmentScheduleRequests\{unifiedRoleAssignmentScheduleRequest-id}
type UnifiedRoleAssignmentScheduleRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentScheduleRequestRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentScheduleRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentScheduleRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters get roleAssignmentScheduleRequests from roleManagement
type UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentScheduleRequestRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentScheduleRequestRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) ActivatedUsing()(*i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442.ActivatedUsingRequestBuilder) {
    return i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442.NewActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) AppScope()(*i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09.AppScopeRequestBuilder) {
    return i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Cancel()(*id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a.CancelRequestBuilder) {
    return id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestRequestBuilder) {
    m := &UnifiedRoleAssignmentScheduleRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentScheduleRequestRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Delete(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderDeleteOptions)(error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) DirectoryScope()(*ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846.DirectoryScopeRequestBuilder) {
    return ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Get(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleAssignmentScheduleRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest), nil
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Patch(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Principal()(*i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf.PrincipalRequestBuilder) {
    return i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) RoleDefinition()(*iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0.RoleDefinitionRequestBuilder) {
    return iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) TargetSchedule()(*i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4.TargetScheduleRequestBuilder) {
    return i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
