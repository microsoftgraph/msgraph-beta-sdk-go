package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder provides operations to call the dismissReminder method.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderInternal instantiates a new ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/dismissReminder", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder instantiates a new ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post dismiss a reminder that has been triggered for an event in a user calendar.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-dismissreminder?view=graph-rest-beta
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation dismiss a reminder that has been triggered for an event in a user calendar.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
