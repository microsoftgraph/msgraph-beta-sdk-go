package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Accept()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Cancel()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCancelRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    m := &MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline provides operations to call the decline method.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Decline()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) DismissReminder()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Forward()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SnoozeReminder()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) TentativelyAccept()(*MeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder) {
    return NewMeCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
