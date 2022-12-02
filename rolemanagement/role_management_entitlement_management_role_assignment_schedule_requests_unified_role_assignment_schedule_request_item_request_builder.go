package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters get roleAssignmentScheduleRequests from roleManagement
type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters
}
// RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemActivatedUsingRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) AppScope()(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemAppScopeRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Cancel()(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemCancelRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    m := &RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get roleAssignmentScheduleRequests from roleManagement
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DirectoryScope()(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemDirectoryScopeRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignmentScheduleRequests from roleManagement
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable), nil
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Principal()(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemPrincipalRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) RoleDefinition()(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemRoleDefinitionRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
func (m *RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*RoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestsItemTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
