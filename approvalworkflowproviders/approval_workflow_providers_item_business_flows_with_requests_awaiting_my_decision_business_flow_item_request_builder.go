package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
type ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetQueryParameters get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
type ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetQueryParameters
}
// ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderInternal instantiates a new BusinessFlowItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) {
    m := &ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}/businessFlowsWithRequestsAwaitingMyDecision/{businessFlow%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder instantiates a new BusinessFlowItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property businessFlowsWithRequestsAwaitingMyDecision in approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, requestConfiguration *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable), nil
}
// Patch update the navigation property businessFlowsWithRequestsAwaitingMyDecision in approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, requestConfiguration *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable), nil
}
// RecordDecisions provides operations to call the recordDecisions method.
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) RecordDecisions()(*ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder) {
    return NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
