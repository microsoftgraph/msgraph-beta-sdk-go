package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i047483f3293527a2951c3b9aa722b408545abf89aea7001af043e0cb335075dc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/cancel"
    i345323339c081b0a5241285458ff2263b8e2113c0b58ea4037495dec94dfdfaa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/calendar"
    i3b0a7b5ab74aecbc2451b3b5bff84142d7c39867bcbcd4207f092ee35ab67f20 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i6088291a59c3aa28168a0aaf187697e5cf5a5fdb0c1f777656538b3f3dede1fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i66f1d3634ab07ffa0b694ce0a2846d488ac96d489a32da96323cc5358cc39695 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i6d08d5a769e86f6028fa1c0b7f33d6fabbe65d340e7537b523c142e6e484a70c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/extensions"
    i812a8745ebf7e8e47b31a7c174ccccebaa6729c9cef5418842c1a8fa8979edc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/attachments"
    i9b0deec2e1d45b5c43051a5e623fcaacaa4bc041828eed3986201e6f690cae11 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/accept"
    ia838cbc3704946216c18dae79363ebfb060d2310672b278306e308387e63759c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/decline"
    ib1ccff1ee35b084a3f28d166816e3ac7f6f372597d83b2bd8ec29f675497a63c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    idec326f2c13044c41cfe2eceab82347df831418f26e77db7e75721f7fd306cc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    iffe7c4934d5612b4c1a2c4c8ec1387a646dc7ca75c212be141440ec0553d0198 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/forward"
    i495412143c49c42452e355b622a455da29bdeaf347e9ea928b681ecaee1b4c02 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i4dcbaea77fa9d6c80eb0e03d5c2a5784dfadd67eaca1fcecacf8628623003724 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i5b64856dd62f57a3d6a5d64e1adbc5c631c38704115c6397c04e89684d4f003e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    iccd5b2db92df6c0ee013280ca71b87bfb94be33bd9953935999f4b9cf5ada7ee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item/exceptionoccurrences/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
}
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i9b0deec2e1d45b5c43051a5e623fcaacaa4bc041828eed3986201e6f690cae11.AcceptRequestBuilder) {
    return i9b0deec2e1d45b5c43051a5e623fcaacaa4bc041828eed3986201e6f690cae11.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i812a8745ebf7e8e47b31a7c174ccccebaa6729c9cef5418842c1a8fa8979edc0.AttachmentsRequestBuilder) {
    return i812a8745ebf7e8e47b31a7c174ccccebaa6729c9cef5418842c1a8fa8979edc0.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*iccd5b2db92df6c0ee013280ca71b87bfb94be33bd9953935999f4b9cf5ada7ee.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return iccd5b2db92df6c0ee013280ca71b87bfb94be33bd9953935999f4b9cf5ada7ee.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i345323339c081b0a5241285458ff2263b8e2113c0b58ea4037495dec94dfdfaa.CalendarRequestBuilder) {
    return i345323339c081b0a5241285458ff2263b8e2113c0b58ea4037495dec94dfdfaa.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i047483f3293527a2951c3b9aa722b408545abf89aea7001af043e0cb335075dc.CancelRequestBuilder) {
    return i047483f3293527a2951c3b9aa722b408545abf89aea7001af043e0cb335075dc.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from users
func (m *EventItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*ia838cbc3704946216c18dae79363ebfb060d2310672b278306e308387e63759c.DeclineRequestBuilder) {
    return ia838cbc3704946216c18dae79363ebfb060d2310672b278306e308387e63759c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*ib1ccff1ee35b084a3f28d166816e3ac7f6f372597d83b2bd8ec29f675497a63c.DismissReminderRequestBuilder) {
    return ib1ccff1ee35b084a3f28d166816e3ac7f6f372597d83b2bd8ec29f675497a63c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i6d08d5a769e86f6028fa1c0b7f33d6fabbe65d340e7537b523c142e6e484a70c.ExtensionsRequestBuilder) {
    return i6d08d5a769e86f6028fa1c0b7f33d6fabbe65d340e7537b523c142e6e484a70c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i5b64856dd62f57a3d6a5d64e1adbc5c631c38704115c6397c04e89684d4f003e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i5b64856dd62f57a3d6a5d64e1adbc5c631c38704115c6397c04e89684d4f003e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*iffe7c4934d5612b4c1a2c4c8ec1387a646dc7ca75c212be141440ec0553d0198.ForwardRequestBuilder) {
    return iffe7c4934d5612b4c1a2c4c8ec1387a646dc7ca75c212be141440ec0553d0198.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get exceptionOccurrences from users
func (m *EventItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i6088291a59c3aa28168a0aaf187697e5cf5a5fdb0c1f777656538b3f3dede1fc.MultiValueExtendedPropertiesRequestBuilder) {
    return i6088291a59c3aa28168a0aaf187697e5cf5a5fdb0c1f777656538b3f3dede1fc.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i495412143c49c42452e355b622a455da29bdeaf347e9ea928b681ecaee1b4c02.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i495412143c49c42452e355b622a455da29bdeaf347e9ea928b681ecaee1b4c02.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i66f1d3634ab07ffa0b694ce0a2846d488ac96d489a32da96323cc5358cc39695.SingleValueExtendedPropertiesRequestBuilder) {
    return i66f1d3634ab07ffa0b694ce0a2846d488ac96d489a32da96323cc5358cc39695.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4dcbaea77fa9d6c80eb0e03d5c2a5784dfadd67eaca1fcecacf8628623003724.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i4dcbaea77fa9d6c80eb0e03d5c2a5784dfadd67eaca1fcecacf8628623003724.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*idec326f2c13044c41cfe2eceab82347df831418f26e77db7e75721f7fd306cc5.SnoozeReminderRequestBuilder) {
    return idec326f2c13044c41cfe2eceab82347df831418f26e77db7e75721f7fd306cc5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3b0a7b5ab74aecbc2451b3b5bff84142d7c39867bcbcd4207f092ee35ab67f20.TentativelyAcceptRequestBuilder) {
    return i3b0a7b5ab74aecbc2451b3b5bff84142d7c39867bcbcd4207f092ee35ab67f20.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
