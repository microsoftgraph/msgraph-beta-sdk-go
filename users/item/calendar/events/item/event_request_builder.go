package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1022eeefd0950e0e70752c5cb0b720b2a7fbd1152f8b2db65dbf1f1fe3785fdd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/tentativelyaccept"
    i16b073d42adfaa42cd2197fe997fd666e6c717513e29bbc1411d84d135e5907d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/dismissreminder"
    i3261245f326c44b01fe5da55ef3eb6fdcf3eec9727c6249445c533e07ad1762b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/cancel"
    i3cc615139480e6b6080500d2592d3b826f8172402d8123a8033c61f3b0f6d2b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/extensions"
    i449d1a85f5abadbfceb7b75147c603057ebd7e4b6fe7fbc07c7854cef9f748be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances"
    i5ccef22bea92cabddac92714e10b7b992bc8b281f5928395c1b100cef0454e2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences"
    i71f34fdb44622e1b8dbd50563e3b844d5b06315264bd4bf7dcfc1262b7101a4f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/attachments"
    i912a64a291cce9b7af41be42dd7d2c73b1358686172dc2d4200444b14724cf3c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/calendar"
    i9fdfb1b218cc10bbbf0b6f6529dc0097d2a03fa8fa6a246d23addca54cc1138f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/snoozereminder"
    ibe8940be56566d710a3152778a1869ba155c8636fa528fc2f319110c15e678a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/multivalueextendedproperties"
    ic365f138404b6574eb593b1e6e3be13031d99faa7fde1b344c069d91fe235dbc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/forward"
    iddfae0fb4ae92b93079596d35b8250f92d21dc1f6215fad3b6818349dc5a0e58 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/accept"
    ie9a1738dc4818849cb5a030a886b3107bfbc115e7d40261896f7c8b5c0a1d983 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/singlevalueextendedproperties"
    if65de254c43ef0974070ea922f0230c5434fa0eaef859e883b04b25ba3f75a0b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/decline"
    i35e38f254ec120c13f755d3867f2fd9952b17b0aef33b67009171560f14276ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/extensions/item"
    i39d2408d6b0b07a8e199b2d29e86bbc35c38980add543123b882bc022c320552 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/singlevalueextendedproperties/item"
    i53b1fc78501c4055752cde988ccaa74dc9af1ec1a4f11c6c09b19a16565c5fad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item"
    i5cd1ea32e6df077b61718a3134fe8e9004c569fa48259151b83a3713a0993009 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item"
    ib21562e8f33ec5c88876706437f50c8b6523900dedd39da40e1e91e9c54470d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/multivalueextendedproperties/item"
    ief2a3897f4132f3372d8f033cd2ebd787a4705005e05e502f64f9441d3294683 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/attachments/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendar\events\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventRequestBuilderDeleteOptions options for Delete
type EventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetOptions options for Get
type EventRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventRequestBuilderPatchOptions options for Patch
type EventRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*iddfae0fb4ae92b93079596d35b8250f92d21dc1f6215fad3b6818349dc5a0e58.AcceptRequestBuilder) {
    return iddfae0fb4ae92b93079596d35b8250f92d21dc1f6215fad3b6818349dc5a0e58.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i71f34fdb44622e1b8dbd50563e3b844d5b06315264bd4bf7dcfc1262b7101a4f.AttachmentsRequestBuilder) {
    return i71f34fdb44622e1b8dbd50563e3b844d5b06315264bd4bf7dcfc1262b7101a4f.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ief2a3897f4132f3372d8f033cd2ebd787a4705005e05e502f64f9441d3294683.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ief2a3897f4132f3372d8f033cd2ebd787a4705005e05e502f64f9441d3294683.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i912a64a291cce9b7af41be42dd7d2c73b1358686172dc2d4200444b14724cf3c.CalendarRequestBuilder) {
    return i912a64a291cce9b7af41be42dd7d2c73b1358686172dc2d4200444b14724cf3c.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i3261245f326c44b01fe5da55ef3eb6fdcf3eec9727c6249445c533e07ad1762b.CancelRequestBuilder) {
    return i3261245f326c44b01fe5da55ef3eb6fdcf3eec9727c6249445c533e07ad1762b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/events/{event_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventRequestBuilder instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) CreateDeleteRequestInformation(options *EventRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) CreatePatchRequestInformation(options *EventRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*if65de254c43ef0974070ea922f0230c5434fa0eaef859e883b04b25ba3f75a0b.DeclineRequestBuilder) {
    return if65de254c43ef0974070ea922f0230c5434fa0eaef859e883b04b25ba3f75a0b.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Delete(options *EventRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) DismissReminder()(*i16b073d42adfaa42cd2197fe997fd666e6c717513e29bbc1411d84d135e5907d.DismissReminderRequestBuilder) {
    return i16b073d42adfaa42cd2197fe997fd666e6c717513e29bbc1411d84d135e5907d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i5ccef22bea92cabddac92714e10b7b992bc8b281f5928395c1b100cef0454e2d.ExceptionOccurrencesRequestBuilder) {
    return i5ccef22bea92cabddac92714e10b7b992bc8b281f5928395c1b100cef0454e2d.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i5cd1ea32e6df077b61718a3134fe8e9004c569fa48259151b83a3713a0993009.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i5cd1ea32e6df077b61718a3134fe8e9004c569fa48259151b83a3713a0993009.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i3cc615139480e6b6080500d2592d3b826f8172402d8123a8033c61f3b0f6d2b0.ExtensionsRequestBuilder) {
    return i3cc615139480e6b6080500d2592d3b826f8172402d8123a8033c61f3b0f6d2b0.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i35e38f254ec120c13f755d3867f2fd9952b17b0aef33b67009171560f14276ed.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i35e38f254ec120c13f755d3867f2fd9952b17b0aef33b67009171560f14276ed.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ic365f138404b6574eb593b1e6e3be13031d99faa7fde1b344c069d91fe235dbc.ForwardRequestBuilder) {
    return ic365f138404b6574eb593b1e6e3be13031d99faa7fde1b344c069d91fe235dbc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i449d1a85f5abadbfceb7b75147c603057ebd7e4b6fe7fbc07c7854cef9f748be.InstancesRequestBuilder) {
    return i449d1a85f5abadbfceb7b75147c603057ebd7e4b6fe7fbc07c7854cef9f748be.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i53b1fc78501c4055752cde988ccaa74dc9af1ec1a4f11c6c09b19a16565c5fad.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i53b1fc78501c4055752cde988ccaa74dc9af1ec1a4f11c6c09b19a16565c5fad.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ibe8940be56566d710a3152778a1869ba155c8636fa528fc2f319110c15e678a5.MultiValueExtendedPropertiesRequestBuilder) {
    return ibe8940be56566d710a3152778a1869ba155c8636fa528fc2f319110c15e678a5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib21562e8f33ec5c88876706437f50c8b6523900dedd39da40e1e91e9c54470d7.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib21562e8f33ec5c88876706437f50c8b6523900dedd39da40e1e91e9c54470d7.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the events in the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Patch(options *EventRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ie9a1738dc4818849cb5a030a886b3107bfbc115e7d40261896f7c8b5c0a1d983.SingleValueExtendedPropertiesRequestBuilder) {
    return ie9a1738dc4818849cb5a030a886b3107bfbc115e7d40261896f7c8b5c0a1d983.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i39d2408d6b0b07a8e199b2d29e86bbc35c38980add543123b882bc022c320552.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i39d2408d6b0b07a8e199b2d29e86bbc35c38980add543123b882bc022c320552.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i9fdfb1b218cc10bbbf0b6f6529dc0097d2a03fa8fa6a246d23addca54cc1138f.SnoozeReminderRequestBuilder) {
    return i9fdfb1b218cc10bbbf0b6f6529dc0097d2a03fa8fa6a246d23addca54cc1138f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i1022eeefd0950e0e70752c5cb0b720b2a7fbd1152f8b2db65dbf1f1fe3785fdd.TentativelyAcceptRequestBuilder) {
    return i1022eeefd0950e0e70752c5cb0b720b2a7fbd1152f8b2db65dbf1f1fe3785fdd.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
