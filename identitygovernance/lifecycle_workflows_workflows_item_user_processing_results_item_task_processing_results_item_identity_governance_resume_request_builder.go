package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder provides operations to call the resume method.
type LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderInternal instantiates a new IdentityGovernanceResumeRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/userProcessingResults/{userProcessingResult%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/identityGovernance.resume";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder instantiates a new IdentityGovernanceResumeRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post resume a task processing result that's `inProgress`. An Azure Logic Apps system-assigned managed identity calls this API.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/identitygovernance-taskprocessingresult-resume?view=graph-rest-1.0
func (m *LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder) Post(ctx context.Context, body LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeResumePostRequestBodyable, requestConfiguration *LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation resume a task processing result that's `inProgress`. An Azure Logic Apps system-assigned managed identity calls this API.
func (m *LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilder) ToPostRequestInformation(ctx context.Context, body LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeResumePostRequestBodyable, requestConfiguration *LifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemIdentityGovernanceResumeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
