package governanceresources

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GovernanceResourcesGovernanceResourceItemRequestBuilder provides operations to manage the collection of governanceResource entities.
type GovernanceResourcesGovernanceResourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GovernanceResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GovernanceResourcesGovernanceResourceItemRequestBuilderGetQueryParameters get entity from governanceResources by key
type GovernanceResourcesGovernanceResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GovernanceResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GovernanceResourcesGovernanceResourceItemRequestBuilderGetQueryParameters
}
// GovernanceResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGovernanceResourcesGovernanceResourceItemRequestBuilderInternal instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewGovernanceResourcesGovernanceResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceResourcesGovernanceResourceItemRequestBuilder) {
    m := &GovernanceResourcesGovernanceResourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceResources/{governanceResource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceResourcesGovernanceResourceItemRequestBuilder instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewGovernanceResourcesGovernanceResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceResourcesGovernanceResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceResourcesGovernanceResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from governanceResources
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GovernanceResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from governanceResources by key
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GovernanceResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in governanceResources
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *GovernanceResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from governanceResources
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GovernanceResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get entity from governanceResources by key
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GovernanceResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
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
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) Parent()(*GovernanceResourcesItemParentRequestBuilder) {
    return NewGovernanceResourcesItemParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in governanceResources
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *GovernanceResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
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
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentRequests()(*GovernanceResourcesItemRoleAssignmentRequestsRequestBuilder) {
    return NewGovernanceResourcesItemRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentRequestsById provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentRequestsById(id string)(*GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest%2Did"] = id
    }
    return NewGovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) RoleAssignments()(*GovernanceResourcesItemRoleAssignmentsRequestBuilder) {
    return NewGovernanceResourcesItemRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentsById(id string)(*GovernanceResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment%2Did"] = id
    }
    return NewGovernanceResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.governanceResource entity.
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) RoleDefinitions()(*GovernanceResourcesItemRoleDefinitionsRequestBuilder) {
    return NewGovernanceResourcesItemRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.governanceResource entity.
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) RoleDefinitionsById(id string)(*GovernanceResourcesItemRoleDefinitionsGovernanceRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition%2Did"] = id
    }
    return NewGovernanceResourcesItemRoleDefinitionsGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleSettings provides operations to manage the roleSettings property of the microsoft.graph.governanceResource entity.
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) RoleSettings()(*GovernanceResourcesItemRoleSettingsRequestBuilder) {
    return NewGovernanceResourcesItemRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSettingsById provides operations to manage the roleSettings property of the microsoft.graph.governanceResource entity.
func (m *GovernanceResourcesGovernanceResourceItemRequestBuilder) RoleSettingsById(id string)(*GovernanceResourcesItemRoleSettingsGovernanceRoleSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting%2Did"] = id
    }
    return NewGovernanceResourcesItemRoleSettingsGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
