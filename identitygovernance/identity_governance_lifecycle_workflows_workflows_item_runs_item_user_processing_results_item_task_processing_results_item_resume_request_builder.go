package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder provides operations to call the resume method.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderInternal instantiates a new ResumeRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder) {
    m := &IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/runs/{run%2Did}/userProcessingResults/{userProcessingResult%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/microsoft.graph.identityGovernance.resume";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder instantiates a new ResumeRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation resume a task processing result that's `inProgress`. An Azure Logic Apps system-assigned managed identity calls this API.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBodyable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post resume a task processing result that's `inProgress`. An Azure Logic Apps system-assigned managed identity calls this API.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBodyable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
