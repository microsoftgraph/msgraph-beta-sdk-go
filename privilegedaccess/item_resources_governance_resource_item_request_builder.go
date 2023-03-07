package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemResourcesGovernanceResourceItemRequestBuilder provides operations to manage the resources property of the microsoft.graph.privilegedAccess entity.
type ItemResourcesGovernanceResourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters a collection of resources for the provider.
type ItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters
}
// ItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemResourcesGovernanceResourceItemRequestBuilderInternal instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewItemResourcesGovernanceResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesGovernanceResourceItemRequestBuilder) {
    m := &ItemResourcesGovernanceResourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemResourcesGovernanceResourceItemRequestBuilder instantiates a new GovernanceResourceItemRequestBuilder and sets the default values.
func NewItemResourcesGovernanceResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesGovernanceResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResourcesGovernanceResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resources for privilegedAccess
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of resources for the provider.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// Parent provides operations to manage the parent property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) Parent()(*ItemResourcesItemParentRequestBuilder) {
    return NewItemResourcesItemParentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property resources in privilegedAccess
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// RoleAssignmentRequests provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentRequests()(*ItemResourcesItemRoleAssignmentRequestsRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleAssignmentRequestsById provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentRequestsById(id string)(*ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest%2Did"] = id
    }
    return NewItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignments()(*ItemResourcesItemRoleAssignmentsRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentsById(id string)(*ItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment%2Did"] = id
    }
    return NewItemResourcesItemRoleAssignmentsGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleDefinitions()(*ItemResourcesItemRoleDefinitionsRequestBuilder) {
    return NewItemResourcesItemRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleDefinitionsById(id string)(*ItemResourcesItemRoleDefinitionsGovernanceRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition%2Did"] = id
    }
    return NewItemResourcesItemRoleDefinitionsGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleSettings provides operations to manage the roleSettings property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleSettings()(*ItemResourcesItemRoleSettingsRequestBuilder) {
    return NewItemResourcesItemRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleSettingsById provides operations to manage the roleSettings property of the microsoft.graph.governanceResource entity.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleSettingsById(id string)(*ItemResourcesItemRoleSettingsGovernanceRoleSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting%2Did"] = id
    }
    return NewItemResourcesItemRoleSettingsGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property resources for privilegedAccess
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation a collection of resources for the provider.
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property resources in privilegedAccess
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
