package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters the workflows in the lifecycle workflows instance.
type IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate provides operations to call the activate method.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Activate()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemActivateRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderInternal instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) {
    m := &IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property workflows for identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the workflows in the lifecycle workflows instance.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateNewVersion provides operations to call the createNewVersion method.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) CreateNewVersion()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemCreateNewVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property workflows in identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property workflows for identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) ExecutionScope()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemExecutionScopeRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemExecutionScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExecutionScopeById provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) ExecutionScopeById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemExecutionScopeUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemExecutionScopeUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the workflows in the lifecycle workflows instance.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, error) {
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
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, error) {
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
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Restore()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemRestoreRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Runs provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Runs()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RunsById provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) RunsById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["run%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsRunItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskReports provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) TaskReports()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskReportsById provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) TaskReportsById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskReport%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) UserProcessingResults()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserProcessingResultsById provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) UserProcessingResultsById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userProcessingResult%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Versions()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemVersionsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) VersionsById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflowVersion%2DversionNumber"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
