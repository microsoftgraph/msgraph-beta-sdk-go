package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder provides operations to call the sendReminder method.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilderInternal instantiates a new AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/sendReminder", pathParameters),
    }
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder instantiates a new AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send a reminder to the reviewers of a currently active accessReviewInstance.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-sendreminder?view=graph-rest-beta
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send a reminder to the reviewers of a currently active accessReviewInstance.
// returns a *RequestInformation when successful
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceSendReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
