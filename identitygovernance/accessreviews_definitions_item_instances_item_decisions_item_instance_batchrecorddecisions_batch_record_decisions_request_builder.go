package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder provides operations to call the batchRecordDecisions method.
type AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/batchRecordDecisions", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-batchrecorddecisions?view=graph-rest-beta
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) Post(ctx context.Context, body AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsPostRequestBodyable, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsPostRequestBodyable, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchrecorddecisionsBatchRecordDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
