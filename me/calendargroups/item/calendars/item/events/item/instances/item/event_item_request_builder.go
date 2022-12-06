package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i10e4b1e3f6bd852c677152292639ccd1e61010e7fa5ca1541ef13603e056f2e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/snoozereminder"
    i348cd244d3cab237a57e44a48084ec499e7c4aabdac5ce635fbcf20e267f51ea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/calendar"
    i5df077c1c5f18f890af1986ddba90e10bcd5be64cc909378ba6c006f2a93aa21 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/extensions"
    i7007d12eb47b506c4cf7d64a4e8e9acef31cb7dbcf8615ee0e17da5c10b9c078 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/dismissreminder"
    i7977fbc8ac3767d9e48f60eb33019a6411f7128bfe1778d43370247e93466bfb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/forward"
    i7a62d23724402bf2a61a7fec266609b01f3d0da42c228de9e1e2f149991db9be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/accept"
    i819662e6f790da706fc1275c925a384bdae2aa2157f2223b3fda7e82524f7f98 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/attachments"
    ia77ad0cb10f3330b8abfdf9a06b5af287dce4d8fcdf6519e8fc42835d36fdcb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences"
    id4879243a275d6c9450549607747ccd0e9789fbf769a81a5dafaf012ab42f5d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/tentativelyaccept"
    ie5ed9d4d554d1486081e85a823f0d6ff07a5df285ca148addfeaf8f01d90692b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/singlevalueextendedproperties"
    ie66e0659a8a15ea9cabcd3956cf1de646daf3de2e1544f62523e8ab049a05716 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/multivalueextendedproperties"
    if0d45f9a70f6bc747ff0c972a7c61cb9ff3a477bf18efdc1d0c88d6485eea368 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/cancel"
    ifbb707e1ea7c6ee3f196caf39c807520a3f880e4a73510141c521c5be140f338 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/decline"
    i1d99abc81fc742c4aacb3ea3eda3e9e1f27c1a09b673450ac066f23c14fc099b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/attachments/item"
    i1e367a5289a3908c9360f722ef6763a5921a23ed4f986a77c9ecbdbb3395414b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/extensions/item"
    i96e4ae565d3208b520a462d5edf4dbc60c6c8d43e60725d9ef1564cb2f0f2230 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item"
    id1576b67ac87783b278f027244a6cfdcf07c4e6ba7c8d063dffae6983467ffce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/singlevalueextendedproperties/item"
    if539699c300c7aa56c467f53218e17fdc3ff76d68ceab01c030cd819b9a71b2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/multivalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
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
// Accept provides operations to call the accept method.
func (m *EventItemRequestBuilder) Accept()(*i7a62d23724402bf2a61a7fec266609b01f3d0da42c228de9e1e2f149991db9be.AcceptRequestBuilder) {
    return i7a62d23724402bf2a61a7fec266609b01f3d0da42c228de9e1e2f149991db9be.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i819662e6f790da706fc1275c925a384bdae2aa2157f2223b3fda7e82524f7f98.AttachmentsRequestBuilder) {
    return i819662e6f790da706fc1275c925a384bdae2aa2157f2223b3fda7e82524f7f98.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i1d99abc81fc742c4aacb3ea3eda3e9e1f27c1a09b673450ac066f23c14fc099b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i1d99abc81fc742c4aacb3ea3eda3e9e1f27c1a09b673450ac066f23c14fc099b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i348cd244d3cab237a57e44a48084ec499e7c4aabdac5ce635fbcf20e267f51ea.CalendarRequestBuilder) {
    return i348cd244d3cab237a57e44a48084ec499e7c4aabdac5ce635fbcf20e267f51ea.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*if0d45f9a70f6bc747ff0c972a7c61cb9ff3a477bf18efdc1d0c88d6485eea368.CancelRequestBuilder) {
    return if0d45f9a70f6bc747ff0c972a7c61cb9ff3a477bf18efdc1d0c88d6485eea368.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) Decline()(*ifbb707e1ea7c6ee3f196caf39c807520a3f880e4a73510141c521c5be140f338.DeclineRequestBuilder) {
    return ifbb707e1ea7c6ee3f196caf39c807520a3f880e4a73510141c521c5be140f338.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i7007d12eb47b506c4cf7d64a4e8e9acef31cb7dbcf8615ee0e17da5c10b9c078.DismissReminderRequestBuilder) {
    return i7007d12eb47b506c4cf7d64a4e8e9acef31cb7dbcf8615ee0e17da5c10b9c078.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ia77ad0cb10f3330b8abfdf9a06b5af287dce4d8fcdf6519e8fc42835d36fdcb5.ExceptionOccurrencesRequestBuilder) {
    return ia77ad0cb10f3330b8abfdf9a06b5af287dce4d8fcdf6519e8fc42835d36fdcb5.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i96e4ae565d3208b520a462d5edf4dbc60c6c8d43e60725d9ef1564cb2f0f2230.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i96e4ae565d3208b520a462d5edf4dbc60c6c8d43e60725d9ef1564cb2f0f2230.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i5df077c1c5f18f890af1986ddba90e10bcd5be64cc909378ba6c006f2a93aa21.ExtensionsRequestBuilder) {
    return i5df077c1c5f18f890af1986ddba90e10bcd5be64cc909378ba6c006f2a93aa21.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i1e367a5289a3908c9360f722ef6763a5921a23ed4f986a77c9ecbdbb3395414b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i1e367a5289a3908c9360f722ef6763a5921a23ed4f986a77c9ecbdbb3395414b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i7977fbc8ac3767d9e48f60eb33019a6411f7128bfe1778d43370247e93466bfb.ForwardRequestBuilder) {
    return i7977fbc8ac3767d9e48f60eb33019a6411f7128bfe1778d43370247e93466bfb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie66e0659a8a15ea9cabcd3956cf1de646daf3de2e1544f62523e8ab049a05716.MultiValueExtendedPropertiesRequestBuilder) {
    return ie66e0659a8a15ea9cabcd3956cf1de646daf3de2e1544f62523e8ab049a05716.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if539699c300c7aa56c467f53218e17fdc3ff76d68ceab01c030cd819b9a71b2e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return if539699c300c7aa56c467f53218e17fdc3ff76d68ceab01c030cd819b9a71b2e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ie5ed9d4d554d1486081e85a823f0d6ff07a5df285ca148addfeaf8f01d90692b.SingleValueExtendedPropertiesRequestBuilder) {
    return ie5ed9d4d554d1486081e85a823f0d6ff07a5df285ca148addfeaf8f01d90692b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id1576b67ac87783b278f027244a6cfdcf07c4e6ba7c8d063dffae6983467ffce.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return id1576b67ac87783b278f027244a6cfdcf07c4e6ba7c8d063dffae6983467ffce.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i10e4b1e3f6bd852c677152292639ccd1e61010e7fa5ca1541ef13603e056f2e1.SnoozeReminderRequestBuilder) {
    return i10e4b1e3f6bd852c677152292639ccd1e61010e7fa5ca1541ef13603e056f2e1.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*id4879243a275d6c9450549607747ccd0e9789fbf769a81a5dafaf012ab42f5d6.TentativelyAcceptRequestBuilder) {
    return id4879243a275d6c9450549607747ccd0e9789fbf769a81a5dafaf012ab42f5d6.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
