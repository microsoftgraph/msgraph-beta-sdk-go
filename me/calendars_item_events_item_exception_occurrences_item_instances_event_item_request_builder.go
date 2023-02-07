package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    m := &CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphAccept()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphAcceptRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphCancel()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphDecline()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDeclineRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphForward()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphForwardRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
