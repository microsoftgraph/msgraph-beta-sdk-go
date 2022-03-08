package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i25d5cacbc332f7c80f6d55863a6d428492b427013d7f7e02591f70c76d580b1b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/forward"
    i263083a993cdb11d56373107896c6d703bbc61ee150798812cf0a179d743bfca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i27a9d3a5d46a574bc703506a155b15b6db0bcdc1b332e230168664dcf238c5c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i297679d1108e0046358c9f2a7c894d5d3a54b4ea2b60dae94d5c45e491992798 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/accept"
    i36822135361ed06fe29d1595a97c920b52594358304f1e4733b2a73310c2fdc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/cancel"
    i430cdd8a7badf301c4fd80bbeaf7da6104a4ae0e5d2f4c2a22029b8bc11e61dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i4953e43cca0dc3b0c53b250fa449ac267acc518de450c9eaee996573fce62357 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/calendar"
    i866a8d8a74dd530c64622d5cc865a55ff79af06bad1b9b0b417686feeb8337fe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/extensions"
    i9914c84b840e3457f1c267f3af8ac42fc5eccc06fc8403b8f13a4863b157812e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ib164dfa1fd0114ca131b8bea09446692ea3fd5b5017d795b3766e154e5235fb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/decline"
    ib17018b7fb7705ca721082cb344a6981b2b895834bd0e8d0bb5001ea795cb568 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/attachments"
    iee4a320aef5c207016c7b92107552b85e50517335291a79d598b92333086f173 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i1f618d495b3c5d7c336d9a185e85e8bd5d40881d124784f7bcdfc5e19f750178 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    i5f9b04d94995a7e5e51b76524467b526c1ffd173b3a2e880395f46dcd276bc8e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    idabf8e5a57753a10fd9f8fca7e75b0e5bb8cd819449af06f6ed5e113bc25eb04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    iec4ffec5d16fb93c8e5806e65d18a9b3dfe689b6fe575d435360b154e7d26b31 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i297679d1108e0046358c9f2a7c894d5d3a54b4ea2b60dae94d5c45e491992798.AcceptRequestBuilder) {
    return i297679d1108e0046358c9f2a7c894d5d3a54b4ea2b60dae94d5c45e491992798.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ib17018b7fb7705ca721082cb344a6981b2b895834bd0e8d0bb5001ea795cb568.AttachmentsRequestBuilder) {
    return ib17018b7fb7705ca721082cb344a6981b2b895834bd0e8d0bb5001ea795cb568.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*idabf8e5a57753a10fd9f8fca7e75b0e5bb8cd819449af06f6ed5e113bc25eb04.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return idabf8e5a57753a10fd9f8fca7e75b0e5bb8cd819449af06f6ed5e113bc25eb04.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i4953e43cca0dc3b0c53b250fa449ac267acc518de450c9eaee996573fce62357.CalendarRequestBuilder) {
    return i4953e43cca0dc3b0c53b250fa449ac267acc518de450c9eaee996573fce62357.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i36822135361ed06fe29d1595a97c920b52594358304f1e4733b2a73310c2fdc8.CancelRequestBuilder) {
    return i36822135361ed06fe29d1595a97c920b52594358304f1e4733b2a73310c2fdc8.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/events/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
func (m *EventItemRequestBuilder) Decline()(*ib164dfa1fd0114ca131b8bea09446692ea3fd5b5017d795b3766e154e5235fb4.DeclineRequestBuilder) {
    return ib164dfa1fd0114ca131b8bea09446692ea3fd5b5017d795b3766e154e5235fb4.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i27a9d3a5d46a574bc703506a155b15b6db0bcdc1b332e230168664dcf238c5c2.DismissReminderRequestBuilder) {
    return i27a9d3a5d46a574bc703506a155b15b6db0bcdc1b332e230168664dcf238c5c2.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i866a8d8a74dd530c64622d5cc865a55ff79af06bad1b9b0b417686feeb8337fe.ExtensionsRequestBuilder) {
    return i866a8d8a74dd530c64622d5cc865a55ff79af06bad1b9b0b417686feeb8337fe.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i1f618d495b3c5d7c336d9a185e85e8bd5d40881d124784f7bcdfc5e19f750178.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i1f618d495b3c5d7c336d9a185e85e8bd5d40881d124784f7bcdfc5e19f750178.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i25d5cacbc332f7c80f6d55863a6d428492b427013d7f7e02591f70c76d580b1b.ForwardRequestBuilder) {
    return i25d5cacbc332f7c80f6d55863a6d428492b427013d7f7e02591f70c76d580b1b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i263083a993cdb11d56373107896c6d703bbc61ee150798812cf0a179d743bfca.MultiValueExtendedPropertiesRequestBuilder) {
    return i263083a993cdb11d56373107896c6d703bbc61ee150798812cf0a179d743bfca.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iec4ffec5d16fb93c8e5806e65d18a9b3dfe689b6fe575d435360b154e7d26b31.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iec4ffec5d16fb93c8e5806e65d18a9b3dfe689b6fe575d435360b154e7d26b31.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i9914c84b840e3457f1c267f3af8ac42fc5eccc06fc8403b8f13a4863b157812e.SingleValueExtendedPropertiesRequestBuilder) {
    return i9914c84b840e3457f1c267f3af8ac42fc5eccc06fc8403b8f13a4863b157812e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5f9b04d94995a7e5e51b76524467b526c1ffd173b3a2e880395f46dcd276bc8e.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5f9b04d94995a7e5e51b76524467b526c1ffd173b3a2e880395f46dcd276bc8e.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*iee4a320aef5c207016c7b92107552b85e50517335291a79d598b92333086f173.SnoozeReminderRequestBuilder) {
    return iee4a320aef5c207016c7b92107552b85e50517335291a79d598b92333086f173.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i430cdd8a7badf301c4fd80bbeaf7da6104a4ae0e5d2f4c2a22029b8bc11e61dc.TentativelyAcceptRequestBuilder) {
    return i430cdd8a7badf301c4fd80bbeaf7da6104a4ae0e5d2f4c2a22029b8bc11e61dc.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
