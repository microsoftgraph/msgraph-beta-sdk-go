package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1022eeefd0950e0e70752c5cb0b720b2a7fbd1152f8b2db65dbf1f1fe3785fdd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/tentativelyaccept"
    i16b073d42adfaa42cd2197fe997fd666e6c717513e29bbc1411d84d135e5907d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/dismissreminder"
    i3261245f326c44b01fe5da55ef3eb6fdcf3eec9727c6249445c533e07ad1762b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/cancel"
    i3cc615139480e6b6080500d2592d3b826f8172402d8123a8033c61f3b0f6d2b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/extensions"
    i449d1a85f5abadbfceb7b75147c603057ebd7e4b6fe7fbc07c7854cef9f748be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances"
    i5ccef22bea92cabddac92714e10b7b992bc8b281f5928395c1b100cef0454e2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences"
    i71f34fdb44622e1b8dbd50563e3b844d5b06315264bd4bf7dcfc1262b7101a4f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/attachments"
    i912a64a291cce9b7af41be42dd7d2c73b1358686172dc2d4200444b14724cf3c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/calendar"
    i9fdfb1b218cc10bbbf0b6f6529dc0097d2a03fa8fa6a246d23addca54cc1138f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/snoozereminder"
    ibe8940be56566d710a3152778a1869ba155c8636fa528fc2f319110c15e678a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/multivalueextendedproperties"
    ic365f138404b6574eb593b1e6e3be13031d99faa7fde1b344c069d91fe235dbc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/forward"
    iddfae0fb4ae92b93079596d35b8250f92d21dc1f6215fad3b6818349dc5a0e58 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/accept"
    ie9a1738dc4818849cb5a030a886b3107bfbc115e7d40261896f7c8b5c0a1d983 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/singlevalueextendedproperties"
    if65de254c43ef0974070ea922f0230c5434fa0eaef859e883b04b25ba3f75a0b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/decline"
    i35e38f254ec120c13f755d3867f2fd9952b17b0aef33b67009171560f14276ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/extensions/item"
    i39d2408d6b0b07a8e199b2d29e86bbc35c38980add543123b882bc022c320552 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/singlevalueextendedproperties/item"
    i53b1fc78501c4055752cde988ccaa74dc9af1ec1a4f11c6c09b19a16565c5fad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item"
    i5cd1ea32e6df077b61718a3134fe8e9004c569fa48259151b83a3713a0993009 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item"
    ib21562e8f33ec5c88876706437f50c8b6523900dedd39da40e1e91e9c54470d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/multivalueextendedproperties/item"
    ief2a3897f4132f3372d8f033cd2ebd787a4705005e05e502f64f9441d3294683 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
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
// EventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*iddfae0fb4ae92b93079596d35b8250f92d21dc1f6215fad3b6818349dc5a0e58.AcceptRequestBuilder) {
    return iddfae0fb4ae92b93079596d35b8250f92d21dc1f6215fad3b6818349dc5a0e58.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i71f34fdb44622e1b8dbd50563e3b844d5b06315264bd4bf7dcfc1262b7101a4f.AttachmentsRequestBuilder) {
    return i71f34fdb44622e1b8dbd50563e3b844d5b06315264bd4bf7dcfc1262b7101a4f.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ief2a3897f4132f3372d8f033cd2ebd787a4705005e05e502f64f9441d3294683.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ief2a3897f4132f3372d8f033cd2ebd787a4705005e05e502f64f9441d3294683.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i912a64a291cce9b7af41be42dd7d2c73b1358686172dc2d4200444b14724cf3c.CalendarRequestBuilder) {
    return i912a64a291cce9b7af41be42dd7d2c73b1358686172dc2d4200444b14724cf3c.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i3261245f326c44b01fe5da55ef3eb6fdcf3eec9727c6249445c533e07ad1762b.CancelRequestBuilder) {
    return i3261245f326c44b01fe5da55ef3eb6fdcf3eec9727c6249445c533e07ad1762b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for users
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the events in the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation update the navigation property events in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in users
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
func (m *EventItemRequestBuilder) Decline()(*if65de254c43ef0974070ea922f0230c5434fa0eaef859e883b04b25ba3f75a0b.DeclineRequestBuilder) {
    return if65de254c43ef0974070ea922f0230c5434fa0eaef859e883b04b25ba3f75a0b.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property events for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i16b073d42adfaa42cd2197fe997fd666e6c717513e29bbc1411d84d135e5907d.DismissReminderRequestBuilder) {
    return i16b073d42adfaa42cd2197fe997fd666e6c717513e29bbc1411d84d135e5907d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i5ccef22bea92cabddac92714e10b7b992bc8b281f5928395c1b100cef0454e2d.ExceptionOccurrencesRequestBuilder) {
    return i5ccef22bea92cabddac92714e10b7b992bc8b281f5928395c1b100cef0454e2d.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i5cd1ea32e6df077b61718a3134fe8e9004c569fa48259151b83a3713a0993009.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i5cd1ea32e6df077b61718a3134fe8e9004c569fa48259151b83a3713a0993009.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i3cc615139480e6b6080500d2592d3b826f8172402d8123a8033c61f3b0f6d2b0.ExtensionsRequestBuilder) {
    return i3cc615139480e6b6080500d2592d3b826f8172402d8123a8033c61f3b0f6d2b0.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i35e38f254ec120c13f755d3867f2fd9952b17b0aef33b67009171560f14276ed.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i35e38f254ec120c13f755d3867f2fd9952b17b0aef33b67009171560f14276ed.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ic365f138404b6574eb593b1e6e3be13031d99faa7fde1b344c069d91fe235dbc.ForwardRequestBuilder) {
    return ic365f138404b6574eb593b1e6e3be13031d99faa7fde1b344c069d91fe235dbc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the events in the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*i449d1a85f5abadbfceb7b75147c603057ebd7e4b6fe7fbc07c7854cef9f748be.InstancesRequestBuilder) {
    return i449d1a85f5abadbfceb7b75147c603057ebd7e4b6fe7fbc07c7854cef9f748be.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i53b1fc78501c4055752cde988ccaa74dc9af1ec1a4f11c6c09b19a16565c5fad.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i53b1fc78501c4055752cde988ccaa74dc9af1ec1a4f11c6c09b19a16565c5fad.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ibe8940be56566d710a3152778a1869ba155c8636fa528fc2f319110c15e678a5.MultiValueExtendedPropertiesRequestBuilder) {
    return ibe8940be56566d710a3152778a1869ba155c8636fa528fc2f319110c15e678a5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib21562e8f33ec5c88876706437f50c8b6523900dedd39da40e1e91e9c54470d7.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ib21562e8f33ec5c88876706437f50c8b6523900dedd39da40e1e91e9c54470d7.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property events in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ie9a1738dc4818849cb5a030a886b3107bfbc115e7d40261896f7c8b5c0a1d983.SingleValueExtendedPropertiesRequestBuilder) {
    return ie9a1738dc4818849cb5a030a886b3107bfbc115e7d40261896f7c8b5c0a1d983.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i39d2408d6b0b07a8e199b2d29e86bbc35c38980add543123b882bc022c320552.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i39d2408d6b0b07a8e199b2d29e86bbc35c38980add543123b882bc022c320552.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i9fdfb1b218cc10bbbf0b6f6529dc0097d2a03fa8fa6a246d23addca54cc1138f.SnoozeReminderRequestBuilder) {
    return i9fdfb1b218cc10bbbf0b6f6529dc0097d2a03fa8fa6a246d23addca54cc1138f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i1022eeefd0950e0e70752c5cb0b720b2a7fbd1152f8b2db65dbf1f1fe3785fdd.TentativelyAcceptRequestBuilder) {
    return i1022eeefd0950e0e70752c5cb0b720b2a7fbd1152f8b2db65dbf1f1fe3785fdd.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
