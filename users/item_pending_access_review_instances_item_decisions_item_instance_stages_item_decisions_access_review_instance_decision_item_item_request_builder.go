package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder provides operations to manage the decisions property of the microsoft.graph.accessReviewStage entity.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal instantiates a new ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    m := &ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder instantiates a new ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property decisions for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
// returns a AccessReviewInstanceDecisionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable), nil
}
// Insights provides operations to manage the insights property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsItemInsightsRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Insights()(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsItemInsightsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsItemInsightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property decisions in users
// returns a AccessReviewInstanceDecisionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable), nil
}
// ToDeleteRequestInformation delete navigation property decisions for users
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property decisions in users
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) WithUrl(rawUrl string)(*ItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInstanceStagesItemDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
