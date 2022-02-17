package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1d7efb0cf0e538f716cbf4ec58ffbd8b2627f35a80b5cc51ad1a719ccda63410 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/tentativelyaccept"
    i594f031e77c7ff1deebf416d7fe1c8b3e48a3898deb3d09b0a51f16b899b9999 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/extensions"
    i5b406bcd2ab417fc4dd03bec5345f447db84b34a3739079d6b77141d501cad42 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/dismissreminder"
    i71a1dfd1790ad38215efe86cb847a7e9aa5b0e4485774ae264b132be08548500 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/snoozereminder"
    i761da5c466c608a108c783f66ae6a90778938de7cac958956283379dfa5f5b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/decline"
    i79f6e8eea7c7b3e6a2bce802b9a7f0a342ccafa8bb76d03da9284afa5af979ba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/attachments"
    i84e35fc82945374b0b4e6d6b60b9a283b295d752e621fa432c44c539b7457a13 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances"
    i95be0951230a29aac635ed6fe9a25887c5bbd1fa02bf4ea7df08ee54d2b119c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/calendar"
    i96f018cc33130e2eb44ba1f382328043fb9aa8cf4cbb32e98de1271009bcb638 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/cancel"
    ia0356016284b401708cc71750c01506b894366ee232c9f3a246c7a5f66855d00 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/accept"
    iba617898bd91d23a58685d21f4d6bd3cc9e930cf5709a670c8cbf8d09826572c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties"
    ic2c0e2ed3b1c4a6edf0119f14ebfd9619d42bb1ec48bfc2d02573e9892f75d7f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences"
    idb15d9c63943dfdd77189c4bc501947f713932b0b8a26a41ec57a1a9ed91bd3f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/forward"
    if32e2baabff5c88751154ffb2c522378538c97134301e17d7a7a463932a40dc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    i18e39b2135201cee257b03e8127c40e270b8ac9c6928d3b1408ca41554f5736f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item"
    i223f69849eb07c52dccbfbd8ef59d3b734f676c9d212707b26943e8a446dea29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item"
    i39da18ddc9fe463d0b5beda214a3ca0a1bbdc9a87511170acbce67edc1dd7aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/extensions/item"
    i7de345c43b4c56542c252f9c3b2494b65cf9d8a173f7713ee3bdee0b5aa8f6e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    ide91c535cba1a01dae33b9ba11703c6da99e6cb97eefde9dbed88118af7df3a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
    ie52936233179f5ad76614301cc056362325b5238f0fd7c3078dcce062a6ebadd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/attachments/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\calendarView\{event-id}
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
func (m *EventRequestBuilder) Accept()(*ia0356016284b401708cc71750c01506b894366ee232c9f3a246c7a5f66855d00.AcceptRequestBuilder) {
    return ia0356016284b401708cc71750c01506b894366ee232c9f3a246c7a5f66855d00.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i79f6e8eea7c7b3e6a2bce802b9a7f0a342ccafa8bb76d03da9284afa5af979ba.AttachmentsRequestBuilder) {
    return i79f6e8eea7c7b3e6a2bce802b9a7f0a342ccafa8bb76d03da9284afa5af979ba.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ie52936233179f5ad76614301cc056362325b5238f0fd7c3078dcce062a6ebadd.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ie52936233179f5ad76614301cc056362325b5238f0fd7c3078dcce062a6ebadd.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i95be0951230a29aac635ed6fe9a25887c5bbd1fa02bf4ea7df08ee54d2b119c0.CalendarRequestBuilder) {
    return i95be0951230a29aac635ed6fe9a25887c5bbd1fa02bf4ea7df08ee54d2b119c0.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i96f018cc33130e2eb44ba1f382328043fb9aa8cf4cbb32e98de1271009bcb638.CancelRequestBuilder) {
    return i96f018cc33130e2eb44ba1f382328043fb9aa8cf4cbb32e98de1271009bcb638.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}{?select}";
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
func (m *EventRequestBuilder) Decline()(*i761da5c466c608a108c783f66ae6a90778938de7cac958956283379dfa5f5b3c.DeclineRequestBuilder) {
    return i761da5c466c608a108c783f66ae6a90778938de7cac958956283379dfa5f5b3c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i5b406bcd2ab417fc4dd03bec5345f447db84b34a3739079d6b77141d501cad42.DismissReminderRequestBuilder) {
    return i5b406bcd2ab417fc4dd03bec5345f447db84b34a3739079d6b77141d501cad42.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*ic2c0e2ed3b1c4a6edf0119f14ebfd9619d42bb1ec48bfc2d02573e9892f75d7f.ExceptionOccurrencesRequestBuilder) {
    return ic2c0e2ed3b1c4a6edf0119f14ebfd9619d42bb1ec48bfc2d02573e9892f75d7f.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i18e39b2135201cee257b03e8127c40e270b8ac9c6928d3b1408ca41554f5736f.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i18e39b2135201cee257b03e8127c40e270b8ac9c6928d3b1408ca41554f5736f.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i594f031e77c7ff1deebf416d7fe1c8b3e48a3898deb3d09b0a51f16b899b9999.ExtensionsRequestBuilder) {
    return i594f031e77c7ff1deebf416d7fe1c8b3e48a3898deb3d09b0a51f16b899b9999.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i39da18ddc9fe463d0b5beda214a3ca0a1bbdc9a87511170acbce67edc1dd7aa3.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i39da18ddc9fe463d0b5beda214a3ca0a1bbdc9a87511170acbce67edc1dd7aa3.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*idb15d9c63943dfdd77189c4bc501947f713932b0b8a26a41ec57a1a9ed91bd3f.ForwardRequestBuilder) {
    return idb15d9c63943dfdd77189c4bc501947f713932b0b8a26a41ec57a1a9ed91bd3f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*i84e35fc82945374b0b4e6d6b60b9a283b295d752e621fa432c44c539b7457a13.InstancesRequestBuilder) {
    return i84e35fc82945374b0b4e6d6b60b9a283b295d752e621fa432c44c539b7457a13.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i223f69849eb07c52dccbfbd8ef59d3b734f676c9d212707b26943e8a446dea29.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i223f69849eb07c52dccbfbd8ef59d3b734f676c9d212707b26943e8a446dea29.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*iba617898bd91d23a58685d21f4d6bd3cc9e930cf5709a670c8cbf8d09826572c.MultiValueExtendedPropertiesRequestBuilder) {
    return iba617898bd91d23a58685d21f4d6bd3cc9e930cf5709a670c8cbf8d09826572c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7de345c43b4c56542c252f9c3b2494b65cf9d8a173f7713ee3bdee0b5aa8f6e1.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i7de345c43b4c56542c252f9c3b2494b65cf9d8a173f7713ee3bdee0b5aa8f6e1.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*if32e2baabff5c88751154ffb2c522378538c97134301e17d7a7a463932a40dc0.SingleValueExtendedPropertiesRequestBuilder) {
    return if32e2baabff5c88751154ffb2c522378538c97134301e17d7a7a463932a40dc0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ide91c535cba1a01dae33b9ba11703c6da99e6cb97eefde9dbed88118af7df3a8.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ide91c535cba1a01dae33b9ba11703c6da99e6cb97eefde9dbed88118af7df3a8.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i71a1dfd1790ad38215efe86cb847a7e9aa5b0e4485774ae264b132be08548500.SnoozeReminderRequestBuilder) {
    return i71a1dfd1790ad38215efe86cb847a7e9aa5b0e4485774ae264b132be08548500.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i1d7efb0cf0e538f716cbf4ec58ffbd8b2627f35a80b5cc51ad1a719ccda63410.TentativelyAcceptRequestBuilder) {
    return i1d7efb0cf0e538f716cbf4ec58ffbd8b2627f35a80b5cc51ad1a719ccda63410.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
