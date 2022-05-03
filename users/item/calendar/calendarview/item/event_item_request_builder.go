package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c329946cfe9f65dc325f2a581d3900729e97e135f4fa5163e961fa1e6e6d508 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/snoozereminder"
    i0ca291604f0b37d4776c12c9384a6eeb886be25245701d1fc67067c0dacb9dab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/dismissreminder"
    i1b00a17c7bb6fd455766112e66daa7f315c3c2273539b6623fd787045e7cfc58 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/decline"
    i44e22bbc8a532c709bfd044f0c1182e6b9ac4fb8ba558f8b9d2015dc92483ee3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/tentativelyaccept"
    i4d7f4fc225d56ea372542c18aeb2c00b17087c1c476c4863be7d926d87d50330 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences"
    i816368b14e36a2a2d37e74b88ac8c6573fe1a447d99db2318bb45ac9cdc24556 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/singlevalueextendedproperties"
    id5bc622193e0269306e87950cc270ebfb747a961d3429a6e3ea98bc8222973a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/cancel"
    idbc896fd8428016fac95ec67fe116bd29913fc8de306ff8bfa377ceebcfb8100 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/attachments"
    ie314fe62a7e9af4b0d7b11457b5626b958f206d50d82443295107c58ef116f61 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances"
    ie50c98b38e8e98a9f53d53b7e33b241c4c543d0bc822b64dfff2b0c9ad8c5c9a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/extensions"
    ie8006d154169d84f839bc2404831e3bb11e537a670efbe3ea0fb678a847af840 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/multivalueextendedproperties"
    ie91f4ccb669f1f8d4f64d79430d177143bde7d6a5365d3543fe193853873802b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/forward"
    ifc93da8efa08649ecce830e5d3781181217fee6eb592ce5b00c098dbddf9dcb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/calendar"
    iff78016271ceddded3c710c7ba3b970caf8701ce868f7f32419f6cd8b5dc8fd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/accept"
    i1376f0a1866a80575b5784426dde2f70d42a6b2774cd1355d41401b86343013e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item"
    i19637f5a41e3ccc2fde05974db4440d5f14a573949916f579e495143f263fa32 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/multivalueextendedproperties/item"
    i739bfffbafaabb840a2f950e7cb9a89e6f3cf0c1eb9e4da98494b920cf8630df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item"
    i8d5280e702197db0c5ffc91ef6e1d06ee120edcff7b29550bf05c1d7e21a43d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/extensions/item"
    iad641c02387a864f90ee4d52b38266ffa5accacfea788b495524a535d81ab8e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/singlevalueextendedproperties/item"
    ic4a1b50dec9f8e4a474bc37be7aa505a20fd2b9ab74f0a81fa330b2fcc1a301b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
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
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type EventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string
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
func (m *EventItemRequestBuilder) Accept()(*iff78016271ceddded3c710c7ba3b970caf8701ce868f7f32419f6cd8b5dc8fd3.AcceptRequestBuilder) {
    return iff78016271ceddded3c710c7ba3b970caf8701ce868f7f32419f6cd8b5dc8fd3.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*idbc896fd8428016fac95ec67fe116bd29913fc8de306ff8bfa377ceebcfb8100.AttachmentsRequestBuilder) {
    return idbc896fd8428016fac95ec67fe116bd29913fc8de306ff8bfa377ceebcfb8100.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ic4a1b50dec9f8e4a474bc37be7aa505a20fd2b9ab74f0a81fa330b2fcc1a301b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ic4a1b50dec9f8e4a474bc37be7aa505a20fd2b9ab74f0a81fa330b2fcc1a301b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ifc93da8efa08649ecce830e5d3781181217fee6eb592ce5b00c098dbddf9dcb8.CalendarRequestBuilder) {
    return ifc93da8efa08649ecce830e5d3781181217fee6eb592ce5b00c098dbddf9dcb8.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*id5bc622193e0269306e87950cc270ebfb747a961d3429a6e3ea98bc8222973a6.CancelRequestBuilder) {
    return id5bc622193e0269306e87950cc270ebfb747a961d3429a6e3ea98bc8222973a6.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}{?startDateTime,endDateTime,%24select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendarView for users
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the calendar view for the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation update the navigation property calendarView in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendarView in users
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
func (m *EventItemRequestBuilder) Decline()(*i1b00a17c7bb6fd455766112e66daa7f315c3c2273539b6623fd787045e7cfc58.DeclineRequestBuilder) {
    return i1b00a17c7bb6fd455766112e66daa7f315c3c2273539b6623fd787045e7cfc58.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property calendarView for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i0ca291604f0b37d4776c12c9384a6eeb886be25245701d1fc67067c0dacb9dab.DismissReminderRequestBuilder) {
    return i0ca291604f0b37d4776c12c9384a6eeb886be25245701d1fc67067c0dacb9dab.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i4d7f4fc225d56ea372542c18aeb2c00b17087c1c476c4863be7d926d87d50330.ExceptionOccurrencesRequestBuilder) {
    return i4d7f4fc225d56ea372542c18aeb2c00b17087c1c476c4863be7d926d87d50330.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i739bfffbafaabb840a2f950e7cb9a89e6f3cf0c1eb9e4da98494b920cf8630df.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i739bfffbafaabb840a2f950e7cb9a89e6f3cf0c1eb9e4da98494b920cf8630df.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ie50c98b38e8e98a9f53d53b7e33b241c4c543d0bc822b64dfff2b0c9ad8c5c9a.ExtensionsRequestBuilder) {
    return ie50c98b38e8e98a9f53d53b7e33b241c4c543d0bc822b64dfff2b0c9ad8c5c9a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8d5280e702197db0c5ffc91ef6e1d06ee120edcff7b29550bf05c1d7e21a43d7.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8d5280e702197db0c5ffc91ef6e1d06ee120edcff7b29550bf05c1d7e21a43d7.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ie91f4ccb669f1f8d4f64d79430d177143bde7d6a5365d3543fe193853873802b.ForwardRequestBuilder) {
    return ie91f4ccb669f1f8d4f64d79430d177143bde7d6a5365d3543fe193853873802b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the calendar view for the calendar. Navigation property. Read-only.
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*ie314fe62a7e9af4b0d7b11457b5626b958f206d50d82443295107c58ef116f61.InstancesRequestBuilder) {
    return ie314fe62a7e9af4b0d7b11457b5626b958f206d50d82443295107c58ef116f61.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i1376f0a1866a80575b5784426dde2f70d42a6b2774cd1355d41401b86343013e.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i1376f0a1866a80575b5784426dde2f70d42a6b2774cd1355d41401b86343013e.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie8006d154169d84f839bc2404831e3bb11e537a670efbe3ea0fb678a847af840.MultiValueExtendedPropertiesRequestBuilder) {
    return ie8006d154169d84f839bc2404831e3bb11e537a670efbe3ea0fb678a847af840.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i19637f5a41e3ccc2fde05974db4440d5f14a573949916f579e495143f263fa32.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i19637f5a41e3ccc2fde05974db4440d5f14a573949916f579e495143f263fa32.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property calendarView in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i816368b14e36a2a2d37e74b88ac8c6573fe1a447d99db2318bb45ac9cdc24556.SingleValueExtendedPropertiesRequestBuilder) {
    return i816368b14e36a2a2d37e74b88ac8c6573fe1a447d99db2318bb45ac9cdc24556.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iad641c02387a864f90ee4d52b38266ffa5accacfea788b495524a535d81ab8e7.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return iad641c02387a864f90ee4d52b38266ffa5accacfea788b495524a535d81ab8e7.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i0c329946cfe9f65dc325f2a581d3900729e97e135f4fa5163e961fa1e6e6d508.SnoozeReminderRequestBuilder) {
    return i0c329946cfe9f65dc325f2a581d3900729e97e135f4fa5163e961fa1e6e6d508.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i44e22bbc8a532c709bfd044f0c1182e6b9ac4fb8ba558f8b9d2015dc92483ee3.TentativelyAcceptRequestBuilder) {
    return i44e22bbc8a532c709bfd044f0c1182e6b9ac4fb8ba558f8b9d2015dc92483ee3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
