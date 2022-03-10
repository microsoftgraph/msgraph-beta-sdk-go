package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i0ada765736df380496639152c96e85c08fa5694a11f604c5f7796cd7bdf41b8d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/decline"
    i379526d3bf7aed5b3d384bf8f27ad0c0a9f8f005463f0a29be80cf7ebd556e77 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/forward"
    i703a58619607122a48d03bd9e139c80af32322f09e198e79b1d3c34c522fcf7b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/cancel"
    i74e20a858042c3b06ec93d51b4cc1fe2a5b91c4cce20b5f732ad7230d532d904 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/extensions"
    i764e7f2f5f8c03269b0f61bbe3cc34b00547f2ea0f526d822316ce462b1cfc8e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/snoozereminder"
    i766249a3a7449417cdc0f889fe91c6f81a4bbf753193e2ce804dd79fd7ee1c39 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/accept"
    i7a162b4c8628bb486f34601f3830c8a54b386639bf323bda94860b399a0351c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/multivalueextendedproperties"
    i9cb3bd5a363a546a32a1b6cbd06eb82d2f7d2bb0c4683ae0efc0ea758f375617 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/singlevalueextendedproperties"
    ia5b83a086e3503c0e5f64ce6bd87c0afdd942ec3a6f20053c49e8117b9eae4a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/tentativelyaccept"
    ib45ff153b6c9e630ce2b424e49604634b173747fb72b74c01cbae60865f5f67a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/attachments"
    ibcc6cdd07105848a4e7f312773f7c56cd76001b195e070731b6b75d111cd8910 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences"
    id79528d2908c97845234f0872cdf44c9c308d6b359e3167e3aa9d9f6c7a69def "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/dismissreminder"
    ied46e231558ee92e395d1342ff6006b6048d1a38a3d9e7f2760f2df151d6414b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/calendar"
    i5d295821cd556be50acc39040c786214b9f54d0655c7edda01907cc12a82b273 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/multivalueextendedproperties/item"
    ibe56593adc981184a6a79bb399418c1d2c2b699e9249d67a27d5f85b8688b3a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item"
    idb89d33a8e1ddbae9cec4785046d3035dc4556f27a636ef8dc755e025ea31ff9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/attachments/item"
    ie1adb99a9a8d1e2e1dbcadb5f04792bcbfbe8480cf5fd059a5f2f5a2dbce2fb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/singlevalueextendedproperties/item"
    if57e644c3ae5776b03ad8b3946980614baa56e29823a40c4f98c68452a9201a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*i766249a3a7449417cdc0f889fe91c6f81a4bbf753193e2ce804dd79fd7ee1c39.AcceptRequestBuilder) {
    return i766249a3a7449417cdc0f889fe91c6f81a4bbf753193e2ce804dd79fd7ee1c39.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ib45ff153b6c9e630ce2b424e49604634b173747fb72b74c01cbae60865f5f67a.AttachmentsRequestBuilder) {
    return ib45ff153b6c9e630ce2b424e49604634b173747fb72b74c01cbae60865f5f67a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*idb89d33a8e1ddbae9cec4785046d3035dc4556f27a636ef8dc755e025ea31ff9.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return idb89d33a8e1ddbae9cec4785046d3035dc4556f27a636ef8dc755e025ea31ff9.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ied46e231558ee92e395d1342ff6006b6048d1a38a3d9e7f2760f2df151d6414b.CalendarRequestBuilder) {
    return ied46e231558ee92e395d1342ff6006b6048d1a38a3d9e7f2760f2df151d6414b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i703a58619607122a48d03bd9e139c80af32322f09e198e79b1d3c34c522fcf7b.CancelRequestBuilder) {
    return i703a58619607122a48d03bd9e139c80af32322f09e198e79b1d3c34c522fcf7b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i0ada765736df380496639152c96e85c08fa5694a11f604c5f7796cd7bdf41b8d.DeclineRequestBuilder) {
    return i0ada765736df380496639152c96e85c08fa5694a11f604c5f7796cd7bdf41b8d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*id79528d2908c97845234f0872cdf44c9c308d6b359e3167e3aa9d9f6c7a69def.DismissReminderRequestBuilder) {
    return id79528d2908c97845234f0872cdf44c9c308d6b359e3167e3aa9d9f6c7a69def.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ibcc6cdd07105848a4e7f312773f7c56cd76001b195e070731b6b75d111cd8910.ExceptionOccurrencesRequestBuilder) {
    return ibcc6cdd07105848a4e7f312773f7c56cd76001b195e070731b6b75d111cd8910.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ibe56593adc981184a6a79bb399418c1d2c2b699e9249d67a27d5f85b8688b3a7.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return ibe56593adc981184a6a79bb399418c1d2c2b699e9249d67a27d5f85b8688b3a7.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i74e20a858042c3b06ec93d51b4cc1fe2a5b91c4cce20b5f732ad7230d532d904.ExtensionsRequestBuilder) {
    return i74e20a858042c3b06ec93d51b4cc1fe2a5b91c4cce20b5f732ad7230d532d904.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*if57e644c3ae5776b03ad8b3946980614baa56e29823a40c4f98c68452a9201a2.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return if57e644c3ae5776b03ad8b3946980614baa56e29823a40c4f98c68452a9201a2.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i379526d3bf7aed5b3d384bf8f27ad0c0a9f8f005463f0a29be80cf7ebd556e77.ForwardRequestBuilder) {
    return i379526d3bf7aed5b3d384bf8f27ad0c0a9f8f005463f0a29be80cf7ebd556e77.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i7a162b4c8628bb486f34601f3830c8a54b386639bf323bda94860b399a0351c4.MultiValueExtendedPropertiesRequestBuilder) {
    return i7a162b4c8628bb486f34601f3830c8a54b386639bf323bda94860b399a0351c4.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5d295821cd556be50acc39040c786214b9f54d0655c7edda01907cc12a82b273.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5d295821cd556be50acc39040c786214b9f54d0655c7edda01907cc12a82b273.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i9cb3bd5a363a546a32a1b6cbd06eb82d2f7d2bb0c4683ae0efc0ea758f375617.SingleValueExtendedPropertiesRequestBuilder) {
    return i9cb3bd5a363a546a32a1b6cbd06eb82d2f7d2bb0c4683ae0efc0ea758f375617.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie1adb99a9a8d1e2e1dbcadb5f04792bcbfbe8480cf5fd059a5f2f5a2dbce2fb0.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ie1adb99a9a8d1e2e1dbcadb5f04792bcbfbe8480cf5fd059a5f2f5a2dbce2fb0.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i764e7f2f5f8c03269b0f61bbe3cc34b00547f2ea0f526d822316ce462b1cfc8e.SnoozeReminderRequestBuilder) {
    return i764e7f2f5f8c03269b0f61bbe3cc34b00547f2ea0f526d822316ce462b1cfc8e.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ia5b83a086e3503c0e5f64ce6bd87c0afdd942ec3a6f20053c49e8117b9eae4a5.TentativelyAcceptRequestBuilder) {
    return ia5b83a086e3503c0e5f64ce6bd87c0afdd942ec3a6f20053c49e8117b9eae4a5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
