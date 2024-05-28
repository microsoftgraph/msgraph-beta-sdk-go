package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder provides operations to call the batchRecordDecisions method.
type ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal instantiates a new ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    m := &ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/batchRecordDecisions", pathParameters),
    }
    return m
}
// NewItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder instantiates a new ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder and sets the default values.
func NewItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-batchrecorddecisions?view=graph-rest-beta
func (m *ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) Post(ctx context.Context, body ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsPostRequestBodyable, requestConfiguration *ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsPostRequestBodyable, requestConfiguration *ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder when successful
func (m *ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) WithUrl(rawUrl string)(*ItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder) {
    return NewItemPendingaccessreviewinstancesItemBatchrecorddecisionsBatchRecordDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
