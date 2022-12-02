package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get exceptionOccurrences from me
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
