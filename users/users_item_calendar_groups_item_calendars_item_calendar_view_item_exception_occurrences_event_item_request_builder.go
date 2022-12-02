package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) InstancesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*UsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewUsersItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
