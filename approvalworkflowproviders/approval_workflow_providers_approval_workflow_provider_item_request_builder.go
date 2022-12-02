package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder provides operations to manage the collection of approvalWorkflowProvider entities.
type ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderGetQueryParameters get entity from approvalWorkflowProviders by key
type ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderGetQueryParameters
}
// ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BusinessFlows provides operations to manage the businessFlows property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) BusinessFlows()(*ApprovalWorkflowProvidersItemBusinessFlowsRequestBuilder) {
    return NewApprovalWorkflowProvidersItemBusinessFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowsById provides operations to manage the businessFlows property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsById(id string)(*ApprovalWorkflowProvidersItemBusinessFlowsBusinessFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlow%2Did"] = id
    }
    return NewApprovalWorkflowProvidersItemBusinessFlowsBusinessFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BusinessFlowsWithRequestsAwaitingMyDecision provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecision()(*ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    return NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowsWithRequestsAwaitingMyDecisionById provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecisionById(id string)(*ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlow%2Did"] = id
    }
    return NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderInternal instantiates a new ApprovalWorkflowProviderItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) {
    m := &ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder instantiates a new ApprovalWorkflowProviderItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from approvalWorkflowProviders by key
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, requestConfiguration *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get entity from approvalWorkflowProviders by key
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalWorkflowProviderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable), nil
}
// Patch update entity in approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, requestConfiguration *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalWorkflowProviderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable), nil
}
// PolicyTemplates provides operations to manage the policyTemplates property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) PolicyTemplates()(*ApprovalWorkflowProvidersItemPolicyTemplatesRequestBuilder) {
    return NewApprovalWorkflowProvidersItemPolicyTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PolicyTemplatesById provides operations to manage the policyTemplates property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProvidersApprovalWorkflowProviderItemRequestBuilder) PolicyTemplatesById(id string)(*ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governancePolicyTemplate%2Did"] = id
    }
    return NewApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
