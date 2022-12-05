package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder provides operations to manage the resources property of the microsoft.graph.privilegedAccess entity.
type PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters a collection of resources for the provider.
type PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderInternal instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) {
    m := &PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resources for privilegedAccess
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of resources for the provider.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property resources in privilegedAccess
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property resources for privilegedAccess
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of resources for the provider.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// Parent provides operations to manage the parent property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) Parent()(*PrivilegedAccessItemResourcesItemParentRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property resources in privilegedAccess
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// RoleAssignmentRequests provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentRequests()(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentRequestsById provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentRequestsById(id string)(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest%2Did"] = id
    }
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignments()(*PrivilegedAccessItemResourcesItemRoleAssignmentsRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentsById(id string)(*PrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment%2Did"] = id
    }
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) RoleDefinitions()(*PrivilegedAccessItemResourcesItemRoleDefinitionsRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) RoleDefinitionsById(id string)(*PrivilegedAccessItemResourcesItemRoleDefinitionsGovernanceRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition%2Did"] = id
    }
    return NewPrivilegedAccessItemResourcesItemRoleDefinitionsGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleSettings provides operations to manage the roleSettings property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) RoleSettings()(*PrivilegedAccessItemResourcesItemRoleSettingsRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSettingsById provides operations to manage the roleSettings property of the microsoft.graph.governanceResource entity.
func (m *PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) RoleSettingsById(id string)(*PrivilegedAccessItemResourcesItemRoleSettingsGovernanceRoleSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting%2Did"] = id
    }
    return NewPrivilegedAccessItemResourcesItemRoleSettingsGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
