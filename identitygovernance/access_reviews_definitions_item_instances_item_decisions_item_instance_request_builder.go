package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder provides operations to manage the instance property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder instantiates a new InstanceRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) ContactedReviewers()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceContactedReviewersRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) ContactedReviewersById(id string)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceContactedReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceContactedReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) Definition()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property instance for identityGovernance
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphAcceptRecommendations()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphAcceptRecommendationsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApplyDecisions provides operations to call the applyDecisions method.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphApplyDecisions()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphApplyDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphBatchRecordDecisions()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetDecisions provides operations to call the resetDecisions method.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphResetDecisions()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphResetDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendReminder provides operations to call the sendReminder method.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphSendReminder()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphSendReminderRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphStop provides operations to call the stop method.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) MicrosoftGraphStop()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphStopRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphStopRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property instance in identityGovernance
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) Stages()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// StagesById provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) StagesById(id string)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesAccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property instance for identityGovernance
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
