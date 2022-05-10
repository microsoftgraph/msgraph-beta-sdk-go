package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0bdcf60800a6b056ee1c8565e7447b9daaec86d9e60ad52ca8b5b8b2b75d41a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions"
    i2636d5b634d5f746a7f01f6ff7685d68dfdd7c4041b6631006c2b30fcd372f2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i2d78d70aea23c58562ba4f9e321a72590f9e770c269e042f117c7cb644bc094f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/calendar"
    i2ecba6bbe776ade2d62fbde0ba756508fec124a122947ddf0b43071c1815238d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/accept"
    i4794613f6931d5adad53d177863b1bf2dc256f5c2cf078f70fcade11984b69b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i5230faf1ccb5f0630c36ae740d8fdfba48b56fc7dce01d9d19b67bebe14f8a29 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/decline"
    i79d88ccfb00f803fa2a59fa9f6e19418c1b46acd423dc38f029cce8bcc9e23c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments"
    ia03d91e3851de2f534af7059494d28f5b3eb2a8e3a31fe3c0ffb3212a37398a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    iaa3859a8fce518756a3b573e9cc29cc15e0a7ba801770bafc7f77079a28d7ec0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/cancel"
    icd4d37ee472be1b09803f8f7ec4533f9fabbcd66b73ff013ee7980f18d535627 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/forward"
    icfabe077b6c5fc569e663a5a5ffa9e216005c443ba86efe98571b723680c8cc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ieb2d4625aac25a0f82039e5c7b1af2e8a27e1f3912b2f13cf21c2666c4665a4f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i29ebde018d55cff5f772b3b9ec238ae3baeb33f5d6c5ce00b4329c6254431c74 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    i4aaa148fe59f51524c0809fd51c618abdb44b296d60d6499482aee98589ef15e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i7ae42577dbdcb983bccb60cb46d673a9c183590f584606acdfc25ce04d68f7c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    ic4f7f1d15191e252a73fe9167b62563e09f646027522e909ae8312239fc93ef8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Accept()(*i2ecba6bbe776ade2d62fbde0ba756508fec124a122947ddf0b43071c1815238d.AcceptRequestBuilder) {
    return i2ecba6bbe776ade2d62fbde0ba756508fec124a122947ddf0b43071c1815238d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i79d88ccfb00f803fa2a59fa9f6e19418c1b46acd423dc38f029cce8bcc9e23c4.AttachmentsRequestBuilder) {
    return i79d88ccfb00f803fa2a59fa9f6e19418c1b46acd423dc38f029cce8bcc9e23c4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i29ebde018d55cff5f772b3b9ec238ae3baeb33f5d6c5ce00b4329c6254431c74.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i29ebde018d55cff5f772b3b9ec238ae3baeb33f5d6c5ce00b4329c6254431c74.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i2d78d70aea23c58562ba4f9e321a72590f9e770c269e042f117c7cb644bc094f.CalendarRequestBuilder) {
    return i2d78d70aea23c58562ba4f9e321a72590f9e770c269e042f117c7cb644bc094f.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*iaa3859a8fce518756a3b573e9cc29cc15e0a7ba801770bafc7f77079a28d7ec0.CancelRequestBuilder) {
    return iaa3859a8fce518756a3b573e9cc29cc15e0a7ba801770bafc7f77079a28d7ec0.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for me
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
// CreateGetRequestInformation get exceptionOccurrences from me
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from me
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) Decline()(*i5230faf1ccb5f0630c36ae740d8fdfba48b56fc7dce01d9d19b67bebe14f8a29.DeclineRequestBuilder) {
    return i5230faf1ccb5f0630c36ae740d8fdfba48b56fc7dce01d9d19b67bebe14f8a29.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property exceptionOccurrences for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*ieb2d4625aac25a0f82039e5c7b1af2e8a27e1f3912b2f13cf21c2666c4665a4f.DismissReminderRequestBuilder) {
    return ieb2d4625aac25a0f82039e5c7b1af2e8a27e1f3912b2f13cf21c2666c4665a4f.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i0bdcf60800a6b056ee1c8565e7447b9daaec86d9e60ad52ca8b5b8b2b75d41a1.ExtensionsRequestBuilder) {
    return i0bdcf60800a6b056ee1c8565e7447b9daaec86d9e60ad52ca8b5b8b2b75d41a1.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i7ae42577dbdcb983bccb60cb46d673a9c183590f584606acdfc25ce04d68f7c7.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i7ae42577dbdcb983bccb60cb46d673a9c183590f584606acdfc25ce04d68f7c7.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*icd4d37ee472be1b09803f8f7ec4533f9fabbcd66b73ff013ee7980f18d535627.ForwardRequestBuilder) {
    return icd4d37ee472be1b09803f8f7ec4533f9fabbcd66b73ff013ee7980f18d535627.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i2636d5b634d5f746a7f01f6ff7685d68dfdd7c4041b6631006c2b30fcd372f2e.MultiValueExtendedPropertiesRequestBuilder) {
    return i2636d5b634d5f746a7f01f6ff7685d68dfdd7c4041b6631006c2b30fcd372f2e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4aaa148fe59f51524c0809fd51c618abdb44b296d60d6499482aee98589ef15e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i4aaa148fe59f51524c0809fd51c618abdb44b296d60d6499482aee98589ef15e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia03d91e3851de2f534af7059494d28f5b3eb2a8e3a31fe3c0ffb3212a37398a0.SingleValueExtendedPropertiesRequestBuilder) {
    return ia03d91e3851de2f534af7059494d28f5b3eb2a8e3a31fe3c0ffb3212a37398a0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic4f7f1d15191e252a73fe9167b62563e09f646027522e909ae8312239fc93ef8.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ic4f7f1d15191e252a73fe9167b62563e09f646027522e909ae8312239fc93ef8.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i4794613f6931d5adad53d177863b1bf2dc256f5c2cf078f70fcade11984b69b1.SnoozeReminderRequestBuilder) {
    return i4794613f6931d5adad53d177863b1bf2dc256f5c2cf078f70fcade11984b69b1.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*icfabe077b6c5fc569e663a5a5ffa9e216005c443ba86efe98571b723680c8cc2.TentativelyAcceptRequestBuilder) {
    return icfabe077b6c5fc569e663a5a5ffa9e216005c443ba86efe98571b723680c8cc2.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
