// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplicationMultiple entity.
type DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters get the properties and relationships of a unifiedRoleAssignmentMultiple object of an RBAC provider.  The following RBAC providers are currently supported:- Cloud PC - device management (Intune)- Defender (Microsoft Defender XDR Unified RBAC) For other Microsoft 365 applications (like Microsoft Entra ID), use unifiedRoleAssignment.
type DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters
}
// DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScopes provides operations to manage the appScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DeviceManagementRoleAssignmentsItemAppScopesRequestBuilder when successful
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopes()(*DeviceManagementRoleAssignmentsItemAppScopesRequestBuilder) {
    return NewDeviceManagementRoleAssignmentsItemAppScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal instantiates a new DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    m := &DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/deviceManagement/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder instantiates a new DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a unifiedRoleAssignmentMultiple object of an RBAC provider.  This is applicable for a RBAC application that supports multiple principals and scopes. The following RBAC providers are currently supported:- Cloud PC - device management (Intune)- Defender (Microsoft Defender XDR Unified RBAC)
// Deprecated: This version is being deprecated and is scheduled for removal on 2025-12-01.Please migrate to the latest version before the removal date. as of 2025-01/PrivatePreview:microsoft.applicationAuthorization on 2025-01-01 and will be removed 2025-12-01
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentmultiple-delete?view=graph-rest-beta
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScopes provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DeviceManagementRoleAssignmentsItemDirectoryScopesRequestBuilder when successful
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopes()(*DeviceManagementRoleAssignmentsItemDirectoryScopesRequestBuilder) {
    return NewDeviceManagementRoleAssignmentsItemDirectoryScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the properties and relationships of a unifiedRoleAssignmentMultiple object of an RBAC provider.  The following RBAC providers are currently supported:- Cloud PC - device management (Intune)- Defender (Microsoft Defender XDR Unified RBAC) For other Microsoft 365 applications (like Microsoft Entra ID), use unifiedRoleAssignment.
// Deprecated: This version is being deprecated and is scheduled for removal on 2025-12-01.Please migrate to the latest version before the removal date. as of 2025-01/PrivatePreview:microsoft.applicationAuthorization on 2025-01-01 and will be removed 2025-12-01
// returns a UnifiedRoleAssignmentMultipleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentmultiple-get?view=graph-rest-beta
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable), nil
}
// Patch update an existing unifiedRoleAssignmentMultiple object of an RBAC provider.  The following RBAC providers are currently supported:- Cloud PC - device management (Intune) In contrast, unifiedRoleAssignment does not support update.
// Deprecated: This version is being deprecated and is scheduled for removal on 2025-12-01.Please migrate to the latest version before the removal date. as of 2025-01/PrivatePreview:microsoft.applicationAuthorization on 2025-01-01 and will be removed 2025-12-01
// returns a UnifiedRoleAssignmentMultipleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentmultiple-update?view=graph-rest-beta
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, requestConfiguration *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable), nil
}
// Principals provides operations to manage the principals property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DeviceManagementRoleAssignmentsItemPrincipalsRequestBuilder when successful
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Principals()(*DeviceManagementRoleAssignmentsItemPrincipalsRequestBuilder) {
    return NewDeviceManagementRoleAssignmentsItemPrincipalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder when successful
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) RoleDefinition()(*DeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    return NewDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a unifiedRoleAssignmentMultiple object of an RBAC provider.  This is applicable for a RBAC application that supports multiple principals and scopes. The following RBAC providers are currently supported:- Cloud PC - device management (Intune)- Defender (Microsoft Defender XDR Unified RBAC)
// Deprecated: This version is being deprecated and is scheduled for removal on 2025-12-01.Please migrate to the latest version before the removal date. as of 2025-01/PrivatePreview:microsoft.applicationAuthorization on 2025-01-01 and will be removed 2025-12-01
// returns a *RequestInformation when successful
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of a unifiedRoleAssignmentMultiple object of an RBAC provider.  The following RBAC providers are currently supported:- Cloud PC - device management (Intune)- Defender (Microsoft Defender XDR Unified RBAC) For other Microsoft 365 applications (like Microsoft Entra ID), use unifiedRoleAssignment.
// Deprecated: This version is being deprecated and is scheduled for removal on 2025-12-01.Please migrate to the latest version before the removal date. as of 2025-01/PrivatePreview:microsoft.applicationAuthorization on 2025-01-01 and will be removed 2025-12-01
// returns a *RequestInformation when successful
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update an existing unifiedRoleAssignmentMultiple object of an RBAC provider.  The following RBAC providers are currently supported:- Cloud PC - device management (Intune) In contrast, unifiedRoleAssignment does not support update.
// Deprecated: This version is being deprecated and is scheduled for removal on 2025-12-01.Please migrate to the latest version before the removal date. as of 2025-01/PrivatePreview:microsoft.applicationAuthorization on 2025-01-01 and will be removed 2025-12-01
// returns a *RequestInformation when successful
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, requestConfiguration *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: This version is being deprecated and is scheduled for removal on 2025-12-01.Please migrate to the latest version before the removal date. as of 2025-01/PrivatePreview:microsoft.applicationAuthorization on 2025-01-01 and will be removed 2025-12-01
// returns a *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder when successful
func (m *DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    return NewDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
