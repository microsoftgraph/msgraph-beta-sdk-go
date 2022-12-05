package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
type RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters get resourceNamespaces from roleManagement
type RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters
}
// RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal instantiates a new UnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    m := &RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder instantiates a new UnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resourceNamespaces for roleManagement
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get resourceNamespaces from roleManagement
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property resourceNamespaces in roleManagement
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property resourceNamespaces for roleManagement
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get resourceNamespaces from roleManagement
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ImportResourceActions provides operations to call the importResourceActions method.
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ImportResourceActions()(*RoleManagementEntitlementManagementResourceNamespacesItemImportResourceActionsRequestBuilder) {
    return NewRoleManagementEntitlementManagementResourceNamespacesItemImportResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property resourceNamespaces in roleManagement
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ResourceActions provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActions()(*RoleManagementEntitlementManagementResourceNamespacesItemResourceActionsRequestBuilder) {
    return NewRoleManagementEntitlementManagementResourceNamespacesItemResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceActionsById provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
func (m *RoleManagementEntitlementManagementResourceNamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActionsById(id string)(*RoleManagementEntitlementManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceAction%2Did"] = id
    }
    return NewRoleManagementEntitlementManagementResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
