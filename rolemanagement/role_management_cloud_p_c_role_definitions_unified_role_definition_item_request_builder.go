package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplicationMultiple entity.
type RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters get roleDefinitions from roleManagement
type RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters
}
// RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeId provides operations to call the assignedPrincipals method.
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) AssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeId()(*RoleManagementCloudPCRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) {
    return NewRoleManagementCloudPCRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal instantiates a new UnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewRoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    m := &RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/cloudPC/roleDefinitions/{unifiedRoleDefinition%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder instantiates a new UnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewRoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleDefinitions for roleManagement
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
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
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) InheritsPermissionsFrom()(*RoleManagementCloudPCRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) {
    return NewRoleManagementCloudPCRoleDefinitionsItemInheritsPermissionsFromRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InheritsPermissionsFromById provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) InheritsPermissionsFromById(id string)(*RoleManagementCloudPCRoleDefinitionsItemInheritsPermissionsFromUnifiedRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition%2Did1"] = id
    }
    return NewRoleManagementCloudPCRoleDefinitionsItemInheritsPermissionsFromUnifiedRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property roleDefinitions in roleManagement
func (m *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *RoleManagementCloudPCRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
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
