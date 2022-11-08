package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflows"
    i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflowswithrequestsawaitingmydecision"
    i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/policytemplates"
    iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/policytemplates/item"
    icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflows/item"
    id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflowswithrequestsawaitingmydecision/item"
)

// ApprovalWorkflowProviderItemRequestBuilder provides operations to manage the collection of approvalWorkflowProvider entities.
type ApprovalWorkflowProviderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApprovalWorkflowProviderItemRequestBuilderGetQueryParameters get entity from approvalWorkflowProviders by key
type ApprovalWorkflowProviderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApprovalWorkflowProviderItemRequestBuilderGetQueryParameters
}
// ApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BusinessFlows provides operations to manage the businessFlows property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProviderItemRequestBuilder) BusinessFlows()(*i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2.BusinessFlowsRequestBuilder) {
    return i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2.NewBusinessFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowsById provides operations to manage the businessFlows property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsById(id string)(*icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7.BusinessFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlow%2Did"] = id
    }
    return icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7.NewBusinessFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BusinessFlowsWithRequestsAwaitingMyDecision provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecision()(*i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532.BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    return i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532.NewBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowsWithRequestsAwaitingMyDecisionById provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecisionById(id string)(*id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28.BusinessFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlow%2Did"] = id
    }
    return id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28.NewBusinessFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewApprovalWorkflowProviderItemRequestBuilderInternal instantiates a new ApprovalWorkflowProviderItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProviderItemRequestBuilder) {
    m := &ApprovalWorkflowProviderItemRequestBuilder{
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
// NewApprovalWorkflowProviderItemRequestBuilder instantiates a new ApprovalWorkflowProviderItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalWorkflowProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ApprovalWorkflowProviderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ApprovalWorkflowProviderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, error) {
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
func (m *ApprovalWorkflowProviderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, error) {
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
func (m *ApprovalWorkflowProviderItemRequestBuilder) PolicyTemplates()(*i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107.PolicyTemplatesRequestBuilder) {
    return i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107.NewPolicyTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PolicyTemplatesById provides operations to manage the policyTemplates property of the microsoft.graph.approvalWorkflowProvider entity.
func (m *ApprovalWorkflowProviderItemRequestBuilder) PolicyTemplatesById(id string)(*iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa.GovernancePolicyTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governancePolicyTemplate%2Did"] = id
    }
    return iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa.NewGovernancePolicyTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
