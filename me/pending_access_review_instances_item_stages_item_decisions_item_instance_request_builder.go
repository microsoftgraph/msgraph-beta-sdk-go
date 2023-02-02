package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder provides operations to manage the instance property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderGetQueryParameters there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderGetQueryParameters
}
// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) {
    m := &PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder instantiates a new InstanceRequestBuilder and sets the default values.
func NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) ContactedReviewers()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceContactedReviewersRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) ContactedReviewersById(id string)(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceContactedReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceContactedReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) Decisions()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) DecisionsById(id string)(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did1"] = id
    }
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) Definition()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property instance for me
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) Delete(ctx context.Context, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) Get(ctx context.Context, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// MicrosoftGraphAcceptRecommendations provides operations to call the acceptRecommendations method.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphAcceptRecommendations()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphAcceptRecommendationsAcceptRecommendationsRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphAcceptRecommendationsAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyDecisions provides operations to call the applyDecisions method.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphApplyDecisions()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphApplyDecisionsApplyDecisionsRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphApplyDecisionsApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphBatchRecordDecisions()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetDecisions provides operations to call the resetDecisions method.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphResetDecisions()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphResetDecisionsResetDecisionsRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphResetDecisionsResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendReminder provides operations to call the sendReminder method.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphSendReminder()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphSendReminderSendReminderRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphSendReminderSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphStop provides operations to call the stop method.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphStop()(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphStopStopRequestBuilder) {
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceMicrosoftGraphStopStopRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property instance in me
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// ToDeleteRequestInformation delete navigation property instance for me
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the navigation property instance in me
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
