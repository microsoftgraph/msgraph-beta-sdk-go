package directory

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i003e247115fdc2cbce0f60f7d8a2c16dfec08ea21559d4270399d03acc2b4081 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests"
    i1674488bab3f723cd92320bded03455762b2d10b76d65717248d2a39b768949d "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/transitiveroleassignments"
    i4196e83a4bae24d0aeb92656129c7f451fb9b312487c6af458d556759a773f07 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentapprovals"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRequestBuilderGetQueryParameters get directory from roleManagement
type DirectoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRequestBuilderGetQueryParameters
}
// DirectoryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRequestBuilderInternal instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    m := &DirectoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRequestBuilder instantiates a new DirectoryRequestBuilder and sets the default values.
func NewDirectoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property directory for roleManagement
func (m *DirectoryRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property directory for roleManagement
func (m *DirectoryRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get directory from roleManagement
func (m *DirectoryRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get directory from roleManagement
func (m *DirectoryRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property directory in roleManagement
func (m *DirectoryRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property directory in roleManagement
func (m *DirectoryRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *DirectoryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property directory for roleManagement
func (m *DirectoryRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property directory for roleManagement
func (m *DirectoryRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DirectoryRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get directory from roleManagement
func (m *DirectoryRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get directory from roleManagement
func (m *DirectoryRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DirectoryRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable), nil
}
// Patch update the navigation property directory in roleManagement
func (m *DirectoryRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property directory in roleManagement
func (m *DirectoryRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *DirectoryRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ResourceNamespaces the resourceNamespaces property
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
        urlTplParams["unifiedRbacResourceNamespace%2Did"] = id
    }
    return i04dea7bcc2de46d7f933671fce512fa0699d5fc5d83576e03b7ac0fe57cc5704.NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentApprovals the roleAssignmentApprovals property
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
        urlTplParams["approval%2Did"] = id
    }
    return i44fbfcce41c9eaa7ba9ff5f6ba24f61e5a9c03cdb24c3bc25beed0e1d6216960.NewApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments the roleAssignments property
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
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return i0dc801eaa7e34f0715fe9868f950fd4b68e0050fd9669d9d65ab28af2e0d7f73.NewUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleInstances the roleAssignmentScheduleInstances property
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
        urlTplParams["unifiedRoleAssignmentScheduleInstance%2Did"] = id
    }
    return if172e05e5a1d156d04fd91e88341824ca485364aaa7f2ee21ded1804029cf1f5.NewUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleRequests the roleAssignmentScheduleRequests property
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
        urlTplParams["unifiedRoleAssignmentScheduleRequest%2Did"] = id
    }
    return i1a1177924b99f8b8668ec4e5b307b5e9d537c2cfa54abdf2bf1fddde3af81044.NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentSchedules the roleAssignmentSchedules property
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
        urlTplParams["unifiedRoleAssignmentSchedule%2Did"] = id
    }
    return i5e4e5e8e6ab57f21346122383935dbe84160f647fde997fe85fe2ae4a427c50e.NewUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions the roleDefinitions property
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
        urlTplParams["unifiedRoleDefinition%2Did"] = id
    }
    return i15fef4d98a25fc03a5a29e7e66cec2ffeccb757e532a20aa82098881625f052a.NewUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleInstances the roleEligibilityScheduleInstances property
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
        urlTplParams["unifiedRoleEligibilityScheduleInstance%2Did"] = id
    }
    return i322381699c08d5273b2d2e8a72a478e9a4b2c80cb977ad926aa502c1353b499a.NewUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleRequests the roleEligibilityScheduleRequests property
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
        urlTplParams["unifiedRoleEligibilityScheduleRequest%2Did"] = id
    }
    return id0ebe1fb3bbbfdf8b2d57215ab0ec7662faabe46ca681bc46803c9c70d11cab5.NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilitySchedules the roleEligibilitySchedules property
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
        urlTplParams["unifiedRoleEligibilitySchedule%2Did"] = id
    }
    return i2b3fa3ea5b323a0c2dca9d12e3d5f7bdadbc4bbf1e526ea51f794a5d0f803d75.NewUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId provides operations to call the roleScheduleInstances method.
func (m *DirectoryRequestBuilder) RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId()(*ib7ee0af232e3f7f843b9fc80b2bf1bcd6bdc1aa1876b6682453a23c520194820.RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) {
    return ib7ee0af232e3f7f843b9fc80b2bf1bcd6bdc1aa1876b6682453a23c520194820.NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId provides operations to call the roleSchedules method.
func (m *DirectoryRequestBuilder) RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId()(*i89936fb8dab16eb5ef6964ec9e5d8aa7c54f3f69934f902cdc4a2a400db45270.RoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) {
    return i89936fb8dab16eb5ef6964ec9e5d8aa7c54f3f69934f902cdc4a2a400db45270.NewRoleSchedulesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveRoleAssignments the transitiveRoleAssignments property
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
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return id3919a2c832ecf67620541df52d40952233a8c3c7f1293166f41f2e20a0f3aa3.NewUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
