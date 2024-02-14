package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
type EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters get the instance of an active role assignment.
type EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters
}
// EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentScheduleInstance entity.
// returns a *EntitlementManagementRoleAssignmentScheduleInstancesItemActivatedUsingRequestBuilder when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ActivatedUsing()(*EntitlementManagementRoleAssignmentScheduleInstancesItemActivatedUsingRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleInstancesItemActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *EntitlementManagementRoleAssignmentScheduleInstancesItemAppScopeRequestBuilder when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) AppScope()(*EntitlementManagementRoleAssignmentScheduleInstancesItemAppScopeRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleInstancesItemAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal instantiates a new EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewEntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    m := &EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleInstances/{unifiedRoleAssignmentScheduleInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder instantiates a new EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewEntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentScheduleInstances for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *EntitlementManagementRoleAssignmentScheduleInstancesItemDirectoryScopeRequestBuilder when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) DirectoryScope()(*EntitlementManagementRoleAssignmentScheduleInstancesItemDirectoryScopeRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleInstancesItemDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the instance of an active role assignment.
// returns a UnifiedRoleAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentscheduleinstance-get?view=graph-rest-1.0
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleInstanceable), nil
}
// Patch update the navigation property roleAssignmentScheduleInstances in roleManagement
// returns a UnifiedRoleAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleInstanceable, requestConfiguration *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleInstanceable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *EntitlementManagementRoleAssignmentScheduleInstancesItemPrincipalRequestBuilder when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Principal()(*EntitlementManagementRoleAssignmentScheduleInstancesItemPrincipalRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleInstancesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *EntitlementManagementRoleAssignmentScheduleInstancesItemRoleDefinitionRequestBuilder when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) RoleDefinition()(*EntitlementManagementRoleAssignmentScheduleInstancesItemRoleDefinitionRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleInstancesItemRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentScheduleInstances for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleInstances/{unifiedRoleAssignmentScheduleInstance%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the instance of an active role assignment.
// returns a *RequestInformation when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property roleAssignmentScheduleInstances in roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleInstanceable, requestConfiguration *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleInstances/{unifiedRoleAssignmentScheduleInstance%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder when successful
func (m *EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    return NewEntitlementManagementRoleAssignmentScheduleInstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
