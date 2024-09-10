package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
type ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters navigation property to get a list of access reviews pending approval by the reviewer.
type ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetQueryParameters
}
// ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
// returns a *ItemPendingAccessReviewInstancesItemAcceptRecommendationsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*ItemPendingAccessReviewInstancesItemAcceptRecommendationsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemAcceptRecommendationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplyDecisions provides operations to call the applyDecisions method.
// returns a *ItemPendingAccessReviewInstancesItemApplyDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*ItemPendingAccessReviewInstancesItemApplyDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemApplyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
// returns a *ItemPendingAccessReviewInstancesItemBatchRecordDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*ItemPendingAccessReviewInstancesItemBatchRecordDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemBatchRecordDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal instantiates a new ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    m := &ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder instantiates a new ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingAccessReviewInstancesItemContactedReviewersRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*ItemPendingAccessReviewInstancesItemContactedReviewersRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemContactedReviewersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Decisions()(*ItemPendingAccessReviewInstancesItemDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Definition()(*ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property pendingAccessReviewInstances for users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get navigation property to get a list of access reviews pending approval by the reviewer.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a AccessReviewInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// Patch update the navigation property pendingAccessReviewInstances in users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a AccessReviewInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// returns a *ItemPendingAccessReviewInstancesItemResetDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ResetDecisions()(*ItemPendingAccessReviewInstancesItemResetDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemResetDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendReminder provides operations to call the sendReminder method.
// returns a *ItemPendingAccessReviewInstancesItemSendReminderRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) SendReminder()(*ItemPendingAccessReviewInstancesItemSendReminderRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemSendReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stages provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
// returns a *ItemPendingAccessReviewInstancesItemStagesRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Stages()(*ItemPendingAccessReviewInstancesItemStagesRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemStagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *ItemPendingAccessReviewInstancesItemStopRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) Stop()(*ItemPendingAccessReviewInstancesItemStopRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StopApplyDecisions provides operations to call the stopApplyDecisions method.
// returns a *ItemPendingAccessReviewInstancesItemStopApplyDecisionsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) StopApplyDecisions()(*ItemPendingAccessReviewInstancesItemStopApplyDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemStopApplyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property pendingAccessReviewInstances for users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation navigation property to get a list of access reviews pending approval by the reviewer.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property pendingAccessReviewInstances in users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) WithUrl(rawUrl string)(*ItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    return NewItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
