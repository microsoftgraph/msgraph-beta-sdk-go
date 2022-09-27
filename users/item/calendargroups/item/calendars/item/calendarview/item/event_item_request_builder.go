package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1d7efb0cf0e538f716cbf4ec58ffbd8b2627f35a80b5cc51ad1a719ccda63410 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/tentativelyaccept"
    i594f031e77c7ff1deebf416d7fe1c8b3e48a3898deb3d09b0a51f16b899b9999 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/extensions"
    i5b406bcd2ab417fc4dd03bec5345f447db84b34a3739079d6b77141d501cad42 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/dismissreminder"
    i71a1dfd1790ad38215efe86cb847a7e9aa5b0e4485774ae264b132be08548500 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/snoozereminder"
    i761da5c466c608a108c783f66ae6a90778938de7cac958956283379dfa5f5b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/decline"
    i79f6e8eea7c7b3e6a2bce802b9a7f0a342ccafa8bb76d03da9284afa5af979ba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/attachments"
    i84e35fc82945374b0b4e6d6b60b9a283b295d752e621fa432c44c539b7457a13 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances"
    i95be0951230a29aac635ed6fe9a25887c5bbd1fa02bf4ea7df08ee54d2b119c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/calendar"
    i96f018cc33130e2eb44ba1f382328043fb9aa8cf4cbb32e98de1271009bcb638 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/cancel"
    ia0356016284b401708cc71750c01506b894366ee232c9f3a246c7a5f66855d00 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/accept"
    iba617898bd91d23a58685d21f4d6bd3cc9e930cf5709a670c8cbf8d09826572c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties"
    ic2c0e2ed3b1c4a6edf0119f14ebfd9619d42bb1ec48bfc2d02573e9892f75d7f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences"
    idb15d9c63943dfdd77189c4bc501947f713932b0b8a26a41ec57a1a9ed91bd3f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/forward"
    if32e2baabff5c88751154ffb2c522378538c97134301e17d7a7a463932a40dc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    i18e39b2135201cee257b03e8127c40e270b8ac9c6928d3b1408ca41554f5736f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item"
    i223f69849eb07c52dccbfbd8ef59d3b734f676c9d212707b26943e8a446dea29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/instances/item"
    i39da18ddc9fe463d0b5beda214a3ca0a1bbdc9a87511170acbce67edc1dd7aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/extensions/item"
    i7de345c43b4c56542c252f9c3b2494b65cf9d8a173f7713ee3bdee0b5aa8f6e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    ide91c535cba1a01dae33b9ba11703c6da99e6cb97eefde9dbed88118af7df3a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
    ie52936233179f5ad76614301cc056362325b5238f0fd7c3078dcce062a6ebadd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*ia0356016284b401708cc71750c01506b894366ee232c9f3a246c7a5f66855d00.AcceptRequestBuilder) {
    return ia0356016284b401708cc71750c01506b894366ee232c9f3a246c7a5f66855d00.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i79f6e8eea7c7b3e6a2bce802b9a7f0a342ccafa8bb76d03da9284afa5af979ba.AttachmentsRequestBuilder) {
    return i79f6e8eea7c7b3e6a2bce802b9a7f0a342ccafa8bb76d03da9284afa5af979ba.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ie52936233179f5ad76614301cc056362325b5238f0fd7c3078dcce062a6ebadd.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ie52936233179f5ad76614301cc056362325b5238f0fd7c3078dcce062a6ebadd.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i95be0951230a29aac635ed6fe9a25887c5bbd1fa02bf4ea7df08ee54d2b119c0.CalendarRequestBuilder) {
    return i95be0951230a29aac635ed6fe9a25887c5bbd1fa02bf4ea7df08ee54d2b119c0.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i96f018cc33130e2eb44ba1f382328043fb9aa8cf4cbb32e98de1271009bcb638.CancelRequestBuilder) {
    return i96f018cc33130e2eb44ba1f382328043fb9aa8cf4cbb32e98de1271009bcb638.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}{?%24select}";
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Decline()(*i761da5c466c608a108c783f66ae6a90778938de7cac958956283379dfa5f5b3c.DeclineRequestBuilder) {
    return i761da5c466c608a108c783f66ae6a90778938de7cac958956283379dfa5f5b3c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i5b406bcd2ab417fc4dd03bec5345f447db84b34a3739079d6b77141d501cad42.DismissReminderRequestBuilder) {
    return i5b406bcd2ab417fc4dd03bec5345f447db84b34a3739079d6b77141d501cad42.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ic2c0e2ed3b1c4a6edf0119f14ebfd9619d42bb1ec48bfc2d02573e9892f75d7f.ExceptionOccurrencesRequestBuilder) {
    return ic2c0e2ed3b1c4a6edf0119f14ebfd9619d42bb1ec48bfc2d02573e9892f75d7f.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i18e39b2135201cee257b03e8127c40e270b8ac9c6928d3b1408ca41554f5736f.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i18e39b2135201cee257b03e8127c40e270b8ac9c6928d3b1408ca41554f5736f.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i594f031e77c7ff1deebf416d7fe1c8b3e48a3898deb3d09b0a51f16b899b9999.ExtensionsRequestBuilder) {
    return i594f031e77c7ff1deebf416d7fe1c8b3e48a3898deb3d09b0a51f16b899b9999.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i39da18ddc9fe463d0b5beda214a3ca0a1bbdc9a87511170acbce67edc1dd7aa3.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i39da18ddc9fe463d0b5beda214a3ca0a1bbdc9a87511170acbce67edc1dd7aa3.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*idb15d9c63943dfdd77189c4bc501947f713932b0b8a26a41ec57a1a9ed91bd3f.ForwardRequestBuilder) {
    return idb15d9c63943dfdd77189c4bc501947f713932b0b8a26a41ec57a1a9ed91bd3f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*i84e35fc82945374b0b4e6d6b60b9a283b295d752e621fa432c44c539b7457a13.InstancesRequestBuilder) {
    return i84e35fc82945374b0b4e6d6b60b9a283b295d752e621fa432c44c539b7457a13.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i223f69849eb07c52dccbfbd8ef59d3b734f676c9d212707b26943e8a446dea29.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i223f69849eb07c52dccbfbd8ef59d3b734f676c9d212707b26943e8a446dea29.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iba617898bd91d23a58685d21f4d6bd3cc9e930cf5709a670c8cbf8d09826572c.MultiValueExtendedPropertiesRequestBuilder) {
    return iba617898bd91d23a58685d21f4d6bd3cc9e930cf5709a670c8cbf8d09826572c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7de345c43b4c56542c252f9c3b2494b65cf9d8a173f7713ee3bdee0b5aa8f6e1.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i7de345c43b4c56542c252f9c3b2494b65cf9d8a173f7713ee3bdee0b5aa8f6e1.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if32e2baabff5c88751154ffb2c522378538c97134301e17d7a7a463932a40dc0.SingleValueExtendedPropertiesRequestBuilder) {
    return if32e2baabff5c88751154ffb2c522378538c97134301e17d7a7a463932a40dc0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ide91c535cba1a01dae33b9ba11703c6da99e6cb97eefde9dbed88118af7df3a8.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ide91c535cba1a01dae33b9ba11703c6da99e6cb97eefde9dbed88118af7df3a8.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i71a1dfd1790ad38215efe86cb847a7e9aa5b0e4485774ae264b132be08548500.SnoozeReminderRequestBuilder) {
    return i71a1dfd1790ad38215efe86cb847a7e9aa5b0e4485774ae264b132be08548500.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i1d7efb0cf0e538f716cbf4ec58ffbd8b2627f35a80b5cc51ad1a719ccda63410.TentativelyAcceptRequestBuilder) {
    return i1d7efb0cf0e538f716cbf4ec58ffbd8b2627f35a80b5cc51ad1a719ccda63410.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
