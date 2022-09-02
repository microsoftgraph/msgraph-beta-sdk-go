package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2fec313c1ff2613627d0501b33555a647f0c1eb2583ea93ca8066c00911e8bbb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i38ba030b1102af0b177262034ef1b13eef1530342e3e935c23d3b3208938a978 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    i59ac7aabe390ca7d5056b690062e1e7f86f719e516f7ca8b971d3814a45939d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    i5b28e94838075cfd25ebe7535803e65d7a9c62a50168473edce73f262717a9a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i5b68e030271873619cabeef13c554c71e175326aac3dc5ad029b28cafd3d3267 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i619bec8debf98977be826c35065ee19ce9a45fd8440ac956b1eb8404550bb424 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    i6d7a4c0c9534c743ca50c52607aaf4213a43b95bb59a7313b0ce1c831769dead "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i85899973653da5aede87150e4dfac54d76a3d4cd360aaaf614b8bc4faf6fbe43 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    icbc1ea2d51290bca281c36180b54412b79ab1b98f05a2e14c4370603756a8175 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    icd1c71bd8000e542dfc930d20696b18d64606f69f7078af0836aa02af3537a7f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    id61003e840d4e7a57641908b95c7f8af8af4531d39b6425a1775037e9ff33659 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    ifc712aab74bb05939c55d02b980d74cb013865fd713a3877b9c731e642eab273 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    i6c1c62f167b45cb8719fc69afc60f0e0a8c813e8f8561199f45e7e645b5ac052 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    ia301ab77319b2410c4f9c7db160699123feaf1a79b09ec12cee63a1c176b47c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    ia415d4ec84429cdb69128a2c8b485787932a7056346abf6b678fd55d55e67ace "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    ia4c3bec308a0d54074746650aa890128e8fcc0e5b7428aa31013466358350db1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
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
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
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
func (m *EventItemRequestBuilder) Accept()(*i38ba030b1102af0b177262034ef1b13eef1530342e3e935c23d3b3208938a978.AcceptRequestBuilder) {
    return i38ba030b1102af0b177262034ef1b13eef1530342e3e935c23d3b3208938a978.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*icd1c71bd8000e542dfc930d20696b18d64606f69f7078af0836aa02af3537a7f.AttachmentsRequestBuilder) {
    return icd1c71bd8000e542dfc930d20696b18d64606f69f7078af0836aa02af3537a7f.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia4c3bec308a0d54074746650aa890128e8fcc0e5b7428aa31013466358350db1.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ia4c3bec308a0d54074746650aa890128e8fcc0e5b7428aa31013466358350db1.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ifc712aab74bb05939c55d02b980d74cb013865fd713a3877b9c731e642eab273.CalendarRequestBuilder) {
    return ifc712aab74bb05939c55d02b980d74cb013865fd713a3877b9c731e642eab273.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*icbc1ea2d51290bca281c36180b54412b79ab1b98f05a2e14c4370603756a8175.CancelRequestBuilder) {
    return icbc1ea2d51290bca281c36180b54412b79ab1b98f05a2e14c4370603756a8175.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property instances for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for users
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property instances in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in users
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
func (m *EventItemRequestBuilder) Decline()(*i85899973653da5aede87150e4dfac54d76a3d4cd360aaaf614b8bc4faf6fbe43.DeclineRequestBuilder) {
    return i85899973653da5aede87150e4dfac54d76a3d4cd360aaaf614b8bc4faf6fbe43.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i6d7a4c0c9534c743ca50c52607aaf4213a43b95bb59a7313b0ce1c831769dead.DismissReminderRequestBuilder) {
    return i6d7a4c0c9534c743ca50c52607aaf4213a43b95bb59a7313b0ce1c831769dead.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i619bec8debf98977be826c35065ee19ce9a45fd8440ac956b1eb8404550bb424.ExtensionsRequestBuilder) {
    return i619bec8debf98977be826c35065ee19ce9a45fd8440ac956b1eb8404550bb424.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia301ab77319b2410c4f9c7db160699123feaf1a79b09ec12cee63a1c176b47c8.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ia301ab77319b2410c4f9c7db160699123feaf1a79b09ec12cee63a1c176b47c8.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i59ac7aabe390ca7d5056b690062e1e7f86f719e516f7ca8b971d3814a45939d1.ForwardRequestBuilder) {
    return i59ac7aabe390ca7d5056b690062e1e7f86f719e516f7ca8b971d3814a45939d1.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i5b68e030271873619cabeef13c554c71e175326aac3dc5ad029b28cafd3d3267.MultiValueExtendedPropertiesRequestBuilder) {
    return i5b68e030271873619cabeef13c554c71e175326aac3dc5ad029b28cafd3d3267.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ia415d4ec84429cdb69128a2c8b485787932a7056346abf6b678fd55d55e67ace.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ia415d4ec84429cdb69128a2c8b485787932a7056346abf6b678fd55d55e67ace.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*id61003e840d4e7a57641908b95c7f8af8af4531d39b6425a1775037e9ff33659.SingleValueExtendedPropertiesRequestBuilder) {
    return id61003e840d4e7a57641908b95c7f8af8af4531d39b6425a1775037e9ff33659.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i6c1c62f167b45cb8719fc69afc60f0e0a8c813e8f8561199f45e7e645b5ac052.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i6c1c62f167b45cb8719fc69afc60f0e0a8c813e8f8561199f45e7e645b5ac052.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5b28e94838075cfd25ebe7535803e65d7a9c62a50168473edce73f262717a9a9.SnoozeReminderRequestBuilder) {
    return i5b28e94838075cfd25ebe7535803e65d7a9c62a50168473edce73f262717a9a9.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i2fec313c1ff2613627d0501b33555a647f0c1eb2583ea93ca8066c00911e8bbb.TentativelyAcceptRequestBuilder) {
    return i2fec313c1ff2613627d0501b33555a647f0c1eb2583ea93ca8066c00911e8bbb.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
