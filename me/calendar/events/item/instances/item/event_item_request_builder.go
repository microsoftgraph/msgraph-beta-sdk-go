package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0356886c0845609492b9bf60a7c488acd69b6c86b232fb198255f609fe93b31f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences"
    i1444a249ac01aafa0735bf7e8b5da62666e2f5f4f981ff9c3eaead8ea876f150 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/attachments"
    i1fe260e215ca0781a61cb33b1b2d12905f0442276d4c2fc0e4c637d24ca16fca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/dismissreminder"
    i6c7a2a249bc358926cda74f02fff64df9ff5886a4df160fd0fc8519827d36071 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/calendar"
    i790167e20913c493f017a6d467297de951bf0f122a51b4388dfc46b398feefdc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/snoozereminder"
    i7a2d2941f0e06491b021dfe7b8f9766bef924677c3eaf11a47c9cb3afb2da641 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/decline"
    i8442e254adaf4b462a3bdf5334936a63edd0b84b501d58a1ff28f9270a2e6b51 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/forward"
    i957a61dd0e69cae772b2942f2c97a85edeb70871e49605fdde860aeda90ba99a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/tentativelyaccept"
    ic0253bacd98fa4b7cf33288f799b41324814b0d5974b51be59fbb199a1c420e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/singlevalueextendedproperties"
    ic1e4e3d17c65f260c4a606bc8c54c7d47e0b62dd73d3e2f0a1af5cc189a0481f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/cancel"
    ic51195672f9785a7d9414beedccec0248a577de7ff5d589516b1ca3f89571008 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/multivalueextendedproperties"
    ie2e5b8126528fff04b0ed490c8d661ddd3070ec664dd6689c0914929dcf2dda3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/extensions"
    if5088fb26ef324fac5d36dc63fa11212bb3431e4d17a8fe905a7a6225b6b5288 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/accept"
    i43a37fac2ef5c960a562463e6c39f54e76c26a06d5e7c24bfa86544c4f0d4d52 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/extensions/item"
    i5d9bc9c397a3755d7ac3f37791085759b4e7e5dcd18bcf7878dba68168ece27a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/singlevalueextendedproperties/item"
    ib0c7f369f9a2f64b7fab964afd56c0b5f29817d76afb03b61f0547a1dc72feb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/attachments/item"
    ibf251d271f90c69e84013601225d7bea0e42862e91588b16bd9184815d9433b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item"
    ide426a1a4aa0853cb0f90be46e3048fb55f5158047f30979fd91b49e2ea240a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*if5088fb26ef324fac5d36dc63fa11212bb3431e4d17a8fe905a7a6225b6b5288.AcceptRequestBuilder) {
    return if5088fb26ef324fac5d36dc63fa11212bb3431e4d17a8fe905a7a6225b6b5288.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i1444a249ac01aafa0735bf7e8b5da62666e2f5f4f981ff9c3eaead8ea876f150.AttachmentsRequestBuilder) {
    return i1444a249ac01aafa0735bf7e8b5da62666e2f5f4f981ff9c3eaead8ea876f150.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib0c7f369f9a2f64b7fab964afd56c0b5f29817d76afb03b61f0547a1dc72feb3.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ib0c7f369f9a2f64b7fab964afd56c0b5f29817d76afb03b61f0547a1dc72feb3.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i6c7a2a249bc358926cda74f02fff64df9ff5886a4df160fd0fc8519827d36071.CalendarRequestBuilder) {
    return i6c7a2a249bc358926cda74f02fff64df9ff5886a4df160fd0fc8519827d36071.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*ic1e4e3d17c65f260c4a606bc8c54c7d47e0b62dd73d3e2f0a1af5cc189a0481f.CancelRequestBuilder) {
    return ic1e4e3d17c65f260c4a606bc8c54c7d47e0b62dd73d3e2f0a1af5cc189a0481f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i7a2d2941f0e06491b021dfe7b8f9766bef924677c3eaf11a47c9cb3afb2da641.DeclineRequestBuilder) {
    return i7a2d2941f0e06491b021dfe7b8f9766bef924677c3eaf11a47c9cb3afb2da641.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i1fe260e215ca0781a61cb33b1b2d12905f0442276d4c2fc0e4c637d24ca16fca.DismissReminderRequestBuilder) {
    return i1fe260e215ca0781a61cb33b1b2d12905f0442276d4c2fc0e4c637d24ca16fca.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i0356886c0845609492b9bf60a7c488acd69b6c86b232fb198255f609fe93b31f.ExceptionOccurrencesRequestBuilder) {
    return i0356886c0845609492b9bf60a7c488acd69b6c86b232fb198255f609fe93b31f.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ibf251d271f90c69e84013601225d7bea0e42862e91588b16bd9184815d9433b7.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return ibf251d271f90c69e84013601225d7bea0e42862e91588b16bd9184815d9433b7.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*ie2e5b8126528fff04b0ed490c8d661ddd3070ec664dd6689c0914929dcf2dda3.ExtensionsRequestBuilder) {
    return ie2e5b8126528fff04b0ed490c8d661ddd3070ec664dd6689c0914929dcf2dda3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i43a37fac2ef5c960a562463e6c39f54e76c26a06d5e7c24bfa86544c4f0d4d52.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i43a37fac2ef5c960a562463e6c39f54e76c26a06d5e7c24bfa86544c4f0d4d52.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i8442e254adaf4b462a3bdf5334936a63edd0b84b501d58a1ff28f9270a2e6b51.ForwardRequestBuilder) {
    return i8442e254adaf4b462a3bdf5334936a63edd0b84b501d58a1ff28f9270a2e6b51.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic51195672f9785a7d9414beedccec0248a577de7ff5d589516b1ca3f89571008.MultiValueExtendedPropertiesRequestBuilder) {
    return ic51195672f9785a7d9414beedccec0248a577de7ff5d589516b1ca3f89571008.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ide426a1a4aa0853cb0f90be46e3048fb55f5158047f30979fd91b49e2ea240a2.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ide426a1a4aa0853cb0f90be46e3048fb55f5158047f30979fd91b49e2ea240a2.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ic0253bacd98fa4b7cf33288f799b41324814b0d5974b51be59fbb199a1c420e3.SingleValueExtendedPropertiesRequestBuilder) {
    return ic0253bacd98fa4b7cf33288f799b41324814b0d5974b51be59fbb199a1c420e3.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5d9bc9c397a3755d7ac3f37791085759b4e7e5dcd18bcf7878dba68168ece27a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i5d9bc9c397a3755d7ac3f37791085759b4e7e5dcd18bcf7878dba68168ece27a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i790167e20913c493f017a6d467297de951bf0f122a51b4388dfc46b398feefdc.SnoozeReminderRequestBuilder) {
    return i790167e20913c493f017a6d467297de951bf0f122a51b4388dfc46b398feefdc.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i957a61dd0e69cae772b2942f2c97a85edeb70871e49605fdde860aeda90ba99a.TentativelyAcceptRequestBuilder) {
    return i957a61dd0e69cae772b2942f2c97a85edeb70871e49605fdde860aeda90ba99a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
