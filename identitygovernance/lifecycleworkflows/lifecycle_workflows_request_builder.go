package lifecycleworkflows

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
    i18009dcca459c246d46cb5828cc0eee8f1d993b91686c706db8e83fb460509d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflowtemplates"
    i37c70f03fae165b58938fc0ae08b3f677d0db4bdb15ca0c09368f6e67555bbbf "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/settings"
    i5707969f667414c398d3c9b22236a007c32e34c6252415a86d6bca6c57e743d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/customtaskextensions"
    i59b00a625634744f30edbe5af244c841751a117e220b6c04a8154bbda71cecbd "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows"
    i6808840b5526833bc69eb1b4e6cf26bd083d4894645ce8021775399f27fc430f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems"
    id28ac37f235479a4d042611c03e18df4abd08307ecf13f724585b68b40cb8401 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/taskdefinitions"
    i43e6642142b931dba778895fb75fdcc74949ad3affdac1525d7089a8698e0c53 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/taskdefinitions/item"
    i5256431ba63928d796b312fa77a2f1c20b3bf8896afabf580168e5adca606b7b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflowtemplates/item"
    i9981a40a98032ec47f106a3d59efd9b83b8c90577133fc2f18f9835fc22e6dcd "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item"
    ib86c43a2f2e75a7aa5fbd68fd99f6cea961e6ae68ee0209e49d5248994f78717 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/customtaskextensions/item"
)

// LifecycleWorkflowsRequestBuilder provides operations to manage the lifecycleWorkflows property of the microsoft.graph.identityGovernance entity.
type LifecycleWorkflowsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LifecycleWorkflowsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LifecycleWorkflowsRequestBuilderGetQueryParameters get lifecycleWorkflows from identityGovernance
type LifecycleWorkflowsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsRequestBuilderGetQueryParameters
}
// LifecycleWorkflowsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleWorkflowsRequestBuilderInternal instantiates a new LifecycleWorkflowsRequestBuilder and sets the default values.
func NewLifecycleWorkflowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsRequestBuilder) {
    m := &LifecycleWorkflowsRequestBuilder{
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
// NewLifecycleWorkflowsRequestBuilder instantiates a new LifecycleWorkflowsRequestBuilder and sets the default values.
func NewLifecycleWorkflowsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property lifecycleWorkflows for identityGovernance
func (m *LifecycleWorkflowsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LifecycleWorkflowsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LifecycleWorkflowsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable, requestConfiguration *LifecycleWorkflowsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LifecycleWorkflowsRequestBuilder) CustomTaskExtensions()(*i5707969f667414c398d3c9b22236a007c32e34c6252415a86d6bca6c57e743d0.CustomTaskExtensionsRequestBuilder) {
    return i5707969f667414c398d3c9b22236a007c32e34c6252415a86d6bca6c57e743d0.NewCustomTaskExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomTaskExtensionsById provides operations to manage the customTaskExtensions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *LifecycleWorkflowsRequestBuilder) CustomTaskExtensionsById(id string)(*ib86c43a2f2e75a7aa5fbd68fd99f6cea961e6ae68ee0209e49d5248994f78717.CustomTaskExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customTaskExtension%2Did"] = id
    }
    return ib86c43a2f2e75a7aa5fbd68fd99f6cea961e6ae68ee0209e49d5248994f78717.NewCustomTaskExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property lifecycleWorkflows for identityGovernance
func (m *LifecycleWorkflowsRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleWorkflowsRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *LifecycleWorkflowsRequestBuilder) DeletedItems()(*i6808840b5526833bc69eb1b4e6cf26bd083d4894645ce8021775399f27fc430f.DeletedItemsRequestBuilder) {
    return i6808840b5526833bc69eb1b4e6cf26bd083d4894645ce8021775399f27fc430f.NewDeletedItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get lifecycleWorkflows from identityGovernance
func (m *LifecycleWorkflowsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable, error) {
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
func (m *LifecycleWorkflowsRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable, requestConfiguration *LifecycleWorkflowsRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.LifecycleWorkflowsContainerable, error) {
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
func (m *LifecycleWorkflowsRequestBuilder) Settings()(*i37c70f03fae165b58938fc0ae08b3f677d0db4bdb15ca0c09368f6e67555bbbf.SettingsRequestBuilder) {
    return i37c70f03fae165b58938fc0ae08b3f677d0db4bdb15ca0c09368f6e67555bbbf.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskDefinitions provides operations to manage the taskDefinitions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *LifecycleWorkflowsRequestBuilder) TaskDefinitions()(*id28ac37f235479a4d042611c03e18df4abd08307ecf13f724585b68b40cb8401.TaskDefinitionsRequestBuilder) {
    return id28ac37f235479a4d042611c03e18df4abd08307ecf13f724585b68b40cb8401.NewTaskDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskDefinitionsById provides operations to manage the taskDefinitions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *LifecycleWorkflowsRequestBuilder) TaskDefinitionsById(id string)(*i43e6642142b931dba778895fb75fdcc74949ad3affdac1525d7089a8698e0c53.TaskDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskDefinition%2Did"] = id
    }
    return i43e6642142b931dba778895fb75fdcc74949ad3affdac1525d7089a8698e0c53.NewTaskDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Workflows provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *LifecycleWorkflowsRequestBuilder) Workflows()(*i59b00a625634744f30edbe5af244c841751a117e220b6c04a8154bbda71cecbd.WorkflowsRequestBuilder) {
    return i59b00a625634744f30edbe5af244c841751a117e220b6c04a8154bbda71cecbd.NewWorkflowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkflowsById provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *LifecycleWorkflowsRequestBuilder) WorkflowsById(id string)(*i9981a40a98032ec47f106a3d59efd9b83b8c90577133fc2f18f9835fc22e6dcd.WorkflowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflow%2Did"] = id
    }
    return i9981a40a98032ec47f106a3d59efd9b83b8c90577133fc2f18f9835fc22e6dcd.NewWorkflowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WorkflowTemplates provides operations to manage the workflowTemplates property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *LifecycleWorkflowsRequestBuilder) WorkflowTemplates()(*i18009dcca459c246d46cb5828cc0eee8f1d993b91686c706db8e83fb460509d0.WorkflowTemplatesRequestBuilder) {
    return i18009dcca459c246d46cb5828cc0eee8f1d993b91686c706db8e83fb460509d0.NewWorkflowTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkflowTemplatesById provides operations to manage the workflowTemplates property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
func (m *LifecycleWorkflowsRequestBuilder) WorkflowTemplatesById(id string)(*i5256431ba63928d796b312fa77a2f1c20b3bf8896afabf580168e5adca606b7b.WorkflowTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflowTemplate%2Did"] = id
    }
    return i5256431ba63928d796b312fa77a2f1c20b3bf8896afabf580168e5adca606b7b.NewWorkflowTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
