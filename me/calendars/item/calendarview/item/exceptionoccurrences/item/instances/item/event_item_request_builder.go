package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1e5c164ac79c2516796935548a2d598bb836d0c0121a070215d73f5b74af4eb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    i2427cf2cf754d500e6fb9ab893a18d3f8b7d34839d40a4b5a5778f684fa85ef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i27846051f36a9f23e9a8a3c3a0ac1e658add43b2ec45aa5897cc1b2302e9b71f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i2d3d49f75b76349cc59e17a34560f312dd11757a37fbe1a56e39c132a5cf7533 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    i2d5bede201267a087678d08b3f08698789c55af7dd1117f7faed5a9218e24cd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i56d844f656b61154c0d99f0e1c111e9c9b90f7fa2e5a8cd9e80bb1cf1a2a4ca3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i5e2a3eea2e47f371e756a926a73d334429c5361186b34c7cbf8f24220b6c54fb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    i7cb7bc461416e62cbca42aa51409df7f5c19c5cf858a4f73901ce242b1b5995c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i82c03a019e3f923a7ab010850cc4fbc97fb775beac1122c58b1c89e76ba2c384 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    ic596d1187abbe8f29cc367123e5db8052adc9d68655d2d86441ecdd0b597ae31 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    id7be5befed1fd9874bc578ad69553a692a1b2132f44f25fe8b91a07274f05a36 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ifabbc968a17ac48c8c8fe818c8ffaafaf853ea3c1b46f0c4d872a13b6ab34327 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    i7abe2a6f1c5284004871e717c40452effde59f4381c532808626eca64fe62a72 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    ib30f91e1e26cfc36a5cd425cf7c426f9028a7802d92cd6572d87739676b62260 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    ib541d9a069049274503971e07bb2de50449cb5bb30f76a0b0f8caf2dc4b7301b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    if5673d0066ed3a8d8691249fa9a9f06fd5fdff1c36e6df8bfe1b5f0f14b08add "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i82c03a019e3f923a7ab010850cc4fbc97fb775beac1122c58b1c89e76ba2c384.AcceptRequestBuilder) {
    return i82c03a019e3f923a7ab010850cc4fbc97fb775beac1122c58b1c89e76ba2c384.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ic596d1187abbe8f29cc367123e5db8052adc9d68655d2d86441ecdd0b597ae31.AttachmentsRequestBuilder) {
    return ic596d1187abbe8f29cc367123e5db8052adc9d68655d2d86441ecdd0b597ae31.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib30f91e1e26cfc36a5cd425cf7c426f9028a7802d92cd6572d87739676b62260.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ib30f91e1e26cfc36a5cd425cf7c426f9028a7802d92cd6572d87739676b62260.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ifabbc968a17ac48c8c8fe818c8ffaafaf853ea3c1b46f0c4d872a13b6ab34327.CalendarRequestBuilder) {
    return ifabbc968a17ac48c8c8fe818c8ffaafaf853ea3c1b46f0c4d872a13b6ab34327.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i27846051f36a9f23e9a8a3c3a0ac1e658add43b2ec45aa5897cc1b2302e9b71f.CancelRequestBuilder) {
    return i27846051f36a9f23e9a8a3c3a0ac1e658add43b2ec45aa5897cc1b2302e9b71f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/calendarView/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i1e5c164ac79c2516796935548a2d598bb836d0c0121a070215d73f5b74af4eb3.DeclineRequestBuilder) {
    return i1e5c164ac79c2516796935548a2d598bb836d0c0121a070215d73f5b74af4eb3.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i2427cf2cf754d500e6fb9ab893a18d3f8b7d34839d40a4b5a5778f684fa85ef6.DismissReminderRequestBuilder) {
    return i2427cf2cf754d500e6fb9ab893a18d3f8b7d34839d40a4b5a5778f684fa85ef6.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i2d3d49f75b76349cc59e17a34560f312dd11757a37fbe1a56e39c132a5cf7533.ExtensionsRequestBuilder) {
    return i2d3d49f75b76349cc59e17a34560f312dd11757a37fbe1a56e39c132a5cf7533.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i7abe2a6f1c5284004871e717c40452effde59f4381c532808626eca64fe62a72.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i7abe2a6f1c5284004871e717c40452effde59f4381c532808626eca64fe62a72.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i5e2a3eea2e47f371e756a926a73d334429c5361186b34c7cbf8f24220b6c54fb.ForwardRequestBuilder) {
    return i5e2a3eea2e47f371e756a926a73d334429c5361186b34c7cbf8f24220b6c54fb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i7cb7bc461416e62cbca42aa51409df7f5c19c5cf858a4f73901ce242b1b5995c.MultiValueExtendedPropertiesRequestBuilder) {
    return i7cb7bc461416e62cbca42aa51409df7f5c19c5cf858a4f73901ce242b1b5995c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib541d9a069049274503971e07bb2de50449cb5bb30f76a0b0f8caf2dc4b7301b.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib541d9a069049274503971e07bb2de50449cb5bb30f76a0b0f8caf2dc4b7301b.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i2d5bede201267a087678d08b3f08698789c55af7dd1117f7faed5a9218e24cd8.SingleValueExtendedPropertiesRequestBuilder) {
    return i2d5bede201267a087678d08b3f08698789c55af7dd1117f7faed5a9218e24cd8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if5673d0066ed3a8d8691249fa9a9f06fd5fdff1c36e6df8bfe1b5f0f14b08add.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if5673d0066ed3a8d8691249fa9a9f06fd5fdff1c36e6df8bfe1b5f0f14b08add.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*id7be5befed1fd9874bc578ad69553a692a1b2132f44f25fe8b91a07274f05a36.SnoozeReminderRequestBuilder) {
    return id7be5befed1fd9874bc578ad69553a692a1b2132f44f25fe8b91a07274f05a36.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i56d844f656b61154c0d99f0e1c111e9c9b90f7fa2e5a8cd9e80bb1cf1a2a4ca3.TentativelyAcceptRequestBuilder) {
    return i56d844f656b61154c0d99f0e1c111e9c9b90f7fa2e5a8cd9e80bb1cf1a2a4ca3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
