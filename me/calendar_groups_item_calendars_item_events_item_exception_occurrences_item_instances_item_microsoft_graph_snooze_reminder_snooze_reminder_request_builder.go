package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderInternal instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/microsoft.graph.snoozeReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post postpone a reminder for an event in a user calendar until a new time.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-snoozereminder?view=graph-rest-1.0
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) Post(ctx context.Context, body CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyable, requestConfiguration *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation postpone a reminder for an event in a user calendar until a new time.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyable, requestConfiguration *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
