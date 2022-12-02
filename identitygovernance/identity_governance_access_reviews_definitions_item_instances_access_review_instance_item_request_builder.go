package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.accessReviewScheduleDefinition entity.
type IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
type IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemAcceptRecommendationsRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions provides operations to call the applyDecisions method.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemApplyDecisionsRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemBatchRecordDecisionsRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) {
    m := &IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemContactedReviewersRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewersById(id string)(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemContactedReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemContactedReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property instances for identityGovernance
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property instances in identityGovernance
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Decisions()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) DecisionsById(id string)(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did"] = id
    }
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Definition()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDefinitionRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for identityGovernance
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// Patch update the navigation property instances in identityGovernance
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// ResetDecisions provides operations to call the resetDecisions method.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) ResetDecisions()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemResetDecisionsRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder provides operations to call the sendReminder method.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) SendReminder()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemSendReminderRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Stages()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemStagesRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) StagesById(id string)(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemStagesAccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemStagesAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesAccessReviewInstanceItemRequestBuilder) Stop()(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemStopRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
