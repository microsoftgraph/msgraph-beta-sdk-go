package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderGetQueryParameters represents the aggregation of task execution data for tasks within a workflow object.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderInternal instantiates a new TaskReportItemRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) {
    m := &IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder instantiates a new TaskReportItemRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property taskReports for identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents the aggregation of task execution data for tasks within a workflow object.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property taskReports in identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskReportable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property taskReports for identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the aggregation of task execution data for tasks within a workflow object.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskReportable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateTaskReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskReportable), nil
}
// Patch update the navigation property taskReports in identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskReportable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskReportable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateTaskReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskReportable), nil
}
// Task provides operations to manage the task property of the microsoft.graph.identityGovernance.taskReport entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) Task()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskDefinition provides operations to manage the taskDefinition property of the microsoft.graph.identityGovernance.taskReport entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) TaskDefinition()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskProcessingResults provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.taskReport entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) TaskProcessingResults()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskProcessingResultsById provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.taskReport entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsTaskReportItemRequestBuilder) TaskProcessingResultsById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskProcessingResult%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
