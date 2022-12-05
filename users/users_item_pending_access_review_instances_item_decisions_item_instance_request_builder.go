package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder provides operations to manage the instance property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetQueryParameters
}
// UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) AcceptRecommendations()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions provides operations to call the applyDecisions method.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ApplyDecisions()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) BatchRecordDecisions()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) {
    m := &UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder instantiates a new InstanceRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ContactedReviewers()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceContactedReviewersRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ContactedReviewersById(id string)(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceContactedReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceContactedReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property instance for users
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property instance in users
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Definition()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instance for users
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// Patch update the navigation property instance in users
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) ResetDecisions()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceResetDecisionsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder provides operations to call the sendReminder method.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) SendReminder()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Stages()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) StagesById(id string)(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesAccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceRequestBuilder) Stop()(*UsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceStopRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsItemInstanceStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
