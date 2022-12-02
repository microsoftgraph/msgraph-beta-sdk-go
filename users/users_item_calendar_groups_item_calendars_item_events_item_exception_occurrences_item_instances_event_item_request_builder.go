package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Accept()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Cancel()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    m := &UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Decline()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) DismissReminder()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Forward()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SnoozeReminder()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) TentativelyAccept()(*UsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
