package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i1392060414ef191e570bf2441a06bf0aa02916b747db286a28fb7dc3c2c92e09 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i2138bb645bcb8a4c21f46dd5c1c182aae362eb9ad12f1032dc2a2a80b4234d40 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i4b74cf1475be79d8c5bd5a96a857dfee3b4a001104265d86db42bc4ccdb139f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i5a9277e9b29a102a4ac15815f43c255e33b63212fd83179cfb06ba9c0095a417 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i5c74e0eceb86a90abe58ab377605dfd9a92559258fc4e477917da5c5ec800911 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i638365c56d4a3a9d8db8ef7a8f8cb86cb08f3af8ba56b38f8f1b65bdf27df1bf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i6bc39607fa012755f044c52399c42bb807b680526e3542b3df42d1f675ce2edc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i759d7b0b434060f3aaed7df848350d3517136226a983ee0c485acde76a53fed9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ic2e3264dc769002999eaf12fe5eec70ab157b8c926d2fd9e8e8b474cc4aed762 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    ice2f5260d34d3c032174d7c833245dd7952e01872c8c372f831f6ccda3b49aa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    ie00181941fec3c9663ffb77f6f64ef400b43cab518310760d1e5894cee015858 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    if08749e141f1fdc527f44f4b0e1c235ad9ab3135dd7323ba322054e32b36a03d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i0067dc107b6785a4cbaa9516df188178c3ed39eca93dd9f677c93380cc0a0d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i4235545d855e808026b5af63ffee955a888c59e75d548a24093ddf21321260f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i700bc3bbf501af9afe6e7f3fbfa81a604caccf7870a64a2b888153813b9c749a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    ia0ef39ad909bd7275646fee8900363fa776a983b1e646976e57cfe18a7c5bcbd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) Accept()(*ice2f5260d34d3c032174d7c833245dd7952e01872c8c372f831f6ccda3b49aa9.AcceptRequestBuilder) {
    return ice2f5260d34d3c032174d7c833245dd7952e01872c8c372f831f6ccda3b49aa9.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i5c74e0eceb86a90abe58ab377605dfd9a92559258fc4e477917da5c5ec800911.AttachmentsRequestBuilder) {
    return i5c74e0eceb86a90abe58ab377605dfd9a92559258fc4e477917da5c5ec800911.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i700bc3bbf501af9afe6e7f3fbfa81a604caccf7870a64a2b888153813b9c749a.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i700bc3bbf501af9afe6e7f3fbfa81a604caccf7870a64a2b888153813b9c749a.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i638365c56d4a3a9d8db8ef7a8f8cb86cb08f3af8ba56b38f8f1b65bdf27df1bf.CalendarRequestBuilder) {
    return i638365c56d4a3a9d8db8ef7a8f8cb86cb08f3af8ba56b38f8f1b65bdf27df1bf.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i1392060414ef191e570bf2441a06bf0aa02916b747db286a28fb7dc3c2c92e09.CancelRequestBuilder) {
    return i1392060414ef191e570bf2441a06bf0aa02916b747db286a28fb7dc3c2c92e09.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/calendarView/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for users
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
// CreateGetRequestInformation get exceptionOccurrences from users
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) Decline()(*i6bc39607fa012755f044c52399c42bb807b680526e3542b3df42d1f675ce2edc.DeclineRequestBuilder) {
    return i6bc39607fa012755f044c52399c42bb807b680526e3542b3df42d1f675ce2edc.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*ic2e3264dc769002999eaf12fe5eec70ab157b8c926d2fd9e8e8b474cc4aed762.DismissReminderRequestBuilder) {
    return ic2e3264dc769002999eaf12fe5eec70ab157b8c926d2fd9e8e8b474cc4aed762.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i5a9277e9b29a102a4ac15815f43c255e33b63212fd83179cfb06ba9c0095a417.ExtensionsRequestBuilder) {
    return i5a9277e9b29a102a4ac15815f43c255e33b63212fd83179cfb06ba9c0095a417.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia0ef39ad909bd7275646fee8900363fa776a983b1e646976e57cfe18a7c5bcbd.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ia0ef39ad909bd7275646fee8900363fa776a983b1e646976e57cfe18a7c5bcbd.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i4b74cf1475be79d8c5bd5a96a857dfee3b4a001104265d86db42bc4ccdb139f2.ForwardRequestBuilder) {
    return i4b74cf1475be79d8c5bd5a96a857dfee3b4a001104265d86db42bc4ccdb139f2.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie00181941fec3c9663ffb77f6f64ef400b43cab518310760d1e5894cee015858.MultiValueExtendedPropertiesRequestBuilder) {
    return ie00181941fec3c9663ffb77f6f64ef400b43cab518310760d1e5894cee015858.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0067dc107b6785a4cbaa9516df188178c3ed39eca93dd9f677c93380cc0a0d1a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0067dc107b6785a4cbaa9516df188178c3ed39eca93dd9f677c93380cc0a0d1a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i2138bb645bcb8a4c21f46dd5c1c182aae362eb9ad12f1032dc2a2a80b4234d40.SingleValueExtendedPropertiesRequestBuilder) {
    return i2138bb645bcb8a4c21f46dd5c1c182aae362eb9ad12f1032dc2a2a80b4234d40.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4235545d855e808026b5af63ffee955a888c59e75d548a24093ddf21321260f9.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i4235545d855e808026b5af63ffee955a888c59e75d548a24093ddf21321260f9.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*if08749e141f1fdc527f44f4b0e1c235ad9ab3135dd7323ba322054e32b36a03d.SnoozeReminderRequestBuilder) {
    return if08749e141f1fdc527f44f4b0e1c235ad9ab3135dd7323ba322054e32b36a03d.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i759d7b0b434060f3aaed7df848350d3517136226a983ee0c485acde76a53fed9.TentativelyAcceptRequestBuilder) {
    return i759d7b0b434060f3aaed7df848350d3517136226a983ee0c485acde76a53fed9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}