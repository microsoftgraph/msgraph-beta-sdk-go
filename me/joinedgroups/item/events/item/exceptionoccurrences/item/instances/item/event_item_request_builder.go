package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07255503c707e039bfbc4e78ad569b1445cfc77f63f78b317b207908c1f01d79 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i08476fd53d4e4dd3af88cb6aeb6644f881912d0c66f31c91f510b8815f4cfce4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    i15813f5dd8c6a6a9e56a888e5d13e16c0d745d61dd339ea241f8ddbafd15a3ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i3670220b1fbdafe8dc54a67b76278edb17e49a60dcdd9b06265f60c851c9f8d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i5f1c8ba186882d1bb33722f293f8e5932050ba4041760dd250100ee42fea07a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i7d0cf792995d2d58ade0071ba5ff5ee3a39dc05ad6a2db36a1f43cb86c7dfcec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    ia79a5ce0c83e86d28d4c6cf24d592a61f3842751e28952201818811dbc934566 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    iaafb5927318c6a3293eabd0d36ca5302c7282fc27a3f347136400543b8be2ad2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    iaf4b39a9c84a8166868b423ba3e13373ab8262080f975aee1c8e1716edd5a068 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/forward"
    ib0380a95b288c294f0a44571524b2586ee4ac11b18edcc2258ed371cf500eb04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/decline"
    id51f4700143ddb5d9eda71508a20f304e9af5c85b4c246895b117ea69feacf3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/accept"
    if1a35c6bf0fb75efc0fb59ebb41e8124031f6bb9e30d88327bec10236cc47fad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i706c8079e528b2203cf0780bedb91ccb0c7e6f225fffe3b9f1fcea973dbf672b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i92299c8817c702174e1e3c7a5fe935a1f8a29adcb3696dd4be6f048a4dfed63c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    id53069efc3892f2ffcfcd70b28f6d001ade2dcd970098b5e94be62f72dcf175c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    id59130fb323a87d647759491c1d07b299700a2c3296ab0379f5d516399acb6f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*id51f4700143ddb5d9eda71508a20f304e9af5c85b4c246895b117ea69feacf3a.AcceptRequestBuilder) {
    return id51f4700143ddb5d9eda71508a20f304e9af5c85b4c246895b117ea69feacf3a.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i07255503c707e039bfbc4e78ad569b1445cfc77f63f78b317b207908c1f01d79.AttachmentsRequestBuilder) {
    return i07255503c707e039bfbc4e78ad569b1445cfc77f63f78b317b207908c1f01d79.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i706c8079e528b2203cf0780bedb91ccb0c7e6f225fffe3b9f1fcea973dbf672b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i706c8079e528b2203cf0780bedb91ccb0c7e6f225fffe3b9f1fcea973dbf672b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i08476fd53d4e4dd3af88cb6aeb6644f881912d0c66f31c91f510b8815f4cfce4.CalendarRequestBuilder) {
    return i08476fd53d4e4dd3af88cb6aeb6644f881912d0c66f31c91f510b8815f4cfce4.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i7d0cf792995d2d58ade0071ba5ff5ee3a39dc05ad6a2db36a1f43cb86c7dfcec.CancelRequestBuilder) {
    return i7d0cf792995d2d58ade0071ba5ff5ee3a39dc05ad6a2db36a1f43cb86c7dfcec.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*ib0380a95b288c294f0a44571524b2586ee4ac11b18edcc2258ed371cf500eb04.DeclineRequestBuilder) {
    return ib0380a95b288c294f0a44571524b2586ee4ac11b18edcc2258ed371cf500eb04.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i5f1c8ba186882d1bb33722f293f8e5932050ba4041760dd250100ee42fea07a1.DismissReminderRequestBuilder) {
    return i5f1c8ba186882d1bb33722f293f8e5932050ba4041760dd250100ee42fea07a1.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i3670220b1fbdafe8dc54a67b76278edb17e49a60dcdd9b06265f60c851c9f8d5.ExtensionsRequestBuilder) {
    return i3670220b1fbdafe8dc54a67b76278edb17e49a60dcdd9b06265f60c851c9f8d5.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*id53069efc3892f2ffcfcd70b28f6d001ade2dcd970098b5e94be62f72dcf175c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return id53069efc3892f2ffcfcd70b28f6d001ade2dcd970098b5e94be62f72dcf175c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*iaf4b39a9c84a8166868b423ba3e13373ab8262080f975aee1c8e1716edd5a068.ForwardRequestBuilder) {
    return iaf4b39a9c84a8166868b423ba3e13373ab8262080f975aee1c8e1716edd5a068.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*if1a35c6bf0fb75efc0fb59ebb41e8124031f6bb9e30d88327bec10236cc47fad.MultiValueExtendedPropertiesRequestBuilder) {
    return if1a35c6bf0fb75efc0fb59ebb41e8124031f6bb9e30d88327bec10236cc47fad.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i92299c8817c702174e1e3c7a5fe935a1f8a29adcb3696dd4be6f048a4dfed63c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i92299c8817c702174e1e3c7a5fe935a1f8a29adcb3696dd4be6f048a4dfed63c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i15813f5dd8c6a6a9e56a888e5d13e16c0d745d61dd339ea241f8ddbafd15a3ba.SingleValueExtendedPropertiesRequestBuilder) {
    return i15813f5dd8c6a6a9e56a888e5d13e16c0d745d61dd339ea241f8ddbafd15a3ba.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id59130fb323a87d647759491c1d07b299700a2c3296ab0379f5d516399acb6f6.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return id59130fb323a87d647759491c1d07b299700a2c3296ab0379f5d516399acb6f6.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*iaafb5927318c6a3293eabd0d36ca5302c7282fc27a3f347136400543b8be2ad2.SnoozeReminderRequestBuilder) {
    return iaafb5927318c6a3293eabd0d36ca5302c7282fc27a3f347136400543b8be2ad2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ia79a5ce0c83e86d28d4c6cf24d592a61f3842751e28952201818811dbc934566.TentativelyAcceptRequestBuilder) {
    return ia79a5ce0c83e86d28d4c6cf24d592a61f3842751e28952201818811dbc934566.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
