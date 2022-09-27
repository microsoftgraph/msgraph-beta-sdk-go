package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0f1a586e02c48d8d20c79097f007f7e47f32f69f0042e6340415f3edc6368e09 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/accept"
    i0f6786d0f015f4a9e6d15f9a02d084210f77652248ff35600e6b2052e34cb8ee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/decline"
    i39da92daed46a78d06088ff7d64bd54ec8ee825fb267cf152af8d81a9a07d14d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/calendar"
    i4402a65cd88371352fb2733c27e5fc14fdb4a77ddd1ee6c058800688fa0c6ff4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/attachments"
    i49ba1c423c143b602af84628f15aaa8269795950d34ec0e5712f027635ba0954 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i6993f9ef530b7249eff9bd6e68ee2c554fb313c4dbf671629445fc6a821307c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/cancel"
    i7ca9fbc65ff1e866aeea0c0b2d05b34a9105a9e06903e82afebb6b87f81ad82e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ia2da1a05683f51fd9cd6900829f4d7a01708df0a55d0566b0babc2bf22f536d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/extensions"
    ia657a65517b4ceccdbbca50613f1d2362e14464ca9e58ae586c6d39acb900681 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    ibbc7dae07f158c8d9b1b2b21a19441045bcfaad92bac0dee00d39ceb37ba4955 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    id2b798642d244212d8e1e2d9927afe32820c6f12099c1202597cb0eef52c5544 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    ie8736a7b60e1ca004d2a30406df2d191282e92c1a70b23210d3c8f6ac1e6b89e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/forward"
    i271633dd931e6f83b538cf85551d51d635881acad169f9f55964748be82ddf45 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i28d929187c18d2d0c35f8165ba26a53887dd7d6395d44306172cf9211803d034 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    ica8d5645131e73b6410654bc7f30c1863daa7b066d19b64654d1a4d1ecd344cb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    id55f41e90eb1b5dd9af679cc1b053c8dbbaaa6c2944925b8b9c3b493d2d9b5b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i0f1a586e02c48d8d20c79097f007f7e47f32f69f0042e6340415f3edc6368e09.AcceptRequestBuilder) {
    return i0f1a586e02c48d8d20c79097f007f7e47f32f69f0042e6340415f3edc6368e09.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i4402a65cd88371352fb2733c27e5fc14fdb4a77ddd1ee6c058800688fa0c6ff4.AttachmentsRequestBuilder) {
    return i4402a65cd88371352fb2733c27e5fc14fdb4a77ddd1ee6c058800688fa0c6ff4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ica8d5645131e73b6410654bc7f30c1863daa7b066d19b64654d1a4d1ecd344cb.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ica8d5645131e73b6410654bc7f30c1863daa7b066d19b64654d1a4d1ecd344cb.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i39da92daed46a78d06088ff7d64bd54ec8ee825fb267cf152af8d81a9a07d14d.CalendarRequestBuilder) {
    return i39da92daed46a78d06088ff7d64bd54ec8ee825fb267cf152af8d81a9a07d14d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i6993f9ef530b7249eff9bd6e68ee2c554fb313c4dbf671629445fc6a821307c6.CancelRequestBuilder) {
    return i6993f9ef530b7249eff9bd6e68ee2c554fb313c4dbf671629445fc6a821307c6.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from users
func (m *EventItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i0f6786d0f015f4a9e6d15f9a02d084210f77652248ff35600e6b2052e34cb8ee.DeclineRequestBuilder) {
    return i0f6786d0f015f4a9e6d15f9a02d084210f77652248ff35600e6b2052e34cb8ee.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*id2b798642d244212d8e1e2d9927afe32820c6f12099c1202597cb0eef52c5544.DismissReminderRequestBuilder) {
    return id2b798642d244212d8e1e2d9927afe32820c6f12099c1202597cb0eef52c5544.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ia2da1a05683f51fd9cd6900829f4d7a01708df0a55d0566b0babc2bf22f536d2.ExtensionsRequestBuilder) {
    return ia2da1a05683f51fd9cd6900829f4d7a01708df0a55d0566b0babc2bf22f536d2.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i28d929187c18d2d0c35f8165ba26a53887dd7d6395d44306172cf9211803d034.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i28d929187c18d2d0c35f8165ba26a53887dd7d6395d44306172cf9211803d034.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ie8736a7b60e1ca004d2a30406df2d191282e92c1a70b23210d3c8f6ac1e6b89e.ForwardRequestBuilder) {
    return ie8736a7b60e1ca004d2a30406df2d191282e92c1a70b23210d3c8f6ac1e6b89e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ia657a65517b4ceccdbbca50613f1d2362e14464ca9e58ae586c6d39acb900681.MultiValueExtendedPropertiesRequestBuilder) {
    return ia657a65517b4ceccdbbca50613f1d2362e14464ca9e58ae586c6d39acb900681.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*id55f41e90eb1b5dd9af679cc1b053c8dbbaaa6c2944925b8b9c3b493d2d9b5b1.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return id55f41e90eb1b5dd9af679cc1b053c8dbbaaa6c2944925b8b9c3b493d2d9b5b1.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i49ba1c423c143b602af84628f15aaa8269795950d34ec0e5712f027635ba0954.SingleValueExtendedPropertiesRequestBuilder) {
    return i49ba1c423c143b602af84628f15aaa8269795950d34ec0e5712f027635ba0954.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i271633dd931e6f83b538cf85551d51d635881acad169f9f55964748be82ddf45.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i271633dd931e6f83b538cf85551d51d635881acad169f9f55964748be82ddf45.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ibbc7dae07f158c8d9b1b2b21a19441045bcfaad92bac0dee00d39ceb37ba4955.SnoozeReminderRequestBuilder) {
    return ibbc7dae07f158c8d9b1b2b21a19441045bcfaad92bac0dee00d39ceb37ba4955.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i7ca9fbc65ff1e866aeea0c0b2d05b34a9105a9e06903e82afebb6b87f81ad82e.TentativelyAcceptRequestBuilder) {
    return i7ca9fbc65ff1e866aeea0c0b2d05b34a9105a9e06903e82afebb6b87f81ad82e.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
