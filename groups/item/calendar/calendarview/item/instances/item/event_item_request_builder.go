package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c378967a46d1c6d8fef8c8e3053a97acbc200a8e2927303fc3d3cc91217d860 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/cancel"
    i1d968c4075b6eb272d892489f3229d1ef48dc2607fd6aae2425e1f3a3efd4306 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/tentativelyaccept"
    i4bf6b5a89713a8cc11aef572a420551e91e349a0a993b31afbf4403109ac15f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/dismissreminder"
    i5f8c26c3be9361840edd7010d7e3573a2b4ccfad6a6ae6694ccc85110987349b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/snoozereminder"
    i64c8b21c35a4481e421c7b111ca9bed4f841aac9bc85689023b38b006c76a674 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/decline"
    i88053a5d3f753504ddb2e9409a5001ac85ad3dc9b3a9248af4622438c629a08d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/accept"
    ia6f6410a642f73f03ee51ff41cff9013001c024599bb34ffcff7b2118083763c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/attachments"
    ib1e73be1eee7475b632f2b470ce9a2158bfcebc8ede29a0269974a7e1e53b3c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/forward"
    ib27434ec32517d676b380899cae20b64a26ccbd9067eacc9c47159d6188af24e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/multivalueextendedproperties"
    ib3c458cfd2a3cfdd387af0e0761aa7d10192ef72e39334f5fcbfaa88e7b5988e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/singlevalueextendedproperties"
    id0a5a19eef9d25f0436de0085830e1baa56792b47135d9619f37cca448b49e7c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/calendar"
    ie85be1c63a17757f7bcc50ffde0b1597ff2f8e3d4921aad5b4689ac9d71962eb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/extensions"
    if7e18153867bed3a71b57f9f4c8e3d351e42f2454f6b183173a5e689b5658b14 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences"
    i2c38e1b456dd671227de8fdc42754f8fa70d2fa246c57d42260fbb8918d414c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/attachments/item"
    i5c8a6f3f69a906442292cc41d2fba66e51db1ee4043971bf5e01f0d6b7368288 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/multivalueextendedproperties/item"
    i9b4f3090d0fb76fc42dccf8a32b33a84c3d487422fcf038f3685a62710b10dd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/extensions/item"
    i9caaa68c46f0f06192e9eab14a6a736a4364bc7b2c830667a175cf425292be1d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item"
    ic9d735d62c225f2c487da109d382dba4e0fd322f1321b427724bc2f9f022e33a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/singlevalueextendedproperties/item"
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i88053a5d3f753504ddb2e9409a5001ac85ad3dc9b3a9248af4622438c629a08d.AcceptRequestBuilder) {
    return i88053a5d3f753504ddb2e9409a5001ac85ad3dc9b3a9248af4622438c629a08d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ia6f6410a642f73f03ee51ff41cff9013001c024599bb34ffcff7b2118083763c.AttachmentsRequestBuilder) {
    return ia6f6410a642f73f03ee51ff41cff9013001c024599bb34ffcff7b2118083763c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i2c38e1b456dd671227de8fdc42754f8fa70d2fa246c57d42260fbb8918d414c8.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i2c38e1b456dd671227de8fdc42754f8fa70d2fa246c57d42260fbb8918d414c8.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*id0a5a19eef9d25f0436de0085830e1baa56792b47135d9619f37cca448b49e7c.CalendarRequestBuilder) {
    return id0a5a19eef9d25f0436de0085830e1baa56792b47135d9619f37cca448b49e7c.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i0c378967a46d1c6d8fef8c8e3053a97acbc200a8e2927303fc3d3cc91217d860.CancelRequestBuilder) {
    return i0c378967a46d1c6d8fef8c8e3053a97acbc200a8e2927303fc3d3cc91217d860.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Decline()(*i64c8b21c35a4481e421c7b111ca9bed4f841aac9bc85689023b38b006c76a674.DeclineRequestBuilder) {
    return i64c8b21c35a4481e421c7b111ca9bed4f841aac9bc85689023b38b006c76a674.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i4bf6b5a89713a8cc11aef572a420551e91e349a0a993b31afbf4403109ac15f1.DismissReminderRequestBuilder) {
    return i4bf6b5a89713a8cc11aef572a420551e91e349a0a993b31afbf4403109ac15f1.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*if7e18153867bed3a71b57f9f4c8e3d351e42f2454f6b183173a5e689b5658b14.ExceptionOccurrencesRequestBuilder) {
    return if7e18153867bed3a71b57f9f4c8e3d351e42f2454f6b183173a5e689b5658b14.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i9caaa68c46f0f06192e9eab14a6a736a4364bc7b2c830667a175cf425292be1d.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i9caaa68c46f0f06192e9eab14a6a736a4364bc7b2c830667a175cf425292be1d.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ie85be1c63a17757f7bcc50ffde0b1597ff2f8e3d4921aad5b4689ac9d71962eb.ExtensionsRequestBuilder) {
    return ie85be1c63a17757f7bcc50ffde0b1597ff2f8e3d4921aad5b4689ac9d71962eb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9b4f3090d0fb76fc42dccf8a32b33a84c3d487422fcf038f3685a62710b10dd3.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9b4f3090d0fb76fc42dccf8a32b33a84c3d487422fcf038f3685a62710b10dd3.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ib1e73be1eee7475b632f2b470ce9a2158bfcebc8ede29a0269974a7e1e53b3c3.ForwardRequestBuilder) {
    return ib1e73be1eee7475b632f2b470ce9a2158bfcebc8ede29a0269974a7e1e53b3c3.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ib27434ec32517d676b380899cae20b64a26ccbd9067eacc9c47159d6188af24e.MultiValueExtendedPropertiesRequestBuilder) {
    return ib27434ec32517d676b380899cae20b64a26ccbd9067eacc9c47159d6188af24e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5c8a6f3f69a906442292cc41d2fba66e51db1ee4043971bf5e01f0d6b7368288.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i5c8a6f3f69a906442292cc41d2fba66e51db1ee4043971bf5e01f0d6b7368288.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib3c458cfd2a3cfdd387af0e0761aa7d10192ef72e39334f5fcbfaa88e7b5988e.SingleValueExtendedPropertiesRequestBuilder) {
    return ib3c458cfd2a3cfdd387af0e0761aa7d10192ef72e39334f5fcbfaa88e7b5988e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic9d735d62c225f2c487da109d382dba4e0fd322f1321b427724bc2f9f022e33a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ic9d735d62c225f2c487da109d382dba4e0fd322f1321b427724bc2f9f022e33a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5f8c26c3be9361840edd7010d7e3573a2b4ccfad6a6ae6694ccc85110987349b.SnoozeReminderRequestBuilder) {
    return i5f8c26c3be9361840edd7010d7e3573a2b4ccfad6a6ae6694ccc85110987349b.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i1d968c4075b6eb272d892489f3229d1ef48dc2607fd6aae2425e1f3a3efd4306.TentativelyAcceptRequestBuilder) {
    return i1d968c4075b6eb272d892489f3229d1ef48dc2607fd6aae2425e1f3a3efd4306.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
