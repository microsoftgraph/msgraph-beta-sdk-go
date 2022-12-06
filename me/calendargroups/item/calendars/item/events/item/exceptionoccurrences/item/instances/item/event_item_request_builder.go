package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c61c219ece89a15a3b1c362293daf307dd56203dc2eb41819519507fb821a95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i1b87d4349875d608b2fcb309f83fbc50fa254967338fa4eb64c0bf6d92aca4f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i578ac15ad5122219997505817b9354c0e48936a98cafe3e7551f14635b82cd9f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    i5ad5005cc0b7805c9c270ecf3f3e598cb43d0c21cb21a5657fed812c6a7b8321 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i6d4f41a32747a54163e709d5145e993de2b2f14bf108f14b7b56af2629bc8476 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i764c1423bfc8f8627cf88717b79626133712d8c00eadf4653fb56066620b8a85 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/decline"
    i8b53255ad0927e543eab23ffd50e9282203cbcf19891e4ad86472f72d815ca20 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/forward"
    i8cc663b5a8e8c8fee06db5c214133be38b3f8e9b2d514de01d6ec4891e4d3f83 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    iad59b69a523dc48fc2eb9dc3d9de66dfac837fb2341dc6690e11ed9072cbaa66 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    iad8ee13b80e46088e7e6bca3b0410f792a1b002bc7b2cd3367c58af73ed2b526 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/accept"
    idd840b3e79ffdbb66997bb899dd0b8f9f2084711fbbe4c30ba5215ce4636b5c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    ifbd62e19f5473073107ffc419dfa8d09e08c8a3d84ca506326d42b3ec2eb4f21 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i0ec8f604e36f8e58b97583932d40948c21ed10d5add5a1ce4dece35ce1f4795f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i230d6ceddb21d1429027ad16976ae7efbd7bd4e4199246f028e100fcdd283524 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    ic27f680db50c40280b447dfc22db1c197bb03ebdf796cd72d2378b87369a7325 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    id831ec7431bc1afe337c9d7b8ba62662c19e7886c20faf9b8c4b0576eb51644d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*iad8ee13b80e46088e7e6bca3b0410f792a1b002bc7b2cd3367c58af73ed2b526.AcceptRequestBuilder) {
    return iad8ee13b80e46088e7e6bca3b0410f792a1b002bc7b2cd3367c58af73ed2b526.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i1b87d4349875d608b2fcb309f83fbc50fa254967338fa4eb64c0bf6d92aca4f9.AttachmentsRequestBuilder) {
    return i1b87d4349875d608b2fcb309f83fbc50fa254967338fa4eb64c0bf6d92aca4f9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i0ec8f604e36f8e58b97583932d40948c21ed10d5add5a1ce4dece35ce1f4795f.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i0ec8f604e36f8e58b97583932d40948c21ed10d5add5a1ce4dece35ce1f4795f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i578ac15ad5122219997505817b9354c0e48936a98cafe3e7551f14635b82cd9f.CalendarRequestBuilder) {
    return i578ac15ad5122219997505817b9354c0e48936a98cafe3e7551f14635b82cd9f.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i6d4f41a32747a54163e709d5145e993de2b2f14bf108f14b7b56af2629bc8476.CancelRequestBuilder) {
    return i6d4f41a32747a54163e709d5145e993de2b2f14bf108f14b7b56af2629bc8476.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i764c1423bfc8f8627cf88717b79626133712d8c00eadf4653fb56066620b8a85.DeclineRequestBuilder) {
    return i764c1423bfc8f8627cf88717b79626133712d8c00eadf4653fb56066620b8a85.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*iad59b69a523dc48fc2eb9dc3d9de66dfac837fb2341dc6690e11ed9072cbaa66.DismissReminderRequestBuilder) {
    return iad59b69a523dc48fc2eb9dc3d9de66dfac837fb2341dc6690e11ed9072cbaa66.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i5ad5005cc0b7805c9c270ecf3f3e598cb43d0c21cb21a5657fed812c6a7b8321.ExtensionsRequestBuilder) {
    return i5ad5005cc0b7805c9c270ecf3f3e598cb43d0c21cb21a5657fed812c6a7b8321.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i230d6ceddb21d1429027ad16976ae7efbd7bd4e4199246f028e100fcdd283524.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i230d6ceddb21d1429027ad16976ae7efbd7bd4e4199246f028e100fcdd283524.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i8b53255ad0927e543eab23ffd50e9282203cbcf19891e4ad86472f72d815ca20.ForwardRequestBuilder) {
    return i8b53255ad0927e543eab23ffd50e9282203cbcf19891e4ad86472f72d815ca20.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ifbd62e19f5473073107ffc419dfa8d09e08c8a3d84ca506326d42b3ec2eb4f21.MultiValueExtendedPropertiesRequestBuilder) {
    return ifbd62e19f5473073107ffc419dfa8d09e08c8a3d84ca506326d42b3ec2eb4f21.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic27f680db50c40280b447dfc22db1c197bb03ebdf796cd72d2378b87369a7325.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ic27f680db50c40280b447dfc22db1c197bb03ebdf796cd72d2378b87369a7325.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i8cc663b5a8e8c8fee06db5c214133be38b3f8e9b2d514de01d6ec4891e4d3f83.SingleValueExtendedPropertiesRequestBuilder) {
    return i8cc663b5a8e8c8fee06db5c214133be38b3f8e9b2d514de01d6ec4891e4d3f83.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id831ec7431bc1afe337c9d7b8ba62662c19e7886c20faf9b8c4b0576eb51644d.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return id831ec7431bc1afe337c9d7b8ba62662c19e7886c20faf9b8c4b0576eb51644d.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i0c61c219ece89a15a3b1c362293daf307dd56203dc2eb41819519507fb821a95.SnoozeReminderRequestBuilder) {
    return i0c61c219ece89a15a3b1c362293daf307dd56203dc2eb41819519507fb821a95.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*idd840b3e79ffdbb66997bb899dd0b8f9f2084711fbbe4c30ba5215ce4636b5c3.TentativelyAcceptRequestBuilder) {
    return idd840b3e79ffdbb66997bb899dd0b8f9f2084711fbbe4c30ba5215ce4636b5c3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
