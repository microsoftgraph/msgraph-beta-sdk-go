package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder provides operations to call the resetDecisions method.
type AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilderInternal instantiates a new AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder) {
    m := &AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/resetDecisions", pathParameters),
    }
    return m
}
// NewAccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder instantiates a new AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post resets decisions of all accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-resetdecisions?view=graph-rest-beta
func (m *AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation resets decisions of all accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
// returns a *RequestInformation when successful
func (m *AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder when successful
func (m *AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder) {
    return NewAccessReviewsDecisionsItemInstanceResetDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
