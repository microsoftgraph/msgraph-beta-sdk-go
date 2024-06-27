package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder provides operations to manage the insights property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetQueryParameters insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetQueryParameters
}
// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderInternal instantiates a new AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions/{accessReviewInstanceDecisionItem%2Did1}/insights/{governanceInsight%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder instantiates a new AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property insights for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, error) {
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
// Patch update the navigation property insights in identityGovernance
// returns a GovernanceInsightable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, error) {
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
// ToDeleteRequestInformation delete navigation property insights for identityGovernance
// returns a *RequestInformation when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property insights in identityGovernance
// returns a *RequestInformation when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
