package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get exceptionOccurrences from me
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) InstancesById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAccept provides operations to call the accept method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphAccept()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphAcceptRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphCancel()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphCancelRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphDecline()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphDeclineRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphForward()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphForwardRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from me
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
