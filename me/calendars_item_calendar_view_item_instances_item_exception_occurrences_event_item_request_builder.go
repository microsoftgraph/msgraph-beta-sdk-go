package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get exceptionOccurrences from me
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphAccept()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphAcceptRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphCancel()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphDecline()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphDeclineRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphForward()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphForwardRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from me
func (m *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
