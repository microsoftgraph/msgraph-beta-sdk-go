package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i019db5d653a643e1a549e917348bbdacf1efb3a8caf89ca71a40988c2d3ca692 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i36bb761cf2b3fd08e22dd0629b6d3414537d80a143da72167abdc43db9fe5176 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i40b295700d1a3dc9d4856ca82158a88dc5d2d35c5958521997664ba4fb1148e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i51fe48770386ecc7a9855817ee2ac09be025a6a9b143895262f400e1b3714012 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i55d48ec4278f3950b9d37bd21c26e0c1f021183da3015589d261582b9ce9978f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i7a291b26188996d4cdd5a95ee9429a4386af4d615fa05f98cba8442103f1969f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/accept"
    i824af3ce9c56f248451ff666dab98ed0a74003056e040dbb1bf5e36d52a9a78b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i8992ab597ecd5e75d6fac72ebb8df63711765f795e1150140c8ca3b713450fc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    i8e6836cb48c3fcffb55577d809f7a3db58227602f22c9c5d5145099bc8b03b20 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    ib97d7ab381f4d3986575b3f368addd550b65ae2dd39dff078276ced80a3047ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/forward"
    ibc8dd027671410060ae3e28393b66723e6b3fa7d64029404faa96df20501488f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/decline"
    ie5295b5cffb686191c2e1152e4dfe27d2a49572591853453888b9e997af70e80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i56325cc99232a0ede8e79a8e95284320babc1e40ece138d8773f76f51025cf3b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    i66a83b5642595f7c324d4fae0b85ed53b98f278d9fbd18fc37cbcc3484190981 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    id8c880825b71c3b54ed5b24f987ef412bf93ec0d06b3b9ac1cf44402719b34aa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    if75358db76287af9368cb02f46e1d783229c00f6337c1d11b9013fb349cf584c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Accept()(*i7a291b26188996d4cdd5a95ee9429a4386af4d615fa05f98cba8442103f1969f.AcceptRequestBuilder) {
    return i7a291b26188996d4cdd5a95ee9429a4386af4d615fa05f98cba8442103f1969f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ie5295b5cffb686191c2e1152e4dfe27d2a49572591853453888b9e997af70e80.AttachmentsRequestBuilder) {
    return ie5295b5cffb686191c2e1152e4dfe27d2a49572591853453888b9e997af70e80.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*id8c880825b71c3b54ed5b24f987ef412bf93ec0d06b3b9ac1cf44402719b34aa.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return id8c880825b71c3b54ed5b24f987ef412bf93ec0d06b3b9ac1cf44402719b34aa.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i8992ab597ecd5e75d6fac72ebb8df63711765f795e1150140c8ca3b713450fc2.CalendarRequestBuilder) {
    return i8992ab597ecd5e75d6fac72ebb8df63711765f795e1150140c8ca3b713450fc2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i019db5d653a643e1a549e917348bbdacf1efb3a8caf89ca71a40988c2d3ca692.CancelRequestBuilder) {
    return i019db5d653a643e1a549e917348bbdacf1efb3a8caf89ca71a40988c2d3ca692.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/events/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
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
// CreateDeleteRequestInformation delete navigation property instances for users
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property instances in users
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
func (m *EventItemRequestBuilder) Decline()(*ibc8dd027671410060ae3e28393b66723e6b3fa7d64029404faa96df20501488f.DeclineRequestBuilder) {
    return ibc8dd027671410060ae3e28393b66723e6b3fa7d64029404faa96df20501488f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i51fe48770386ecc7a9855817ee2ac09be025a6a9b143895262f400e1b3714012.DismissReminderRequestBuilder) {
    return i51fe48770386ecc7a9855817ee2ac09be025a6a9b143895262f400e1b3714012.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i8e6836cb48c3fcffb55577d809f7a3db58227602f22c9c5d5145099bc8b03b20.ExtensionsRequestBuilder) {
    return i8e6836cb48c3fcffb55577d809f7a3db58227602f22c9c5d5145099bc8b03b20.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*if75358db76287af9368cb02f46e1d783229c00f6337c1d11b9013fb349cf584c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return if75358db76287af9368cb02f46e1d783229c00f6337c1d11b9013fb349cf584c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ib97d7ab381f4d3986575b3f368addd550b65ae2dd39dff078276ced80a3047ac.ForwardRequestBuilder) {
    return ib97d7ab381f4d3986575b3f368addd550b65ae2dd39dff078276ced80a3047ac.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i40b295700d1a3dc9d4856ca82158a88dc5d2d35c5958521997664ba4fb1148e2.MultiValueExtendedPropertiesRequestBuilder) {
    return i40b295700d1a3dc9d4856ca82158a88dc5d2d35c5958521997664ba4fb1148e2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i56325cc99232a0ede8e79a8e95284320babc1e40ece138d8773f76f51025cf3b.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i56325cc99232a0ede8e79a8e95284320babc1e40ece138d8773f76f51025cf3b.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i824af3ce9c56f248451ff666dab98ed0a74003056e040dbb1bf5e36d52a9a78b.SingleValueExtendedPropertiesRequestBuilder) {
    return i824af3ce9c56f248451ff666dab98ed0a74003056e040dbb1bf5e36d52a9a78b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i66a83b5642595f7c324d4fae0b85ed53b98f278d9fbd18fc37cbcc3484190981.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i66a83b5642595f7c324d4fae0b85ed53b98f278d9fbd18fc37cbcc3484190981.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i55d48ec4278f3950b9d37bd21c26e0c1f021183da3015589d261582b9ce9978f.SnoozeReminderRequestBuilder) {
    return i55d48ec4278f3950b9d37bd21c26e0c1f021183da3015589d261582b9ce9978f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i36bb761cf2b3fd08e22dd0629b6d3414537d80a143da72167abdc43db9fe5176.TentativelyAcceptRequestBuilder) {
    return i36bb761cf2b3fd08e22dd0629b6d3414537d80a143da72167abdc43db9fe5176.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
