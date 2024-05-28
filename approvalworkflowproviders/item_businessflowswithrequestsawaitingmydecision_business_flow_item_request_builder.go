package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder provides operations to manage the businessFlowsWithRequestsAwaitingMyDecision property of the microsoft.graph.approvalWorkflowProvider entity.
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderGetQueryParameters get businessFlowsWithRequestsAwaitingMyDecision from approvalWorkflowProviders
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderGetQueryParameters
}
// ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderInternal instantiates a new ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder and sets the default values.
func NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) {
    m := &ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}/businessFlowsWithRequestsAwaitingMyDecision/{businessFlow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder instantiates a new ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder and sets the default values.
func NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, error) {
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
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, error) {
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
// returns a *ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) RecordDecisions()(*ItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilder) {
    return NewItemBusinessflowswithrequestsawaitingmydecisionItemRecorddecisionsRecordDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property businessFlowsWithRequestsAwaitingMyDecision for approvalWorkflowProviders
// returns a *RequestInformation when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessFlowable, requestConfiguration *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder when successful
func (m *ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) WithUrl(rawUrl string)(*ItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder) {
    return NewItemBusinessflowswithrequestsawaitingmydecisionBusinessFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
