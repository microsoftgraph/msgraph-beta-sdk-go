package directory

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i003e247115fdc2cbce0f60f7d8a2c16dfec08ea21559d4270399d03acc2b4081 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests"
    i1674488bab3f723cd92320bded03455762b2d10b76d65717248d2a39b768949d "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/transitiveroleassignments"
    i4196e83a4bae24d0aeb92656129c7f451fb9b312487c6af458d556759a773f07 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentapprovals"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i89936fb8dab16eb5ef6964ec9e5d8aa7c54f3f69934f902cdc4a2a400db45270 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/rolescheduleswithdirectoryscopeidwithappscopeidwithprincipalidwithroledefinitionid"
    i955424cf34f796ced961a38c0bb409d0018463520cd1d610c06c4e0187f8f095 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentscheduleinstances"
    i99441748b67ef831e7259b916e604bc0f3a134120b98e3f3cec7689707945a5b "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests"
    ib3f3f9905e601fb21a0927cffd109f585906d82dd65ac7b379b6df162c0beec5 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedules"
    ib7ee0af232e3f7f843b9fc80b2bf1bcd6bdc1aa1876b6682453a23c520194820 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/rolescheduleinstanceswithdirectoryscopeidwithappscopeidwithprincipalidwithroledefinitionid"
    ic09a099122154ece987c7f85636fcc11aac7d11f4704eb7ea18e15d32cb118be "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roledefinitions"
    ie00125ef63b154ab9fd6cbeba46882ef7ecf0a1bbe044d47b4e5fd94f87ff817 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityscheduleinstances"
    ie88f7587d0b612d6aad4a4cc6e4adf2df6596695dcfb937bd88d739a30fec4d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedules"
    if3967cc171135f8ae89f635c0a197e13e64649a0043e1dcc96293c3c4f5ebbbe "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignments"
    if7ad3988c943ab6c47fb727c18a8349d5e50ebec53ab6cc066c11a20717d46af "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/resourcenamespaces"
    i04dea7bcc2de46d7f933671fce512fa0699d5fc5d83576e03b7ac0fe57cc5704 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/resourcenamespaces/item"
    i0dc801eaa7e34f0715fe9868f950fd4b68e0050fd9669d9d65ab28af2e0d7f73 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignments/item"
    i15fef4d98a25fc03a5a29e7e66cec2ffeccb757e532a20aa82098881625f052a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roledefinitions/item"
    i1a1177924b99f8b8668ec4e5b307b5e9d537c2cfa54abdf2bf1fddde3af81044 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item"
    i2b3fa3ea5b323a0c2dca9d12e3d5f7bdadbc4bbf1e526ea51f794a5d0f803d75 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedules/item"
    i322381699c08d5273b2d2e8a72a478e9a4b2c80cb977ad926aa502c1353b499a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityscheduleinstances/item"
    i44fbfcce41c9eaa7ba9ff5f6ba24f61e5a9c03cdb24c3bc25beed0e1d6216960 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentapprovals/item"
    i5e4e5e8e6ab57f21346122383935dbe84160f647fde997fe85fe2ae4a427c50e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedules/item"
    id0ebe1fb3bbbfdf8b2d57215ab0ec7662faabe46ca681bc46803c9c70d11cab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleeligibilityschedulerequests/item"
    id3919a2c832ecf67620541df52d40952233a8c3c7f1293166f41f2e20a0f3aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/transitiveroleassignments/item"
    if172e05e5a1d156d04fd91e88341824ca485364aaa7f2ee21ded1804029cf1f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentscheduleinstances/item"
)

// DirectoryRequestBuilder provides operations to manage the directory property of the microsoft.graph.roleManagement entity.
type DirectoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRequestBuilderDeleteOptions options for Delete
type DirectoryRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRequestBuilderGetOptions options for Get
type DirectoryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRequestBuilderGetQueryParameters read-only. Nullable.
type DirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRequestBuilderPatchOptions options for Patch
type DirectoryRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property directory for roleManagement
func (m *DirectoryRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable.
func (m *DirectoryRequestBuilder) CreateGetRequestInformation(options *DirectoryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property directory in roleManagement
func (m *DirectoryRequestBuilder) CreatePatchRequestInformation(options *DirectoryRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property directory for roleManagement
func (m *DirectoryRequestBuilder) Delete(options *DirectoryRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable.
func (m *DirectoryRequestBuilder) Get(options *DirectoryRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateRbacApplicationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationable), nil
}
// Patch update the navigation property directory in roleManagement
func (m *DirectoryRequestBuilder) Patch(options *DirectoryRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DirectoryRequestBuilder) ResourceNamespaces()(*if7ad3988c943ab6c47fb727c18a8349d5e50ebec53ab6cc066c11a20717d46af.ResourceNamespacesRequestBuilder) {
    return if7ad3988c943ab6c47fb727c18a8349d5e50ebec53ab6cc066c11a20717d46af.NewResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceNamespacesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.resourceNamespaces.item collection
func (m *DirectoryRequestBuilder) ResourceNamespacesById(id string)(*i04dea7bcc2de46d7f933671fce512fa0699d5fc5d83576e03b7ac0fe57cc5704.UnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace_id"] = id
    }
    return i04dea7bcc2de46d7f933671fce512fa0699d5fc5d83576e03b7ac0fe57cc5704.NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleAssignmentApprovals()(*i4196e83a4bae24d0aeb92656129c7f451fb9b312487c6af458d556759a773f07.RoleAssignmentApprovalsRequestBuilder) {
    return i4196e83a4bae24d0aeb92656129c7f451fb9b312487c6af458d556759a773f07.NewRoleAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentApprovalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleAssignmentApprovals.item collection
func (m *DirectoryRequestBuilder) RoleAssignmentApprovalsById(id string)(*i44fbfcce41c9eaa7ba9ff5f6ba24f61e5a9c03cdb24c3bc25beed0e1d6216960.ApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval_id"] = id
    }
    return i44fbfcce41c9eaa7ba9ff5f6ba24f61e5a9c03cdb24c3bc25beed0e1d6216960.NewApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleAssignments()(*if3967cc171135f8ae89f635c0a197e13e64649a0043e1dcc96293c3c4f5ebbbe.RoleAssignmentsRequestBuilder) {
    return if3967cc171135f8ae89f635c0a197e13e64649a0043e1dcc96293c3c4f5ebbbe.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleAssignments.item collection
func (m *DirectoryRequestBuilder) RoleAssignmentsById(id string)(*i0dc801eaa7e34f0715fe9868f950fd4b68e0050fd9669d9d65ab28af2e0d7f73.UnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment_id"] = id
    }
    return i0dc801eaa7e34f0715fe9868f950fd4b68e0050fd9669d9d65ab28af2e0d7f73.NewUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleInstances()(*i955424cf34f796ced961a38c0bb409d0018463520cd1d610c06c4e0187f8f095.RoleAssignmentScheduleInstancesRequestBuilder) {
    return i955424cf34f796ced961a38c0bb409d0018463520cd1d610c06c4e0187f8f095.NewRoleAssignmentScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleAssignmentScheduleInstances.item collection
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleInstancesById(id string)(*if172e05e5a1d156d04fd91e88341824ca485364aaa7f2ee21ded1804029cf1f5.UnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance_id"] = id
    }
    return if172e05e5a1d156d04fd91e88341824ca485364aaa7f2ee21ded1804029cf1f5.NewUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleRequests()(*i003e247115fdc2cbce0f60f7d8a2c16dfec08ea21559d4270399d03acc2b4081.RoleAssignmentScheduleRequestsRequestBuilder) {
    return i003e247115fdc2cbce0f60f7d8a2c16dfec08ea21559d4270399d03acc2b4081.NewRoleAssignmentScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleAssignmentScheduleRequests.item collection
func (m *DirectoryRequestBuilder) RoleAssignmentScheduleRequestsById(id string)(*i1a1177924b99f8b8668ec4e5b307b5e9d537c2cfa54abdf2bf1fddde3af81044.UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleRequest_id"] = id
    }
    return i1a1177924b99f8b8668ec4e5b307b5e9d537c2cfa54abdf2bf1fddde3af81044.NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleAssignmentSchedules()(*ie88f7587d0b612d6aad4a4cc6e4adf2df6596695dcfb937bd88d739a30fec4d4.RoleAssignmentSchedulesRequestBuilder) {
    return ie88f7587d0b612d6aad4a4cc6e4adf2df6596695dcfb937bd88d739a30fec4d4.NewRoleAssignmentSchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentSchedulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleAssignmentSchedules.item collection
func (m *DirectoryRequestBuilder) RoleAssignmentSchedulesById(id string)(*i5e4e5e8e6ab57f21346122383935dbe84160f647fde997fe85fe2ae4a427c50e.UnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentSchedule_id"] = id
    }
    return i5e4e5e8e6ab57f21346122383935dbe84160f647fde997fe85fe2ae4a427c50e.NewUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleDefinitions()(*ic09a099122154ece987c7f85636fcc11aac7d11f4704eb7ea18e15d32cb118be.RoleDefinitionsRequestBuilder) {
    return ic09a099122154ece987c7f85636fcc11aac7d11f4704eb7ea18e15d32cb118be.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleDefinitions.item collection
func (m *DirectoryRequestBuilder) RoleDefinitionsById(id string)(*i15fef4d98a25fc03a5a29e7e66cec2ffeccb757e532a20aa82098881625f052a.UnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition_id"] = id
    }
    return i15fef4d98a25fc03a5a29e7e66cec2ffeccb757e532a20aa82098881625f052a.NewUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleEligibilityScheduleInstances()(*ie00125ef63b154ab9fd6cbeba46882ef7ecf0a1bbe044d47b4e5fd94f87ff817.RoleEligibilityScheduleInstancesRequestBuilder) {
    return ie00125ef63b154ab9fd6cbeba46882ef7ecf0a1bbe044d47b4e5fd94f87ff817.NewRoleEligibilityScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleEligibilityScheduleInstances.item collection
func (m *DirectoryRequestBuilder) RoleEligibilityScheduleInstancesById(id string)(*i322381699c08d5273b2d2e8a72a478e9a4b2c80cb977ad926aa502c1353b499a.UnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleInstance_id"] = id
    }
    return i322381699c08d5273b2d2e8a72a478e9a4b2c80cb977ad926aa502c1353b499a.NewUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleEligibilityScheduleRequests()(*i99441748b67ef831e7259b916e604bc0f3a134120b98e3f3cec7689707945a5b.RoleEligibilityScheduleRequestsRequestBuilder) {
    return i99441748b67ef831e7259b916e604bc0f3a134120b98e3f3cec7689707945a5b.NewRoleEligibilityScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleEligibilityScheduleRequests.item collection
func (m *DirectoryRequestBuilder) RoleEligibilityScheduleRequestsById(id string)(*id0ebe1fb3bbbfdf8b2d57215ab0ec7662faabe46ca681bc46803c9c70d11cab5.UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest_id"] = id
    }
    return id0ebe1fb3bbbfdf8b2d57215ab0ec7662faabe46ca681bc46803c9c70d11cab5.NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DirectoryRequestBuilder) RoleEligibilitySchedules()(*ib3f3f9905e601fb21a0927cffd109f585906d82dd65ac7b379b6df162c0beec5.RoleEligibilitySchedulesRequestBuilder) {
    return ib3f3f9905e601fb21a0927cffd109f585906d82dd65ac7b379b6df162c0beec5.NewRoleEligibilitySchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilitySchedulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.roleEligibilitySchedules.item collection
func (m *DirectoryRequestBuilder) RoleEligibilitySchedulesById(id string)(*i2b3fa3ea5b323a0c2dca9d12e3d5f7bdadbc4bbf1e526ea51f794a5d0f803d75.UnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilitySchedule_id"] = id
    }
    return i2b3fa3ea5b323a0c2dca9d12e3d5f7bdadbc4bbf1e526ea51f794a5d0f803d75.NewUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId provides operations to call the roleScheduleInstances method.
func (m *DirectoryRequestBuilder) RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId(directoryScopeId *string, principalId *string, roleDefinitionId *string, appScopeId *string)(*ib7ee0af232e3f7f843b9fc80b2bf1bcd6bdc1aa1876b6682453a23c520194820.RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) {
    return ib7ee0af232e3f7f843b9fc80b2bf1bcd6bdc1aa1876b6682453a23c520194820.NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, directoryScopeId, principalId, roleDefinitionId, appScopeId);
}
// RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId provides operations to call the roleSchedules method.
func (m *DirectoryRequestBuilder) RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId(directoryScopeId *string, principalId *string, roleDefinitionId *string, appScopeId *string)(*i89936fb8dab16eb5ef6964ec9e5d8aa7c54f3f69934f902cdc4a2a400db45270.RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) {
    return i89936fb8dab16eb5ef6964ec9e5d8aa7c54f3f69934f902cdc4a2a400db45270.NewRoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, directoryScopeId, principalId, roleDefinitionId, appScopeId);
}
func (m *DirectoryRequestBuilder) TransitiveRoleAssignments()(*i1674488bab3f723cd92320bded03455762b2d10b76d65717248d2a39b768949d.TransitiveRoleAssignmentsRequestBuilder) {
    return i1674488bab3f723cd92320bded03455762b2d10b76d65717248d2a39b768949d.NewTransitiveRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.transitiveRoleAssignments.item collection
func (m *DirectoryRequestBuilder) TransitiveRoleAssignmentsById(id string)(*id3919a2c832ecf67620541df52d40952233a8c3c7f1293166f41f2e20a0f3aa3.UnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment_id"] = id
    }
    return id3919a2c832ecf67620541df52d40952233a8c3c7f1293166f41f2e20a0f3aa3.NewUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
