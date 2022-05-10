package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i08c38a7614e628d6eb73c62c2cf91ba60191d565c6929a5692197c5e0061406d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i10971b307bd2b45c1674e96c2561abade70e29dcaafcc43b3d11a5429e4690e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i25824d8895850b9479382eb0e0d9bd5f3b5c19f41681cf3de89d0796c36f106c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i3201be0e73ff8adbc694ecc5f656d3126dfb8c6616a504161599447efea5492a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i36193968a4070aad107015950395cd4dad04b8881a9740dcba517f850f18b0e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i4c5cb5a0c84bc5b195db548dd48ccd72b1d0bc3ead10726f67aed21922cf058f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i8fb64369143681c94f4298ec9d3a575828a1928fcd19ccdc559dd557ee036c6d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    id3b58c6ca30fbbb4f50e8afd90395632bfd64558500020858bab064aa66313db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    if1c2017c4385cc72a48f5c653066f42b27d05e128855b96ec70f4cabbd836ec4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    if435cdb9b8a88b14c6f3f4bd40285b7e962b3052b74684810723dc8ae3f161ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    if895636c1f6ed4402e3fecd7ba57af4aa5dc239fc670fbbf5dee12e34a298167 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    if9658a7f46e29259cd8a60aba6426a6009ef6765ec068c1725cd8c7f2cbd7b2f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i18b1a627488fb9d8678ad5c79c5626038c6a77996d402949e9bb61da32fb684c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    ib257529959f0c2cc3771287de99e32489d3e61fb356fab908f22f2dc03993f2e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    ic0e266cd92b2fd2eaa4fb551d1919c9124ff010f1f94aab44798a8f73b55b0ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ida47a625620035b4d4c8f6e43033d5ef9fb4897fd032d5a14a899289980d0991 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*if435cdb9b8a88b14c6f3f4bd40285b7e962b3052b74684810723dc8ae3f161ff.AcceptRequestBuilder) {
    return if435cdb9b8a88b14c6f3f4bd40285b7e962b3052b74684810723dc8ae3f161ff.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i4c5cb5a0c84bc5b195db548dd48ccd72b1d0bc3ead10726f67aed21922cf058f.AttachmentsRequestBuilder) {
    return i4c5cb5a0c84bc5b195db548dd48ccd72b1d0bc3ead10726f67aed21922cf058f.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i18b1a627488fb9d8678ad5c79c5626038c6a77996d402949e9bb61da32fb684c.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i18b1a627488fb9d8678ad5c79c5626038c6a77996d402949e9bb61da32fb684c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i08c38a7614e628d6eb73c62c2cf91ba60191d565c6929a5692197c5e0061406d.CalendarRequestBuilder) {
    return i08c38a7614e628d6eb73c62c2cf91ba60191d565c6929a5692197c5e0061406d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*id3b58c6ca30fbbb4f50e8afd90395632bfd64558500020858bab064aa66313db.CancelRequestBuilder) {
    return id3b58c6ca30fbbb4f50e8afd90395632bfd64558500020858bab064aa66313db.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i8fb64369143681c94f4298ec9d3a575828a1928fcd19ccdc559dd557ee036c6d.DeclineRequestBuilder) {
    return i8fb64369143681c94f4298ec9d3a575828a1928fcd19ccdc559dd557ee036c6d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i25824d8895850b9479382eb0e0d9bd5f3b5c19f41681cf3de89d0796c36f106c.DismissReminderRequestBuilder) {
    return i25824d8895850b9479382eb0e0d9bd5f3b5c19f41681cf3de89d0796c36f106c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i10971b307bd2b45c1674e96c2561abade70e29dcaafcc43b3d11a5429e4690e1.ExtensionsRequestBuilder) {
    return i10971b307bd2b45c1674e96c2561abade70e29dcaafcc43b3d11a5429e4690e1.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib257529959f0c2cc3771287de99e32489d3e61fb356fab908f22f2dc03993f2e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ib257529959f0c2cc3771287de99e32489d3e61fb356fab908f22f2dc03993f2e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*if895636c1f6ed4402e3fecd7ba57af4aa5dc239fc670fbbf5dee12e34a298167.ForwardRequestBuilder) {
    return if895636c1f6ed4402e3fecd7ba57af4aa5dc239fc670fbbf5dee12e34a298167.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i36193968a4070aad107015950395cd4dad04b8881a9740dcba517f850f18b0e6.MultiValueExtendedPropertiesRequestBuilder) {
    return i36193968a4070aad107015950395cd4dad04b8881a9740dcba517f850f18b0e6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic0e266cd92b2fd2eaa4fb551d1919c9124ff010f1f94aab44798a8f73b55b0ed.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ic0e266cd92b2fd2eaa4fb551d1919c9124ff010f1f94aab44798a8f73b55b0ed.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i3201be0e73ff8adbc694ecc5f656d3126dfb8c6616a504161599447efea5492a.SingleValueExtendedPropertiesRequestBuilder) {
    return i3201be0e73ff8adbc694ecc5f656d3126dfb8c6616a504161599447efea5492a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ida47a625620035b4d4c8f6e43033d5ef9fb4897fd032d5a14a899289980d0991.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ida47a625620035b4d4c8f6e43033d5ef9fb4897fd032d5a14a899289980d0991.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*if1c2017c4385cc72a48f5c653066f42b27d05e128855b96ec70f4cabbd836ec4.SnoozeReminderRequestBuilder) {
    return if1c2017c4385cc72a48f5c653066f42b27d05e128855b96ec70f4cabbd836ec4.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*if9658a7f46e29259cd8a60aba6426a6009ef6765ec068c1725cd8c7f2cbd7b2f.TentativelyAcceptRequestBuilder) {
    return if9658a7f46e29259cd8a60aba6426a6009ef6765ec068c1725cd8c7f2cbd7b2f.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
