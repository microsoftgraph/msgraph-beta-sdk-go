package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0062b298e2ffb7e864aa5cf6d632cd823ed9566653be5f6e8889231b84ce4ac6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i130975e0dfdb91688100b9e129e8f10467053f976e74ac87dc6d4854a54d2269 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i2eade44e5de8ec4344095b0746f2d6960e63a87bcf9112e70c1211232da9cdd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i4eea483bb3b07759a8e23e0b484f1387be1d4a551467220bf81b7cacb9229b14 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i5fd86f654e0f3b974006e93b1a60dfa42ec4b2a4c807a4f2d19c0db3f187d8ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i6f26a813084343024cbe1c3286df0b8f2f26416eba0a1775a482cdce89fb7dbc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i7bd7e43ce7d9bea156f2b9631c5f4813d23f9c1845297c6ae689fa909021aaf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    ia9a4ad63c67b1c65d087f4752cd189fa43de9d39f8f63d194308ed8fc36359f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    ib0d47db12e79afad843f2e24dc210d801526edb509046a1b212735c8e5becf63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    icce2f9aa51075b8ad56182a928bc9ad87c78b053af67e5c992385b87dc201261 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    icf61d9d88f414ffb4011d10abea9efba1294d34aa148f99e68a74aff7f08aa65 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    idcf18fa49cece6031f6eacf311f33950e2dcf50388e29fa130e3faa25cbf55cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i120cbebfaf8ba152ed4efd5cf712e7e7b635f0772df6ec8110dc5b4ab24ff598 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    i88adc21805da074a5a0b41c23c3b02be605d362660d5cdb45436d56c427209fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i9d03403c0367899199389a99a42e39415397ebb34b7e1daa91b81bed69ba97aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ida5e3f89ae23131e5b26fed27044589841ae5730177f129f83f50a7bff858f54 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
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
// Accept provides operations to call the accept method.
func (m *EventItemRequestBuilder) Accept()(*ia9a4ad63c67b1c65d087f4752cd189fa43de9d39f8f63d194308ed8fc36359f6.AcceptRequestBuilder) {
    return ia9a4ad63c67b1c65d087f4752cd189fa43de9d39f8f63d194308ed8fc36359f6.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i130975e0dfdb91688100b9e129e8f10467053f976e74ac87dc6d4854a54d2269.AttachmentsRequestBuilder) {
    return i130975e0dfdb91688100b9e129e8f10467053f976e74ac87dc6d4854a54d2269.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i120cbebfaf8ba152ed4efd5cf712e7e7b635f0772df6ec8110dc5b4ab24ff598.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i120cbebfaf8ba152ed4efd5cf712e7e7b635f0772df6ec8110dc5b4ab24ff598.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*ib0d47db12e79afad843f2e24dc210d801526edb509046a1b212735c8e5becf63.CalendarRequestBuilder) {
    return ib0d47db12e79afad843f2e24dc210d801526edb509046a1b212735c8e5becf63.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*idcf18fa49cece6031f6eacf311f33950e2dcf50388e29fa130e3faa25cbf55cd.CancelRequestBuilder) {
    return idcf18fa49cece6031f6eacf311f33950e2dcf50388e29fa130e3faa25cbf55cd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Decline()(*icce2f9aa51075b8ad56182a928bc9ad87c78b053af67e5c992385b87dc201261.DeclineRequestBuilder) {
    return icce2f9aa51075b8ad56182a928bc9ad87c78b053af67e5c992385b87dc201261.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i2eade44e5de8ec4344095b0746f2d6960e63a87bcf9112e70c1211232da9cdd5.DismissReminderRequestBuilder) {
    return i2eade44e5de8ec4344095b0746f2d6960e63a87bcf9112e70c1211232da9cdd5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*icf61d9d88f414ffb4011d10abea9efba1294d34aa148f99e68a74aff7f08aa65.ExtensionsRequestBuilder) {
    return icf61d9d88f414ffb4011d10abea9efba1294d34aa148f99e68a74aff7f08aa65.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ida5e3f89ae23131e5b26fed27044589841ae5730177f129f83f50a7bff858f54.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ida5e3f89ae23131e5b26fed27044589841ae5730177f129f83f50a7bff858f54.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i0062b298e2ffb7e864aa5cf6d632cd823ed9566653be5f6e8889231b84ce4ac6.ForwardRequestBuilder) {
    return i0062b298e2ffb7e864aa5cf6d632cd823ed9566653be5f6e8889231b84ce4ac6.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4eea483bb3b07759a8e23e0b484f1387be1d4a551467220bf81b7cacb9229b14.MultiValueExtendedPropertiesRequestBuilder) {
    return i4eea483bb3b07759a8e23e0b484f1387be1d4a551467220bf81b7cacb9229b14.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9d03403c0367899199389a99a42e39415397ebb34b7e1daa91b81bed69ba97aa.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i9d03403c0367899199389a99a42e39415397ebb34b7e1daa91b81bed69ba97aa.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i5fd86f654e0f3b974006e93b1a60dfa42ec4b2a4c807a4f2d19c0db3f187d8ee.SingleValueExtendedPropertiesRequestBuilder) {
    return i5fd86f654e0f3b974006e93b1a60dfa42ec4b2a4c807a4f2d19c0db3f187d8ee.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i88adc21805da074a5a0b41c23c3b02be605d362660d5cdb45436d56c427209fa.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i88adc21805da074a5a0b41c23c3b02be605d362660d5cdb45436d56c427209fa.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i7bd7e43ce7d9bea156f2b9631c5f4813d23f9c1845297c6ae689fa909021aaf7.SnoozeReminderRequestBuilder) {
    return i7bd7e43ce7d9bea156f2b9631c5f4813d23f9c1845297c6ae689fa909021aaf7.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6f26a813084343024cbe1c3286df0b8f2f26416eba0a1775a482cdce89fb7dbc.TentativelyAcceptRequestBuilder) {
    return i6f26a813084343024cbe1c3286df0b8f2f26416eba0a1775a482cdce89fb7dbc.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
