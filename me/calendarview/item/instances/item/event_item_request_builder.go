package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0604ac7d512df3b4eb96ae20cd98c288d9968e4ab07d7c0fe8a5c4be771297ca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/attachments"
    i189328c79d22236d9b92041360d53f1928257dbf143e22032d480030b14a50cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/tentativelyaccept"
    i22abce52ee562d957626d5f4328048cfe6c6870984265dd6ad55fed38f73b7ce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences"
    i242657b864a63c52e28f8ba6c2a6570723e67ca53a4cd1134a38829e729af30b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/forward"
    i26b773c466ffee56e8f9756889040753ff9526053bf8c72669976815deacd01f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/accept"
    i389c4160f712471be558936e65480058e8cb908fa6cc6b4ef1c99f8515475c3d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/multivalueextendedproperties"
    i3d5cf2470c4628035ec05b471318fbf876f6c321e41ae2e24742ab2db6236e4c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/dismissreminder"
    i61e1b50424c7b8596bff365635e22852af40b33f0f5da423b57f9dedc525a403 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/extensions"
    i6bf22a3ab7bf110327b17625ea3863d6482afb72c57887f3e2d590213d8b7798 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/decline"
    i82452a1a1495cabdf9de5f574f4dc3e2dfaa9836a0ea9aef7d3537ac46543fa4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/cancel"
    i9975a1bc9a2653c445caa7c1e6bbdef52618f11657dd0027849c2cad3411cc00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/snoozereminder"
    ie18ef437e4857f158d5d172faee6ae086c39a1d31faf79f00efaf43d64fd48f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/calendar"
    iedabb476d940cc60507c1504cb4d8d8b89e61e6f11ee1a607bdf8fbc9dd59715 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/singlevalueextendedproperties"
    i3031a482b76c5a7a7bf88758979a58c92e98bce384a98705bce0f87ed0b1110c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/extensions/item"
    i36924185f9e82035b48655f56a14bf0431b2facfafe06788e519970590edbefe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/attachments/item"
    i3d54f93e91d5ed6654e10959b0f1a04fbae75f6f2bc65d90ac87dcbb210b2420 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/multivalueextendedproperties/item"
    if9e51c796cc6289089cb878d40f0a4958e538ab5eb849aaafa95b8492c41740b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/singlevalueextendedproperties/item"
    iff47ad4dcdf3bba9885779bc97cc39f17995a28116daf285208c79ff00f82691 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*i26b773c466ffee56e8f9756889040753ff9526053bf8c72669976815deacd01f.AcceptRequestBuilder) {
    return i26b773c466ffee56e8f9756889040753ff9526053bf8c72669976815deacd01f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i0604ac7d512df3b4eb96ae20cd98c288d9968e4ab07d7c0fe8a5c4be771297ca.AttachmentsRequestBuilder) {
    return i0604ac7d512df3b4eb96ae20cd98c288d9968e4ab07d7c0fe8a5c4be771297ca.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i36924185f9e82035b48655f56a14bf0431b2facfafe06788e519970590edbefe.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i36924185f9e82035b48655f56a14bf0431b2facfafe06788e519970590edbefe.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ie18ef437e4857f158d5d172faee6ae086c39a1d31faf79f00efaf43d64fd48f5.CalendarRequestBuilder) {
    return ie18ef437e4857f158d5d172faee6ae086c39a1d31faf79f00efaf43d64fd48f5.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i82452a1a1495cabdf9de5f574f4dc3e2dfaa9836a0ea9aef7d3537ac46543fa4.CancelRequestBuilder) {
    return i82452a1a1495cabdf9de5f574f4dc3e2dfaa9836a0ea9aef7d3537ac46543fa4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for me
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
// CreatePatchRequestInformation update the navigation property instances in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in me
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
func (m *EventItemRequestBuilder) Decline()(*i6bf22a3ab7bf110327b17625ea3863d6482afb72c57887f3e2d590213d8b7798.DeclineRequestBuilder) {
    return i6bf22a3ab7bf110327b17625ea3863d6482afb72c57887f3e2d590213d8b7798.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property instances for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i3d5cf2470c4628035ec05b471318fbf876f6c321e41ae2e24742ab2db6236e4c.DismissReminderRequestBuilder) {
    return i3d5cf2470c4628035ec05b471318fbf876f6c321e41ae2e24742ab2db6236e4c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i22abce52ee562d957626d5f4328048cfe6c6870984265dd6ad55fed38f73b7ce.ExceptionOccurrencesRequestBuilder) {
    return i22abce52ee562d957626d5f4328048cfe6c6870984265dd6ad55fed38f73b7ce.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*iff47ad4dcdf3bba9885779bc97cc39f17995a28116daf285208c79ff00f82691.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return iff47ad4dcdf3bba9885779bc97cc39f17995a28116daf285208c79ff00f82691.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i61e1b50424c7b8596bff365635e22852af40b33f0f5da423b57f9dedc525a403.ExtensionsRequestBuilder) {
    return i61e1b50424c7b8596bff365635e22852af40b33f0f5da423b57f9dedc525a403.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i3031a482b76c5a7a7bf88758979a58c92e98bce384a98705bce0f87ed0b1110c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i3031a482b76c5a7a7bf88758979a58c92e98bce384a98705bce0f87ed0b1110c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i242657b864a63c52e28f8ba6c2a6570723e67ca53a4cd1134a38829e729af30b.ForwardRequestBuilder) {
    return i242657b864a63c52e28f8ba6c2a6570723e67ca53a4cd1134a38829e729af30b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i389c4160f712471be558936e65480058e8cb908fa6cc6b4ef1c99f8515475c3d.MultiValueExtendedPropertiesRequestBuilder) {
    return i389c4160f712471be558936e65480058e8cb908fa6cc6b4ef1c99f8515475c3d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i3d54f93e91d5ed6654e10959b0f1a04fbae75f6f2bc65d90ac87dcbb210b2420.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i3d54f93e91d5ed6654e10959b0f1a04fbae75f6f2bc65d90ac87dcbb210b2420.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property instances in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*iedabb476d940cc60507c1504cb4d8d8b89e61e6f11ee1a607bdf8fbc9dd59715.SingleValueExtendedPropertiesRequestBuilder) {
    return iedabb476d940cc60507c1504cb4d8d8b89e61e6f11ee1a607bdf8fbc9dd59715.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if9e51c796cc6289089cb878d40f0a4958e538ab5eb849aaafa95b8492c41740b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return if9e51c796cc6289089cb878d40f0a4958e538ab5eb849aaafa95b8492c41740b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i9975a1bc9a2653c445caa7c1e6bbdef52618f11657dd0027849c2cad3411cc00.SnoozeReminderRequestBuilder) {
    return i9975a1bc9a2653c445caa7c1e6bbdef52618f11657dd0027849c2cad3411cc00.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i189328c79d22236d9b92041360d53f1928257dbf143e22032d480030b14a50cb.TentativelyAcceptRequestBuilder) {
    return i189328c79d22236d9b92041360d53f1928257dbf143e22032d480030b14a50cb.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
