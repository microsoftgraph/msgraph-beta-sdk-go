// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder provides operations to call the applyDecisions method.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilderInternal instantiates a new AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/applyDecisions", pathParameters),
    }
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder instantiates a new AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post apply review decisions on an accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled property is false in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be Completed to call this method.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-applydecisions?view=graph-rest-beta
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation apply review decisions on an accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled property is false in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be Completed to call this method.
// returns a *RequestInformation when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceApplyDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
