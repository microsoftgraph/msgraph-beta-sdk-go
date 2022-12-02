package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
type RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters get roleDefinitions from roleManagement
type RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters
}
// RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeId provides operations to call the assignedPrincipals method.
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) AssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeId()(*RoleManagementEntitlementManagementRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal instantiates a new UnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    m := &RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleDefinitions/{unifiedRoleDefinition%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder instantiates a new UnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleDefinitions for roleManagement
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get roleDefinitions from roleManagement
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleDefinitions in roleManagement
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleDefinitions for roleManagement
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get roleDefinitions from roleManagement
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
// InheritsPermissionsFrom provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) InheritsPermissionsFrom()(*RoleManagementEntitlementManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) {
    return NewRoleManagementEntitlementManagementRoleDefinitionsItemInheritsPermissionsFromRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InheritsPermissionsFromById provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) InheritsPermissionsFromById(id string)(*RoleManagementEntitlementManagementRoleDefinitionsItemInheritsPermissionsFromUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition%2Did1"] = id
    }
    return NewRoleManagementEntitlementManagementRoleDefinitionsItemInheritsPermissionsFromUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property roleDefinitions in roleManagement
func (m *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *RoleManagementEntitlementManagementRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
