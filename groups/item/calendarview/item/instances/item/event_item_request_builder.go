package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0027058e0d423583daf072ee040468c478d8c0b4a73f27f5e8e3b9ddc8f0b8ef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/accept"
    i1ee292b4bfd12a5a08055138ca95fe8ed6add2074501c36b47296e7f5395245f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/singlevalueextendedproperties"
    i363571edd07fdc45dbb90546ad94b140da516bf5e2fd0dbaf7680e4c8707322c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/attachments"
    i3bf21c93f9b204d8479c6377f731bd400351887fe8d7d09975644ba521f5b23a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/dismissreminder"
    i5c50128020109de077f303070b1cac0deb045542050a39b2fa16819b453d96e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/cancel"
    i71980c447d3f65db94f98311252461c988f1732c4ea524944a51d04bd11201c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/extensions"
    i8ab054005cfc72f5f53ccbe56e16e607855b409640461c91f1caf8374e96d5ce "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/decline"
    ia33fabf348fe94674450060f87ab70e22e2a1d37ddb7fec896d5e7927b45401e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/tentativelyaccept"
    id092f5929cbe8e80f499a8c21f338e37a241f22460530789c26cc88a3d5124e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences"
    idbcc7da3fb1085f0b67f61b46ef0465bcf59b18f59fa703c8a3a88d3ba55bc30 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/calendar"
    ie1eb08c9d3ef722744db5b679b78aa278799d737f8954921462a06075ddbc912 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/snoozereminder"
    ifc52a0f3ded1600cd5a892645e0a01d9eb828a226b3fc5c8219891078e6f5f21 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/forward"
    ifdacb98e6858687ab06fc69533dc5f97395b24db7ab6183a2e002991d680b7b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/multivalueextendedproperties"
    i5e40965db50dbf358b9015f0b8a9368a0541be1015434490eece4ef00e9f76a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item"
    i9851bf1eefcb7934e564b0788919ac8a4c95264839244901a4eb905091851d0b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/multivalueextendedproperties/item"
    i9a64bb2f39b4439fc46040622b1650edaf6903491944b03cbcbb71abda6232c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    ib78b9da451da716385cd805130aff8e089ea5fb4714ae6a61538391c2a6f41e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/extensions/item"
    id936d5cffe86d1e439b54635db6706e9bb870cec0df54698d0629c632a88c131 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/attachments/item"
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i0027058e0d423583daf072ee040468c478d8c0b4a73f27f5e8e3b9ddc8f0b8ef.AcceptRequestBuilder) {
    return i0027058e0d423583daf072ee040468c478d8c0b4a73f27f5e8e3b9ddc8f0b8ef.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i363571edd07fdc45dbb90546ad94b140da516bf5e2fd0dbaf7680e4c8707322c.AttachmentsRequestBuilder) {
    return i363571edd07fdc45dbb90546ad94b140da516bf5e2fd0dbaf7680e4c8707322c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*id936d5cffe86d1e439b54635db6706e9bb870cec0df54698d0629c632a88c131.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return id936d5cffe86d1e439b54635db6706e9bb870cec0df54698d0629c632a88c131.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*idbcc7da3fb1085f0b67f61b46ef0465bcf59b18f59fa703c8a3a88d3ba55bc30.CalendarRequestBuilder) {
    return idbcc7da3fb1085f0b67f61b46ef0465bcf59b18f59fa703c8a3a88d3ba55bc30.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i5c50128020109de077f303070b1cac0deb045542050a39b2fa16819b453d96e3.CancelRequestBuilder) {
    return i5c50128020109de077f303070b1cac0deb045542050a39b2fa16819b453d96e3.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i8ab054005cfc72f5f53ccbe56e16e607855b409640461c91f1caf8374e96d5ce.DeclineRequestBuilder) {
    return i8ab054005cfc72f5f53ccbe56e16e607855b409640461c91f1caf8374e96d5ce.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i3bf21c93f9b204d8479c6377f731bd400351887fe8d7d09975644ba521f5b23a.DismissReminderRequestBuilder) {
    return i3bf21c93f9b204d8479c6377f731bd400351887fe8d7d09975644ba521f5b23a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*id092f5929cbe8e80f499a8c21f338e37a241f22460530789c26cc88a3d5124e2.ExceptionOccurrencesRequestBuilder) {
    return id092f5929cbe8e80f499a8c21f338e37a241f22460530789c26cc88a3d5124e2.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i5e40965db50dbf358b9015f0b8a9368a0541be1015434490eece4ef00e9f76a1.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i5e40965db50dbf358b9015f0b8a9368a0541be1015434490eece4ef00e9f76a1.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i71980c447d3f65db94f98311252461c988f1732c4ea524944a51d04bd11201c1.ExtensionsRequestBuilder) {
    return i71980c447d3f65db94f98311252461c988f1732c4ea524944a51d04bd11201c1.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib78b9da451da716385cd805130aff8e089ea5fb4714ae6a61538391c2a6f41e8.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ib78b9da451da716385cd805130aff8e089ea5fb4714ae6a61538391c2a6f41e8.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ifc52a0f3ded1600cd5a892645e0a01d9eb828a226b3fc5c8219891078e6f5f21.ForwardRequestBuilder) {
    return ifc52a0f3ded1600cd5a892645e0a01d9eb828a226b3fc5c8219891078e6f5f21.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ifdacb98e6858687ab06fc69533dc5f97395b24db7ab6183a2e002991d680b7b5.MultiValueExtendedPropertiesRequestBuilder) {
    return ifdacb98e6858687ab06fc69533dc5f97395b24db7ab6183a2e002991d680b7b5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9851bf1eefcb7934e564b0788919ac8a4c95264839244901a4eb905091851d0b.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i9851bf1eefcb7934e564b0788919ac8a4c95264839244901a4eb905091851d0b.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i1ee292b4bfd12a5a08055138ca95fe8ed6add2074501c36b47296e7f5395245f.SingleValueExtendedPropertiesRequestBuilder) {
    return i1ee292b4bfd12a5a08055138ca95fe8ed6add2074501c36b47296e7f5395245f.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9a64bb2f39b4439fc46040622b1650edaf6903491944b03cbcbb71abda6232c9.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i9a64bb2f39b4439fc46040622b1650edaf6903491944b03cbcbb71abda6232c9.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie1eb08c9d3ef722744db5b679b78aa278799d737f8954921462a06075ddbc912.SnoozeReminderRequestBuilder) {
    return ie1eb08c9d3ef722744db5b679b78aa278799d737f8954921462a06075ddbc912.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ia33fabf348fe94674450060f87ab70e22e2a1d37ddb7fec896d5e7927b45401e.TentativelyAcceptRequestBuilder) {
    return ia33fabf348fe94674450060f87ab70e22e2a1d37ddb7fec896d5e7927b45401e.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
