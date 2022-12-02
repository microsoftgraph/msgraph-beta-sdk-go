package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    m := &MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/microsoft.graph.snoozeReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation postpone a reminder for an event in a user calendar until a new time.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderPostRequestBodyable, requestConfiguration *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post postpone a reminder for an event in a user calendar until a new time.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderPostRequestBodyable, requestConfiguration *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
