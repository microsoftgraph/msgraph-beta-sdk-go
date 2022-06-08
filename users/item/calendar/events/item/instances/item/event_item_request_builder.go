package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b36e7fbc97a2d8370891893c3c221c35d5170abe34df872e61110332d277b91 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/singlevalueextendedproperties"
    i1b9574753683cc8d202bd6ebe4d47660d0c4af6c9b6e916baae1434101fdc2cf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/extensions"
    i1f5e3d65a5cf771aca552b386199739fe6316961b86a8d245c7ee77e728e3042 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/calendar"
    i2c5a6d255f388fbe9f458474c3d7b8375b48a7a3427dcf464b1a2d6f3d8a0a5a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/multivalueextendedproperties"
    i59cd209bb367a3ab914363b6be8f5fb451216cf21050e0d5a896fa664fc34276 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/decline"
    i5b094fdc86d92954de5d4355ef2be265bc48f21e126689ede44ba2eee10350e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/attachments"
    i8e02c23b1d6b8b21ee7ec54c3b7d52509b13487c70ba47e4d17cd88e1f51cec4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/cancel"
    i99342ad7888c3b1b3de1ebe64b60c40c790ddd95c47e43ccecef6ef4862cfbf1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/forward"
    ia0961fac91912f701735f9967d33ee12b5c710fb6ded8ba0937d6628742ca96f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences"
    ib5416c6361187b3a360d9a310bfeae26f774901100a1d92c2d62e00cf601b6fe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/accept"
    ib9b9b411c5dd682c0d1ff5f4867c9cd6f857e0f0132e9460a7153ca0b4713af3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/dismissreminder"
    ic5a32a2b62a0bf106cf3425ad8844f8aff7ec2c750471a939a605ee262f4bce0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/tentativelyaccept"
    id825ca3df52f5a3f6c09527f1b9c03b850d3da5afd13f533f8d806dd8a94f968 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/snoozereminder"
    i28aca109418a72a7f76736bcf0260278de4d6157ed8f33f3f1c3c23942637431 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/singlevalueextendedproperties/item"
    i3a75008e9baf8acede4ca3413c6986388307a21f4074dba1b6f99b21d130eb68 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/multivalueextendedproperties/item"
    i89d4ea45299649e9b18492aaa78288497d03472775e3df598917d86ed072b972 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item"
    ia0eebe8027560a990309ae2f1973622881ab677fc092d3f9a9c72a48ae18c04b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/attachments/item"
    ied016554a0be1e56a06c14b4213fff5b2a35322c0b4bd752e9f5a25a7a40e738 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*ib5416c6361187b3a360d9a310bfeae26f774901100a1d92c2d62e00cf601b6fe.AcceptRequestBuilder) {
    return ib5416c6361187b3a360d9a310bfeae26f774901100a1d92c2d62e00cf601b6fe.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i5b094fdc86d92954de5d4355ef2be265bc48f21e126689ede44ba2eee10350e4.AttachmentsRequestBuilder) {
    return i5b094fdc86d92954de5d4355ef2be265bc48f21e126689ede44ba2eee10350e4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia0eebe8027560a990309ae2f1973622881ab677fc092d3f9a9c72a48ae18c04b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ia0eebe8027560a990309ae2f1973622881ab677fc092d3f9a9c72a48ae18c04b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i1f5e3d65a5cf771aca552b386199739fe6316961b86a8d245c7ee77e728e3042.CalendarRequestBuilder) {
    return i1f5e3d65a5cf771aca552b386199739fe6316961b86a8d245c7ee77e728e3042.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i8e02c23b1d6b8b21ee7ec54c3b7d52509b13487c70ba47e4d17cd88e1f51cec4.CancelRequestBuilder) {
    return i8e02c23b1d6b8b21ee7ec54c3b7d52509b13487c70ba47e4d17cd88e1f51cec4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i59cd209bb367a3ab914363b6be8f5fb451216cf21050e0d5a896fa664fc34276.DeclineRequestBuilder) {
    return i59cd209bb367a3ab914363b6be8f5fb451216cf21050e0d5a896fa664fc34276.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*ib9b9b411c5dd682c0d1ff5f4867c9cd6f857e0f0132e9460a7153ca0b4713af3.DismissReminderRequestBuilder) {
    return ib9b9b411c5dd682c0d1ff5f4867c9cd6f857e0f0132e9460a7153ca0b4713af3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ia0961fac91912f701735f9967d33ee12b5c710fb6ded8ba0937d6628742ca96f.ExceptionOccurrencesRequestBuilder) {
    return ia0961fac91912f701735f9967d33ee12b5c710fb6ded8ba0937d6628742ca96f.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i89d4ea45299649e9b18492aaa78288497d03472775e3df598917d86ed072b972.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i89d4ea45299649e9b18492aaa78288497d03472775e3df598917d86ed072b972.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i1b9574753683cc8d202bd6ebe4d47660d0c4af6c9b6e916baae1434101fdc2cf.ExtensionsRequestBuilder) {
    return i1b9574753683cc8d202bd6ebe4d47660d0c4af6c9b6e916baae1434101fdc2cf.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ied016554a0be1e56a06c14b4213fff5b2a35322c0b4bd752e9f5a25a7a40e738.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ied016554a0be1e56a06c14b4213fff5b2a35322c0b4bd752e9f5a25a7a40e738.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i99342ad7888c3b1b3de1ebe64b60c40c790ddd95c47e43ccecef6ef4862cfbf1.ForwardRequestBuilder) {
    return i99342ad7888c3b1b3de1ebe64b60c40c790ddd95c47e43ccecef6ef4862cfbf1.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i2c5a6d255f388fbe9f458474c3d7b8375b48a7a3427dcf464b1a2d6f3d8a0a5a.MultiValueExtendedPropertiesRequestBuilder) {
    return i2c5a6d255f388fbe9f458474c3d7b8375b48a7a3427dcf464b1a2d6f3d8a0a5a.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i3a75008e9baf8acede4ca3413c6986388307a21f4074dba1b6f99b21d130eb68.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i3a75008e9baf8acede4ca3413c6986388307a21f4074dba1b6f99b21d130eb68.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i0b36e7fbc97a2d8370891893c3c221c35d5170abe34df872e61110332d277b91.SingleValueExtendedPropertiesRequestBuilder) {
    return i0b36e7fbc97a2d8370891893c3c221c35d5170abe34df872e61110332d277b91.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i28aca109418a72a7f76736bcf0260278de4d6157ed8f33f3f1c3c23942637431.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i28aca109418a72a7f76736bcf0260278de4d6157ed8f33f3f1c3c23942637431.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*id825ca3df52f5a3f6c09527f1b9c03b850d3da5afd13f533f8d806dd8a94f968.SnoozeReminderRequestBuilder) {
    return id825ca3df52f5a3f6c09527f1b9c03b850d3da5afd13f533f8d806dd8a94f968.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ic5a32a2b62a0bf106cf3425ad8844f8aff7ec2c750471a939a605ee262f4bce0.TentativelyAcceptRequestBuilder) {
    return ic5a32a2b62a0bf106cf3425ad8844f8aff7ec2c750471a939a605ee262f4bce0.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
