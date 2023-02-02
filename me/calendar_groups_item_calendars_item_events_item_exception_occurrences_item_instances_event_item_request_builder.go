package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, eventId2 *string)(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if eventId2 != nil {
        urlTplParams["event%2Did2"] = *eventId2
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MicrosoftGraphAccept provides operations to call the accept method.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphAccept()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphAcceptAcceptRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphAcceptAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphCancel()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphDecline()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDeclineDeclineRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDeclineDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphForward()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphForwardForwardRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphForwardForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphTentativelyAcceptTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
