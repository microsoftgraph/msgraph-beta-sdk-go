package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder provides operations to call the batchRecordDecisions method.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderInternal instantiates a new BatchRecordDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/batchRecordDecisions", pathParameters),
    }
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder instantiates a new BatchRecordDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-batchrecorddecisions?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder) Post(ctx context.Context, body AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsPostRequestBodyable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsPostRequestBodyable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
