package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderGetQueryParameters get userProcessingResults from identityGovernance
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderInternal instantiates a new UserProcessingResultItemRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) {
    m := &IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/userProcessingResults/{userProcessingResult%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder instantiates a new UserProcessingResultItemRequestBuilder and sets the default values.
func NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userProcessingResults for identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get userProcessingResults from identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userProcessingResults in identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property userProcessingResults for identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get userProcessingResults from identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateUserProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable), nil
}
// Patch update the navigation property userProcessingResults in identityGovernance
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable, requestConfiguration *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateUserProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable), nil
}
// Subject provides operations to manage the subject property of the microsoft.graph.identityGovernance.userProcessingResult entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) Subject()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemSubjectRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskProcessingResults provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.userProcessingResult entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) TaskProcessingResults()(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsRequestBuilder) {
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskProcessingResultsById provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.userProcessingResult entity.
func (m *IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) TaskProcessingResultsById(id string)(*IdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskProcessingResult%2Did"] = id
    }
    return NewIdentityGovernanceLifecycleWorkflowsWorkflowsItemUserProcessingResultsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
