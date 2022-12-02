package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemCalendarEventsEventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
type GroupsItemCalendarEventsEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemCalendarEventsEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemCalendarEventsEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemCalendarEventsEventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type GroupsItemCalendarEventsEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemCalendarEventsEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemCalendarEventsEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemCalendarEventsEventItemRequestBuilderGetQueryParameters
}
// GroupsItemCalendarEventsEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemCalendarEventsEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept provides operations to call the accept method.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Accept()(*GroupsItemCalendarEventsItemAcceptRequestBuilder) {
    return NewGroupsItemCalendarEventsItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Attachments()(*GroupsItemCalendarEventsItemAttachmentsRequestBuilder) {
    return NewGroupsItemCalendarEventsItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) AttachmentsById(id string)(*GroupsItemCalendarEventsItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewGroupsItemCalendarEventsItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Calendar()(*GroupsItemCalendarEventsItemCalendarRequestBuilder) {
    return NewGroupsItemCalendarEventsItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Cancel()(*GroupsItemCalendarEventsItemCancelRequestBuilder) {
    return NewGroupsItemCalendarEventsItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsItemCalendarEventsEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarEventsEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarEventsEventItemRequestBuilder) {
    m := &GroupsItemCalendarEventsEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemCalendarEventsEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewGroupsItemCalendarEventsEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemCalendarEventsEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemCalendarEventsEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property events for groups
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemCalendarEventsEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemCalendarEventsEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property events in groups
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *GroupsItemCalendarEventsEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline provides operations to call the decline method.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Decline()(*GroupsItemCalendarEventsItemDeclineRequestBuilder) {
    return NewGroupsItemCalendarEventsItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for groups
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemCalendarEventsEventItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) DismissReminder()(*GroupsItemCalendarEventsItemDismissReminderRequestBuilder) {
    return NewGroupsItemCalendarEventsItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) ExceptionOccurrences()(*GroupsItemCalendarEventsItemExceptionOccurrencesRequestBuilder) {
    return NewGroupsItemCalendarEventsItemExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) ExceptionOccurrencesById(id string)(*GroupsItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewGroupsItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Extensions()(*GroupsItemCalendarEventsItemExtensionsRequestBuilder) {
    return NewGroupsItemCalendarEventsItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) ExtensionsById(id string)(*GroupsItemCalendarEventsItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewGroupsItemCalendarEventsItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Forward()(*GroupsItemCalendarEventsItemForwardRequestBuilder) {
    return NewGroupsItemCalendarEventsItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemCalendarEventsEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Instances()(*GroupsItemCalendarEventsItemInstancesRequestBuilder) {
    return NewGroupsItemCalendarEventsItemInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) InstancesById(id string)(*GroupsItemCalendarEventsItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return NewGroupsItemCalendarEventsItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) MultiValueExtendedProperties()(*GroupsItemCalendarEventsItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemCalendarEventsItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*GroupsItemCalendarEventsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemCalendarEventsItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in groups
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *GroupsItemCalendarEventsEventItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) SingleValueExtendedProperties()(*GroupsItemCalendarEventsItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewGroupsItemCalendarEventsItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*GroupsItemCalendarEventsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewGroupsItemCalendarEventsItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) SnoozeReminder()(*GroupsItemCalendarEventsItemSnoozeReminderRequestBuilder) {
    return NewGroupsItemCalendarEventsItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *GroupsItemCalendarEventsEventItemRequestBuilder) TentativelyAccept()(*GroupsItemCalendarEventsItemTentativelyAcceptRequestBuilder) {
    return NewGroupsItemCalendarEventsItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
