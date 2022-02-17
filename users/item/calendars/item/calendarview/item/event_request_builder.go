package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i09464dfde46ea9b5614331a441462441658822ee493fbccebc73230d1194d842 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences"
    i0ae79ce850436f0587525bcaec70b835a2cef9711988a12de9cb93e24414b430 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/snoozereminder"
    i0b099560e4b2efb93ab572b92e7f1a76ec97b902d61ab9becc2defcef681ac25 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/dismissreminder"
    i2bcbda220eee3efb3381726cc4041421ea7e20a03d0304ac407dea06fa354f27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    i4b428c5cf9a69e44aa972695ae72775d3e28ccc810fe60f6b821b483200c436d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances"
    i5debd82ad007fe3e7b83a2eeb29c28f5da3b9563dd9ca463f34104a008b9eee9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/calendar"
    i88e3730e4bce338f224d440e746456a0359fb801e5e58784dbfe1799783810bf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/forward"
    i8c425fcf68846ce1fd43b66e601e98461e65f7444df3be897e0da2f43a8df63c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/cancel"
    i9dd05f797cf579bc4dc0e79bc00cc79d1351d71707d91b30098d9d3ae9ea28b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/attachments"
    iac7fd37b42a89ce1966b5f0ea72249cc45d99faf64966f63c5ebe88241edbbe9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/extensions"
    iaf2491bc5b9ff1ab39daeb0d781b18170fe54aa51fb51480ae6422705aa157c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/accept"
    ic31b3a2cd538f90211d1e9cf6de4bb7092f7aacee83586f1ab97621d45d8b815 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/decline"
    id6ac304b8cd98e6aec12a2df7396b7ef53304b70df9d56b58094a06bb593da74 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/multivalueextendedproperties"
    ifd62164a910c345faf910bed37d632b88189cfa74bdd65db738cf2f49dabf4d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/tentativelyaccept"
    i0f9b7621256a1b0f4c88adedcc6a9f1bbfd62cef422da8242f344ebe0128feda "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/extensions/item"
    i47e3e498cae25dc161c5f4d25d399c625874e2e39d1692d4216a0b825941c4ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    i91c7117765263d0e129264c87641bfc35cd44799a05cde4e3181a86d7ab46ccd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item"
    ib0b41a465e6f0cb684c983aa2ebdcc8e6295fb076929a3ad3e741eb6fc148feb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
    ib21fde6c453d56158f7bf6413ec85dcf8c8b15c7d91c031a7cccedeacb0d80c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item"
    ief3b9e1dfba882027be7bcb7ecb89490b77e45266ac9933e953debb528098d78 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/attachments/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendars\{calendar-id}\calendarView\{event-id}
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
// EventRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
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
func (m *EventRequestBuilder) Accept()(*iaf2491bc5b9ff1ab39daeb0d781b18170fe54aa51fb51480ae6422705aa157c8.AcceptRequestBuilder) {
    return iaf2491bc5b9ff1ab39daeb0d781b18170fe54aa51fb51480ae6422705aa157c8.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i9dd05f797cf579bc4dc0e79bc00cc79d1351d71707d91b30098d9d3ae9ea28b5.AttachmentsRequestBuilder) {
    return i9dd05f797cf579bc4dc0e79bc00cc79d1351d71707d91b30098d9d3ae9ea28b5.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ief3b9e1dfba882027be7bcb7ecb89490b77e45266ac9933e953debb528098d78.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ief3b9e1dfba882027be7bcb7ecb89490b77e45266ac9933e953debb528098d78.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i5debd82ad007fe3e7b83a2eeb29c28f5da3b9563dd9ca463f34104a008b9eee9.CalendarRequestBuilder) {
    return i5debd82ad007fe3e7b83a2eeb29c28f5da3b9563dd9ca463f34104a008b9eee9.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i8c425fcf68846ce1fd43b66e601e98461e65f7444df3be897e0da2f43a8df63c.CancelRequestBuilder) {
    return i8c425fcf68846ce1fd43b66e601e98461e65f7444df3be897e0da2f43a8df63c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// CreateDeleteRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*ic31b3a2cd538f90211d1e9cf6de4bb7092f7aacee83586f1ab97621d45d8b815.DeclineRequestBuilder) {
    return ic31b3a2cd538f90211d1e9cf6de4bb7092f7aacee83586f1ab97621d45d8b815.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i0b099560e4b2efb93ab572b92e7f1a76ec97b902d61ab9becc2defcef681ac25.DismissReminderRequestBuilder) {
    return i0b099560e4b2efb93ab572b92e7f1a76ec97b902d61ab9becc2defcef681ac25.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i09464dfde46ea9b5614331a441462441658822ee493fbccebc73230d1194d842.ExceptionOccurrencesRequestBuilder) {
    return i09464dfde46ea9b5614331a441462441658822ee493fbccebc73230d1194d842.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i91c7117765263d0e129264c87641bfc35cd44799a05cde4e3181a86d7ab46ccd.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i91c7117765263d0e129264c87641bfc35cd44799a05cde4e3181a86d7ab46ccd.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*iac7fd37b42a89ce1966b5f0ea72249cc45d99faf64966f63c5ebe88241edbbe9.ExtensionsRequestBuilder) {
    return iac7fd37b42a89ce1966b5f0ea72249cc45d99faf64966f63c5ebe88241edbbe9.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i0f9b7621256a1b0f4c88adedcc6a9f1bbfd62cef422da8242f344ebe0128feda.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i0f9b7621256a1b0f4c88adedcc6a9f1bbfd62cef422da8242f344ebe0128feda.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i88e3730e4bce338f224d440e746456a0359fb801e5e58784dbfe1799783810bf.ForwardRequestBuilder) {
    return i88e3730e4bce338f224d440e746456a0359fb801e5e58784dbfe1799783810bf.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*i4b428c5cf9a69e44aa972695ae72775d3e28ccc810fe60f6b821b483200c436d.InstancesRequestBuilder) {
    return i4b428c5cf9a69e44aa972695ae72775d3e28ccc810fe60f6b821b483200c436d.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*ib21fde6c453d56158f7bf6413ec85dcf8c8b15c7d91c031a7cccedeacb0d80c9.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ib21fde6c453d56158f7bf6413ec85dcf8c8b15c7d91c031a7cccedeacb0d80c9.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*id6ac304b8cd98e6aec12a2df7396b7ef53304b70df9d56b58094a06bb593da74.MultiValueExtendedPropertiesRequestBuilder) {
    return id6ac304b8cd98e6aec12a2df7396b7ef53304b70df9d56b58094a06bb593da74.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i47e3e498cae25dc161c5f4d25d399c625874e2e39d1692d4216a0b825941c4ed.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i47e3e498cae25dc161c5f4d25d399c625874e2e39d1692d4216a0b825941c4ed.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i2bcbda220eee3efb3381726cc4041421ea7e20a03d0304ac407dea06fa354f27.SingleValueExtendedPropertiesRequestBuilder) {
    return i2bcbda220eee3efb3381726cc4041421ea7e20a03d0304ac407dea06fa354f27.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib0b41a465e6f0cb684c983aa2ebdcc8e6295fb076929a3ad3e741eb6fc148feb.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib0b41a465e6f0cb684c983aa2ebdcc8e6295fb076929a3ad3e741eb6fc148feb.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i0ae79ce850436f0587525bcaec70b835a2cef9711988a12de9cb93e24414b430.SnoozeReminderRequestBuilder) {
    return i0ae79ce850436f0587525bcaec70b835a2cef9711988a12de9cb93e24414b430.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ifd62164a910c345faf910bed37d632b88189cfa74bdd65db738cf2f49dabf4d1.TentativelyAcceptRequestBuilder) {
    return ifd62164a910c345faf910bed37d632b88189cfa74bdd65db738cf2f49dabf4d1.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
