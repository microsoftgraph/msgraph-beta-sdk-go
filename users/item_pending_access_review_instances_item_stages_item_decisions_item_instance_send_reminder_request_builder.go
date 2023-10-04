package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder provides operations to call the sendReminder method.
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal instantiates a new SendReminderRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    m := &ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/sendReminder", pathParameters),
    }
    return m
}
// NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder instantiates a new SendReminderRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send a reminder to the reviewers of a currently active accessReviewInstance. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-sendreminder?view=graph-rest-1.0
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation send a reminder to the reviewers of a currently active accessReviewInstance. This API is supported in the following national cloud deployments.
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) WithUrl(rawUrl string)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
