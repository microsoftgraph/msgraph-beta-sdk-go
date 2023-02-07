package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*ItemCalendarsItemCalendarViewItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarsItemCalendarViewItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) ExceptionOccurrences()(*ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return NewItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarsItemCalendarViewItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*ItemCalendarsItemCalendarViewItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphAccept()(*ItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphCancel()(*ItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphCancelRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphDecline()(*ItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDeclineRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*ItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphForward()(*ItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphForwardRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*ItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*ItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
