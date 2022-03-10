package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*i478733ca896fd701e010f337cda235511583204d47386ecfa56b62e0a98443fe.AcceptRequestBuilder) {
    return i478733ca896fd701e010f337cda235511583204d47386ecfa56b62e0a98443fe.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["attachment_id"] = id
    }
    return i2c9a832a9f6a17e5461f70b0de4c07e5cb5dc07c2042e5ce892821c9b1d0707f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i8e55b29a738e2fbe5af7a71e39a259f99d831e9d0b3177ce1998729ccc6e2929.CalendarRequestBuilder) {
    return i8e55b29a738e2fbe5af7a71e39a259f99d831e9d0b3177ce1998729ccc6e2929.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i4f89703cf441bd902b89cb4c99f75bac13b63aeb7b5845a3436d1e7c689a5195.CancelRequestBuilder) {
    return i4f89703cf441bd902b89cb4c99f75bac13b63aeb7b5845a3436d1e7c689a5195.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/calendarView/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instances in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventItemRequestBuilder) Decline()(*id3274881ae65cccae5d315d739c891094c129050f5fa22b31f4913301f689576.DeclineRequestBuilder) {
    return id3274881ae65cccae5d315d739c891094c129050f5fa22b31f4913301f689576.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for groups
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i49b1c97d15f2316447a9bb0b4269fb4e294ae3fc9bedd58a3d50639ea9002dd5.DismissReminderRequestBuilder) {
    return i49b1c97d15f2316447a9bb0b4269fb4e294ae3fc9bedd58a3d50639ea9002dd5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["extension_id"] = id
    }
    return ia78fa6be731642296ed3ffa2cdcf04348a39532233073ef6c1264a6b6d2b8956.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*id7510f91b0460b935816b2dd68aadb1cbae0b5f64d6c8070ba8e618aa4a40ead.ForwardRequestBuilder) {
    return id7510f91b0460b935816b2dd68aadb1cbae0b5f64d6c8070ba8e618aa4a40ead.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
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
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie2c165457950b83643c10f0b0e117a8512ee93363019e33f30e36cca0be3bc03.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in groups
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i605921d2853da5fb2c5ce39e6919ad5721219379375ce52846d806d6982d1d9b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ia5dd167797f9cb74408b3a6f5aac9532212dd80dc5f95b576582190028fe4e15.SnoozeReminderRequestBuilder) {
    return ia5dd167797f9cb74408b3a6f5aac9532212dd80dc5f95b576582190028fe4e15.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3de6a5b64466af0fe93de1a2414ed49bd849ea3938b865cb40d1364b532a7165.TentativelyAcceptRequestBuilder) {
    return i3de6a5b64466af0fe93de1a2414ed49bd849ea3938b865cb40d1364b532a7165.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
