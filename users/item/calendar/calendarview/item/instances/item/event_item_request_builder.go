package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i020b3e967fcc436548d613398edcfa37d8cd604ea60cbdd79b8a97c6a5f11b6f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/calendar"
    i136f1a7ad510e8e001985daa897b2a78a1914cee3f52b15549a1f0b4424e0cf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/forward"
    i260edc1ba1a1f3fa3af0880409accb18c06e9725182ebd32c4cf25d2150d9267 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/decline"
    i2aa9b2f5b4f50fd61e144342ebb89287b492d8e9c976ffe9230efe948cf4757d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/multivalueextendedproperties"
    i2b16cb92a36dba3cc55cbf667fbf35227381b9da75c0515c9563076088882e83 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/accept"
    i4c0063fc77cf077ec34f2c230b5148f750d0f146e59b424dc135bb93348af1b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/snoozereminder"
    i70c0fe78c2812f618e8b9989b5c951957a3a9e445ea6edde566406d7709191ec "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/extensions"
    i7b284eeb1cf95e6a769ebcdea6d679260bef93e3d7e013aae029d953b3c8d304 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/cancel"
    i86ef006b8fbe2c6aa4646ed5da5ffff77f989cc21a8bae94e82832cad17edbf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/tentativelyaccept"
    i9619f56b877cd8640053989b275eb5be3bcf2415fdecdccd842b2ec5d2eed4cd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/attachments"
    ia94f3fd847f66d103d5769b324cc71c32465960b44bb1e33dd61f55477012b10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/dismissreminder"
    if5ef4784bec4d20e8afe960ab3e144e629291bf3d4fd3269060fce9125b3ccc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/singlevalueextendedproperties"
    ifd92068e93bd6f7d74e2537ebe39d44f31f925d12d10c557ebc783c77b09cc76 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences"
    i3ea001ffee843c59e46463559b32e5645be9217c3aa308786764738be19fd804 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/singlevalueextendedproperties/item"
    i4a351f8736d296c5398c04ce13de45771f4c5665fb79a4e6663ea05a999e22ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/attachments/item"
    i62b4d40fd53de31f3077d7bc77cb9b85c73f78bca34235dc7bc96d2fea5851c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item"
    i971bdb194a77b4ec456ecc0c0122edbbb8106d4a0188dfd9743f485fd4224cff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/multivalueextendedproperties/item"
    idaf785b11968a1f16125a204196482f45cd007b84a4381c61503e486ff710960 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*i2b16cb92a36dba3cc55cbf667fbf35227381b9da75c0515c9563076088882e83.AcceptRequestBuilder) {
    return i2b16cb92a36dba3cc55cbf667fbf35227381b9da75c0515c9563076088882e83.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i9619f56b877cd8640053989b275eb5be3bcf2415fdecdccd842b2ec5d2eed4cd.AttachmentsRequestBuilder) {
    return i9619f56b877cd8640053989b275eb5be3bcf2415fdecdccd842b2ec5d2eed4cd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i4a351f8736d296c5398c04ce13de45771f4c5665fb79a4e6663ea05a999e22ff.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i4a351f8736d296c5398c04ce13de45771f4c5665fb79a4e6663ea05a999e22ff.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i020b3e967fcc436548d613398edcfa37d8cd604ea60cbdd79b8a97c6a5f11b6f.CalendarRequestBuilder) {
    return i020b3e967fcc436548d613398edcfa37d8cd604ea60cbdd79b8a97c6a5f11b6f.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i7b284eeb1cf95e6a769ebcdea6d679260bef93e3d7e013aae029d953b3c8d304.CancelRequestBuilder) {
    return i7b284eeb1cf95e6a769ebcdea6d679260bef93e3d7e013aae029d953b3c8d304.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i260edc1ba1a1f3fa3af0880409accb18c06e9725182ebd32c4cf25d2150d9267.DeclineRequestBuilder) {
    return i260edc1ba1a1f3fa3af0880409accb18c06e9725182ebd32c4cf25d2150d9267.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property instances for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*ia94f3fd847f66d103d5769b324cc71c32465960b44bb1e33dd61f55477012b10.DismissReminderRequestBuilder) {
    return ia94f3fd847f66d103d5769b324cc71c32465960b44bb1e33dd61f55477012b10.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ifd92068e93bd6f7d74e2537ebe39d44f31f925d12d10c557ebc783c77b09cc76.ExceptionOccurrencesRequestBuilder) {
    return ifd92068e93bd6f7d74e2537ebe39d44f31f925d12d10c557ebc783c77b09cc76.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i62b4d40fd53de31f3077d7bc77cb9b85c73f78bca34235dc7bc96d2fea5851c9.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i62b4d40fd53de31f3077d7bc77cb9b85c73f78bca34235dc7bc96d2fea5851c9.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i70c0fe78c2812f618e8b9989b5c951957a3a9e445ea6edde566406d7709191ec.ExtensionsRequestBuilder) {
    return i70c0fe78c2812f618e8b9989b5c951957a3a9e445ea6edde566406d7709191ec.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*idaf785b11968a1f16125a204196482f45cd007b84a4381c61503e486ff710960.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return idaf785b11968a1f16125a204196482f45cd007b84a4381c61503e486ff710960.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i136f1a7ad510e8e001985daa897b2a78a1914cee3f52b15549a1f0b4424e0cf4.ForwardRequestBuilder) {
    return i136f1a7ad510e8e001985daa897b2a78a1914cee3f52b15549a1f0b4424e0cf4.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i2aa9b2f5b4f50fd61e144342ebb89287b492d8e9c976ffe9230efe948cf4757d.MultiValueExtendedPropertiesRequestBuilder) {
    return i2aa9b2f5b4f50fd61e144342ebb89287b492d8e9c976ffe9230efe948cf4757d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i971bdb194a77b4ec456ecc0c0122edbbb8106d4a0188dfd9743f485fd4224cff.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i971bdb194a77b4ec456ecc0c0122edbbb8106d4a0188dfd9743f485fd4224cff.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property instances in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if5ef4784bec4d20e8afe960ab3e144e629291bf3d4fd3269060fce9125b3ccc0.SingleValueExtendedPropertiesRequestBuilder) {
    return if5ef4784bec4d20e8afe960ab3e144e629291bf3d4fd3269060fce9125b3ccc0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3ea001ffee843c59e46463559b32e5645be9217c3aa308786764738be19fd804.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i3ea001ffee843c59e46463559b32e5645be9217c3aa308786764738be19fd804.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i4c0063fc77cf077ec34f2c230b5148f750d0f146e59b424dc135bb93348af1b6.SnoozeReminderRequestBuilder) {
    return i4c0063fc77cf077ec34f2c230b5148f750d0f146e59b424dc135bb93348af1b6.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i86ef006b8fbe2c6aa4646ed5da5ffff77f989cc21a8bae94e82832cad17edbf6.TentativelyAcceptRequestBuilder) {
    return i86ef006b8fbe2c6aa4646ed5da5ffff77f989cc21a8bae94e82832cad17edbf6.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
