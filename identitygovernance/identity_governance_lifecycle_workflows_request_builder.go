package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// IdentityGovernanceLifecycleWorkflowsRequestBuilder provides operations to manage the lifecycleWorkflows property of the microsoft.graph.identityGovernance entity.
type IdentityGovernanceLifecycleWorkflowsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceLifecycleWorkflowsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceLifecycleWorkflowsRequestBuilderGetQueryParameters get lifecycleWorkflows from identityGovernance
type IdentityGovernanceLifecycleWorkflowsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceLifecycleWorkflowsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceLifecycleWorkflowsRequestBuilderGetQueryParameters
}
// IdentityGovernanceLifecycleWorkflowsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceLifecycleWorkflowsRequestBuilderInternal instantiates a new LifecycleWorkflowsRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsRequestBuilder) {
    m := &IdentityGovernanceLifecycleWorkflowsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceLifecycleWorkflowsRequestBuilder instantiates a new LifecycleWorkflowsRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceLifecycleWorkflowsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property lifecycleWorkflows for identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get lifecycleWorkflows from identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property lifecycleWorkflows in identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CustomTaskExtensions provides operations to manage the customTaskExtensions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) CustomTaskExtensions()(*IdentityGovernanceLifecycleWorkflowsCustomTaskExtensionsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsCustomTaskExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomTaskExtensionsById provides operations to manage the customTaskExtensions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) CustomTaskExtensionsById(id string)(*IdentityGovernanceLifecycleWorkflowsCustomTaskExtensionsCustomTaskExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customTaskExtension%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsCustomTaskExtensionsCustomTaskExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property lifecycleWorkflows for identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeletedItems provides operations to manage the deletedItems property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) DeletedItems()(*IdentityGovernanceLifecycleWorkflowsDeletedItemsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsDeletedItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get lifecycleWorkflows from identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateLifecycleWorkflowsContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable), nil
}
// Patch update the navigation property lifecycleWorkflows in identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateLifecycleWorkflowsContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) Settings()(*IdentityGovernanceLifecycleWorkflowsSettingsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskDefinitions provides operations to manage the taskDefinitions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) TaskDefinitions()(*IdentityGovernanceLifecycleWorkflowsTaskDefinitionsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsTaskDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskDefinitionsById provides operations to manage the taskDefinitions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) TaskDefinitionsById(id string)(*IdentityGovernanceLifecycleWorkflowsTaskDefinitionsTaskDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskDefinition%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsTaskDefinitionsTaskDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Workflows provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) Workflows()(*IdentityGovernanceLifecycleWorkflowsWorkflowsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkflowsById provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) WorkflowsById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflow%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WorkflowTemplates provides operations to manage the workflowTemplates property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) WorkflowTemplates()(*IdentityGovernanceLifecycleWorkflowsWorkflowTemplatesRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkflowTemplatesById provides operations to manage the workflowTemplates property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *IdentityGovernanceLifecycleWorkflowsRequestBuilder) WorkflowTemplatesById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowTemplatesWorkflowTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflowTemplate%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowTemplatesWorkflowTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
