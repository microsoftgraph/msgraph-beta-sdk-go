package accessreviews

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInstancesAccessReviewItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.accessReview entity.
type ItemInstancesAccessReviewItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstancesAccessReviewItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesAccessReviewItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemInstancesAccessReviewItemRequestBuilderGetQueryParameters the collection of access reviews instances past, present and future, if this object is a recurring access review.
type ItemInstancesAccessReviewItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInstancesAccessReviewItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesAccessReviewItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInstancesAccessReviewItemRequestBuilderGetQueryParameters
}
// ItemInstancesAccessReviewItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesAccessReviewItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyDecisions provides operations to call the applyDecisions method.
// returns a *ItemInstancesItemApplyDecisionsRequestBuilder when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) ApplyDecisions()(*ItemInstancesItemApplyDecisionsRequestBuilder) {
    return NewItemInstancesItemApplyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInstancesAccessReviewItemRequestBuilderInternal instantiates a new ItemInstancesAccessReviewItemRequestBuilder and sets the default values.
func NewItemInstancesAccessReviewItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesAccessReviewItemRequestBuilder) {
    m := &ItemInstancesAccessReviewItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemInstancesAccessReviewItemRequestBuilder instantiates a new ItemInstancesAccessReviewItemRequestBuilder and sets the default values.
func NewItemInstancesAccessReviewItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesAccessReviewItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstancesAccessReviewItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReview entity.
// returns a *ItemInstancesItemDecisionsRequestBuilder when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) Decisions()(*ItemInstancesItemDecisionsRequestBuilder) {
    return NewItemInstancesItemDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property instances for accessReviews
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInstancesAccessReviewItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of access reviews instances past, present and future, if this object is a recurring access review.
// returns a AccessReviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInstancesAccessReviewItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable), nil
}
// MyDecisions provides operations to manage the myDecisions property of the microsoft.graph.accessReview entity.
// returns a *ItemInstancesItemMyDecisionsRequestBuilder when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) MyDecisions()(*ItemInstancesItemMyDecisionsRequestBuilder) {
    return NewItemInstancesItemMyDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property instances in accessReviews
// returns a AccessReviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInstancesAccessReviewItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable), nil
}
// ResetDecisions provides operations to call the resetDecisions method.
// returns a *ItemInstancesItemResetDecisionsRequestBuilder when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) ResetDecisions()(*ItemInstancesItemResetDecisionsRequestBuilder) {
    return NewItemInstancesItemResetDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reviewers provides operations to manage the reviewers property of the microsoft.graph.accessReview entity.
// returns a *ItemInstancesItemReviewersRequestBuilder when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) Reviewers()(*ItemInstancesItemReviewersRequestBuilder) {
    return NewItemInstancesItemReviewersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendReminder provides operations to call the sendReminder method.
// returns a *ItemInstancesItemSendReminderRequestBuilder when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) SendReminder()(*ItemInstancesItemSendReminderRequestBuilder) {
    return NewItemInstancesItemSendReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *ItemInstancesItemStopRequestBuilder when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) Stop()(*ItemInstancesItemStopRequestBuilder) {
    return NewItemInstancesItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property instances for accessReviews
// returns a *RequestInformation when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of access reviews instances past, present and future, if this object is a recurring access review.
// returns a *RequestInformation when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property instances in accessReviews
// returns a *RequestInformation when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemInstancesAccessReviewItemRequestBuilder when successful
func (m *ItemInstancesAccessReviewItemRequestBuilder) WithUrl(rawUrl string)(*ItemInstancesAccessReviewItemRequestBuilder) {
    return NewItemInstancesAccessReviewItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
