package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0062b298e2ffb7e864aa5cf6d632cd823ed9566653be5f6e8889231b84ce4ac6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i130975e0dfdb91688100b9e129e8f10467053f976e74ac87dc6d4854a54d2269 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i2eade44e5de8ec4344095b0746f2d6960e63a87bcf9112e70c1211232da9cdd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i4eea483bb3b07759a8e23e0b484f1387be1d4a551467220bf81b7cacb9229b14 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i5fd86f654e0f3b974006e93b1a60dfa42ec4b2a4c807a4f2d19c0db3f187d8ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i6f26a813084343024cbe1c3286df0b8f2f26416eba0a1775a482cdce89fb7dbc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i7bd7e43ce7d9bea156f2b9631c5f4813d23f9c1845297c6ae689fa909021aaf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    ia9a4ad63c67b1c65d087f4752cd189fa43de9d39f8f63d194308ed8fc36359f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    ib0d47db12e79afad843f2e24dc210d801526edb509046a1b212735c8e5becf63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    icce2f9aa51075b8ad56182a928bc9ad87c78b053af67e5c992385b87dc201261 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    icf61d9d88f414ffb4011d10abea9efba1294d34aa148f99e68a74aff7f08aa65 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    idcf18fa49cece6031f6eacf311f33950e2dcf50388e29fa130e3faa25cbf55cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i120cbebfaf8ba152ed4efd5cf712e7e7b635f0772df6ec8110dc5b4ab24ff598 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    i88adc21805da074a5a0b41c23c3b02be605d362660d5cdb45436d56c427209fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i9d03403c0367899199389a99a42e39415397ebb34b7e1daa91b81bed69ba97aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ida5e3f89ae23131e5b26fed27044589841ae5730177f129f83f50a7bff858f54 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
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
func (m *EventItemRequestBuilder) Accept()(*ia9a4ad63c67b1c65d087f4752cd189fa43de9d39f8f63d194308ed8fc36359f6.AcceptRequestBuilder) {
    return ia9a4ad63c67b1c65d087f4752cd189fa43de9d39f8f63d194308ed8fc36359f6.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i130975e0dfdb91688100b9e129e8f10467053f976e74ac87dc6d4854a54d2269.AttachmentsRequestBuilder) {
    return i130975e0dfdb91688100b9e129e8f10467053f976e74ac87dc6d4854a54d2269.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i120cbebfaf8ba152ed4efd5cf712e7e7b635f0772df6ec8110dc5b4ab24ff598.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i120cbebfaf8ba152ed4efd5cf712e7e7b635f0772df6ec8110dc5b4ab24ff598.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ib0d47db12e79afad843f2e24dc210d801526edb509046a1b212735c8e5becf63.CalendarRequestBuilder) {
    return ib0d47db12e79afad843f2e24dc210d801526edb509046a1b212735c8e5becf63.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*idcf18fa49cece6031f6eacf311f33950e2dcf50388e29fa130e3faa25cbf55cd.CancelRequestBuilder) {
    return idcf18fa49cece6031f6eacf311f33950e2dcf50388e29fa130e3faa25cbf55cd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for me
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) Decline()(*icce2f9aa51075b8ad56182a928bc9ad87c78b053af67e5c992385b87dc201261.DeclineRequestBuilder) {
    return icce2f9aa51075b8ad56182a928bc9ad87c78b053af67e5c992385b87dc201261.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i2eade44e5de8ec4344095b0746f2d6960e63a87bcf9112e70c1211232da9cdd5.DismissReminderRequestBuilder) {
    return i2eade44e5de8ec4344095b0746f2d6960e63a87bcf9112e70c1211232da9cdd5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*icf61d9d88f414ffb4011d10abea9efba1294d34aa148f99e68a74aff7f08aa65.ExtensionsRequestBuilder) {
    return icf61d9d88f414ffb4011d10abea9efba1294d34aa148f99e68a74aff7f08aa65.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ida5e3f89ae23131e5b26fed27044589841ae5730177f129f83f50a7bff858f54.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ida5e3f89ae23131e5b26fed27044589841ae5730177f129f83f50a7bff858f54.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i0062b298e2ffb7e864aa5cf6d632cd823ed9566653be5f6e8889231b84ce4ac6.ForwardRequestBuilder) {
    return i0062b298e2ffb7e864aa5cf6d632cd823ed9566653be5f6e8889231b84ce4ac6.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4eea483bb3b07759a8e23e0b484f1387be1d4a551467220bf81b7cacb9229b14.MultiValueExtendedPropertiesRequestBuilder) {
    return i4eea483bb3b07759a8e23e0b484f1387be1d4a551467220bf81b7cacb9229b14.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9d03403c0367899199389a99a42e39415397ebb34b7e1daa91b81bed69ba97aa.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i9d03403c0367899199389a99a42e39415397ebb34b7e1daa91b81bed69ba97aa.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i5fd86f654e0f3b974006e93b1a60dfa42ec4b2a4c807a4f2d19c0db3f187d8ee.SingleValueExtendedPropertiesRequestBuilder) {
    return i5fd86f654e0f3b974006e93b1a60dfa42ec4b2a4c807a4f2d19c0db3f187d8ee.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i88adc21805da074a5a0b41c23c3b02be605d362660d5cdb45436d56c427209fa.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i88adc21805da074a5a0b41c23c3b02be605d362660d5cdb45436d56c427209fa.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i7bd7e43ce7d9bea156f2b9631c5f4813d23f9c1845297c6ae689fa909021aaf7.SnoozeReminderRequestBuilder) {
    return i7bd7e43ce7d9bea156f2b9631c5f4813d23f9c1845297c6ae689fa909021aaf7.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6f26a813084343024cbe1c3286df0b8f2f26416eba0a1775a482cdce89fb7dbc.TentativelyAcceptRequestBuilder) {
    return i6f26a813084343024cbe1c3286df0b8f2f26416eba0a1775a482cdce89fb7dbc.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
