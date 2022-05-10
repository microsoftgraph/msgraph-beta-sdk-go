package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09945128fb2098fd125db19f2d3a012411ca00415620a211a60cb78dae4da987 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences"
    i122652aab3c9611379a14b91ad1b1cbeeb158fd3a61c5acafcf1a4d5f2ad1e9a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/multivalueextendedproperties"
    i1b5edb77fd4bba6a74bb44f8d76ccb3850b37a5b3b94ea1302971b700b9084ad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/extensions"
    i1b832b6b3c70ad9609443b3f0629fb38ca7ad6198da06f6ce7287070ee25b96e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/singlevalueextendedproperties"
    i1e05e4e9a66ce98b59fb198c0e3d294070bc4d66543b70577b898ed6186cb66e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/accept"
    i3c611a0581372d65489a4a01456ddfd4ead9d5e3d04a519627ab77a060e9c31e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/attachments"
    i484854fe4dfae6e2f052e247c46d552a6cb8fada250be481f44e62d30408f4a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/tentativelyaccept"
    i6a978c47e3b6a25d3ec5226f805046a442ac59b6bc46a71d84ea767055d60b72 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/snoozereminder"
    i76b94b74c1e6c17b8ecab4ce256d6ee727d1494756f17cedd5dcc29671e9bcbc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/decline"
    i77055dde4712106efd1645afad743597ad157c1d13ef37fd6b51aa9ea21a3c34 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/forward"
    i9817fa24ab735d53e6c5febb448dfe2aeb42eac28177d6427580b26a74a61071 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/cancel"
    ib12d09ae64159438f436cd2b799ad49c5dd1984119d43d79bd7c7c60753cdc30 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/calendar"
    ib965dc5f494822f35f6679782c6c41a8fdedd998f873ab891169b6fb59a45bec "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances"
    ica8a1c2a9cf40482694d5792659e29bbc43335ebaeed53d5621ce7e21b7da9f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/dismissreminder"
    i0ceeb653a84da85c05438f9e7df41a660a8da775310fd4140bd14f0b463aca2b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/multivalueextendedproperties/item"
    i2c41fa27d22a52f94f541373431839a5f4854fe38a916d21f07890d63ab6b22e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/attachments/item"
    i5df05e2f083146b00860db930564e797df1f5582dc904b390a93bfd43fe13e75 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item"
    i80393276f9718a86232f7367f1e59a496dac83876ad76c3389d4616a8ab0bc7a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/extensions/item"
    ib74e22f89003f716354e1ab576c7f7ae6439318e8da5bdc7a22d82984ae0ad50 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/instances/item"
    ifb49ed00957e1912aa177e865dc48a1c790ae39194cb1e937f0b3b189ac368ce "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i1e05e4e9a66ce98b59fb198c0e3d294070bc4d66543b70577b898ed6186cb66e.AcceptRequestBuilder) {
    return i1e05e4e9a66ce98b59fb198c0e3d294070bc4d66543b70577b898ed6186cb66e.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i3c611a0581372d65489a4a01456ddfd4ead9d5e3d04a519627ab77a060e9c31e.AttachmentsRequestBuilder) {
    return i3c611a0581372d65489a4a01456ddfd4ead9d5e3d04a519627ab77a060e9c31e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i2c41fa27d22a52f94f541373431839a5f4854fe38a916d21f07890d63ab6b22e.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i2c41fa27d22a52f94f541373431839a5f4854fe38a916d21f07890d63ab6b22e.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ib12d09ae64159438f436cd2b799ad49c5dd1984119d43d79bd7c7c60753cdc30.CalendarRequestBuilder) {
    return ib12d09ae64159438f436cd2b799ad49c5dd1984119d43d79bd7c7c60753cdc30.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i9817fa24ab735d53e6c5febb448dfe2aeb42eac28177d6427580b26a74a61071.CancelRequestBuilder) {
    return i9817fa24ab735d53e6c5febb448dfe2aeb42eac28177d6427580b26a74a61071.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendar/events/{event%2Did}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i76b94b74c1e6c17b8ecab4ce256d6ee727d1494756f17cedd5dcc29671e9bcbc.DeclineRequestBuilder) {
    return i76b94b74c1e6c17b8ecab4ce256d6ee727d1494756f17cedd5dcc29671e9bcbc.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*ica8a1c2a9cf40482694d5792659e29bbc43335ebaeed53d5621ce7e21b7da9f4.DismissReminderRequestBuilder) {
    return ica8a1c2a9cf40482694d5792659e29bbc43335ebaeed53d5621ce7e21b7da9f4.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i09945128fb2098fd125db19f2d3a012411ca00415620a211a60cb78dae4da987.ExceptionOccurrencesRequestBuilder) {
    return i09945128fb2098fd125db19f2d3a012411ca00415620a211a60cb78dae4da987.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i5df05e2f083146b00860db930564e797df1f5582dc904b390a93bfd43fe13e75.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i5df05e2f083146b00860db930564e797df1f5582dc904b390a93bfd43fe13e75.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i1b5edb77fd4bba6a74bb44f8d76ccb3850b37a5b3b94ea1302971b700b9084ad.ExtensionsRequestBuilder) {
    return i1b5edb77fd4bba6a74bb44f8d76ccb3850b37a5b3b94ea1302971b700b9084ad.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i80393276f9718a86232f7367f1e59a496dac83876ad76c3389d4616a8ab0bc7a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i80393276f9718a86232f7367f1e59a496dac83876ad76c3389d4616a8ab0bc7a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i77055dde4712106efd1645afad743597ad157c1d13ef37fd6b51aa9ea21a3c34.ForwardRequestBuilder) {
    return i77055dde4712106efd1645afad743597ad157c1d13ef37fd6b51aa9ea21a3c34.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*ib965dc5f494822f35f6679782c6c41a8fdedd998f873ab891169b6fb59a45bec.InstancesRequestBuilder) {
    return ib965dc5f494822f35f6679782c6c41a8fdedd998f873ab891169b6fb59a45bec.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ib74e22f89003f716354e1ab576c7f7ae6439318e8da5bdc7a22d82984ae0ad50.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ib74e22f89003f716354e1ab576c7f7ae6439318e8da5bdc7a22d82984ae0ad50.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i122652aab3c9611379a14b91ad1b1cbeeb158fd3a61c5acafcf1a4d5f2ad1e9a.MultiValueExtendedPropertiesRequestBuilder) {
    return i122652aab3c9611379a14b91ad1b1cbeeb158fd3a61c5acafcf1a4d5f2ad1e9a.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0ceeb653a84da85c05438f9e7df41a660a8da775310fd4140bd14f0b463aca2b.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i0ceeb653a84da85c05438f9e7df41a660a8da775310fd4140bd14f0b463aca2b.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i1b832b6b3c70ad9609443b3f0629fb38ca7ad6198da06f6ce7287070ee25b96e.SingleValueExtendedPropertiesRequestBuilder) {
    return i1b832b6b3c70ad9609443b3f0629fb38ca7ad6198da06f6ce7287070ee25b96e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifb49ed00957e1912aa177e865dc48a1c790ae39194cb1e937f0b3b189ac368ce.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ifb49ed00957e1912aa177e865dc48a1c790ae39194cb1e937f0b3b189ac368ce.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i6a978c47e3b6a25d3ec5226f805046a442ac59b6bc46a71d84ea767055d60b72.SnoozeReminderRequestBuilder) {
    return i6a978c47e3b6a25d3ec5226f805046a442ac59b6bc46a71d84ea767055d60b72.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i484854fe4dfae6e2f052e247c46d552a6cb8fada250be481f44e62d30408f4a0.TentativelyAcceptRequestBuilder) {
    return i484854fe4dfae6e2f052e247c46d552a6cb8fada250be481f44e62d30408f4a0.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
