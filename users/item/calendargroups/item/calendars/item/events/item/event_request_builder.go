package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i25de4e25092b7e755133987907dde688413aeca8870d959074f2d5a29b215781 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/calendar"
    i32ed368ce5ee12bfd01e4b23d9aa3f8f0a53f6b58ebcba5ff956c739eae37b41 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties"
    i5a22110dcba479e1eff1b0c29b99c5190def78fa9a90e3dd939e20db4a0d09a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/tentativelyaccept"
    i5e8df9fbcb3652f57651023f51df8a28d379b1376ab981cb84ef6aab28544051 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/multivalueextendedproperties"
    i62828a10360c0b31d201e58ffa11ef310f19d0858ade98c0b16d15ad0f526ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/forward"
    i6d0fc705adb0d375d549d85597118478eea44ed85a34e1950b70ca1d40ee2263 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/snoozereminder"
    iae504bcbafa7e7be1fcbdb6c9e41f1aa2973c490c6635c4f814c6f3028746793 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/cancel"
    ib4d53f5f0a9c528946157ab89269d7751b4b8b78a5c47d34fee1c32f6c90666e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences"
    ib52fa7b6ab178bc9ca1668bc1f4528d57028a61a8512bd987abdbcc172cb752c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances"
    ib9a5addb47b9731b352af65c12df457afa53f7c42ba5324b4fee1645ea69a8d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/accept"
    ic7c445d04ee5fee0a0fb119ac2f55d9c0abac409a27e95723d315a2fe1546355 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/decline"
    ic7ced464a710579a513ff3b400431f16298c62341fee23031a0922695d61671d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/extensions"
    ie2d0e0c974228c2b4059332b63a9a93487ca86935ea8ca735848ba174a09d839 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/attachments"
    if515108435b34c8569b4ad6dc8db7c465e9a50c988cc32a3f51d4c541cb6d284 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/dismissreminder"
    i1034f1aca18807ae946587c94d7baea7e891feccc77cd9c27536801c3d38edfd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/attachments/item"
    i14a109fa65988885cbc9233c0f6bbb1830d25a2157a41eaec2043a62d8a956bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item"
    i6b26c9235a27137c3d11dddffd6b9e0a151857c05bb49489992064460bbfe0f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/extensions/item"
    i6cf98b0ba15623a64ec587ad6a80d5690b2a34c0ba7e0016c5f13c28d749dfc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/multivalueextendedproperties/item"
    i8adf45a44048de2db7e88293c362739387a38abd925b2c94e5fab81ca24cd81b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item"
    id8a23edf2d82fd94fc95e02eb04d228bf11f41aa3452bbd39108f71fb7b96f2c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\events\{event-id}
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
// EventRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
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
func (m *EventRequestBuilder) Accept()(*ib9a5addb47b9731b352af65c12df457afa53f7c42ba5324b4fee1645ea69a8d9.AcceptRequestBuilder) {
    return ib9a5addb47b9731b352af65c12df457afa53f7c42ba5324b4fee1645ea69a8d9.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ie2d0e0c974228c2b4059332b63a9a93487ca86935ea8ca735848ba174a09d839.AttachmentsRequestBuilder) {
    return ie2d0e0c974228c2b4059332b63a9a93487ca86935ea8ca735848ba174a09d839.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*i1034f1aca18807ae946587c94d7baea7e891feccc77cd9c27536801c3d38edfd.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i1034f1aca18807ae946587c94d7baea7e891feccc77cd9c27536801c3d38edfd.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i25de4e25092b7e755133987907dde688413aeca8870d959074f2d5a29b215781.CalendarRequestBuilder) {
    return i25de4e25092b7e755133987907dde688413aeca8870d959074f2d5a29b215781.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*iae504bcbafa7e7be1fcbdb6c9e41f1aa2973c490c6635c4f814c6f3028746793.CancelRequestBuilder) {
    return iae504bcbafa7e7be1fcbdb6c9e41f1aa2973c490c6635c4f814c6f3028746793.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*ic7c445d04ee5fee0a0fb119ac2f55d9c0abac409a27e95723d315a2fe1546355.DeclineRequestBuilder) {
    return ic7c445d04ee5fee0a0fb119ac2f55d9c0abac409a27e95723d315a2fe1546355.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*if515108435b34c8569b4ad6dc8db7c465e9a50c988cc32a3f51d4c541cb6d284.DismissReminderRequestBuilder) {
    return if515108435b34c8569b4ad6dc8db7c465e9a50c988cc32a3f51d4c541cb6d284.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*ib4d53f5f0a9c528946157ab89269d7751b4b8b78a5c47d34fee1c32f6c90666e.ExceptionOccurrencesRequestBuilder) {
    return ib4d53f5f0a9c528946157ab89269d7751b4b8b78a5c47d34fee1c32f6c90666e.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i8adf45a44048de2db7e88293c362739387a38abd925b2c94e5fab81ca24cd81b.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i8adf45a44048de2db7e88293c362739387a38abd925b2c94e5fab81ca24cd81b.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ic7ced464a710579a513ff3b400431f16298c62341fee23031a0922695d61671d.ExtensionsRequestBuilder) {
    return ic7ced464a710579a513ff3b400431f16298c62341fee23031a0922695d61671d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i6b26c9235a27137c3d11dddffd6b9e0a151857c05bb49489992064460bbfe0f9.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i6b26c9235a27137c3d11dddffd6b9e0a151857c05bb49489992064460bbfe0f9.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i62828a10360c0b31d201e58ffa11ef310f19d0858ade98c0b16d15ad0f526ddd.ForwardRequestBuilder) {
    return i62828a10360c0b31d201e58ffa11ef310f19d0858ade98c0b16d15ad0f526ddd.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*ib52fa7b6ab178bc9ca1668bc1f4528d57028a61a8512bd987abdbcc172cb752c.InstancesRequestBuilder) {
    return ib52fa7b6ab178bc9ca1668bc1f4528d57028a61a8512bd987abdbcc172cb752c.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i14a109fa65988885cbc9233c0f6bbb1830d25a2157a41eaec2043a62d8a956bc.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i14a109fa65988885cbc9233c0f6bbb1830d25a2157a41eaec2043a62d8a956bc.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i5e8df9fbcb3652f57651023f51df8a28d379b1376ab981cb84ef6aab28544051.MultiValueExtendedPropertiesRequestBuilder) {
    return i5e8df9fbcb3652f57651023f51df8a28d379b1376ab981cb84ef6aab28544051.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6cf98b0ba15623a64ec587ad6a80d5690b2a34c0ba7e0016c5f13c28d749dfc1.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i6cf98b0ba15623a64ec587ad6a80d5690b2a34c0ba7e0016c5f13c28d749dfc1.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i32ed368ce5ee12bfd01e4b23d9aa3f8f0a53f6b58ebcba5ff956c739eae37b41.SingleValueExtendedPropertiesRequestBuilder) {
    return i32ed368ce5ee12bfd01e4b23d9aa3f8f0a53f6b58ebcba5ff956c739eae37b41.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id8a23edf2d82fd94fc95e02eb04d228bf11f41aa3452bbd39108f71fb7b96f2c.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return id8a23edf2d82fd94fc95e02eb04d228bf11f41aa3452bbd39108f71fb7b96f2c.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i6d0fc705adb0d375d549d85597118478eea44ed85a34e1950b70ca1d40ee2263.SnoozeReminderRequestBuilder) {
    return i6d0fc705adb0d375d549d85597118478eea44ed85a34e1950b70ca1d40ee2263.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i5a22110dcba479e1eff1b0c29b99c5190def78fa9a90e3dd939e20db4a0d09a9.TentativelyAcceptRequestBuilder) {
    return i5a22110dcba479e1eff1b0c29b99c5190def78fa9a90e3dd939e20db4a0d09a9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
