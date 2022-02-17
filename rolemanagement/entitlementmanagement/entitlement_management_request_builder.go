package entitlementmanagement

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1baa975bee8bc1fcf62ee7c7d05b837d65d8d022e7556de9e284932ea41c5624 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedules"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5e09c2b6a4bc955be405d7fbb453204f96818f8aae7508864d080577c5491cd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests"
    i75ab2bc44492a09208082e3507fb5b1add85563cc4e6f15226038f431107cc36 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/resourcenamespaces"
    i7f3b8558d0a039f46114bbef7c414667ed2c5dd0d28b58faa7b09850032d46db "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/rolescheduleswithdirectoryscopeidwithappscopeidwithprincipalidwithroledefinitionid"
    i8efb8a53874fa360f1b73def8f68e4e407323fb5958457abad962b8900d25355 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests"
    ia143a4210940ae95c6f24da8e1def30da0952553539019a9c5e037acaae6d536 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedules"
    ic1ac1019e9fecf7777c6fabe94d5263065004cf3d093ea9a50afcb7544792459 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/rolescheduleinstanceswithdirectoryscopeidwithappscopeidwithprincipalidwithroledefinitionid"
    ic74edb001b0ce5f97639bc39b485bc37fc8d4b308c0e093fb317822e211e74ab "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityscheduleinstances"
    icfcbd750b986c89f05029c8a1c38ed5d1e850b098c1de07332c9c8f4dba9e621 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roledefinitions"
    idaef4275a9c6a8af3da51ed597b3772552b76155312fa2174e893913ec186b9e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentapprovals"
    ie577e072e7edde2d3f423e1ceb296400868e7f5e7bddbe2ee934c71bd8d1d637 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentscheduleinstances"
    if94439b0eb19f66d0fe98bc2e50b7664597fad69cfac1e1132b6b13d328ae667 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignments"
    i1094763ed645fa459b11fe2fd0dbbb4cd5b4e5abf8fb494633f130f88a198ab7 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item"
    i1f3fe61534b224c6fac2a1c5c69508f0db3518d4530ee88be6548eae659bb75a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests/item"
    i2e767808e39342c5e075a8c93310b2c0591c0c806546f5b58190d82a94993717 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignments/item"
    i6f70cfe1f611776ca2d5cdc9b36b624ba1b69dda981099ccaaf42d45453807a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentscheduleinstances/item"
    i94f6c4184f5f2637cb72ccacd72d05ca460673ced94157d6b941e79b19f921e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentapprovals/item"
    ib2c51c75dd96d44353372f4ec3d1bd31e89c02a24c89600c2d3fb288758d1bf3 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roledefinitions/item"
    ib97806c66498b497946ee83f7bb51d300c1952475468eabbe9d487576ab31775 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedules/item"
    icf1ccc36301949ddd8b0254d07f44b375f6f95d7adb58a22b17ccfc55ae25abc "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityscheduleinstances/item"
    id9a0f06cc5af154fe717e7b85c2d5db50242eb2a5eb5b84cd9b6f2bfff25c1f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedules/item"
    ie58dbe4c0951e87ffefa6af6faabb1e4ad0651b2ea404363bfd8a2d38cc97a53 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/resourcenamespaces/item"
)

// EntitlementManagementRequestBuilder builds and executes requests for operations under \roleManagement\entitlementManagement
type EntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EntitlementManagementRequestBuilderDeleteOptions options for Delete
type EntitlementManagementRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EntitlementManagementRequestBuilderGetOptions options for Get
type EntitlementManagementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EntitlementManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EntitlementManagementRequestBuilderGetQueryParameters container for all entitlement management resources in Azure AD identity governance.
type EntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EntitlementManagementRequestBuilderPatchOptions options for Patch
type EntitlementManagementRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplication;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    m := &EntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation container for all entitlement management resources in Azure AD identity governance.
func (m *EntitlementManagementRequestBuilder) CreateDeleteRequestInformation(options *EntitlementManagementRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation container for all entitlement management resources in Azure AD identity governance.
func (m *EntitlementManagementRequestBuilder) CreateGetRequestInformation(options *EntitlementManagementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation container for all entitlement management resources in Azure AD identity governance.
func (m *EntitlementManagementRequestBuilder) CreatePatchRequestInformation(options *EntitlementManagementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete container for all entitlement management resources in Azure AD identity governance.
func (m *EntitlementManagementRequestBuilder) Delete(options *EntitlementManagementRequestBuilderDeleteOptions)(error) {
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
// Get container for all entitlement management resources in Azure AD identity governance.
func (m *EntitlementManagementRequestBuilder) Get(options *EntitlementManagementRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplication, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRbacApplication() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplication), nil
}
// Patch container for all entitlement management resources in Azure AD identity governance.
func (m *EntitlementManagementRequestBuilder) Patch(options *EntitlementManagementRequestBuilderPatchOptions)(error) {
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
func (m *EntitlementManagementRequestBuilder) ResourceNamespaces()(*i75ab2bc44492a09208082e3507fb5b1add85563cc4e6f15226038f431107cc36.ResourceNamespacesRequestBuilder) {
    return i75ab2bc44492a09208082e3507fb5b1add85563cc4e6f15226038f431107cc36.NewResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceNamespacesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.resourceNamespaces.item collection
func (m *EntitlementManagementRequestBuilder) ResourceNamespacesById(id string)(*ie58dbe4c0951e87ffefa6af6faabb1e4ad0651b2ea404363bfd8a2d38cc97a53.UnifiedRbacResourceNamespaceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace_id"] = id
    }
    return ie58dbe4c0951e87ffefa6af6faabb1e4ad0651b2ea404363bfd8a2d38cc97a53.NewUnifiedRbacResourceNamespaceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleAssignmentApprovals()(*idaef4275a9c6a8af3da51ed597b3772552b76155312fa2174e893913ec186b9e.RoleAssignmentApprovalsRequestBuilder) {
    return idaef4275a9c6a8af3da51ed597b3772552b76155312fa2174e893913ec186b9e.NewRoleAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentApprovalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignmentApprovals.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentApprovalsById(id string)(*i94f6c4184f5f2637cb72ccacd72d05ca460673ced94157d6b941e79b19f921e8.ApprovalRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval_id"] = id
    }
    return i94f6c4184f5f2637cb72ccacd72d05ca460673ced94157d6b941e79b19f921e8.NewApprovalRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleAssignments()(*if94439b0eb19f66d0fe98bc2e50b7664597fad69cfac1e1132b6b13d328ae667.RoleAssignmentsRequestBuilder) {
    return if94439b0eb19f66d0fe98bc2e50b7664597fad69cfac1e1132b6b13d328ae667.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignments.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentsById(id string)(*i2e767808e39342c5e075a8c93310b2c0591c0c806546f5b58190d82a94993717.UnifiedRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment_id"] = id
    }
    return i2e767808e39342c5e075a8c93310b2c0591c0c806546f5b58190d82a94993717.NewUnifiedRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleInstances()(*ie577e072e7edde2d3f423e1ceb296400868e7f5e7bddbe2ee934c71bd8d1d637.RoleAssignmentScheduleInstancesRequestBuilder) {
    return ie577e072e7edde2d3f423e1ceb296400868e7f5e7bddbe2ee934c71bd8d1d637.NewRoleAssignmentScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignmentScheduleInstances.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleInstancesById(id string)(*i6f70cfe1f611776ca2d5cdc9b36b624ba1b69dda981099ccaaf42d45453807a7.UnifiedRoleAssignmentScheduleInstanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance_id"] = id
    }
    return i6f70cfe1f611776ca2d5cdc9b36b624ba1b69dda981099ccaaf42d45453807a7.NewUnifiedRoleAssignmentScheduleInstanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleRequests()(*i8efb8a53874fa360f1b73def8f68e4e407323fb5958457abad962b8900d25355.RoleAssignmentScheduleRequestsRequestBuilder) {
    return i8efb8a53874fa360f1b73def8f68e4e407323fb5958457abad962b8900d25355.NewRoleAssignmentScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignmentScheduleRequests.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleRequestsById(id string)(*i1094763ed645fa459b11fe2fd0dbbb4cd5b4e5abf8fb494633f130f88a198ab7.UnifiedRoleAssignmentScheduleRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleRequest_id"] = id
    }
    return i1094763ed645fa459b11fe2fd0dbbb4cd5b4e5abf8fb494633f130f88a198ab7.NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleAssignmentSchedules()(*i1baa975bee8bc1fcf62ee7c7d05b837d65d8d022e7556de9e284932ea41c5624.RoleAssignmentSchedulesRequestBuilder) {
    return i1baa975bee8bc1fcf62ee7c7d05b837d65d8d022e7556de9e284932ea41c5624.NewRoleAssignmentSchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentSchedulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignmentSchedules.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentSchedulesById(id string)(*ib97806c66498b497946ee83f7bb51d300c1952475468eabbe9d487576ab31775.UnifiedRoleAssignmentScheduleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentSchedule_id"] = id
    }
    return ib97806c66498b497946ee83f7bb51d300c1952475468eabbe9d487576ab31775.NewUnifiedRoleAssignmentScheduleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleDefinitions()(*icfcbd750b986c89f05029c8a1c38ed5d1e850b098c1de07332c9c8f4dba9e621.RoleDefinitionsRequestBuilder) {
    return icfcbd750b986c89f05029c8a1c38ed5d1e850b098c1de07332c9c8f4dba9e621.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleDefinitions.item collection
func (m *EntitlementManagementRequestBuilder) RoleDefinitionsById(id string)(*ib2c51c75dd96d44353372f4ec3d1bd31e89c02a24c89600c2d3fb288758d1bf3.UnifiedRoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition_id"] = id
    }
    return ib2c51c75dd96d44353372f4ec3d1bd31e89c02a24c89600c2d3fb288758d1bf3.NewUnifiedRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleInstances()(*ic74edb001b0ce5f97639bc39b485bc37fc8d4b308c0e093fb317822e211e74ab.RoleEligibilityScheduleInstancesRequestBuilder) {
    return ic74edb001b0ce5f97639bc39b485bc37fc8d4b308c0e093fb317822e211e74ab.NewRoleEligibilityScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleEligibilityScheduleInstances.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleInstancesById(id string)(*icf1ccc36301949ddd8b0254d07f44b375f6f95d7adb58a22b17ccfc55ae25abc.UnifiedRoleEligibilityScheduleInstanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleInstance_id"] = id
    }
    return icf1ccc36301949ddd8b0254d07f44b375f6f95d7adb58a22b17ccfc55ae25abc.NewUnifiedRoleEligibilityScheduleInstanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleRequests()(*i5e09c2b6a4bc955be405d7fbb453204f96818f8aae7508864d080577c5491cd0.RoleEligibilityScheduleRequestsRequestBuilder) {
    return i5e09c2b6a4bc955be405d7fbb453204f96818f8aae7508864d080577c5491cd0.NewRoleEligibilityScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleEligibilityScheduleRequests.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleRequestsById(id string)(*i1f3fe61534b224c6fac2a1c5c69508f0db3518d4530ee88be6548eae659bb75a.UnifiedRoleEligibilityScheduleRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest_id"] = id
    }
    return i1f3fe61534b224c6fac2a1c5c69508f0db3518d4530ee88be6548eae659bb75a.NewUnifiedRoleEligibilityScheduleRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EntitlementManagementRequestBuilder) RoleEligibilitySchedules()(*ia143a4210940ae95c6f24da8e1def30da0952553539019a9c5e037acaae6d536.RoleEligibilitySchedulesRequestBuilder) {
    return ia143a4210940ae95c6f24da8e1def30da0952553539019a9c5e037acaae6d536.NewRoleEligibilitySchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilitySchedulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleEligibilitySchedules.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilitySchedulesById(id string)(*id9a0f06cc5af154fe717e7b85c2d5db50242eb2a5eb5b84cd9b6f2bfff25c1f4.UnifiedRoleEligibilityScheduleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilitySchedule_id"] = id
    }
    return id9a0f06cc5af154fe717e7b85c2d5db50242eb2a5eb5b84cd9b6f2bfff25c1f4.NewUnifiedRoleEligibilityScheduleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId builds and executes requests for operations under \roleManagement\entitlementManagement\microsoft.graph.roleScheduleInstances(directoryScopeId='{directoryScopeId}',appScopeId='{appScopeId}',principalId='{principalId}',roleDefinitionId='{roleDefinitionId}')
func (m *EntitlementManagementRequestBuilder) RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId(directoryScopeId *string, appScopeId *string, principalId *string, roleDefinitionId *string)(*ic1ac1019e9fecf7777c6fabe94d5263065004cf3d093ea9a50afcb7544792459.RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) {
    return ic1ac1019e9fecf7777c6fabe94d5263065004cf3d093ea9a50afcb7544792459.NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, directoryScopeId, appScopeId, principalId, roleDefinitionId);
}
// RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId builds and executes requests for operations under \roleManagement\entitlementManagement\microsoft.graph.roleSchedules(directoryScopeId='{directoryScopeId}',appScopeId='{appScopeId}',principalId='{principalId}',roleDefinitionId='{roleDefinitionId}')
func (m *EntitlementManagementRequestBuilder) RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId(directoryScopeId *string, appScopeId *string, principalId *string, roleDefinitionId *string)(*i7f3b8558d0a039f46114bbef7c414667ed2c5dd0d28b58faa7b09850032d46db.RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) {
    return i7f3b8558d0a039f46114bbef7c414667ed2c5dd0d28b58faa7b09850032d46db.NewRoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, directoryScopeId, appScopeId, principalId, roleDefinitionId);
}
