package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i17bbabd4b7e0afc45d3051f2c0350e75601541e217fe30496f14e667b8169539 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/calendar"
    i17c636c9c5b327bbcebb22b208626f3e71a0b3f502525d0b4dfd4795b41e74cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/forward"
    i34dbff51a82965b295784df99987ce16f5210a72810d56b90ed7cb42f7898dbc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/extensions"
    i36895a73e3b224fdc423d474f2e4d7dc2b463b6f0ab9b2ecab2273afb68f8d20 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i61de77e6ec9f0da71793cded3ee656f14755cd42e2cfd2ff94e678f7f12d1e8b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/attachments"
    i6962f67ff4eaf82f3e66e63d6c5c4405194c697dc18d751ee80f7c8c98e16f14 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/accept"
    i698d73ecf8a258445aae8195026eed534dbaf766a1497dd0d6b8c98eca01033a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/cancel"
    i6b73520083bc75edf4180e4f25edd94ce24af3f3359f25ca510e011234e21ea5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    ic3be1b61162600578ab76ef4fab16774e6031f3d5218e2460f310c66f798c672 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    ic97ec42fa1cf319d5887645b1f3278a6b84f36118fc066e1d12e54d6b7378393 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/decline"
    if5b87a9e479dc390dcfaa94a4c829c447d1f4e87f667db66bcd841af1d7dcf1a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    if668ead7715e702c63ea74d76571e4257105f50f768ef0b1819df9b241283a50 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i0059db08895374d8a5f7140fd0e26b1354c4a33848317fff10e940f63f7ab30d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i2ee9df08472806a82c50c687236022a2f19583b4a0cd09230781c0b0498ca850 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    i6326255e85299972c17320c8baf4927e4c2266aae8bc0f6ce848eed27dc0c173 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    ia218b8e540e39885aaa5e98492d4ee9bd524307563da15df39360093be0ec72d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i6962f67ff4eaf82f3e66e63d6c5c4405194c697dc18d751ee80f7c8c98e16f14.AcceptRequestBuilder) {
    return i6962f67ff4eaf82f3e66e63d6c5c4405194c697dc18d751ee80f7c8c98e16f14.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i61de77e6ec9f0da71793cded3ee656f14755cd42e2cfd2ff94e678f7f12d1e8b.AttachmentsRequestBuilder) {
    return i61de77e6ec9f0da71793cded3ee656f14755cd42e2cfd2ff94e678f7f12d1e8b.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i0059db08895374d8a5f7140fd0e26b1354c4a33848317fff10e940f63f7ab30d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i0059db08895374d8a5f7140fd0e26b1354c4a33848317fff10e940f63f7ab30d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i17bbabd4b7e0afc45d3051f2c0350e75601541e217fe30496f14e667b8169539.CalendarRequestBuilder) {
    return i17bbabd4b7e0afc45d3051f2c0350e75601541e217fe30496f14e667b8169539.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i698d73ecf8a258445aae8195026eed534dbaf766a1497dd0d6b8c98eca01033a.CancelRequestBuilder) {
    return i698d73ecf8a258445aae8195026eed534dbaf766a1497dd0d6b8c98eca01033a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/events/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
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
// CreateDeleteRequestInformation delete navigation property instances for me
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
// CreatePatchRequestInformation update the navigation property instances in me
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
func (m *EventItemRequestBuilder) Decline()(*ic97ec42fa1cf319d5887645b1f3278a6b84f36118fc066e1d12e54d6b7378393.DeclineRequestBuilder) {
    return ic97ec42fa1cf319d5887645b1f3278a6b84f36118fc066e1d12e54d6b7378393.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*if668ead7715e702c63ea74d76571e4257105f50f768ef0b1819df9b241283a50.DismissReminderRequestBuilder) {
    return if668ead7715e702c63ea74d76571e4257105f50f768ef0b1819df9b241283a50.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i34dbff51a82965b295784df99987ce16f5210a72810d56b90ed7cb42f7898dbc.ExtensionsRequestBuilder) {
    return i34dbff51a82965b295784df99987ce16f5210a72810d56b90ed7cb42f7898dbc.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i2ee9df08472806a82c50c687236022a2f19583b4a0cd09230781c0b0498ca850.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i2ee9df08472806a82c50c687236022a2f19583b4a0cd09230781c0b0498ca850.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i17c636c9c5b327bbcebb22b208626f3e71a0b3f502525d0b4dfd4795b41e74cc.ForwardRequestBuilder) {
    return i17c636c9c5b327bbcebb22b208626f3e71a0b3f502525d0b4dfd4795b41e74cc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic3be1b61162600578ab76ef4fab16774e6031f3d5218e2460f310c66f798c672.MultiValueExtendedPropertiesRequestBuilder) {
    return ic3be1b61162600578ab76ef4fab16774e6031f3d5218e2460f310c66f798c672.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6326255e85299972c17320c8baf4927e4c2266aae8bc0f6ce848eed27dc0c173.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i6326255e85299972c17320c8baf4927e4c2266aae8bc0f6ce848eed27dc0c173.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if5b87a9e479dc390dcfaa94a4c829c447d1f4e87f667db66bcd841af1d7dcf1a.SingleValueExtendedPropertiesRequestBuilder) {
    return if5b87a9e479dc390dcfaa94a4c829c447d1f4e87f667db66bcd841af1d7dcf1a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia218b8e540e39885aaa5e98492d4ee9bd524307563da15df39360093be0ec72d.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ia218b8e540e39885aaa5e98492d4ee9bd524307563da15df39360093be0ec72d.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i36895a73e3b224fdc423d474f2e4d7dc2b463b6f0ab9b2ecab2273afb68f8d20.SnoozeReminderRequestBuilder) {
    return i36895a73e3b224fdc423d474f2e4d7dc2b463b6f0ab9b2ecab2273afb68f8d20.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6b73520083bc75edf4180e4f25edd94ce24af3f3359f25ca510e011234e21ea5.TentativelyAcceptRequestBuilder) {
    return i6b73520083bc75edf4180e4f25edd94ce24af3f3359f25ca510e011234e21ea5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
