package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetQueryParameters get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetQueryParameters
}
// ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderInternal instantiates a new ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) {
    m := &ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}/businessFlowsWithRequestsAwaitingMyDecision/{businessFlow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder instantiates a new ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
// returns a BusinessFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable), nil
}
// Patch update the navigation property businessFlowsWithRequestsAwaitingMyDecision in approvalWorkflowProviders
// returns a BusinessFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable), nil
}
// RecordDecisions provides operations to call the recordDecisions method.
// returns a *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder when successful
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) RecordDecisions()(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder) {
    return NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
// returns a *RequestInformation when successful
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
// returns a *RequestInformation when successful
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property businessFlowsWithRequestsAwaitingMyDecision in approvalWorkflowProviders
// returns a *RequestInformation when successful
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder when successful
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) WithUrl(rawUrl string)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder) {
    return NewItemBusinessFlowsWithRequestsAwaitingMyDecisionBusinessFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
