package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1a1b8e96337bb7483cd78fdab2be9dc44c98f840c502bb01623bf2c8635861d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i3e0a8a2d61feeca0b653d4e764381c75f333d29e9a09dec9e6635b943cf9a442 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/extensions"
    i474e3769a1767d1d69d84c46a7163657df993d904b05e638908eae6b74e0222f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/calendar"
    i48741b209367b947e80c69f8dafc5395630bf399c3d30c3e70c93e56a0458fec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/attachments"
    i6d4b26098c6d76561c8632d71cb718afdd8fad1e34731d673df511805f92e5c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    iad4c8c603072a05a962c46f93ca2878e8e14e9a12dfff682d5bdf8dc173c5a1a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ib1bf60d851d8d6e9d204f1374cc4ddb1161bd577880143a8bbf57c029c09c70f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/decline"
    ic2e266faf9f785da86116cb80205b5c271ba4d18f79f696c49c81faa408f380a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/cancel"
    ic87a4a083820dc975406d5c3e92133fc615efe8758b56d91bed9a68b2a587ac4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    ieae867b50b9e152a1ab48742be0f1f4abfa26db00262db4ccda4b53c034ce012 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    ied1372a1dd50f4b8c1fc7d1bd03cbe4c8f9fcf0b6a561e3134afbc9f83239c5c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/accept"
    if3e3ba4a4b11173947afda79839dfa5758bf4f69e33002dd325bdc306bdd87d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/forward"
    i25792bfcf313e985fa426ee30854e45096449111352e45745170588d0784798a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i31a1d7487fb1f3c869c614e2016b3bdca8b10cd0671e6d3becd10465fba76ee2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    i564d51fe6cdcce1006d5d9ff834a56883aea8737c8b5641b6aee47127dd72c2b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    iff3cb40fe773031a6efad6cf45453191bc2cc27d0941c6369e0bf3ddc4f204b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item/exceptionoccurrences/item/extensions/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Accept()(*ied1372a1dd50f4b8c1fc7d1bd03cbe4c8f9fcf0b6a561e3134afbc9f83239c5c.AcceptRequestBuilder) {
    return ied1372a1dd50f4b8c1fc7d1bd03cbe4c8f9fcf0b6a561e3134afbc9f83239c5c.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i48741b209367b947e80c69f8dafc5395630bf399c3d30c3e70c93e56a0458fec.AttachmentsRequestBuilder) {
    return i48741b209367b947e80c69f8dafc5395630bf399c3d30c3e70c93e56a0458fec.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i31a1d7487fb1f3c869c614e2016b3bdca8b10cd0671e6d3becd10465fba76ee2.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i31a1d7487fb1f3c869c614e2016b3bdca8b10cd0671e6d3becd10465fba76ee2.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i474e3769a1767d1d69d84c46a7163657df993d904b05e638908eae6b74e0222f.CalendarRequestBuilder) {
    return i474e3769a1767d1d69d84c46a7163657df993d904b05e638908eae6b74e0222f.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ic2e266faf9f785da86116cb80205b5c271ba4d18f79f696c49c81faa408f380a.CancelRequestBuilder) {
    return ic2e266faf9f785da86116cb80205b5c271ba4d18f79f696c49c81faa408f380a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/events/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for me
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) Decline()(*ib1bf60d851d8d6e9d204f1374cc4ddb1161bd577880143a8bbf57c029c09c70f.DeclineRequestBuilder) {
    return ib1bf60d851d8d6e9d204f1374cc4ddb1161bd577880143a8bbf57c029c09c70f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*ieae867b50b9e152a1ab48742be0f1f4abfa26db00262db4ccda4b53c034ce012.DismissReminderRequestBuilder) {
    return ieae867b50b9e152a1ab48742be0f1f4abfa26db00262db4ccda4b53c034ce012.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i3e0a8a2d61feeca0b653d4e764381c75f333d29e9a09dec9e6635b943cf9a442.ExtensionsRequestBuilder) {
    return i3e0a8a2d61feeca0b653d4e764381c75f333d29e9a09dec9e6635b943cf9a442.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iff3cb40fe773031a6efad6cf45453191bc2cc27d0941c6369e0bf3ddc4f204b6.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iff3cb40fe773031a6efad6cf45453191bc2cc27d0941c6369e0bf3ddc4f204b6.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*if3e3ba4a4b11173947afda79839dfa5758bf4f69e33002dd325bdc306bdd87d7.ForwardRequestBuilder) {
    return if3e3ba4a4b11173947afda79839dfa5758bf4f69e33002dd325bdc306bdd87d7.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic87a4a083820dc975406d5c3e92133fc615efe8758b56d91bed9a68b2a587ac4.MultiValueExtendedPropertiesRequestBuilder) {
    return ic87a4a083820dc975406d5c3e92133fc615efe8758b56d91bed9a68b2a587ac4.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i564d51fe6cdcce1006d5d9ff834a56883aea8737c8b5641b6aee47127dd72c2b.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i564d51fe6cdcce1006d5d9ff834a56883aea8737c8b5641b6aee47127dd72c2b.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6d4b26098c6d76561c8632d71cb718afdd8fad1e34731d673df511805f92e5c1.SingleValueExtendedPropertiesRequestBuilder) {
    return i6d4b26098c6d76561c8632d71cb718afdd8fad1e34731d673df511805f92e5c1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i25792bfcf313e985fa426ee30854e45096449111352e45745170588d0784798a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i25792bfcf313e985fa426ee30854e45096449111352e45745170588d0784798a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i1a1b8e96337bb7483cd78fdab2be9dc44c98f840c502bb01623bf2c8635861d3.SnoozeReminderRequestBuilder) {
    return i1a1b8e96337bb7483cd78fdab2be9dc44c98f840c502bb01623bf2c8635861d3.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*iad4c8c603072a05a962c46f93ca2878e8e14e9a12dfff682d5bdf8dc173c5a1a.TentativelyAcceptRequestBuilder) {
    return iad4c8c603072a05a962c46f93ca2878e8e14e9a12dfff682d5bdf8dc173c5a1a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
