package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07670fc2b5ad7c38ca83d3d7eb784dd93260f9ccf700f64d07b9dc665cd1b103 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/extensions"
    i0f2be0567981d3d7b15fec27c27723680a229b3e7e4b9fed06f2607c07770ad9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/singlevalueextendedproperties"
    i11c912cdfbd103a89a06ba951d85b3c9d1190e38c8b3bb179ac92e5125182bdc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/forward"
    i182dff33927f9786e027071b857a1a96a045caed7e85102c1485a8dc7a6e7da6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/decline"
    i21697c1216cd0ef5deb2fe67c016f7d746dbbbc7261df5b764553bfba26c1534 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/calendar"
    i33d96815c909a72df38e4c4f32fa1cfbe727a2b082d161f29271aa585526eec3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/dismissreminder"
    i5bd86f7c92cb52bb6b263420e8fb37db6aa04e2e0dc3e75357c0a902b70ba2c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/snoozereminder"
    iaea2d91b6ae4abdb34fce9c73ce7fc298668b470ff75f7c6be768373d4ceb7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/accept"
    ib476f24c9486b87ddfa3f348795ef1933f1dd919371bcf3ac5250911b6bb6aab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/cancel"
    iba1109b87b7945a7efdd28f97e737eec8eda0d4fb61056727e18e23a152ceead "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences"
    ibaa894fa9c6a960e2a22b240311826e310a835659caafe52952f5b6101ca7577 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/tentativelyaccept"
    ibe7786e98ed0784850c47015e64ba32d8eebfbe0e81f5ce5f0b792bb7a4e7a73 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/attachments"
    id845c6766f509138b8b505270cfd0134b1ebb08f71c52755c46cd534eafbf8af "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/multivalueextendedproperties"
    i0aee271214b6a56cc7b5056c01fc075fc528ef43643695282c7fc53c4e64c78b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/singlevalueextendedproperties/item"
    i8cbd4cb6c8fbebc086df152ebf27c1a10ba7eb596b0b0470e1eaaab73c1af610 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item"
    ibc9fe37d5f6e899679dca4270239fc80a74b20fb17745d1ee1beda217d4645b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/attachments/item"
    ie4df7fdf56671232abf33881312e91277f16b7f276bee14cb78a58a8550d808e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/extensions/item"
    ieac1c8fac973b18dba018d942d20070a8b080c6ad207ac4af3c77cf66f0e440d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*iaea2d91b6ae4abdb34fce9c73ce7fc298668b470ff75f7c6be768373d4ceb7bc.AcceptRequestBuilder) {
    return iaea2d91b6ae4abdb34fce9c73ce7fc298668b470ff75f7c6be768373d4ceb7bc.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ibe7786e98ed0784850c47015e64ba32d8eebfbe0e81f5ce5f0b792bb7a4e7a73.AttachmentsRequestBuilder) {
    return ibe7786e98ed0784850c47015e64ba32d8eebfbe0e81f5ce5f0b792bb7a4e7a73.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ibc9fe37d5f6e899679dca4270239fc80a74b20fb17745d1ee1beda217d4645b9.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ibc9fe37d5f6e899679dca4270239fc80a74b20fb17745d1ee1beda217d4645b9.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i21697c1216cd0ef5deb2fe67c016f7d746dbbbc7261df5b764553bfba26c1534.CalendarRequestBuilder) {
    return i21697c1216cd0ef5deb2fe67c016f7d746dbbbc7261df5b764553bfba26c1534.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ib476f24c9486b87ddfa3f348795ef1933f1dd919371bcf3ac5250911b6bb6aab.CancelRequestBuilder) {
    return ib476f24c9486b87ddfa3f348795ef1933f1dd919371bcf3ac5250911b6bb6aab.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i182dff33927f9786e027071b857a1a96a045caed7e85102c1485a8dc7a6e7da6.DeclineRequestBuilder) {
    return i182dff33927f9786e027071b857a1a96a045caed7e85102c1485a8dc7a6e7da6.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i33d96815c909a72df38e4c4f32fa1cfbe727a2b082d161f29271aa585526eec3.DismissReminderRequestBuilder) {
    return i33d96815c909a72df38e4c4f32fa1cfbe727a2b082d161f29271aa585526eec3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*iba1109b87b7945a7efdd28f97e737eec8eda0d4fb61056727e18e23a152ceead.ExceptionOccurrencesRequestBuilder) {
    return iba1109b87b7945a7efdd28f97e737eec8eda0d4fb61056727e18e23a152ceead.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i8cbd4cb6c8fbebc086df152ebf27c1a10ba7eb596b0b0470e1eaaab73c1af610.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i8cbd4cb6c8fbebc086df152ebf27c1a10ba7eb596b0b0470e1eaaab73c1af610.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i07670fc2b5ad7c38ca83d3d7eb784dd93260f9ccf700f64d07b9dc665cd1b103.ExtensionsRequestBuilder) {
    return i07670fc2b5ad7c38ca83d3d7eb784dd93260f9ccf700f64d07b9dc665cd1b103.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ie4df7fdf56671232abf33881312e91277f16b7f276bee14cb78a58a8550d808e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ie4df7fdf56671232abf33881312e91277f16b7f276bee14cb78a58a8550d808e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i11c912cdfbd103a89a06ba951d85b3c9d1190e38c8b3bb179ac92e5125182bdc.ForwardRequestBuilder) {
    return i11c912cdfbd103a89a06ba951d85b3c9d1190e38c8b3bb179ac92e5125182bdc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*id845c6766f509138b8b505270cfd0134b1ebb08f71c52755c46cd534eafbf8af.MultiValueExtendedPropertiesRequestBuilder) {
    return id845c6766f509138b8b505270cfd0134b1ebb08f71c52755c46cd534eafbf8af.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ieac1c8fac973b18dba018d942d20070a8b080c6ad207ac4af3c77cf66f0e440d.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ieac1c8fac973b18dba018d942d20070a8b080c6ad207ac4af3c77cf66f0e440d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i0f2be0567981d3d7b15fec27c27723680a229b3e7e4b9fed06f2607c07770ad9.SingleValueExtendedPropertiesRequestBuilder) {
    return i0f2be0567981d3d7b15fec27c27723680a229b3e7e4b9fed06f2607c07770ad9.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0aee271214b6a56cc7b5056c01fc075fc528ef43643695282c7fc53c4e64c78b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i0aee271214b6a56cc7b5056c01fc075fc528ef43643695282c7fc53c4e64c78b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5bd86f7c92cb52bb6b263420e8fb37db6aa04e2e0dc3e75357c0a902b70ba2c8.SnoozeReminderRequestBuilder) {
    return i5bd86f7c92cb52bb6b263420e8fb37db6aa04e2e0dc3e75357c0a902b70ba2c8.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibaa894fa9c6a960e2a22b240311826e310a835659caafe52952f5b6101ca7577.TentativelyAcceptRequestBuilder) {
    return ibaa894fa9c6a960e2a22b240311826e310a835659caafe52952f5b6101ca7577.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
