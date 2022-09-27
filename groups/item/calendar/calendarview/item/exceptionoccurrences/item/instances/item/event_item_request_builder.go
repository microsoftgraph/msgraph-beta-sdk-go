package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0e6e2d6d6307fb89a517741c32edd0fefa5725f3848956f868ff5f68a7d48e29 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    i3de6a5b64466af0fe93de1a2414ed49bd849ea3938b865cb40d1364b532a7165 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i478733ca896fd701e010f337cda235511583204d47386ecfa56b62e0a98443fe "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    i49b1c97d15f2316447a9bb0b4269fb4e294ae3fc9bedd58a3d50639ea9002dd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i4f89703cf441bd902b89cb4c99f75bac13b63aeb7b5845a3436d1e7c689a5195 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i860357b7b59c0fcfab261744812712e787557aa40b1f31c4d0413f0856299c67 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i8e55b29a738e2fbe5af7a71e39a259f99d831e9d0b3177ce1998729ccc6e2929 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    ia5dd167797f9cb74408b3a6f5aac9532212dd80dc5f95b576582190028fe4e15 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ib1e00580978e78ae9d4262278b018462aa0b361c4072b6c05107ffdb05fc2c58 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    id3274881ae65cccae5d315d739c891094c129050f5fa22b31f4913301f689576 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    id7510f91b0460b935816b2dd68aadb1cbae0b5f64d6c8070ba8e618aa4a40ead "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    ie94159ca7519a9ab32026053d2eeb7e205e35ca22049f25882d343db5c60f5a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i2c9a832a9f6a17e5461f70b0de4c07e5cb5dc07c2042e5ce892821c9b1d0707f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    i605921d2853da5fb2c5ce39e6919ad5721219379375ce52846d806d6982d1d9b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    ia78fa6be731642296ed3ffa2cdcf04348a39532233073ef6c1264a6b6d2b8956 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    ie2c165457950b83643c10f0b0e117a8512ee93363019e33f30e36cca0be3bc03 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i478733ca896fd701e010f337cda235511583204d47386ecfa56b62e0a98443fe.AcceptRequestBuilder) {
    return i478733ca896fd701e010f337cda235511583204d47386ecfa56b62e0a98443fe.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ib1e00580978e78ae9d4262278b018462aa0b361c4072b6c05107ffdb05fc2c58.AttachmentsRequestBuilder) {
    return ib1e00580978e78ae9d4262278b018462aa0b361c4072b6c05107ffdb05fc2c58.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i2c9a832a9f6a17e5461f70b0de4c07e5cb5dc07c2042e5ce892821c9b1d0707f.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i2c9a832a9f6a17e5461f70b0de4c07e5cb5dc07c2042e5ce892821c9b1d0707f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i8e55b29a738e2fbe5af7a71e39a259f99d831e9d0b3177ce1998729ccc6e2929.CalendarRequestBuilder) {
    return i8e55b29a738e2fbe5af7a71e39a259f99d831e9d0b3177ce1998729ccc6e2929.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i4f89703cf441bd902b89cb4c99f75bac13b63aeb7b5845a3436d1e7c689a5195.CancelRequestBuilder) {
    return i4f89703cf441bd902b89cb4c99f75bac13b63aeb7b5845a3436d1e7c689a5195.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*id3274881ae65cccae5d315d739c891094c129050f5fa22b31f4913301f689576.DeclineRequestBuilder) {
    return id3274881ae65cccae5d315d739c891094c129050f5fa22b31f4913301f689576.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i49b1c97d15f2316447a9bb0b4269fb4e294ae3fc9bedd58a3d50639ea9002dd5.DismissReminderRequestBuilder) {
    return i49b1c97d15f2316447a9bb0b4269fb4e294ae3fc9bedd58a3d50639ea9002dd5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i0e6e2d6d6307fb89a517741c32edd0fefa5725f3848956f868ff5f68a7d48e29.ExtensionsRequestBuilder) {
    return i0e6e2d6d6307fb89a517741c32edd0fefa5725f3848956f868ff5f68a7d48e29.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia78fa6be731642296ed3ffa2cdcf04348a39532233073ef6c1264a6b6d2b8956.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ia78fa6be731642296ed3ffa2cdcf04348a39532233073ef6c1264a6b6d2b8956.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*id7510f91b0460b935816b2dd68aadb1cbae0b5f64d6c8070ba8e618aa4a40ead.ForwardRequestBuilder) {
    return id7510f91b0460b935816b2dd68aadb1cbae0b5f64d6c8070ba8e618aa4a40ead.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i860357b7b59c0fcfab261744812712e787557aa40b1f31c4d0413f0856299c67.MultiValueExtendedPropertiesRequestBuilder) {
    return i860357b7b59c0fcfab261744812712e787557aa40b1f31c4d0413f0856299c67.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie2c165457950b83643c10f0b0e117a8512ee93363019e33f30e36cca0be3bc03.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ie2c165457950b83643c10f0b0e117a8512ee93363019e33f30e36cca0be3bc03.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ie94159ca7519a9ab32026053d2eeb7e205e35ca22049f25882d343db5c60f5a3.SingleValueExtendedPropertiesRequestBuilder) {
    return ie94159ca7519a9ab32026053d2eeb7e205e35ca22049f25882d343db5c60f5a3.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i605921d2853da5fb2c5ce39e6919ad5721219379375ce52846d806d6982d1d9b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i605921d2853da5fb2c5ce39e6919ad5721219379375ce52846d806d6982d1d9b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ia5dd167797f9cb74408b3a6f5aac9532212dd80dc5f95b576582190028fe4e15.SnoozeReminderRequestBuilder) {
    return ia5dd167797f9cb74408b3a6f5aac9532212dd80dc5f95b576582190028fe4e15.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3de6a5b64466af0fe93de1a2414ed49bd849ea3938b865cb40d1364b532a7165.TentativelyAcceptRequestBuilder) {
    return i3de6a5b64466af0fe93de1a2414ed49bd849ea3938b865cb40d1364b532a7165.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
