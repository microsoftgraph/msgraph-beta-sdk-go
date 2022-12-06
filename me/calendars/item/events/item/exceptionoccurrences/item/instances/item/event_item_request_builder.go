package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i13ca2ae1acbace613d6d3be66cc9452c0d770a6528fcab85b29a52b3c6f040a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i2a921ec71a89667427cf20125da19b5a8fb2908741f78c1fe6f3f97bc3e8b6a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/forward"
    i2c72860be341111e303e7055357e65539cf937295dc6c580f8f2225ee8e5464a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i3598f99f347fac630875d2ce261c36924aba0f9771a99458d209f2c2e6fd97aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i3ea35641728a04999309cc0d22f521bf1a4e9ce82dbc4f1be02453cd161e6760 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/decline"
    i66a5e2fe3cadab88d66dae06f1b1c5095ce422597e40c8a0bee91694d0f8ff3e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i77727a6d99b4ac119324bb2712b1fcf112bd8cc32b8d2a5cfb765200e81a0869 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i89d95640cae9611da2b6ad9e68fd7c24d3a7c8e7112408806523e6f2bab082a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ib89e08cdbe3e5c3152add0842332a59a8acde4fb4a24dfd361e670fa60c39c7a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    id89660e80708f330cb2a7abaca3937a3d10567ca2d17b208f12ab6644256ae27 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    idb7add6fe03fa68c3b4d4e08ca47db0642dd83a5729c945990f999225e15223f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    if2617f7c54055c759722b2340e27f24d6e8efaf7993109eebb8007ed66d87cb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/accept"
    i31b269866bcb6b70a0f120c92bfb4301725963629321dae678377411eac48dda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i65022a937c245df99276fdbc4bcd868c2572a195571bacb0456d97e2e3eb1256 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    ia52fb8bdd3d432e46306f31d580cec6a80425239cb5b34363af0801f1757bd64 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    ic307e30e757b58447196969d338b33a1ae0bad7fcc17369d701a6f85e83f41c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*if2617f7c54055c759722b2340e27f24d6e8efaf7993109eebb8007ed66d87cb5.AcceptRequestBuilder) {
    return if2617f7c54055c759722b2340e27f24d6e8efaf7993109eebb8007ed66d87cb5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i66a5e2fe3cadab88d66dae06f1b1c5095ce422597e40c8a0bee91694d0f8ff3e.AttachmentsRequestBuilder) {
    return i66a5e2fe3cadab88d66dae06f1b1c5095ce422597e40c8a0bee91694d0f8ff3e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia52fb8bdd3d432e46306f31d580cec6a80425239cb5b34363af0801f1757bd64.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ia52fb8bdd3d432e46306f31d580cec6a80425239cb5b34363af0801f1757bd64.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*idb7add6fe03fa68c3b4d4e08ca47db0642dd83a5729c945990f999225e15223f.CalendarRequestBuilder) {
    return idb7add6fe03fa68c3b4d4e08ca47db0642dd83a5729c945990f999225e15223f.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i3598f99f347fac630875d2ce261c36924aba0f9771a99458d209f2c2e6fd97aa.CancelRequestBuilder) {
    return i3598f99f347fac630875d2ce261c36924aba0f9771a99458d209f2c2e6fd97aa.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i3ea35641728a04999309cc0d22f521bf1a4e9ce82dbc4f1be02453cd161e6760.DeclineRequestBuilder) {
    return i3ea35641728a04999309cc0d22f521bf1a4e9ce82dbc4f1be02453cd161e6760.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*id89660e80708f330cb2a7abaca3937a3d10567ca2d17b208f12ab6644256ae27.DismissReminderRequestBuilder) {
    return id89660e80708f330cb2a7abaca3937a3d10567ca2d17b208f12ab6644256ae27.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i13ca2ae1acbace613d6d3be66cc9452c0d770a6528fcab85b29a52b3c6f040a9.ExtensionsRequestBuilder) {
    return i13ca2ae1acbace613d6d3be66cc9452c0d770a6528fcab85b29a52b3c6f040a9.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic307e30e757b58447196969d338b33a1ae0bad7fcc17369d701a6f85e83f41c7.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ic307e30e757b58447196969d338b33a1ae0bad7fcc17369d701a6f85e83f41c7.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i2a921ec71a89667427cf20125da19b5a8fb2908741f78c1fe6f3f97bc3e8b6a7.ForwardRequestBuilder) {
    return i2a921ec71a89667427cf20125da19b5a8fb2908741f78c1fe6f3f97bc3e8b6a7.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i77727a6d99b4ac119324bb2712b1fcf112bd8cc32b8d2a5cfb765200e81a0869.MultiValueExtendedPropertiesRequestBuilder) {
    return i77727a6d99b4ac119324bb2712b1fcf112bd8cc32b8d2a5cfb765200e81a0869.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i65022a937c245df99276fdbc4bcd868c2572a195571bacb0456d97e2e3eb1256.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i65022a937c245df99276fdbc4bcd868c2572a195571bacb0456d97e2e3eb1256.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib89e08cdbe3e5c3152add0842332a59a8acde4fb4a24dfd361e670fa60c39c7a.SingleValueExtendedPropertiesRequestBuilder) {
    return ib89e08cdbe3e5c3152add0842332a59a8acde4fb4a24dfd361e670fa60c39c7a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i31b269866bcb6b70a0f120c92bfb4301725963629321dae678377411eac48dda.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i31b269866bcb6b70a0f120c92bfb4301725963629321dae678377411eac48dda.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i89d95640cae9611da2b6ad9e68fd7c24d3a7c8e7112408806523e6f2bab082a3.SnoozeReminderRequestBuilder) {
    return i89d95640cae9611da2b6ad9e68fd7c24d3a7c8e7112408806523e6f2bab082a3.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i2c72860be341111e303e7055357e65539cf937295dc6c580f8f2225ee8e5464a.TentativelyAcceptRequestBuilder) {
    return i2c72860be341111e303e7055357e65539cf937295dc6c580f8f2225ee8e5464a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
