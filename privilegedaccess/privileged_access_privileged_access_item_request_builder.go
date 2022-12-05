package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessPrivilegedAccessItemRequestBuilder provides operations to manage the collection of privilegedAccess entities.
type PrivilegedAccessPrivilegedAccessItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedAccessPrivilegedAccessItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessPrivilegedAccessItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessPrivilegedAccessItemRequestBuilderGetQueryParameters get entity from privilegedAccess by key
type PrivilegedAccessPrivilegedAccessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessPrivilegedAccessItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessPrivilegedAccessItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessPrivilegedAccessItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessPrivilegedAccessItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessPrivilegedAccessItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedAccessPrivilegedAccessItemRequestBuilderInternal instantiates a new PrivilegedAccessItemRequestBuilder and sets the default values.
func NewPrivilegedAccessPrivilegedAccessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessPrivilegedAccessItemRequestBuilder) {
    m := &PrivilegedAccessPrivilegedAccessItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedAccessPrivilegedAccessItemRequestBuilder instantiates a new PrivilegedAccessItemRequestBuilder and sets the default values.
func NewPrivilegedAccessPrivilegedAccessItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessPrivilegedAccessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessPrivilegedAccessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from privilegedAccess
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessPrivilegedAccessItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from privilegedAccess by key
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessPrivilegedAccessItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in privilegedAccess
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessable, requestConfiguration *PrivilegedAccessPrivilegedAccessItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from privilegedAccess
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessPrivilegedAccessItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get entity from privilegedAccess by key
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessPrivilegedAccessItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessable), nil
}
// Patch update entity in privilegedAccess
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessable, requestConfiguration *PrivilegedAccessPrivilegedAccessItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessable), nil
}
// Resources provides operations to manage the resources property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) Resources()(*PrivilegedAccessItemResourcesRequestBuilder) {
    return NewPrivilegedAccessItemResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) ResourcesById(id string)(*PrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceResource%2Did"] = id
    }
    return NewPrivilegedAccessItemResourcesGovernanceResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignmentRequests provides operations to manage the roleAssignmentRequests property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) RoleAssignmentRequests()(*PrivilegedAccessItemRoleAssignmentRequestsRequestBuilder) {
    return NewPrivilegedAccessItemRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentRequestsById provides operations to manage the roleAssignmentRequests property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) RoleAssignmentRequestsById(id string)(*PrivilegedAccessItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest%2Did"] = id
    }
    return NewPrivilegedAccessItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) RoleAssignments()(*PrivilegedAccessItemRoleAssignmentsRequestBuilder) {
    return NewPrivilegedAccessItemRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) RoleAssignmentsById(id string)(*PrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment%2Did"] = id
    }
    return NewPrivilegedAccessItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) RoleDefinitions()(*PrivilegedAccessItemRoleDefinitionsRequestBuilder) {
    return NewPrivilegedAccessItemRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) RoleDefinitionsById(id string)(*PrivilegedAccessItemRoleDefinitionsGovernanceRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition%2Did"] = id
    }
    return NewPrivilegedAccessItemRoleDefinitionsGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleSettings provides operations to manage the roleSettings property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) RoleSettings()(*PrivilegedAccessItemRoleSettingsRequestBuilder) {
    return NewPrivilegedAccessItemRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSettingsById provides operations to manage the roleSettings property of the microsoft.graph.privilegedAccess entity.
func (m *PrivilegedAccessPrivilegedAccessItemRequestBuilder) RoleSettingsById(id string)(*PrivilegedAccessItemRoleSettingsGovernanceRoleSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting%2Did"] = id
    }
    return NewPrivilegedAccessItemRoleSettingsGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
