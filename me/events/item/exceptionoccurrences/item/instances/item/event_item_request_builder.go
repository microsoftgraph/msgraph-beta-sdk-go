package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i147ca5bd9a5aae750aed77db12eae5e11e85e45b4093e66627295e6bd521e988 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/extensions"
    i28b9f46301313545f25b9d8f34435f2d1ee8d32f1a60d6ec9e55702ca985fec3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/calendar"
    i2eac05f2e95375fed0a145b4bf25402689fb36da05cdba3c8b019545d24880cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/forward"
    i4acc795b5eba0be575ce078037dae652c46c68043009976d20a7f74d28612df2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/attachments"
    i5924364d8cf986cfd5824f048ad865f0a58573a2c1c0cae08615d1453a8c8eda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/accept"
    i5d3ed573b41ca382216556f70e4248b07375d5c1b5d9c6f451142c49f7288d47 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/cancel"
    i6601f78c7843912d177aac249910ea39a5ab2645a472967f6048ee09e6561cce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i784ddad8f6113e1afe9958fd066baf595bffb7099bd5c2226f47e687c6e525bb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i80bddbd6354676a0f38c89b600b454da1efa77e47dd8460cc668f25335006b0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i84e6fc913cb1afe4bc42026decabeae4a344402eb709b3016ab69bfaef118bec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i8b1d1134181c3b38646dd9b746f684444a5b3dbee03718862ef65415ef341631 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i96c095c6229d8e8c75d31b5ae00216a48c8d7d2b233dd3cd40bcb61255105ad8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/decline"
    i1110646ae74c4d176a75bc9dfadda7baad3827d18a90f41495c6e9228c3ab04d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    i3b64b4f523bf181b6794bc57e9fda58e4864997acb7170cad2d07aa56833e459 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i4b215a6ef9a56911a44ea7fdbbaba2bf0b2ac89631d6bdb86aeca0da455f873a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i9fb2bb9d83f127b66d4c17a08138e8d3b945089862524f0be88de8189eb92df4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i5924364d8cf986cfd5824f048ad865f0a58573a2c1c0cae08615d1453a8c8eda.AcceptRequestBuilder) {
    return i5924364d8cf986cfd5824f048ad865f0a58573a2c1c0cae08615d1453a8c8eda.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i4acc795b5eba0be575ce078037dae652c46c68043009976d20a7f74d28612df2.AttachmentsRequestBuilder) {
    return i4acc795b5eba0be575ce078037dae652c46c68043009976d20a7f74d28612df2.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i3b64b4f523bf181b6794bc57e9fda58e4864997acb7170cad2d07aa56833e459.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i3b64b4f523bf181b6794bc57e9fda58e4864997acb7170cad2d07aa56833e459.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i28b9f46301313545f25b9d8f34435f2d1ee8d32f1a60d6ec9e55702ca985fec3.CalendarRequestBuilder) {
    return i28b9f46301313545f25b9d8f34435f2d1ee8d32f1a60d6ec9e55702ca985fec3.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i5d3ed573b41ca382216556f70e4248b07375d5c1b5d9c6f451142c49f7288d47.CancelRequestBuilder) {
    return i5d3ed573b41ca382216556f70e4248b07375d5c1b5d9c6f451142c49f7288d47.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property instances for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instances in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i96c095c6229d8e8c75d31b5ae00216a48c8d7d2b233dd3cd40bcb61255105ad8.DeclineRequestBuilder) {
    return i96c095c6229d8e8c75d31b5ae00216a48c8d7d2b233dd3cd40bcb61255105ad8.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i784ddad8f6113e1afe9958fd066baf595bffb7099bd5c2226f47e687c6e525bb.DismissReminderRequestBuilder) {
    return i784ddad8f6113e1afe9958fd066baf595bffb7099bd5c2226f47e687c6e525bb.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i147ca5bd9a5aae750aed77db12eae5e11e85e45b4093e66627295e6bd521e988.ExtensionsRequestBuilder) {
    return i147ca5bd9a5aae750aed77db12eae5e11e85e45b4093e66627295e6bd521e988.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i1110646ae74c4d176a75bc9dfadda7baad3827d18a90f41495c6e9228c3ab04d.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i1110646ae74c4d176a75bc9dfadda7baad3827d18a90f41495c6e9228c3ab04d.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i2eac05f2e95375fed0a145b4bf25402689fb36da05cdba3c8b019545d24880cb.ForwardRequestBuilder) {
    return i2eac05f2e95375fed0a145b4bf25402689fb36da05cdba3c8b019545d24880cb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i80bddbd6354676a0f38c89b600b454da1efa77e47dd8460cc668f25335006b0c.MultiValueExtendedPropertiesRequestBuilder) {
    return i80bddbd6354676a0f38c89b600b454da1efa77e47dd8460cc668f25335006b0c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9fb2bb9d83f127b66d4c17a08138e8d3b945089862524f0be88de8189eb92df4.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i9fb2bb9d83f127b66d4c17a08138e8d3b945089862524f0be88de8189eb92df4.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6601f78c7843912d177aac249910ea39a5ab2645a472967f6048ee09e6561cce.SingleValueExtendedPropertiesRequestBuilder) {
    return i6601f78c7843912d177aac249910ea39a5ab2645a472967f6048ee09e6561cce.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4b215a6ef9a56911a44ea7fdbbaba2bf0b2ac89631d6bdb86aeca0da455f873a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i4b215a6ef9a56911a44ea7fdbbaba2bf0b2ac89631d6bdb86aeca0da455f873a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i8b1d1134181c3b38646dd9b746f684444a5b3dbee03718862ef65415ef341631.SnoozeReminderRequestBuilder) {
    return i8b1d1134181c3b38646dd9b746f684444a5b3dbee03718862ef65415ef341631.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i84e6fc913cb1afe4bc42026decabeae4a344402eb709b3016ab69bfaef118bec.TentativelyAcceptRequestBuilder) {
    return i84e6fc913cb1afe4bc42026decabeae4a344402eb709b3016ab69bfaef118bec.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
