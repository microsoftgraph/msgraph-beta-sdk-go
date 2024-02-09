package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal instantiates a new ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    m := &ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/snoozeReminder", pathParameters),
    }
    return m
}
// NewItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder instantiates a new ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post postpone a reminder for an event in a user calendar until a new time.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-snoozereminder?view=graph-rest-1.0
func (m *ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) Post(ctx context.Context, body ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderPostRequestBodyable, requestConfiguration *ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation postpone a reminder for an event in a user calendar until a new time.
// returns a *RequestInformation when successful
func (m *ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderPostRequestBodyable, requestConfiguration *ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) WithUrl(rawUrl string)(*ItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
