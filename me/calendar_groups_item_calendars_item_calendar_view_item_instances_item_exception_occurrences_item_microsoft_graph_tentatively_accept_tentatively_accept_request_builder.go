package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder provides operations to call the tentativelyAccept method.
type CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilderInternal instantiates a new TentativelyAcceptRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/microsoft.graph.tentativelyAccept";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder instantiates a new TentativelyAcceptRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post tentatively accept the specified event in a user calendar. If the event allows proposals for new times, on responding tentative to the event, an invitee can choose to suggest an alternative time by including the **proposedNewTime** parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-tentativelyaccept?view=graph-rest-1.0
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder) Post(ctx context.Context, body CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBodyable, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation tentatively accept the specified event in a user calendar. If the event allows proposals for new times, on responding tentative to the event, an invitee can choose to suggest an alternative time by including the **proposedNewTime** parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder) ToPostRequestInformation(ctx context.Context, body CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptPostRequestBodyable, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
