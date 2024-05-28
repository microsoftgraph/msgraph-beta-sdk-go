package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder provides operations to call the stop method.
type AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilderInternal instantiates a new AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder and sets the default values.
func NewAccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder) {
    m := &AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/stages/{accessReviewStage%2Did}/stop", pathParameters),
    }
    return m
}
// NewAccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder instantiates a new AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder and sets the default values.
func NewAccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilderInternal(urlParams, requestAdapter)
}
// Post stop an access review stage that is inProgress. After the access review stage stops, the stage status will be Completed and the reviewers can no longer give input. If there are subsequent stages that depend on the completed stage, the next stage will be created.  The accessReviewInstanceDecisionItem objects will always reflect the last decisions recorded across all stages at that given time, regardless of the status of the stages.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewstage-stop?view=graph-rest-beta
func (m *AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation stop an access review stage that is inProgress. After the access review stage stops, the stage status will be Completed and the reviewers can no longer give input. If there are subsequent stages that depend on the completed stage, the next stage will be created.  The accessReviewInstanceDecisionItem objects will always reflect the last decisions recorded across all stages at that given time, regardless of the status of the stages.
// returns a *RequestInformation when successful
func (m *AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder when successful
func (m *AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder) {
    return NewAccessreviewsDecisionsItemInstanceStagesItemStopRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
