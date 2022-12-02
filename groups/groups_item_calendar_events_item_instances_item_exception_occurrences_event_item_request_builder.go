package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get exceptionOccurrences from groups
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*GroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
