package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i059dab5e0aa78e83cfd615dda918a2c44d52daaa45372cb50f4bdcb514105cde "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/dismissreminder"
    i0d5d4fce97db38d2a5b93fbc1501a96a95c48e6dbb0a1641f9d615536e8fdb85 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/decline"
    i1201defebd7ce1e7d618f96de982cfb8aed706c05f9bfdecd38c1058138330fa "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/attachments"
    i14c2916673848ffbd3aceff4d7b49e2dc70fa60d7e7b789282b202b5b22bca0b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/multivalueextendedproperties"
    i2fa8f54576ab0073b0e114c1abe63c305fdf346eb1ba9dad9a8229c2238b7a4f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances"
    i3cbdd98e0afa2ffda5ae80ddc70ba716047bcfe5b400f6c3e7809c3d891b03ad "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/snoozereminder"
    i76c4305997d6cb13bf9da13aa31e7e7b4ea8c55cdf2ea79d14c2fdeeb448c185 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences"
    i770a5145fd01ee602a18c426ac71a2cdead8830b24f6ad3719436886bcc9d7c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/tentativelyaccept"
    i9128f3c60135f22c7a47ea9daa2314158aa7788abd19f1fbca2ad975760a7e07 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/extensions"
    ia44285a0d899332adefbad5c8c67cf547debbcaf746514cf5302c11d36802773 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/singlevalueextendedproperties"
    ic52a4b74b1cf2455bad81a20dc724df4cd71fe40313d94035c5b81bf66ea5133 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/calendar"
    ie81dfddf719cafd2afba4c80536ed434eec18f9c2171a25468bb1d24ba3bedeb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/forward"
    iefcbcbbdae997919fc5de05ed7114f2a54e10b4e2ac8918402b4edff67f30424 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/accept"
    if3156c0c94138e02899ab847edde1d4e045eb5a3b668b3dede6f5559615c6a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/cancel"
    i251d233f5532efc8f12bbf88b9ca31fa771e05cdc0f1ae5875f7add75a4c74c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/extensions/item"
    i2ebc6aacd0f573f384bf8d2d129513cff7da5807dab9633bd6908d1883c8ee6a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item"
    i47b8ece28c607d20c15c3a51208eede6e26c773b4fa6fbb653ec711cdf217df4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/attachments/item"
    i9855b087292f7172af644d450f0c55e1127a8a59839510cce0e9db7dae78568c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item"
    icf5853047b0325830c18b342a5e684e64a614c8eb80b63cdba2cfe39a25a6ebf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/singlevalueextendedproperties/item"
    ida403eba8e8430247eef196fbc4ec9d9c0bd649127d0bf203ade9a926cf3f55c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/multivalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \groups\{group-id}\calendar\events\{event-id}
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
func (m *EventRequestBuilder) Accept()(*iefcbcbbdae997919fc5de05ed7114f2a54e10b4e2ac8918402b4edff67f30424.AcceptRequestBuilder) {
    return iefcbcbbdae997919fc5de05ed7114f2a54e10b4e2ac8918402b4edff67f30424.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i1201defebd7ce1e7d618f96de982cfb8aed706c05f9bfdecd38c1058138330fa.AttachmentsRequestBuilder) {
    return i1201defebd7ce1e7d618f96de982cfb8aed706c05f9bfdecd38c1058138330fa.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*i47b8ece28c607d20c15c3a51208eede6e26c773b4fa6fbb653ec711cdf217df4.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i47b8ece28c607d20c15c3a51208eede6e26c773b4fa6fbb653ec711cdf217df4.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ic52a4b74b1cf2455bad81a20dc724df4cd71fe40313d94035c5b81bf66ea5133.CalendarRequestBuilder) {
    return ic52a4b74b1cf2455bad81a20dc724df4cd71fe40313d94035c5b81bf66ea5133.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*if3156c0c94138e02899ab847edde1d4e045eb5a3b668b3dede6f5559615c6a91.CancelRequestBuilder) {
    return if3156c0c94138e02899ab847edde1d4e045eb5a3b668b3dede6f5559615c6a91.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/events/{event_id}{?select}";
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
func (m *EventRequestBuilder) Decline()(*i0d5d4fce97db38d2a5b93fbc1501a96a95c48e6dbb0a1641f9d615536e8fdb85.DeclineRequestBuilder) {
    return i0d5d4fce97db38d2a5b93fbc1501a96a95c48e6dbb0a1641f9d615536e8fdb85.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i059dab5e0aa78e83cfd615dda918a2c44d52daaa45372cb50f4bdcb514105cde.DismissReminderRequestBuilder) {
    return i059dab5e0aa78e83cfd615dda918a2c44d52daaa45372cb50f4bdcb514105cde.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i76c4305997d6cb13bf9da13aa31e7e7b4ea8c55cdf2ea79d14c2fdeeb448c185.ExceptionOccurrencesRequestBuilder) {
    return i76c4305997d6cb13bf9da13aa31e7e7b4ea8c55cdf2ea79d14c2fdeeb448c185.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i2ebc6aacd0f573f384bf8d2d129513cff7da5807dab9633bd6908d1883c8ee6a.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i2ebc6aacd0f573f384bf8d2d129513cff7da5807dab9633bd6908d1883c8ee6a.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i9128f3c60135f22c7a47ea9daa2314158aa7788abd19f1fbca2ad975760a7e07.ExtensionsRequestBuilder) {
    return i9128f3c60135f22c7a47ea9daa2314158aa7788abd19f1fbca2ad975760a7e07.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i251d233f5532efc8f12bbf88b9ca31fa771e05cdc0f1ae5875f7add75a4c74c8.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i251d233f5532efc8f12bbf88b9ca31fa771e05cdc0f1ae5875f7add75a4c74c8.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ie81dfddf719cafd2afba4c80536ed434eec18f9c2171a25468bb1d24ba3bedeb.ForwardRequestBuilder) {
    return ie81dfddf719cafd2afba4c80536ed434eec18f9c2171a25468bb1d24ba3bedeb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*i2fa8f54576ab0073b0e114c1abe63c305fdf346eb1ba9dad9a8229c2238b7a4f.InstancesRequestBuilder) {
    return i2fa8f54576ab0073b0e114c1abe63c305fdf346eb1ba9dad9a8229c2238b7a4f.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i9855b087292f7172af644d450f0c55e1127a8a59839510cce0e9db7dae78568c.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i9855b087292f7172af644d450f0c55e1127a8a59839510cce0e9db7dae78568c.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i14c2916673848ffbd3aceff4d7b49e2dc70fa60d7e7b789282b202b5b22bca0b.MultiValueExtendedPropertiesRequestBuilder) {
    return i14c2916673848ffbd3aceff4d7b49e2dc70fa60d7e7b789282b202b5b22bca0b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ida403eba8e8430247eef196fbc4ec9d9c0bd649127d0bf203ade9a926cf3f55c.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ida403eba8e8430247eef196fbc4ec9d9c0bd649127d0bf203ade9a926cf3f55c.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ia44285a0d899332adefbad5c8c67cf547debbcaf746514cf5302c11d36802773.SingleValueExtendedPropertiesRequestBuilder) {
    return ia44285a0d899332adefbad5c8c67cf547debbcaf746514cf5302c11d36802773.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*icf5853047b0325830c18b342a5e684e64a614c8eb80b63cdba2cfe39a25a6ebf.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return icf5853047b0325830c18b342a5e684e64a614c8eb80b63cdba2cfe39a25a6ebf.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i3cbdd98e0afa2ffda5ae80ddc70ba716047bcfe5b400f6c3e7809c3d891b03ad.SnoozeReminderRequestBuilder) {
    return i3cbdd98e0afa2ffda5ae80ddc70ba716047bcfe5b400f6c3e7809c3d891b03ad.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i770a5145fd01ee602a18c426ac71a2cdead8830b24f6ad3719436886bcc9d7c9.TentativelyAcceptRequestBuilder) {
    return i770a5145fd01ee602a18c426ac71a2cdead8830b24f6ad3719436886bcc9d7c9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
