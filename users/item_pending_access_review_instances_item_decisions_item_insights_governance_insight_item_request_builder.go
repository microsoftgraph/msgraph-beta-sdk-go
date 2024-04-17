package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder provides operations to manage the insights property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetQueryParameters insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
type ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetQueryParameters
}
// ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderInternal instantiates a new ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) {
    m := &ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/insights/{governanceInsight%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder instantiates a new ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property insights for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
// returns a GovernanceInsightable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceInsightFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable), nil
}
// Patch update the navigation property insights in users
// returns a GovernanceInsightable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceInsightFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable), nil
}
// ToDeleteRequestInformation delete navigation property insights for users
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property insights in users
// returns a *RequestInformation when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, requestConfiguration *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder when successful
func (m *ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) WithUrl(rawUrl string)(*ItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDecisionsItemInsightsGovernanceInsightItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
