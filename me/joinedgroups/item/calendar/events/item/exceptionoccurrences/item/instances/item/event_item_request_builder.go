package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i16f879f27d6974703591181304a59908a67b0f0f810f948d13a416513a6fd406 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i1b488a7e379d2b5e2813602e108f28269b5220b655ff8ab05c5882901829c391 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/cancel"
    i3422ce8167813ae98ab978944a7f9b5f60403edbb948d62ea8eef3c5669bd256 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/extensions"
    i3d7c92b071f231ee10215734f773133ffbc4586db2eeb3f3273080eace80297f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/forward"
    i3e8320657b569e4014c8380d535fa8745ed693caec9225939208cee69d9f4bb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i512705d0d0e7995dc057e30309bca0a90a9e81e673002c2c4472307f2135762f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i6d40b00d929ec39f8c0b5578eb60a598b3700a5ed1b3ad9895e776b62ffe82d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/decline"
    ia6cc28432d7c6eef55835a2acf5cad55ffe1804a30de69cf1cc450a371646349 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/calendar"
    iab6d1bfc72a15e9574f736193f46251c319ed25b86f6d766f38517e343327660 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    iad1242f7b25ef6737d3b56ef2b0b72332c2a69c287c120962f393608946c8d11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/accept"
    id69fcf7e22af9df1aab3199619ac64781181d7e4ec1048deb2ce426ce135cb15 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/attachments"
    idfec674514b7bd2e1f71e52b4b41ab51363b520aef80d02d084c50df12ac9798 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i3f235c785e0b7de3733b4d3ec26e4b5dcaf5f614763126e9322f0a29a6ae0c38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i7818a9edbeeee4d5a5ed1ebffb913e4021af6da1ea7839c3ec7cadbe11cb7e13 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i8f76ec027dcd83aba61496a5eb121e8b8343d5db289ec04c9e98f892ad323da0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    iac65f70456e1ce9942c5e8ffe3853a731412b7531c961326b6c30983a4c77848 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*iad1242f7b25ef6737d3b56ef2b0b72332c2a69c287c120962f393608946c8d11.AcceptRequestBuilder) {
    return iad1242f7b25ef6737d3b56ef2b0b72332c2a69c287c120962f393608946c8d11.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*id69fcf7e22af9df1aab3199619ac64781181d7e4ec1048deb2ce426ce135cb15.AttachmentsRequestBuilder) {
    return id69fcf7e22af9df1aab3199619ac64781181d7e4ec1048deb2ce426ce135cb15.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i3f235c785e0b7de3733b4d3ec26e4b5dcaf5f614763126e9322f0a29a6ae0c38.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i3f235c785e0b7de3733b4d3ec26e4b5dcaf5f614763126e9322f0a29a6ae0c38.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ia6cc28432d7c6eef55835a2acf5cad55ffe1804a30de69cf1cc450a371646349.CalendarRequestBuilder) {
    return ia6cc28432d7c6eef55835a2acf5cad55ffe1804a30de69cf1cc450a371646349.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i1b488a7e379d2b5e2813602e108f28269b5220b655ff8ab05c5882901829c391.CancelRequestBuilder) {
    return i1b488a7e379d2b5e2813602e108f28269b5220b655ff8ab05c5882901829c391.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i6d40b00d929ec39f8c0b5578eb60a598b3700a5ed1b3ad9895e776b62ffe82d6.DeclineRequestBuilder) {
    return i6d40b00d929ec39f8c0b5578eb60a598b3700a5ed1b3ad9895e776b62ffe82d6.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i512705d0d0e7995dc057e30309bca0a90a9e81e673002c2c4472307f2135762f.DismissReminderRequestBuilder) {
    return i512705d0d0e7995dc057e30309bca0a90a9e81e673002c2c4472307f2135762f.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i3422ce8167813ae98ab978944a7f9b5f60403edbb948d62ea8eef3c5669bd256.ExtensionsRequestBuilder) {
    return i3422ce8167813ae98ab978944a7f9b5f60403edbb948d62ea8eef3c5669bd256.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8f76ec027dcd83aba61496a5eb121e8b8343d5db289ec04c9e98f892ad323da0.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8f76ec027dcd83aba61496a5eb121e8b8343d5db289ec04c9e98f892ad323da0.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i3d7c92b071f231ee10215734f773133ffbc4586db2eeb3f3273080eace80297f.ForwardRequestBuilder) {
    return i3d7c92b071f231ee10215734f773133ffbc4586db2eeb3f3273080eace80297f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*idfec674514b7bd2e1f71e52b4b41ab51363b520aef80d02d084c50df12ac9798.MultiValueExtendedPropertiesRequestBuilder) {
    return idfec674514b7bd2e1f71e52b4b41ab51363b520aef80d02d084c50df12ac9798.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iac65f70456e1ce9942c5e8ffe3853a731412b7531c961326b6c30983a4c77848.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iac65f70456e1ce9942c5e8ffe3853a731412b7531c961326b6c30983a4c77848.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*iab6d1bfc72a15e9574f736193f46251c319ed25b86f6d766f38517e343327660.SingleValueExtendedPropertiesRequestBuilder) {
    return iab6d1bfc72a15e9574f736193f46251c319ed25b86f6d766f38517e343327660.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7818a9edbeeee4d5a5ed1ebffb913e4021af6da1ea7839c3ec7cadbe11cb7e13.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i7818a9edbeeee4d5a5ed1ebffb913e4021af6da1ea7839c3ec7cadbe11cb7e13.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3e8320657b569e4014c8380d535fa8745ed693caec9225939208cee69d9f4bb5.SnoozeReminderRequestBuilder) {
    return i3e8320657b569e4014c8380d535fa8745ed693caec9225939208cee69d9f4bb5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i16f879f27d6974703591181304a59908a67b0f0f810f948d13a416513a6fd406.TentativelyAcceptRequestBuilder) {
    return i16f879f27d6974703591181304a59908a67b0f0f810f948d13a416513a6fd406.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
