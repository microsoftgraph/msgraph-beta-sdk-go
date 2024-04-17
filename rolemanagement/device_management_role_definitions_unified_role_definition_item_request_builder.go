package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplicationMultiple entity.
type DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters get the properties and relationships of a unifiedRoleDefinition object of an RBAC provider.  The following RBAC providers are currently supported:- Cloud PC - device management (Intune)- directory (Microsoft Entra directory roles)- entitlement management (Microsoft Entra entitlement management)- Exchange Online
type DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters
}
// DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeId provides operations to call the assignedPrincipals method.
// returns a *DeviceManagementRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder when successful
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) AssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeId()(*DeviceManagementRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) {
    return NewDeviceManagementRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal instantiates a new DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewDeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    m := &DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/deviceManagement/roleDefinitions/{unifiedRoleDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder instantiates a new DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewDeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a unifiedRoleDefinition object for an RBAC provider. You cannot delete built-in roles. This feature requires a Microsoft Entra ID P1 or P2 license. The following RBAC providers are currently supported:- Cloud PC- device management (Intune)- directory (Microsoft Entra ID) 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroledefinition-delete?view=graph-rest-1.0
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties and relationships of a unifiedRoleDefinition object of an RBAC provider.  The following RBAC providers are currently supported:- Cloud PC - device management (Intune)- directory (Microsoft Entra directory roles)- entitlement management (Microsoft Entra entitlement management)- Exchange Online
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroledefinition-get?view=graph-rest-1.0
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
// InheritsPermissionsFrom provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
// returns a *DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder when successful
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) InheritsPermissionsFrom()(*DeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) {
    return NewDeviceManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a unifiedRoleDefinition object for an RBAC provider. You cannot update built-in roles. This feature requires a Microsoft Entra ID P1 or P2 license. The following RBAC providers are currently supported:- Cloud PC- device management (Intune)- directory (Microsoft Entra ID) 
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroledefinition-update?view=graph-rest-1.0
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
// ToDeleteRequestInformation delete a unifiedRoleDefinition object for an RBAC provider. You cannot delete built-in roles. This feature requires a Microsoft Entra ID P1 or P2 license. The following RBAC providers are currently supported:- Cloud PC- device management (Intune)- directory (Microsoft Entra ID) 
// returns a *RequestInformation when successful
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of a unifiedRoleDefinition object of an RBAC provider.  The following RBAC providers are currently supported:- Cloud PC - device management (Intune)- directory (Microsoft Entra directory roles)- entitlement management (Microsoft Entra entitlement management)- Exchange Online
// returns a *RequestInformation when successful
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a unifiedRoleDefinition object for an RBAC provider. You cannot update built-in roles. This feature requires a Microsoft Entra ID P1 or P2 license. The following RBAC providers are currently supported:- Cloud PC- device management (Intune)- directory (Microsoft Entra ID) 
// returns a *RequestInformation when successful
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder when successful
func (m *DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    return NewDeviceManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
