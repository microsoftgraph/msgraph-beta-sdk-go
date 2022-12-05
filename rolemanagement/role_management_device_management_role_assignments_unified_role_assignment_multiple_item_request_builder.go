package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplicationMultiple entity.
type RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters get roleAssignments from roleManagement
type RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters
}
// RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScopes provides operations to manage the appScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopes()(*RoleManagementDeviceManagementRoleAssignmentsItemAppScopesRequestBuilder) {
    return NewRoleManagementDeviceManagementRoleAssignmentsItemAppScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppScopesById provides operations to manage the appScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopesById(id string)(*RoleManagementDeviceManagementRoleAssignmentsItemAppScopesAppScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appScope%2Did"] = id
    }
    return NewRoleManagementDeviceManagementRoleAssignmentsItemAppScopesAppScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewRoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    m := &RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder instantiates a new UnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewRoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for roleManagement
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get roleAssignments from roleManagement
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignments in roleManagement
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, requestConfiguration *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleAssignments for roleManagement
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScopes provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopes()(*RoleManagementDeviceManagementRoleAssignmentsItemDirectoryScopesRequestBuilder) {
    return NewRoleManagementDeviceManagementRoleAssignmentsItemDirectoryScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryScopesById provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopesById(id string)(*RoleManagementDeviceManagementRoleAssignmentsItemDirectoryScopesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewRoleManagementDeviceManagementRoleAssignmentsItemDirectoryScopesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get roleAssignments from roleManagement
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable), nil
}
// Patch update the navigation property roleAssignments in roleManagement
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, requestConfiguration *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable), nil
}
// Principals provides operations to manage the principals property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Principals()(*RoleManagementDeviceManagementRoleAssignmentsItemPrincipalsRequestBuilder) {
    return NewRoleManagementDeviceManagementRoleAssignmentsItemPrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrincipalsById provides operations to manage the principals property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) PrincipalsById(id string)(*RoleManagementDeviceManagementRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewRoleManagementDeviceManagementRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementDeviceManagementRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) RoleDefinition()(*RoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    return NewRoleManagementDeviceManagementRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
