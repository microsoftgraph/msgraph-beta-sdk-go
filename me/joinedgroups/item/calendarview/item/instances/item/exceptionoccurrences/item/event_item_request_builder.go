package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04b3feec5267a634232164bcf260a24411c98338623a8132e7efe66a7466f81c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i271aa4a6fb0a48ab53ca826971488fc8e063df12daac1442c500cbcda1e51405 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i3957a702f9e5d9c66f8fdb1c17499dd3182e09dcee40f0918fa39af3c199b8aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i43ef85cf1c61bccb44443270379d9119f5dc8fb7ba9aae856796f9f5060e92ae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i4ece8594178013f3500f82c7e94b4677a1665b571ade0f6534f3d1ff19dfb2cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i53d7dbd1870d405befa5c7eabb14f06d2e25881b0ff83960ff80013f32d12566 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i615d0f236aa956b1444f7475b83d557537ae038cd8e3544c87da9709c561912a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i67ba307c9832031ee187511c847aac384642a5b0dac02a5e717f673f74615f08 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i762ad8929a10a82c5298204eef66c908c90f2f84fc61ff76bd39402a8651e23c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    ib2f9c33b14738dda5cb85fde00bd16ae09d5deb707d62069d73a767b03c0858a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    ib795057e50dd0b1f96d602fba500d01860f7862b32603bdbaef262777b99ca76 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    if4995258cbb6e028e5680d1d10eca0271183f06d4c25fcf979a39ddf1e2510d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i8ab6e68207abd0bbfa4123a8468d145a2ccc3d3b6d3e04bcb25f04abdedffd2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    iae6e74e396ad558371bce2178e9ae3c5a4a555051386e82831882fe9a8795c87 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ic7126f3d8672c0fe6a5a1c8b8853531acbf78a666408a475e7dd959081bd1c35 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ie2b0934b599cd2bb880e40b6093f621a54774a7015483d03d4528c233ecfb2a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*ib795057e50dd0b1f96d602fba500d01860f7862b32603bdbaef262777b99ca76.AcceptRequestBuilder) {
    return ib795057e50dd0b1f96d602fba500d01860f7862b32603bdbaef262777b99ca76.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i04b3feec5267a634232164bcf260a24411c98338623a8132e7efe66a7466f81c.AttachmentsRequestBuilder) {
    return i04b3feec5267a634232164bcf260a24411c98338623a8132e7efe66a7466f81c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ie2b0934b599cd2bb880e40b6093f621a54774a7015483d03d4528c233ecfb2a2.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ie2b0934b599cd2bb880e40b6093f621a54774a7015483d03d4528c233ecfb2a2.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i615d0f236aa956b1444f7475b83d557537ae038cd8e3544c87da9709c561912a.CalendarRequestBuilder) {
    return i615d0f236aa956b1444f7475b83d557537ae038cd8e3544c87da9709c561912a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i53d7dbd1870d405befa5c7eabb14f06d2e25881b0ff83960ff80013f32d12566.CancelRequestBuilder) {
    return i53d7dbd1870d405befa5c7eabb14f06d2e25881b0ff83960ff80013f32d12566.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i3957a702f9e5d9c66f8fdb1c17499dd3182e09dcee40f0918fa39af3c199b8aa.DeclineRequestBuilder) {
    return i3957a702f9e5d9c66f8fdb1c17499dd3182e09dcee40f0918fa39af3c199b8aa.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i762ad8929a10a82c5298204eef66c908c90f2f84fc61ff76bd39402a8651e23c.DismissReminderRequestBuilder) {
    return i762ad8929a10a82c5298204eef66c908c90f2f84fc61ff76bd39402a8651e23c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i4ece8594178013f3500f82c7e94b4677a1665b571ade0f6534f3d1ff19dfb2cb.ExtensionsRequestBuilder) {
    return i4ece8594178013f3500f82c7e94b4677a1665b571ade0f6534f3d1ff19dfb2cb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8ab6e68207abd0bbfa4123a8468d145a2ccc3d3b6d3e04bcb25f04abdedffd2e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8ab6e68207abd0bbfa4123a8468d145a2ccc3d3b6d3e04bcb25f04abdedffd2e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i43ef85cf1c61bccb44443270379d9119f5dc8fb7ba9aae856796f9f5060e92ae.ForwardRequestBuilder) {
    return i43ef85cf1c61bccb44443270379d9119f5dc8fb7ba9aae856796f9f5060e92ae.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ib2f9c33b14738dda5cb85fde00bd16ae09d5deb707d62069d73a767b03c0858a.MultiValueExtendedPropertiesRequestBuilder) {
    return ib2f9c33b14738dda5cb85fde00bd16ae09d5deb707d62069d73a767b03c0858a.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iae6e74e396ad558371bce2178e9ae3c5a4a555051386e82831882fe9a8795c87.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iae6e74e396ad558371bce2178e9ae3c5a4a555051386e82831882fe9a8795c87.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if4995258cbb6e028e5680d1d10eca0271183f06d4c25fcf979a39ddf1e2510d5.SingleValueExtendedPropertiesRequestBuilder) {
    return if4995258cbb6e028e5680d1d10eca0271183f06d4c25fcf979a39ddf1e2510d5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic7126f3d8672c0fe6a5a1c8b8853531acbf78a666408a475e7dd959081bd1c35.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ic7126f3d8672c0fe6a5a1c8b8853531acbf78a666408a475e7dd959081bd1c35.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i67ba307c9832031ee187511c847aac384642a5b0dac02a5e717f673f74615f08.SnoozeReminderRequestBuilder) {
    return i67ba307c9832031ee187511c847aac384642a5b0dac02a5e717f673f74615f08.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i271aa4a6fb0a48ab53ca826971488fc8e063df12daac1442c500cbcda1e51405.TentativelyAcceptRequestBuilder) {
    return i271aa4a6fb0a48ab53ca826971488fc8e063df12daac1442c500cbcda1e51405.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
