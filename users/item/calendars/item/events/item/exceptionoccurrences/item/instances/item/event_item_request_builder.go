package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0b9f80f8fac79c3a9cf2e537f859b1c2e21df1fa7ae49ddb0e041c048e2a9b92 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i1ed714b778d8b7d2dea09689275b50005e4d6639baaec71a0dbe9aa6adc2f31f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/accept"
    i22dbf9f458bfeb5e2a193184182a4c810cd7d72886f8040446e2625460755d45 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i4895c52c8b3aef844d9b277586d6bd877f0354270c0fb4615e2752e505d9435d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i53e0452291f095707b5942cca00bac9131e66dd37adfcea3833107ea99f4b1ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i77f20c4b4594205b842d02086d88fd51932204c7169f06e344b6189cb87d468a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    ia22b68672e63e164db65f2d6433cd00e6dfef6588859a76f352301ecbade573c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    ic2577072b12a75e0f45116330ffea1ed21cd5ce0880c405b313b87ef925038e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    icde9fc17d59fb2667ba21b41bb23a93f3e7ef492ef80db24a749927f783187a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    id87abc09906ded7d7954d4fd07f8614f5ef2e06fe29a2001e0aad0a02e6cd304 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/forward"
    ie01b911721cc06fe9220cc6e311e48277d2c5d1e83c2f85c954f45d058023887 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ieb053bc43aed21e38a0dc172dfa07fcdf56eb34b1a45324f6429b41778dcc174 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/decline"
    i1b58f0aecd662360ac86db0cb8d5dc6847453e51a07519e2eaeecf55c28f53c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i5787a60d94ced5c7eacb9c5b114c730d82e74fe1740d214aaeeb47a39a2fe71d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i8e3cf8a70d3ff34d13e9fb27e223a5ac2f39b2dc223bf83bc8e66b97e1132129 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    i9349f129ad1cf4d0c89e359bfda1097c31ea5c39cbc03c7c1aa4b773be138b31 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i1ed714b778d8b7d2dea09689275b50005e4d6639baaec71a0dbe9aa6adc2f31f.AcceptRequestBuilder) {
    return i1ed714b778d8b7d2dea09689275b50005e4d6639baaec71a0dbe9aa6adc2f31f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i4895c52c8b3aef844d9b277586d6bd877f0354270c0fb4615e2752e505d9435d.AttachmentsRequestBuilder) {
    return i4895c52c8b3aef844d9b277586d6bd877f0354270c0fb4615e2752e505d9435d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i5787a60d94ced5c7eacb9c5b114c730d82e74fe1740d214aaeeb47a39a2fe71d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i5787a60d94ced5c7eacb9c5b114c730d82e74fe1740d214aaeeb47a39a2fe71d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i77f20c4b4594205b842d02086d88fd51932204c7169f06e344b6189cb87d468a.CalendarRequestBuilder) {
    return i77f20c4b4594205b842d02086d88fd51932204c7169f06e344b6189cb87d468a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i22dbf9f458bfeb5e2a193184182a4c810cd7d72886f8040446e2625460755d45.CancelRequestBuilder) {
    return i22dbf9f458bfeb5e2a193184182a4c810cd7d72886f8040446e2625460755d45.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/events/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
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
// CreateDeleteRequestInformation delete navigation property instances for users
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
// CreatePatchRequestInformation update the navigation property instances in users
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
func (m *EventItemRequestBuilder) Decline()(*ieb053bc43aed21e38a0dc172dfa07fcdf56eb34b1a45324f6429b41778dcc174.DeclineRequestBuilder) {
    return ieb053bc43aed21e38a0dc172dfa07fcdf56eb34b1a45324f6429b41778dcc174.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*ia22b68672e63e164db65f2d6433cd00e6dfef6588859a76f352301ecbade573c.DismissReminderRequestBuilder) {
    return ia22b68672e63e164db65f2d6433cd00e6dfef6588859a76f352301ecbade573c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i0b9f80f8fac79c3a9cf2e537f859b1c2e21df1fa7ae49ddb0e041c048e2a9b92.ExtensionsRequestBuilder) {
    return i0b9f80f8fac79c3a9cf2e537f859b1c2e21df1fa7ae49ddb0e041c048e2a9b92.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8e3cf8a70d3ff34d13e9fb27e223a5ac2f39b2dc223bf83bc8e66b97e1132129.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i8e3cf8a70d3ff34d13e9fb27e223a5ac2f39b2dc223bf83bc8e66b97e1132129.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*id87abc09906ded7d7954d4fd07f8614f5ef2e06fe29a2001e0aad0a02e6cd304.ForwardRequestBuilder) {
    return id87abc09906ded7d7954d4fd07f8614f5ef2e06fe29a2001e0aad0a02e6cd304.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic2577072b12a75e0f45116330ffea1ed21cd5ce0880c405b313b87ef925038e6.MultiValueExtendedPropertiesRequestBuilder) {
    return ic2577072b12a75e0f45116330ffea1ed21cd5ce0880c405b313b87ef925038e6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9349f129ad1cf4d0c89e359bfda1097c31ea5c39cbc03c7c1aa4b773be138b31.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i9349f129ad1cf4d0c89e359bfda1097c31ea5c39cbc03c7c1aa4b773be138b31.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i53e0452291f095707b5942cca00bac9131e66dd37adfcea3833107ea99f4b1ef.SingleValueExtendedPropertiesRequestBuilder) {
    return i53e0452291f095707b5942cca00bac9131e66dd37adfcea3833107ea99f4b1ef.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i1b58f0aecd662360ac86db0cb8d5dc6847453e51a07519e2eaeecf55c28f53c6.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i1b58f0aecd662360ac86db0cb8d5dc6847453e51a07519e2eaeecf55c28f53c6.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie01b911721cc06fe9220cc6e311e48277d2c5d1e83c2f85c954f45d058023887.SnoozeReminderRequestBuilder) {
    return ie01b911721cc06fe9220cc6e311e48277d2c5d1e83c2f85c954f45d058023887.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*icde9fc17d59fb2667ba21b41bb23a93f3e7ef492ef80db24a749927f783187a3.TentativelyAcceptRequestBuilder) {
    return icde9fc17d59fb2667ba21b41bb23a93f3e7ef492ef80db24a749927f783187a3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
