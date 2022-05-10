package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0a6857a29775f3527ade48d34c6b10728a71af6dd78f6079d11d7a8c550d33d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/decline"
    i2b2f3d27aa9083c49b8996cc9e960d32e3e186f3a4d0747f2f5d7f7a012876f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/forward"
    i2d5ec4a49b750fc5ad295afe48c94a5ab6cef5a9fc91371d7fb50fa0313d3d36 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    i2fa86fdbf8669734c4ff4baec8a07f4167e970aa9cea188b2a558f8575787f5e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i3fae8342aedac4a38054d7a1b7f39f1aba0b6d7dab4f6bc060223ed82b262e66 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/cancel"
    i51370417203f2995dc2c3d12892585081cbbe62d6fc6ab078d6ab9a032c583de "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/calendar"
    i600d2d25136676bc3455d985eb1fcb7cba8c523738d73efc6515bbcbbd5f81c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i92f67b392c482516e6602b38183e458635d5e30dd2198233a40313a09d9ae4fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/attachments"
    ib043c85c14658a9046e5ec7b35d56973a9523498c18c990c3d2c449ece96a914 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    ib9679d22124b56eee2727407ed877bdcb26037fbd160079af337cae5fb4d0daa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/dismissreminder"
    id538d185b63affdf922df6fd539fc310753e12da2cf320d1b88f99a274eb630c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/extensions"
    ie741db865b9092fbc0b4e545d3139568a5f72efb36f3174af384bac9f32a913f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/accept"
    ifebec4f242fc2fe3f2ea086e8d1cbf3af35e7102fe2ffa1cb26ae92db66cf2fe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances"
    i57e38ffd72170d9573806545b3dcb9b944291909edbb83dba3cdb0286be63625 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/extensions/item"
    i5d87c261930d44d33fc6debc682d69d91880d07c4002721d4eeaa087a2285b13 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i67a2dce04b02ebfce753343824b3c5282d62c261c455c73f23e7cea86b9d9ac5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i74f629d6bb0be831d2a137d49ec5069ad5842bccb2e79a63092c5a89772f6977 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item"
    ie2b9a5053f203680ed3694469faffe46cc31da2f3c7315c5f201ea7b4537f244 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*ie741db865b9092fbc0b4e545d3139568a5f72efb36f3174af384bac9f32a913f.AcceptRequestBuilder) {
    return ie741db865b9092fbc0b4e545d3139568a5f72efb36f3174af384bac9f32a913f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i92f67b392c482516e6602b38183e458635d5e30dd2198233a40313a09d9ae4fd.AttachmentsRequestBuilder) {
    return i92f67b392c482516e6602b38183e458635d5e30dd2198233a40313a09d9ae4fd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ie2b9a5053f203680ed3694469faffe46cc31da2f3c7315c5f201ea7b4537f244.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ie2b9a5053f203680ed3694469faffe46cc31da2f3c7315c5f201ea7b4537f244.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i51370417203f2995dc2c3d12892585081cbbe62d6fc6ab078d6ab9a032c583de.CalendarRequestBuilder) {
    return i51370417203f2995dc2c3d12892585081cbbe62d6fc6ab078d6ab9a032c583de.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i3fae8342aedac4a38054d7a1b7f39f1aba0b6d7dab4f6bc060223ed82b262e66.CancelRequestBuilder) {
    return i3fae8342aedac4a38054d7a1b7f39f1aba0b6d7dab4f6bc060223ed82b262e66.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i0a6857a29775f3527ade48d34c6b10728a71af6dd78f6079d11d7a8c550d33d0.DeclineRequestBuilder) {
    return i0a6857a29775f3527ade48d34c6b10728a71af6dd78f6079d11d7a8c550d33d0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*ib9679d22124b56eee2727407ed877bdcb26037fbd160079af337cae5fb4d0daa.DismissReminderRequestBuilder) {
    return ib9679d22124b56eee2727407ed877bdcb26037fbd160079af337cae5fb4d0daa.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*id538d185b63affdf922df6fd539fc310753e12da2cf320d1b88f99a274eb630c.ExtensionsRequestBuilder) {
    return id538d185b63affdf922df6fd539fc310753e12da2cf320d1b88f99a274eb630c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i57e38ffd72170d9573806545b3dcb9b944291909edbb83dba3cdb0286be63625.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i57e38ffd72170d9573806545b3dcb9b944291909edbb83dba3cdb0286be63625.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i2b2f3d27aa9083c49b8996cc9e960d32e3e186f3a4d0747f2f5d7f7a012876f9.ForwardRequestBuilder) {
    return i2b2f3d27aa9083c49b8996cc9e960d32e3e186f3a4d0747f2f5d7f7a012876f9.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*ifebec4f242fc2fe3f2ea086e8d1cbf3af35e7102fe2ffa1cb26ae92db66cf2fe.InstancesRequestBuilder) {
    return ifebec4f242fc2fe3f2ea086e8d1cbf3af35e7102fe2ffa1cb26ae92db66cf2fe.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i74f629d6bb0be831d2a137d49ec5069ad5842bccb2e79a63092c5a89772f6977.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i74f629d6bb0be831d2a137d49ec5069ad5842bccb2e79a63092c5a89772f6977.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i2d5ec4a49b750fc5ad295afe48c94a5ab6cef5a9fc91371d7fb50fa0313d3d36.MultiValueExtendedPropertiesRequestBuilder) {
    return i2d5ec4a49b750fc5ad295afe48c94a5ab6cef5a9fc91371d7fb50fa0313d3d36.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i67a2dce04b02ebfce753343824b3c5282d62c261c455c73f23e7cea86b9d9ac5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i67a2dce04b02ebfce753343824b3c5282d62c261c455c73f23e7cea86b9d9ac5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i600d2d25136676bc3455d985eb1fcb7cba8c523738d73efc6515bbcbbd5f81c4.SingleValueExtendedPropertiesRequestBuilder) {
    return i600d2d25136676bc3455d985eb1fcb7cba8c523738d73efc6515bbcbbd5f81c4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5d87c261930d44d33fc6debc682d69d91880d07c4002721d4eeaa087a2285b13.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i5d87c261930d44d33fc6debc682d69d91880d07c4002721d4eeaa087a2285b13.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i2fa86fdbf8669734c4ff4baec8a07f4167e970aa9cea188b2a558f8575787f5e.SnoozeReminderRequestBuilder) {
    return i2fa86fdbf8669734c4ff4baec8a07f4167e970aa9cea188b2a558f8575787f5e.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ib043c85c14658a9046e5ec7b35d56973a9523498c18c990c3d2c449ece96a914.TentativelyAcceptRequestBuilder) {
    return ib043c85c14658a9046e5ec7b35d56973a9523498c18c990c3d2c449ece96a914.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
