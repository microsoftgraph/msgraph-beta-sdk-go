package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i243c01333083b5c295c2309644c3df489a5cc335e9c51b4ffc2e035f9b9f0bb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/cancel"
    i30abab7738629124ac93833fc45023fe2b12fa3b590a762461e788d2bd838144 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/accept"
    i35ef6d6d8a8e50f3bf23b92484f453495489cb5b4dc08d8edc763d59516147b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/tentativelyaccept"
    i475eabc4be5936a74902d42990b29cbd9a9033e66f0f2aec0030db907f66b21a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences"
    i4ac671f44d47edc4c021ddc55f4fdd12076762fcf64e914f2b0c04cb8b1e0a95 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/extensions"
    i61c902510f86d1e684f6c489dc73d7a7031b3593e7b6bd33e761517cecda4509 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/calendar"
    i7ac887e020adfb20866da87ff9a788ba9581e708bbd9d4291d8220b2153f4abb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/attachments"
    i7d75c48be078c5d331f0743c295b6c5fc2fd88bec4f7970104907e5d7fd7d12b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/decline"
    i88af56ab1c01174d6be5137278d7993043f7111dbeea83dceae2e7697494e83f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/multivalueextendedproperties"
    ic5d6286fccff030cac0efa4d73ecf9251f387232ffc5b54ccec72c20c94d328b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/singlevalueextendedproperties"
    id3c9a84fc0b599211eeac42aacc7a3e569153d3c26f16a6c7fb67345f4591394 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/dismissreminder"
    idfa57c16d7c78464fcafb51ab0cae08c5d12332c0ff83770101d9aa469ace446 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/forward"
    iefe518378ee4448e26acfe4ac7308ecd3cc5cc39fee63160908ccd4fc6893d65 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/snoozereminder"
    i012a10d89efea0e588aec70b1be4ea3c5656a9af0b4164617343c7ee268ef30c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/multivalueextendedproperties/item"
    i24c4c3857dbdfbda065a30a2b3dc16699f20f178a367b64a96b7b413ef4b38db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/singlevalueextendedproperties/item"
    i5f29dbb8d13eb3b56698507c4607a355f3c05d8fb3e02fb0a50daeb535727a19 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/attachments/item"
    i88f1c4cf34b11638871974f9373f2a13f0c76e1716b13f49b66dd81c124fcb1e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/extensions/item"
    i8ed72975bceb80778c820f7be5ee8614b645106718dd99664c9e6980f7dd6bb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*i30abab7738629124ac93833fc45023fe2b12fa3b590a762461e788d2bd838144.AcceptRequestBuilder) {
    return i30abab7738629124ac93833fc45023fe2b12fa3b590a762461e788d2bd838144.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i7ac887e020adfb20866da87ff9a788ba9581e708bbd9d4291d8220b2153f4abb.AttachmentsRequestBuilder) {
    return i7ac887e020adfb20866da87ff9a788ba9581e708bbd9d4291d8220b2153f4abb.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i5f29dbb8d13eb3b56698507c4607a355f3c05d8fb3e02fb0a50daeb535727a19.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i5f29dbb8d13eb3b56698507c4607a355f3c05d8fb3e02fb0a50daeb535727a19.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i61c902510f86d1e684f6c489dc73d7a7031b3593e7b6bd33e761517cecda4509.CalendarRequestBuilder) {
    return i61c902510f86d1e684f6c489dc73d7a7031b3593e7b6bd33e761517cecda4509.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i243c01333083b5c295c2309644c3df489a5cc335e9c51b4ffc2e035f9b9f0bb6.CancelRequestBuilder) {
    return i243c01333083b5c295c2309644c3df489a5cc335e9c51b4ffc2e035f9b9f0bb6.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i7d75c48be078c5d331f0743c295b6c5fc2fd88bec4f7970104907e5d7fd7d12b.DeclineRequestBuilder) {
    return i7d75c48be078c5d331f0743c295b6c5fc2fd88bec4f7970104907e5d7fd7d12b.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*id3c9a84fc0b599211eeac42aacc7a3e569153d3c26f16a6c7fb67345f4591394.DismissReminderRequestBuilder) {
    return id3c9a84fc0b599211eeac42aacc7a3e569153d3c26f16a6c7fb67345f4591394.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i475eabc4be5936a74902d42990b29cbd9a9033e66f0f2aec0030db907f66b21a.ExceptionOccurrencesRequestBuilder) {
    return i475eabc4be5936a74902d42990b29cbd9a9033e66f0f2aec0030db907f66b21a.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i8ed72975bceb80778c820f7be5ee8614b645106718dd99664c9e6980f7dd6bb6.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i8ed72975bceb80778c820f7be5ee8614b645106718dd99664c9e6980f7dd6bb6.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i4ac671f44d47edc4c021ddc55f4fdd12076762fcf64e914f2b0c04cb8b1e0a95.ExtensionsRequestBuilder) {
    return i4ac671f44d47edc4c021ddc55f4fdd12076762fcf64e914f2b0c04cb8b1e0a95.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i88f1c4cf34b11638871974f9373f2a13f0c76e1716b13f49b66dd81c124fcb1e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i88f1c4cf34b11638871974f9373f2a13f0c76e1716b13f49b66dd81c124fcb1e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*idfa57c16d7c78464fcafb51ab0cae08c5d12332c0ff83770101d9aa469ace446.ForwardRequestBuilder) {
    return idfa57c16d7c78464fcafb51ab0cae08c5d12332c0ff83770101d9aa469ace446.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i88af56ab1c01174d6be5137278d7993043f7111dbeea83dceae2e7697494e83f.MultiValueExtendedPropertiesRequestBuilder) {
    return i88af56ab1c01174d6be5137278d7993043f7111dbeea83dceae2e7697494e83f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i012a10d89efea0e588aec70b1be4ea3c5656a9af0b4164617343c7ee268ef30c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i012a10d89efea0e588aec70b1be4ea3c5656a9af0b4164617343c7ee268ef30c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ic5d6286fccff030cac0efa4d73ecf9251f387232ffc5b54ccec72c20c94d328b.SingleValueExtendedPropertiesRequestBuilder) {
    return ic5d6286fccff030cac0efa4d73ecf9251f387232ffc5b54ccec72c20c94d328b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i24c4c3857dbdfbda065a30a2b3dc16699f20f178a367b64a96b7b413ef4b38db.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i24c4c3857dbdfbda065a30a2b3dc16699f20f178a367b64a96b7b413ef4b38db.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*iefe518378ee4448e26acfe4ac7308ecd3cc5cc39fee63160908ccd4fc6893d65.SnoozeReminderRequestBuilder) {
    return iefe518378ee4448e26acfe4ac7308ecd3cc5cc39fee63160908ccd4fc6893d65.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i35ef6d6d8a8e50f3bf23b92484f453495489cb5b4dc08d8edc763d59516147b7.TentativelyAcceptRequestBuilder) {
    return i35ef6d6d8a8e50f3bf23b92484f453495489cb5b4dc08d8edc763d59516147b7.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
