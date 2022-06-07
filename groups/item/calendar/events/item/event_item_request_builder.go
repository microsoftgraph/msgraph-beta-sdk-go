package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i059dab5e0aa78e83cfd615dda918a2c44d52daaa45372cb50f4bdcb514105cde "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/dismissreminder"
    i0d5d4fce97db38d2a5b93fbc1501a96a95c48e6dbb0a1641f9d615536e8fdb85 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/decline"
    i1201defebd7ce1e7d618f96de982cfb8aed706c05f9bfdecd38c1058138330fa "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/attachments"
    i14c2916673848ffbd3aceff4d7b49e2dc70fa60d7e7b789282b202b5b22bca0b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/multivalueextendedproperties"
    i2fa8f54576ab0073b0e114c1abe63c305fdf346eb1ba9dad9a8229c2238b7a4f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances"
    i3cbdd98e0afa2ffda5ae80ddc70ba716047bcfe5b400f6c3e7809c3d891b03ad "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/snoozereminder"
    i76c4305997d6cb13bf9da13aa31e7e7b4ea8c55cdf2ea79d14c2fdeeb448c185 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences"
    i770a5145fd01ee602a18c426ac71a2cdead8830b24f6ad3719436886bcc9d7c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/tentativelyaccept"
    i9128f3c60135f22c7a47ea9daa2314158aa7788abd19f1fbca2ad975760a7e07 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/extensions"
    ia44285a0d899332adefbad5c8c67cf547debbcaf746514cf5302c11d36802773 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/singlevalueextendedproperties"
    ic52a4b74b1cf2455bad81a20dc724df4cd71fe40313d94035c5b81bf66ea5133 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/calendar"
    ie81dfddf719cafd2afba4c80536ed434eec18f9c2171a25468bb1d24ba3bedeb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/forward"
    iefcbcbbdae997919fc5de05ed7114f2a54e10b4e2ac8918402b4edff67f30424 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/accept"
    if3156c0c94138e02899ab847edde1d4e045eb5a3b668b3dede6f5559615c6a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/cancel"
    i251d233f5532efc8f12bbf88b9ca31fa771e05cdc0f1ae5875f7add75a4c74c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/extensions/item"
    i2ebc6aacd0f573f384bf8d2d129513cff7da5807dab9633bd6908d1883c8ee6a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item"
    i47b8ece28c607d20c15c3a51208eede6e26c773b4fa6fbb653ec711cdf217df4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/attachments/item"
    i9855b087292f7172af644d450f0c55e1127a8a59839510cce0e9db7dae78568c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item"
    icf5853047b0325830c18b342a5e684e64a614c8eb80b63cdba2cfe39a25a6ebf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/singlevalueextendedproperties/item"
    ida403eba8e8430247eef196fbc4ec9d9c0bd649127d0bf203ade9a926cf3f55c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*iefcbcbbdae997919fc5de05ed7114f2a54e10b4e2ac8918402b4edff67f30424.AcceptRequestBuilder) {
    return iefcbcbbdae997919fc5de05ed7114f2a54e10b4e2ac8918402b4edff67f30424.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i1201defebd7ce1e7d618f96de982cfb8aed706c05f9bfdecd38c1058138330fa.AttachmentsRequestBuilder) {
    return i1201defebd7ce1e7d618f96de982cfb8aed706c05f9bfdecd38c1058138330fa.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i47b8ece28c607d20c15c3a51208eede6e26c773b4fa6fbb653ec711cdf217df4.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i47b8ece28c607d20c15c3a51208eede6e26c773b4fa6fbb653ec711cdf217df4.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ic52a4b74b1cf2455bad81a20dc724df4cd71fe40313d94035c5b81bf66ea5133.CalendarRequestBuilder) {
    return ic52a4b74b1cf2455bad81a20dc724df4cd71fe40313d94035c5b81bf66ea5133.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*if3156c0c94138e02899ab847edde1d4e045eb5a3b668b3dede6f5559615c6a91.CancelRequestBuilder) {
    return if3156c0c94138e02899ab847edde1d4e045eb5a3b668b3dede6f5559615c6a91.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for groups
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
// CreatePatchRequestInformation update the navigation property events in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in groups
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
func (m *EventItemRequestBuilder) Decline()(*i0d5d4fce97db38d2a5b93fbc1501a96a95c48e6dbb0a1641f9d615536e8fdb85.DeclineRequestBuilder) {
    return i0d5d4fce97db38d2a5b93fbc1501a96a95c48e6dbb0a1641f9d615536e8fdb85.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for groups
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property events for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*i059dab5e0aa78e83cfd615dda918a2c44d52daaa45372cb50f4bdcb514105cde.DismissReminderRequestBuilder) {
    return i059dab5e0aa78e83cfd615dda918a2c44d52daaa45372cb50f4bdcb514105cde.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i76c4305997d6cb13bf9da13aa31e7e7b4ea8c55cdf2ea79d14c2fdeeb448c185.ExceptionOccurrencesRequestBuilder) {
    return i76c4305997d6cb13bf9da13aa31e7e7b4ea8c55cdf2ea79d14c2fdeeb448c185.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i2ebc6aacd0f573f384bf8d2d129513cff7da5807dab9633bd6908d1883c8ee6a.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i2ebc6aacd0f573f384bf8d2d129513cff7da5807dab9633bd6908d1883c8ee6a.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i9128f3c60135f22c7a47ea9daa2314158aa7788abd19f1fbca2ad975760a7e07.ExtensionsRequestBuilder) {
    return i9128f3c60135f22c7a47ea9daa2314158aa7788abd19f1fbca2ad975760a7e07.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i251d233f5532efc8f12bbf88b9ca31fa771e05cdc0f1ae5875f7add75a4c74c8.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i251d233f5532efc8f12bbf88b9ca31fa771e05cdc0f1ae5875f7add75a4c74c8.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ie81dfddf719cafd2afba4c80536ed434eec18f9c2171a25468bb1d24ba3bedeb.ForwardRequestBuilder) {
    return ie81dfddf719cafd2afba4c80536ed434eec18f9c2171a25468bb1d24ba3bedeb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*i2fa8f54576ab0073b0e114c1abe63c305fdf346eb1ba9dad9a8229c2238b7a4f.InstancesRequestBuilder) {
    return i2fa8f54576ab0073b0e114c1abe63c305fdf346eb1ba9dad9a8229c2238b7a4f.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i9855b087292f7172af644d450f0c55e1127a8a59839510cce0e9db7dae78568c.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i9855b087292f7172af644d450f0c55e1127a8a59839510cce0e9db7dae78568c.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i14c2916673848ffbd3aceff4d7b49e2dc70fa60d7e7b789282b202b5b22bca0b.MultiValueExtendedPropertiesRequestBuilder) {
    return i14c2916673848ffbd3aceff4d7b49e2dc70fa60d7e7b789282b202b5b22bca0b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ida403eba8e8430247eef196fbc4ec9d9c0bd649127d0bf203ade9a926cf3f55c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ida403eba8e8430247eef196fbc4ec9d9c0bd649127d0bf203ade9a926cf3f55c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in groups
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property events in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia44285a0d899332adefbad5c8c67cf547debbcaf746514cf5302c11d36802773.SingleValueExtendedPropertiesRequestBuilder) {
    return ia44285a0d899332adefbad5c8c67cf547debbcaf746514cf5302c11d36802773.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*icf5853047b0325830c18b342a5e684e64a614c8eb80b63cdba2cfe39a25a6ebf.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return icf5853047b0325830c18b342a5e684e64a614c8eb80b63cdba2cfe39a25a6ebf.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3cbdd98e0afa2ffda5ae80ddc70ba716047bcfe5b400f6c3e7809c3d891b03ad.SnoozeReminderRequestBuilder) {
    return i3cbdd98e0afa2ffda5ae80ddc70ba716047bcfe5b400f6c3e7809c3d891b03ad.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i770a5145fd01ee602a18c426ac71a2cdead8830b24f6ad3719436886bcc9d7c9.TentativelyAcceptRequestBuilder) {
    return i770a5145fd01ee602a18c426ac71a2cdead8830b24f6ad3719436886bcc9d7c9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
