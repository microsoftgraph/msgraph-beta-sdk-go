package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.taskReport entity.
type LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters the related lifecycle workflow taskProcessingResults.
type LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal instantiates a new TaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}/taskProcessingResults/{taskProcessingResult%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder instantiates a new TaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the related lifecycle workflow taskProcessingResults.
func (m *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskProcessingResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateTaskProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.TaskProcessingResultable), nil
}
// MicrosoftGraphIdentityGovernanceResume provides operations to call the resume method.
func (m *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) MicrosoftGraphIdentityGovernanceResume()(*LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Subject provides operations to manage the subject property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
func (m *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Subject()(*LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsItemSubjectRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsItemSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Task provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
func (m *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Task()(*LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation the related lifecycle workflow taskProcessingResults.
func (m *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
