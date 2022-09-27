package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i189cf07452004a59fdc9a312b85613cf42e674b1268f40892744850e80393bfa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i18f5bf7222ff32fb24ccc8f6d32737d602e9f7bc09f143b35f9d132c2a2021aa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i245482dd46f7d17bb831b43d6b1ad96ff3e81020caf86ffb9b6d599037540c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i30d9d0230e5d2db706ed53b91c53ae308e2b8358c9cd7c22d46c1c29877c1af0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i33447d1fb62e68bdd02fc6c549dfff416b77c7baeee5cb43d7ecc14cb29557b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i3c17378b532dfd56069cdb9f5ea5ea3b8978bbd5e503ae433d1354f572cecad9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i603d786505660fcad5c74082c44a538a12268bbe066125af061fc34e8fbf6311 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i83c05d9ac07f820cfa6d1ad0162434a1c6dd103ab5066bcab1cd38cb219140d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    iaac5fefd553b7432398b65c83dc236c81f390dff713b631a9a2ccd7de84936d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    id09cdc1412d932b6e8acaff79d58551587c44cdbab5dc27b7a61f10007f4ff85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    idc76f6a8ee03d34512ffad53b611cabbd8b9015cc991aae550d8697bde22588e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    if75001e11938f08c700df05d71ae477ea3cde9a9c6cf9d17a91b35a55ab4261d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i0bf2d76aebec29acdb979979d4b8a74c9a8352845eb6dd72584313051773f121 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i3288dde21fb1b1441b04dff2b72ba0b3a4daf0269a8bde7c09b26d780dc938e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    ia3428c7cb98f457b22c0c9745b715d96ec19ce770aed036ac37ac456f2e90bdf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ida2e6f17dbd1a69a5e2c6bc9d98e7ffc20ce68a999e0c960b044302c74ccd843 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*iaac5fefd553b7432398b65c83dc236c81f390dff713b631a9a2ccd7de84936d5.AcceptRequestBuilder) {
    return iaac5fefd553b7432398b65c83dc236c81f390dff713b631a9a2ccd7de84936d5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i603d786505660fcad5c74082c44a538a12268bbe066125af061fc34e8fbf6311.AttachmentsRequestBuilder) {
    return i603d786505660fcad5c74082c44a538a12268bbe066125af061fc34e8fbf6311.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ida2e6f17dbd1a69a5e2c6bc9d98e7ffc20ce68a999e0c960b044302c74ccd843.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ida2e6f17dbd1a69a5e2c6bc9d98e7ffc20ce68a999e0c960b044302c74ccd843.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*idc76f6a8ee03d34512ffad53b611cabbd8b9015cc991aae550d8697bde22588e.CalendarRequestBuilder) {
    return idc76f6a8ee03d34512ffad53b611cabbd8b9015cc991aae550d8697bde22588e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i33447d1fb62e68bdd02fc6c549dfff416b77c7baeee5cb43d7ecc14cb29557b0.CancelRequestBuilder) {
    return i33447d1fb62e68bdd02fc6c549dfff416b77c7baeee5cb43d7ecc14cb29557b0.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i30d9d0230e5d2db706ed53b91c53ae308e2b8358c9cd7c22d46c1c29877c1af0.DeclineRequestBuilder) {
    return i30d9d0230e5d2db706ed53b91c53ae308e2b8358c9cd7c22d46c1c29877c1af0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i245482dd46f7d17bb831b43d6b1ad96ff3e81020caf86ffb9b6d599037540c0f.DismissReminderRequestBuilder) {
    return i245482dd46f7d17bb831b43d6b1ad96ff3e81020caf86ffb9b6d599037540c0f.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i18f5bf7222ff32fb24ccc8f6d32737d602e9f7bc09f143b35f9d132c2a2021aa.ExtensionsRequestBuilder) {
    return i18f5bf7222ff32fb24ccc8f6d32737d602e9f7bc09f143b35f9d132c2a2021aa.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i3288dde21fb1b1441b04dff2b72ba0b3a4daf0269a8bde7c09b26d780dc938e9.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i3288dde21fb1b1441b04dff2b72ba0b3a4daf0269a8bde7c09b26d780dc938e9.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i189cf07452004a59fdc9a312b85613cf42e674b1268f40892744850e80393bfa.ForwardRequestBuilder) {
    return i189cf07452004a59fdc9a312b85613cf42e674b1268f40892744850e80393bfa.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*if75001e11938f08c700df05d71ae477ea3cde9a9c6cf9d17a91b35a55ab4261d.MultiValueExtendedPropertiesRequestBuilder) {
    return if75001e11938f08c700df05d71ae477ea3cde9a9c6cf9d17a91b35a55ab4261d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0bf2d76aebec29acdb979979d4b8a74c9a8352845eb6dd72584313051773f121.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i0bf2d76aebec29acdb979979d4b8a74c9a8352845eb6dd72584313051773f121.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i83c05d9ac07f820cfa6d1ad0162434a1c6dd103ab5066bcab1cd38cb219140d9.SingleValueExtendedPropertiesRequestBuilder) {
    return i83c05d9ac07f820cfa6d1ad0162434a1c6dd103ab5066bcab1cd38cb219140d9.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia3428c7cb98f457b22c0c9745b715d96ec19ce770aed036ac37ac456f2e90bdf.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ia3428c7cb98f457b22c0c9745b715d96ec19ce770aed036ac37ac456f2e90bdf.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3c17378b532dfd56069cdb9f5ea5ea3b8978bbd5e503ae433d1354f572cecad9.SnoozeReminderRequestBuilder) {
    return i3c17378b532dfd56069cdb9f5ea5ea3b8978bbd5e503ae433d1354f572cecad9.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*id09cdc1412d932b6e8acaff79d58551587c44cdbab5dc27b7a61f10007f4ff85.TentativelyAcceptRequestBuilder) {
    return id09cdc1412d932b6e8acaff79d58551587c44cdbab5dc27b7a61f10007f4ff85.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
