package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i010dc92817b077230bc03affd8a0be7a0090d9cfe8b9254a530a31cae9fcac9b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/calendar"
    i09d23257ca0f093bc1f460a1212bb51550e9c91c114555949e7e1eae66b4b6f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/decline"
    i18dcfdc47d5b00a31c42955b2c1daf51d8f445dd6125e98688eee5566bf49f60 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/forward"
    i31adfcd54b0ecce7c0db45f55e6021aa8787a527ce1a8c77b2a50e79f6603373 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/accept"
    i3494e6d78bf5ff9497af85edcdc616816ed8508abed745689a72c95a0df4084c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/tentativelyaccept"
    i5fdeb8bcb6b030c474deb3e72cd00c8ac2f8c7a526ee31a85fa200a260f4ae44 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/dismissreminder"
    i7eef8e5a0af06b56d5b56184daf0502d0abc4c50e6760595ff90b109a079f7c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances"
    i9c0dfb1b4b9236fb38dfb3028eb53f5a87b4f0eb858f42474a9dd2fe3227c20f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    ibec62fde36d2d5a62e923ab9eee993b35364c67f43fa6fc46158217aac880610 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ic388180a42e7d63134f8cf66516a8ab8f954060ad563b216983411b62c734ebb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/attachments"
    icbdded2aa5f0ea464a9c43b99cdae8a5fed1b0f1eb1b0151d546a8c6c341e8cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/cancel"
    iee90ee1946a38203ab4367eb687eced6a41dc30199fea98bacc28eee5d12c175 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/extensions"
    if8d9aab55457e361c2195e127fc3ff10ac89c0f4da97de974eb0d4fc7453bd05 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/snoozereminder"
    i4741aad84db11cb1a6aaea11baf7573dd20560672a18d89aebb04e6ae6fe6970 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/extensions/item"
    i773ffd1b6a16c3fe3658c8ff7079c04c1f36d857425d57ae9cf733d67fa8acaf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i93b236d1087168ddf74a7ee4d96a2a410b735f0f68481bb6c43986163642bc4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item"
    ibaf269aec317466f43a658d489ecfde1b8c75e0ddbc536ad686b9286a0ecf2b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    idb5681b2fc93cc50e3c63dba4be0cd097e3e18ef67650a6e1a00669c4fcf45e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i31adfcd54b0ecce7c0db45f55e6021aa8787a527ce1a8c77b2a50e79f6603373.AcceptRequestBuilder) {
    return i31adfcd54b0ecce7c0db45f55e6021aa8787a527ce1a8c77b2a50e79f6603373.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ic388180a42e7d63134f8cf66516a8ab8f954060ad563b216983411b62c734ebb.AttachmentsRequestBuilder) {
    return ic388180a42e7d63134f8cf66516a8ab8f954060ad563b216983411b62c734ebb.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*idb5681b2fc93cc50e3c63dba4be0cd097e3e18ef67650a6e1a00669c4fcf45e1.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return idb5681b2fc93cc50e3c63dba4be0cd097e3e18ef67650a6e1a00669c4fcf45e1.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i010dc92817b077230bc03affd8a0be7a0090d9cfe8b9254a530a31cae9fcac9b.CalendarRequestBuilder) {
    return i010dc92817b077230bc03affd8a0be7a0090d9cfe8b9254a530a31cae9fcac9b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*icbdded2aa5f0ea464a9c43b99cdae8a5fed1b0f1eb1b0151d546a8c6c341e8cd.CancelRequestBuilder) {
    return icbdded2aa5f0ea464a9c43b99cdae8a5fed1b0f1eb1b0151d546a8c6c341e8cd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i09d23257ca0f093bc1f460a1212bb51550e9c91c114555949e7e1eae66b4b6f1.DeclineRequestBuilder) {
    return i09d23257ca0f093bc1f460a1212bb51550e9c91c114555949e7e1eae66b4b6f1.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i5fdeb8bcb6b030c474deb3e72cd00c8ac2f8c7a526ee31a85fa200a260f4ae44.DismissReminderRequestBuilder) {
    return i5fdeb8bcb6b030c474deb3e72cd00c8ac2f8c7a526ee31a85fa200a260f4ae44.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*iee90ee1946a38203ab4367eb687eced6a41dc30199fea98bacc28eee5d12c175.ExtensionsRequestBuilder) {
    return iee90ee1946a38203ab4367eb687eced6a41dc30199fea98bacc28eee5d12c175.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i4741aad84db11cb1a6aaea11baf7573dd20560672a18d89aebb04e6ae6fe6970.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i4741aad84db11cb1a6aaea11baf7573dd20560672a18d89aebb04e6ae6fe6970.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i18dcfdc47d5b00a31c42955b2c1daf51d8f445dd6125e98688eee5566bf49f60.ForwardRequestBuilder) {
    return i18dcfdc47d5b00a31c42955b2c1daf51d8f445dd6125e98688eee5566bf49f60.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*i7eef8e5a0af06b56d5b56184daf0502d0abc4c50e6760595ff90b109a079f7c2.InstancesRequestBuilder) {
    return i7eef8e5a0af06b56d5b56184daf0502d0abc4c50e6760595ff90b109a079f7c2.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i93b236d1087168ddf74a7ee4d96a2a410b735f0f68481bb6c43986163642bc4e.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i93b236d1087168ddf74a7ee4d96a2a410b735f0f68481bb6c43986163642bc4e.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i9c0dfb1b4b9236fb38dfb3028eb53f5a87b4f0eb858f42474a9dd2fe3227c20f.MultiValueExtendedPropertiesRequestBuilder) {
    return i9c0dfb1b4b9236fb38dfb3028eb53f5a87b4f0eb858f42474a9dd2fe3227c20f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibaf269aec317466f43a658d489ecfde1b8c75e0ddbc536ad686b9286a0ecf2b3.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ibaf269aec317466f43a658d489ecfde1b8c75e0ddbc536ad686b9286a0ecf2b3.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ibec62fde36d2d5a62e923ab9eee993b35364c67f43fa6fc46158217aac880610.SingleValueExtendedPropertiesRequestBuilder) {
    return ibec62fde36d2d5a62e923ab9eee993b35364c67f43fa6fc46158217aac880610.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i773ffd1b6a16c3fe3658c8ff7079c04c1f36d857425d57ae9cf733d67fa8acaf.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i773ffd1b6a16c3fe3658c8ff7079c04c1f36d857425d57ae9cf733d67fa8acaf.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*if8d9aab55457e361c2195e127fc3ff10ac89c0f4da97de974eb0d4fc7453bd05.SnoozeReminderRequestBuilder) {
    return if8d9aab55457e361c2195e127fc3ff10ac89c0f4da97de974eb0d4fc7453bd05.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3494e6d78bf5ff9497af85edcdc616816ed8508abed745689a72c95a0df4084c.TentativelyAcceptRequestBuilder) {
    return i3494e6d78bf5ff9497af85edcdc616816ed8508abed745689a72c95a0df4084c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
