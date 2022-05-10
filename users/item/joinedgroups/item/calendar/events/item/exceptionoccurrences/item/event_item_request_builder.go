package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i153a0a97e1dd9111f2dd9c4ce5dfd3c88f458d17791272ba7f632d0128bd49ea "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/accept"
    i200a66f50b5e64c96854d7ecf0a318c636a3532690b53451a0bf46e4ae6f126c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/cancel"
    i3459f89d7603e7d51a564229e61abddfd3f54699437d2014d49fb1a91c949ff5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/dismissreminder"
    i3b4711e1bba2c9458a3e21ec3a56732725e524e25e5bd5497f2352dbb028c759 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/tentativelyaccept"
    i433cf23cd5cacb2cf90540c1c55d782a85a4f91842467a3404a347b2dd5b60a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/forward"
    i4f75f03e04438ee220fb56a3830cfdfc22a274dd129c6f24e5f752ccbf4bd744 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/attachments"
    i6bca958f99fb0d144e17eb20b94775fd4b81455ebc8e8775ce1d3350855ffa31 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ia43fbec6322a2214a5651b1916a4a1528f95970a87628bd6bf61330c3bf96cd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/calendar"
    iaef0109334e3592eb6c06cc7d5c620b37f803fc8f46dda1f0eda988fc933b7a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/decline"
    ic69a08cb97c01851679fc169a6ad9e85507b2d6b78bb86f9fd893d0e5335492a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances"
    icd9443a6e76dc12f828924ed4edc176c7dec495512e0269934d0361380945532 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/extensions"
    if796084a67aae15a5a51ed233d4c7b886238a212203d725fc6be4f4a3e05ddd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/snoozereminder"
    ifff1d668db99941454301e6eed20e5e43d1cb251a7a13225d382d4b984f97333 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    i09f72ad8f1b2a60ddeaffcebb602829ffe90519509213aad0d82d9aaa5891d48 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i256ec336933b33992cfc665b473a47d64d9388ad5536592c44fbec02614b28aa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/instances/item"
    i5f0a8f67d2b176847981fe83929935861e0a48b46027b53ef59d884f7c1d1813 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/extensions/item"
    i6f4a1f18db97af96dbc535d8e22bc60032694af676e3d97c632922a2ca65ca28 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ib073245943b80fffa2d4f9cdbb1a572a3119b78f790a83a2432f7b5201c27879 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/events/item/exceptionoccurrences/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
func (m *EventItemRequestBuilder) Accept()(*i153a0a97e1dd9111f2dd9c4ce5dfd3c88f458d17791272ba7f632d0128bd49ea.AcceptRequestBuilder) {
    return i153a0a97e1dd9111f2dd9c4ce5dfd3c88f458d17791272ba7f632d0128bd49ea.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i4f75f03e04438ee220fb56a3830cfdfc22a274dd129c6f24e5f752ccbf4bd744.AttachmentsRequestBuilder) {
    return i4f75f03e04438ee220fb56a3830cfdfc22a274dd129c6f24e5f752ccbf4bd744.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib073245943b80fffa2d4f9cdbb1a572a3119b78f790a83a2432f7b5201c27879.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ib073245943b80fffa2d4f9cdbb1a572a3119b78f790a83a2432f7b5201c27879.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ia43fbec6322a2214a5651b1916a4a1528f95970a87628bd6bf61330c3bf96cd3.CalendarRequestBuilder) {
    return ia43fbec6322a2214a5651b1916a4a1528f95970a87628bd6bf61330c3bf96cd3.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i200a66f50b5e64c96854d7ecf0a318c636a3532690b53451a0bf46e4ae6f126c.CancelRequestBuilder) {
    return i200a66f50b5e64c96854d7ecf0a318c636a3532690b53451a0bf46e4ae6f126c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for users
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
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from users
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) Decline()(*iaef0109334e3592eb6c06cc7d5c620b37f803fc8f46dda1f0eda988fc933b7a1.DeclineRequestBuilder) {
    return iaef0109334e3592eb6c06cc7d5c620b37f803fc8f46dda1f0eda988fc933b7a1.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property exceptionOccurrences for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i3459f89d7603e7d51a564229e61abddfd3f54699437d2014d49fb1a91c949ff5.DismissReminderRequestBuilder) {
    return i3459f89d7603e7d51a564229e61abddfd3f54699437d2014d49fb1a91c949ff5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*icd9443a6e76dc12f828924ed4edc176c7dec495512e0269934d0361380945532.ExtensionsRequestBuilder) {
    return icd9443a6e76dc12f828924ed4edc176c7dec495512e0269934d0361380945532.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i5f0a8f67d2b176847981fe83929935861e0a48b46027b53ef59d884f7c1d1813.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i5f0a8f67d2b176847981fe83929935861e0a48b46027b53ef59d884f7c1d1813.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i433cf23cd5cacb2cf90540c1c55d782a85a4f91842467a3404a347b2dd5b60a3.ForwardRequestBuilder) {
    return i433cf23cd5cacb2cf90540c1c55d782a85a4f91842467a3404a347b2dd5b60a3.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) Instances()(*ic69a08cb97c01851679fc169a6ad9e85507b2d6b78bb86f9fd893d0e5335492a.InstancesRequestBuilder) {
    return ic69a08cb97c01851679fc169a6ad9e85507b2d6b78bb86f9fd893d0e5335492a.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i256ec336933b33992cfc665b473a47d64d9388ad5536592c44fbec02614b28aa.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i256ec336933b33992cfc665b473a47d64d9388ad5536592c44fbec02614b28aa.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ifff1d668db99941454301e6eed20e5e43d1cb251a7a13225d382d4b984f97333.MultiValueExtendedPropertiesRequestBuilder) {
    return ifff1d668db99941454301e6eed20e5e43d1cb251a7a13225d382d4b984f97333.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i09f72ad8f1b2a60ddeaffcebb602829ffe90519509213aad0d82d9aaa5891d48.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i09f72ad8f1b2a60ddeaffcebb602829ffe90519509213aad0d82d9aaa5891d48.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6bca958f99fb0d144e17eb20b94775fd4b81455ebc8e8775ce1d3350855ffa31.SingleValueExtendedPropertiesRequestBuilder) {
    return i6bca958f99fb0d144e17eb20b94775fd4b81455ebc8e8775ce1d3350855ffa31.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i6f4a1f18db97af96dbc535d8e22bc60032694af676e3d97c632922a2ca65ca28.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i6f4a1f18db97af96dbc535d8e22bc60032694af676e3d97c632922a2ca65ca28.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*if796084a67aae15a5a51ed233d4c7b886238a212203d725fc6be4f4a3e05ddd6.SnoozeReminderRequestBuilder) {
    return if796084a67aae15a5a51ed233d4c7b886238a212203d725fc6be4f4a3e05ddd6.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3b4711e1bba2c9458a3e21ec3a56732725e524e25e5bd5497f2352dbb028c759.TentativelyAcceptRequestBuilder) {
    return i3b4711e1bba2c9458a3e21ec3a56732725e524e25e5bd5497f2352dbb028c759.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
