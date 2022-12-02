package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Accept()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemAcceptRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Attachments()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) AttachmentsById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Calendar()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemCalendarRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Cancel()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    m := &UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Decline()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemDeclineRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) DismissReminder()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemDismissReminderRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) ExceptionOccurrences()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) ExceptionOccurrencesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Extensions()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) ExtensionsById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Forward()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemForwardRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) Instances()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) InstancesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) SnoozeReminder()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewEventItemRequestBuilder) TentativelyAccept()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemTentativelyAcceptRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
