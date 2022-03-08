package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0bd35b1fb9b5cd3a509560e24ec57494bb35171b5ffa0a6e007c91ce2ba419c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/accept"
    i187c9e6816ffbc23be0deb69b1a127130ae37c10edf1254e9b5826ff6031ba01 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    i332f925e36b2fe945c50b1b6d90c949acde78fca9bdddf4f0487fd904ab368b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i4476c63e8cf22dbd915a41d4ab33b7425e126d395bceb40bad6ab35569ae35d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i4956ec783cf4129a29c7531c7d13a3f7bb3f5c830d1c90ee045cecaefe3288be "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i4cab08f986acca114d1529f5b639030b91f1da7c714ffc08f80e31221192dcd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i7a60b93f7d0221f3b1402a67920fc82eb7d36db595295f22b24ff301b1ec7098 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i9248a10c6abd9ed71d01c49cf1e461e9e9b06c0ef44578f7d33b2c9e99cc2aba "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ia4ef4f75b96a0768357c29a43a6454f3b1a8ae188bee0c585960a7babc17654f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/forward"
    ia812ae1dc4b1f30fa3f2a099e78c863c203a7f19deb5c36ac0ed96f2e1514a3c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/decline"
    ic66635521102afb8cd7214f483378adefcbe7d0f153f0a7ef85b1da938364825 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    icca8571980095ab1b8115b8579ce64b33631d293ab438a53752a6360e4fd2154 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i4f5f241ea2012d39eb24707ce6d64e3cb86fc78310cdccb1e0caeff5547fa816 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    i558f96fb156da918ec2245ed183a904c7c5e6ce361c4ca7e14bb110bbf083ee2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    ic3ef47a5c760d3b34d87fdce711d26557f319a1afff6d092446c7dc99656869b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    ie3ae8b75285bede0c6f450453901c8aea50863d5c526729b8f466fc8a489ce60 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*i0bd35b1fb9b5cd3a509560e24ec57494bb35171b5ffa0a6e007c91ce2ba419c9.AcceptRequestBuilder) {
    return i0bd35b1fb9b5cd3a509560e24ec57494bb35171b5ffa0a6e007c91ce2ba419c9.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i7a60b93f7d0221f3b1402a67920fc82eb7d36db595295f22b24ff301b1ec7098.AttachmentsRequestBuilder) {
    return i7a60b93f7d0221f3b1402a67920fc82eb7d36db595295f22b24ff301b1ec7098.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i558f96fb156da918ec2245ed183a904c7c5e6ce361c4ca7e14bb110bbf083ee2.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i558f96fb156da918ec2245ed183a904c7c5e6ce361c4ca7e14bb110bbf083ee2.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i187c9e6816ffbc23be0deb69b1a127130ae37c10edf1254e9b5826ff6031ba01.CalendarRequestBuilder) {
    return i187c9e6816ffbc23be0deb69b1a127130ae37c10edf1254e9b5826ff6031ba01.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*icca8571980095ab1b8115b8579ce64b33631d293ab438a53752a6360e4fd2154.CancelRequestBuilder) {
    return icca8571980095ab1b8115b8579ce64b33631d293ab438a53752a6360e4fd2154.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/events/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
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
// CreateDeleteRequestInformation delete navigation property instances for groups
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
// CreatePatchRequestInformation update the navigation property instances in groups
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
func (m *EventItemRequestBuilder) Decline()(*ia812ae1dc4b1f30fa3f2a099e78c863c203a7f19deb5c36ac0ed96f2e1514a3c.DeclineRequestBuilder) {
    return ia812ae1dc4b1f30fa3f2a099e78c863c203a7f19deb5c36ac0ed96f2e1514a3c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*i4cab08f986acca114d1529f5b639030b91f1da7c714ffc08f80e31221192dcd3.DismissReminderRequestBuilder) {
    return i4cab08f986acca114d1529f5b639030b91f1da7c714ffc08f80e31221192dcd3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i4956ec783cf4129a29c7531c7d13a3f7bb3f5c830d1c90ee045cecaefe3288be.ExtensionsRequestBuilder) {
    return i4956ec783cf4129a29c7531c7d13a3f7bb3f5c830d1c90ee045cecaefe3288be.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ie3ae8b75285bede0c6f450453901c8aea50863d5c526729b8f466fc8a489ce60.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ie3ae8b75285bede0c6f450453901c8aea50863d5c526729b8f466fc8a489ce60.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ia4ef4f75b96a0768357c29a43a6454f3b1a8ae188bee0c585960a7babc17654f.ForwardRequestBuilder) {
    return ia4ef4f75b96a0768357c29a43a6454f3b1a8ae188bee0c585960a7babc17654f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4476c63e8cf22dbd915a41d4ab33b7425e126d395bceb40bad6ab35569ae35d9.MultiValueExtendedPropertiesRequestBuilder) {
    return i4476c63e8cf22dbd915a41d4ab33b7425e126d395bceb40bad6ab35569ae35d9.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4f5f241ea2012d39eb24707ce6d64e3cb86fc78310cdccb1e0caeff5547fa816.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i4f5f241ea2012d39eb24707ce6d64e3cb86fc78310cdccb1e0caeff5547fa816.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i332f925e36b2fe945c50b1b6d90c949acde78fca9bdddf4f0487fd904ab368b1.SingleValueExtendedPropertiesRequestBuilder) {
    return i332f925e36b2fe945c50b1b6d90c949acde78fca9bdddf4f0487fd904ab368b1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic3ef47a5c760d3b34d87fdce711d26557f319a1afff6d092446c7dc99656869b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic3ef47a5c760d3b34d87fdce711d26557f319a1afff6d092446c7dc99656869b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i9248a10c6abd9ed71d01c49cf1e461e9e9b06c0ef44578f7d33b2c9e99cc2aba.SnoozeReminderRequestBuilder) {
    return i9248a10c6abd9ed71d01c49cf1e461e9e9b06c0ef44578f7d33b2c9e99cc2aba.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ic66635521102afb8cd7214f483378adefcbe7d0f153f0a7ef85b1da938364825.TentativelyAcceptRequestBuilder) {
    return ic66635521102afb8cd7214f483378adefcbe7d0f153f0a7ef85b1da938364825.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
