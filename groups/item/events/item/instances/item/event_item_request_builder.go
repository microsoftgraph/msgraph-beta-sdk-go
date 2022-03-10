package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i079fc82ebc66b91d2f1e5db2f46e14a93047f055730594a3c555217f5240b7d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/forward"
    i145ee4ad288f28c56f916c936201460b07ba8171fb97b6049f9c5b75da266ad2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/snoozereminder"
    i39b113e521727af9278341f568dc06471a9cdc7b19e99eedae87afbe5bcacb40 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/singlevalueextendedproperties"
    i5b0755b4a44cd9894e5545390c7357eb1b4714cdea7d58057e113f210e1e2c9a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/extensions"
    i5bf66d6159a00f290e9608799894e6a26635043d285c7f964cfdef336a4e70fc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/attachments"
    i677a95ac0890a8b4be07a112c61b91c390467b7f02a7cb3c8ad0e5f612695e9e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/tentativelyaccept"
    i689c547b0466816be72f6ed7126d1d817413b9d6a49872d2c88986517f12d48d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences"
    i6bf706d5338b74ee361af755175f5ec65af545ae444e30c37b562789b5935ef0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/accept"
    i74ccd3f22fbacb438df996a11eab6efcecf784f34164f51b4e63d80efbf60810 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/multivalueextendedproperties"
    i9d94ecbd137503e20f96ee1ce31d8792e15b81445f5394a489a51af80c70fede "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/decline"
    ia90f3bd5bd033bb6b91fdd2f5f002750f3f2f7943bff6ac7d13877cce888f9b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/dismissreminder"
    ic84861045d14b6e24fc34141d241b6e078098ca5abb7a26e747cb444172fb401 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/cancel"
    id7357db391ff461daa943136d564c93fd47f99013d9fdef7dd53b196022da292 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/calendar"
    i4419e2cbdd63fd5f009ec291e624ddf6e7f7e7e969ba42696d7e0a300f5e32c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/extensions/item"
    i554a181b0fba8aaa263c329ec9f30c8c0170b08d3f8ec9364659a4b1ce07ed8f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/multivalueextendedproperties/item"
    i7dfeae46798520fb15d6b541a64a5451c69c35f2a0deed35ba1e83b44eac510b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/singlevalueextendedproperties/item"
    i92a640ce65e53720f133cc9c105b150c718e27e7fb81d273bf327aee939f0bd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/attachments/item"
    ib74fe045468a2fd2c8f9a50d505bb56fb3152d793e1e1d7004f932eda018c8b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*i6bf706d5338b74ee361af755175f5ec65af545ae444e30c37b562789b5935ef0.AcceptRequestBuilder) {
    return i6bf706d5338b74ee361af755175f5ec65af545ae444e30c37b562789b5935ef0.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i5bf66d6159a00f290e9608799894e6a26635043d285c7f964cfdef336a4e70fc.AttachmentsRequestBuilder) {
    return i5bf66d6159a00f290e9608799894e6a26635043d285c7f964cfdef336a4e70fc.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i92a640ce65e53720f133cc9c105b150c718e27e7fb81d273bf327aee939f0bd8.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i92a640ce65e53720f133cc9c105b150c718e27e7fb81d273bf327aee939f0bd8.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*id7357db391ff461daa943136d564c93fd47f99013d9fdef7dd53b196022da292.CalendarRequestBuilder) {
    return id7357db391ff461daa943136d564c93fd47f99013d9fdef7dd53b196022da292.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ic84861045d14b6e24fc34141d241b6e078098ca5abb7a26e747cb444172fb401.CancelRequestBuilder) {
    return ic84861045d14b6e24fc34141d241b6e078098ca5abb7a26e747cb444172fb401.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/events/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i9d94ecbd137503e20f96ee1ce31d8792e15b81445f5394a489a51af80c70fede.DeclineRequestBuilder) {
    return i9d94ecbd137503e20f96ee1ce31d8792e15b81445f5394a489a51af80c70fede.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*ia90f3bd5bd033bb6b91fdd2f5f002750f3f2f7943bff6ac7d13877cce888f9b9.DismissReminderRequestBuilder) {
    return ia90f3bd5bd033bb6b91fdd2f5f002750f3f2f7943bff6ac7d13877cce888f9b9.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i689c547b0466816be72f6ed7126d1d817413b9d6a49872d2c88986517f12d48d.ExceptionOccurrencesRequestBuilder) {
    return i689c547b0466816be72f6ed7126d1d817413b9d6a49872d2c88986517f12d48d.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ib74fe045468a2fd2c8f9a50d505bb56fb3152d793e1e1d7004f932eda018c8b9.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return ib74fe045468a2fd2c8f9a50d505bb56fb3152d793e1e1d7004f932eda018c8b9.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i5b0755b4a44cd9894e5545390c7357eb1b4714cdea7d58057e113f210e1e2c9a.ExtensionsRequestBuilder) {
    return i5b0755b4a44cd9894e5545390c7357eb1b4714cdea7d58057e113f210e1e2c9a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i4419e2cbdd63fd5f009ec291e624ddf6e7f7e7e969ba42696d7e0a300f5e32c1.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i4419e2cbdd63fd5f009ec291e624ddf6e7f7e7e969ba42696d7e0a300f5e32c1.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i079fc82ebc66b91d2f1e5db2f46e14a93047f055730594a3c555217f5240b7d9.ForwardRequestBuilder) {
    return i079fc82ebc66b91d2f1e5db2f46e14a93047f055730594a3c555217f5240b7d9.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i74ccd3f22fbacb438df996a11eab6efcecf784f34164f51b4e63d80efbf60810.MultiValueExtendedPropertiesRequestBuilder) {
    return i74ccd3f22fbacb438df996a11eab6efcecf784f34164f51b4e63d80efbf60810.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i554a181b0fba8aaa263c329ec9f30c8c0170b08d3f8ec9364659a4b1ce07ed8f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i554a181b0fba8aaa263c329ec9f30c8c0170b08d3f8ec9364659a4b1ce07ed8f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i39b113e521727af9278341f568dc06471a9cdc7b19e99eedae87afbe5bcacb40.SingleValueExtendedPropertiesRequestBuilder) {
    return i39b113e521727af9278341f568dc06471a9cdc7b19e99eedae87afbe5bcacb40.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7dfeae46798520fb15d6b541a64a5451c69c35f2a0deed35ba1e83b44eac510b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7dfeae46798520fb15d6b541a64a5451c69c35f2a0deed35ba1e83b44eac510b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i145ee4ad288f28c56f916c936201460b07ba8171fb97b6049f9c5b75da266ad2.SnoozeReminderRequestBuilder) {
    return i145ee4ad288f28c56f916c936201460b07ba8171fb97b6049f9c5b75da266ad2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i677a95ac0890a8b4be07a112c61b91c390467b7f02a7cb3c8ad0e5f612695e9e.TentativelyAcceptRequestBuilder) {
    return i677a95ac0890a8b4be07a112c61b91c390467b7f02a7cb3c8ad0e5f612695e9e.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
