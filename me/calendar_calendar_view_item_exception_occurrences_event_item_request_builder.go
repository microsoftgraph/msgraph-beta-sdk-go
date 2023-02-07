package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*CalendarCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*CalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*CalendarCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*CalendarCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*CalendarCalendarViewItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewCalendarCalendarViewItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get exceptionOccurrences from me
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*CalendarCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) InstancesById(id string)(*CalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return NewCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphAccept provides operations to call the accept method.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphAccept()(*CalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphAcceptRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphCancel()(*CalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphCancelRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphDecline()(*CalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphDeclineRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*CalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphForward()(*CalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphForwardRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*CalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*CalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*CalendarCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*CalendarCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*CalendarCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewCalendarCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*CalendarCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewCalendarCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from me
func (m *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
