package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01a75427b87be21975974ead88a87bc11a0709c3ff257bc26be186d41a0a6dc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i39bf7a839393dc55a8bbb350624604426be4999755136228fe076c7234ea89ca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    i701e00d5a1b1d13bdc306da2a2bcdd9028c711c8791d2e5325e0dc194a99bc1c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i779386181751c7be4dcd5eecdeea7ef777c045863fa59bb91e82cce74fe9c19c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i84df2d056251cdd6082430796dd7fd209517d0feb80e351ed8f93d8454af1ab9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    ia181efcbfba17c32b15057303477f154fce8a6cabd3846442d3b994a42ac9f2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    ib2be4b1670cd95a37592c65221138de112c1f6b84c0bd1d4d9599382cf23e369 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    ib5994a70546fff6d7f021a1da13430dca6b4edfec486a19d28e1f0a7df299c8b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    idbd684b8d115faa9d0d056d9bf7dbb363740d3288193f856123f40497fb7a1b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    ie7d4c862a793da46a2b38ca9c2d436536cbed1a4bb12f3454cd4b855975ec063 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    iea749971fa4855d622f98306a1ee3a54e6bbd33b8092b21e051ef99a0d7f370d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    iedf2c38e0d51fb87fadecdd8dfbb8405b9d4efc6a97521d0b3663da103fe9fe0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    iad7b042a3d32e58feea6a4c422187120eb3573c5757226f0739c18e7235f4862 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    ic8e9be5855152cb6a1adc53e3b7fec175c6ae633758aea5b11e4754c5309584c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    id3ad06d897bec8965054a252d38bd3ef7033676897d2fc2bb4d1745b8b4b59b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    idf1b3b046766220af9380b38583e7488e6124bef311e7d2b561f44cb58b30157 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*idbd684b8d115faa9d0d056d9bf7dbb363740d3288193f856123f40497fb7a1b5.AcceptRequestBuilder) {
    return idbd684b8d115faa9d0d056d9bf7dbb363740d3288193f856123f40497fb7a1b5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ib2be4b1670cd95a37592c65221138de112c1f6b84c0bd1d4d9599382cf23e369.AttachmentsRequestBuilder) {
    return ib2be4b1670cd95a37592c65221138de112c1f6b84c0bd1d4d9599382cf23e369.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ic8e9be5855152cb6a1adc53e3b7fec175c6ae633758aea5b11e4754c5309584c.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ic8e9be5855152cb6a1adc53e3b7fec175c6ae633758aea5b11e4754c5309584c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ib5994a70546fff6d7f021a1da13430dca6b4edfec486a19d28e1f0a7df299c8b.CalendarRequestBuilder) {
    return ib5994a70546fff6d7f021a1da13430dca6b4edfec486a19d28e1f0a7df299c8b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*iea749971fa4855d622f98306a1ee3a54e6bbd33b8092b21e051ef99a0d7f370d.CancelRequestBuilder) {
    return iea749971fa4855d622f98306a1ee3a54e6bbd33b8092b21e051ef99a0d7f370d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i39bf7a839393dc55a8bbb350624604426be4999755136228fe076c7234ea89ca.DeclineRequestBuilder) {
    return i39bf7a839393dc55a8bbb350624604426be4999755136228fe076c7234ea89ca.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i701e00d5a1b1d13bdc306da2a2bcdd9028c711c8791d2e5325e0dc194a99bc1c.DismissReminderRequestBuilder) {
    return i701e00d5a1b1d13bdc306da2a2bcdd9028c711c8791d2e5325e0dc194a99bc1c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ia181efcbfba17c32b15057303477f154fce8a6cabd3846442d3b994a42ac9f2d.ExtensionsRequestBuilder) {
    return ia181efcbfba17c32b15057303477f154fce8a6cabd3846442d3b994a42ac9f2d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*idf1b3b046766220af9380b38583e7488e6124bef311e7d2b561f44cb58b30157.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return idf1b3b046766220af9380b38583e7488e6124bef311e7d2b561f44cb58b30157.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i84df2d056251cdd6082430796dd7fd209517d0feb80e351ed8f93d8454af1ab9.ForwardRequestBuilder) {
    return i84df2d056251cdd6082430796dd7fd209517d0feb80e351ed8f93d8454af1ab9.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iedf2c38e0d51fb87fadecdd8dfbb8405b9d4efc6a97521d0b3663da103fe9fe0.MultiValueExtendedPropertiesRequestBuilder) {
    return iedf2c38e0d51fb87fadecdd8dfbb8405b9d4efc6a97521d0b3663da103fe9fe0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*id3ad06d897bec8965054a252d38bd3ef7033676897d2fc2bb4d1745b8b4b59b9.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return id3ad06d897bec8965054a252d38bd3ef7033676897d2fc2bb4d1745b8b4b59b9.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i01a75427b87be21975974ead88a87bc11a0709c3ff257bc26be186d41a0a6dc4.SingleValueExtendedPropertiesRequestBuilder) {
    return i01a75427b87be21975974ead88a87bc11a0709c3ff257bc26be186d41a0a6dc4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iad7b042a3d32e58feea6a4c422187120eb3573c5757226f0739c18e7235f4862.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return iad7b042a3d32e58feea6a4c422187120eb3573c5757226f0739c18e7235f4862.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie7d4c862a793da46a2b38ca9c2d436536cbed1a4bb12f3454cd4b855975ec063.SnoozeReminderRequestBuilder) {
    return ie7d4c862a793da46a2b38ca9c2d436536cbed1a4bb12f3454cd4b855975ec063.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i779386181751c7be4dcd5eecdeea7ef777c045863fa59bb91e82cce74fe9c19c.TentativelyAcceptRequestBuilder) {
    return i779386181751c7be4dcd5eecdeea7ef777c045863fa59bb91e82cce74fe9c19c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
