package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDecisionsItemInstanceRequestBuilder provides operations to manage the instance property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type AccessReviewsDecisionsItemInstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewsDecisionsItemInstanceRequestBuilderGetQueryParameters there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type AccessReviewsDecisionsItemInstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewsDecisionsItemInstanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDecisionsItemInstanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDecisionsItemInstanceRequestBuilderGetQueryParameters
}
// AccessReviewsDecisionsItemInstanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDecisionsItemInstanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) AcceptRecommendations()(*AccessReviewsDecisionsItemInstanceAcceptRecommendationsRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplyDecisions provides operations to call the applyDecisions method.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) ApplyDecisions()(*AccessReviewsDecisionsItemInstanceApplyDecisionsRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) BatchRecordDecisions()(*AccessReviewsDecisionsItemInstanceBatchRecordDecisionsRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewAccessReviewsDecisionsItemInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewAccessReviewsDecisionsItemInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDecisionsItemInstanceRequestBuilder) {
    m := &AccessReviewsDecisionsItemInstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAccessReviewsDecisionsItemInstanceRequestBuilder instantiates a new InstanceRequestBuilder and sets the default values.
func NewAccessReviewsDecisionsItemInstanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDecisionsItemInstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDecisionsItemInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) ContactedReviewers()(*AccessReviewsDecisionsItemInstanceContactedReviewersRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) ContactedReviewersById(id string)(*AccessReviewsDecisionsItemInstanceContactedReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewAccessReviewsDecisionsItemInstanceContactedReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) Decisions()(*AccessReviewsDecisionsItemInstanceDecisionsRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) DecisionsById(id string)(*AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did1"] = id
    }
    return NewAccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) Definition()(*AccessReviewsDecisionsItemInstanceDefinitionRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property instance for identityGovernance
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// Patch update the navigation property instance in identityGovernance
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *AccessReviewsDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// ResetDecisions provides operations to call the resetDecisions method.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) ResetDecisions()(*AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SendReminder provides operations to call the sendReminder method.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) SendReminder()(*AccessReviewsDecisionsItemInstanceSendReminderRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) Stages()(*AccessReviewsDecisionsItemInstanceStagesRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// StagesById provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) StagesById(id string)(*AccessReviewsDecisionsItemInstanceStagesAccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return NewAccessReviewsDecisionsItemInstanceStagesAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Stop provides operations to call the stop method.
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) Stop()(*AccessReviewsDecisionsItemInstanceStopRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceStopRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property instance for identityGovernance
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property instance in identityGovernance
func (m *AccessReviewsDecisionsItemInstanceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *AccessReviewsDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
