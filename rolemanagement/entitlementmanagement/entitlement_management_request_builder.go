package entitlementmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11fc614d278bf1739fd916a26cc01c0f55b7e6b82f3788fd0a49e4d3d058b2c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleschedulesdirectoryscopeiddirectoryscopeidappscopeidappscopeidprincipalidprincipalidroledefinitionidroledefinitionid"
    i1baa975bee8bc1fcf62ee7c7d05b837d65d8d022e7556de9e284932ea41c5624 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedules"
    i1f9f0f605bf919aaf8802400f663bbfe35dc4f923ea29fc6cb19acd46f7d7e52 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/rolescheduleinstancesdirectoryscopeiddirectoryscopeidappscopeidappscopeidprincipalidprincipalidroledefinitionidroledefinitionid"
    i3b3dd9d3bb14498220a703a0b62e7d53250e7202ee87f41c8d4aff4c9404d695 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/transitiveroleassignments"
    i5e09c2b6a4bc955be405d7fbb453204f96818f8aae7508864d080577c5491cd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedulerequests"
    i75ab2bc44492a09208082e3507fb5b1add85563cc4e6f15226038f431107cc36 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/resourcenamespaces"
    i8efb8a53874fa360f1b73def8f68e4e407323fb5958457abad962b8900d25355 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests"
    ia143a4210940ae95c6f24da8e1def30da0952553539019a9c5e037acaae6d536 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedules"
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
    ic784fd081f4d06fa9b67547dc793f66881addca2cbbf91d449ea61a1f7e5476e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/transitiveroleassignments/item"
    icf1ccc36301949ddd8b0254d07f44b375f6f95d7adb58a22b17ccfc55ae25abc "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityscheduleinstances/item"
    id9a0f06cc5af154fe717e7b85c2d5db50242eb2a5eb5b84cd9b6f2bfff25c1f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleeligibilityschedules/item"
    ie58dbe4c0951e87ffefa6af6faabb1e4ad0651b2ea404363bfd8a2d38cc97a53 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/resourcenamespaces/item"
)

// EntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.roleManagement entity.
type EntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementRequestBuilderGetQueryParameters the RbacApplication for Entitlement Management
type EntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementRequestBuilderGetQueryParameters
}
// EntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    m := &EntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property entitlementManagement for roleManagement
func (m *EntitlementManagementRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property entitlementManagement for roleManagement
func (m *EntitlementManagementRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property entitlementManagement in roleManagement
func (m *EntitlementManagementRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property entitlementManagement in roleManagement
func (m *EntitlementManagementRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *EntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property entitlementManagement for roleManagement
func (m *EntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the RbacApplication for Entitlement Management
func (m *EntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable), nil
}
// Patch update the navigation property entitlementManagement in roleManagement
func (m *EntitlementManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, requestConfiguration *EntitlementManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RbacApplicationable), nil
}
// ResourceNamespaces the resourceNamespaces property
func (m *EntitlementManagementRequestBuilder) ResourceNamespaces()(*i75ab2bc44492a09208082e3507fb5b1add85563cc4e6f15226038f431107cc36.ResourceNamespacesRequestBuilder) {
    return i75ab2bc44492a09208082e3507fb5b1add85563cc4e6f15226038f431107cc36.NewResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceNamespacesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.resourceNamespaces.item collection
func (m *EntitlementManagementRequestBuilder) ResourceNamespacesById(id string)(*ie58dbe4c0951e87ffefa6af6faabb1e4ad0651b2ea404363bfd8a2d38cc97a53.UnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace%2Did"] = id
    }
    return ie58dbe4c0951e87ffefa6af6faabb1e4ad0651b2ea404363bfd8a2d38cc97a53.NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentApprovals the roleAssignmentApprovals property
func (m *EntitlementManagementRequestBuilder) RoleAssignmentApprovals()(*idaef4275a9c6a8af3da51ed597b3772552b76155312fa2174e893913ec186b9e.RoleAssignmentApprovalsRequestBuilder) {
    return idaef4275a9c6a8af3da51ed597b3772552b76155312fa2174e893913ec186b9e.NewRoleAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentApprovalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignmentApprovals.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentApprovalsById(id string)(*i94f6c4184f5f2637cb72ccacd72d05ca460673ced94157d6b941e79b19f921e8.ApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return i94f6c4184f5f2637cb72ccacd72d05ca460673ced94157d6b941e79b19f921e8.NewApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments the roleAssignments property
func (m *EntitlementManagementRequestBuilder) RoleAssignments()(*if94439b0eb19f66d0fe98bc2e50b7664597fad69cfac1e1132b6b13d328ae667.RoleAssignmentsRequestBuilder) {
    return if94439b0eb19f66d0fe98bc2e50b7664597fad69cfac1e1132b6b13d328ae667.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignments.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentsById(id string)(*i2e767808e39342c5e075a8c93310b2c0591c0c806546f5b58190d82a94993717.UnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return i2e767808e39342c5e075a8c93310b2c0591c0c806546f5b58190d82a94993717.NewUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleInstances the roleAssignmentScheduleInstances property
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleInstances()(*ie577e072e7edde2d3f423e1ceb296400868e7f5e7bddbe2ee934c71bd8d1d637.RoleAssignmentScheduleInstancesRequestBuilder) {
    return ie577e072e7edde2d3f423e1ceb296400868e7f5e7bddbe2ee934c71bd8d1d637.NewRoleAssignmentScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignmentScheduleInstances.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleInstancesById(id string)(*i6f70cfe1f611776ca2d5cdc9b36b624ba1b69dda981099ccaaf42d45453807a7.UnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleInstance%2Did"] = id
    }
    return i6f70cfe1f611776ca2d5cdc9b36b624ba1b69dda981099ccaaf42d45453807a7.NewUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentScheduleRequests the roleAssignmentScheduleRequests property
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleRequests()(*i8efb8a53874fa360f1b73def8f68e4e407323fb5958457abad962b8900d25355.RoleAssignmentScheduleRequestsRequestBuilder) {
    return i8efb8a53874fa360f1b73def8f68e4e407323fb5958457abad962b8900d25355.NewRoleAssignmentScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentScheduleRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignmentScheduleRequests.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentScheduleRequestsById(id string)(*i1094763ed645fa459b11fe2fd0dbbb4cd5b4e5abf8fb494633f130f88a198ab7.UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentScheduleRequest%2Did"] = id
    }
    return i1094763ed645fa459b11fe2fd0dbbb4cd5b4e5abf8fb494633f130f88a198ab7.NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentSchedules the roleAssignmentSchedules property
func (m *EntitlementManagementRequestBuilder) RoleAssignmentSchedules()(*i1baa975bee8bc1fcf62ee7c7d05b837d65d8d022e7556de9e284932ea41c5624.RoleAssignmentSchedulesRequestBuilder) {
    return i1baa975bee8bc1fcf62ee7c7d05b837d65d8d022e7556de9e284932ea41c5624.NewRoleAssignmentSchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentSchedulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleAssignmentSchedules.item collection
func (m *EntitlementManagementRequestBuilder) RoleAssignmentSchedulesById(id string)(*ib97806c66498b497946ee83f7bb51d300c1952475468eabbe9d487576ab31775.UnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentSchedule%2Did"] = id
    }
    return ib97806c66498b497946ee83f7bb51d300c1952475468eabbe9d487576ab31775.NewUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions the roleDefinitions property
func (m *EntitlementManagementRequestBuilder) RoleDefinitions()(*icfcbd750b986c89f05029c8a1c38ed5d1e850b098c1de07332c9c8f4dba9e621.RoleDefinitionsRequestBuilder) {
    return icfcbd750b986c89f05029c8a1c38ed5d1e850b098c1de07332c9c8f4dba9e621.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleDefinitions.item collection
func (m *EntitlementManagementRequestBuilder) RoleDefinitionsById(id string)(*ib2c51c75dd96d44353372f4ec3d1bd31e89c02a24c89600c2d3fb288758d1bf3.UnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition%2Did"] = id
    }
    return ib2c51c75dd96d44353372f4ec3d1bd31e89c02a24c89600c2d3fb288758d1bf3.NewUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleInstances the roleEligibilityScheduleInstances property
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleInstances()(*ic74edb001b0ce5f97639bc39b485bc37fc8d4b308c0e093fb317822e211e74ab.RoleEligibilityScheduleInstancesRequestBuilder) {
    return ic74edb001b0ce5f97639bc39b485bc37fc8d4b308c0e093fb317822e211e74ab.NewRoleEligibilityScheduleInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleEligibilityScheduleInstances.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleInstancesById(id string)(*icf1ccc36301949ddd8b0254d07f44b375f6f95d7adb58a22b17ccfc55ae25abc.UnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleInstance%2Did"] = id
    }
    return icf1ccc36301949ddd8b0254d07f44b375f6f95d7adb58a22b17ccfc55ae25abc.NewUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilityScheduleRequests the roleEligibilityScheduleRequests property
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleRequests()(*i5e09c2b6a4bc955be405d7fbb453204f96818f8aae7508864d080577c5491cd0.RoleEligibilityScheduleRequestsRequestBuilder) {
    return i5e09c2b6a4bc955be405d7fbb453204f96818f8aae7508864d080577c5491cd0.NewRoleEligibilityScheduleRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilityScheduleRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleEligibilityScheduleRequests.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilityScheduleRequestsById(id string)(*i1f3fe61534b224c6fac2a1c5c69508f0db3518d4530ee88be6548eae659bb75a.UnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilityScheduleRequest%2Did"] = id
    }
    return i1f3fe61534b224c6fac2a1c5c69508f0db3518d4530ee88be6548eae659bb75a.NewUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleEligibilitySchedules the roleEligibilitySchedules property
func (m *EntitlementManagementRequestBuilder) RoleEligibilitySchedules()(*ia143a4210940ae95c6f24da8e1def30da0952553539019a9c5e037acaae6d536.RoleEligibilitySchedulesRequestBuilder) {
    return ia143a4210940ae95c6f24da8e1def30da0952553539019a9c5e037acaae6d536.NewRoleEligibilitySchedulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleEligibilitySchedulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.roleEligibilitySchedules.item collection
func (m *EntitlementManagementRequestBuilder) RoleEligibilitySchedulesById(id string)(*id9a0f06cc5af154fe717e7b85c2d5db50242eb2a5eb5b84cd9b6f2bfff25c1f4.UnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleEligibilitySchedule%2Did"] = id
    }
    return id9a0f06cc5af154fe717e7b85c2d5db50242eb2a5eb5b84cd9b6f2bfff25c1f4.NewUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId provides operations to call the roleScheduleInstances method.
func (m *EntitlementManagementRequestBuilder) RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId()(*i1f9f0f605bf919aaf8802400f663bbfe35dc4f923ea29fc6cb19acd46f7d7e52.RoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return i1f9f0f605bf919aaf8802400f663bbfe35dc4f923ea29fc6cb19acd46f7d7e52.NewRoleScheduleInstancesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId provides operations to call the roleSchedules method.
func (m *EntitlementManagementRequestBuilder) RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionId()(*i11fc614d278bf1739fd916a26cc01c0f55b7e6b82f3788fd0a49e4d3d058b2c8.RoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilder) {
    return i11fc614d278bf1739fd916a26cc01c0f55b7e6b82f3788fd0a49e4d3d058b2c8.NewRoleSchedulesdirectoryScopeIdDirectoryScopeIdAppScopeIdAppScopeIdPrincipalIdPrincipalIdRoleDefinitionIdRoleDefinitionIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveRoleAssignments the transitiveRoleAssignments property
func (m *EntitlementManagementRequestBuilder) TransitiveRoleAssignments()(*i3b3dd9d3bb14498220a703a0b62e7d53250e7202ee87f41c8d4aff4c9404d695.TransitiveRoleAssignmentsRequestBuilder) {
    return i3b3dd9d3bb14498220a703a0b62e7d53250e7202ee87f41c8d4aff4c9404d695.NewTransitiveRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.entitlementManagement.transitiveRoleAssignments.item collection
func (m *EntitlementManagementRequestBuilder) TransitiveRoleAssignmentsById(id string)(*ic784fd081f4d06fa9b67547dc793f66881addca2cbbf91d449ea61a1f7e5476e.UnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = id
    }
    return ic784fd081f4d06fa9b67547dc793f66881addca2cbbf91d449ea61a1f7e5476e.NewUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
