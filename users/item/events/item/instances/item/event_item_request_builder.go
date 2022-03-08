package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1720a7ceaf3a740e21c9791ba9f21c078ab46ff77082619efd43c176037bb89e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences"
    i3d659f8c84dfc1f6b2b006a12a81e7be6b96dfab59caa73e60d4efc55240c65a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/cancel"
    i5067c919585ed58aa0ff4419be188695f4ea434f95e904e133fb55c243e0f412 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/decline"
    i50f61aabcf2a79623925778b32a0fec15688127120f50d91d5020d507c6871f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/extensions"
    i6264160d0e28137d4fa29f7b97bf6a5fbf4630d97b9ac0f0dec8732517c798c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/dismissreminder"
    i7f3e2c5a1718c681447bf953dd3677a4b16d3e85d9acd12da0c931a4b09bb247 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/singlevalueextendedproperties"
    i96523e4b6e04037100c5f389ff5cd546c675baee933f2d6f865989bc7f0a1d62 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/tentativelyaccept"
    i9dba6787329bb8194a05f7c8fd25bf2bc96a52038aa19d2f8ef2d81f23971ed1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/attachments"
    ia1a281eee7df2f7f6df2e895554b89c3b22dedfa1d2e72c4091910213a41d2b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/accept"
    ib75eaee188679d344e02b1c694adf0e831340071613135df137739b7c2da40d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/multivalueextendedproperties"
    ibe8f97064ae05d0c65f04090a3f4567fc177dfe0cfee1e9fd999a0a800b49306 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/calendar"
    icced7e83203c2072101b4334ccbb6da1f2954bbd8165212bebc60f6b2b4c580f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/forward"
    ie2ed44c27ab48f75b6e944495110e64c56f311242dfb0d911a1575d3cfcbcb58 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/snoozereminder"
    i4b97a44a86b32435a8809147dcc12ac5bfcfbcad953e60210b88761b82f9993a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/multivalueextendedproperties/item"
    i75800b91c84735d858668798cd39fe4fe96e3985383f60b7823ea901d12ac7a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/attachments/item"
    ia15eb645068ac102741848bdc27f33ed7d20533b0b43980ca6d4507bf29f4dab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item"
    ibbaea04b51d7b7766d5ad02712ef5717a7974fef9b22c7011c5dc94e72d5cd81 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/singlevalueextendedproperties/item"
    idba18abf5aea5eb2239f0ecdd22b5f56420648397142d15f778264867414b2cc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*ia1a281eee7df2f7f6df2e895554b89c3b22dedfa1d2e72c4091910213a41d2b8.AcceptRequestBuilder) {
    return ia1a281eee7df2f7f6df2e895554b89c3b22dedfa1d2e72c4091910213a41d2b8.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i9dba6787329bb8194a05f7c8fd25bf2bc96a52038aa19d2f8ef2d81f23971ed1.AttachmentsRequestBuilder) {
    return i9dba6787329bb8194a05f7c8fd25bf2bc96a52038aa19d2f8ef2d81f23971ed1.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i75800b91c84735d858668798cd39fe4fe96e3985383f60b7823ea901d12ac7a7.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i75800b91c84735d858668798cd39fe4fe96e3985383f60b7823ea901d12ac7a7.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ibe8f97064ae05d0c65f04090a3f4567fc177dfe0cfee1e9fd999a0a800b49306.CalendarRequestBuilder) {
    return ibe8f97064ae05d0c65f04090a3f4567fc177dfe0cfee1e9fd999a0a800b49306.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i3d659f8c84dfc1f6b2b006a12a81e7be6b96dfab59caa73e60d4efc55240c65a.CancelRequestBuilder) {
    return i3d659f8c84dfc1f6b2b006a12a81e7be6b96dfab59caa73e60d4efc55240c65a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/events/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i5067c919585ed58aa0ff4419be188695f4ea434f95e904e133fb55c243e0f412.DeclineRequestBuilder) {
    return i5067c919585ed58aa0ff4419be188695f4ea434f95e904e133fb55c243e0f412.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i6264160d0e28137d4fa29f7b97bf6a5fbf4630d97b9ac0f0dec8732517c798c8.DismissReminderRequestBuilder) {
    return i6264160d0e28137d4fa29f7b97bf6a5fbf4630d97b9ac0f0dec8732517c798c8.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i1720a7ceaf3a740e21c9791ba9f21c078ab46ff77082619efd43c176037bb89e.ExceptionOccurrencesRequestBuilder) {
    return i1720a7ceaf3a740e21c9791ba9f21c078ab46ff77082619efd43c176037bb89e.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ia15eb645068ac102741848bdc27f33ed7d20533b0b43980ca6d4507bf29f4dab.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return ia15eb645068ac102741848bdc27f33ed7d20533b0b43980ca6d4507bf29f4dab.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i50f61aabcf2a79623925778b32a0fec15688127120f50d91d5020d507c6871f8.ExtensionsRequestBuilder) {
    return i50f61aabcf2a79623925778b32a0fec15688127120f50d91d5020d507c6871f8.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*idba18abf5aea5eb2239f0ecdd22b5f56420648397142d15f778264867414b2cc.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return idba18abf5aea5eb2239f0ecdd22b5f56420648397142d15f778264867414b2cc.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*icced7e83203c2072101b4334ccbb6da1f2954bbd8165212bebc60f6b2b4c580f.ForwardRequestBuilder) {
    return icced7e83203c2072101b4334ccbb6da1f2954bbd8165212bebc60f6b2b4c580f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ib75eaee188679d344e02b1c694adf0e831340071613135df137739b7c2da40d6.MultiValueExtendedPropertiesRequestBuilder) {
    return ib75eaee188679d344e02b1c694adf0e831340071613135df137739b7c2da40d6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4b97a44a86b32435a8809147dcc12ac5bfcfbcad953e60210b88761b82f9993a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i4b97a44a86b32435a8809147dcc12ac5bfcfbcad953e60210b88761b82f9993a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i7f3e2c5a1718c681447bf953dd3677a4b16d3e85d9acd12da0c931a4b09bb247.SingleValueExtendedPropertiesRequestBuilder) {
    return i7f3e2c5a1718c681447bf953dd3677a4b16d3e85d9acd12da0c931a4b09bb247.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ibbaea04b51d7b7766d5ad02712ef5717a7974fef9b22c7011c5dc94e72d5cd81.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ibbaea04b51d7b7766d5ad02712ef5717a7974fef9b22c7011c5dc94e72d5cd81.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie2ed44c27ab48f75b6e944495110e64c56f311242dfb0d911a1575d3cfcbcb58.SnoozeReminderRequestBuilder) {
    return ie2ed44c27ab48f75b6e944495110e64c56f311242dfb0d911a1575d3cfcbcb58.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i96523e4b6e04037100c5f389ff5cd546c675baee933f2d6f865989bc7f0a1d62.TentativelyAcceptRequestBuilder) {
    return i96523e4b6e04037100c5f389ff5cd546c675baee933f2d6f865989bc7f0a1d62.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
