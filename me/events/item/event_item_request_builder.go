package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i76672221c83dc38179cae2efca7d08816ee9e597b77ddba836865157c3630aed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/accept"
    i79d1a9f8030b80fbfd719d034510b53c2ba1b45bb5d4273c13475346e7572187 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences"
    i8682926cfffc5b787c793b8059664b7eca389ad3c56f9ab1d1a400cd8d871693 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/multivalueextendedproperties"
    iac6711c67e4a14446e97de026ecbbcb570f073f3ef4fb36f8499a89c841b1251 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/tentativelyaccept"
    ib358e52ce114f81aaf3b8031c457fd405d0b7d8f5713d6d8f672f2b7da55c472 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances"
    ibe2dcd2c69e9c3f481a78bb4fcc0f38e09c50fb8b31605fc25647b0929a41625 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/snoozereminder"
    ic5c41c78200267f0ce789a523ec30cd54d492fed024852556a3861cd0fb2a5c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/extensions"
    ic6c4b9d285f086fc8176d404623fb1cdc195066c64b4f3b919160e4aaf83e9c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/singlevalueextendedproperties"
    ic8031f21cab994735c81a8358ed9272736d6c94e67001ce6e339aff6aadffa98 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/attachments"
    id1adeec63eafeac63e2a323cb6de5aad1091c62b1c34922e2bd17933aa7b90ce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/forward"
    iddd81fea2710e6f8c30fe8e6d8a6be2de8d8e183fbf870d75b4eb59fe43a869a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/cancel"
    ie0787d9509cd426004407c065c019a6b3b286a59c1bede43d1b0e756af114de6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar"
    if8385c0725bebab3400ac53450aa0a6fe210717d6702a895487379d999cababc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/decline"
    ifed870770ab60b2acfbdd9a09c0805c4a33032e5e0cf47c2a65312c4d1793b10 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/dismissreminder"
    i0b64b3ef52dbd3e72d88b3002e2346ed309a759fff57c4e03d03441f16395a61 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/multivalueextendedproperties/item"
    i333124becbdebfb8e2cde251c85664f8ce17b7da39c808fe1028c2198d118049 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/attachments/item"
    i4c10bf2ba0135e1c96fd2955ae7b4b167eb0eac796f4a9dc0131cb21bec5a5ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/singlevalueextendedproperties/item"
    i5c003d1900b66ad028b06d24a1b318a1bf2bd5542c6653c57d1ce9cdb602ae09 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item"
    i5d36ace97f3f76411ae8d4c97f0a2bfcdaefecd85098e56269cf2ce48f78c0ea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/extensions/item"
    iefbbcc1cd090052db6ae026949f16812722efc15791251bb87ca2d4861adb422 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.user entity.
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
// EventItemRequestBuilderGetQueryParameters the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Accept()(*i76672221c83dc38179cae2efca7d08816ee9e597b77ddba836865157c3630aed.AcceptRequestBuilder) {
    return i76672221c83dc38179cae2efca7d08816ee9e597b77ddba836865157c3630aed.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ic8031f21cab994735c81a8358ed9272736d6c94e67001ce6e339aff6aadffa98.AttachmentsRequestBuilder) {
    return ic8031f21cab994735c81a8358ed9272736d6c94e67001ce6e339aff6aadffa98.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i333124becbdebfb8e2cde251c85664f8ce17b7da39c808fe1028c2198d118049.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i333124becbdebfb8e2cde251c85664f8ce17b7da39c808fe1028c2198d118049.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ie0787d9509cd426004407c065c019a6b3b286a59c1bede43d1b0e756af114de6.CalendarRequestBuilder) {
    return ie0787d9509cd426004407c065c019a6b3b286a59c1bede43d1b0e756af114de6.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*iddd81fea2710e6f8c30fe8e6d8a6be2de8d8e183fbf870d75b4eb59fe43a869a.CancelRequestBuilder) {
    return iddd81fea2710e6f8c30fe8e6d8a6be2de8d8e183fbf870d75b4eb59fe43a869a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property events for me
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
// CreateGetRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property events in me
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
func (m *EventItemRequestBuilder) Decline()(*if8385c0725bebab3400ac53450aa0a6fe210717d6702a895487379d999cababc.DeclineRequestBuilder) {
    return if8385c0725bebab3400ac53450aa0a6fe210717d6702a895487379d999cababc.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*ifed870770ab60b2acfbdd9a09c0805c4a33032e5e0cf47c2a65312c4d1793b10.DismissReminderRequestBuilder) {
    return ifed870770ab60b2acfbdd9a09c0805c4a33032e5e0cf47c2a65312c4d1793b10.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i79d1a9f8030b80fbfd719d034510b53c2ba1b45bb5d4273c13475346e7572187.ExceptionOccurrencesRequestBuilder) {
    return i79d1a9f8030b80fbfd719d034510b53c2ba1b45bb5d4273c13475346e7572187.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i5c003d1900b66ad028b06d24a1b318a1bf2bd5542c6653c57d1ce9cdb602ae09.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i5c003d1900b66ad028b06d24a1b318a1bf2bd5542c6653c57d1ce9cdb602ae09.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*ic5c41c78200267f0ce789a523ec30cd54d492fed024852556a3861cd0fb2a5c5.ExtensionsRequestBuilder) {
    return ic5c41c78200267f0ce789a523ec30cd54d492fed024852556a3861cd0fb2a5c5.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i5d36ace97f3f76411ae8d4c97f0a2bfcdaefecd85098e56269cf2ce48f78c0ea.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i5d36ace97f3f76411ae8d4c97f0a2bfcdaefecd85098e56269cf2ce48f78c0ea.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*id1adeec63eafeac63e2a323cb6de5aad1091c62b1c34922e2bd17933aa7b90ce.ForwardRequestBuilder) {
    return id1adeec63eafeac63e2a323cb6de5aad1091c62b1c34922e2bd17933aa7b90ce.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
func (m *EventItemRequestBuilder) Instances()(*ib358e52ce114f81aaf3b8031c457fd405d0b7d8f5713d6d8f672f2b7da55c472.InstancesRequestBuilder) {
    return ib358e52ce114f81aaf3b8031c457fd405d0b7d8f5713d6d8f672f2b7da55c472.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*iefbbcc1cd090052db6ae026949f16812722efc15791251bb87ca2d4861adb422.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return iefbbcc1cd090052db6ae026949f16812722efc15791251bb87ca2d4861adb422.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i8682926cfffc5b787c793b8059664b7eca389ad3c56f9ab1d1a400cd8d871693.MultiValueExtendedPropertiesRequestBuilder) {
    return i8682926cfffc5b787c793b8059664b7eca389ad3c56f9ab1d1a400cd8d871693.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0b64b3ef52dbd3e72d88b3002e2346ed309a759fff57c4e03d03441f16395a61.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0b64b3ef52dbd3e72d88b3002e2346ed309a759fff57c4e03d03441f16395a61.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ic6c4b9d285f086fc8176d404623fb1cdc195066c64b4f3b919160e4aaf83e9c8.SingleValueExtendedPropertiesRequestBuilder) {
    return ic6c4b9d285f086fc8176d404623fb1cdc195066c64b4f3b919160e4aaf83e9c8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4c10bf2ba0135e1c96fd2955ae7b4b167eb0eac796f4a9dc0131cb21bec5a5ad.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i4c10bf2ba0135e1c96fd2955ae7b4b167eb0eac796f4a9dc0131cb21bec5a5ad.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ibe2dcd2c69e9c3f481a78bb4fcc0f38e09c50fb8b31605fc25647b0929a41625.SnoozeReminderRequestBuilder) {
    return ibe2dcd2c69e9c3f481a78bb4fcc0f38e09c50fb8b31605fc25647b0929a41625.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*iac6711c67e4a14446e97de026ecbbcb570f073f3ef4fb36f8499a89c841b1251.TentativelyAcceptRequestBuilder) {
    return iac6711c67e4a14446e97de026ecbbcb570f073f3ef4fb36f8499a89c841b1251.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
