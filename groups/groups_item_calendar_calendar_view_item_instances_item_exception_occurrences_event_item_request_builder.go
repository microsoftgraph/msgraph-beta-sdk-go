package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get exceptionOccurrences from groups
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*GroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewGroupsItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
