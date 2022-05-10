package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b52517d59d1b2ef50e4b1f39c2ffe7a4e235f102dcc0360178d060447b3c8c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances"
    i10e1932d04eefbae53bc87ff5d4651a07bb957be5ee58f3330e245e72700771e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/calendar"
    i2d9b8945dc83b907c9d47da84dc390145686ab3a96ceac68ca9d06092180d049 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/cancel"
    i3980aa2b2211e7b4e8cf1b9a66f3b5526fcbea7d2565bfa34bd4b4c24612bc96 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/accept"
    i3a48c6c8a417970e14685018993cf88049fd653210ae9d92d6b1c063cde3fdda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/forward"
    i833140e6983341926ed7fd99e255a277ce4d9ab2e9e94acb4c70f514be9b55da "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/attachments"
    i90bcb515892cd8998a863795a455b613e233b281e8ceb58b505448feef06ea38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/snoozereminder"
    i91f3e58f9d7797b49f76a0d5e1baba5f0439eeb53a72ddc07992348408368950 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/dismissreminder"
    i945502080217c8479dce813dc3ef672305f1d3bde81b27065a1b3d057fe54a31 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/extensions"
    i98af978ee1e6690e48a9bfab9d843eba2365af9cadb8d10fa1e707a12f589491 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences"
    i9d8b4e2f21133fa5a3651e4da784c89a64aa502a92d2a693f53c10b830dfe801 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/tentativelyaccept"
    ib890c01be51b75e4e9b18e45379e359e8f9a10900097e4bca033afcb87977c6b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/decline"
    iedec73f58177a1cce2e41cf4e9a6552ad45ee382e46a4e57087fe4b03043b466 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/multivalueextendedproperties"
    ifce7cfad8f10cb24abf3ba2eaf429b15ac6df184328a35a92788cdc03111d3f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/singlevalueextendedproperties"
    i18fe398a0807dfdb23d9467ce436c82e91409e0bb2aac6ff34d55675d2a1c7b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/exceptionoccurrences/item"
    i1dc7028d59f21825234ce8972e972ca08b8bb10dce1c9865fcae494431064f72 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item"
    ia900b769fbe5c138dd549d477ed249263e9e4602bbca6c405e7dc2643171d4ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/singlevalueextendedproperties/item"
    ib0a58dc0bd2c5cc647647993270a7fef3291e0f731d19793d4cefefaa0db650f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/attachments/item"
    ib4a35552498b15bb2130344b26cd1077bc511c75347dd7fe61da7d3bd4e9d5fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/extensions/item"
    ie44eda0817815b25010f4b9c46c14d6f0158152a662360c8e8d063239fc1159f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i3980aa2b2211e7b4e8cf1b9a66f3b5526fcbea7d2565bfa34bd4b4c24612bc96.AcceptRequestBuilder) {
    return i3980aa2b2211e7b4e8cf1b9a66f3b5526fcbea7d2565bfa34bd4b4c24612bc96.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i833140e6983341926ed7fd99e255a277ce4d9ab2e9e94acb4c70f514be9b55da.AttachmentsRequestBuilder) {
    return i833140e6983341926ed7fd99e255a277ce4d9ab2e9e94acb4c70f514be9b55da.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib0a58dc0bd2c5cc647647993270a7fef3291e0f731d19793d4cefefaa0db650f.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ib0a58dc0bd2c5cc647647993270a7fef3291e0f731d19793d4cefefaa0db650f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i10e1932d04eefbae53bc87ff5d4651a07bb957be5ee58f3330e245e72700771e.CalendarRequestBuilder) {
    return i10e1932d04eefbae53bc87ff5d4651a07bb957be5ee58f3330e245e72700771e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i2d9b8945dc83b907c9d47da84dc390145686ab3a96ceac68ca9d06092180d049.CancelRequestBuilder) {
    return i2d9b8945dc83b907c9d47da84dc390145686ab3a96ceac68ca9d06092180d049.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendar/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for me
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
// CreatePatchRequestInformation update the navigation property events in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in me
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
func (m *EventItemRequestBuilder) Decline()(*ib890c01be51b75e4e9b18e45379e359e8f9a10900097e4bca033afcb87977c6b.DeclineRequestBuilder) {
    return ib890c01be51b75e4e9b18e45379e359e8f9a10900097e4bca033afcb87977c6b.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property events for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i91f3e58f9d7797b49f76a0d5e1baba5f0439eeb53a72ddc07992348408368950.DismissReminderRequestBuilder) {
    return i91f3e58f9d7797b49f76a0d5e1baba5f0439eeb53a72ddc07992348408368950.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i98af978ee1e6690e48a9bfab9d843eba2365af9cadb8d10fa1e707a12f589491.ExceptionOccurrencesRequestBuilder) {
    return i98af978ee1e6690e48a9bfab9d843eba2365af9cadb8d10fa1e707a12f589491.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i18fe398a0807dfdb23d9467ce436c82e91409e0bb2aac6ff34d55675d2a1c7b9.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i18fe398a0807dfdb23d9467ce436c82e91409e0bb2aac6ff34d55675d2a1c7b9.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i945502080217c8479dce813dc3ef672305f1d3bde81b27065a1b3d057fe54a31.ExtensionsRequestBuilder) {
    return i945502080217c8479dce813dc3ef672305f1d3bde81b27065a1b3d057fe54a31.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib4a35552498b15bb2130344b26cd1077bc511c75347dd7fe61da7d3bd4e9d5fc.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ib4a35552498b15bb2130344b26cd1077bc511c75347dd7fe61da7d3bd4e9d5fc.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i3a48c6c8a417970e14685018993cf88049fd653210ae9d92d6b1c063cde3fdda.ForwardRequestBuilder) {
    return i3a48c6c8a417970e14685018993cf88049fd653210ae9d92d6b1c063cde3fdda.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*i0b52517d59d1b2ef50e4b1f39c2ffe7a4e235f102dcc0360178d060447b3c8c0.InstancesRequestBuilder) {
    return i0b52517d59d1b2ef50e4b1f39c2ffe7a4e235f102dcc0360178d060447b3c8c0.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i1dc7028d59f21825234ce8972e972ca08b8bb10dce1c9865fcae494431064f72.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i1dc7028d59f21825234ce8972e972ca08b8bb10dce1c9865fcae494431064f72.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iedec73f58177a1cce2e41cf4e9a6552ad45ee382e46a4e57087fe4b03043b466.MultiValueExtendedPropertiesRequestBuilder) {
    return iedec73f58177a1cce2e41cf4e9a6552ad45ee382e46a4e57087fe4b03043b466.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie44eda0817815b25010f4b9c46c14d6f0158152a662360c8e8d063239fc1159f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ie44eda0817815b25010f4b9c46c14d6f0158152a662360c8e8d063239fc1159f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property events in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ifce7cfad8f10cb24abf3ba2eaf429b15ac6df184328a35a92788cdc03111d3f1.SingleValueExtendedPropertiesRequestBuilder) {
    return ifce7cfad8f10cb24abf3ba2eaf429b15ac6df184328a35a92788cdc03111d3f1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia900b769fbe5c138dd549d477ed249263e9e4602bbca6c405e7dc2643171d4ff.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ia900b769fbe5c138dd549d477ed249263e9e4602bbca6c405e7dc2643171d4ff.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i90bcb515892cd8998a863795a455b613e233b281e8ceb58b505448feef06ea38.SnoozeReminderRequestBuilder) {
    return i90bcb515892cd8998a863795a455b613e233b281e8ceb58b505448feef06ea38.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i9d8b4e2f21133fa5a3651e4da784c89a64aa502a92d2a693f53c10b830dfe801.TentativelyAcceptRequestBuilder) {
    return i9d8b4e2f21133fa5a3651e4da784c89a64aa502a92d2a693f53c10b830dfe801.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
