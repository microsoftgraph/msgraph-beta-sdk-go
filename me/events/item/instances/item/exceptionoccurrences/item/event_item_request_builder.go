package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i235c5ce0d1f74d524ef46dbde3ad46ff72b77628085dbc6bfdac40c1fb48c23c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i29a8545864fc805f0e2701a8b3eec5af1142195c0187b2c6968094d3b343fc07 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/accept"
    i3bd23c78fcfbc6b0f4baf49aa2633f97e3d708ae8375a9cce895ad4a6a73a76a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/attachments"
    i5e5077c321fc6e74b2a81f7770dae70a1666eda8cfcd968f487525cb202d1690 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/forward"
    i62af91d8429dd5f4cb597ef41ed6e71511fb493b4b5fcbff8a4e59b7b859e6be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i70176bee18712bf6effaf06a937591a269c52462ddd3254f0c639fa8193e8257 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/decline"
    i76a3e45766626f05b2e1bfda4f0a8e614736e1ed200ed9d1ecd953a7c21f29b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i782c18674076087ebdcded13180ab6c2573441036e3b850dbacd6508343032fb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/extensions"
    i782f63d0c32eb12865e32070ff0d30a940280f9a88402ee6b94497170ac21eda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i81438a5134e8af20e7bb64724e4a869db0f1040f8c7d4b6651b4a8c4e66758bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/cancel"
    iaef4803cc86320b3f857eae04b3c7a0ea4111e2a487f8e3cc015656cfbd2f476 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/calendar"
    ibe378cc2630ababf55a71d5b746f507d3433811e5b1747b548c290c5f774ebd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i5a7159f62dfc28b5ce002654b03796b6b64a8bf6334f48f84d36040205821d66 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    i6030ecc1ae42e4cc03f29341a683734a1c1a632a2e1bc1df69f744da21821d07 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    iacd3eefe1e8b57f639bf54342be2037e4646af4aa9296f8557c19be991af64bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    ife47cecb8e20b7ce7a99756a6a698281988672e18019307b3efe508888854f1e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i29a8545864fc805f0e2701a8b3eec5af1142195c0187b2c6968094d3b343fc07.AcceptRequestBuilder) {
    return i29a8545864fc805f0e2701a8b3eec5af1142195c0187b2c6968094d3b343fc07.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i3bd23c78fcfbc6b0f4baf49aa2633f97e3d708ae8375a9cce895ad4a6a73a76a.AttachmentsRequestBuilder) {
    return i3bd23c78fcfbc6b0f4baf49aa2633f97e3d708ae8375a9cce895ad4a6a73a76a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i5a7159f62dfc28b5ce002654b03796b6b64a8bf6334f48f84d36040205821d66.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i5a7159f62dfc28b5ce002654b03796b6b64a8bf6334f48f84d36040205821d66.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*iaef4803cc86320b3f857eae04b3c7a0ea4111e2a487f8e3cc015656cfbd2f476.CalendarRequestBuilder) {
    return iaef4803cc86320b3f857eae04b3c7a0ea4111e2a487f8e3cc015656cfbd2f476.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i81438a5134e8af20e7bb64724e4a869db0f1040f8c7d4b6651b4a8c4e66758bd.CancelRequestBuilder) {
    return i81438a5134e8af20e7bb64724e4a869db0f1040f8c7d4b6651b4a8c4e66758bd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i70176bee18712bf6effaf06a937591a269c52462ddd3254f0c639fa8193e8257.DeclineRequestBuilder) {
    return i70176bee18712bf6effaf06a937591a269c52462ddd3254f0c639fa8193e8257.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i235c5ce0d1f74d524ef46dbde3ad46ff72b77628085dbc6bfdac40c1fb48c23c.DismissReminderRequestBuilder) {
    return i235c5ce0d1f74d524ef46dbde3ad46ff72b77628085dbc6bfdac40c1fb48c23c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i782c18674076087ebdcded13180ab6c2573441036e3b850dbacd6508343032fb.ExtensionsRequestBuilder) {
    return i782c18674076087ebdcded13180ab6c2573441036e3b850dbacd6508343032fb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iacd3eefe1e8b57f639bf54342be2037e4646af4aa9296f8557c19be991af64bd.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iacd3eefe1e8b57f639bf54342be2037e4646af4aa9296f8557c19be991af64bd.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i5e5077c321fc6e74b2a81f7770dae70a1666eda8cfcd968f487525cb202d1690.ForwardRequestBuilder) {
    return i5e5077c321fc6e74b2a81f7770dae70a1666eda8cfcd968f487525cb202d1690.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i62af91d8429dd5f4cb597ef41ed6e71511fb493b4b5fcbff8a4e59b7b859e6be.MultiValueExtendedPropertiesRequestBuilder) {
    return i62af91d8429dd5f4cb597ef41ed6e71511fb493b4b5fcbff8a4e59b7b859e6be.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ife47cecb8e20b7ce7a99756a6a698281988672e18019307b3efe508888854f1e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ife47cecb8e20b7ce7a99756a6a698281988672e18019307b3efe508888854f1e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i76a3e45766626f05b2e1bfda4f0a8e614736e1ed200ed9d1ecd953a7c21f29b5.SingleValueExtendedPropertiesRequestBuilder) {
    return i76a3e45766626f05b2e1bfda4f0a8e614736e1ed200ed9d1ecd953a7c21f29b5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i6030ecc1ae42e4cc03f29341a683734a1c1a632a2e1bc1df69f744da21821d07.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i6030ecc1ae42e4cc03f29341a683734a1c1a632a2e1bc1df69f744da21821d07.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i782f63d0c32eb12865e32070ff0d30a940280f9a88402ee6b94497170ac21eda.SnoozeReminderRequestBuilder) {
    return i782f63d0c32eb12865e32070ff0d30a940280f9a88402ee6b94497170ac21eda.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibe378cc2630ababf55a71d5b746f507d3433811e5b1747b548c290c5f774ebd7.TentativelyAcceptRequestBuilder) {
    return ibe378cc2630ababf55a71d5b746f507d3433811e5b1747b548c290c5f774ebd7.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
