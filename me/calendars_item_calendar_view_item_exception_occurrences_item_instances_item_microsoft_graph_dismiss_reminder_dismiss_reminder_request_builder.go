package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder provides operations to call the dismissReminder method.
type CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilderInternal instantiates a new DismissReminderRequestBuilder and sets the default values.
func NewCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder) {
    m := &CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/microsoft.graph.dismissReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder instantiates a new DismissReminderRequestBuilder and sets the default values.
func NewCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post dismiss a reminder that has been triggered for an event in a user calendar.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-dismissreminder?view=graph-rest-1.0
func (m *CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation dismiss a reminder that has been triggered for an event in a user calendar.
func (m *CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
