package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
type RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderGetQueryParameters get roleAssignments from roleManagement
type RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderGetQueryParameters
}
// RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleAssignment entity.
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) AppScope()(*RoleManagementDirectoryRoleAssignmentsItemAppScopeRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentsItemAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentItemRequestBuilder and sets the default values.
func NewRoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    m := &RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleAssignments/{unifiedRoleAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder instantiates a new UnifiedRoleAssignmentItemRequestBuilder and sets the default values.
func NewRoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for roleManagement
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, requestConfiguration *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleAssignment entity.
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) DirectoryScope()(*RoleManagementDirectoryRoleAssignmentsItemDirectoryScopeRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentsItemDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignments from roleManagement
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable), nil
}
// Patch update the navigation property roleAssignments in roleManagement
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, requestConfiguration *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleAssignment entity.
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) Principal()(*RoleManagementDirectoryRoleAssignmentsItemPrincipalRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentsItemPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignment entity.
func (m *RoleManagementDirectoryRoleAssignmentsUnifiedRoleAssignmentItemRequestBuilder) RoleDefinition()(*RoleManagementDirectoryRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    return NewRoleManagementDirectoryRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
