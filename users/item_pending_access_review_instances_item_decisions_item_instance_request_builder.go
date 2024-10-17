package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder provides operations to manage the instance property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters there's exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) AcceptRecommendations()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyDecisions provides operations to call the applyDecisions method.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ApplyDecisions()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) BatchRecordDecisions()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderInternal instantiates a new ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) {
    m := &ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder instantiates a new ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceContactedReviewersRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ContactedReviewers()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceContactedReviewersRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceContactedReviewersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceDefinitionRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Definition()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property instance for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceResetDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ResetDecisions()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceResetDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceResetDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendReminder provides operations to call the sendReminder method.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceSendReminderRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) SendReminder()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Stages()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStopRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Stop()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStopRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StopApplyDecisions provides operations to call the stopApplyDecisions method.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStopApplyDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) StopApplyDecisions()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStopApplyDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStopApplyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property instance for users
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) WithUrl(rawUrl string)(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
