package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder provides operations to call the resume method.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderInternal instantiates a new ResumeRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder) {
    m := &IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/userProcessingResults/{userProcessingResult%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/microsoft.graph.identityGovernance.resume";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder instantiates a new ResumeRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation resume a task processing result that's `inProgress`. An Azure Logic Apps system-assigned managed identity calls this API.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBodyable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumePostRequestBodyable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsItemResumeRequestBuilderPostRequestConfiguration)(error) {
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
