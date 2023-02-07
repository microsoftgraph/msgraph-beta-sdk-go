package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*EventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*EventsItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewEventsItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*EventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*EventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*EventsItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewEventsItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get exceptionOccurrences from me
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphAccept()(*EventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphAcceptRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphCancel()(*EventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDecline provides operations to call the decline method.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphDecline()(*EventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphDeclineRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDismissReminder provides operations to call the dismissReminder method.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphDismissReminder()(*EventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphDismissReminderRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphForward provides operations to call the forward method.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphForward()(*EventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphForwardRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphSnoozeReminder()(*EventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MicrosoftGraphTentativelyAccept()(*EventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemMicrosoftGraphTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*EventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*EventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*EventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*EventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from me
func (m *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
