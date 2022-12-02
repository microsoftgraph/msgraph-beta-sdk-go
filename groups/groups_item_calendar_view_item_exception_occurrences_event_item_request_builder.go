package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*GroupsItemCalendarViewItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*GroupsItemCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) AttachmentsById(id string)(*GroupsItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*GroupsItemCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*GroupsItemCalendarViewItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get exceptionOccurrences from groups
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*GroupsItemCalendarViewItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*GroupsItemCalendarViewItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*GroupsItemCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) ExtensionsById(id string)(*GroupsItemCalendarViewItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*GroupsItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*GroupsItemCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) InstancesById(id string)(*GroupsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedProperties()(*GroupsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*GroupsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedProperties()(*GroupsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*GroupsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*GroupsItemCalendarViewItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *GroupsItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*GroupsItemCalendarViewItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewGroupsItemCalendarViewItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
