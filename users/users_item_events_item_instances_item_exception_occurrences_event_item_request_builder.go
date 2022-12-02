package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *UsersItemEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*UsersItemEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewUsersItemEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
