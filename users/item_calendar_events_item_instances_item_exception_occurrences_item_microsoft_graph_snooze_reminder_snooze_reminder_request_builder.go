package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderInternal instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) {
    m := &ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/microsoft.graph.snoozeReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post postpone a reminder for an event in a user calendar until a new time.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-snoozereminder?view=graph-rest-1.0
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) Post(ctx context.Context, body ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
