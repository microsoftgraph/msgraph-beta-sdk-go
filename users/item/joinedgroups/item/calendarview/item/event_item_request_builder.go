package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i065986b8c5d8df4cf90720db4f296dc1a8085d310718e5c11d0c2d8a31ed8ab2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/multivalueextendedproperties"
    i1e2a85064dcaa31f5f2a3d1c44d6ad2012ede667d4d0d5c7b5acac2b0e8d5947 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/forward"
    i2afe22cc8a461a1d9b3587c577266d15f3a6166dd5dad68ea24f25ca956d03ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/accept"
    i48a9022ef079c23a69542192b5f5f4ab279ad2c97ccb6b0d5492744bfc6d7f5a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/cancel"
    i53e13c31f472c1ed281680a83af9d60cc11b2bd756c1ecb7e2e09cbfd7ab9bf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences"
    i5967a248b5184ad86c50051d57efb7b38c8a28a8e9be10135129976682ed9da6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/dismissreminder"
    i5a80a74d7db2e608145814f4d27d65ca2ce410d54297ec02523074599c0201c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances"
    i646690d084bb69b301583c45c20013a491f30406aaae0ee1df69573cee57cfeb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/calendar"
    i6d2d8db2c23b7ac6d419b67258c67ad33d110bb46825e2975ba3ad5298d43027 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/tentativelyaccept"
    ia5db82cc34be700c510d72aaf2653ce4880f0fb684f71bad186501d29ac63b52 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/decline"
    id1e9f51d567993ecf0dd729a04f8235a424da0a59e4b0d61b5eb8667d6c39f7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/singlevalueextendedproperties"
    id5142bed4eb4bd355390abca58f5da89d398c0b3dbd23fa1d05990febc368ab9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/extensions"
    ie8841531914f8cb75d2301c9d628aca846bc55c546b2b067a434b51dcc1653e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/attachments"
    iec65631fd6fb0a3a9f77fa9ed83a89920471afc4f2df0310d4940db54fa20613 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/snoozereminder"
    i2d3ce52d36c111564342f43ce0a66923cbd01b873167eab39af6ebb938486b30 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/multivalueextendedproperties/item"
    i3276f61a687aea842f66bf35f5aca8227e2fac398c456252ac0406cca4892e69 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/extensions/item"
    i540e03337568617035b245fc00de4eb3144909562d376b3baac152e7fbd66f0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item"
    i691e973b6c96673a7b3428b60eb4acad694b42fc79a9ff564a9f92242425547b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item"
    idcb89e9103645f752e5648b8b598ef2d3b01a806c43843433196a120f6c09d8e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/singlevalueextendedproperties/item"
    iebd60e21229f7c7aa27d83cc2af07579adbc09555db1f44d6d7b31e21aeda4b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.group entity.
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
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*i2afe22cc8a461a1d9b3587c577266d15f3a6166dd5dad68ea24f25ca956d03ef.AcceptRequestBuilder) {
    return i2afe22cc8a461a1d9b3587c577266d15f3a6166dd5dad68ea24f25ca956d03ef.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ie8841531914f8cb75d2301c9d628aca846bc55c546b2b067a434b51dcc1653e0.AttachmentsRequestBuilder) {
    return ie8841531914f8cb75d2301c9d628aca846bc55c546b2b067a434b51dcc1653e0.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*iebd60e21229f7c7aa27d83cc2af07579adbc09555db1f44d6d7b31e21aeda4b8.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return iebd60e21229f7c7aa27d83cc2af07579adbc09555db1f44d6d7b31e21aeda4b8.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i646690d084bb69b301583c45c20013a491f30406aaae0ee1df69573cee57cfeb.CalendarRequestBuilder) {
    return i646690d084bb69b301583c45c20013a491f30406aaae0ee1df69573cee57cfeb.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i48a9022ef079c23a69542192b5f5f4ab279ad2c97ccb6b0d5492744bfc6d7f5a.CancelRequestBuilder) {
    return i48a9022ef079c23a69542192b5f5f4ab279ad2c97ccb6b0d5492744bfc6d7f5a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendarView/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendarView for users
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
// CreateGetRequestInformation the calendar view for the calendar. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the calendar view for the calendar. Read-only.
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
// CreatePatchRequestInformation update the navigation property calendarView in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendarView in users
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
func (m *EventItemRequestBuilder) Decline()(*ia5db82cc34be700c510d72aaf2653ce4880f0fb684f71bad186501d29ac63b52.DeclineRequestBuilder) {
    return ia5db82cc34be700c510d72aaf2653ce4880f0fb684f71bad186501d29ac63b52.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property calendarView for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i5967a248b5184ad86c50051d57efb7b38c8a28a8e9be10135129976682ed9da6.DismissReminderRequestBuilder) {
    return i5967a248b5184ad86c50051d57efb7b38c8a28a8e9be10135129976682ed9da6.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i53e13c31f472c1ed281680a83af9d60cc11b2bd756c1ecb7e2e09cbfd7ab9bf5.ExceptionOccurrencesRequestBuilder) {
    return i53e13c31f472c1ed281680a83af9d60cc11b2bd756c1ecb7e2e09cbfd7ab9bf5.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i691e973b6c96673a7b3428b60eb4acad694b42fc79a9ff564a9f92242425547b.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i691e973b6c96673a7b3428b60eb4acad694b42fc79a9ff564a9f92242425547b.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*id5142bed4eb4bd355390abca58f5da89d398c0b3dbd23fa1d05990febc368ab9.ExtensionsRequestBuilder) {
    return id5142bed4eb4bd355390abca58f5da89d398c0b3dbd23fa1d05990febc368ab9.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i3276f61a687aea842f66bf35f5aca8227e2fac398c456252ac0406cca4892e69.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i3276f61a687aea842f66bf35f5aca8227e2fac398c456252ac0406cca4892e69.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i1e2a85064dcaa31f5f2a3d1c44d6ad2012ede667d4d0d5c7b5acac2b0e8d5947.ForwardRequestBuilder) {
    return i1e2a85064dcaa31f5f2a3d1c44d6ad2012ede667d4d0d5c7b5acac2b0e8d5947.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Read-only.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the calendar view for the calendar. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*i5a80a74d7db2e608145814f4d27d65ca2ce410d54297ec02523074599c0201c2.InstancesRequestBuilder) {
    return i5a80a74d7db2e608145814f4d27d65ca2ce410d54297ec02523074599c0201c2.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i540e03337568617035b245fc00de4eb3144909562d376b3baac152e7fbd66f0e.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i540e03337568617035b245fc00de4eb3144909562d376b3baac152e7fbd66f0e.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i065986b8c5d8df4cf90720db4f296dc1a8085d310718e5c11d0c2d8a31ed8ab2.MultiValueExtendedPropertiesRequestBuilder) {
    return i065986b8c5d8df4cf90720db4f296dc1a8085d310718e5c11d0c2d8a31ed8ab2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2d3ce52d36c111564342f43ce0a66923cbd01b873167eab39af6ebb938486b30.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i2d3ce52d36c111564342f43ce0a66923cbd01b873167eab39af6ebb938486b30.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property calendarView in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*id1e9f51d567993ecf0dd729a04f8235a424da0a59e4b0d61b5eb8667d6c39f7c.SingleValueExtendedPropertiesRequestBuilder) {
    return id1e9f51d567993ecf0dd729a04f8235a424da0a59e4b0d61b5eb8667d6c39f7c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*idcb89e9103645f752e5648b8b598ef2d3b01a806c43843433196a120f6c09d8e.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return idcb89e9103645f752e5648b8b598ef2d3b01a806c43843433196a120f6c09d8e.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*iec65631fd6fb0a3a9f77fa9ed83a89920471afc4f2df0310d4940db54fa20613.SnoozeReminderRequestBuilder) {
    return iec65631fd6fb0a3a9f77fa9ed83a89920471afc4f2df0310d4940db54fa20613.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6d2d8db2c23b7ac6d419b67258c67ad33d110bb46825e2975ba3ad5298d43027.TentativelyAcceptRequestBuilder) {
    return i6d2d8db2c23b7ac6d419b67258c67ad33d110bb46825e2975ba3ad5298d43027.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
