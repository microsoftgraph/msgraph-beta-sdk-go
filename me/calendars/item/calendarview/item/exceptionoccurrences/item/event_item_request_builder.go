package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i1aa418bf331327a659bb0f32332d28f62ab63752666539a7cf47e919f7893a90 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/forward"
    i2b590e8188cb62f0d0507032e3fc005f1fa0198caa0b2c631653e2bb9be5763d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/calendar"
    i4289ce2674bfb1acf7e13a2172bb68f483e5fc47e6b2fa219b0aef0d147240ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    i4383183ffea8f4c5769f6500a11d17f34633bc6dbd43569fbb6d5957f82210cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/cancel"
    i55e2fb9e1894a140fe83226946bd723401b8ef0fcaa94c379b46f41f5cd6aec0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i7d8b6f8f1bea8e5a3c38e9d318b814fc4a976a012a7a3b8578a572d2b8eec7cf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/dismissreminder"
    iaca54173f05380f3018de48ca9fb3022722d85707c953870e4865e7a0e5a44ca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/attachments"
    id65df392189d92c808919db0c5a2ce2853e0d16197cb40450aebcb4a6c6b018f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    id9f4a56bec987ed0a8df8d454790fb86facf9aff7c2f2d974f816c31c0b7eab1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances"
    idf9e9d4f08f0e447e6fe08ce4b1772706c78c73a75c3f37d57526a5d61ffa930 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/accept"
    ie2d7e68438cae9387ea2671fd9e98de2291fe31d55e24d18d9f8497a8022d4cf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/decline"
    ie4c1b4b94974a554fee856ef1b687bd849c870d9f74ef2b1d77e3b42ed9a0307 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/extensions"
    ie8c4dc60d53be27e14111c89821c074b63729660e7a9dc7e7ebe8c1bf06d417d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i045ec5be0c83060670c48a7ba40d93b8f8a18cc8fc77ba3c77c93eb3b30724f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i0465a24ffed9a2dbc02873ffc2894a6d821a3c3eb48c93e16709e537310d2e46 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/extensions/item"
    i7da262cb2f9511b88ac19b3fe075289e4b07a46755bf5622bc61d222e44091ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i8f4c1f9d63a10d955d056120e32a26065c23b29aa6c9c2a6c32f199c81f789b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item"
    ica3af7a4d17ffae0defd0dacca718db93c518b5cd2e348546a5b001b5b681dd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*idf9e9d4f08f0e447e6fe08ce4b1772706c78c73a75c3f37d57526a5d61ffa930.AcceptRequestBuilder) {
    return idf9e9d4f08f0e447e6fe08ce4b1772706c78c73a75c3f37d57526a5d61ffa930.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*iaca54173f05380f3018de48ca9fb3022722d85707c953870e4865e7a0e5a44ca.AttachmentsRequestBuilder) {
    return iaca54173f05380f3018de48ca9fb3022722d85707c953870e4865e7a0e5a44ca.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ica3af7a4d17ffae0defd0dacca718db93c518b5cd2e348546a5b001b5b681dd8.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ica3af7a4d17ffae0defd0dacca718db93c518b5cd2e348546a5b001b5b681dd8.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i2b590e8188cb62f0d0507032e3fc005f1fa0198caa0b2c631653e2bb9be5763d.CalendarRequestBuilder) {
    return i2b590e8188cb62f0d0507032e3fc005f1fa0198caa0b2c631653e2bb9be5763d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i4383183ffea8f4c5769f6500a11d17f34633bc6dbd43569fbb6d5957f82210cd.CancelRequestBuilder) {
    return i4383183ffea8f4c5769f6500a11d17f34633bc6dbd43569fbb6d5957f82210cd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/calendarView/{event_id}/exceptionOccurrences/{event_id1}{?select,expand}";
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
func (m *EventItemRequestBuilder) Decline()(*ie2d7e68438cae9387ea2671fd9e98de2291fe31d55e24d18d9f8497a8022d4cf.DeclineRequestBuilder) {
    return ie2d7e68438cae9387ea2671fd9e98de2291fe31d55e24d18d9f8497a8022d4cf.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i7d8b6f8f1bea8e5a3c38e9d318b814fc4a976a012a7a3b8578a572d2b8eec7cf.DismissReminderRequestBuilder) {
    return i7d8b6f8f1bea8e5a3c38e9d318b814fc4a976a012a7a3b8578a572d2b8eec7cf.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*ie4c1b4b94974a554fee856ef1b687bd849c870d9f74ef2b1d77e3b42ed9a0307.ExtensionsRequestBuilder) {
    return ie4c1b4b94974a554fee856ef1b687bd849c870d9f74ef2b1d77e3b42ed9a0307.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i0465a24ffed9a2dbc02873ffc2894a6d821a3c3eb48c93e16709e537310d2e46.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i0465a24ffed9a2dbc02873ffc2894a6d821a3c3eb48c93e16709e537310d2e46.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i1aa418bf331327a659bb0f32332d28f62ab63752666539a7cf47e919f7893a90.ForwardRequestBuilder) {
    return i1aa418bf331327a659bb0f32332d28f62ab63752666539a7cf47e919f7893a90.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Instances()(*id9f4a56bec987ed0a8df8d454790fb86facf9aff7c2f2d974f816c31c0b7eab1.InstancesRequestBuilder) {
    return id9f4a56bec987ed0a8df8d454790fb86facf9aff7c2f2d974f816c31c0b7eab1.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i8f4c1f9d63a10d955d056120e32a26065c23b29aa6c9c2a6c32f199c81f789b7.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return i8f4c1f9d63a10d955d056120e32a26065c23b29aa6c9c2a6c32f199c81f789b7.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4289ce2674bfb1acf7e13a2172bb68f483e5fc47e6b2fa219b0aef0d147240ed.MultiValueExtendedPropertiesRequestBuilder) {
    return i4289ce2674bfb1acf7e13a2172bb68f483e5fc47e6b2fa219b0aef0d147240ed.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i045ec5be0c83060670c48a7ba40d93b8f8a18cc8fc77ba3c77c93eb3b30724f4.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i045ec5be0c83060670c48a7ba40d93b8f8a18cc8fc77ba3c77c93eb3b30724f4.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i55e2fb9e1894a140fe83226946bd723401b8ef0fcaa94c379b46f41f5cd6aec0.SingleValueExtendedPropertiesRequestBuilder) {
    return i55e2fb9e1894a140fe83226946bd723401b8ef0fcaa94c379b46f41f5cd6aec0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7da262cb2f9511b88ac19b3fe075289e4b07a46755bf5622bc61d222e44091ee.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7da262cb2f9511b88ac19b3fe075289e4b07a46755bf5622bc61d222e44091ee.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie8c4dc60d53be27e14111c89821c074b63729660e7a9dc7e7ebe8c1bf06d417d.SnoozeReminderRequestBuilder) {
    return ie8c4dc60d53be27e14111c89821c074b63729660e7a9dc7e7ebe8c1bf06d417d.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*id65df392189d92c808919db0c5a2ce2853e0d16197cb40450aebcb4a6c6b018f.TentativelyAcceptRequestBuilder) {
    return id65df392189d92c808919db0c5a2ce2853e0d16197cb40450aebcb4a6c6b018f.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
