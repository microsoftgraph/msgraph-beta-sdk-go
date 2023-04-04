package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder provides operations to manage the workflows property of the microsoft.graph.deletedItemContainer entity.
type LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderInternal instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property workflows for identityGovernance
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ExecutionScope provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) ExecutionScope()(*LifecycleWorkflowsDeletedItemsWorkflowsItemExecutionScopeRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemExecutionScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExecutionScopeById provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) ExecutionScopeById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemExecutionScopeUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemExecutionScopeUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Get deleted workflows that end up in the deletedItemsContainer.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable), nil
}
// IdentityGovernanceActivate provides operations to call the activate method.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) IdentityGovernanceActivate()(*LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IdentityGovernanceCreateNewVersion provides operations to call the createNewVersion method.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) IdentityGovernanceCreateNewVersion()(*LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceCreateNewVersionRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceCreateNewVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IdentityGovernanceRestore provides operations to call the restore method.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) IdentityGovernanceRestore()(*LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceRestoreRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Runs provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Runs()(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RunsById provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) RunsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["run%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// TaskReports provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) TaskReports()(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskReportsById provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) TaskReportsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskReport%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsTaskReportItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property workflows for identityGovernance
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation deleted workflows that end up in the deletedItemsContainer.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) UserProcessingResults()(*LifecycleWorkflowsDeletedItemsWorkflowsItemUserProcessingResultsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemUserProcessingResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserProcessingResultsById provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) UserProcessingResultsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userProcessingResult%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) Versions()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsWorkflowItemRequestBuilder) VersionsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflowVersion%2DversionNumber"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
