package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
    i33107b1fbb2c6dc8d5c107552777e6d9083049971b55368abaf6a448675cb970 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/activate"
    i3a1a0eb953932015b6ecdca974faf4c262c5ecea8b6b8c8707e8a512a9430647 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/executionscope"
    i6b97b143636762ca7cad16e985bc443f0126bf521630dc036ebdf1b1d32afdb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/createnewversion"
    i72ce7d58e9852fd932435a826c49ca0ceb68d0fface3230435dcb902fc0e1ac0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/versions"
    ia11651440d0da0b9d4ca2e8dc624d1fbc24b791b414523ff004d45697c22b89a "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/taskreports"
    icc73530d36c372edf33706373bb5a31710c24e0832d3a561aeee7d04cae9d1d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/runs"
    if1129de09b88d2836dc73f021e9a09f3854ffa329edebc9e5eefde3e9f3f2f4c "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/userprocessingresults"
    ifec8ad6d85be4ee106958d5dd7e0c25171275c5871ff476af8af0a3dd693af93 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/restore"
    i6ad171a2e2184a06bf3c8381d3823f8984897689a23d81f145ceddf152a6368d "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/userprocessingresults/item"
    i6ef5d654ec18954956bcc3f06c9a25a2996d35ad34f5c9bedbdf3063fa0fc2e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/versions/item"
    ia541f436e8d4506f2fa7170547fac1b08f1ab90aa101530b5143693c6d53790a "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/taskreports/item"
    ic89dd7714b62d0dbd4595b1ea7fef4e63a6625693bc6d58ce9a2cbc1f5ffcc20 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/runs/item"
    idbdb585d24d72a7d4a6280168e5b9b47a815e87a6055e162778e0dae48b4328e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/executionscope/item"
)

// WorkflowItemRequestBuilder provides operations to manage the workflows property of the microsoft.graph.deletedItemContainer entity.
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
// WorkflowItemRequestBuilderGetQueryParameters deleted workflows that end up in the deletedItemsContainer.
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
// Activate provides operations to call the activate method.
func (m *WorkflowItemRequestBuilder) Activate()(*i33107b1fbb2c6dc8d5c107552777e6d9083049971b55368abaf6a448675cb970.ActivateRequestBuilder) {
    return i33107b1fbb2c6dc8d5c107552777e6d9083049971b55368abaf6a448675cb970.NewActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkflowItemRequestBuilderInternal instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewWorkflowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WorkflowItemRequestBuilder) {
    m := &WorkflowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation deleted workflows that end up in the deletedItemsContainer.
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
// CreateNewVersion provides operations to call the createNewVersion method.
func (m *WorkflowItemRequestBuilder) CreateNewVersion()(*i6b97b143636762ca7cad16e985bc443f0126bf521630dc036ebdf1b1d32afdb9.CreateNewVersionRequestBuilder) {
    return i6b97b143636762ca7cad16e985bc443f0126bf521630dc036ebdf1b1d32afdb9.NewCreateNewVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// ExecutionScope provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) ExecutionScope()(*i3a1a0eb953932015b6ecdca974faf4c262c5ecea8b6b8c8707e8a512a9430647.ExecutionScopeRequestBuilder) {
    return i3a1a0eb953932015b6ecdca974faf4c262c5ecea8b6b8c8707e8a512a9430647.NewExecutionScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExecutionScopeById provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) ExecutionScopeById(id string)(*idbdb585d24d72a7d4a6280168e5b9b47a815e87a6055e162778e0dae48b4328e.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return idbdb585d24d72a7d4a6280168e5b9b47a815e87a6055e162778e0dae48b4328e.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get deleted workflows that end up in the deletedItemsContainer.
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
// Restore provides operations to call the restore method.
func (m *WorkflowItemRequestBuilder) Restore()(*ifec8ad6d85be4ee106958d5dd7e0c25171275c5871ff476af8af0a3dd693af93.RestoreRequestBuilder) {
    return ifec8ad6d85be4ee106958d5dd7e0c25171275c5871ff476af8af0a3dd693af93.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Runs provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) Runs()(*icc73530d36c372edf33706373bb5a31710c24e0832d3a561aeee7d04cae9d1d4.RunsRequestBuilder) {
    return icc73530d36c372edf33706373bb5a31710c24e0832d3a561aeee7d04cae9d1d4.NewRunsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RunsById provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) RunsById(id string)(*ic89dd7714b62d0dbd4595b1ea7fef4e63a6625693bc6d58ce9a2cbc1f5ffcc20.RunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["run%2Did"] = id
    }
    return ic89dd7714b62d0dbd4595b1ea7fef4e63a6625693bc6d58ce9a2cbc1f5ffcc20.NewRunItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaskReports provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) TaskReports()(*ia11651440d0da0b9d4ca2e8dc624d1fbc24b791b414523ff004d45697c22b89a.TaskReportsRequestBuilder) {
    return ia11651440d0da0b9d4ca2e8dc624d1fbc24b791b414523ff004d45697c22b89a.NewTaskReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskReportsById provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) TaskReportsById(id string)(*ia541f436e8d4506f2fa7170547fac1b08f1ab90aa101530b5143693c6d53790a.TaskReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskReport%2Did"] = id
    }
    return ia541f436e8d4506f2fa7170547fac1b08f1ab90aa101530b5143693c6d53790a.NewTaskReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) UserProcessingResults()(*if1129de09b88d2836dc73f021e9a09f3854ffa329edebc9e5eefde3e9f3f2f4c.UserProcessingResultsRequestBuilder) {
    return if1129de09b88d2836dc73f021e9a09f3854ffa329edebc9e5eefde3e9f3f2f4c.NewUserProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserProcessingResultsById provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) UserProcessingResultsById(id string)(*i6ad171a2e2184a06bf3c8381d3823f8984897689a23d81f145ceddf152a6368d.UserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userProcessingResult%2Did"] = id
    }
    return i6ad171a2e2184a06bf3c8381d3823f8984897689a23d81f145ceddf152a6368d.NewUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) Versions()(*i72ce7d58e9852fd932435a826c49ca0ceb68d0fface3230435dcb902fc0e1ac0.VersionsRequestBuilder) {
    return i72ce7d58e9852fd932435a826c49ca0ceb68d0fface3230435dcb902fc0e1ac0.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *WorkflowItemRequestBuilder) VersionsById(id string)(*i6ef5d654ec18954956bcc3f06c9a25a2996d35ad34f5c9bedbdf3063fa0fc2e1.WorkflowVersionVersionNumberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workflowVersion%2DversionNumber"] = id
    }
    return i6ef5d654ec18954956bcc3f06c9a25a2996d35ad34f5c9bedbdf3063fa0fc2e1.NewWorkflowVersionVersionNumberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
