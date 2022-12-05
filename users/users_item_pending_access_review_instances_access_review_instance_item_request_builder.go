package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
type UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters navigation property to get list of access reviews pending approval by reviewer.
type UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters
}
// UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*UsersItemPendingAccessReviewInstancesItemAcceptRecommendationsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions provides operations to call the applyDecisions method.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*UsersItemPendingAccessReviewInstancesItemApplyDecisionsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*UsersItemPendingAccessReviewInstancesItemBatchRecordDecisionsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    m := &UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*UsersItemPendingAccessReviewInstancesItemContactedReviewersRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewersById(id string)(*UsersItemPendingAccessReviewInstancesItemContactedReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewUsersItemPendingAccessReviewInstancesItemContactedReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pendingAccessReviewInstances for users
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property pendingAccessReviewInstances in users
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Decisions()(*UsersItemPendingAccessReviewInstancesItemDecisionsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) DecisionsById(id string)(*UsersItemPendingAccessReviewInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did"] = id
    }
    return NewUsersItemPendingAccessReviewInstancesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Definition()(*UsersItemPendingAccessReviewInstancesItemDefinitionRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property pendingAccessReviewInstances for users
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// Patch update the navigation property pendingAccessReviewInstances in users
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ResetDecisions()(*UsersItemPendingAccessReviewInstancesItemResetDecisionsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder provides operations to call the sendReminder method.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) SendReminder()(*UsersItemPendingAccessReviewInstancesItemSendReminderRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Stages()(*UsersItemPendingAccessReviewInstancesItemStagesRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) StagesById(id string)(*UsersItemPendingAccessReviewInstancesItemStagesAccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return NewUsersItemPendingAccessReviewInstancesItemStagesAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Stop()(*UsersItemPendingAccessReviewInstancesItemStopRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
