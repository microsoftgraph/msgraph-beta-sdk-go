package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
type AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters read the properties and relationships of an accessReviewInstanceDecisionItem object. This API is available in the following national cloud deployments.
type AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters
}
// AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal instantiates a new AccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewAccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    m := &AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions/{accessReviewInstanceDecisionItem%2Did1}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder instantiates a new AccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewAccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property decisions for identityGovernance
func (m *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of an accessReviewInstanceDecisionItem object. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstancedecisionitem-get?view=graph-rest-1.0
func (m *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Insights()(*AccessReviewsDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceDecisionsItemInsightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property decisions in identityGovernance
func (m *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete navigation property decisions for identityGovernance
func (m *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an accessReviewInstanceDecisionItem object. This API is available in the following national cloud deployments.
func (m *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property decisions in identityGovernance
func (m *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
