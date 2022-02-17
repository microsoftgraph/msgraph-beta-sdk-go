package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i07fe24d1dec166fbef03b00ab5b46e976baffb6c06b5ef5e6347d446ff25bca3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/extensions"
    i5dbf769c1f25e1b0b17d30509a4a645cfa242a1c4d7c348c060dccc53cf78dd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/multivalueextendedproperties"
    i7cda4870e61701a73f65b06cb1c454e1fa2c00f69cd309e8d501fd4a70e14a22 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/dismissreminder"
    i86ef866cc2d984e3066c38be80d8b147c286544fa4fecd6ab59fe1d8ff248c1d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/decline"
    i8f07c66b2cf1a3f6d8da9f8cf5d8442d87cb10892aad82c5dd44556cc8ec85bb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/forward"
    i97b528d359cdfaf66013514b9af5ce682328c510bc03248fb0810ebf226f19e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/attachments"
    ib2a51806722e21c6ec788fdf6b0eee10e907a4dc8e354b7d914a32f4368d0864 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/accept"
    ibd61d302d825b6e73918b5027ff3b14b0b17c3c96a153d5cffabf8cf98cbd736 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/tentativelyaccept"
    icae7b5ff991f4631b46cc3802546c55d85ad10eaeb0238bf2f6e3ff7bf554f0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/cancel"
    idf6baae4d48e4fcb9e8ccf745a9b91eb8ec7f7f41d6f2c4074fca23712f7aa70 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/calendar"
    ie4cb09c95e0c35a55010b7327a11926471dad8ff28619f453f661c22aee3f94a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/singlevalueextendedproperties"
    ie506b5a0873b4dc903852664779e9434b8ef3c1ada875aa7f0551abf90de2d3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/snoozereminder"
    if0032e7643530222733c0d1a82c01d7975494927aa723dfe9e7505cb06dc839e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences"
    if607cd43935ec64edb92794109173f3c0d5910d8dfebec67b6e22ea9bffd7433 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances"
    i12298c30cb81e9f690ef9152f8caab34fa4d1864ebd36ed5577020082370bd73 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/extensions/item"
    i5956806f1c4b998abe4fcd54db1586d7465f29b5e18161bba26e275d033eab11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item"
    i73e700365b2c2d1e86684192b02ce3e5f7775ba91bb65b0fe4eb04ec97e7b592 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item"
    i916f3cc62a9632691b92c06ae4a196bfbbab686ba1d7e0ec5f5b48c10806fd6b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/singlevalueextendedproperties/item"
    ia14123345e5038209c9a2a037b69700ff79c35780a03e5454afed921703f8320 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/attachments/item"
    if2b1e2a9807085574644316d2e02a3a9aa2335ecc40c7e9a8828e945fbc62be9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/multivalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendar\events\{event-id}
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
func (m *EventRequestBuilder) Accept()(*ib2a51806722e21c6ec788fdf6b0eee10e907a4dc8e354b7d914a32f4368d0864.AcceptRequestBuilder) {
    return ib2a51806722e21c6ec788fdf6b0eee10e907a4dc8e354b7d914a32f4368d0864.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i97b528d359cdfaf66013514b9af5ce682328c510bc03248fb0810ebf226f19e4.AttachmentsRequestBuilder) {
    return i97b528d359cdfaf66013514b9af5ce682328c510bc03248fb0810ebf226f19e4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ia14123345e5038209c9a2a037b69700ff79c35780a03e5454afed921703f8320.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia14123345e5038209c9a2a037b69700ff79c35780a03e5454afed921703f8320.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*idf6baae4d48e4fcb9e8ccf745a9b91eb8ec7f7f41d6f2c4074fca23712f7aa70.CalendarRequestBuilder) {
    return idf6baae4d48e4fcb9e8ccf745a9b91eb8ec7f7f41d6f2c4074fca23712f7aa70.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*icae7b5ff991f4631b46cc3802546c55d85ad10eaeb0238bf2f6e3ff7bf554f0e.CancelRequestBuilder) {
    return icae7b5ff991f4631b46cc3802546c55d85ad10eaeb0238bf2f6e3ff7bf554f0e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/events/{event_id}{?select}";
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
func (m *EventRequestBuilder) Decline()(*i86ef866cc2d984e3066c38be80d8b147c286544fa4fecd6ab59fe1d8ff248c1d.DeclineRequestBuilder) {
    return i86ef866cc2d984e3066c38be80d8b147c286544fa4fecd6ab59fe1d8ff248c1d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i7cda4870e61701a73f65b06cb1c454e1fa2c00f69cd309e8d501fd4a70e14a22.DismissReminderRequestBuilder) {
    return i7cda4870e61701a73f65b06cb1c454e1fa2c00f69cd309e8d501fd4a70e14a22.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*if0032e7643530222733c0d1a82c01d7975494927aa723dfe9e7505cb06dc839e.ExceptionOccurrencesRequestBuilder) {
    return if0032e7643530222733c0d1a82c01d7975494927aa723dfe9e7505cb06dc839e.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i73e700365b2c2d1e86684192b02ce3e5f7775ba91bb65b0fe4eb04ec97e7b592.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i73e700365b2c2d1e86684192b02ce3e5f7775ba91bb65b0fe4eb04ec97e7b592.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i07fe24d1dec166fbef03b00ab5b46e976baffb6c06b5ef5e6347d446ff25bca3.ExtensionsRequestBuilder) {
    return i07fe24d1dec166fbef03b00ab5b46e976baffb6c06b5ef5e6347d446ff25bca3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i12298c30cb81e9f690ef9152f8caab34fa4d1864ebd36ed5577020082370bd73.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i12298c30cb81e9f690ef9152f8caab34fa4d1864ebd36ed5577020082370bd73.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i8f07c66b2cf1a3f6d8da9f8cf5d8442d87cb10892aad82c5dd44556cc8ec85bb.ForwardRequestBuilder) {
    return i8f07c66b2cf1a3f6d8da9f8cf5d8442d87cb10892aad82c5dd44556cc8ec85bb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*if607cd43935ec64edb92794109173f3c0d5910d8dfebec67b6e22ea9bffd7433.InstancesRequestBuilder) {
    return if607cd43935ec64edb92794109173f3c0d5910d8dfebec67b6e22ea9bffd7433.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i5956806f1c4b998abe4fcd54db1586d7465f29b5e18161bba26e275d033eab11.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i5956806f1c4b998abe4fcd54db1586d7465f29b5e18161bba26e275d033eab11.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i5dbf769c1f25e1b0b17d30509a4a645cfa242a1c4d7c348c060dccc53cf78dd6.MultiValueExtendedPropertiesRequestBuilder) {
    return i5dbf769c1f25e1b0b17d30509a4a645cfa242a1c4d7c348c060dccc53cf78dd6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if2b1e2a9807085574644316d2e02a3a9aa2335ecc40c7e9a8828e945fbc62be9.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return if2b1e2a9807085574644316d2e02a3a9aa2335ecc40c7e9a8828e945fbc62be9.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ie4cb09c95e0c35a55010b7327a11926471dad8ff28619f453f661c22aee3f94a.SingleValueExtendedPropertiesRequestBuilder) {
    return ie4cb09c95e0c35a55010b7327a11926471dad8ff28619f453f661c22aee3f94a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i916f3cc62a9632691b92c06ae4a196bfbbab686ba1d7e0ec5f5b48c10806fd6b.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i916f3cc62a9632691b92c06ae4a196bfbbab686ba1d7e0ec5f5b48c10806fd6b.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ie506b5a0873b4dc903852664779e9434b8ef3c1ada875aa7f0551abf90de2d3a.SnoozeReminderRequestBuilder) {
    return ie506b5a0873b4dc903852664779e9434b8ef3c1ada875aa7f0551abf90de2d3a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ibd61d302d825b6e73918b5027ff3b14b0b17c3c96a153d5cffabf8cf98cbd736.TentativelyAcceptRequestBuilder) {
    return ibd61d302d825b6e73918b5027ff3b14b0b17c3c96a153d5cffabf8cf98cbd736.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
