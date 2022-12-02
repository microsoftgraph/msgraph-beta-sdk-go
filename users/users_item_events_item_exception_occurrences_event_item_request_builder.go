package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*UsersItemEventsItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*UsersItemEventsItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*UsersItemEventsItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewUsersItemEventsItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*UsersItemEventsItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*UsersItemEventsItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewUsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*UsersItemEventsItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*UsersItemEventsItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*UsersItemEventsItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*UsersItemEventsItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemEventsItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*UsersItemEventsItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*UsersItemEventsItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) InstancesById(id string)(*UsersItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return NewUsersItemEventsItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemEventsItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemEventsItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemEventsItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemEventsItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemEventsItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemEventsItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*UsersItemEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *UsersItemEventsItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*UsersItemEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewUsersItemEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
