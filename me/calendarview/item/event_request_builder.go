package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/extensions"
    i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/multivalueextendedproperties"
    i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances"
    i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/attachments"
    i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/accept"
    i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/dismissreminder"
    i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences"
    i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/tentativelyaccept"
    ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/snoozereminder"
    ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/singlevalueextendedproperties"
    iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/cancel"
    ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/decline"
    iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/forward"
    if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar"
    i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/singlevalueextendedproperties/item"
    i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item"
    ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/extensions/item"
    ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/multivalueextendedproperties/item"
    if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/attachments/item"
    if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendarView\{event-id}
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
// EventRequestBuilderGetQueryParameters the calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) Accept()(*i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602.AcceptRequestBuilder) {
    return i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9.AttachmentsRequestBuilder) {
    return i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af.CalendarRequestBuilder) {
    return if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b.CancelRequestBuilder) {
    return iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// CreateDeleteRequestInformation the calendar view for the calendar. Read-only. Nullable.
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
// CreateGetRequestInformation the calendar view for the calendar. Read-only. Nullable.
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
// CreatePatchRequestInformation the calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) Decline()(*ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6.DeclineRequestBuilder) {
    return ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) DismissReminder()(*i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34.DismissReminderRequestBuilder) {
    return i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5.ExceptionOccurrencesRequestBuilder) {
    return i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f.ExtensionsRequestBuilder) {
    return i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95.ForwardRequestBuilder) {
    return iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) Instances()(*i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954.InstancesRequestBuilder) {
    return i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f.MultiValueExtendedPropertiesRequestBuilder) {
    return i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar view for the calendar. Read-only. Nullable.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476.SingleValueExtendedPropertiesRequestBuilder) {
    return ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4.SnoozeReminderRequestBuilder) {
    return ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63.TentativelyAcceptRequestBuilder) {
    return i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
