package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4dee44f6ca6933f0d202e222b4a177b72d807831149a9b155943116a9fc6c608 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i58d26e13a36f9ef40d2af09c4e8691b48132d0c9426cbf53d095c5170731109d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    i6e665073402a7c287f185798d81d13998809adcd52296fa21e8872685d14f0f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    i746ca6978b7e4e298f4b3a396747dc93eabc3b1af8630010cdfef674a8b2ba5a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    i8591b3851ec293b18b7b25a8487164392ad4c94065887645688a90391ba3ed7c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ia5c1595b125cd86fbdb0b88405847b5ab79dd7d2e82d6cb655063bdfb88687a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    iaabe423d98853e5679e2818f79c48cf8118446ab0e21e9277dfa9e13c72563ce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    ib1a5b88aef5ff8f68dc100aa6864744698883c2b761dbed0c4d16d621f1e0589 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    ic2e82087cc46ed3985fa760f476712dd98b190edcfc4f8f468227dbc289b7dc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    ida1717c6d2c83acb1b27edfb4be4c14af316e8fa40c5765c30a79f3925918af8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    ie188e205d8571822e0c7ee1d8d8f45eb480f9c1275d387580b42ee3da4e4daa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    if3b007988b97770c3ad7134cd0dcef2daa2ad3a078b4962bce639aa83f37e6a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i8c3d3b2a3a519127b6269b0369442d6e6cb2986e9f9ddba9b16e43f6f896a968 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i9eddaccbda0ff78804a00c5558444b2178859466c90e5da168ade311c1aa6cae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    ia1a9fbc268691d060abeb1b734aebb4136a647b1bc84c38af59edadc0c236071 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    id25a8573715e00d920f41a8560aa207878b8007cb2dcc2713421f69cb65c93d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i6e665073402a7c287f185798d81d13998809adcd52296fa21e8872685d14f0f0.AcceptRequestBuilder) {
    return i6e665073402a7c287f185798d81d13998809adcd52296fa21e8872685d14f0f0.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i746ca6978b7e4e298f4b3a396747dc93eabc3b1af8630010cdfef674a8b2ba5a.AttachmentsRequestBuilder) {
    return i746ca6978b7e4e298f4b3a396747dc93eabc3b1af8630010cdfef674a8b2ba5a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*id25a8573715e00d920f41a8560aa207878b8007cb2dcc2713421f69cb65c93d4.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return id25a8573715e00d920f41a8560aa207878b8007cb2dcc2713421f69cb65c93d4.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ie188e205d8571822e0c7ee1d8d8f45eb480f9c1275d387580b42ee3da4e4daa0.CalendarRequestBuilder) {
    return ie188e205d8571822e0c7ee1d8d8f45eb480f9c1275d387580b42ee3da4e4daa0.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*if3b007988b97770c3ad7134cd0dcef2daa2ad3a078b4962bce639aa83f37e6a7.CancelRequestBuilder) {
    return if3b007988b97770c3ad7134cd0dcef2daa2ad3a078b4962bce639aa83f37e6a7.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i58d26e13a36f9ef40d2af09c4e8691b48132d0c9426cbf53d095c5170731109d.DeclineRequestBuilder) {
    return i58d26e13a36f9ef40d2af09c4e8691b48132d0c9426cbf53d095c5170731109d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i4dee44f6ca6933f0d202e222b4a177b72d807831149a9b155943116a9fc6c608.DismissReminderRequestBuilder) {
    return i4dee44f6ca6933f0d202e222b4a177b72d807831149a9b155943116a9fc6c608.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ic2e82087cc46ed3985fa760f476712dd98b190edcfc4f8f468227dbc289b7dc6.ExtensionsRequestBuilder) {
    return ic2e82087cc46ed3985fa760f476712dd98b190edcfc4f8f468227dbc289b7dc6.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia1a9fbc268691d060abeb1b734aebb4136a647b1bc84c38af59edadc0c236071.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ia1a9fbc268691d060abeb1b734aebb4136a647b1bc84c38af59edadc0c236071.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ida1717c6d2c83acb1b27edfb4be4c14af316e8fa40c5765c30a79f3925918af8.ForwardRequestBuilder) {
    return ida1717c6d2c83acb1b27edfb4be4c14af316e8fa40c5765c30a79f3925918af8.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ib1a5b88aef5ff8f68dc100aa6864744698883c2b761dbed0c4d16d621f1e0589.MultiValueExtendedPropertiesRequestBuilder) {
    return ib1a5b88aef5ff8f68dc100aa6864744698883c2b761dbed0c4d16d621f1e0589.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9eddaccbda0ff78804a00c5558444b2178859466c90e5da168ade311c1aa6cae.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i9eddaccbda0ff78804a00c5558444b2178859466c90e5da168ade311c1aa6cae.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia5c1595b125cd86fbdb0b88405847b5ab79dd7d2e82d6cb655063bdfb88687a6.SingleValueExtendedPropertiesRequestBuilder) {
    return ia5c1595b125cd86fbdb0b88405847b5ab79dd7d2e82d6cb655063bdfb88687a6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8c3d3b2a3a519127b6269b0369442d6e6cb2986e9f9ddba9b16e43f6f896a968.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i8c3d3b2a3a519127b6269b0369442d6e6cb2986e9f9ddba9b16e43f6f896a968.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i8591b3851ec293b18b7b25a8487164392ad4c94065887645688a90391ba3ed7c.SnoozeReminderRequestBuilder) {
    return i8591b3851ec293b18b7b25a8487164392ad4c94065887645688a90391ba3ed7c.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*iaabe423d98853e5679e2818f79c48cf8118446ab0e21e9277dfa9e13c72563ce.TentativelyAcceptRequestBuilder) {
    return iaabe423d98853e5679e2818f79c48cf8118446ab0e21e9277dfa9e13c72563ce.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
