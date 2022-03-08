package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0b59187fa51909a371f8b9858abba00a04491a0a4774c16f37ed7ad230dd3c80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/multivalueextendedproperties"
    i181320aa93ef4321014b74609a44e3f5a32133f9580d21d62327fa7bbb5cda54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/snoozereminder"
    i19dd778a11375dddbb2d2611e805800876e81cba206c523af607ec0876d42638 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/dismissreminder"
    i336acfa1272f48c2fe797109baa0a33d28eb0913b952cd433351e0dca2bfea27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/calendar"
    i5793f33136eb9238f8d107847a380f870d05d5e35af7d529b95b28bce2609cf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/tentativelyaccept"
    i60fcaa43c34fa0b127246ffcfd3c522b1bb552c2093011690608d937556cadcb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/cancel"
    i83b5cc71f24decf0fa568a1fea8c63f1115a6c0cda7a6f7a48cb1cc942e4eb75 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/attachments"
    i8553f6cdfc8767ac4a225c751996e9393fc28a393eea6cf2b2cb3fb7537c37d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/singlevalueextendedproperties"
    ia64b443dfec1fdf0f96b97e8ce3dc6342bfb217cb244a28fb51f4e2fac80098a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/extensions"
    ib22faf4889a4ff940189a07d0634017ab1700b96946ffe4c7859b1c600c95b13 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/accept"
    ic699aa89cf0fa5dce893a4e798799256b1490ee3a0e7945a5380b385a6e1ebf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/forward"
    id81837152e19e46f34e39a5f29933b7bf6301a3b447bcdab80bfcac087daa8df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/decline"
    ieb69cb4984db33666e55e73bcf1987aa77b5daea0acac2114d5c716cc8ee9590 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances"
    if3620ac1eeab8b06c743846feb104c5b5c1c860caa9b835409096fbd72f4dbc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences"
    i08d2438ecad8b31d14bf1588846d430634768d2d73ce954a9c7459f7128dee29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item"
    i0e2c2c7317d5772b90fb64d735c3dd692b4af05af875d54fd676524e724cd512 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item"
    i2c60bfae6ea65be12b4d02f34ad3d2bfe71fdc157260307ad234bfb2acc44811 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/multivalueextendedproperties/item"
    i5f7bc65efefc422b0d02ea65c0c11a8e28f4805da2f6d82655cafcb3c05ae1a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/singlevalueextendedproperties/item"
    ia9e0f6249aba2f425cdb963bff747c708eaa818659d217e6df3cd92448321a0f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/attachments/item"
    ic6abe1dc87b2423ac4d607f9d80cfadb179e722bf3ebc7a79351f62c697aff2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
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
// EventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*ib22faf4889a4ff940189a07d0634017ab1700b96946ffe4c7859b1c600c95b13.AcceptRequestBuilder) {
    return ib22faf4889a4ff940189a07d0634017ab1700b96946ffe4c7859b1c600c95b13.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i83b5cc71f24decf0fa568a1fea8c63f1115a6c0cda7a6f7a48cb1cc942e4eb75.AttachmentsRequestBuilder) {
    return i83b5cc71f24decf0fa568a1fea8c63f1115a6c0cda7a6f7a48cb1cc942e4eb75.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia9e0f6249aba2f425cdb963bff747c708eaa818659d217e6df3cd92448321a0f.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia9e0f6249aba2f425cdb963bff747c708eaa818659d217e6df3cd92448321a0f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i336acfa1272f48c2fe797109baa0a33d28eb0913b952cd433351e0dca2bfea27.CalendarRequestBuilder) {
    return i336acfa1272f48c2fe797109baa0a33d28eb0913b952cd433351e0dca2bfea27.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i60fcaa43c34fa0b127246ffcfd3c522b1bb552c2093011690608d937556cadcb.CancelRequestBuilder) {
    return i60fcaa43c34fa0b127246ffcfd3c522b1bb552c2093011690608d937556cadcb.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property events for users
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation update the navigation property events in users
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
func (m *EventItemRequestBuilder) Decline()(*id81837152e19e46f34e39a5f29933b7bf6301a3b447bcdab80bfcac087daa8df.DeclineRequestBuilder) {
    return id81837152e19e46f34e39a5f29933b7bf6301a3b447bcdab80bfcac087daa8df.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i19dd778a11375dddbb2d2611e805800876e81cba206c523af607ec0876d42638.DismissReminderRequestBuilder) {
    return i19dd778a11375dddbb2d2611e805800876e81cba206c523af607ec0876d42638.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*if3620ac1eeab8b06c743846feb104c5b5c1c860caa9b835409096fbd72f4dbc2.ExceptionOccurrencesRequestBuilder) {
    return if3620ac1eeab8b06c743846feb104c5b5c1c860caa9b835409096fbd72f4dbc2.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i0e2c2c7317d5772b90fb64d735c3dd692b4af05af875d54fd676524e724cd512.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0e2c2c7317d5772b90fb64d735c3dd692b4af05af875d54fd676524e724cd512.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*ia64b443dfec1fdf0f96b97e8ce3dc6342bfb217cb244a28fb51f4e2fac80098a.ExtensionsRequestBuilder) {
    return ia64b443dfec1fdf0f96b97e8ce3dc6342bfb217cb244a28fb51f4e2fac80098a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic6abe1dc87b2423ac4d607f9d80cfadb179e722bf3ebc7a79351f62c697aff2d.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic6abe1dc87b2423ac4d607f9d80cfadb179e722bf3ebc7a79351f62c697aff2d.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ic699aa89cf0fa5dce893a4e798799256b1490ee3a0e7945a5380b385a6e1ebf4.ForwardRequestBuilder) {
    return ic699aa89cf0fa5dce893a4e798799256b1490ee3a0e7945a5380b385a6e1ebf4.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*ieb69cb4984db33666e55e73bcf1987aa77b5daea0acac2114d5c716cc8ee9590.InstancesRequestBuilder) {
    return ieb69cb4984db33666e55e73bcf1987aa77b5daea0acac2114d5c716cc8ee9590.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i08d2438ecad8b31d14bf1588846d430634768d2d73ce954a9c7459f7128dee29.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i08d2438ecad8b31d14bf1588846d430634768d2d73ce954a9c7459f7128dee29.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i0b59187fa51909a371f8b9858abba00a04491a0a4774c16f37ed7ad230dd3c80.MultiValueExtendedPropertiesRequestBuilder) {
    return i0b59187fa51909a371f8b9858abba00a04491a0a4774c16f37ed7ad230dd3c80.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2c60bfae6ea65be12b4d02f34ad3d2bfe71fdc157260307ad234bfb2acc44811.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2c60bfae6ea65be12b4d02f34ad3d2bfe71fdc157260307ad234bfb2acc44811.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i8553f6cdfc8767ac4a225c751996e9393fc28a393eea6cf2b2cb3fb7537c37d6.SingleValueExtendedPropertiesRequestBuilder) {
    return i8553f6cdfc8767ac4a225c751996e9393fc28a393eea6cf2b2cb3fb7537c37d6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5f7bc65efefc422b0d02ea65c0c11a8e28f4805da2f6d82655cafcb3c05ae1a7.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5f7bc65efefc422b0d02ea65c0c11a8e28f4805da2f6d82655cafcb3c05ae1a7.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i181320aa93ef4321014b74609a44e3f5a32133f9580d21d62327fa7bbb5cda54.SnoozeReminderRequestBuilder) {
    return i181320aa93ef4321014b74609a44e3f5a32133f9580d21d62327fa7bbb5cda54.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i5793f33136eb9238f8d107847a380f870d05d5e35af7d529b95b28bce2609cf5.TentativelyAcceptRequestBuilder) {
    return i5793f33136eb9238f8d107847a380f870d05d5e35af7d529b95b28bce2609cf5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
