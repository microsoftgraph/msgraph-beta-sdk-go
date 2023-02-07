package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder provides operations to call the getSchedule method.
type CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderInternal instantiates a new MicrosoftGraphGetScheduleRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/microsoft.graph.getSchedule";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder instantiates a new MicrosoftGraphGetScheduleRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the free/busy availability information for a collection of users, distributions lists, or resources (rooms or equipment) for a specified time period.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/calendar-getschedule?view=graph-rest-1.0
func (m *CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder) Post(ctx context.Context, body CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetSchedulePostRequestBodyable, requestConfiguration *CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderPostRequestConfiguration)(CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponseable), nil
}
// ToPostRequestInformation get the free/busy availability information for a collection of users, distributions lists, or resources (rooms or equipment) for a specified time period.
func (m *CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilder) ToPostRequestInformation(ctx context.Context, body CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetSchedulePostRequestBodyable, requestConfiguration *CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
