package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Accept()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemAcceptRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Attachments()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) AttachmentsById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Calendar()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemCalendarRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Cancel()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    m := &MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Decline()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemDeclineRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) DismissReminder()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemDismissReminderRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) ExceptionOccurrences()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) ExceptionOccurrencesById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Extensions()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) ExtensionsById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Forward()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemForwardRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Instances()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) InstancesById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedProperties()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedProperties()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) SnoozeReminder()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) TentativelyAccept()(*MeCalendarGroupsItemCalendarsItemCalendarViewItemTentativelyAcceptRequestBuilder) {
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
