package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal instantiates a new ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    m := &ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/snoozeReminder", pathParameters),
    }
    return m
}
// NewItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder instantiates a new ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post postpone a reminder for an event in a user calendar until a new time.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-snoozereminder?view=graph-rest-beta
func (m *ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) Post(ctx context.Context, body ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
