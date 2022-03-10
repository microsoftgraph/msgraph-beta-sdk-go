package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i06614cf97b25e4f1cc95b099db1bff5aa2d15eb9fe48ceeec06d30f041cec7b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i20818696d027c8339014df26f5cff39fca939fea5e71a491efd9e8a8feb34d96 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i245e54c9299bc73f21d507de0abc77a91a3ee273ede39d36f8cc4b30d3d929d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/cancel"
    i2c6f242e034d183656c5cbb51a72e875198d972c7654e24ca256a037cd60fa2a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/attachments"
    i342c127c12be2e6048b7dff13a836a0b2d04c12eba2f598c2e211d0dd0666e5a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/accept"
    i380e854f772bf5bc9a128e8415ee4ad551ae07887f6a329120e013966758c29c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/forward"
    i7968aa204c89c1bd548200a458a991bebe04ca6a16068768c1269ebf170e26f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/extensions"
    ibd54d8108ed0df1725e8c7a096107a962956e60e3b9b4d0851e4b09aa0d4bd1e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances"
    ic5673a6bd5616076eb617c73ae999646f4b1e1254759dd671a8f249fb54a8a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/dismissreminder"
    id0d2742f6d706f93431cd309866bd747cae6f6fb721136a8478052f87a40452d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/decline"
    ide9cc24b38c043ff97d2882e83ae818dc56b4d7ed8bfa6dab5ac8b8e766435da "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/calendar"
    ie2bc1bf78859cacf27c559fad849551bf1cd6552993f78d6f1e7e4c484dfc5d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    ie48d50a4a462cf7ba0136119b50d82579a435f443e9085a25c31f262598ced12 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i15f5819471788997963c36e490f6a51e9f5550aa4436173f156b761ef03a91ff "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/extensions/item"
    i30151f5bb8a6840dfd789ab1d79f1791912103c60c88280ee2b9070461ca730f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i50f22ae6af4b1db80a0da3d8151a9ddaba5fcae0f9669dd5503fe1945048a7b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ic98255458c7c86b6d4c11d49c572809d991cd1cfd656b375f4e29a936576cdd4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item"
    ifa2e8f16b45156022dbea12e9f778c0713f79a848df0dd18efd079d168eba76e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/attachments/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Accept()(*i342c127c12be2e6048b7dff13a836a0b2d04c12eba2f598c2e211d0dd0666e5a.AcceptRequestBuilder) {
    return i342c127c12be2e6048b7dff13a836a0b2d04c12eba2f598c2e211d0dd0666e5a.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i2c6f242e034d183656c5cbb51a72e875198d972c7654e24ca256a037cd60fa2a.AttachmentsRequestBuilder) {
    return i2c6f242e034d183656c5cbb51a72e875198d972c7654e24ca256a037cd60fa2a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ifa2e8f16b45156022dbea12e9f778c0713f79a848df0dd18efd079d168eba76e.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ifa2e8f16b45156022dbea12e9f778c0713f79a848df0dd18efd079d168eba76e.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ide9cc24b38c043ff97d2882e83ae818dc56b4d7ed8bfa6dab5ac8b8e766435da.CalendarRequestBuilder) {
    return ide9cc24b38c043ff97d2882e83ae818dc56b4d7ed8bfa6dab5ac8b8e766435da.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i245e54c9299bc73f21d507de0abc77a91a3ee273ede39d36f8cc4b30d3d929d0.CancelRequestBuilder) {
    return i245e54c9299bc73f21d507de0abc77a91a3ee273ede39d36f8cc4b30d3d929d0.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/calendarView/{event_id}/exceptionOccurrences/{event_id1}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for groups
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
// CreateGetRequestInformation get exceptionOccurrences from groups
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in groups
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
func (m *EventItemRequestBuilder) Decline()(*id0d2742f6d706f93431cd309866bd747cae6f6fb721136a8478052f87a40452d.DeclineRequestBuilder) {
    return id0d2742f6d706f93431cd309866bd747cae6f6fb721136a8478052f87a40452d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*ic5673a6bd5616076eb617c73ae999646f4b1e1254759dd671a8f249fb54a8a37.DismissReminderRequestBuilder) {
    return ic5673a6bd5616076eb617c73ae999646f4b1e1254759dd671a8f249fb54a8a37.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i7968aa204c89c1bd548200a458a991bebe04ca6a16068768c1269ebf170e26f3.ExtensionsRequestBuilder) {
    return i7968aa204c89c1bd548200a458a991bebe04ca6a16068768c1269ebf170e26f3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i15f5819471788997963c36e490f6a51e9f5550aa4436173f156b761ef03a91ff.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i15f5819471788997963c36e490f6a51e9f5550aa4436173f156b761ef03a91ff.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i380e854f772bf5bc9a128e8415ee4ad551ae07887f6a329120e013966758c29c.ForwardRequestBuilder) {
    return i380e854f772bf5bc9a128e8415ee4ad551ae07887f6a329120e013966758c29c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Instances()(*ibd54d8108ed0df1725e8c7a096107a962956e60e3b9b4d0851e4b09aa0d4bd1e.InstancesRequestBuilder) {
    return ibd54d8108ed0df1725e8c7a096107a962956e60e3b9b4d0851e4b09aa0d4bd1e.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ic98255458c7c86b6d4c11d49c572809d991cd1cfd656b375f4e29a936576cdd4.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return ic98255458c7c86b6d4c11d49c572809d991cd1cfd656b375f4e29a936576cdd4.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie2bc1bf78859cacf27c559fad849551bf1cd6552993f78d6f1e7e4c484dfc5d0.MultiValueExtendedPropertiesRequestBuilder) {
    return ie2bc1bf78859cacf27c559fad849551bf1cd6552993f78d6f1e7e4c484dfc5d0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i50f22ae6af4b1db80a0da3d8151a9ddaba5fcae0f9669dd5503fe1945048a7b5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i50f22ae6af4b1db80a0da3d8151a9ddaba5fcae0f9669dd5503fe1945048a7b5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i20818696d027c8339014df26f5cff39fca939fea5e71a491efd9e8a8feb34d96.SingleValueExtendedPropertiesRequestBuilder) {
    return i20818696d027c8339014df26f5cff39fca939fea5e71a491efd9e8a8feb34d96.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i30151f5bb8a6840dfd789ab1d79f1791912103c60c88280ee2b9070461ca730f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i30151f5bb8a6840dfd789ab1d79f1791912103c60c88280ee2b9070461ca730f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i06614cf97b25e4f1cc95b099db1bff5aa2d15eb9fe48ceeec06d30f041cec7b3.SnoozeReminderRequestBuilder) {
    return i06614cf97b25e4f1cc95b099db1bff5aa2d15eb9fe48ceeec06d30f041cec7b3.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ie48d50a4a462cf7ba0136119b50d82579a435f443e9085a25c31f262598ced12.TentativelyAcceptRequestBuilder) {
    return ie48d50a4a462cf7ba0136119b50d82579a435f443e9085a25c31f262598ced12.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
