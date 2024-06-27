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
// LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters read the properties and relationships of a workflowVersion object.
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
// NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal instantiates a new LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder instantiates a new LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedBy provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleWorkflowsWorkflowsItemVersionsItemCreatedByRequestBuilder when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) CreatedBy()(*LifecycleWorkflowsWorkflowsItemVersionsItemCreatedByRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemCreatedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a workflowVersion object.
// returns a WorkflowVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflowversion-get?view=graph-rest-beta
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.WorkflowVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) LastModifiedBy()(*LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleWorkflowsWorkflowsItemVersionsItemTasksRequestBuilder when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Tasks()(*LifecycleWorkflowsWorkflowsItemVersionsItemTasksRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read the properties and relationships of a workflowVersion object.
// returns a *RequestInformation when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
