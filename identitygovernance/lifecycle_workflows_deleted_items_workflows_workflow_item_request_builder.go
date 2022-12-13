package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder provides operations to manage the workflows property of the microsoft.graph.deletedItemContainer entity.
type LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetQueryParameters deleted workflows that end up in the deletedItemsContainer.
type LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetQueryParameters
}
// LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate provides operations to call the activate method.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Activate()(*LifecycleWorkflowsDeletedItemsWorkflowsItemActivateRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderInternal instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property workflows for identityGovernance
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation deleted workflows that end up in the deletedItemsContainer.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateNewVersion provides operations to call the createNewVersion method.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) CreateNewVersion()(*LifecycleWorkflowsDeletedItemsWorkflowsItemCreateNewVersionRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemCreateNewVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property workflows in identityGovernance
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property workflows for identityGovernance
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExecutionScope provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) ExecutionScope()(*LifecycleWorkflowsDeletedItemsWorkflowsItemExecutionScopeRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemExecutionScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExecutionScopeById provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) ExecutionScopeById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemExecutionScopeUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemExecutionScopeUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get deleted workflows that end up in the deletedItemsContainer.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable), nil
}
// Patch update the navigation property workflows in identityGovernance
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable), nil
}
// Restore provides operations to call the restore method.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Restore()(*LifecycleWorkflowsDeletedItemsWorkflowsItemRestoreRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Runs provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Runs()(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RunsById provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) RunsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["run%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskReports provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) TaskReports()(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskReportsById provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) TaskReportsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskReport%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsTaskReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) UserProcessingResults()(*LifecycleWorkflowsDeletedItemsWorkflowsItemUserProcessingResultsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemUserProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserProcessingResultsById provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) UserProcessingResultsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userProcessingResult%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Versions()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) VersionsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflowVersion%2DversionNumber"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
