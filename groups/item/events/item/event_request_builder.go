package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i05a42ea7d5e1884a8e1dae2793ab110b1e1f8d988db71891436832059ff2ba3c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/cancel"
    i5fb746eca0c242c8984c8a879184ad9d08b364aa7b9c285a3bc182d2ddfc503e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/attachments"
    i79a7405307cbbdb21bedf1701ef943fd18c4517f47986e866d5dd325fbd2c118 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/decline"
    i8aae6ec506cc62628c5af6e10369ee326d5030e6a9f6dd3af2f33eed318509cb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/accept"
    iae63acc96b63dde476df5d572aa9d992f16b648b4365e066f9f52d4d87bb923c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/tentativelyaccept"
    ib4e29bfd4281e36f35dc582f9c214a6f560296bfeb558b3a54ea7b8bf8ae5b8e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/extensions"
    ib9a7b515d2c33bd771c3637dd035f38dd054dfc1c3bedb54b7c1c3fc35bd0c4f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/snoozereminder"
    ibc5ca5cafb1fc0ba72521ab6774586b362f987eb4515dfa602dbad582436708c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/multivalueextendedproperties"
    ic0681472e18a3022a4b2b729ab4483580fb7b4e6b7ae1d4f181ec9dbd1da88d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences"
    id83297e5327484cd1ced89db57bf0a85fa33d5cae6cdbb44636d29ebe5ae6f36 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances"
    iea920758cfc75c1ccf77797bcdbc550af58e71c9f109ec8a94722a33d74c7cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/singlevalueextendedproperties"
    ieea5b0e4f6283da8e5578128baf2e8078539a018e85731134404d8b5aa938c76 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar"
    if96aa91366f276971bcdb05b3ac53fb867427708b271afba983647fc1f3609cd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/forward"
    if9e65d705bb4056db822cab18c1d189374708a5ba4fbdacd32700342862d9952 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/dismissreminder"
    i1e6b45e04eb5c15511813d0040d37a7573d659ce2474195041adf65e7a3da179 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item"
    i51135f6535d37302d8527c79893cfa3567ae1a2cc4fec29b051f79c3f3b706b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/extensions/item"
    i792a485ff185e2fdffa2e1feaabbe0871f5a892a62e0d8d85b4f2a593554431a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item"
    i81aa3005b6422532d1fe02077a5f622ae66cc374dced98aa4b6578b09fbfba01 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/singlevalueextendedproperties/item"
    iaf6ee02f9752677d650b6f2d4444fe3664c88db77ed2d14f02597638d09e85d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/multivalueextendedproperties/item"
    ife7ef1d1a994cdb0804ec1d15de22eaf8f8e9bc910b45c55f19b4d13e63ecff9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/attachments/item"
)

// EventRequestBuilder builds and executes requests for operations under \groups\{group-id}\events\{event-id}
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
// EventRequestBuilderGetQueryParameters the group's calendar events.
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
func (m *EventRequestBuilder) Accept()(*i8aae6ec506cc62628c5af6e10369ee326d5030e6a9f6dd3af2f33eed318509cb.AcceptRequestBuilder) {
    return i8aae6ec506cc62628c5af6e10369ee326d5030e6a9f6dd3af2f33eed318509cb.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i5fb746eca0c242c8984c8a879184ad9d08b364aa7b9c285a3bc182d2ddfc503e.AttachmentsRequestBuilder) {
    return i5fb746eca0c242c8984c8a879184ad9d08b364aa7b9c285a3bc182d2ddfc503e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ife7ef1d1a994cdb0804ec1d15de22eaf8f8e9bc910b45c55f19b4d13e63ecff9.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ife7ef1d1a994cdb0804ec1d15de22eaf8f8e9bc910b45c55f19b4d13e63ecff9.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ieea5b0e4f6283da8e5578128baf2e8078539a018e85731134404d8b5aa938c76.CalendarRequestBuilder) {
    return ieea5b0e4f6283da8e5578128baf2e8078539a018e85731134404d8b5aa938c76.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i05a42ea7d5e1884a8e1dae2793ab110b1e1f8d988db71891436832059ff2ba3c.CancelRequestBuilder) {
    return i05a42ea7d5e1884a8e1dae2793ab110b1e1f8d988db71891436832059ff2ba3c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation the group's calendar events.
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
// CreateGetRequestInformation the group's calendar events.
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
// CreatePatchRequestInformation the group's calendar events.
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
func (m *EventRequestBuilder) Decline()(*i79a7405307cbbdb21bedf1701ef943fd18c4517f47986e866d5dd325fbd2c118.DeclineRequestBuilder) {
    return i79a7405307cbbdb21bedf1701ef943fd18c4517f47986e866d5dd325fbd2c118.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the group's calendar events.
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
func (m *EventRequestBuilder) DismissReminder()(*if9e65d705bb4056db822cab18c1d189374708a5ba4fbdacd32700342862d9952.DismissReminderRequestBuilder) {
    return if9e65d705bb4056db822cab18c1d189374708a5ba4fbdacd32700342862d9952.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*ic0681472e18a3022a4b2b729ab4483580fb7b4e6b7ae1d4f181ec9dbd1da88d9.ExceptionOccurrencesRequestBuilder) {
    return ic0681472e18a3022a4b2b729ab4483580fb7b4e6b7ae1d4f181ec9dbd1da88d9.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i1e6b45e04eb5c15511813d0040d37a7573d659ce2474195041adf65e7a3da179.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i1e6b45e04eb5c15511813d0040d37a7573d659ce2474195041adf65e7a3da179.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ib4e29bfd4281e36f35dc582f9c214a6f560296bfeb558b3a54ea7b8bf8ae5b8e.ExtensionsRequestBuilder) {
    return ib4e29bfd4281e36f35dc582f9c214a6f560296bfeb558b3a54ea7b8bf8ae5b8e.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i51135f6535d37302d8527c79893cfa3567ae1a2cc4fec29b051f79c3f3b706b1.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i51135f6535d37302d8527c79893cfa3567ae1a2cc4fec29b051f79c3f3b706b1.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*if96aa91366f276971bcdb05b3ac53fb867427708b271afba983647fc1f3609cd.ForwardRequestBuilder) {
    return if96aa91366f276971bcdb05b3ac53fb867427708b271afba983647fc1f3609cd.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the group's calendar events.
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
func (m *EventRequestBuilder) Instances()(*id83297e5327484cd1ced89db57bf0a85fa33d5cae6cdbb44636d29ebe5ae6f36.InstancesRequestBuilder) {
    return id83297e5327484cd1ced89db57bf0a85fa33d5cae6cdbb44636d29ebe5ae6f36.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*i792a485ff185e2fdffa2e1feaabbe0871f5a892a62e0d8d85b4f2a593554431a.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i792a485ff185e2fdffa2e1feaabbe0871f5a892a62e0d8d85b4f2a593554431a.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*ibc5ca5cafb1fc0ba72521ab6774586b362f987eb4515dfa602dbad582436708c.MultiValueExtendedPropertiesRequestBuilder) {
    return ibc5ca5cafb1fc0ba72521ab6774586b362f987eb4515dfa602dbad582436708c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iaf6ee02f9752677d650b6f2d4444fe3664c88db77ed2d14f02597638d09e85d1.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iaf6ee02f9752677d650b6f2d4444fe3664c88db77ed2d14f02597638d09e85d1.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the group's calendar events.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*iea920758cfc75c1ccf77797bcdbc550af58e71c9f109ec8a94722a33d74c7cc4.SingleValueExtendedPropertiesRequestBuilder) {
    return iea920758cfc75c1ccf77797bcdbc550af58e71c9f109ec8a94722a33d74c7cc4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i81aa3005b6422532d1fe02077a5f622ae66cc374dced98aa4b6578b09fbfba01.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i81aa3005b6422532d1fe02077a5f622ae66cc374dced98aa4b6578b09fbfba01.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ib9a7b515d2c33bd771c3637dd035f38dd054dfc1c3bedb54b7c1c3fc35bd0c4f.SnoozeReminderRequestBuilder) {
    return ib9a7b515d2c33bd771c3637dd035f38dd054dfc1c3bedb54b7c1c3fc35bd0c4f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*iae63acc96b63dde476df5d572aa9d992f16b648b4365e066f9f52d4d87bb923c.TentativelyAcceptRequestBuilder) {
    return iae63acc96b63dde476df5d572aa9d992f16b648b4365e066f9f52d4d87bb923c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
