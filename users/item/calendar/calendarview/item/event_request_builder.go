package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0c329946cfe9f65dc325f2a581d3900729e97e135f4fa5163e961fa1e6e6d508 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/snoozereminder"
    i0ca291604f0b37d4776c12c9384a6eeb886be25245701d1fc67067c0dacb9dab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/dismissreminder"
    i1b00a17c7bb6fd455766112e66daa7f315c3c2273539b6623fd787045e7cfc58 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/decline"
    i44e22bbc8a532c709bfd044f0c1182e6b9ac4fb8ba558f8b9d2015dc92483ee3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/tentativelyaccept"
    i4d7f4fc225d56ea372542c18aeb2c00b17087c1c476c4863be7d926d87d50330 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences"
    i816368b14e36a2a2d37e74b88ac8c6573fe1a447d99db2318bb45ac9cdc24556 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/singlevalueextendedproperties"
    id5bc622193e0269306e87950cc270ebfb747a961d3429a6e3ea98bc8222973a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/cancel"
    idbc896fd8428016fac95ec67fe116bd29913fc8de306ff8bfa377ceebcfb8100 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/attachments"
    ie314fe62a7e9af4b0d7b11457b5626b958f206d50d82443295107c58ef116f61 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances"
    ie50c98b38e8e98a9f53d53b7e33b241c4c543d0bc822b64dfff2b0c9ad8c5c9a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/extensions"
    ie8006d154169d84f839bc2404831e3bb11e537a670efbe3ea0fb678a847af840 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/multivalueextendedproperties"
    ie91f4ccb669f1f8d4f64d79430d177143bde7d6a5365d3543fe193853873802b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/forward"
    ifc93da8efa08649ecce830e5d3781181217fee6eb592ce5b00c098dbddf9dcb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/calendar"
    iff78016271ceddded3c710c7ba3b970caf8701ce868f7f32419f6cd8b5dc8fd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/accept"
    i1376f0a1866a80575b5784426dde2f70d42a6b2774cd1355d41401b86343013e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item"
    i19637f5a41e3ccc2fde05974db4440d5f14a573949916f579e495143f263fa32 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/multivalueextendedproperties/item"
    i739bfffbafaabb840a2f950e7cb9a89e6f3cf0c1eb9e4da98494b920cf8630df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item"
    i8d5280e702197db0c5ffc91ef6e1d06ee120edcff7b29550bf05c1d7e21a43d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/extensions/item"
    iad641c02387a864f90ee4d52b38266ffa5accacfea788b495524a535d81ab8e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/singlevalueextendedproperties/item"
    ic4a1b50dec9f8e4a474bc37be7aa505a20fd2b9ab74f0a81fa330b2fcc1a301b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/attachments/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendar\calendarView\{event-id}
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
// EventRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
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
func (m *EventRequestBuilder) Accept()(*iff78016271ceddded3c710c7ba3b970caf8701ce868f7f32419f6cd8b5dc8fd3.AcceptRequestBuilder) {
    return iff78016271ceddded3c710c7ba3b970caf8701ce868f7f32419f6cd8b5dc8fd3.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*idbc896fd8428016fac95ec67fe116bd29913fc8de306ff8bfa377ceebcfb8100.AttachmentsRequestBuilder) {
    return idbc896fd8428016fac95ec67fe116bd29913fc8de306ff8bfa377ceebcfb8100.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ic4a1b50dec9f8e4a474bc37be7aa505a20fd2b9ab74f0a81fa330b2fcc1a301b.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ic4a1b50dec9f8e4a474bc37be7aa505a20fd2b9ab74f0a81fa330b2fcc1a301b.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ifc93da8efa08649ecce830e5d3781181217fee6eb592ce5b00c098dbddf9dcb8.CalendarRequestBuilder) {
    return ifc93da8efa08649ecce830e5d3781181217fee6eb592ce5b00c098dbddf9dcb8.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*id5bc622193e0269306e87950cc270ebfb747a961d3429a6e3ea98bc8222973a6.CancelRequestBuilder) {
    return id5bc622193e0269306e87950cc270ebfb747a961d3429a6e3ea98bc8222973a6.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// CreateDeleteRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*i1b00a17c7bb6fd455766112e66daa7f315c3c2273539b6623fd787045e7cfc58.DeclineRequestBuilder) {
    return i1b00a17c7bb6fd455766112e66daa7f315c3c2273539b6623fd787045e7cfc58.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i0ca291604f0b37d4776c12c9384a6eeb886be25245701d1fc67067c0dacb9dab.DismissReminderRequestBuilder) {
    return i0ca291604f0b37d4776c12c9384a6eeb886be25245701d1fc67067c0dacb9dab.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i4d7f4fc225d56ea372542c18aeb2c00b17087c1c476c4863be7d926d87d50330.ExceptionOccurrencesRequestBuilder) {
    return i4d7f4fc225d56ea372542c18aeb2c00b17087c1c476c4863be7d926d87d50330.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i739bfffbafaabb840a2f950e7cb9a89e6f3cf0c1eb9e4da98494b920cf8630df.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i739bfffbafaabb840a2f950e7cb9a89e6f3cf0c1eb9e4da98494b920cf8630df.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ie50c98b38e8e98a9f53d53b7e33b241c4c543d0bc822b64dfff2b0c9ad8c5c9a.ExtensionsRequestBuilder) {
    return ie50c98b38e8e98a9f53d53b7e33b241c4c543d0bc822b64dfff2b0c9ad8c5c9a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i8d5280e702197db0c5ffc91ef6e1d06ee120edcff7b29550bf05c1d7e21a43d7.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i8d5280e702197db0c5ffc91ef6e1d06ee120edcff7b29550bf05c1d7e21a43d7.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ie91f4ccb669f1f8d4f64d79430d177143bde7d6a5365d3543fe193853873802b.ForwardRequestBuilder) {
    return ie91f4ccb669f1f8d4f64d79430d177143bde7d6a5365d3543fe193853873802b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*ie314fe62a7e9af4b0d7b11457b5626b958f206d50d82443295107c58ef116f61.InstancesRequestBuilder) {
    return ie314fe62a7e9af4b0d7b11457b5626b958f206d50d82443295107c58ef116f61.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i1376f0a1866a80575b5784426dde2f70d42a6b2774cd1355d41401b86343013e.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i1376f0a1866a80575b5784426dde2f70d42a6b2774cd1355d41401b86343013e.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ie8006d154169d84f839bc2404831e3bb11e537a670efbe3ea0fb678a847af840.MultiValueExtendedPropertiesRequestBuilder) {
    return ie8006d154169d84f839bc2404831e3bb11e537a670efbe3ea0fb678a847af840.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i19637f5a41e3ccc2fde05974db4440d5f14a573949916f579e495143f263fa32.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i19637f5a41e3ccc2fde05974db4440d5f14a573949916f579e495143f263fa32.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i816368b14e36a2a2d37e74b88ac8c6573fe1a447d99db2318bb45ac9cdc24556.SingleValueExtendedPropertiesRequestBuilder) {
    return i816368b14e36a2a2d37e74b88ac8c6573fe1a447d99db2318bb45ac9cdc24556.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iad641c02387a864f90ee4d52b38266ffa5accacfea788b495524a535d81ab8e7.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return iad641c02387a864f90ee4d52b38266ffa5accacfea788b495524a535d81ab8e7.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i0c329946cfe9f65dc325f2a581d3900729e97e135f4fa5163e961fa1e6e6d508.SnoozeReminderRequestBuilder) {
    return i0c329946cfe9f65dc325f2a581d3900729e97e135f4fa5163e961fa1e6e6d508.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i44e22bbc8a532c709bfd044f0c1182e6b9ac4fb8ba558f8b9d2015dc92483ee3.TentativelyAcceptRequestBuilder) {
    return i44e22bbc8a532c709bfd044f0c1182e6b9ac4fb8ba558f8b9d2015dc92483ee3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
