package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
    i05a5138a4e33f7e83896eacd4c6fd83d689731f5d0a139cd1bf922f9f49c4aec "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/versions"
    i313445eff5f5b36a7c96ea00a400493cef88d5c9739c4a86f54522dc923afb5f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/userprocessingresults"
    i8cf137125336cf3388d0d941867bc83ad6c65b4792df2a56c0d38480e18514e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/createnewversion"
    i934fe4f55a2953e59cde9d68812afe753793bccb49bd983f66eb03de7818ddb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/executionscope"
    i9c7801e61117ebd031766d3bba5d64703219864956b206a667ccb7693e19be01 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/runs"
    ia296fd3f1e5ff3c40ad60f799e79db0663a31846e3148df92a8094390a84cbc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/activate"
    id1a69d66fcdba1734751a42ddd76d3290528f1a35ebb6abc86009b8bd9bc7ae9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/restore"
    ie017ea425d3c3d3e3a14e4eba029e59ff7d30b45b86a6679525d524b2394f958 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/taskreports"
    i142e5b6ba183b550be2173bc3636299e41d86cf0262dccc248a5ffdc7efb99f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/taskreports/item"
    i83e8cbd4cf7df11f171e7aa3cad7bef500b6ee10a4be4b785640618dad1171ce "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/userprocessingresults/item"
    i8da9737a36642b2e91c4eff981addd7f2de8405fe0688fe557f971af4a13ac18 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/executionscope/item"
    iaf85afa1d6b8f410137ec7672bd355b194e453bcdfa855d80caa7244c13e5984 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/runs/item"
    ie76c3e625e2dc0bf84c56645d4951120ee9f62d232c50d3eddcf264dc8035d22 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/workflows/item/versions/item"
)

// WorkflowItemRequestBuilder provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type WorkflowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WorkflowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WorkflowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WorkflowItemRequestBuilderGetQueryParameters the workflows in the lifecycle workflows instance.
type WorkflowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WorkflowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WorkflowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WorkflowItemRequestBuilderGetQueryParameters
}
// WorkflowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WorkflowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate the activate property
func (m *WorkflowItemRequestBuilder) Activate()(*ia296fd3f1e5ff3c40ad60f799e79db0663a31846e3148df92a8094390a84cbc6.ActivateRequestBuilder) {
    return ia296fd3f1e5ff3c40ad60f799e79db0663a31846e3148df92a8094390a84cbc6.NewActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkflowItemRequestBuilderInternal instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewWorkflowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WorkflowItemRequestBuilder) {
    m := &WorkflowItemRequestBuilder{
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
// NewWorkflowItemRequestBuilder instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewWorkflowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WorkflowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkflowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property workflows for identityGovernance
func (m *WorkflowItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *WorkflowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WorkflowItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *WorkflowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateNewVersion the createNewVersion property
func (m *WorkflowItemRequestBuilder) CreateNewVersion()(*i8cf137125336cf3388d0d941867bc83ad6c65b4792df2a56c0d38480e18514e6.CreateNewVersionRequestBuilder) {
    return i8cf137125336cf3388d0d941867bc83ad6c65b4792df2a56c0d38480e18514e6.NewCreateNewVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property workflows in identityGovernance
func (m *WorkflowItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, requestConfiguration *WorkflowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WorkflowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WorkflowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExecutionScope the executionScope property
func (m *WorkflowItemRequestBuilder) ExecutionScope()(*i934fe4f55a2953e59cde9d68812afe753793bccb49bd983f66eb03de7818ddb3.ExecutionScopeRequestBuilder) {
    return i934fe4f55a2953e59cde9d68812afe753793bccb49bd983f66eb03de7818ddb3.NewExecutionScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExecutionScopeById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.lifecycleWorkflows.workflows.item.executionScope.item collection
func (m *WorkflowItemRequestBuilder) ExecutionScopeById(id string)(*i8da9737a36642b2e91c4eff981addd7f2de8405fe0688fe557f971af4a13ac18.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return i8da9737a36642b2e91c4eff981addd7f2de8405fe0688fe557f971af4a13ac18.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the workflows in the lifecycle workflows instance.
func (m *WorkflowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WorkflowItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, error) {
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
func (m *WorkflowItemRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, requestConfiguration *WorkflowItemRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Workflowable, error) {
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
// Restore the restore property
func (m *WorkflowItemRequestBuilder) Restore()(*id1a69d66fcdba1734751a42ddd76d3290528f1a35ebb6abc86009b8bd9bc7ae9.RestoreRequestBuilder) {
    return id1a69d66fcdba1734751a42ddd76d3290528f1a35ebb6abc86009b8bd9bc7ae9.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Runs the runs property
func (m *WorkflowItemRequestBuilder) Runs()(*i9c7801e61117ebd031766d3bba5d64703219864956b206a667ccb7693e19be01.RunsRequestBuilder) {
    return i9c7801e61117ebd031766d3bba5d64703219864956b206a667ccb7693e19be01.NewRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RunsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.lifecycleWorkflows.workflows.item.runs.item collection
func (m *WorkflowItemRequestBuilder) RunsById(id string)(*iaf85afa1d6b8f410137ec7672bd355b194e453bcdfa855d80caa7244c13e5984.RunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["run%2Did"] = id
    }
    return iaf85afa1d6b8f410137ec7672bd355b194e453bcdfa855d80caa7244c13e5984.NewRunItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskReports the taskReports property
func (m *WorkflowItemRequestBuilder) TaskReports()(*ie017ea425d3c3d3e3a14e4eba029e59ff7d30b45b86a6679525d524b2394f958.TaskReportsRequestBuilder) {
    return ie017ea425d3c3d3e3a14e4eba029e59ff7d30b45b86a6679525d524b2394f958.NewTaskReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.lifecycleWorkflows.workflows.item.taskReports.item collection
func (m *WorkflowItemRequestBuilder) TaskReportsById(id string)(*i142e5b6ba183b550be2173bc3636299e41d86cf0262dccc248a5ffdc7efb99f4.TaskReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskReport%2Did"] = id
    }
    return i142e5b6ba183b550be2173bc3636299e41d86cf0262dccc248a5ffdc7efb99f4.NewTaskReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserProcessingResults the userProcessingResults property
func (m *WorkflowItemRequestBuilder) UserProcessingResults()(*i313445eff5f5b36a7c96ea00a400493cef88d5c9739c4a86f54522dc923afb5f.UserProcessingResultsRequestBuilder) {
    return i313445eff5f5b36a7c96ea00a400493cef88d5c9739c4a86f54522dc923afb5f.NewUserProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserProcessingResultsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.lifecycleWorkflows.workflows.item.userProcessingResults.item collection
func (m *WorkflowItemRequestBuilder) UserProcessingResultsById(id string)(*i83e8cbd4cf7df11f171e7aa3cad7bef500b6ee10a4be4b785640618dad1171ce.UserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userProcessingResult%2Did"] = id
    }
    return i83e8cbd4cf7df11f171e7aa3cad7bef500b6ee10a4be4b785640618dad1171ce.NewUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *WorkflowItemRequestBuilder) Versions()(*i05a5138a4e33f7e83896eacd4c6fd83d689731f5d0a139cd1bf922f9f49c4aec.VersionsRequestBuilder) {
    return i05a5138a4e33f7e83896eacd4c6fd83d689731f5d0a139cd1bf922f9f49c4aec.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.lifecycleWorkflows.workflows.item.versions.item collection
func (m *WorkflowItemRequestBuilder) VersionsById(id string)(*ie76c3e625e2dc0bf84c56645d4951120ee9f62d232c50d3eddcf264dc8035d22.WorkflowVersionVersionNumberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflowVersion%2DversionNumber"] = id
    }
    return ie76c3e625e2dc0bf84c56645d4951120ee9f62d232c50d3eddcf264dc8035d22.NewWorkflowVersionVersionNumberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
