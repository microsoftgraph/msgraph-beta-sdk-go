package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i01fb2df2892ac0830c87fcf62df419a5b957e82a94a89a349b90c39f6326679f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i13c2c43066b385d189c402f1f3f02386c685d77e6f9642b8d003c4bd88889914 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i15bbfef361fe7f43c4a0cd40a59b5ec6588061e6a4197451c9519398b88a3cfa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i699849b0e1615887fb1c6d4922bcb1386bae6fd4b85c8e86f0c9da8ab8202e40 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i6feed7c975422a5f9086f766d03ab0a24ed1b17fba89878fafb57aabe69c93dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i72df339a7df0d82419938509403b21bf282ed3744852880a72066041563e6758 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i799529df9f53556c8af69e2520763915522836d6aafef34252ada69b68d59e9c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    ia1da8b1ea95847818016a9356e6270d909e6d418a0dfd0e51e8de07b93f58774 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ic1ad4ca23474b81e1226dfd7992a27af68a71109111a888e15e745ba61403c56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    icc15b479516c13d95b6f684dfe0766b6706c9d7d143ab550d41c2ec0a10dfd63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    ide723326b393119009a6f0c17c949555227f138f4d770bd2a2b909f1d33c93df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    iebfd6610409ba87469bf28d1f2ca772c68e234f4525294e6e0ef11f925f69988 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i2c01bdb15aebb4c846849d6e815ae936248297bb97e58ed479654e0c45204217 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    i72e02a1bf8eaa5d8a8bbb72de7ef1af36bf0bcdb78c896a4043e45a41db53468 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ic7bcb62e34588a967c7d97ac3c1f396f8903ff050ce91b506399f6b2a748bec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    ie6d99adf383d06bf9aead85124fcf4eccb5db52694831a69544b844c86e96fea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ide723326b393119009a6f0c17c949555227f138f4d770bd2a2b909f1d33c93df.AcceptRequestBuilder) {
    return ide723326b393119009a6f0c17c949555227f138f4d770bd2a2b909f1d33c93df.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i13c2c43066b385d189c402f1f3f02386c685d77e6f9642b8d003c4bd88889914.AttachmentsRequestBuilder) {
    return i13c2c43066b385d189c402f1f3f02386c685d77e6f9642b8d003c4bd88889914.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i2c01bdb15aebb4c846849d6e815ae936248297bb97e58ed479654e0c45204217.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i2c01bdb15aebb4c846849d6e815ae936248297bb97e58ed479654e0c45204217.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i699849b0e1615887fb1c6d4922bcb1386bae6fd4b85c8e86f0c9da8ab8202e40.CalendarRequestBuilder) {
    return i699849b0e1615887fb1c6d4922bcb1386bae6fd4b85c8e86f0c9da8ab8202e40.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ic1ad4ca23474b81e1226dfd7992a27af68a71109111a888e15e745ba61403c56.CancelRequestBuilder) {
    return ic1ad4ca23474b81e1226dfd7992a27af68a71109111a888e15e745ba61403c56.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/calendarView/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i799529df9f53556c8af69e2520763915522836d6aafef34252ada69b68d59e9c.DeclineRequestBuilder) {
    return i799529df9f53556c8af69e2520763915522836d6aafef34252ada69b68d59e9c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i15bbfef361fe7f43c4a0cd40a59b5ec6588061e6a4197451c9519398b88a3cfa.DismissReminderRequestBuilder) {
    return i15bbfef361fe7f43c4a0cd40a59b5ec6588061e6a4197451c9519398b88a3cfa.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*icc15b479516c13d95b6f684dfe0766b6706c9d7d143ab550d41c2ec0a10dfd63.ExtensionsRequestBuilder) {
    return icc15b479516c13d95b6f684dfe0766b6706c9d7d143ab550d41c2ec0a10dfd63.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic7bcb62e34588a967c7d97ac3c1f396f8903ff050ce91b506399f6b2a748bec6.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic7bcb62e34588a967c7d97ac3c1f396f8903ff050ce91b506399f6b2a748bec6.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i01fb2df2892ac0830c87fcf62df419a5b957e82a94a89a349b90c39f6326679f.ForwardRequestBuilder) {
    return i01fb2df2892ac0830c87fcf62df419a5b957e82a94a89a349b90c39f6326679f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iebfd6610409ba87469bf28d1f2ca772c68e234f4525294e6e0ef11f925f69988.MultiValueExtendedPropertiesRequestBuilder) {
    return iebfd6610409ba87469bf28d1f2ca772c68e234f4525294e6e0ef11f925f69988.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i72e02a1bf8eaa5d8a8bbb72de7ef1af36bf0bcdb78c896a4043e45a41db53468.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i72e02a1bf8eaa5d8a8bbb72de7ef1af36bf0bcdb78c896a4043e45a41db53468.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia1da8b1ea95847818016a9356e6270d909e6d418a0dfd0e51e8de07b93f58774.SingleValueExtendedPropertiesRequestBuilder) {
    return ia1da8b1ea95847818016a9356e6270d909e6d418a0dfd0e51e8de07b93f58774.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie6d99adf383d06bf9aead85124fcf4eccb5db52694831a69544b844c86e96fea.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ie6d99adf383d06bf9aead85124fcf4eccb5db52694831a69544b844c86e96fea.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i72df339a7df0d82419938509403b21bf282ed3744852880a72066041563e6758.SnoozeReminderRequestBuilder) {
    return i72df339a7df0d82419938509403b21bf282ed3744852880a72066041563e6758.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6feed7c975422a5f9086f766d03ab0a24ed1b17fba89878fafb57aabe69c93dd.TentativelyAcceptRequestBuilder) {
    return i6feed7c975422a5f9086f766d03ab0a24ed1b17fba89878fafb57aabe69c93dd.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
