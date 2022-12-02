package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
type IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderGetQueryParameters the related workflow task
type IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderGetQueryParameters
}
// NewIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderInternal instantiates a new TaskRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder) {
    m := &IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/task{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder instantiates a new TaskRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the related workflow task
func (m *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the related workflow task
func (m *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Taskable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Taskable), nil
}
