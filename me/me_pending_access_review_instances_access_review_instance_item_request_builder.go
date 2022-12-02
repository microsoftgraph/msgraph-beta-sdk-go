package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
type MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters navigation property to get list of access reviews pending approval by reviewer.
type MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters
}
// MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*MePendingAccessReviewInstancesItemAcceptRecommendationsRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions provides operations to call the applyDecisions method.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*MePendingAccessReviewInstancesItemApplyDecisionsRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*MePendingAccessReviewInstancesItemBatchRecordDecisionsRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewMePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    m := &MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewMePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*MePendingAccessReviewInstancesItemContactedReviewersRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewersById(id string)(*MePendingAccessReviewInstancesItemContactedReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewMePendingAccessReviewInstancesItemContactedReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pendingAccessReviewInstances for me
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation navigation property to get list of access reviews pending approval by reviewer.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property pendingAccessReviewInstances in me
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Decisions()(*MePendingAccessReviewInstancesItemDecisionsRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) DecisionsById(id string)(*MePendingAccessReviewInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did"] = id
    }
    return NewMePendingAccessReviewInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Definition()(*MePendingAccessReviewInstancesItemDefinitionRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property pendingAccessReviewInstances for me
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get navigation property to get list of access reviews pending approval by reviewer.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// Patch update the navigation property pendingAccessReviewInstances in me
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ResetDecisions()(*MePendingAccessReviewInstancesItemResetDecisionsRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder provides operations to call the sendReminder method.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) SendReminder()(*MePendingAccessReviewInstancesItemSendReminderRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Stages()(*MePendingAccessReviewInstancesItemStagesRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) StagesById(id string)(*MePendingAccessReviewInstancesItemStagesAccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return NewMePendingAccessReviewInstancesItemStagesAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Stop()(*MePendingAccessReviewInstancesItemStopRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
