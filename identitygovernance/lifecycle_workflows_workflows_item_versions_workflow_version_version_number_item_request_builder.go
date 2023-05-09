package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters the workflow versions that are available.
type LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal instantiates a new WorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder instantiates a new WorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedBy provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) CreatedBy()(*LifecycleWorkflowsWorkflowsItemVersionsItemCreatedByRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemCreatedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the workflow versions that are available.
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.WorkflowVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateWorkflowVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.WorkflowVersionable), nil
}
// LastModifiedBy provides operations to manage the lastModifiedBy property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) LastModifiedBy()(*LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Tasks()(*LifecycleWorkflowsWorkflowsItemVersionsItemTasksRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the workflow versions that are available.
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
