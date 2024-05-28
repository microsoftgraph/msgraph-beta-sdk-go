package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder provides operations to manage the instance property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters there's exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters
}
// ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) AcceptRecommendations()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceAcceptrecommendationsAcceptRecommendationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyDecisions provides operations to call the applyDecisions method.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceApplydecisionsApplyDecisionsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) ApplyDecisions()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceApplydecisionsApplyDecisionsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceApplydecisionsApplyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) BatchRecordDecisions()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderInternal instantiates a new ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) {
    m := &ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder instantiates a new ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceContactedreviewersContactedReviewersRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) ContactedReviewers()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceContactedreviewersContactedReviewersRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceContactedreviewersContactedReviewersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceDefinitionRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) Definition()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property instance for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get there's exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
// returns a AccessReviewInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// Patch update the navigation property instance in users
// returns a AccessReviewInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// ResetDecisions provides operations to call the resetDecisions method.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceResetdecisionsResetDecisionsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) ResetDecisions()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceResetdecisionsResetDecisionsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceResetdecisionsResetDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendReminder provides operations to call the sendReminder method.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceSendreminderSendReminderRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) SendReminder()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceSendreminderSendReminderRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceSendreminderSendReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) Stages()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceStagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStopRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) Stop()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStopRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StopApplyDecisions provides operations to call the stopApplyDecisions method.
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) StopApplyDecisions()(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceStopapplydecisionsStopApplyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property instance for users
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation there's exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property instance in users
// returns a *RequestInformation when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) WithUrl(rawUrl string)(*ItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemDecisionsItemInstanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
