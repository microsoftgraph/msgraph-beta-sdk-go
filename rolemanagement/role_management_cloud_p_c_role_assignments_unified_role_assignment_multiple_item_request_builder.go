package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplicationMultiple entity.
type RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters get roleAssignments from roleManagement
type RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters
}
// RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScopes provides operations to manage the appScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopes()(*RoleManagementCloudPCRoleAssignmentsItemAppScopesRequestBuilder) {
    return NewRoleManagementCloudPCRoleAssignmentsItemAppScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppScopesById provides operations to manage the appScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopesById(id string)(*RoleManagementCloudPCRoleAssignmentsItemAppScopesAppScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appScope%2Did"] = id
    }
    return NewRoleManagementCloudPCRoleAssignmentsItemAppScopesAppScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewRoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    m := &RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/cloudPC/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder instantiates a new UnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewRoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for roleManagement
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, requestConfiguration *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopes()(*RoleManagementCloudPCRoleAssignmentsItemDirectoryScopesRequestBuilder) {
    return NewRoleManagementCloudPCRoleAssignmentsItemDirectoryScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryScopesById provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopesById(id string)(*RoleManagementCloudPCRoleAssignmentsItemDirectoryScopesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewRoleManagementCloudPCRoleAssignmentsItemDirectoryScopesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get roleAssignments from roleManagement
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, error) {
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
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, requestConfiguration *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, error) {
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
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Principals()(*RoleManagementCloudPCRoleAssignmentsItemPrincipalsRequestBuilder) {
    return NewRoleManagementCloudPCRoleAssignmentsItemPrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrincipalsById provides operations to manage the principals property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) PrincipalsById(id string)(*RoleManagementCloudPCRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewRoleManagementCloudPCRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
func (m *RoleManagementCloudPCRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) RoleDefinition()(*RoleManagementCloudPCRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    return NewRoleManagementCloudPCRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
