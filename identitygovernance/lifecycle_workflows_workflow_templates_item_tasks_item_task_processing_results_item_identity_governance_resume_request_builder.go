package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder provides operations to call the resume method.
type LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderInternal instantiates a new IdentityGovernanceResumeRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflowTemplates/{workflowTemplate%2Did}/tasks/{task%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/identityGovernance.resume", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder instantiates a new IdentityGovernanceResumeRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post resume a task processing result that's `inProgress`. In the default case an Azure Logic Apps system-assigned managed identity calls this API. For more information read about Lifecycle Workflows extensibility approach.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/identitygovernance-taskprocessingresult-resume?view=graph-rest-1.0
func (m *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder) Post(ctx context.Context, body LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeResumePostRequestBodyable, requestConfiguration *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation resume a task processing result that's `inProgress`. In the default case an Azure Logic Apps system-assigned managed identity calls this API. For more information read about Lifecycle Workflows extensibility approach.
func (m *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder) ToPostRequestInformation(ctx context.Context, body LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeResumePostRequestBodyable, requestConfiguration *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
